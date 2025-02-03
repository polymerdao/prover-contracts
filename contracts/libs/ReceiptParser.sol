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

pragma solidity ^0.8.0;

import {RLPReader} from "optimism/libraries/rlp/RLPReader.sol";
import {Bytes} from "optimism/libraries/Bytes.sol";

// OpIcs23ProofPath represents a commitment path in an ICS23 proof, which consists of a commitment prefix and a suffix.
struct OpIcs23ProofPath {
    bytes prefix;
    bytes suffix;
}

// OpIcs23Proof represents an ICS23 proof
struct OpIcs23Proof {
    OpIcs23ProofPath[] path;
    bytes key;
    bytes value;
    bytes prefix;
}

// the Ics23 proof related structs are used to do membership verification. These are not the actual Ics23
// format but a "solidity friendly" version of it - data is the same just packaged differently
struct Ics23Proof {
    OpIcs23Proof[] proof;
    uint256 height;
}

// This is the proof we use to verify the apphash (state) updates.
struct OpL2StateProof {
    bytes[] accountProof;
    bytes[] outputRootProof;
    bytes32 l2OutputProposalKey;
    bytes32 l2BlockHash;
}

/**
 * A library for helpers for proving peptide state
 */
library ReceiptParser {
    error invalidAddressBytes();

    function toStr(uint256 _number) public pure returns (string memory outStr) {
        if (_number == 0) {
            return "0";
        }

        uint256 length;
        uint256 number = _number;

        // Determine the length of the string
        while (number != 0) {
            length++;
            number /= 10;
        }

        bytes memory buffer = new bytes(length);

        // Convert each digit to its ASCII representation
        for (uint256 i = length; i > 0; i--) {
            buffer[i - 1] = bytes1(uint8(48 + (_number % 10)));
            _number /= 10;
        }

        outStr = string(buffer);
    }

    function bytesToAddr(bytes memory a) public pure returns (address addr) {
        if (a.length != 20) {
            revert invalidAddressBytes();
        }
        assembly {
            addr := mload(add(a, 20))
        }
    }

    function parseLog(uint256 logIndex, bytes memory receiptRLP)
        internal
        pure
        returns (address emittingContract, bytes[] memory topics, bytes memory unindexedData)
    {
        // The first byte is a RLP encoded receipt type so slice it off.
        uint8 typeByte;
        assembly {
            typeByte := byte(0, mload(add(receiptRLP, 32)))
        }
        if (typeByte < 0x80) {
            // Typed receipt: strip the type byte
            receiptRLP = Bytes.slice(receiptRLP, 1, receiptRLP.length - 1);
        }

        RLPReader.RLPItem[] memory receipt = RLPReader.readList(receiptRLP);
        /*
            // RLP encoded receipt has the following structure. Logs are the 4th RLP list item.
            type ReceiptRLP struct {
                    PostStateOrStatus []byte
                   CumulativeGasUsed uint64
                    Bloom             Bloom
                    Logs              []*Log
            }
        */

        // Each log itself is an rlp encoded datatype of 3 properties:
        // type Log struct {
        //         senderAddress bytes // contract address where this log was emitted from
        //         topics bytes        // Array of indexed topics. The first element is the 32-byte selector of the
        // event (can use TransmitToHouston.selector), and the following  elements in this array are the abi encoded
        // arguments individually
        //         topics data         // abi encoded raw bytes of unindexed data
        // }
        RLPReader.RLPItem[] memory log = RLPReader.readList(RLPReader.readList(receipt[3])[logIndex]);

        emittingContract = bytesToAddr(RLPReader.readBytes(log[0]));

        RLPReader.RLPItem[] memory encodedTopics = RLPReader.readList(log[1]);
        unindexedData = (RLPReader.readBytes(log[2])); // This is the raw unindexed data. in this case it's
        // just an abi encoded uint64

        topics = new bytes[](encodedTopics.length);
        for (uint256 i = 0; i < encodedTopics.length; i++) {
            topics[i] = RLPReader.readBytes(encodedTopics[i]);
        }
    }

    function receiptRootKey(string memory chainId, string memory clientType, uint256 height)
        internal
        pure
        returns (bytes memory proofKey)
    {
        proofKey = abi.encodePacked("chain/", chainId, "/storedReceipts/", clientType, "/receiptRoot/", toStr(height));
    }

    function eventRootKey(uint32 chainId, string memory clientType, uint256 height, uint16 receiptIndex, uint8 logIndex)
        internal
        pure
        returns (bytes memory proofKey)
    {
        // TODO actually change this to the decided structure
        return abi.encodePacked(
            "chain/",
            toStr(uint256(chainId)),
            "/storedLogs/",
            clientType,
            "/",
            toStr(height),
            "/",
            toStr(receiptIndex),
            "/",
            toStr(logIndex)
        );
    }
}
