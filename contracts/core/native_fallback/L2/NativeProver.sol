// SPDX-License-Identifier: Apache-2.0
/*
 * Copyright 2024, Polymer Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity 0.8.15;

import {SecureMerkleTrie} from "@eth-optimism/contracts-bedrock/src/libraries/trie/SecureMerkleTrie.sol";
import {RLPReader} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPReader.sol";
import {IL1Block} from "../../../interfaces/IL1Block.sol";
import {INativeProver} from "../../../interfaces/INativeProver.sol";
import {ISettledStateProver} from "../../../interfaces/ISettledStateProver.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {
    L2Configuration,
    L1Configuration,
    ProveScalarArgs,
    ProveL1ScalarArgs,
    UpdateL2ConfigArgs
} from "../../../libs/RegistryTypes.sol";
import {ProverHelpers} from "../../../libs/ProverHelpers.sol";

contract NativeProver is Ownable, INativeProver {
    struct InitialL2Configuration {
        uint256 chainID;
        L2Configuration config;
    }

    /**
     * @notice Stores proven block data for a chain
     * @param blockNumber Number of the proven block
     * @param blockHash Hash of the proven block
     * @param stateRoot State root of the proven block
     */
    struct BlockProof {
        uint256 blockNumber;
        bytes32 blockHash;
        bytes32 stateRoot;
    }

    uint256 public immutable L1_CHAIN_ID; // Chain ID of the settlement chain

    L1Configuration public L1_CONFIGURATION; // Configuration for L1

    mapping(uint256 => L2Configuration) public l2ChainConfigurations; // Mapping of counterparty chainIDs to their
        // configurations

    mapping(uint256 => BlockProof) public provenStates;

    mapping(uint256 => ISettledStateProver) public stateProvers;

    event L1WorldStateProven(uint256 indexed _blockNumber, bytes32 _L1WorldStateRoot);

    event L2WorldStateProven(
        uint256 indexed _destinationChainID, uint256 indexed _blockNumber, bytes32 _L2WorldStateRoot
    );

    /**
     * @notice Block number is too recent to prove
     */
    error NeedLaterBlock(uint256 _inputBlockNumber, uint256 _nextProvableBlockNumber);

    /**
     * @notice Block number is older than currently proven block
     */
    error OutdatedBlock(uint256 _inputBlockNumber, uint256 _latestBlockNumber);

    /**
     * @notice RLP encoded block data hash mismatch
     */
    error InvalidRLPEncodedBlock(bytes32 _expectedBlockHash, bytes32 _calculatedBlockHash);

    /**
     * @notice Failed storage proof verification
     */
    error InvalidStorageProof(bytes _key, bytes _val, bytes[] _proof, bytes32 _root);

    /**
     * @notice Failed account proof verification
     */
    error InvalidAccountProof(bytes _address, bytes _data, bytes[] _proof, bytes32 _root);

    /**
     * @notice Settlement chain state not yet proven
     */
    error SettlementChainStateRootNotProven(bytes32 _blockProofStateRoot, bytes32 _l1WorldStateRoot);

    /**
     * @notice Destination chain state not yet proven
     */
    error DestinationChainStateRootNotProved(bytes32 _blockProofStateRoot, bytes32 _l2WorldStateRoot);

    /**
     * @notice Invalid storage root found in contract
     */
    error IncorrectContractStorageRoot(bytes _contractStorageRoot);

    /**
     * @notice Invalid L1 registry proof for L2 configuration update
     */
    error InvalidL2ConfigurationProof(uint256 _chainID, L2Configuration _config);

    /**
     * @notice Invalid L1 registry proof for L1 configuration update
     */
    error InvalidL1ConfigurationProof(L1Configuration _config);

    /**
     * @notice Invalid settled state proof
     */
    error InvalidSettledStateProof(uint256 _chainID, bytes32 _l2WorldStateRoot);

    /**
     * @notice Constructor briefly sets an owner which can set the l1 registry information one time and then transfers
     * ownership away
     * This is to allow for NativeProver contracts to have the same addresses when using create2 even if they have
     * different l1 configs.
     *
     */
    constructor(address _owner, uint256 _l1ChainID, InitialL2Configuration[] memory _initialL2Configurations)
        Ownable()
    {
        L1_CHAIN_ID = _l1ChainID;
        for (uint256 i = 0; i < _initialL2Configurations.length; ++i) {
            _setInitialChainConfiguration(_initialL2Configurations[i].chainID, _initialL2Configurations[i].config);
        }
        transferOwnership(_owner);
    }

    function setInitialL1Config(L1Configuration memory _l1Configuration) external onlyOwner {
        L1_CONFIGURATION = _l1Configuration;
        renounceOwnership();
    }

    /**
     * @notice Proves a storage value in a settled L2 state root
     * @dev Single view function encapsulating all three steps of the proof process: L1 view; settled L2 state; L2
     * storage
     * @param _updateArgs holds the ProveScalarArgs
     * @param _rlpEncodedL1Header L1 header to be verified against L1 block hash oracle to access L1 state root
     * @param _rlpEncodedL2Header L2 header proven in the L1 settlement
     * @param _settledStateProof proof of the L2 world state root in the L1 settlement
     * @param _l2StorageProof proof of the storage value in the L2 _contractAddr
     * @param _rlpEncodedContractAccount RLP encoded _contractAddr account
     * @param _l2AccountProof proof of the _contractAddr account in the L2 world state
     *
     */
    function proveNative(
        UpdateL2ConfigArgs calldata _updateArgs,
        ProveScalarArgs calldata _proveArgs,
        bytes calldata _rlpEncodedL1Header,
        bytes memory _rlpEncodedL2Header,
        bytes calldata _settledStateProof,
        bytes[] calldata _l2StorageProof,
        bytes calldata _rlpEncodedContractAccount,
        bytes[] calldata _l2AccountProof
    ) external view returns (uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue) {
        // First prove the L1 view
        bytes32 _l1StateRoot = _validateL1BlockAndGetStateRoot(_rlpEncodedL1Header);

        // Use the L1 state root to prove the L2 configuration, but don't store it
        if (
            !_proveL2Configuration(
                _proveArgs.chainID,
                _updateArgs.config,
                _updateArgs.l1StorageProof,
                _updateArgs.rlpEncodedRegistryAccountData,
                _updateArgs.l1RegistryProof,
                _l1StateRoot
            )
        ) {
            revert InvalidL2ConfigurationProof(_proveArgs.chainID, _updateArgs.config);
        }

        // Then prove the settled state
        _proveSettledState(
            _proveArgs.chainID, _proveArgs.l2WorldStateRoot, _l1StateRoot, _rlpEncodedL2Header, _settledStateProof
        );

        // Now prove storage against the verified settled L2 state root using the verified L2 configuration
        _proveStorageInState(_proveArgs, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof);

        return (_proveArgs.chainID, _proveArgs.contractAddr, _proveArgs.storageSlot, _proveArgs.storageValue);
    }

    function proveL1Native(
        ProveL1ScalarArgs calldata _args,
        bytes calldata _rlpEncodedL1Header,
        bytes[] calldata _l1StorageProof,
        bytes calldata _rlpEncodedContractAccount,
        bytes[] calldata _l1AccountProof
    ) external view returns (uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue) {
        // First prove the L1 view
        bytes32 _l1StateRoot = _validateL1BlockAndGetStateRoot(_rlpEncodedL1Header);

        // Check that the expected L1 state root matches the verified one
        if (_l1StateRoot != _args.l1WorldStateRoot) {
            revert SettlementChainStateRootNotProven(_l1StateRoot, _args.l1WorldStateRoot);
        }

        // Now prove storage against the L1 state root
        _proveL1StorageInState(_args, _l1StorageProof, _rlpEncodedContractAccount, _l1AccountProof);

        return (L1_CHAIN_ID, _args.contractAddr, _args.storageSlot, _args.storageValue);
    }

    function _setInitialChainConfiguration(uint256 _chainID, L2Configuration memory _config) internal {
        l2ChainConfigurations[_chainID] = _config;
    }

    function _updateL2ChainConfiguration(
        uint256 _chainID,
        L2Configuration calldata _config,
        bytes[] calldata _l1StorageProof,
        bytes calldata _rlpEncodedRegistryAccountData,
        bytes[] calldata _l1RegistryProof,
        bytes32 _l1WorldStateRoot
    ) internal {
        if (
            !_proveL2Configuration(
                _chainID, _config, _l1StorageProof, _rlpEncodedRegistryAccountData, _l1RegistryProof, _l1WorldStateRoot
            )
        ) {
            revert InvalidL2ConfigurationProof(_chainID, _config);
        }
        l2ChainConfigurations[_chainID] = _config;
    }

    /**
     * @notice Updates the proven state if newer
     * @param _chainID Chain ID to update
     * @param blockProof Block proof data to update with
     */
    function _updateProvenState(uint256 _chainID, BlockProof memory blockProof) internal {
        BlockProof memory existingBlockProof = provenStates[_chainID];

        if (existingBlockProof.blockNumber < blockProof.blockNumber) {
            provenStates[_chainID] = blockProof;
            emit L2WorldStateProven(_chainID, blockProof.blockNumber, blockProof.stateRoot);
        } else {
            if (existingBlockProof.blockNumber > blockProof.blockNumber) {
                revert OutdatedBlock(blockProof.blockNumber, existingBlockProof.blockNumber);
            }
        }
    }

    /**
     * @notice Proves an L2 configuration in the L1 registry contract through the L1 view
     * @dev Validates the L2 configuration against the L1 registry contract's state using the provided proof
     * @param _chainID chain ID of the L2 configuration being proven
     * @param _config L2 configuration being proven
     * @param _l1StorageProof proof of the config hash in the registry contract
     * @param _rlpEncodedRegistryAccountData RLP encoded registry contract account data
     * @param _l1RegistryProof proof of the registry contract account
     * @param _l1WorldStateRoot L1 world state root
     *
     */
    function _proveL2Configuration(
        uint256 _chainID,
        L2Configuration calldata _config,
        bytes[] calldata _l1StorageProof,
        bytes calldata _rlpEncodedRegistryAccountData,
        bytes[] calldata _l1RegistryProof,
        bytes32 _l1WorldStateRoot
    ) internal view returns (bool) {
        BlockProof memory existingSettlementBlockProof = provenStates[L1_CHAIN_ID];

        // Verify settlement chain state root
        if (existingSettlementBlockProof.stateRoot != _l1WorldStateRoot) {
            revert SettlementChainStateRootNotProven(existingSettlementBlockProof.stateRoot, _l1WorldStateRoot);
        }

        // Verify proof of the registry contract account
        _proveAccount(
            abi.encodePacked(L1_CONFIGURATION.settlementRegistry),
            _rlpEncodedRegistryAccountData,
            _l1RegistryProof,
            _l1WorldStateRoot
        );

        // Verify the L2 configuration against the registry contract's state
        bytes memory registryStorageRoot = RLPReader.readBytes(RLPReader.readList(_rlpEncodedRegistryAccountData)[2]);

        if (registryStorageRoot.length > 32) {
            revert IncorrectContractStorageRoot(registryStorageRoot);
        }
        bytes32 configHash = keccak256(abi.encode(_config));
        bytes32 configHashStorageSlot =
            bytes32((uint256(keccak256(abi.encode(_chainID))) + L1_CONFIGURATION.settlementRegistryL2ConfigMappingSlot));
        ProverHelpers.proveStorageBytes32(
            abi.encodePacked(configHashStorageSlot), configHash, _l1StorageProof, bytes32(registryStorageRoot)
        );

        return true;
    }

    /**
     * @notice Proves an L1 configuration in the L1 registry contract through the L1 view
     * @dev Validates the L1 configuration against the L1 registry contract's state using the provided proof
     * @param _config L1 configuration being proven
     * @param _l1StorageProof proof of the config hash in the registry contract
     * @param _rlpEncodedRegistryAccountData RLP encoded registry contract account data
     * @param _l1RegistryProof proof of the registry contract account
     * @param _l1WorldStateRoot L1 world state root
     *
     */
    function _proveL1Configuration(
        L1Configuration calldata _config,
        bytes[] calldata _l1StorageProof,
        bytes calldata _rlpEncodedRegistryAccountData,
        bytes[] calldata _l1RegistryProof,
        bytes32 _l1WorldStateRoot
    ) internal view returns (bool) {
        uint256 _l1ChainID = L1_CHAIN_ID;
        BlockProof memory existingSettlementBlockProof = provenStates[_l1ChainID];

        // Verify settlement chain state root
        if (existingSettlementBlockProof.stateRoot != _l1WorldStateRoot) {
            revert SettlementChainStateRootNotProven(existingSettlementBlockProof.stateRoot, _l1WorldStateRoot);
        }

        // Verify proof of the registry contract account
        _proveAccount(
            abi.encodePacked(L1_CONFIGURATION.settlementRegistry),
            _rlpEncodedRegistryAccountData,
            _l1RegistryProof,
            _l1WorldStateRoot
        );

        // Verify the L2 configuration against the registry contract's state
        bytes memory registryStorageRoot = RLPReader.readBytes(RLPReader.readList(_rlpEncodedRegistryAccountData)[2]);

        if (registryStorageRoot.length > 32) {
            revert IncorrectContractStorageRoot(registryStorageRoot);
        }
        bytes32 configHash = keccak256(abi.encode(_config));
        bytes32 configHashStorageSlot = bytes32(
            (uint256(keccak256(abi.encode(_l1ChainID))) + L1_CONFIGURATION.settlementRegistryL1ConfigMappingSlot)
        );
        ProverHelpers.proveStorageBytes32(
            abi.encodePacked(configHashStorageSlot), configHash, _l1StorageProof, bytes32(registryStorageRoot)
        );

        return true;
    }

    /**
     * @notice Validates the L1 block data and extracts state root
     * @param _rlpEncodedL1Header The encoded L1 header
     * @return L1 state root from the header
     */
    function _validateL1BlockAndGetStateRoot(bytes calldata _rlpEncodedL1Header) internal view returns (bytes32) {
        bytes32 _calculatedBlockHash = keccak256(_rlpEncodedL1Header);
        bytes32 _expectedBlockHash = IL1Block(L1_CONFIGURATION.blockHashOracle).hash();
        if (_calculatedBlockHash != _expectedBlockHash) {
            revert InvalidRLPEncodedBlock(_expectedBlockHash, _calculatedBlockHash);
        }

        return bytes32(RLPReader.readBytes(RLPReader.readList(_rlpEncodedL1Header)[3]));
    }

    function _proveSettledState(
        uint256 _chainID,
        bytes32 _l2WorldStateRoot,
        bytes32 _l1WorldStateRoot,
        bytes memory _rlpEncodedL2Header,
        bytes calldata _settledStateProof
    ) internal view {
        // Get the L2Configuration for this chainID
        L2Configuration memory conf = l2ChainConfigurations[_chainID];

        // Call out to the configured prover to verify proof of the settled L2 state root
        if (
            !ISettledStateProver(conf.prover).proveSettledState(
                conf, _l2WorldStateRoot, _rlpEncodedL2Header, _l1WorldStateRoot, _settledStateProof
            )
        ) {
            revert InvalidSettledStateProof(_chainID, _l2WorldStateRoot);
        }
    }

    function _proveSettledStateWithL2Config(
        L2Configuration memory _conf,
        uint256 _chainID,
        bytes32 _l2WorldStateRoot,
        bytes32 _l1WorldStateRoot,
        bytes memory _rlpEncodedL2Header,
        bytes calldata _settledStateProof
    ) internal view {
        // Call out to the configured prover to verify proof of the settled L2 state root
        if (
            !ISettledStateProver(_conf.prover).proveSettledState(
                _conf, _l2WorldStateRoot, _rlpEncodedL2Header, _l1WorldStateRoot, _settledStateProof
            )
        ) {
            revert InvalidSettledStateProof(_chainID, _l2WorldStateRoot);
        }
    }

    /**
     * @notice Validates an account proof against a root
     * @dev Uses SecureMerkleTrie for verification
     * @param _address Account address
     * @param _data Expected account data
     * @param _proof Merkle proof
     * @param _root Expected root
     */
    function _proveAccount(bytes memory _address, bytes memory _data, bytes[] memory _proof, bytes32 _root)
        internal
        pure
    {
        if (!SecureMerkleTrie.verifyInclusionProof(_address, _data, _proof, _root)) {
            revert InvalidAccountProof(_address, _data, _proof, _root);
        }
    }

    /**
     * @notice Converts bytes to uint256
     * @dev Manual byte-by-byte conversion
     * @param _b Bytes to convert
     * @return Converted uint256 value
     */
    function _bytesToUint(bytes memory _b) internal pure returns (uint256) {
        uint256 number;
        for (uint256 i = 0; i < _b.length; i++) {
            number = number + uint256(uint8(_b[i])) * (2 ** (8 * (_b.length - (i + 1))));
        }
        return number;
    }

    /**
     * @notice Internal function to create and verify block proof
     * @param _l2WorldStateRoot L2 state root to prove
     * @param _rlpEncodedL2Header RLP encoded L2 block header
     * @return blockProof Verified block proof data
     */
    function _createBlockProof(bytes32 _l2WorldStateRoot, bytes memory _rlpEncodedL2Header)
        internal
        pure
        returns (BlockProof memory blockProof)
    {
        blockProof = BlockProof({
            blockNumber: _bytesToUint(RLPReader.readBytes(RLPReader.readList(_rlpEncodedL2Header)[8])),
            blockHash: keccak256(_rlpEncodedL2Header),
            stateRoot: _l2WorldStateRoot
        });
    }

    function _proveStorageInState(
        ProveScalarArgs calldata _args,
        bytes[] calldata _l2StorageProof,
        bytes calldata _rlpEncodedContractAccount,
        bytes[] calldata _l2AccountProof
    ) internal pure {
        bytes memory storageRoot = RLPReader.readBytes(RLPReader.readList(_rlpEncodedContractAccount)[2]);

        if (storageRoot.length > 32) {
            revert IncorrectContractStorageRoot(storageRoot);
        }

        // Verify the account exists in the state tree
        _proveAccount(
            abi.encodePacked(_args.contractAddr), _rlpEncodedContractAccount, _l2AccountProof, _args.l2WorldStateRoot
        );

        // Verify the storage value exists in the storage tree
        ProverHelpers.proveStorageBytes32(
            abi.encodePacked(_args.storageSlot), _args.storageValue, _l2StorageProof, bytes32(storageRoot)
        );
    }

    function _proveL1StorageInState(
        ProveL1ScalarArgs calldata _args,
        bytes[] calldata _l2StorageProof,
        bytes calldata _rlpEncodedContractAccount,
        bytes[] calldata _l2AccountProof
    ) internal pure {
        bytes memory storageRoot = RLPReader.readBytes(RLPReader.readList(_rlpEncodedContractAccount)[2]);

        if (storageRoot.length > 32) {
            revert IncorrectContractStorageRoot(storageRoot);
        }

        // Verify the account exists in the state tree
        _proveAccount(
            abi.encodePacked(_args.contractAddr), _rlpEncodedContractAccount, _l2AccountProof, _args.l1WorldStateRoot
        );

        // Verify the storage value exists in the storage tree
        ProverHelpers.proveStorageBytes32(
            abi.encodePacked(_args.storageSlot), _args.storageValue, _l2StorageProof, bytes32(storageRoot)
        );
    }
}
