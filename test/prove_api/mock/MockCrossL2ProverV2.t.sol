// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import {stdStorage, StdStorage} from "forge-std/Test.sol";
import {Ics23Proof} from "../../../contracts/libs/ReceiptParser.sol";
import {MockCrossL2ProverV2} from "../../../contracts/mocks/MockCrossL2ProverV2.sol";

contract TestMockCrossL2ProverV2 is Test {
    MockCrossL2ProverV2 mockProver;

    string clientType = "mock-proof";
    address sequencer = vm.addr(uint256(keccak256(unicode"🤑")));
    bytes32 peptideChainId = keccak256(unicode"🌎");
    bytes32[] topics;

    function setUp() public {
        mockProver = new MockCrossL2ProverV2(clientType, sequencer, peptideChainId);
    }

    function test_generateAndEmitProofMock() public {
        topics.push(keccak256("event"));
        topics.push(keccak256("topic 1"));

        // Test that the mocked proof will return the original chain id & topics
        bytes memory mockProof =
            mockProver.generateAndEmitProof(1, vm.addr(uint256(1231)), topics, abi.encode(unicode"💃🏻🧱🫡🔥"));

        (uint32 chainId, address emittingContract, bytes memory emittedTopics, bytes memory unindexedData) =
            mockProver.validateEvent(mockProof);

        assertEq(chainId, 1);
        assertEq(emittingContract, vm.addr(uint256(1231)));
        assertEq(emittedTopics, abi.encodePacked(topics[0], topics[1]));
        assertEq(unindexedData, abi.encode(unicode"💃🏻🧱🫡🔥"));
    }

    function testRevertIf_generateAndEmitProofMock() public {
        topics.push(keccak256("event"));
        topics.push(keccak256("topic 1"));

        vm.expectRevert();
        (uint32 chainId, address emittingContract, bytes memory emittedTopics, bytes memory unindexedData) =
            mockProver.validateEvent(abi.encode(1, vm.addr(uint256(1231)), topics, abi.encode(unicode"💃🏻🧱🫡🔥")));
    }
}
