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

import {ReceiptParser} from "../../libs/ReceiptParser.sol";
import {ICrossL2ProverV2} from "../../interfaces/ICrossL2ProverV2.sol";
import {LightClientType} from "../../interfaces/IClientUpdates.sol";
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
    //  +--------------------------------------------------+
    //  |  state root (32 bytes)                           | 0:32
    //  +--------------------------------------------------+
    //  |  signature (65 bytes)                            | 32:97
    //  +--------------------------------------------------+
    //  |  source chain ID (big endian, 4 bytes)           | 97:101
    //  +--------------------------------------------------+
    //  |  peptide height (big endian, 8 bytes)            | 101:109
    //  +--------------------------------------------------+
    //  |  source chain block height (big endian, 8 bytes) | 109:117
    //  +--------------------------------------------------+
    //  |  receipt index (big endian, 4 bytes)             | 117:121
    //  +--------------------------------------------------+
    //  |  event index (big endian, 4 bytes)               | 121:125
    //  +--------------------------------------------------+
    //  |  number of topics (1 byte)                       | 125
    //  +--------------------------------------------------+
    //  |  event data end (big endian, 2 bytes)            | 126:128
    //  +--------------------------------------------------+
    //  |  event emitter (contract address) (20 bytes)     | 128:148
    //  +--------------------------------------------------+
    //  |  topics (32 bytes * number of topics)            | 148 + 32 * number of topics: eventDatEnd
    //  +--------------------------------------------------+
    //  |  event data (x bytes)                            | eventDataEnd:
    //  +--------------------------------------------------+
    //  |  iavl proof (x bytes)                            |
    //  +--------------------------------------------------+

    function validateEvent(bytes calldata proof)
        external
        view
        virtual
        returns (uint32 chainId, address emittingContract, bytes memory topics, bytes memory unindexedData)
    {
        chainId = uint32(bytes4(proof[97:101]));
        _verifySequencerSignature(
            bytes32(proof[:32]),
            uint64(bytes8(proof[101:109])),
            uint8(proof[96]),
            bytes32(proof[32:64]),
            bytes32(proof[64:96])
        );

        uint256 eventEnd = uint16(bytes2(proof[126:128]));
        bytes memory rawEvent = proof[128:eventEnd];
        this.verifyMembership(
            bytes32(proof[:32]),
            ReceiptParser.eventRootKey(
                chainId,
                clientType,
                uint64(bytes8(proof[109:117])),
                uint32(bytes4(proof[117:121])),
                uint32(bytes4(proof[121:125]))
            ),
            keccak256(rawEvent),
            proof[eventEnd:]
        );

        (emittingContract, topics, unindexedData) = this.parseEvent(rawEvent, uint8(proof[125]));
    }

    // SOLANA DECODING
    //          +---------------------------------------------------+
    // 0:32     |  state root                 (32 bytes)            |
    //          +---------------------------------------------------+
    // 32:97    |  signature                  (65 bytes)            |
    //          +---------------------------------------------------+
    // 97:101   |  source chain ID            (big endian, 4 bytes) |
    //          +---------------------------------------------------+
    // 101:109  |  peptide height             (big endian, 8 bytes) |
    //          +---------------------------------------------------+
    // 109:117  |  source chain block height  (big endian, 8 bytes) |
    //          +---------------------------------------------------+
    // 117      |  number of log messages     (1 byte)              |
    //          +---------------------------------------------------+
    // 118:150  |  txSignature (high)         (32 bytes)            |
    //          +---------------------------------------------------+
    // 150:182  |  txSignature (low)          (32 bytes)            |
    //          +---------------------------------------------------+
    // 182:214  |  programID                  (32 bytes)            |
    //          +---------------------------------------------------+
    // 214:216  |  (currLogMsgDataEnd,logMsg) (2 bytes, X bytes)    |
    //          +---------------------------------------------------+
    //          |  iavl proof                 (x bytes)             |
    //          +---------------------------------------------------+
    function validateSolLogs(bytes calldata proof)
        external
        view
        virtual
        returns (uint32 chainId, bytes32 programID, string[] memory logMessages)
    {
        chainId = uint32(bytes4(proof[97:101]));

        _verifySequencerSignature(
            bytes32(proof[:32]), // apphash
            uint64(bytes8(proof[101:109])), // peptide height
            uint8(proof[96]), // signature: v component
            bytes32(proof[32:64]), // signature: r component
            bytes32(proof[64:96]) // signature: s component
        );
        programID = bytes32(proof[182:214]);

        uint256 currLogMessageStart = 214;
        uint256 currentLogMessageEnd = 214; // Edge case for 0 log messages

        logMessages = new string[](uint8((proof[117]))); // number of log messages

        for (uint256 i = 0; i < logMessages.length; ++i) {
            currentLogMessageEnd = uint16(bytes2(proof[currLogMessageStart:currLogMessageStart + 2]));
            logMessages[i] = string(proof[currLogMessageStart + 2:currentLogMessageEnd]);
            currLogMessageStart = currentLogMessageEnd;
        }

        bytes memory rawEvent = abi.encodePacked(programID);
        for (uint256 i = 0; i < uint8((proof[117])); ++i) {
            rawEvent = abi.encodePacked(rawEvent, logMessages[i]);
        }

        this.verifyMembership(
            bytes32(proof[:32]), // apphash
            ReceiptParser.solanaEventRootKey(
                chainId,
                clientType,
                uint64(bytes8(proof[109:117])), // height
                bytes32(proof[118:150]), // tx signature high
                bytes32(proof[150:182]), // tx signature low
                programID
            ),
            keccak256(rawEvent),
            proof[currentLogMessageEnd:]
        );
    }

    function inspectLogIdentifier(bytes calldata proof)
        external
        pure
        virtual
        returns (uint32 srcChain, uint64 blockNumber, uint32 receiptIndex, uint32 logIndex)
    {
        return (
            uint32(bytes4(proof[97:101])),
            uint64(bytes8(proof[109:117])),
            uint32(bytes4(proof[117:121])),
            uint32(bytes4(proof[121:125]))
        );
    }

    function inspectPolymerState(bytes calldata proof)
        external
        pure
        virtual
        returns (bytes32 stateRoot, uint64 height, bytes calldata signature)
    {
        return (bytes32(proof[:32]), uint64(bytes8(proof[101:109])), proof[32:97]);
    }

    /*
    header: key start (abs) (2B), key end (abs) (2B), value start (abs) (2B), value end (abs) (2B), num paths (1B),
    layer-0: prefix, varint(key.length), key, varint(hash(value).length), hash(value)
    path-n: [header: suffix start (rel) (1B), suffix end (rel) (1B)],  path[n].prefix, path[n].suffix
    */
    function verifyMembership(bytes32 root, bytes memory key, bytes32 value, bytes calldata proof)
        public
        pure
        virtual
    {
        uint256 path0start = uint256(uint8(proof[1]));
        bytes32 prehash = sha256(abi.encodePacked(proof[2:path0start], key, hex"20", sha256(abi.encodePacked(value))));
        uint256 offset = path0start;

        for (uint256 i = 0; i < uint256(uint8(proof[0])); ++i) {
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

    function parseEvent(bytes calldata rawEvent, uint8 numTopics)
        public
        pure
        virtual
        returns (address emittingContract, bytes memory topics, bytes memory unindexedData)
    {
        uint256 topicsEnd = 32 * numTopics + 20;
        return (address(bytes20(rawEvent[:20])), rawEvent[20:topicsEnd], rawEvent[topicsEnd:]);
    }
}
