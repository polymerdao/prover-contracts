// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Test.sol";
import {CrossL2Executor} from "../../contracts/core/prove_api/CrossL2Executor.sol";
import {CrossL2ProverV2} from "../../contracts/core/prove_api/CrossL2ProverV2.sol";
import {SigningBase} from "./utils/Signing.base.t.sol";

contract CrossL2ExecutorTest is SigningBase {
    CrossL2ProverV2 prover;
    CrossL2Executor executor;

    // Events for testing
    event ValidationSuccess(uint32 chainId, address emittingContract, bytes topics, bytes unindexedData);
    event Ping();

    function setUp() public {
        // Deploy the prover contract
        prover = new CrossL2ProverV2(
            "proof_api",
            0x8D3921B96A3815F403Fb3a4c7fF525969d16f9E0,
            0x0000000000000000000000000000000000000000000000000000000000000385
        );

        // Deploy the executor contract with the prover address
        executor = new CrossL2Executor(address(prover));
    }

    /**
     * @notice Test successful execution with matching topics
     */
    function test_executeValidateEvent_success() public {
        // Skip if no test payload available
        string memory proofPath = string.concat(vm.projectRoot(), "/test/prove_api/payload/op-proof-v2.hex");

        try vm.readFile(proofPath) returns (string memory proofHex) {
            bytes memory proof = vm.parseBytes(proofHex);

            // Get expected topics by calling validateEvent directly on prover
            (,, bytes memory expectedTopics,) = prover.validateEvent(proof);

            // Get the expected values from validateEvent
            (uint32 chainId, address emittingContract,, bytes memory unindexedData) = prover.validateEvent(proof);

            // Expect ValidationSuccess event to be emitted
            vm.expectEmit(true, true, false, true);
            emit ValidationSuccess(chainId, emittingContract, expectedTopics, unindexedData);

            // Execute the validation - should succeed
            executor.executeValidateEvent(proof, expectedTopics);
        } catch {
            vm.skip(true); // Skip test if payload file doesn't exist
        }
    }

    /**
     * @notice Test execution reverts with mismatched topics
     */
    function test_executeValidateEvent_revert_topicsMismatch() public {
        // Skip if no test payload available
        string memory proofPath = string.concat(vm.projectRoot(), "/test/prove_api/payload/op-proof-v2.hex");

        try vm.readFile(proofPath) returns (string memory proofHex) {
            bytes memory proof = vm.parseBytes(proofHex);

            // Use wrong topics - just zeros
            bytes memory wrongTopics = abi.encodePacked(bytes32(0), bytes32(0));

            // Get actual topics for the revert message
            (,, bytes memory actualTopics,) = prover.validateEvent(proof);

            // Expect revert with TopicsDoNotMatch error
            vm.expectRevert(
                abi.encodeWithSelector(CrossL2Executor.TopicsDoNotMatch.selector, wrongTopics, actualTopics)
            );

            // Execute the validation - should revert
            executor.executeValidateEvent(proof, wrongTopics);
        } catch {
            vm.skip(true); // Skip test if payload file doesn't exist
        }
    }

    /**
     * @notice Test ping function
     */
    function test_ping() public {
        // Expect Ping event to be emitted
        vm.expectEmit(true, true, true, true);
        emit Ping();

        // Call ping function
        executor.ping();
    }

    /**
     * @notice Test that executor correctly references the prover
     */
    function test_prover_reference() public view {
        assertEq(address(executor.prover()), address(prover));
    }

    /**
     * @notice Test constructor with different prover address
     */
    function test_constructor_different_prover() public {
        address mockProver = address(0x123);
        CrossL2Executor newExecutor = new CrossL2Executor(mockProver);

        assertEq(address(newExecutor.prover()), mockProver);
    }
}
