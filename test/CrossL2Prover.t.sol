// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import {SigningBase} from "test/utils/Signing.base.t.sol";
import {CrossL2Prover} from "contracts/core/CrossL2Prover.sol";
import {SequencerSignatureVerifier} from "contracts/core/SequencerSignatureVerifier.sol";
import {Ics23Proof, OpIcs23Proof} from "contracts/libs/ReceiptParser.sol";
import {ISignatureVerifier} from "contracts/interfaces/ISignatureVerifier.sol";
import {ICrossL2Prover} from "contracts/interfaces/ICrossL2Prover.sol";
import {MerkleTrie} from "optimism/libraries/trie/MerkleTrie.sol";
import {IAppStateVerifier} from "contracts/interfaces/IAppStateVerifier.sol";

// OpIcs23ProofPath represents a commitment path in an ICS23 proof, which consists of a commitment prefix and a
// suffix.
struct OpIcs23ProofPath {
    bytes prefix;
    bytes suffix;
}

// OpIcs23Proof represents an ICS23 proof
struct OldOpIcs23Proof {
    OpIcs23ProofPath[] path;
    bytes key;
    bytes value;
    bytes prefix;
}

struct OldIcs23Proof {
    OldOpIcs23Proof[] proof;
    uint256 height;
}

contract CrossL2ProverTest is SigningBase {
    using stdJson for string;

    CrossL2Prover crossProver;

    // Proof
    bytes proof;
    bytes peptideAppProofBytes;
    bytes[] receiptMMPTProof;
    bytes32 receiptRoot;
    uint64 eventHeight;
    string srcChainId;
    bytes receiptIndex;
    bytes rlpEncodedReceipt;
    bytes32 logHash;
    uint256 logIdx;

    function setUp() public {
        peptideAppHash = hex"ebf1f4c1903364fef5ed8c01dfae471e3cad3c521faa4b555b091797ad3715fd";
        crossProver = new CrossL2Prover(sigVerifier, "proof_api");

        OldIcs23Proof memory peptideAppProof;
        (proof) = load_proof("/test/payload/RLP/proof1.hex");
        (peptideAppProof, receiptMMPTProof, receiptRoot, eventHeight, srcChainId, receiptIndex) =
        abi.decode(proof, (OldIcs23Proof, bytes[], bytes32, uint64, string, bytes));

        rlpEncodedReceipt = load_bytes_from_hex("/test/payload/RLP/receipt1.hex");

        peptideAppProofBytes = abi.encode(peptideAppProof);
        store_peptide_apphash(peptideAppHash, address(crossProver), peptideAppProof.height);
    }

    // Happy path for CrossEventProver.updateClient()
    function test_clientUpdate_success() public {
        Ics23Proof memory iavlProof = abi.decode(peptideAppProofBytes, (Ics23Proof));
        bytes32 hashToSign = keccak256(
            bytes.concat(
                bytes32(0),
                PEPTIDE_CHAIN_ID,
                keccak256(bytes.concat(bytes32(iavlProof.height), peptideAppHash, l1BlockHash))
            )
        );

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(sequencerPkey, hashToSign);
        bytes memory signature = abi.encodePacked(r, s, v); // Annoyingly, v, r, s are in a different order than those

        crossProver.updateClient(bytes.concat(l1BlockHash, signature), iavlProof.height, uint256(peptideAppHash));
        assertEq(crossProver.getState(iavlProof.height), uint256(peptideAppHash));
    }

    // Test that a client update will fail if it doesn't have a valid sequencer signature
    function test_revert_invalidClentUpdateSignature() public {
        Ics23Proof memory iavlProof = abi.decode(peptideAppProofBytes, (Ics23Proof));
        bytes32 hashToSign = keccak256(
            bytes.concat(
                bytes32(0),
                PEPTIDE_CHAIN_ID,
                keccak256(bytes.concat(bytes32(iavlProof.height), peptideAppHash, l1BlockHash))
            )
        );

        (uint8 v1, bytes32 r1, bytes32 s1) = vm.sign(notSequencerPkey, hashToSign);
        bytes memory invalidSignature = abi.encodePacked(r1, s1, v1); // Annoingly, v, r, s are in a different order
            // than those

        vm.expectRevert(ISignatureVerifier.InvalidSequencerSignature.selector);
        crossProver.updateClient(
            abi.encodePacked(l1BlockHash, invalidSignature), iavlProof.height, uint256(peptideAppHash)
        );
        assertEq(crossProver.getState(iavlProof.height), 0);
    }

    // // Happy path for CrossEventProver.validateReceipt()
    function test_validate_receipt_success() public {
        OldIcs23Proof memory oldP = abi.decode(peptideAppProofBytes, (OldIcs23Proof));

        OpIcs23Proof[] memory newProof = new OpIcs23Proof[](oldP.proof.length);

        for (uint256 i = 0; i < oldP.proof.length; i++) {
            bytes[] memory proofPrefix = new bytes[](oldP.proof[i].path.length);
            bytes[] memory proofSuffix = new bytes[](oldP.proof[i].path.length);

            for (uint256 j = 0; j < oldP.proof[i].path.length; j++) {
                proofPrefix[j] = oldP.proof[i].path[j].prefix;
                proofSuffix[j] = oldP.proof[i].path[j].suffix;
            }

            newProof[i] = OpIcs23Proof({
                proofPrefix: proofPrefix,
                proofSuffix: proofSuffix,
                key: oldP.proof[i].key,
                value: oldP.proof[i].value,
                prefix: oldP.proof[i].prefix
            });
        }

        Ics23Proof memory iavlProof = Ics23Proof({proof: newProof, height: oldP.height});

        bytes memory newProofBytes = abi.encode(iavlProof, receiptMMPTProof, receiptRoot, eventHeight, srcChainId, receiptIndex);
        (string memory chainId, bytes memory RLPBytes) = crossProver.validateReceipt(newProofBytes);

        assertEq(chainId, "11155420", "Chain ID mismatch");
        assertEq(RLPBytes, rlpEncodedReceipt, "RLP encoded receipt mismatch");
        crossProver.validateEvent(0, newProofBytes);
    }

    // Happy path for CrossEventProver.validateEvent()
    function test_validate_event_success() public {
        (string memory chainId, address emittingContract, bytes[] memory topics, bytes memory unindexedData) =
            crossProver.validateEvent(0, proof);
        assertEq(chainId, "11155420", "Chain ID mismatch");
        assertEq(emittingContract, 0x464C8ec100F2F42fB4e42E07E203DA2324f9FC67);
    }

    // Test valid peptide proof but invalid MMPT receipt proof
    function test_revert_invalidReceiptProof() public {
        // Change a few bytes in the receipt proof to cause the validateReceipt to fail
        receiptMMPTProof[1][0] = 0x01;
        Ics23Proof memory peptideAppProof = abi.decode(peptideAppProofBytes, (Ics23Proof));
        proof = abi.encode(peptideAppProof, receiptMMPTProof, receiptRoot, eventHeight, srcChainId, receiptIndex);

        vm.expectRevert("RLPReader: decoded item type for list is not a list item");
        crossProver.validateReceipt(proof);

        vm.expectRevert("RLPReader: decoded item type for list is not a list item");
        crossProver.validateEvent(0, proof);
    }

    // Test revert to prove a peptide apphash which has been seen but which doesn't prove the MPT receipt root given
    function test_revert_invalidReceiptRoot() public {
        // Change a few bytes in the receipt proof to cause the validateReceipt to fail
        bytes32 invalidReceiptRoot = keccak256(bytes("invalidReceiptRoot"));
        Ics23Proof memory peptideAppProof = abi.decode(peptideAppProofBytes, (Ics23Proof));
        proof = abi.encode(peptideAppProof, receiptMMPTProof, invalidReceiptRoot, eventHeight, srcChainId, receiptIndex);

        vm.expectRevert(IAppStateVerifier.InvalidProofValue.selector);
        crossProver.validateReceipt(proof);

        vm.expectRevert(IAppStateVerifier.InvalidProofValue.selector);
        crossProver.validateEvent(0, proof);
    }

    // Test trying to prove with a non-existent peptideApphash that hasn't yet been seen
    function test_revert_nonexistingPeptideAppHash() public {
        Ics23Proof memory peptideProof = abi.decode(peptideAppProofBytes, (Ics23Proof));
        store_peptide_apphash(bytes32(uint256(0)), address(crossProver), peptideProof.height); // clear peptide app hash
            // at height

        vm.expectRevert(IAppStateVerifier.InvalidIbcStateProof.selector);
        crossProver.validateReceipt(proof);

        vm.expectRevert(IAppStateVerifier.InvalidIbcStateProof.selector);
        crossProver.validateEvent(0, proof);
    }
}
