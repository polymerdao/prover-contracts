// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import {Test} from "forge-std/Test.sol";
import {Registry} from "../../contracts/core/native_fallback/L1/Registry.sol";
import {L2Configuration, L1Configuration} from "../../contracts/libs/RegistryTypes.sol";

contract RegistryTest is Test {
    Registry public registry;
    address public owner;
    address public user;
    address public user2;
    uint256 public chainID = 123;
    L2Configuration public l2Config;
    L1Configuration public l1Config;

    event L2ChainConfigurationUpdated(uint256 indexed chainID, bytes32 indexed configHash);
    event L1ChainConfigurationUpdated(uint256 indexed chainID, bytes32 indexed configHash);

    function setUp() public {
        owner = address(0x1);
        user = address(0x2);
        user2 = address(0x3);

        address[] memory addresses = new address[](2);
        addresses[0] = address(0x4);
        addresses[1] = address(0x5);

        uint256[] memory storageSlots = new uint256[](2);
        storageSlots[0] = 10;
        storageSlots[1] = 11;

        l2Config = L2Configuration({
            prover: address(0x6),
            addresses: addresses,
            storageSlots: storageSlots,
            versionNumber: 1,
            finalityDelaySeconds: 3600
        });

        l1Config = L1Configuration({
            blockHashOracle: address(0x1234),
            settlementBlocksDelay: 10,
            settlementRegistry: address(0x5678),
            settlementRegistryL2ConfigMappingSlot: 5,
            settlementRegistryL1ConfigMappingSlot: 6
        });

        Registry.InitialL2Configuration[] memory initialL2Configs = new Registry.InitialL2Configuration[](1);
        initialL2Configs[0] = Registry.InitialL2Configuration({chainID: chainID, config: l2Config});

        Registry.InitialL1Configuration[] memory initialL1Configs = new Registry.InitialL1Configuration[](1);
        initialL1Configs[0] = Registry.InitialL1Configuration({chainID: chainID, config: l1Config});

        vm.prank(owner);
        registry = new Registry(owner, false, initialL2Configs, initialL1Configs);
    }

    function testConstructor() public view {
        assertEq(registry.multiSigOwner(), owner);

        // Test initial configuration by calculating expected hash
        bytes32 l1ConfigHash = keccak256(abi.encode(l1Config));
        bytes32 l2ConfigHash = keccak256(abi.encode(l2Config));
        assertEq(registry.l2ChainConfigurationHashMap(chainID), l2ConfigHash);
        assertEq(registry.l1ChainConfigurationHashMap(chainID), l1ConfigHash);
    }

    function testGrantChainID() public {
        uint256 testChainID = 456;

        // Verify user doesn't have access before grant
        assertFalse(registry.isGrantee(user, testChainID));

        // Grant access
        vm.prank(owner);
        registry.grantChainID(user, testChainID);

        // Verify user has access after grant
        assertTrue(registry.isGrantee(user, testChainID));

        // Verify other users don't have access
        assertFalse(registry.isGrantee(user2, testChainID));
    }

    function testGrantChainIDRange() public {
        uint256 startChainID = 1000;
        uint256 stopChainID = 1010;

        // Verify user doesn't have access before grant
        for (uint256 i = startChainID; i <= stopChainID; i++) {
            assertFalse(registry.isGrantee(user, i));
        }

        // Grant access to range
        vm.prank(owner);
        registry.grantChainIDRange(user, startChainID, stopChainID);

        // Verify user has access after grant
        for (uint256 i = startChainID; i <= stopChainID; i++) {
            assertTrue(registry.isGrantee(user, i));
        }

        // Verify user doesn't have access to chains outside the range
        assertFalse(registry.isGrantee(user, startChainID - 1));
        assertFalse(registry.isGrantee(user, stopChainID + 1));
    }

    function testUpdateChainConfiguration() public {
        uint256 newChainID = 789;

        // Setup test configuration
        address[] memory newAddresses = new address[](1);
        newAddresses[0] = address(0x7);

        uint256[] memory newStorageSlots = new uint256[](1);
        newStorageSlots[0] = 20;

        L2Configuration memory newConfig = L2Configuration({
            prover: address(0x8),
            addresses: newAddresses,
            storageSlots: newStorageSlots,
            versionNumber: 2,
            finalityDelaySeconds: 7200
        });

        // Grant access to user
        vm.prank(owner);
        registry.grantChainID(user, newChainID);

        // Expect an event to be emitted
        bytes32 expectedConfigHash = keccak256(abi.encode(newConfig));
        vm.expectEmit(true, true, true, true);
        emit L2ChainConfigurationUpdated(newChainID, expectedConfigHash);

        // Update configuration
        vm.prank(user);
        registry.updateL2ChainConfiguration(newChainID, newConfig);

        // Verify configuration was updated
        assertEq(registry.l2ChainConfigurationHashMap(newChainID), expectedConfigHash);
    }

    function testUpdateChainConfigurationFlow() public {
        uint256 newChainID = 789;

        // Setup initial configuration
        address[] memory initialAddresses = new address[](1);
        initialAddresses[0] = address(0x7);

        uint256[] memory initialStorageSlots = new uint256[](1);
        initialStorageSlots[0] = 20;

        L2Configuration memory initialConfig = L2Configuration({
            prover: address(0x8),
            addresses: initialAddresses,
            storageSlots: initialStorageSlots,
            versionNumber: 1,
            finalityDelaySeconds: 3600
        });

        // Setup updated configuration
        address[] memory updatedAddresses = new address[](2);
        updatedAddresses[0] = address(0x7);
        updatedAddresses[1] = address(0x9);

        uint256[] memory updatedStorageSlots = new uint256[](2);
        updatedStorageSlots[0] = 20;
        updatedStorageSlots[1] = 30;

        L2Configuration memory updatedConfig = L2Configuration({
            prover: address(0xA),
            addresses: updatedAddresses,
            storageSlots: updatedStorageSlots,
            versionNumber: 2,
            finalityDelaySeconds: 7200
        });

        // Grant access to user for this chain ID
        vm.prank(owner);
        registry.grantChainID(user, newChainID);

        // Verify user is granted access
        assertTrue(registry.isGrantee(user, newChainID));

        // Create initial configuration
        bytes32 initialConfigHash = keccak256(abi.encode(initialConfig));
        vm.expectEmit(true, true, true, true);
        emit L2ChainConfigurationUpdated(newChainID, initialConfigHash);

        vm.prank(user);
        registry.updateL2ChainConfiguration(newChainID, initialConfig);

        // Verify initial configuration is set correctly
        assertEq(registry.l2ChainConfigurationHashMap(newChainID), initialConfigHash);

        // Update the configuration
        bytes32 updatedConfigHash = keccak256(abi.encode(updatedConfig));
        vm.expectEmit(true, true, true, true);
        emit L2ChainConfigurationUpdated(newChainID, updatedConfigHash);

        vm.prank(user);
        registry.updateL2ChainConfiguration(newChainID, updatedConfig);

        // Verify updated configuration
        assertEq(registry.l2ChainConfigurationHashMap(newChainID), updatedConfigHash);

        // Verify user2 cannot update the config (non-grantee)
        vm.prank(user2);
        vm.expectRevert("Not authorized");
        registry.updateL2ChainConfiguration(newChainID, initialConfig);

        // Now we'll test that someone with a fresh granted permission can update properly
        vm.prank(owner);
        registry.grantChainID(user2, newChainID);

        // Verify user2 has access
        assertTrue(registry.isGrantee(user2, newChainID));

        // Update with user2
        L2Configuration memory finalConfig = L2Configuration({
            prover: address(0xB),
            addresses: updatedAddresses,
            storageSlots: updatedStorageSlots,
            versionNumber: 3,
            finalityDelaySeconds: 10_800
        });

        bytes32 finalConfigHash = keccak256(abi.encode(finalConfig));
        vm.expectEmit(true, true, true, true);
        emit L2ChainConfigurationUpdated(newChainID, finalConfigHash);

        vm.prank(user2);
        registry.updateL2ChainConfiguration(newChainID, finalConfig);

        // Verify final configuration
        assertEq(registry.l2ChainConfigurationHashMap(newChainID), finalConfigHash);
    }

    function testOnlyGranteeCanUpdateChainConfig() public {
        uint256 newChainID = 789;

        // Setup test configuration
        address[] memory newAddresses = new address[](1);
        newAddresses[0] = address(0x7);

        uint256[] memory newStorageSlots = new uint256[](1);
        newStorageSlots[0] = 20;

        L2Configuration memory newConfig = L2Configuration({
            prover: address(0x8),
            addresses: newAddresses,
            storageSlots: newStorageSlots,
            versionNumber: 2,
            finalityDelaySeconds: 7200
        });

        // Try to update without permission
        vm.prank(user);
        vm.expectRevert("Not authorized");
        registry.updateL2ChainConfiguration(newChainID, newConfig);
    }

    function testOnlyOwnerCanGrantChainID() public {
        // Try to grant as non-owner
        vm.prank(user);
        vm.expectRevert("Not authorized");
        registry.grantChainID(user2, 999);
    }

    function testOnlyOwnerCanGrantChainIDRange() public {
        // Try to grant range as non-owner
        vm.prank(user);
        vm.expectRevert("Not authorized");
        registry.grantChainIDRange(user2, 1000, 1010);
    }

    function testInvalidRangeReverts() public {
        // Try invalid range (start > stop)
        vm.prank(owner);
        vm.expectRevert("Invalid range");
        registry.grantChainIDRange(user, 1010, 1000);
    }

    function testConstructorRejectsZeroAddress() public {
        Registry.InitialL2Configuration[] memory initialL2Configs = new Registry.InitialL2Configuration[](0);
        Registry.InitialL1Configuration[] memory initialL1Configs = new Registry.InitialL1Configuration[](0);

        vm.expectRevert("Invalid multiSig address");
        new Registry(address(0), false, initialL2Configs, initialL1Configs);
    }

    function testUpdateL1ChainConfiguration() public {
        uint256 l1ChainID = 1;

        // Create L1 config
        L1Configuration memory newL1Config = L1Configuration({
            blockHashOracle: address(0x1234),
            settlementBlocksDelay: 10,
            settlementRegistry: address(0x5678),
            settlementRegistryL2ConfigMappingSlot: 5,
            settlementRegistryL1ConfigMappingSlot: 6
        });

        // Grant access to user
        vm.prank(owner);
        registry.grantChainID(user, l1ChainID);

        // Expect an event to be emitted
        bytes32 expectedConfigHash = keccak256(abi.encode(newL1Config));
        vm.expectEmit(true, true, true, true);
        emit L1ChainConfigurationUpdated(l1ChainID, expectedConfigHash);

        // Update configuration
        vm.prank(user);
        registry.updateL1ChainConfiguration(l1ChainID, newL1Config);

        // Verify configuration was updated
        assertEq(registry.l1ChainConfigurationHashMap(l1ChainID), expectedConfigHash);

        // Try to update without permission
        vm.prank(user2);
        vm.expectRevert("Not authorized");
        registry.updateL1ChainConfiguration(l1ChainID, newL1Config);
    }

    function testIrrevocableChainID() public {
        uint256 testChainID = 777;

        // Grant irrevocable access to user
        vm.prank(owner);
        registry.grantChainIDIrrevocable(user, testChainID);

        // Verify user has access
        assertTrue(registry.isGrantee(user, testChainID));

        // Attempt to grant the same chain ID to another user should fail
        vm.prank(owner);
        vm.expectRevert("ChainID is irrevocable");
        registry.grantChainID(user2, testChainID);

        // Attempt to grant irrevocable access to the same chain ID should fail
        vm.prank(owner);
        vm.expectRevert("ChainID is irrevocable");
        registry.grantChainIDIrrevocable(user2, testChainID);
    }

    function testIrrevocableChainIDRange() public {
        uint256 startChainID = 2000;
        uint256 stopChainID = 2010;

        // Grant irrevocable access to range
        vm.prank(owner);
        registry.grantChainIDRangeIrrevocable(user, startChainID, stopChainID);

        // Verify user has access to all chain IDs in the range
        for (uint256 i = startChainID; i <= stopChainID; i++) {
            assertTrue(registry.isGrantee(user, i));
        }

        // Attempt to grant any chain ID in the range should fail
        for (uint256 i = startChainID; i <= stopChainID; i++) {
            vm.prank(owner);
            vm.expectRevert("ChainID is irrevocable");
            registry.grantChainID(user2, i);
        }

        // Verify invalid range handling
        vm.prank(owner);
        vm.expectRevert("Invalid range");
        registry.grantChainIDRangeIrrevocable(user, stopChainID, startChainID);
    }

    function testGetL2ConfigAddresses() public view {
        // Get the addresses from the initial config
        address[] memory addresses = registry.getL2ConfigAddresses(chainID);

        // Verify against the addresses in l2Config
        assertEq(addresses.length, l2Config.addresses.length);
        for (uint256 i = 0; i < addresses.length; i++) {
            assertEq(addresses[i], l2Config.addresses[i]);
        }
    }

    function testGetL2ConfigStorageSlots() public view {
        // Get the storage slots from the initial config
        uint256[] memory slots = registry.getL2ConfigStorageSlots(chainID);

        // Verify against the storage slots in l2Config
        assertEq(slots.length, l2Config.storageSlots.length);
        for (uint256 i = 0; i < slots.length; i++) {
            assertEq(slots[i], l2Config.storageSlots[i]);
        }
    }

    function testOnlyAddLimitation() public {
        // Deploy a registry with onlyAdd set to true
        Registry.InitialL2Configuration[] memory initialL2Configs = new Registry.InitialL2Configuration[](0);
        Registry.InitialL1Configuration[] memory initialL1Configs = new Registry.InitialL1Configuration[](0);

        Registry restrictedRegistry = new Registry(owner, true, initialL2Configs, initialL1Configs);

        uint256 testChainID = 888;

        // Grant access to user
        vm.prank(owner);
        restrictedRegistry.grantChainID(user, testChainID);

        // Setup test configuration
        address[] memory addresses = new address[](1);
        addresses[0] = address(0x7);

        uint256[] memory storageSlots = new uint256[](1);
        storageSlots[0] = 20;

        L2Configuration memory config = L2Configuration({
            prover: address(0x8),
            addresses: addresses,
            storageSlots: storageSlots,
            versionNumber: 2,
            finalityDelaySeconds: 7200
        });

        // First update should work (adding a new config)
        vm.prank(user);
        restrictedRegistry.updateL2ChainConfiguration(testChainID, config);

        // Second update to the same chain ID should fail
        L2Configuration memory updatedConfig = L2Configuration({
            prover: address(0x9),
            addresses: addresses,
            storageSlots: storageSlots,
            versionNumber: 3,
            finalityDelaySeconds: 7200
        });

        vm.prank(user);
        vm.expectRevert("Existing config cannot be updated");
        restrictedRegistry.updateL2ChainConfiguration(testChainID, updatedConfig);

        // Same test for L1 configs
        L1Configuration memory testL1Config = L1Configuration({
            blockHashOracle: address(0x1234),
            settlementBlocksDelay: 10,
            settlementRegistry: address(0x5678),
            settlementRegistryL2ConfigMappingSlot: 5,
            settlementRegistryL1ConfigMappingSlot: 6
        });

        // First update should work (adding a new config)
        vm.prank(user);
        restrictedRegistry.updateL1ChainConfiguration(testChainID, testL1Config);

        // Second update to the same chain ID should fail
        L1Configuration memory updatedL1Config = L1Configuration({
            blockHashOracle: address(0x9876),
            settlementBlocksDelay: 20,
            settlementRegistry: address(0x5432),
            settlementRegistryL2ConfigMappingSlot: 7,
            settlementRegistryL1ConfigMappingSlot: 8
        });

        vm.prank(user);
        vm.expectRevert("Existing config cannot be updated");
        restrictedRegistry.updateL1ChainConfiguration(testChainID, updatedL1Config);
    }
}
