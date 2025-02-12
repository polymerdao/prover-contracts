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

import {ReceiptParser} from "../libs/ReceiptParser.sol";
import {ICrossL2ProverV2} from "../interfaces/ICrossL2ProverV2.sol";
import {LightClientType} from "../interfaces/IClientUpdates.sol";
import {SequencerSignatureVerifierV2} from "./SequencerSignatureVerifierV2.sol";

contract CrossL2ProverV2 is SequencerSignatureVerifierV2, ICrossL2ProverV2 {
    LightClientType public constant LIGHT_CLIENT_TYPE = LightClientType.SequencerLightClient; // Stored as a constant
        // for cheap on-chain use

    string public clientType;

    error InvalidProofKey();
    error InvalidProofValue();
    error InvalidProofRoot();

    constructor(string memory clientType_, address sequencer_, bytes32 chainId_)
        SequencerSignatureVerifierV2(sequencer_, chainId_)
    {
        clientType = clientType_;
    }

    /**
     * @inheritdoc ICrossL2ProverV2
     */
    function validateEvent(bytes calldata proof)
        external
        view
        returns (uint32 chainId, address emittingContract, bytes memory topics, bytes memory unindexedData)
    {
        // Proof calldata is encoded in the following format:
        //  +--------------------------------------------------+
        //  |  state root (32 bytes)                           | index: 0:32
        //  +--------------------------------------------------+
        //  |  signature (65 bytes)                            | index: 32:97
        //  +--------------------------------------------------+
        //  |  source chain ID (big endian, 4 bytes)           | index: 97:101
        //  +--------------------------------------------------+
        //  |  peptide height (big endian, 8 bytes)            | index: 101:109
        //  +--------------------------------------------------+
        //  |  source chain block height (big endian, 8 bytes) | index: 109:117
        //  +--------------------------------------------------+
        //  |  receipt index (big endian, 2 bytes)             | index: 117:119
        //  +--------------------------------------------------+
        //  |  event index (1 byte)                            | index: 119
        //  +--------------------------------------------------+
        //  |  number of topics (1 byte)                       | index: 120
        //  +--------------------------------------------------+
        //  |  event data end (big endian, 2 bytes)            | index: 121:123
        //  +--------------------------------------------------+
        //  |  event emitter (contract address) (20 bytes)     | index: 123:143
        //  +--------------------------------------------------+
        //  |  topics (32 bytes * number of topics)            | index: 143 + 32 * number of topics: eventDatEnd
        //  +--------------------------------------------------+
        //  |  event data (x bytes)                            | index: 123:eventDataEnd
        //  +--------------------------------------------------+
        //  |  iavl proof (x bytes)                            | index: eventDataEnd:
        //  +--------------------------------------------------+

        chainId = uint32(bytes4(proof[97:101]));
        _verifySequencerSignature(
            bytes32(proof[:32]),
            uint64(bytes8(proof[101:109])),
            uint8(proof[96]),
            bytes32(proof[32:64]),
            bytes32(proof[64:96])
        );

        uint256 eventEnd = uint16(bytes2(proof[121:123]));
        bytes memory rawEvent = proof[123:eventEnd];
        this.verifyMembership(
            bytes32(proof[:32]),
            ReceiptParser.eventRootKey(
                chainId, clientType, uint64(bytes8(proof[109:117])), uint16(bytes2(proof[117:119])), uint8(proof[119])
            ),
            keccak256(rawEvent),
            proof[eventEnd:]
        );

        (emittingContract, topics, unindexedData) = this.parseEvent(rawEvent, uint8(proof[120]));
    }

    /**
     * @inheritdoc ICrossL2ProverV2
     */
    function inspectLogIdentifier(bytes calldata proof)
        external
        pure
        returns (uint32 srcChain, uint64 blockNumber, uint16 receiptIndex, uint8 logIndex)
    {
        return (
            uint32(bytes4(proof[97:101])),
            uint64(bytes8(proof[109:117])),
            uint16(bytes2(proof[117:119])),
            uint8(proof[119])
        );
    }

    /**
     * @inheritdoc ICrossL2ProverV2
     */
    function inspectPolymerState(bytes calldata proof)
        external
        pure
        returns (bytes32 stateRoot, uint64 height, bytes calldata signature)
    {
        return (bytes32(proof[:32]), uint64(bytes8(proof[101:109])), proof[32:97]);
    }

    /**
     * @inheritdoc ICrossL2ProverV2
     */
    function verifyMembership(bytes32 root, bytes memory key, bytes32 value, bytes calldata proof) public pure {
        uint256 path0start = uint256(uint8(proof[1]));
        bytes32 prehash = sha256(abi.encodePacked(proof[2:path0start], key, hex"20", sha256(abi.encodePacked(value))));
        uint256 offset = path0start;

        for (uint256 i = 0; i < uint256(uint8(proof[0])); i++) {
            uint256 suffixstart = uint256(uint8(proof[offset]));
            uint256 suffixend = uint256(uint8(proof[offset + 1]));

            // add +2 to account for path header
            prehash = sha256(
                abi.encodePacked(
                    proof[offset + 2:offset + suffixstart], prehash, proof[offset + suffixstart:offset + suffixend]
                )
            );

            offset = offset + suffixend;
        }

        if (prehash != root) revert InvalidProofRoot();
    }

    /**
     * @inheritdoc ICrossL2ProverV2
     */
    function parseEvent(bytes calldata rawEvent, uint8 numTopics)
        public
        pure
        returns (address emittingContract, bytes memory topics, bytes memory unindexedData)
    {
        uint256 topicsEnd = 32 * numTopics + 20;
        return (address(bytes20(rawEvent[:20])), rawEvent[20:topicsEnd], rawEvent[topicsEnd:]);
    }
}
