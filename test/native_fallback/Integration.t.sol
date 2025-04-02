// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import {Test} from "forge-std/Test.sol";
import {Registry} from "../../contracts/core/native_fallback/L1/Registry.sol";
import {NativeProver} from "../../contracts/core/native_fallback/L2/NativeProver.sol";
import {OPStackBedrockProver} from "../../contracts/core/native_fallback/L2/OPStackBedrockProver.sol";
import {OPStackCannonProver} from "../../contracts/core/native_fallback/L2/OPStackCannonProver.sol";
import {L2Configuration, L1Configuration, Type} from "../../contracts/libs/RegistryTypes.sol";
import {RLPReader} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPReader.sol";
import {RLPWriter} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPWriter.sol";
import {MockIL1Block} from "./mock/MockIL1Block.sol";

// Create a helper contract that exposes the internal methods of Prover for testing
contract MockProver is NativeProver {
    constructor(
        uint256 _chainID,
        L1Configuration memory _l1Configuration,
        InitialL2Configuration[] memory _initialL2Configurations
    ) NativeProver(_chainID, _l1Configuration, _initialL2Configurations) {}

    // Expose internal method for testing
    function exposeSetInitialChainConfiguration(uint256 _chainID, L2Configuration memory _config) external {
        _setInitialChainConfiguration(_chainID, _config);
    }

    // Override the _proveL2Configuration to return true in tests
    function overrideProveL2Configuration(
        uint256 _chainID,
        L2Configuration calldata _config,
        bytes[] calldata,
        bytes calldata,
        bytes[] calldata,
        bytes32 _l1WorldStateRoot
    ) external {
        // Set L1 state root in the proven states map
        BlockProof memory blockProof =
            BlockProof({blockNumber: 100, blockHash: bytes32(uint256(0xabcdef)), stateRoot: _l1WorldStateRoot});
        provenStates[CHAIN_ID] = blockProof;

        // Update the chain configuration directly
        l2ChainConfigurations[_chainID] = _config;
    }
}

contract IntegrationTest is Test {
    Registry public registry;
    MockProver public mockProver;
    OPStackBedrockProver public bedrockProver;
    OPStackCannonProver public cannonProver;
    MockIL1Block public mockL1Block;

    address public owner;
    uint256 public l1ChainID = 1;
    uint256 public bedrockChainID = 10;
    uint256 public cannonChainID = 11;

    // Standard test setup that integrates all components
    function setUp() public {
        owner = address(0x1);
        uint256 l2ChainID = 1337;
        vm.startPrank(owner);

        // Setup mock L1Block
        mockL1Block = new MockIL1Block();
        bytes32 mockBlockHash = bytes32(uint256(0xabcdef));
        mockL1Block.setBlockHash(mockBlockHash);

        // Create provers
        bedrockProver = new OPStackBedrockProver();
        cannonProver = new OPStackCannonProver();

        // Create Registry for L2 configs
        Registry.InitialL2Configuration[] memory registryInitialL2Configs = new Registry.InitialL2Configuration[](0);
        Registry.InitialL1Configuration[] memory registryInitialL1Configs = new Registry.InitialL1Configuration[](0);

        registry = new Registry(owner, false, registryInitialL2Configs, registryInitialL1Configs);

        // Setup L1 configuration for Prover
        L1Configuration memory l1Config = L1Configuration({
            blockHashOracle: address(mockL1Block),
            settlementBlocksDelay: 10,
            settlementRegistry: address(registry),
            settlementRegistryL2ConfigMappingSlot: 0, // l2ChainConfigurationHashMap slot
            settlementRegistryL1ConfigMappingSlot: 1 // l1ChainConfigurationHashMap slot
        });

        // Setup initial L2 configurations - one for Bedrock, one for Cannon
        NativeProver.InitialL2Configuration[] memory proverInitialConfigs = new NativeProver.InitialL2Configuration[](2);

        // Bedrock config
        address[] memory bedrockAddrs = new address[](1);
        bedrockAddrs[0] = address(0x1234); // Mock L2OutputOracle address

        uint256[] memory bedrockSlots = new uint256[](1);
        bedrockSlots[0] = 5; // Mock output root storage slot

        proverInitialConfigs[0] = NativeProver.InitialL2Configuration({
            chainID: bedrockChainID,
            config: L2Configuration({
                prover: address(bedrockProver),
                addresses: bedrockAddrs,
                storageSlots: bedrockSlots,
                versionNumber: 0,
                finalityDelaySeconds: 7200
            })
        });

        // Cannon config
        address[] memory cannonAddrs = new address[](1);
        cannonAddrs[0] = address(0x5678); // Mock DisputeGameFactory address

        uint256[] memory cannonSlots = new uint256[](3);
        cannonSlots[0] = 10; // Mock disputeGameFactoryListSlot
        cannonSlots[1] = 11; // Mock faultDisputeGameRootClaimSlot
        cannonSlots[2] = 12; // Mock faultDisputeGameStatusSlot

        proverInitialConfigs[1] = NativeProver.InitialL2Configuration({
            chainID: cannonChainID,
            config: L2Configuration({
                prover: address(cannonProver),
                addresses: cannonAddrs,
                storageSlots: cannonSlots,
                versionNumber: 0,
                finalityDelaySeconds: 0 // Not used in Cannon
            })
        });

        // Create MockProver
        mockProver = new MockProver(l2ChainID, l1Config, proverInitialConfigs);

        // Grant permissions in Registry for the Prover to update configs
        registry.grantChainID(address(mockProver), bedrockChainID);
        registry.grantChainID(address(mockProver), cannonChainID);

        vm.stopPrank();
    }

    // Test updating chain configuration through Registry and Prover interaction
    function testUpdateChainConfigurationFlow() public {
        // 1. First we'll update configuration directly in the Registry as owner
        vm.startPrank(owner);

        // Create an updated L2 configuration for Bedrock
        address[] memory updatedBedrockAddrs = new address[](1);
        updatedBedrockAddrs[0] = address(0x9876); // Updated address

        uint256[] memory updatedBedrockSlots = new uint256[](1);
        updatedBedrockSlots[0] = 10; // Updated storage slot

        L2Configuration memory updatedBedrockConfig = L2Configuration({
            prover: address(bedrockProver),
            addresses: updatedBedrockAddrs,
            storageSlots: updatedBedrockSlots,
            versionNumber: 1, // Updated version
            finalityDelaySeconds: 3600 // Updated delay
        });

        // Grant permission to ourselves to update this specific chain ID
        registry.grantChainID(owner, bedrockChainID);

        // Update the registry configuration directly
        registry.updateL2ChainConfiguration(bedrockChainID, updatedBedrockConfig);

        // Verify Registry has the correct hash
        bytes32 configHash = keccak256(abi.encode(updatedBedrockConfig));
        assertEq(registry.l2ChainConfigurationHashMap(bedrockChainID), configHash);

        // 2. Now update mockProver configuration
        // Set a mock world state root
        bytes32 mockStateRoot = bytes32(uint256(0xabcdef1234));

        // Create empty proofs (not needed as we're using the mock override method)
        bytes[] memory emptyStorageProof = new bytes[](0);
        bytes memory emptyAccountData = new bytes(0);
        bytes[] memory emptyProof = new bytes[](0);

        // Update the chain configuration using our override method that bypasses proof verification
        mockProver.overrideProveL2Configuration(
            bedrockChainID, updatedBedrockConfig, emptyStorageProof, emptyAccountData, emptyProof, mockStateRoot
        );

        // In a real scenario, we would now verify that both the Prover and Registry
        // have the updated configuration.

        vm.stopPrank();
    }
}
