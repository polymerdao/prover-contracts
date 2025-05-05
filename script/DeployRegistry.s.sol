// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/core/native_fallback/L2/NativeProver.sol";
import "../contracts/core/native_fallback/L1/Registry.sol";
import "../contracts/libs/RegistryTypes.sol";

contract DeployRegistryScript is Script {
    uint256 constant _STARTING_L2_MAPPING_SLOT = 2;
    uint256 constant _STARTING_L1_MAPPING_SLOT = 4;
    uint256 deployerPrivateKey = vm.envUint("PKEY");
    address deployerAddr = vm.addr(deployerPrivateKey);
    // Get L1 configuration from environment variables
    address blockHashOracle = address(0x4200000000000000000000000000000000000015);

    uint256 ethChainId = uint256(1);

    uint256 opChainId = uint256(10);
    address opProver = vm.envAddress("OP_PROVER");

    uint256 baseChainId = uint256(8453);
    address baseProver = vm.envAddress("BASE_PROVER");

    uint256 blocksDelay = 0;

    function run() external virtual returns (address) {
        vm.startBroadcast(deployerPrivateKey);

        Registry.InitialL2Configuration[] memory emptyl2Configs = new Registry.InitialL2Configuration[](0);
        Registry.InitialL1Configuration[] memory emptyl1Configs = new Registry.InitialL1Configuration[](0);

        Registry settlementRegistry = new Registry(deployerAddr, emptyl2Configs, emptyl1Configs);

        // Create an L1Configuration
        L1Configuration memory baseL1Config = L1Configuration({
            blockHashOracle: blockHashOracle,
            settlementRegistry: address(settlementRegistry),
            settlementRegistryL2ConfigMappingSlot: l2StorageSlot(baseChainId),
            settlementRegistryL1ConfigMappingSlot: l1StorageSlot(baseChainId),
            settlementBlocksDelay: blocksDelay
        });

        L1Configuration memory opL1Config = L1Configuration({
            blockHashOracle: blockHashOracle,
            settlementRegistry: address(settlementRegistry),
            settlementRegistryL2ConfigMappingSlot: l2StorageSlot(opChainId),
            settlementRegistryL1ConfigMappingSlot: l1StorageSlot(opChainId),
            settlementBlocksDelay: blocksDelay
        });

        L2Configuration memory baseL2Config = L2Configuration({
            prover: baseProver,
            addresses: addresses(),
            storageSlots: storageSlots(),
            versionNumber: 0,
            finalityDelaySeconds: 0,
            l2Type: Type.OPStackCannon
        });

        L2Configuration memory opL2Config = L2Configuration({
            prover: opProver,
            addresses: addresses(),
            storageSlots: storageSlots(),
            versionNumber: 0,
            finalityDelaySeconds: 0,
            l2Type: Type.OPStackCannon
        });

        settlementRegistry.grantChainID(deployerAddr, baseChainId);
        settlementRegistry.grantChainID(deployerAddr, opChainId);

        settlementRegistry.updateL1ChainConfiguration(baseChainId, baseL1Config);

        settlementRegistry.updateL1ChainConfiguration(opChainId, opL1Config);

        settlementRegistry.updateL2ChainConfiguration(baseChainId, baseL2Config);
        settlementRegistry.updateL2ChainConfiguration(opChainId, opL2Config);
        vm.stopBroadcast();

        console2.log("CONTRACT settlementRegistry: ", address(settlementRegistry));
        return address(settlementRegistry);
    }

    // Given a chain id, return the storage slot for the L2 configuration
    function l2StorageSlot(uint256 chainID) internal pure returns (uint256) {
        // The storage slot for the chain ID is calculated as follows:
        // 1. The first 32 bytes of the storage slot are used for the chain ID.
        // 2. The remaining bytes are used for other data.
        // This is a simple example and may need to be adjusted based on your contract's storage layout.
        uint256 storageSlot = uint256(keccak256(abi.encode(chainID, _STARTING_L2_MAPPING_SLOT)));
        console2.log("L2 CONTRACT storage slot for chain ID: ", chainID, storageSlot);
        console2.logBytes32(keccak256(abi.encode(chainID, _STARTING_L2_MAPPING_SLOT)));

        return storageSlot;
    }

    // Given a chain id, return the storage slot for the L1 configuration
    function l1StorageSlot(uint256 chainID) internal pure returns (uint256) {
        // The storage slot for the chain ID is calculated as follows:
        // 1. The first 32 bytes of the storage slot are used for the chain ID.
        // 2. The remaining bytes are used for other data.
        // This is a simple example and may need to be adjusted based on your contract's storage layout.
        uint256 storageSlot = uint256(keccak256(abi.encode(chainID, _STARTING_L1_MAPPING_SLOT)));
        console2.log("L1 CONTRACT storage slot for chain ID: ", chainID, storageSlot);
        return storageSlot;
    }

    function addresses() public returns (address[] memory) {
        // Addresses and slots are the same for both op and base
        address[] memory _addresses = new address[](1);
        _addresses[0] = 0xe5965Ab5962eDc7477C8520243A95517CD252fA9;
        return _addresses;
    }

    function storageSlots() public returns (uint256[] memory) {
        uint256[] memory _storageSlots = new uint256[](3);
        _storageSlots[0] = 104;
        _storageSlots[1] = 0x405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad1;
        _storageSlots[2] = 0;
        return _storageSlots;
    }
}
