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
import {RLPWriter} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPWriter.sol";
import {IL1Block} from "../../../interfaces/IL1Block.sol";
import {INativeProver} from "../../../interfaces/INativeProver.sol";
import {IProverHelper} from "../../../interfaces/IProverHelper.sol";
import {ISettledStateProver} from "../../../interfaces/ISettledStateProver.sol";
import {L2Configuration, L1Configuration, Type, ProveScalarArgs} from "../../../libs/RegistryTypes.sol";

contract NativeProver is INativeProver, IProverHelper {
    uint256 public CHAIN_ID; // ChainID of the L2 chain this contract is deployed on

    L1Configuration public L1_CONFIGURATION; // Configuration for L1

    mapping(uint256 => L2Configuration) public l2ChainConfigurations; // Mapping of counterparty chainIDs to their
        // configurations

    mapping(uint256 => BlockProof) public provenStates;

    mapping(uint256 => ISettledStateProver) public stateProvers;

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

    /**
     * @notice Emitted when L1 world state is successfully proven
     * @param _blockNumber Block number of proven state
     * @param _L1WorldStateRoot World state root that was proven
     */
    event L1WorldStateProven(uint256 indexed _blockNumber, bytes32 _L1WorldStateRoot);

    /**
     * @notice Emitted when L2 world state is successfully proven
     * @param _destinationChainID Chain ID of the L2
     * @param _blockNumber Block number of proven state
     * @param _L2WorldStateRoot World state root that was proven
     */
    event L2WorldStateProven(
        uint256 indexed _destinationChainID, uint256 indexed _blockNumber, bytes32 _L2WorldStateRoot
    );

    /**
     * @notice Block number is too recent to prove
     * @param _inputBlockNumber Block attempted to prove
     * @param _nextProvableBlockNumber Next valid block number
     */
    error NeedLaterBlock(uint256 _inputBlockNumber, uint256 _nextProvableBlockNumber);

    /**
     * @notice Block number is older than currently proven block
     * @param _inputBlockNumber Block attempted to prove
     * @param _latestBlockNumber Current proven block
     */
    error OutdatedBlock(uint256 _inputBlockNumber, uint256 _latestBlockNumber);

    /**
     * @notice RLP encoded block data hash mismatch
     * @param _expectedBlockHash Expected hash
     * @param _calculatedBlockHash Actual hash
     */
    error InvalidRLPEncodedBlock(bytes32 _expectedBlockHash, bytes32 _calculatedBlockHash);

    /**
     * @notice Failed storage proof verification
     * @param _key Storage key
     * @param _val Storage value
     * @param _proof Merkle proof
     * @param _root Expected root
     */
    error InvalidStorageProof(bytes _key, bytes _val, bytes[] _proof, bytes32 _root);

    /**
     * @notice Failed account proof verification
     * @param _address Account address
     * @param _data Account data
     * @param _proof Merkle proof
     * @param _root Expected root
     */
    error InvalidAccountProof(bytes _address, bytes _data, bytes[] _proof, bytes32 _root);

    /**
     * @notice Settlement chain state not yet proven
     * @param _blockProofStateRoot State root attempted to prove
     * @param _l1WorldStateRoot Current proven state root
     */
    error SettlementChainStateRootNotProved(bytes32 _blockProofStateRoot, bytes32 _l1WorldStateRoot);

    /**
     * @notice Destination chain state not yet proven
     * @param _blockProofStateRoot State root attempted to prove
     * @param _l2WorldStateRoot Current proven state root
     */
    error DestinationChainStateRootNotProved(bytes32 _blockProofStateRoot, bytes32 _l2WorldStateRoot);

    /**
     * @notice Invalid storage root found in contract
     * @param _contractStorageRoot Invalid storage root
     */
    error IncorrectContractStorageRoot(bytes _contractStorageRoot);

    /**
     * @notice Validates RLP encoded block data matches expected hash
     * @param _rlpEncodedHeader Encoded block data
     * @param _expectedBlockHash Expected block hash
     */
    modifier validRLPEncodeBlock(bytes memory _rlpEncodedHeader, bytes32 _expectedBlockHash) {
        bytes32 calculatedBlockHash = keccak256(_rlpEncodedHeader);
        if (calculatedBlockHash == _expectedBlockHash) {
            _;
        } else {
            revert InvalidRLPEncodedBlock(_expectedBlockHash, calculatedBlockHash);
        }
    }

    constructor(
        uint256 _chainID,
        L1Configuration memory _l1Configuration,
        InitialL2Configuration[] memory _initialL2Configurations
    ) {
        L1_CONFIGURATION = _l1Configuration;
        CHAIN_ID = _chainID;
        for (uint256 i = 0; i < _initialL2Configurations.length; ++i) {
            _setInitialChainConfiguration(_initialL2Configurations[i].chainID, _initialL2Configurations[i].config);
        }
    }

    function _setInitialChainConfiguration(uint256 _chainID, L2Configuration memory _config) internal {
        l2ChainConfigurations[_chainID] = _config;
    }

    /**
     * @notice Updates an L2 configuration using the provided L1 registry proof
     * @dev Validates the L2 configuration against the L1 registry contract's state using the provided proof
     * @param _chainID chain ID of the L2 configuration being proven
     * @param _config L2 configuration being proven
     * @param _l1StorageProof proof of the config hash in the registry contract
     * @param _rlpEncodedRegistryAccountData RLP encoded registry contract account data
     * @param _l1RegistryProof proof of the registry contract account
     * @param _l1WorldStateRoot L1 world state root
     *
     */
    function updateL2ChainConfiguration(
        uint256 _chainID,
        L2Configuration calldata _config,
        bytes[] calldata _l1StorageProof,
        bytes calldata _rlpEncodedRegistryAccountData,
        bytes[] calldata _l1RegistryProof,
        bytes32 _l1WorldStateRoot
    ) external {
        require(
            _proveL2Configuration(
                _chainID, _config, _l1StorageProof, _rlpEncodedRegistryAccountData, _l1RegistryProof, _l1WorldStateRoot
            ),
            "Invalid L2Configuration proof"
        );
        l2ChainConfigurations[_chainID] = _config;
    }

    /**
     * @notice Updates an L1 configuration using the provided L1 registry proof
     * @dev Validates the L1 configuration against the L1 registry contract's state using the provided proof
     * @param _config L1 configuration being proven
     * @param _l1StorageProof proof of the config hash in the registry contract
     * @param _rlpEncodedRegistryAccountData RLP encoded registry contract account data
     * @param _l1RegistryProof proof of the registry contract account
     * @param _l1WorldStateRoot L1 world state root
     *
     */
    function updateL1ChainConfiguration(
        L1Configuration calldata _config,
        bytes[] calldata _l1StorageProof,
        bytes calldata _rlpEncodedRegistryAccountData,
        bytes[] calldata _l1RegistryProof,
        bytes32 _l1WorldStateRoot
    ) external {
        require(
            _proveL1Configuration(
                _config, _l1StorageProof, _rlpEncodedRegistryAccountData, _l1RegistryProof, _l1WorldStateRoot
            ),
            "Invalid L1Configuration proof"
        );
        L1_CONFIGURATION = _config;
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
        BlockProof memory existingSettlementBlockProof = provenStates[CHAIN_ID];

        // Verify settlement chain state root
        if (existingSettlementBlockProof.stateRoot != _l1WorldStateRoot) {
            revert SettlementChainStateRootNotProved(existingSettlementBlockProof.stateRoot, _l1WorldStateRoot);
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
        _proveStorageBytes32(
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
        BlockProof memory existingSettlementBlockProof = provenStates[CHAIN_ID];

        // Verify settlement chain state root
        if (existingSettlementBlockProof.stateRoot != _l1WorldStateRoot) {
            revert SettlementChainStateRootNotProved(existingSettlementBlockProof.stateRoot, _l1WorldStateRoot);
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
            bytes32((uint256(keccak256(abi.encode(CHAIN_ID))) + L1_CONFIGURATION.settlementRegistryL1ConfigMappingSlot));
        _proveStorageBytes32(
            abi.encodePacked(configHashStorageSlot), configHash, _l1StorageProof, bytes32(registryStorageRoot)
        );

        return true;
    }

    /**
     * @notice Validates a storage proof against a root
     * @dev Uses SecureMerkleTrie for verification
     * @param _key Storage slot key
     * @param _val Expected value
     * @param _proof Merkle proof
     * @param _root Expected root
     */
    function proveStorage(bytes memory _key, bytes memory _val, bytes[] memory _proof, bytes32 _root) external pure {
        if (!SecureMerkleTrie.verifyInclusionProof(_key, _val, _proof, _root)) {
            revert InvalidStorageProof(_key, _val, _proof, _root);
        }
    }

    /**
     * @notice Validates a bytes32 storage value against a root
     * @dev Encodes value as RLP before verification
     * @param _key Storage slot key
     * @param _val Expected bytes32 value
     * @param _proof Merkle proof
     * @param _root Expected root
     */
    function _proveStorageBytes32(bytes memory _key, bytes32 _val, bytes[] memory _proof, bytes32 _root)
        internal
        pure
    {
        // `RLPWriter.writeUint` properly encodes values by removing any leading zeros.
        bytes memory rlpEncodedValue = RLPWriter.writeUint(uint256(_val));
        if (!SecureMerkleTrie.verifyInclusionProof(_key, rlpEncodedValue, _proof, _root)) {
            revert InvalidStorageProof(_key, rlpEncodedValue, _proof, _root);
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
     * @notice RLP encodes a list of data elements
     * @dev Helper function for batch encoding
     * @param _dataList List of data elements to encode
     * @return RLP encoded bytes
     */
    function rlpEncodeDataLibList(bytes[] memory _dataList) external pure returns (bytes memory) {
        for (uint256 i = 0; i < _dataList.length; ++i) {
            _dataList[i] = RLPWriter.writeBytes(_dataList[i]);
        }

        return RLPWriter.writeList(_dataList);
    }

    /**
     * @notice Packs game metadata into a 32-byte GameId
     * @dev Combines type, timestamp, and proxy address into single identifier
     * @param _gameType Game type identifier
     * @param _timestamp Creation timestamp
     * @param _gameProxy Proxy contract address
     * @return gameId_ Packed game identifier
     */
    function packGameID(uint32 _gameType, uint64 _timestamp, address _gameProxy)
        external
        pure
        returns (bytes32 gameId_)
    {
        assembly {
            gameId_ := or(or(shl(224, _gameType), shl(160, _timestamp)), _gameProxy)
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
     * @notice Proves L1 settlement layer state against oracle
     * @dev Validates block data against L1 blockhash oracle and updates proven state
     * @param _rlpEncodedL1Header RLP encoded block data
     */
    function proveSettlementLayerState(bytes calldata _rlpEncodedL1Header)
        public
        validRLPEncodeBlock(_rlpEncodedL1Header, IL1Block(L1_CONFIGURATION.blockHashOracle).hash())
    {
        uint256 settlementChainId = CHAIN_ID;
        // not necessary because we already confirm that the data is correct by ensuring that it hashes to the block
        // hash
        // require(l1WorldStateRoot.length <= 32); // ensure lossless casting to bytes32

        // Extract block proof from encoded data
        BlockProof memory blockProof = BlockProof({
            blockNumber: _bytesToUint(RLPReader.readBytes(RLPReader.readList(_rlpEncodedL1Header)[8])),
            blockHash: keccak256(_rlpEncodedL1Header),
            stateRoot: bytes32(RLPReader.readBytes(RLPReader.readList(_rlpEncodedL1Header)[3]))
        });

        // Verify block delay and update state
        BlockProof memory existingBlockProof = provenStates[settlementChainId];
        if (existingBlockProof.blockNumber + L1_CONFIGURATION.settlementBlocksDelay < blockProof.blockNumber) {
            provenStates[settlementChainId] = blockProof;
            emit L1WorldStateProven(blockProof.blockNumber, blockProof.stateRoot);
        } else {
            revert NeedLaterBlock(
                blockProof.blockNumber, existingBlockProof.blockNumber + L1_CONFIGURATION.settlementBlocksDelay
            );
        }
    }

    /**
     * @notice Proves a settled L2 state
     * @dev Validates a L2 settlement proof depending on the type of the L2 and its configuration
     * @param _chainID L2 chain ID
     * @param _l2WorldStateRoot L2 state root we want to prove
     * @param _rlpEncodedL2Header RLP encoded L2 block header for that state
     * @param _proof RLP encoded proof data
     */
    function proveSettledState(
        uint256 _chainID,
        bytes32 _l2WorldStateRoot,
        bytes memory _rlpEncodedL2Header,
        bytes32 _l1WorldStateRoot,
        bytes calldata _proof
    ) public {
        L2Configuration memory conf = l2ChainConfigurations[_chainID];

        // Verify settlement chain state root
        BlockProof memory existingSettlementBlockProof = provenStates[CHAIN_ID];
        if (existingSettlementBlockProof.stateRoot != _l1WorldStateRoot) {
            revert SettlementChainStateRootNotProved(existingSettlementBlockProof.stateRoot, _l1WorldStateRoot);
        }

        // Call out to the configured prover to verify proof of the settled L2 state root
        require(
            ISettledStateProver(conf.prover).proveSettledState(
                conf, _l2WorldStateRoot, _rlpEncodedL2Header, _l1WorldStateRoot, _proof
            ),
            "Invalid settled state proof"
        );

        // Update proven state if newer block
        BlockProof memory existingBlockProof = provenStates[_chainID];
        BlockProof memory blockProof = BlockProof({
            blockNumber: _bytesToUint(RLPReader.readBytes(RLPReader.readList(_rlpEncodedL2Header)[8])),
            blockHash: keccak256(_rlpEncodedL2Header),
            stateRoot: _l2WorldStateRoot
        });

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
     * @notice Proves a storage value in a settled L2 state root
     * @dev Verifies a storage value against a verified settle L2 state root
     * @param _args holds the ProveScalarArgs
     * @param _l2StorageProof Proof data for storage value verification
     * @param _rlpEncodedContractState RLP encoded contract state
     * @param _l2AccountProof Proof data for storing contract on the L2
     * @return chainID Destination chain ID
     * @return storingContract Address of storing contract
     * @return storageValue Verified storage value
     */
    function proveStorageValue(
        ProveScalarArgs calldata _args,
        bytes[] calldata _l2StorageProof,
        bytes calldata _rlpEncodedContractState,
        bytes[] calldata _l2AccountProof
    ) public view returns (uint256 chainID, address storingContract, bytes32 storageValue) {
        // Verify L2 state root is proven
        BlockProof memory existingBlockProof = provenStates[_args.chainID];
        if (existingBlockProof.stateRoot != _args.l2WorldStateRoot) {
            revert DestinationChainStateRootNotProved(existingBlockProof.stateRoot, _args.l2WorldStateRoot);
        }

        bytes memory storageRoot = RLPReader.readBytes(RLPReader.readList(_rlpEncodedContractState)[2]);

        if (storageRoot.length > 32) {
            revert IncorrectContractStorageRoot(storageRoot);
        }

        // Verify the account exists in the state tree
        _proveAccount(
            abi.encodePacked(_args.contractAddr), _rlpEncodedContractState, _l2AccountProof, _args.l2WorldStateRoot
        );

        // Verify the storage value exists in the storage tree
        _proveStorageBytes32(
            abi.encodePacked(_args.storageSlot), _args.storageValue, _l2StorageProof, bytes32(storageRoot)
        );

        return (_args.chainID, _args.contractAddr, _args.storageValue);
    }

    /**
     * @notice Proves a storage value in a settled L2 state root
     * @dev Single view function encapsulating all three steps of the proof process: L1 view; settled L2 state; L2
     * storage
     * @param _args holds the ProveScalarArgs
     * @param _rlpEncodedL1Header L1 header to be verified against L1 block hash oracle to access L1 state root
     * @param _rlpEncodedL2Header L2 header proven in the L1 settlement
     * @param _settledStateProof proof of the L2 world state root in the L1 settlement
     * @param _l2StorageProof proof of the storage value in the L2 _contractAddr
     * @param _rlpEncodedContractAccount RLP encoded _contractAddr account
     * @param _l2AccountProof proof of the _contractAddr account in the L2 world state
     *
     */
    function prove(
        ProveScalarArgs calldata _args,
        bytes calldata _rlpEncodedL1Header,
        bytes memory _rlpEncodedL2Header,
        bytes calldata _settledStateProof,
        bytes[] calldata _l2StorageProof,
        bytes calldata _rlpEncodedContractAccount,
        bytes[] calldata _l2AccountProof
    ) external view returns (uint256 chainID, address storingContract, bytes32 storageValue) {
        // First prove the settled state
        _proveSettledState(_args, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof);

        // Now prove storage against the verified settled L2 state root
        _proveStorageInState(_args, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof);

        return (_args.chainID, _args.contractAddr, _args.storageValue);
    }

    function _proveSettledState(
        ProveScalarArgs calldata _args,
        bytes calldata _rlpEncodedL1Header,
        bytes memory _rlpEncodedL2Header,
        bytes calldata _settledStateProof
    ) internal view validRLPEncodeBlock(_rlpEncodedL1Header, IL1Block(L1_CONFIGURATION.blockHashOracle).hash()) {
        bytes32 _l1StateRoot = bytes32(RLPReader.readBytes(RLPReader.readList(_rlpEncodedL1Header)[3]));

        // Get the L2Configuration for this chainID
        L2Configuration memory conf = l2ChainConfigurations[_args.chainID];

        // Call out to the configured prover to verify proof of the settled L2 state root
        require(
            ISettledStateProver(conf.prover).proveSettledState(
                conf, _args.l2WorldStateRoot, _rlpEncodedL2Header, _l1StateRoot, _settledStateProof
            ),
            "Invalid settled state proof"
        );
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
        _proveStorageBytes32(
            abi.encodePacked(_args.storageSlot), _args.storageValue, _l2StorageProof, bytes32(storageRoot)
        );
    }
}
