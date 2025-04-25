// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/examples/TestContract.sol";

/**
 * @title DeployTestContractCreate2Script
 * @notice Deploys the TestContract with CREATE2 for deterministic addresses
 * @dev This script deploys the TestContract using CREATE2 with a fixed salt.
 *      When this script is run with the same:
 *      1. Bytecode and constructor parameters
 *      2. Deployer address (same private key)
 *      3. Salt value
 *      The contract will be deployed to the same address on all chains.
 */
contract DeployTestContractCreate2Script is Script {
    function run() external returns (address testContract) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        // Get initial value from environment, default to 42
        uint256 initialValue = vm.envOr("TEST_CONTRACT_INITIAL_VALUE", uint256(42));
        
        // Use a fixed salt to ensure consistent deployment addresses
        // Include the initial value in the salt to ensure unique addresses for different values
        bytes32 salt = keccak256(abi.encodePacked("polymer-test-contract", initialValue));
        
        console.log("Using CREATE2 salt:", vm.toString(salt));
        console.log("Initial value:", initialValue);
        
        vm.startBroadcast(deployerPrivateKey);
        
        // Deploy with CREATE2 using the salt parameter
        TestContract testContractInstance = new TestContract{salt: salt}(initialValue);
        
        console.log("Deployed TestContract using CREATE2 at address:", address(testContractInstance));
        
        vm.stopBroadcast();
        
        return address(testContractInstance);
    }
}