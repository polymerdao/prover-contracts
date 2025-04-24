// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/core/native_fallback/L2/NativeProver.sol";
import "../contracts/libs/RegistryTypes.sol";

contract DeployNativeProverScript is Script {
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
        
        vm.startBroadcast(deployerPrivateKey);
        
        // Deploy NativeProver contract
        NativeProver nativeProverContract = new NativeProver(
            l1ChainId,
            l1Config,
            initialL2Configs
        );
        
        vm.stopBroadcast();
        
        return address(nativeProverContract);
    }
}