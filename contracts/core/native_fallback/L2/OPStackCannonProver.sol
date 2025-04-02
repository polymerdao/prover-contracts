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

import {L2Configuration} from "../../../libs/RegistryTypes.sol";
import {SecureMerkleTrie} from "@eth-optimism/contracts-bedrock/src/libraries/trie/SecureMerkleTrie.sol";
import {RLPReader} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPReader.sol";
import {RLPWriter} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPWriter.sol";
import {ISettledStateProver} from "../../../interfaces/ISettledStateProver.sol";

contract OPStackCannonProver is ISettledStateProver {
    struct DisputeGameFactoryProofData {
        bytes32 messagePasserStateRoot;
        bytes32 latestBlockHash;
        uint256 gameIndex;
        bytes32 gameId;
        bytes[] disputeFaultGameStorageProof;
        bytes rlpEncodedDisputeGameFactoryData;
        bytes[] disputeGameFactoryAccountProof;
    }

    struct FaultDisputeGameStatusSlotData {
        uint64 createdAt;
        uint64 resolvedAt;
        uint8 gameStatus;
        bool initialized;
        bool l2BlockNumberChallenged;
    }

    struct FaultDisputeGameProofData {
        bytes32 faultDisputeGameStateRoot;
        bytes[] faultDisputeGameRootClaimStorageProof;
        FaultDisputeGameStatusSlotData faultDisputeGameStatusSlotData;
        bytes[] faultDisputeGameStatusStorageProof;
        bytes rlpEncodedFaultDisputeGameData;
        bytes[] faultDisputeGameAccountProof;
    }

    /**
     * @notice Invalid dispute game factory state root encoding
     * @param _disputeGameFactoryStateRoot Invalid state root
     */
    error IncorrectDisputeGameFactoryStateRoot(bytes _disputeGameFactoryStateRoot);

    /**
     * @notice Fault dispute game not yet resolved
     * @param _gameStatus Current game status
     */
    error FaultDisputeGameUnresolved(uint8 _gameStatus);

    /**
     * @notice Failed account proof verification
     * @param _address Account address
     * @param _data Account data
     * @param _proof Merkle proof
     * @param _root Expected root
     */
    error InvalidAccountProof(bytes _address, bytes _data, bytes[] _proof, bytes32 _root);

    /**
     * @notice Failed storage proof verification
     * @param _key Storage key
     * @param _val Storage value
     * @param _proof Merkle proof
     * @param _root Expected root
     */
    error InvalidStorageProof(bytes _key, bytes _val, bytes[] _proof, bytes32 _root);

    /**
     * @notice RLP encoded block data hash mismatch
     * @param _expectedBlockHash Expected hash
     * @param _calculatedBlockHash Actual hash
     */
    error InvalidRLPEncodedBlock(bytes32 _expectedBlockHash, bytes32 _calculatedBlockHash);

    /**
     * @notice Validates RLP encoded block data matches expected hash
     * @param _rlpEncodedBlockData Encoded block data
     * @param _expectedBlockHash Expected block hash
     */
    modifier validRLPEncodeBlock(bytes memory _rlpEncodedBlockData, bytes32 _expectedBlockHash) {
        bytes32 calculatedBlockHash = keccak256(_rlpEncodedBlockData);
        if (calculatedBlockHash == _expectedBlockHash) {
            _;
        } else {
            revert InvalidRLPEncodedBlock(_expectedBlockHash, calculatedBlockHash);
        }
    }

    /**
     * @notice Handles Cannon L2 world state validation
     * @dev Verifies through fault dispute game resolution
     * @param _l2WorldStateRoot L2 state root
     * @param _rlpEncodedL2Header RLP encoded block header for that L2 state
     * @param _l1WorldStateRoot Proven L1 world state root
     * @param _proof rest of the proof data: DisputeGameFactoryProofData, FaultDisputeGameProofData
     */
    function proveSettledState(
        L2Configuration memory _chainConfig,
        bytes32 _l2WorldStateRoot,
        bytes memory _rlpEncodedL2Header,
        bytes32 _l1WorldStateRoot,
        bytes calldata _proof
    ) external pure returns (bool) {
        (
            DisputeGameFactoryProofData memory disputeGameFactoryProofData,
            FaultDisputeGameProofData memory faultDisputeGameProofData
        ) = abi.decode(_proof, (DisputeGameFactoryProofData, FaultDisputeGameProofData));

        require(
            _proveWorldStateCannon(
                _chainConfig,
                _l2WorldStateRoot,
                _rlpEncodedL2Header,
                disputeGameFactoryProofData,
                faultDisputeGameProofData,
                _l1WorldStateRoot
            ),
            "Invalid Bedrock proof"
        );
        return true;
    }

    /**
     * @notice Proves L2 world state using Cannon verification
     * @dev Verifies through fault dispute game resolution
     * @param _chainConfig L2 chain configuration
     * @param _rlpEncodedL2Header RLP encoded L2 header
     * @param _disputeGameFactoryProofData Proof data for factory verification
     * @param _faultDisputeGameProofData Proof data for game verification
     * @param _l1WorldStateRoot Proven L1 world state root
     */
    function _proveWorldStateCannon(
        L2Configuration memory _chainConfig,
        bytes32 _l2WorldStateRoot,
        bytes memory _rlpEncodedL2Header,
        DisputeGameFactoryProofData memory _disputeGameFactoryProofData,
        FaultDisputeGameProofData memory _faultDisputeGameProofData,
        bytes32 _l1WorldStateRoot // trusted at this point
    )
        internal
        pure
        validRLPEncodeBlock(_rlpEncodedL2Header, _disputeGameFactoryProofData.latestBlockHash)
        returns (bool)
    {
        // prove that the FaultDisputeGame was created by the Dispute Game Factory

        // Verify dispute game creation and resolution
        bytes32 rootClaim;
        address faultDisputeGameProxyAddress;
        (faultDisputeGameProxyAddress, rootClaim) = _faultDisputeGameFromFactory(
            _chainConfig.versionNumber,
            _chainConfig.addresses[0], // For OPStackCannon the factory address is stored in the first address slot
            _chainConfig.storageSlots[0], // For OPStackCannon the disputeGameFactoryListSlot is the 1st storage slot in
                // the config
            _l2WorldStateRoot,
            _disputeGameFactoryProofData,
            _l1WorldStateRoot
        );

        _faultDisputeGameIsResolved(
            rootClaim,
            faultDisputeGameProxyAddress,
            _chainConfig.storageSlots[1], // For OPStackCannon faultDisputeGameRootClaimSlot is the 2nd storage slot in
                // the config
            _chainConfig.storageSlots[2], // For OPStackCannon faultDisputeGameStatusSlot is the 3rd storage slot in the
                // config
            _faultDisputeGameProofData,
            _l1WorldStateRoot
        );

        return true;
    }

    /**
     * @notice Verifies fault dispute game resolution
     * @dev Verifies game status and root claim
     * @param _rootClaim Expected root claim value
     * @param _faultDisputeGameProxyAddress Game proxy contract
     * @param _faultDisputeGameRootClaimSlot Storage slot for the root claim
     * @param _faultDisputeGameStatusSlot Storage slot for the game status
     * @param _faultDisputeGameProofData Proof data for game verification
     * @param _l1WorldStateRoot Proven L1 world state root
     */
    function _faultDisputeGameIsResolved(
        bytes32 _rootClaim,
        address _faultDisputeGameProxyAddress,
        uint256 _faultDisputeGameRootClaimSlot, // 2nd storage slot in the config
        uint256 _faultDisputeGameStatusSlot, // 3rd storage slot in the config
        FaultDisputeGameProofData memory _faultDisputeGameProofData,
        bytes32 _l1WorldStateRoot
    ) internal pure {
        // Verify game is resolved
        if (_faultDisputeGameProofData.faultDisputeGameStatusSlotData.gameStatus != 2) {
            revert FaultDisputeGameUnresolved(_faultDisputeGameProofData.faultDisputeGameStatusSlotData.gameStatus);
        }

        // ensure faultDisputeGame is resolved
        // Prove that the FaultDispute game has been settled
        // storage proof for FaultDisputeGame rootClaim (means block is valid)
        _proveStorageBytes32(
            abi.encodePacked(uint256(_faultDisputeGameRootClaimSlot)),
            _rootClaim,
            _faultDisputeGameProofData.faultDisputeGameRootClaimStorageProof,
            bytes32(_faultDisputeGameProofData.faultDisputeGameStateRoot)
        );

        // Assemble and verify game status
        bytes32 faultDisputeGameStatusStorage = _assembleGameStatusStorage(
            _faultDisputeGameProofData.faultDisputeGameStatusSlotData.createdAt,
            _faultDisputeGameProofData.faultDisputeGameStatusSlotData.resolvedAt,
            _faultDisputeGameProofData.faultDisputeGameStatusSlotData.gameStatus,
            _faultDisputeGameProofData.faultDisputeGameStatusSlotData.initialized,
            _faultDisputeGameProofData.faultDisputeGameStatusSlotData.l2BlockNumberChallenged
        );

        // Verify game status storage proof
        _proveStorageBytes32(
            abi.encodePacked(uint256(_faultDisputeGameStatusSlot)),
            faultDisputeGameStatusStorage,
            _faultDisputeGameProofData.faultDisputeGameStatusStorageProof,
            bytes32(
                RLPReader.readBytes(RLPReader.readList(_faultDisputeGameProofData.rlpEncodedFaultDisputeGameData)[2])
            )
        );

        // Verify game contract account proof
        _proveAccount(
            abi.encodePacked(_faultDisputeGameProxyAddress),
            _faultDisputeGameProofData.rlpEncodedFaultDisputeGameData,
            _faultDisputeGameProofData.faultDisputeGameAccountProof,
            _l1WorldStateRoot
        );
    }

    /**
     * @notice Validates fault dispute game from factory configuration
     * @dev Internal helper for Cannon proving
     * @param _disputeGameFactoryAddress Factory contract address
     * @param _l2WorldStateRoot L2 state root to verify
     * @param _disputeGameFactoryProofData Proof data for factory validation
     * @param _l1WorldStateRoot Proven L1 world state root
     * @return faultDisputeGameProxyAddress Address of game proxy
     * @return rootClaim Generated root claim
     */
    function _faultDisputeGameFromFactory(
        uint256 _outputVersion,
        address _disputeGameFactoryAddress, // 1st address in the config
        uint256 _disputeGameFactoryListSlot, // 1st storage slot in the config
        bytes32 _l2WorldStateRoot,
        DisputeGameFactoryProofData memory _disputeGameFactoryProofData,
        bytes32 _l1WorldStateRoot
    ) internal pure returns (address faultDisputeGameProxyAddress, bytes32 rootClaim) {
        // Generate root claim from state data
        bytes32 _rootClaim = _generateOutputRoot(
            _outputVersion,
            _l2WorldStateRoot,
            _disputeGameFactoryProofData.messagePasserStateRoot,
            _disputeGameFactoryProofData.latestBlockHash
        );

        // Verify game exists in factory
        bytes32 disputeGameFactoryStorageSlot = bytes32(
            abi.encode(
                (uint256(keccak256(abi.encode(_disputeGameFactoryListSlot))) + _disputeGameFactoryProofData.gameIndex)
            )
        );

        bytes memory disputeGameFactoryStateRoot =
            RLPReader.readBytes(RLPReader.readList(_disputeGameFactoryProofData.rlpEncodedDisputeGameFactoryData)[2]);

        if (disputeGameFactoryStateRoot.length > 32) {
            revert IncorrectDisputeGameFactoryStateRoot(disputeGameFactoryStateRoot);
        }

        // Verify storage and account proofs
        _proveStorageBytes32(
            abi.encodePacked(disputeGameFactoryStorageSlot),
            _disputeGameFactoryProofData.gameId,
            _disputeGameFactoryProofData.disputeFaultGameStorageProof,
            bytes32(disputeGameFactoryStateRoot)
        );

        _proveAccount(
            abi.encodePacked(_disputeGameFactoryAddress),
            _disputeGameFactoryProofData.rlpEncodedDisputeGameFactoryData,
            _disputeGameFactoryProofData.disputeGameFactoryAccountProof,
            _l1WorldStateRoot
        );

        // Get proxy address from game ID
        (,, address _faultDisputeGameProxyAddress) = _unpackGameID(_disputeGameFactoryProofData.gameId);

        return (_faultDisputeGameProxyAddress, _rootClaim);
    }

    /**
     * @notice Generates an output root for Bedrock and Cannon proving
     * @param _outputRootVersion Version number (usually 0)
     * @param _worldStateRoot State root
     * @param _messagePasserStateRoot Message passer state root
     * @param _latestBlockHash Latest block hash
     * @return Output root hash
     */
    function _generateOutputRoot(
        uint256 _outputRootVersion,
        bytes32 _worldStateRoot,
        bytes32 _messagePasserStateRoot,
        bytes32 _latestBlockHash
    ) internal pure returns (bytes32) {
        return keccak256(abi.encode(_outputRootVersion, _worldStateRoot, _messagePasserStateRoot, _latestBlockHash));
    }

    /**
     * @notice Unpacks a 32-byte GameId into its components
     * @param _gameId Packed game identifier
     * @return gameType_ Game type identifier
     * @return timestamp_ Creation timestamp
     * @return gameProxy_ Proxy contract address
     */
    function _unpackGameID(bytes32 _gameId)
        internal
        pure
        returns (uint32 gameType_, uint64 timestamp_, address gameProxy_)
    {
        assembly {
            gameType_ := shr(224, _gameId)
            timestamp_ := and(shr(160, _gameId), 0xFFFFFFFFFFFFFFFF)
            gameProxy_ := and(_gameId, 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)
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
     * @notice Assembles game status storage slot data
     * @dev Packs status fields into a single bytes32
     * @param _createdAt Creation timestamp
     * @param _resolvedAt Resolution timestamp
     * @param _gameStatus Game status code
     * @param _initialized Initialization status
     * @param _l2BlockNumberChallenged Block number challenge status
     * @return gameStatusStorageSlotRLP Packed status data
     */
    function _assembleGameStatusStorage(
        uint64 _createdAt,
        uint64 _resolvedAt,
        uint8 _gameStatus,
        bool _initialized,
        bool _l2BlockNumberChallenged
    ) internal pure returns (bytes32 gameStatusStorageSlotRLP) {
        // Packed data is 64 + 64 + 8 + 8 + 8 = 152 bits / 19 bytes.
        // Need to convert to `uint152` to preserve right alignment.
        return bytes32(
            uint256(
                uint152(
                    bytes19(
                        abi.encodePacked(_l2BlockNumberChallenged, _initialized, _gameStatus, _resolvedAt, _createdAt)
                    )
                )
            )
        );
    }
}
