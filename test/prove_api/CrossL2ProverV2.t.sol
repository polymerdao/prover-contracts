// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import {SigningBase} from "./utils/Signing.base.t.sol";
import {CrossL2Prover} from "../../contracts/core/prove_api/CrossL2Prover.sol";
import {CrossL2ProverV2} from "../../contracts/core/prove_api/CrossL2ProverV2.sol";
import {SequencerSignatureVerifier} from "../../contracts/core/prove_api/SequencerSignatureVerifier.sol";
import {ReceiptParser, Ics23Proof, OpIcs23Proof} from "../../contracts/libs/ReceiptParser.sol";
import {ISignatureVerifier} from "../../contracts/interfaces/ISignatureVerifier.sol";
import {ICrossL2Prover} from "../../contracts/interfaces/ICrossL2Prover.sol";
import {MerkleTrie} from "optimism/libraries/trie/MerkleTrie.sol";
import {IAppStateVerifier} from "../../contracts/interfaces/IAppStateVerifier.sol";
import {Base64} from "@openzeppelin/contracts/utils/Base64.sol";

contract CrossL2ProverTest is SigningBase {
    using stdJson for string;

    CrossL2ProverV2 crossProverV2;

    function setUp() public {
        crossProverV2 = new CrossL2ProverV2(
            "proof_api",
            0x8D3921B96A3815F403Fb3a4c7fF525969d16f9E0,
            0x0000000000000000000000000000000000000000000000000000000000000385
        );
    }

    // Test that a client update will fail if it doesn't have a valid sequencer signature
    function test_revert_invalidClentUpdateSignatureV2() public {
        vm.skip(true);
    }

    function test_validate_event_and_getters() public view {
        bytes memory proof = load_proof("/test/prove_api/payload/op-proof-v2.hex");

        console2.log("calldata length", proof.length);
        string memory expected = vm.readFile(string.concat(rootDir, "/test/prove_api/payload/op-event-v2.json"));

        (uint32 chainId, address addr, bytes memory topics, bytes memory data) = crossProverV2.validateEvent(proof);

        assertEq(chainId, 11_155_420);
        assertEq(addr, abi.decode(expected.parseRaw(".address"), (address)));
        bytes memory expectedTopics = abi.encodePacked(
            abi.decode(expected.parseRaw(".topics[0]"), (bytes32)),
            abi.decode(expected.parseRaw(".topics[1]"), (bytes32)),
            abi.decode(expected.parseRaw(".topics[2]"), (bytes32)),
            abi.decode(expected.parseRaw(".topics[3]"), (bytes32))
        );
        assertEq(topics, expectedTopics);
        assertEq(data, abi.encodePacked(abi.decode(expected.parseRaw(".data"), (bytes32))));

        // all these data come from the relayer log_test.go test

        // inspect polymer state
        (bytes32 stateRoot, uint64 height, bytes memory signature) = crossProverV2.inspectPolymerState(proof);
        assertEq(stateRoot, hex"00354047cc8b969297de60e7bcbffe8e524749634e358b1a06b2296fb924659d");
        assertEq(height, 3_130_134);
        // Base 64 instead of other way since go defaults to encoding raw bytes in base64
        assertEq(
            Base64.encode(signature),
            "Qfxc7kZXlD2c5A4wf5Q6sbtjcIYeLs1G3w1fbRt2kHdzaTz1PlckZ6vvynXF04uI6i2TpDruoJ3EN9AfxbM56xw="
        );

        // inspect log identifier
        (uint32 srcChain, uint64 blockNumber, uint16 receiptIndex, uint16 logIndex) =
            crossProverV2.inspectLogIdentifier(proof);

        assertEq(srcChain, 11_155_420);
        assertEq(blockNumber, 23_562_439);
        assertEq(receiptIndex, 1);
        assertEq(logIndex, 0);
    }

    function test_solana_validate() public view {
        bytes memory solProof = load_proof("/test/prove_api/payload/solana-proof.hex");
        string memory expected = vm.readFile(string.concat(rootDir, "/test/prove_api/payload/solana-event.json"));
        console2.log("calldata length", solProof.length);

        (uint32 chainId, bytes32 programID, string[] memory logMsgs) = crossProverV2.validateSolLogs(solProof);

        assertEq(chainId, 2);
        assertEq(programID, abi.decode(expected.parseRaw(".programID"), (bytes32)));
        assertEq(0, logMsgs.length);
    }

    function test_solana_validate_with_program_log() public view {
        bytes memory solProof = load_proof("/test/prove_api/payload/solana-proof-2.hex");
        string memory expected = vm.readFile(string.concat(rootDir, "/test/prove_api/payload/solana-event-2.json"));
        console2.log("calldata length", solProof.length);

        (uint32 chainId, bytes32 programID, string[] memory logMsges) = crossProverV2.validateSolLogs(solProof);
        console2.log("executor", toHexString(programID));
        console2.log("log msg", logMsges[0]);

        assertEq(chainId, 2);
        assertEq(programID, abi.decode(expected.parseRaw(".programID"), (bytes32)));
        assertEq(logMsges, abi.decode(expected.parseRaw(".logMessages"), (string[])));
    }

    function test_solana_event_root_key() public pure {
        bytes32 txSignatureHigh = hex"abb0a5d6f2eeb8f87a7549a3ef8f67d805e4985a8e6a51d0070f2f87d5b789d0";
        bytes32 txSignatureLow = hex"2c6b1bfb36f0b5a69c57f4947f1d0a847bc12a98cfdfc6e2fc8c5e5d2b68d47b";
        bytes32 programID = hex"3d2cf5c36c4571fc624b1e7dc0552b997c8e1f5e5b405afad1ad8b2f146e63f9";

        bytes memory expected = abi.encodePacked(
            "chain/2/storedLogs/proof_api/", uint64(42), "/", txSignatureHigh, txSignatureLow, "/", programID
        );
        assertEq(
            expected, ReceiptParser.solanaEventRootKey(2, "proof_api", 42, txSignatureHigh, txSignatureLow, programID)
        );
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

    function toHexString(bytes32 input) internal pure returns (string memory) {
        bytes memory result = new bytes(64); // 32 bytes * 2 hex digits per byte = 64 hex digits
        for (uint256 i = 0; i < 32; i++) {
            result[i * 2] = _byteToHexChar(uint8(input[i]) / 16);
            result[i * 2 + 1] = _byteToHexChar(uint8(input[i]) % 16);
        }
        return string(result);
    }

    function _byteToHexChar(uint8 byteValue) internal pure returns (bytes1) {
        if (byteValue < 10) {
            return bytes1(byteValue + 48); // '0' to '9'
        } else {
            return bytes1(byteValue + 87); // 'a' to 'f'
        }
    }
}
