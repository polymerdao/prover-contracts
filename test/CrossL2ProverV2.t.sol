// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import {SigningBase} from "test/utils/Signing.base.t.sol";
import {CrossL2Prover} from "contracts/core/CrossL2Prover.sol";
import {CrossL2ProverV2} from "contracts/core/CrossL2ProverV2.sol";
import {SequencerSignatureVerifier} from "contracts/core/SequencerSignatureVerifier.sol";
import {Ics23Proof, OpIcs23Proof} from "contracts/libs/ReceiptParser.sol";
import {ISignatureVerifier} from "contracts/interfaces/ISignatureVerifier.sol";
import {ICrossL2Prover} from "contracts/interfaces/ICrossL2Prover.sol";
import {MerkleTrie} from "optimism/libraries/trie/MerkleTrie.sol";
import {IAppStateVerifier} from "contracts/interfaces/IAppStateVerifier.sol";
import {Base64} from "@openzeppelin/contracts/utils/Base64.sol";

contract CrossL2ProverTest is SigningBase {
    using stdJson for string;

    CrossL2ProverV2 crossProverV2;
    bytes proof;

    function setUp() public {
        peptideAppHash = hex"f3198683334bc574225bbadc18a07ff3c7496c5c844cf54d7b54a4341c63f004";
        crossProverV2 = new CrossL2ProverV2(
            "proof_api",
            0x8D3921B96A3815F403Fb3a4c7fF525969d16f9E0,
            0x0000000000000000000000000000000000000000000000000000000000000385
        );
        proof = load_proof("/test/payload/op-proof-v2.hex");

        bytes32 signHash = keccak256(
            bytes.concat(
                bytes32(0),
                crossProverV2.CHAIN_ID(),
                bytes32(0x6a418705ca9f0fd2f41b467c86904648db83c61680a0924f35e83e2aa0505874)
            )
        );

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(sequencerPkey, signHash);
    }

    // Test that a client update will fail if it doesn't have a valid sequencer signature
    function test_revert_invalidClentUpdateSignatureV2() public {
        vm.skip(true);
    }

    function test_validate_event_and_getters() public {
        console2.log("calldata length", proof.length);
        string memory expected = vm.readFile(string.concat(rootDir, "/test/payload/op-event-v2-1.json"));

        proof = load_proof("/test/payload/op-proof-v2-1.hex");
        (uint32 chainId, address addr, bytes memory topics, bytes memory data) = crossProverV2.validateEvent(proof);

        uint32 expectedChainId = abi.decode(expected.parseRaw(".ChainID"), (uint32));
        // address expectedAddress = abi.decode(expected.parseRaw(".Address"), (address));
        assertEq(chainId, expectedChainId);
        assertEq(addr, abi.decode(expected.parseRaw(".Event.address"), (address)));
        bytes memory expectedTopics = abi.encodePacked(
            abi.decode(expected.parseRaw(".Event.topics[0]"), (bytes32)),
            abi.decode(expected.parseRaw(".Event.topics[1]"), (bytes32)),
            abi.decode(expected.parseRaw(".Event.topics[2]"), (bytes32)),
            abi.decode(expected.parseRaw(".Event.topics[3]"), (bytes32))
        );

        assertEq(topics, expectedTopics);
        assertEq(data, abi.encodePacked(abi.decode(expected.parseRaw(".Event.data"), (bytes32))));

        _checkInspectPolymerState(proof, expected);
        _checkInspectLogIndentifier(proof, expected);
    }

    function test_validate_receipt_new_2() public {
        console2.log("calldata length", proof.length);
        string memory expected = vm.readFile(string.concat(rootDir, "/test/payload/op-event-v2.json"));

        (uint32 chainId, address addr, bytes memory topics, bytes memory data) = crossProverV2.validateEvent(proof);

        assertEq(chainId, 84_532);
        assertEq(addr, abi.decode(expected.parseRaw(".address"), (address)));
        bytes memory expectedTopics = abi.encodePacked(
            abi.decode(expected.parseRaw(".topics[0]"), (bytes32)),
            abi.decode(expected.parseRaw(".topics[1]"), (bytes32)),
            abi.decode(expected.parseRaw(".topics[2]"), (bytes32)),
            abi.decode(expected.parseRaw(".topics[3]"), (bytes32))
        );
        assertEq(topics, expectedTopics);
        assertEq(data, abi.encodePacked(abi.decode(expected.parseRaw(".data"), (bytes32))));
    }

    // Test valid peptide proof but invalid event hash
    function test_revert_invalidEventBytes() public {}

    // Test valid peptide proof but invalid local log index
    function test_revert_invalidEventIndex() public {}

    // Test valid peptide proof but invalid receipt index
    function test_revert_invalidReceiptIndex() public {}

    // Test revert to prove a state apphash does not prove the hash
    function test_revert_invalidIavlProof() public {}

    // Test revert to prove an invalid prefix
    function test_revert_invalidPrefix() public {}

    // Test trying to prove with a non-existent peptideApphash that hasn't yet been seen
    function test_revert_chain_ID() public {}

    // Annoyingly add internal methods needed to avoid stack too deep - why does foundry check for stack to deep in
    // tester contracts ? 😱
    function _checkInspectPolymerState(bytes memory proof, string memory expected) internal {
        (bytes32 stateRoot, uint64 height, bytes memory signature) = crossProverV2.inspectPolymerState(proof);
        assertEq(stateRoot, abi.decode(expected.parseRaw(".StateRoot"), (bytes32)));
        assertEq(height, abi.decode(expected.parseRaw(".PeptideHeight"), (uint64)));
        // Base 64 instead of other way since go defaults to encoding raw bytes in base64
        assertEq(Base64.encode(signature), expected.readString(".Signature"));
    }

    function _checkInspectLogIndentifier(bytes memory proof, string memory expected) internal {
        (uint32 srcChain, uint64 blockNumber, uint16 receiptIndex, uint8 logIndex) =
            crossProverV2.inspectLogIdentifier(proof);

        assertEq(srcChain, abi.decode(expected.parseRaw(".ChainID"), (uint32)));
        assertEq(blockNumber, abi.decode(expected.parseRaw(".BlockHeight"), (uint64)));
        assertEq(receiptIndex, abi.decode(expected.parseRaw(".ReceiptIndex"), (uint16)));

        assertEq(logIndex, expected.readUint(".Event.logIndex"));
    }
}
