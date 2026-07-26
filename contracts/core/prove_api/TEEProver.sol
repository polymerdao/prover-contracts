// SPDX-License-Identifier: Apache-2.0
/*
 * Copyright 2026, Polymer Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 */

pragma solidity 0.8.15;

import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {ReceiptParser} from "../../libs/ReceiptParser.sol";

/**
 * @title TEEProver
 * @notice Verifies cross-chain event proofs produced by Polymer's TEE-hosted
 *         payload-signer. Extends the V2 prover model with two properties:
 *
 *         1. Versioned prover registry. Each version of the enclave EIF has
 *            its own signing address; the contract stores `version ->
 *            (proverAddress, pcrsHash)`. Callers specify which version a
 *            proof should validate against, so multiple EIFs can coexist
 *            (overlap windows during rotation, per-chain dedicated provers,
 *            etc.) without any address-freeze migration.
 *
 *         2. Provenance enforcement. Every TEE-signed proof comes with a
 *            second signature committing to the set of RPC upstream domains
 *            (eTLD+1 — e.g. "alchemy.com", "quiknode.pro") whose responses
 *            backed the consensus result. Callers pass a `requiredUpstreams`
 *            allowlist and `blockedUpstreams` blocklist. The contract
 *            enforces both before falling through to IAVL membership.
 *
 *         On-chain enforcement moves the "which RPCs do I trust" policy out
 *         of the TEE itself. The enclave honestly reports witnesses; the
 *         contract (or any off-chain verifier) decides what to accept.
 */
contract TEEProver {
    using ECDSA for bytes32;

    // ------------------------------------------------------------------
    // Storage
    // ------------------------------------------------------------------

    struct ProverVersion {
        // Ethereum address derived from the TEE enclave's signing key. Both
        // the primary V1 sig and the provenance sig must recover to this.
        address proverAddress;
        // keccak256(pcr0 || pcr1 || pcr2) — the PCRs are 48 bytes each
        // (SHA-384), too wide to store directly, so we commit to a digest.
        // Full PCRs live off-chain in the EIF's run-describe.txt; auditors
        // recompute this hash to confirm binding.
        bytes32 pcrsHash;
        bool deprecated;
        bool registered;
    }

    mapping(uint256 => ProverVersion) public provers;

    address public owner;

    // Client type string (e.g. "proof_api") is part of the IAVL key the
    // enclave writes and the contract must reconstruct bit-for-bit.
    string public clientType;

    // Chain ID that goes into the signature's domain-separated hash. For
    // Polymer's peptide this is 901. Immutable — each version already has
    // its own address, and changing the signing chain ID would require
    // re-deploying the contract anyway.
    bytes32 public immutable SIGNING_CHAIN_ID;

    // ------------------------------------------------------------------
    // Events
    // ------------------------------------------------------------------

    event ProverRegistered(uint256 indexed version, address proverAddress, bytes32 pcrsHash);
    event ProverDeprecated(uint256 indexed version);
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    // ------------------------------------------------------------------
    // Errors
    // ------------------------------------------------------------------

    error NotOwner();
    error AlreadyRegistered();
    error UnknownProverVersion();
    error DeprecatedProverVersion();
    error InvalidProverSignature();
    error InvalidProvenanceSignature();
    error InvalidProvenanceSigLength();
    error DomainsNotSortedOrUnique();
    error MissingRequiredDomain();
    error BlockedDomainPresent();
    error InvalidProofRoot();

    // ------------------------------------------------------------------
    // Construction + admin
    // ------------------------------------------------------------------

    constructor(string memory clientType_, bytes32 signingChainId_, address initialOwner) {
        clientType = clientType_;
        SIGNING_CHAIN_ID = signingChainId_;
        owner = initialOwner;
        emit OwnershipTransferred(address(0), initialOwner);
    }

    modifier onlyOwner() {
        if (msg.sender != owner) revert NotOwner();
        _;
    }

    function transferOwnership(address newOwner) external onlyOwner {
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }

    /**
     * @notice Register a new prover version. Binds `version` to an attested
     *         enclave's signing address and a commitment to its PCR set.
     * @dev Idempotent-by-failure: re-registering the same version reverts
     *      rather than silently overwriting. To replace, deprecate first.
     */
    function registerProver(uint256 version, address proverAddress, bytes32 pcrsHash) external onlyOwner {
        if (provers[version].registered) revert AlreadyRegistered();
        provers[version] = ProverVersion({
            proverAddress: proverAddress,
            pcrsHash: pcrsHash,
            deprecated: false,
            registered: true
        });
        emit ProverRegistered(version, proverAddress, pcrsHash);
    }

    /**
     * @notice Mark a prover version as deprecated. validateEvent calls
     *         against it will revert. Used during EIF rollover — register
     *         the new version, then deprecate the old.
     */
    function deprecateProver(uint256 version) external onlyOwner {
        ProverVersion storage pv = provers[version];
        if (!pv.registered) revert UnknownProverVersion();
        pv.deprecated = true;
        emit ProverDeprecated(version);
    }

    // ------------------------------------------------------------------
    // Core validation
    // ------------------------------------------------------------------

    /**
     * @notice Validate a TEE-signed proof against a specific prover version,
     *         subject to caller-provided domain policy.
     *
     * @param proverVersion       Which registered enclave produced this proof.
     * @param proof               The V1-wire proof bytes (same layout as
     *                             CrossL2ProverV2 — contents are unchanged
     *                             from that format).
     * @param provenanceSignature 65-byte ECDSA signature by the same enclave
     *                             key, over the provenance-extended payload.
     * @param provenanceDomains   Sorted, strictly-ascending unique list of
     *                             eTLD+1 domains that witnessed the event.
     *                             Must match exactly what the enclave emitted.
     * @param requiredUpstreams   Caller's allowlist — every entry must
     *                             appear in `provenanceDomains`. Empty =
     *                             no requirement.
     * @param blockedUpstreams    Caller's denylist — no entry may appear
     *                             in `provenanceDomains`. Empty = no deny.
     *
     * @return chainId            The source chain ID of the event.
     * @return emittingContract   The contract address that emitted the event.
     * @return topics             ABI-encoded indexed topics.
     * @return unindexedData      The unindexed data of the event.
     */
    function validateEvent(
        uint256 proverVersion,
        bytes calldata proof,
        bytes calldata provenanceSignature,
        string[] calldata provenanceDomains,
        string[] calldata requiredUpstreams,
        string[] calldata blockedUpstreams
    ) external view returns (uint32 chainId, address emittingContract, bytes memory topics, bytes memory unindexedData) {
        address proverAddress = _lookupProver(proverVersion);

        // Domain policy first — cheapest to reject on.
        _enforceRequired(provenanceDomains, requiredUpstreams);
        _enforceBlocked(provenanceDomains, blockedUpstreams);

        chainId = uint32(bytes4(proof[97:101]));

        // Signature recovery moved into helpers so the top-level frame
        // doesn't overflow the EVM stack. Both must recover to the same
        // address bound to this prover version.
        _verifySignatures(
            proof, provenanceSignature, _hashProvenanceDomains(provenanceDomains), proverAddress
        );

        // IAVL membership — identical to CrossL2ProverV2.
        uint256 eventEnd = uint16(bytes2(proof[126:128]));
        bytes memory rawEvent = proof[128:eventEnd];
        this.verifyMembership(
            bytes32(proof[0:32]),
            ReceiptParser.eventRootKey(
                chainId,
                clientType,
                uint64(bytes8(proof[109:117])),
                uint32(bytes4(proof[117:121])),
                uint32(bytes4(proof[121:125]))
            ),
            keccak256(rawEvent),
            proof[eventEnd:]
        );

        (emittingContract, topics, unindexedData) = this.parseEvent(rawEvent, uint8(proof[125]));
    }

    /// Load + status-check a prover version. Reverts if unknown or deprecated.
    function _lookupProver(uint256 version) internal view returns (address) {
        ProverVersion memory pv = provers[version];
        if (!pv.registered) revert UnknownProverVersion();
        if (pv.deprecated) revert DeprecatedProverVersion();
        return pv.proverAddress;
    }

    /// Verify both the primary V1 signature and the provenance signature
    /// recover to `expected`. Extracted so validateEvent's stack stays shallow.
    function _verifySignatures(
        bytes calldata proof,
        bytes calldata provenanceSignature,
        bytes32 provHash,
        address expected
    ) internal view {
        bytes32 stateRoot = bytes32(proof[0:32]);
        uint64 peptideHeight = uint64(bytes8(proof[101:109]));

        if (
            ECDSA.recover(
                _primarySigningHash(stateRoot, peptideHeight),
                bytes.concat(bytes32(proof[32:64]), bytes32(proof[64:96]), bytes1(proof[96]))
            ) != expected
        ) {
            revert InvalidProverSignature();
        }

        if (provenanceSignature.length != 65) revert InvalidProvenanceSigLength();
        if (
            ECDSA.recover(
                _provenanceSigningHash(stateRoot, peptideHeight, provHash),
                bytes.concat(
                    bytes32(provenanceSignature[0:32]),
                    bytes32(provenanceSignature[32:64]),
                    bytes1(provenanceSignature[64])
                )
            ) != expected
        ) {
            revert InvalidProvenanceSignature();
        }
    }

    // ------------------------------------------------------------------
    // Hashing helpers
    // ------------------------------------------------------------------

    /// Matches the enclave's SignBlockV1 hashing:
    ///   keccak256( 0x00*32 || chainID || keccak256(stateRoot || peptideHeight) )
    function _primarySigningHash(bytes32 stateRoot, uint64 peptideHeight) internal view returns (bytes32) {
        return keccak256(
            bytes.concat(bytes32(0), SIGNING_CHAIN_ID, keccak256(abi.encodePacked(stateRoot, peptideHeight)))
        );
    }

    /// Same domain separation as above, but the inner payload also commits
    /// to the provenance hash:
    ///   keccak256( 0x00*32 || chainID || keccak256(stateRoot || peptideHeight || provHash) )
    function _provenanceSigningHash(bytes32 stateRoot, uint64 peptideHeight, bytes32 provHash)
        internal
        view
        returns (bytes32)
    {
        return keccak256(
            bytes.concat(
                bytes32(0), SIGNING_CHAIN_ID, keccak256(abi.encodePacked(stateRoot, peptideHeight, provHash))
            )
        );
    }

    /**
     * @notice Reconstruct the domain-set digest the enclave signed. The
     *         enclave emits the sorted unique domain list and hashes it as
     *         `keccak256(strings.Join(sorted_domains, "\n"))`. We reproduce
     *         that here, enforcing strict ascending order as a side-effect
     *         (prevents callers from reordering to change the hash).
     *
     *         Empty `domains` -> keccak256("") — still well-defined.
     */
    function _hashProvenanceDomains(string[] calldata domains) internal pure returns (bytes32) {
        if (domains.length == 0) return keccak256("");

        // Enforce strict ascending, no duplicates.
        for (uint256 i = 1; i < domains.length; ++i) {
            if (!_strictlyLess(domains[i - 1], domains[i])) revert DomainsNotSortedOrUnique();
        }

        bytes memory buf = bytes(domains[0]);
        for (uint256 i = 1; i < domains.length; ++i) {
            buf = bytes.concat(buf, "\n", bytes(domains[i]));
        }
        return keccak256(buf);
    }

    function _strictlyLess(string calldata a, string calldata b) internal pure returns (bool) {
        bytes memory ba = bytes(a);
        bytes memory bb = bytes(b);
        uint256 n = ba.length < bb.length ? ba.length : bb.length;
        for (uint256 i = 0; i < n; ++i) {
            if (ba[i] < bb[i]) return true;
            if (ba[i] > bb[i]) return false;
        }
        return ba.length < bb.length;
    }

    // ------------------------------------------------------------------
    // Domain policy helpers
    // ------------------------------------------------------------------

    /// Every element of `required` must be present in `domains`. domains
    /// is sorted, so membership is `_contains`.
    function _enforceRequired(string[] calldata domains, string[] calldata required) internal pure {
        for (uint256 i = 0; i < required.length; ++i) {
            if (!_contains(domains, required[i])) revert MissingRequiredDomain();
        }
    }

    /// No element of `blocked` may appear in `domains`.
    function _enforceBlocked(string[] calldata domains, string[] calldata blocked) internal pure {
        for (uint256 i = 0; i < blocked.length; ++i) {
            if (_contains(domains, blocked[i])) revert BlockedDomainPresent();
        }
    }

    function _contains(string[] calldata haystack, string calldata needle) internal pure returns (bool) {
        bytes32 h = keccak256(bytes(needle));
        for (uint256 i = 0; i < haystack.length; ++i) {
            if (keccak256(bytes(haystack[i])) == h) return true;
        }
        return false;
    }

    // ------------------------------------------------------------------
    // IAVL membership + event parsing (copied from CrossL2ProverV2 so this
    // contract is a drop-in replacement for the on-chain verifier)
    // ------------------------------------------------------------------

    function verifyMembership(bytes32 root, bytes memory key, bytes32 value, bytes calldata proof)
        public
        pure
        virtual
    {
        uint256 path0start = uint256(uint8(proof[1]));
        bytes32 prehash = sha256(abi.encodePacked(proof[2:path0start], key, hex"20", sha256(abi.encodePacked(value))));
        uint256 offset = path0start;

        for (uint256 i = 0; i < uint256(uint8(proof[0])); ++i) {
            uint256 suffixstart = uint256(uint8(proof[offset]));
            uint256 suffixend = uint256(uint8(proof[offset + 1]));
            prehash = sha256(
                abi.encodePacked(
                    proof[offset + 2:offset + suffixstart], prehash, proof[offset + suffixstart:offset + suffixend]
                )
            );
            offset = offset + suffixend;
        }

        if (prehash != root) revert InvalidProofRoot();
    }

    function parseEvent(bytes calldata rawEvent, uint8 numTopics)
        public
        pure
        virtual
        returns (address emittingContract, bytes memory topics, bytes memory unindexedData)
    {
        uint256 topicsEnd = 32 * numTopics + 20;
        return (address(bytes20(rawEvent[:20])), rawEvent[20:topicsEnd], rawEvent[topicsEnd:]);
    }
}
