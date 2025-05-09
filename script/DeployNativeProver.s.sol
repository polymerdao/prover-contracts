// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/core/native_fallback/L2/NativeProver.sol";
import "../contracts/libs/RegistryTypes.sol";
import "./DeployRegistry.s.sol";

contract DeployNativeProverScript is DeployRegistryScript {
    function run() external override returns (address nativeProver) {
        uint256 chainId = vm.envUint("CHAIN_ID");
        address settlementRegistry = vm.envAddress("SETTLEMENT_REGISTRY");
        // Create an L1Configuration
        L1Configuration memory l1Config = L1Configuration({
            blockHashOracle: blockHashOracle,
            settlementRegistry: settlementRegistry,
            settlementRegistryL2ConfigMappingSlot: l2StorageSlot(chainId),
            settlementRegistryL1ConfigMappingSlot: l1StorageSlot(chainId),
            settlementBlocksDelay: blocksDelay
        });

        L2Configuration memory baseL2Config = L2Configuration({
            prover: baseProver,
            addresses: addresses(0xd6E6dBf4F7EA0ac412fD8b65ED297e64BB7a06E1),
            storageSlots: storageSlots(),
            versionNumber: 0,
            finalityDelaySeconds: 0,
            l2Type: Type.OPStackCannon
        });

        L2Configuration memory opL2Config = L2Configuration({
            prover: opProver,
            addresses: addresses(0x05F9613aDB30026FFd634f38e5C4dFd30a197Fa1),
            storageSlots: storageSlots(),
            versionNumber: 0,
            finalityDelaySeconds: 0,
            l2Type: Type.OPStackCannon
        });

        // // Create empty array for initial L2 configurations
        // NativeProver.InitialL2Configuration[] memory initialL2Configs = new NativeProver.InitialL2Configuration[](0);

        NativeProver.InitialL2Configuration[] memory initialL2Configs = new NativeProver.InitialL2Configuration[](2);
        initialL2Configs[0] = NativeProver.InitialL2Configuration({chainID: baseChainId, config: baseL2Config});

        initialL2Configs[1] = NativeProver.InitialL2Configuration({chainID: opChainId, config: opL2Config});

        vm.startBroadcast(deployerPrivateKey);

        // // Deploy NativeProver contract
        NativeProver nativeProverContract = new NativeProver(msg.sender, ethChainId, initialL2Configs);
        nativeProverContract.setInitialL1Config(l1Config);

        vm.stopBroadcast();
        console2.log("nativeProver: ", address(nativeProverContract));

        return address(nativeProverContract);
    }
}
