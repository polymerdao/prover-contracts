// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {SigningBase} from "./utils/Signing.base.t.sol";
import {CrossL2ProverV2} from "../../contracts/core/prove_api/CrossL2ProverV2.sol";

// Fuzzes the exported entrypoints of CrossL2ProverV2 that take untrusted input.
// - validators (validateEvent, validateSolLogs, verifyMembership): arbitrary input must never be
//   accepted, and tampering with a genuine proof must be detected.
// - pure extractors (parseEvent, inspectLogIdentifier, inspectPolymerState): structural/layout
//   consistency and no unexpected panics.
// ping() (no args) and const/immutable getters have nothing to fuzz and are excluded.
contract CrossL2ProverV2FuzzTest is SigningBase {
    CrossL2ProverV2 crossProverV2;
    bytes validProof;
    bytes validSolProof;

    function setUp() public {
        crossProverV2 = new CrossL2ProverV2(
            "proof_api",
            0x8D3921B96A3815F403Fb3a4c7fF525969d16f9E0,
            0x0000000000000000000000000000000000000000000000000000000000000385
        );
        validProof = load_proof("/test/prove_api/payload/op-proof-v2.hex");
        validSolProof = load_proof("/test/prove_api/payload/solana-proof.hex");
    }

    // ---------------------------------------------------------------- validateEvent

    function testFuzz_validateEvent_arbitraryNeverValidates(bytes calldata proof) public view {
        try crossProverV2.validateEvent(proof) returns (uint32, address, bytes memory, bytes memory) {
            revert("arbitrary proof unexpectedly validated");
        } catch {} // expected: not signed by the sequencer / malformed
    }

    // The topics/data split (numTopics, byte 125) is not committed, so it is malleable. A correct
    // consumer pins the topic count it expects. This mimics that consumer and is EXHAUSTIVE over every
    // byte (the mutation space is only validProof.length, so this beats random sampling): for each
    // single-byte flip that still validates and matches the topic count, the full event
    // (chainId, emitter, topics, data) must be byte-identical to the genuine one.
    function test_validateEvent_everyByteMutationIsTamperEvident() public view {
        (uint32 cId, address addr, bytes memory topics, bytes memory data) = crossProverV2.validateEvent(validProof);
        uint256 expectedTopicsLen = topics.length; // the count a consumer independently expects

        for (uint256 i = 0; i < validProof.length; i++) {
            bytes memory mutated = bytes.concat(validProof);
            mutated[i] = ~mutated[i]; // flip every bit of byte i

            try crossProverV2.validateEvent(mutated) returns (
                uint32 cId2, address addr2, bytes memory t2, bytes memory d2
            ) {
                if (t2.length != expectedTopicsLen) continue; // consumer rejects the unexpected topic count
                assertTrue(
                    cId2 == cId && addr2 == addr && keccak256(t2) == keccak256(topics)
                        && keccak256(d2) == keccak256(data),
                    "tampered byte passed the topic-count check with a different payload"
                );
            } catch {} // expected: tamper detected
        }
    }

    // ---------------------------------------------------------------- validateSolLogs

    function testFuzz_validateSolLogs_arbitraryNeverValidates(bytes calldata proof) public view {
        try crossProverV2.validateSolLogs(proof) returns (uint32, bytes32, string[] memory) {
            revert("arbitrary solana proof unexpectedly validated");
        } catch {} // expected: not signed by the sequencer / malformed
    }

    // Exhaustive single-byte tamper-evidence: every flip either reverts or returns the identical
    // (chainId, programID, logs). Solana commits all returned data, so there is no malleable split.
    function test_validateSolLogs_everyByteMutationIsTamperEvident() public view {
        (uint32 cId, bytes32 prog, string[] memory logs) = crossProverV2.validateSolLogs(validSolProof);
        bytes32 logsHash = keccak256(abi.encode(logs));

        for (uint256 i = 0; i < validSolProof.length; i++) {
            bytes memory mutated = bytes.concat(validSolProof);
            mutated[i] = ~mutated[i]; // flip every bit of byte i

            try crossProverV2.validateSolLogs(mutated) returns (uint32 cId2, bytes32 prog2, string[] memory logs2) {
                assertTrue(
                    cId2 == cId && prog2 == prog && keccak256(abi.encode(logs2)) == logsHash,
                    "tampered byte produced different accepted solana logs"
                );
            } catch {} // expected: tamper detected
        }
    }

    // ---------------------------------------------------------------- verifyMembership

    // Membership cannot be forged from arbitrary input: with an attacker-chosen root that is
    // independent of the sha256-derived prehash, this must revert (matching a 256-bit hash by
    // chance is infeasible). Real safety comes from validateEvent passing the *signed* root here.
    function testFuzz_verifyMembership_arbitraryNeverProves(
        bytes32 root,
        bytes calldata key,
        bytes32 value,
        bytes calldata proof
    ) public view {
        try crossProverV2.verifyMembership(root, key, value, proof) {
            revert("arbitrary membership unexpectedly proven");
        } catch {} // expected: InvalidProofRoot or out-of-bounds
    }

    // ---------------------------------------------------------------- parseEvent

    // Scoped to 0-4 topics: the valid domain for an Ethereum log (1 signature + up to 3 indexed).
    // NOTE: parseEvent overflows for numTopics >= 8 (it computes `32 * numTopics` in uint8); that is
    // a separate contract bug, out of the honest input range exercised here.
    function testFuzz_parseEvent_structuralConsistency(bytes calldata rawEvent, uint8 numTopics) public {
        numTopics = uint8(bound(numTopics, 0, 4));
        uint256 topicsEnd = uint256(numTopics) * 32 + 20;

        if (rawEvent.length < topicsEnd) {
            vm.expectRevert();
            crossProverV2.parseEvent(rawEvent, numTopics);
            return;
        }

        (address emitter, bytes memory topics, bytes memory data) = crossProverV2.parseEvent(rawEvent, numTopics);
        assertEq(emitter, address(bytes20(rawEvent[:20])));
        assertEq(topics.length, uint256(numTopics) * 32);
        assertEq(data.length, rawEvent.length - topicsEnd);
    }

    // ---------------------------------------------------------------- inspectors (pure)
    // These are unsigned offset extractors, so fuzzing raw bytes only re-tests slicing. Instead we
    // mutate a genuine (full-length) proof and assert the inspector agrees with what validateEvent
    // actually validated — a real cross-function consistency property, with no panic on valid input.

    // Any mutation that still validates must report the same source chain validateEvent committed to.
    function testFuzz_inspectLogIdentifier_matchesValidatedChain(uint256 pos, uint8 mask) public view {
        bytes memory m = bytes.concat(validProof);
        pos = bound(pos, 0, m.length - 1);
        if (mask != 0) m[pos] = bytes1(uint8(m[pos]) ^ mask);

        (uint32 srcChain,,,) = crossProverV2.inspectLogIdentifier(m);
        try crossProverV2.validateEvent(m) returns (uint32 cId, address, bytes memory, bytes memory) {
            assertEq(cId, srcChain, "inspector chainId disagrees with validated chainId");
        } catch {}
    }

    // The signature commits to (stateRoot, height); any mutation that still validates cannot have
    // altered the state root or height the inspector reports.
    function testFuzz_inspectPolymerState_matchesValidatedState(uint256 pos, uint8 mask) public view {
        (bytes32 root0, uint64 height0,) = crossProverV2.inspectPolymerState(validProof);

        bytes memory m = bytes.concat(validProof);
        pos = bound(pos, 0, m.length - 1);
        if (mask != 0) m[pos] = bytes1(uint8(m[pos]) ^ mask);

        try crossProverV2.validateEvent(m) returns (uint32, address, bytes memory, bytes memory) {
            (bytes32 root1, uint64 height1,) = crossProverV2.inspectPolymerState(m);
            assertTrue(root1 == root0 && height1 == height0, "validated state root/height drifted");
        } catch {}
    }
}
