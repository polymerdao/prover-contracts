// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/core/native_fallback/L2/NativeProver.sol";
import "../contracts/libs/RegistryTypes.sol";

/**
 * @title DeployNativeProverCreate2Script
 * @notice Deploys the NativeProver contract with CREATE2 for deterministic addresses
 * @dev This script deploys the NativeProver contract using CREATE2 with a fixed salt.
 *      When this script is run on multiple L2 chains with the same:
 *      1. Bytecode and constructor parameters
 *      2. Deployer address (same private key)
 *      3. Salt value (hardcoded in this script)
 *      The contract will be deployed to the same address on all chains.
 */
contract DeployNativeProverCreate2Script is Script {
    function run() external returns (address nativeProver) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        // Get L1 configuration from environment variables
        uint256 l1ChainId = vm.envOr("L1_CHAIN_ID", uint256(1));
        address blockHashOracle = vm.envOr("BLOCK_HASH_ORACLE", address(0));
        address settlementRegistry = vm.envOr("SETTLEMENT_REGISTRY", address(0));
        uint256 l2ConfigSlot = vm.envOr("L2_CONFIG_MAPPING_SLOT", uint256(0));
        uint256 l1ConfigSlot = vm.envOr("L1_CONFIG_MAPPING_SLOT", uint256(0));
        uint256 blocksDelay = vm.envOr("SETTLEMENT_BLOCKS_DELAY", uint256(0));
        
        // Create an L1Configuration
        L1Configuration memory l1Config = L1Configuration({
            blockHashOracle: blockHashOracle,
            settlementRegistry: settlementRegistry,
            settlementRegistryL2ConfigMappingSlot: l2ConfigSlot,
            settlementRegistryL1ConfigMappingSlot: l1ConfigSlot,
            settlementBlocksDelay: blocksDelay
        });
        
        // Create empty array for initial L2 configurations
        NativeProver.InitialL2Configuration[] memory initialL2Configs = new NativeProver.InitialL2Configuration[](0);
        
        // Use a fixed salt to ensure consistent deployment addresses across all chains
        // This salt value should not be changed once in production
        // For contracts with constructor args, it's good to include key parameters in the salt
        bytes32 salt = keccak256(abi.encodePacked(
            "polymer-native-prover",
            l1ChainId,
            blockHashOracle,
            settlementRegistry
        ));
        
        console.log("Using CREATE2 salt:", vm.toString(salt));
        
        vm.startBroadcast(deployerPrivateKey);
        
        // Deploy with CREATE2 using the salt parameter
        NativeProver nativeProverContract = new NativeProver{salt: salt}(
            l1ChainId,
            l1Config,
            initialL2Configs
        );
        
        console.log("Deployed NativeProver using CREATE2 at address:", address(nativeProverContract));
        
        vm.stopBroadcast();
        
        return address(nativeProverContract);
    }
}