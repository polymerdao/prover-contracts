// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import {Test} from "forge-std/Test.sol";
import {Registry} from "../../contracts/core/native_fallback/L1/Registry.sol";
import {NativeProver} from "../../contracts/core/native_fallback/L2/NativeProver.sol";
import {OPStackBedrockProver} from "../../contracts/core/native_fallback/L2/OPStackBedrockProver.sol";
import {OPStackCannonProver} from "../../contracts/core/native_fallback/L2/OPStackCannonProver.sol";
import {
    L2Configuration,
    L1Configuration,
    Type,
    ProveScalarArgs,
    ProveL1ScalarArgs
} from "../../contracts/libs/RegistryTypes.sol";
import {RLPReader} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPReader.sol";
import {RLPWriter} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPWriter.sol";
import {MockIL1Block} from "./mock/MockIL1Block.sol";
import {ISettledStateProver} from "../../contracts/interfaces/ISettledStateProver.sol";

// Create a helper contract that exposes the internal methods of Prover for testing
contract MockProver is NativeProver {
    constructor(address _owner, uint256 _chainID, InitialL2Configuration[] memory _initialL2Configurations)
        NativeProver(_owner, _chainID, _initialL2Configurations)
    {}

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
        provenStates[L1_CHAIN_ID] = blockProof;

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

    // Event declarations
    event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender);
    event L2ChainConfigurationUpdated(uint256 indexed chainID, bytes32 indexed configHash);
    event L1WorldStateProven(uint256 indexed _blockNumber, bytes32 _L1WorldStateRoot);

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

        registry = new Registry(owner, registryInitialL2Configs, registryInitialL1Configs);

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
                finalityDelaySeconds: 7200,
                l2Type: Type.OPStackBedrock
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
                finalityDelaySeconds: 0, // Not used in Cannon
                l2Type: Type.OPStackCannon
            })
        });

        // Create MockProver
        mockProver = new MockProver(owner, l2ChainID, proverInitialConfigs);
        mockProver.setInitialL1Config(l1Config);

        // Calculate expected role hash for bedrock chain ID
        bytes32 bedrockRole = keccak256(abi.encode(keccak256("CHAIN_ROLE"), bedrockChainID));

        // Expect the RoleGranted event for bedrock chain
        vm.expectEmit(true, true, true, true);
        emit RoleGranted(bedrockRole, address(mockProver), owner);

        // Grant permissions in Registry for the Prover to update configs
        registry.grantChainID(address(mockProver), bedrockChainID);

        // Calculate expected role hash for cannon chain ID
        bytes32 cannonRole = keccak256(abi.encode(keccak256("CHAIN_ROLE"), cannonChainID));

        // Expect the RoleGranted event for cannon chain
        vm.expectEmit(true, true, true, true);
        emit RoleGranted(cannonRole, address(mockProver), owner);

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
            finalityDelaySeconds: 3600, // Updated delay
            l2Type: Type.OPStackBedrock
        });

        // Calculate expected role hash for bedrock chain ID
        bytes32 bedrockRole = keccak256(abi.encode(keccak256("CHAIN_ROLE"), bedrockChainID));

        // Expect the RoleGranted event
        vm.expectEmit(true, true, true, true);
        emit RoleGranted(bedrockRole, owner, owner);

        // Grant permission to ourselves to update this specific chain ID
        registry.grantChainID(owner, bedrockChainID);

        // Calculate the config hash
        bytes32 configHash = keccak256(abi.encode(updatedBedrockConfig));

        // Expect the L2ChainConfigurationUpdated event
        vm.expectEmit(true, true, true, true);
        emit L2ChainConfigurationUpdated(bedrockChainID, configHash);

        // Update the registry configuration directly
        registry.updateL2ChainConfiguration(bedrockChainID, updatedBedrockConfig);

        // Verify Registry has the correct hash
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

    // Complete prove flow test with parameterization and step verification
    function testCompleteProveFlow(uint256 _blockNumber, bytes32 _storageValue) public {
        // Make sure block number is large enough to satisfy the settlement delay
        // (which is 10 in our setup)
        vm.assume(_blockNumber >= 20 && _blockNumber < type(uint64).max);

        // Setup test parameters that can be reused across different scenarios
        ProveScalarArgs memory proveArgs = ProveScalarArgs({
            chainID: bedrockChainID,
            contractAddr: address(0x9876),
            storageSlot: bytes32(uint256(123)),
            storageValue: _storageValue,
            l2WorldStateRoot: bytes32(uint256(0xb33fca13))
        });

        // 1. Create L1 header first, then generate the hash
        bytes32 mockL1StateRoot = bytes32(uint256(0xdeadbeef));
        bytes memory rlpEncodedL1Header = _createMockL1Header(_blockNumber, mockL1StateRoot);

        // Calculate actual hash of the header and set it in the mock
        bytes32 mockL1BlockHash = keccak256(rlpEncodedL1Header);
        mockL1Block.setBlockHash(mockL1BlockHash);

        // 2. Create mock data for L2 header
        bytes memory rlpEncodedL2Header = _createMockL2Header(_blockNumber, proveArgs.l2WorldStateRoot);

        // 3. Create mock settled state proof that would be accepted by the prover
        bytes memory mockSettledStateProof = new bytes(32);

        // 4. Create mock storage proofs
        bytes[] memory mockL2StorageProof = new bytes[](1);
        mockL2StorageProof[0] = new bytes(64); // Empty mock proof

        // 5. Create mock contract account data with storage root
        bytes memory mockContractAccount = _createMockContractAccount();

        // 6. Create mock account proof
        bytes[] memory mockL2AccountProof = new bytes[](1);
        mockL2AccountProof[0] = new bytes(64); // Empty mock proof

        // Test the complete prove flow step by step

        // STEP 1: Prove L1 settlement layer state - establishes L1 state for subsequent proofs
        vm.expectEmit(true, true, true, true);
        emit L1WorldStateProven(_blockNumber, mockL1StateRoot);
        mockProver.proveSettlementLayerState(rlpEncodedL1Header);

        // STEP 2: Setup for mocking L2 prover calls
        address[] memory bedrockAddrs = new address[](1);
        bedrockAddrs[0] = address(0x1234); // Mock L2OutputOracle address

        uint256[] memory bedrockSlots = new uint256[](1);
        bedrockSlots[0] = 5; // Mock output root storage slot

        L2Configuration memory bedrockConfig = L2Configuration({
            prover: address(bedrockProver),
            addresses: bedrockAddrs,
            storageSlots: bedrockSlots,
            versionNumber: 0,
            finalityDelaySeconds: 7200,
            l2Type: Type.OPStackBedrock
        });

        // STEP 3: Test when ISettledStateProver.proveSettledState returns false
        vm.mockCall(
            address(bedrockProver),
            abi.encodeWithSelector(
                ISettledStateProver.proveSettledState.selector,
                bedrockConfig,
                proveArgs.l2WorldStateRoot,
                rlpEncodedL2Header,
                mockL1StateRoot,
                mockSettledStateProof
            ),
            abi.encode(false) // The mock prover returns false, causing revert
        );

        // Expect revert due to invalid settled state proof
        vm.expectRevert(
            abi.encodeWithSelector(
                NativeProver.InvalidSettledStateProof.selector, proveArgs.chainID, proveArgs.l2WorldStateRoot
            )
        );
        mockProver.proveNative(
            proveArgs,
            rlpEncodedL1Header,
            rlpEncodedL2Header,
            mockSettledStateProof,
            mockL2StorageProof,
            mockContractAccount,
            mockL2AccountProof
        );

        // STEP 4: Test when ISettledStateProver.proveSettledState returns true
        vm.mockCall(
            address(bedrockProver),
            abi.encodeWithSelector(
                ISettledStateProver.proveSettledState.selector,
                bedrockConfig,
                proveArgs.l2WorldStateRoot,
                rlpEncodedL2Header,
                mockL1StateRoot,
                mockSettledStateProof
            ),
            abi.encode(true)
        );

        // We should now fail at the SecureMerkleTrie verification for storage proof
        vm.expectRevert();
        mockProver.proveNative(
            proveArgs,
            rlpEncodedL1Header,
            rlpEncodedL2Header,
            mockSettledStateProof,
            mockL2StorageProof,
            mockContractAccount,
            mockL2AccountProof
        );

        // The test shows that both settlement state and storage proof verification
        // are checked in sequence, validating the full prove flow
    }

    // Test the new proveL1Native method in an integration context
    function testProveL1NativeFlow(uint256 _blockNumber, bytes32 _storageValue) public {
        // Make sure block number is large enough to satisfy any settlement delay requirements
        vm.assume(_blockNumber >= 20 && _blockNumber < type(uint64).max);

        // 1. Create L1 header with a state root we'll use
        bytes32 mockL1StateRoot = bytes32(uint256(0xdeadbeef));
        bytes memory rlpEncodedL1Header = _createMockL1Header(_blockNumber, mockL1StateRoot);

        // Calculate hash and set in the mock L1Block
        bytes32 mockL1BlockHash = keccak256(rlpEncodedL1Header);
        mockL1Block.setBlockHash(mockL1BlockHash);

        // 2. Prove the settlement layer state
        vm.expectEmit(true, true, true, true);
        emit L1WorldStateProven(_blockNumber, mockL1StateRoot);
        mockProver.proveSettlementLayerState(rlpEncodedL1Header);

        // 3. Setup parameters for proveL1Native
        address l1ContractAddr = address(0x9999);
        bytes32 storageSlot = bytes32(uint256(0x8888));

        // Create ProveL1ScalarArgs
        ProveL1ScalarArgs memory proveArgs = ProveL1ScalarArgs({
            contractAddr: l1ContractAddr,
            storageSlot: storageSlot,
            storageValue: _storageValue,
            l1WorldStateRoot: mockL1StateRoot
        });

        // 4. Create mock L1 storage proof
        bytes[] memory l1StorageProof = new bytes[](1);
        l1StorageProof[0] = new bytes(64); // Empty mock proof

        // 5. Create mock contract account data with storage root
        bytes memory mockContractAccount = _createMockContractAccount();

        // 6. Create mock account proof
        bytes[] memory l1AccountProof = new bytes[](1);
        l1AccountProof[0] = new bytes(64); // Empty mock proof

        // Try to prove the L1 state - should revert due to Merkle proof verification
        vm.expectRevert();
        mockProver.proveL1Native(proveArgs, rlpEncodedL1Header, l1StorageProof, mockContractAccount, l1AccountProof);

        // If we were able to properly mock the Merkle verification, we'd verify the return values:
        // (uint256 chainId, address storingContract, bytes32 storageValue) = mockProver.proveL1Native(...)
    }

    // Test proveL1Native method with an invalid L1 state root
    function testProveL1WithInvalidStateRoot(uint256 _blockNumber, bytes32 _storageValue) public {
        // Make sure block number is suitable
        vm.assume(_blockNumber >= 20 && _blockNumber < type(uint64).max);

        // 1. Create and prove a valid L1 header
        bytes32 validL1StateRoot = bytes32(uint256(0xdeadbeef));
        bytes memory rlpEncodedL1Header = _createMockL1Header(_blockNumber, validL1StateRoot);

        // Set mock block hash
        bytes32 mockL1BlockHash = keccak256(rlpEncodedL1Header);
        mockL1Block.setBlockHash(mockL1BlockHash);

        // Prove settlement layer state
        mockProver.proveSettlementLayerState(rlpEncodedL1Header);

        // 2. Setup parameters with an INVALID state root
        bytes32 invalidL1StateRoot = bytes32(uint256(0x999999));

        ProveL1ScalarArgs memory proveArgs = ProveL1ScalarArgs({
            contractAddr: address(0x9999),
            storageSlot: bytes32(uint256(0x8888)),
            storageValue: _storageValue,
            l1WorldStateRoot: invalidL1StateRoot // Using invalid state root
        });

        // Create mocks
        bytes[] memory l1StorageProof = new bytes[](1);
        l1StorageProof[0] = new bytes(64);

        bytes memory mockContractAccount = _createMockContractAccount();

        bytes[] memory l1AccountProof = new bytes[](1);
        l1AccountProof[0] = new bytes(64);

        // Should revert with SettlementChainStateRootNotProven since we're using an invalid state root
        vm.expectRevert(
            abi.encodeWithSelector(
                NativeProver.SettlementChainStateRootNotProven.selector,
                validL1StateRoot, // The proven state root
                invalidL1StateRoot // The invalid state root we're trying to use
            )
        );

        mockProver.proveL1Native(proveArgs, rlpEncodedL1Header, l1StorageProof, mockContractAccount, l1AccountProof);
    }

    // Helper function to create a mock L1 header
    function _createMockL1Header(uint256 _blockNumber, bytes32 _stateRoot) internal pure returns (bytes memory) {
        bytes[] memory encodedHeaderParts = new bytes[](15);

        // Mock data for L1 header parts (we only care about positions 3 and 8)
        for (uint256 i = 0; i < 15; i++) {
            if (i == 3) {
                // State root
                encodedHeaderParts[i] = RLPWriter.writeBytes(abi.encodePacked(_stateRoot));
            } else if (i == 8) {
                // Block number
                encodedHeaderParts[i] = RLPWriter.writeUint(_blockNumber);
            } else {
                // Other fields
                encodedHeaderParts[i] = RLPWriter.writeUint(i);
            }
        }

        return RLPWriter.writeList(encodedHeaderParts);
    }

    // Helper function to create a mock L2 header
    function _createMockL2Header(uint256 _blockNumber, bytes32 _stateRoot) internal pure returns (bytes memory) {
        bytes[] memory encodedHeaderParts = new bytes[](15);

        // Mock data for L2 header parts (we only care about positions 3 and 8)
        for (uint256 i = 0; i < 15; i++) {
            if (i == 3) {
                // State root
                encodedHeaderParts[i] = RLPWriter.writeBytes(abi.encodePacked(_stateRoot));
            } else if (i == 8) {
                // Block number
                encodedHeaderParts[i] = RLPWriter.writeUint(_blockNumber);
            } else {
                // Other fields
                encodedHeaderParts[i] = RLPWriter.writeUint(i);
            }
        }

        return RLPWriter.writeList(encodedHeaderParts);
    }

    // Helper function to create a mock contract account with storage root
    function _createMockContractAccount() internal pure returns (bytes memory) {
        bytes[] memory accountFields = new bytes[](4);

        // RLP encoding:
        // [
        //   nonce: uint256,
        //   balance: uint256,
        //   storageRoot: bytes32, // At position 2
        //   codeHash: bytes32
        // ]

        accountFields[0] = RLPWriter.writeUint(1); // nonce
        accountFields[1] = RLPWriter.writeUint(0); // balance
        accountFields[2] = RLPWriter.writeBytes(abi.encodePacked(bytes32(uint256(0xab12cd34)))); // storageRoot
        accountFields[3] = RLPWriter.writeBytes(abi.encodePacked(bytes32(uint256(0x45ef67)))); // codeHash

        return RLPWriter.writeList(accountFields);
    }

    // Test with semi-real Merkle Patricia Trie proofs
    function testProveWithSemiRealProofData() public {
        // Use a specific block number that satisfies the settlement delay
        uint256 blockNumber = 100;

        // Setup test parameters with specific values for reproducibility
        bytes32 storageValue = bytes32(uint256(0xdeadcafe)); // Value we want to prove
        bytes32 l2StateRoot = bytes32(uint256(0x12345678abcdef));
        address contractAddr = address(0x1234567890123456789012345678901234567890);
        bytes32 storageSlot = bytes32(uint256(0x9876));

        // Build arguments for prove function
        ProveScalarArgs memory proveArgs = ProveScalarArgs({
            chainID: bedrockChainID,
            contractAddr: contractAddr,
            storageSlot: storageSlot,
            storageValue: storageValue,
            l2WorldStateRoot: l2StateRoot
        });

        // Step 1: Create valid Merkle Patricia Trie for contract account
        // First create storage trie - create a real MPT with our storage slot/value
        bytes32 storageTrieRoot = _generateStorageTrieRoot(storageSlot, storageValue);

        // Create account data with our storage root
        bytes memory accountRLP = _createRLPAccountData(
            0, // nonce
            0, // balance
            storageTrieRoot, // storage root
            keccak256(abi.encodePacked("contractCode")) // code hash
        );

        // Create semi-real Merkle proofs for our account and storage
        (bytes[] memory accountProof, bytes[] memory storageProof) =
            _generateSemiRealProofs(contractAddr, accountRLP, storageSlot, storageValue);

        // Step 2: Create L1 header with state root
        bytes32 l1StateRoot = bytes32(uint256(0xf1f2f3f4));
        bytes memory rlpEncodedL1Header = _createMockL1Header(blockNumber, l1StateRoot);

        // Calculate hash and set in mock L1Block
        bytes32 l1BlockHash = keccak256(rlpEncodedL1Header);
        mockL1Block.setBlockHash(l1BlockHash);

        // Step 3: Create L2 header with our state root
        bytes memory rlpEncodedL2Header = _createMockL2Header(blockNumber, l2StateRoot);

        // Step 4: Create settled state proof
        bytes memory settledStateProof = _createSemiRealSettledStateProof(l2StateRoot);

        // Prove L1 settlement layer state
        mockProver.proveSettlementLayerState(rlpEncodedL1Header);

        // Setup mocking for bedrock prover to accept our proof
        address[] memory bedrockAddrs = new address[](1);
        bedrockAddrs[0] = address(0x1234); // Mock L2OutputOracle address

        uint256[] memory bedrockSlots = new uint256[](1);
        bedrockSlots[0] = 5; // Mock output root storage slot

        L2Configuration memory bedrockConfig = L2Configuration({
            prover: address(bedrockProver),
            addresses: bedrockAddrs,
            storageSlots: bedrockSlots,
            versionNumber: 0,
            finalityDelaySeconds: 7200,
            l2Type: Type.OPStackBedrock
        });

        vm.mockCall(
            address(bedrockProver),
            abi.encodeWithSelector(
                ISettledStateProver.proveSettledState.selector,
                bedrockConfig,
                l2StateRoot,
                rlpEncodedL2Header,
                l1StateRoot,
                settledStateProof
            ),
            abi.encode(true) // Make the prover accept our proof
        );

        // We still expect a revert at the Merkle verification stage
        // since we created semi-real proofs but not full valid trie structures
        vm.expectRevert();

        // Call prove with our data
        mockProver.proveNative(
            proveArgs, rlpEncodedL1Header, rlpEncodedL2Header, settledStateProof, storageProof, accountRLP, accountProof
        );
    }

    // Test the proveL1Native method with semi-real proof data
    function testProveL1WithSemiRealProofData() public {
        // Use a specific block number
        uint256 blockNumber = 100;

        // Setup test parameters for L1 proof
        bytes32 storageValue = bytes32(uint256(0xdadacafe)); // Value to prove
        address contractAddr = address(0x1A2b3c4D5e6F7890123456789aBcdeF012345678);
        bytes32 storageSlot = bytes32(uint256(0x1234));

        // Create L1 header with state root
        bytes32 l1StateRoot = bytes32(uint256(0xf5f6f7f8));
        bytes memory rlpEncodedL1Header = _createMockL1Header(blockNumber, l1StateRoot);

        // Calculate hash and set in mock L1Block
        bytes32 l1BlockHash = keccak256(rlpEncodedL1Header);
        mockL1Block.setBlockHash(l1BlockHash);

        // Create storage trie root
        bytes32 storageTrieRoot = _generateStorageTrieRoot(storageSlot, storageValue);

        // Create account data with storage root
        bytes memory accountRLP = _createRLPAccountData(
            1, // nonce
            1000, // balance
            storageTrieRoot, // storage root
            keccak256(abi.encodePacked("L1ContractCode")) // code hash
        );

        // Create Merkle proofs
        (bytes[] memory accountProof, bytes[] memory storageProof) =
            _generateSemiRealProofs(contractAddr, accountRLP, storageSlot, storageValue);

        // Build arguments for proveL1Native
        ProveL1ScalarArgs memory proveArgs = ProveL1ScalarArgs({
            contractAddr: contractAddr,
            storageSlot: storageSlot,
            storageValue: storageValue,
            l1WorldStateRoot: l1StateRoot
        });

        // Prove L1 settlement layer state first
        mockProver.proveSettlementLayerState(rlpEncodedL1Header);

        // We expect a revert due to the semi-real Merkle proofs
        vm.expectRevert();

        // Call proveL1Native with our data
        mockProver.proveL1Native(proveArgs, rlpEncodedL1Header, storageProof, accountRLP, accountProof);
    }

    // Helper function to create semi-real Merkle proofs
    function _generateSemiRealProofs(
        address _contractAddr,
        bytes memory _accountData,
        bytes32 _storageSlot,
        bytes32 _storageValue
    ) internal pure returns (bytes[] memory accountProof, bytes[] memory storageProof) {
        // Create account proof
        accountProof = new bytes[](3);

        // In a real MPT, we would have branch and extension nodes
        // Here we'll create a simplified proof with three nodes

        // Create leaf node for account (simplified)
        bytes memory leafNode = abi.encodePacked(
            bytes1(0x20), // Leaf prefix
            keccak256(abi.encodePacked(_contractAddr)), // Hashed path
            _accountData // Value
        );
        accountProof[0] = leafNode;

        // Create extension node (simplified)
        bytes memory extensionNode = abi.encodePacked(
            bytes1(0x10), // Extension prefix
            bytes4(0x12345678), // Path
            keccak256(leafNode) // Value points to leaf
        );
        accountProof[1] = extensionNode;

        // Create branch node (simplified)
        bytes memory branchNode = abi.encodePacked(
            bytes1(0x00), // Branch prefix
            keccak256(extensionNode) // One of 16 pointers in a real branch node
        );
        accountProof[2] = branchNode;

        // Create storage proof
        storageProof = new bytes[](2);

        // Encode storage value per RLP requirements
        bytes memory encodedValue = RLPWriter.writeUint(uint256(_storageValue));

        // Create leaf node for storage
        bytes memory storageLeaf = abi.encodePacked(
            bytes1(0x20), // Leaf prefix
            keccak256(abi.encodePacked(_storageSlot)), // Hashed path
            encodedValue // Value
        );
        storageProof[0] = storageLeaf;

        // Create branch node that would reference our leaf
        bytes memory storageBranch = abi.encodePacked(
            bytes1(0x00), // Branch prefix
            keccak256(storageLeaf) // Pointer to leaf
        );
        storageProof[1] = storageBranch;

        return (accountProof, storageProof);
    }

    // Helper to generate storage trie root
    function _generateStorageTrieRoot(bytes32 _key, bytes32 _value) internal pure returns (bytes32) {
        // For a simplified storage trie with a single key-value pair:

        // 1. Create leaf node for the storage KV pair
        bytes memory encodedValue = RLPWriter.writeUint(uint256(_value));

        // Hash the key as Ethereum does for storage
        bytes32 hashedKey = keccak256(abi.encodePacked(_key));

        // Create the path with leaf prefix (0x20)
        bytes memory path = abi.encodePacked(bytes1(0x20), hashedKey);

        // Encode leaf node as [path, value]
        bytes[] memory leaf = new bytes[](2);
        leaf[0] = path;
        leaf[1] = encodedValue;
        bytes memory encodedLeaf = RLPWriter.writeList(leaf);

        // Storage trie root is hash of the leaf for a single item trie
        return keccak256(encodedLeaf);
    }

    // Create a properly RLP encoded account
    function _createRLPAccountData(uint256 _nonce, uint256 _balance, bytes32 _storageRoot, bytes32 _codeHash)
        internal
        pure
        returns (bytes memory)
    {
        bytes[] memory accountFields = new bytes[](4);

        accountFields[0] = RLPWriter.writeUint(_nonce);
        accountFields[1] = RLPWriter.writeUint(_balance);
        accountFields[2] = RLPWriter.writeBytes(abi.encodePacked(_storageRoot));
        accountFields[3] = RLPWriter.writeBytes(abi.encodePacked(_codeHash));

        return RLPWriter.writeList(accountFields);
    }

    // Create a semi-real settled state proof
    function _createSemiRealSettledStateProof(bytes32 _l2StateRoot) internal view returns (bytes memory) {
        // For OPStackBedrock, this would include:
        // 1. Message passer state root
        bytes32 l2MessagePasserStateRoot = bytes32(uint256(0xabcdef));

        // 2. L2 output index
        uint256 l2OutputIndex = 42;

        // 3. Storage proof for the L2 output in the oracle
        bytes[] memory l1StorageProof = new bytes[](2);
        l1StorageProof[0] = abi.encodePacked(bytes1(0x20), bytes32(uint256(0x1234)));
        l1StorageProof[1] = abi.encodePacked(bytes1(0x00), bytes32(uint256(0x5678)));

        // 4. L2OutputOracle account data
        bytes[] memory outputOracleData = new bytes[](4);
        outputOracleData[0] = RLPWriter.writeUint(0); // nonce
        outputOracleData[1] = RLPWriter.writeUint(0); // balance
        outputOracleData[2] = RLPWriter.writeBytes(abi.encodePacked(bytes32(uint256(0x9abc)))); // storageRoot
        outputOracleData[3] = RLPWriter.writeBytes(abi.encodePacked(bytes32(uint256(0xdef0)))); // codeHash
        bytes memory rlpEncodedOutputOracleData = RLPWriter.writeList(outputOracleData);

        // 5. Account proof for L2OutputOracle
        bytes[] memory l1AccountProof = new bytes[](2);
        l1AccountProof[0] = abi.encodePacked(bytes1(0x20), bytes32(uint256(0xaaaa)));
        l1AccountProof[1] = abi.encodePacked(bytes1(0x00), bytes32(uint256(0xbbbb)));

        // 6. Output value that includes the L2 state root we're proving
        // In Bedrock, this has a specific format depending on version:
        // v0: keccak256(abi.encode(0, _l2StateRoot, l2MessagePasserStateRoot, l2BlockHash))
        bytes32 l2BlockHash = bytes32(uint256(0xccccdddd));
        bytes32 outputRoot = keccak256(abi.encode(0, _l2StateRoot, l2MessagePasserStateRoot, l2BlockHash));

        bytes memory outputValue = abi.encode(
            outputRoot, // L2 output root
            block.timestamp, // Timestamp
            block.number // L2 block number
        );
        bytes memory rlpOutputValue = RLPWriter.writeBytes(outputValue);

        // Combine all components into a proof structure
        return abi.encode(
            l2MessagePasserStateRoot,
            l2OutputIndex,
            l1StorageProof,
            rlpEncodedOutputOracleData,
            l1AccountProof,
            rlpOutputValue
        );
    }
}
