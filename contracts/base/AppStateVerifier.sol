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

import "forge-std/Test.sol";

pragma solidity 0.8.15;

import {RLPReader} from "optimism/libraries/rlp/RLPReader.sol";

import {IAppStateVerifier} from "../interfaces/IAppStateVerifier.sol";
import {Ics23Proof, OpIcs23Proof} from "../libs/ReceiptParser.sol";

/**
 * @title OptimisticProofVerifier
 * @notice Verifies proofs related to Optimistic Rollup state updates
 * @author Polymer Labs
 */
abstract contract AppStateVerifier is IAppStateVerifier {
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for bytes;

    /**
     * @dev Prove that a given state is not part of a proof
     * @dev this method is mainly used for packet timeouts, which is currently not implemented
     */
    function verifyNonMembership(bytes32, bytes calldata, Ics23Proof calldata) external pure {
        revert MethodNotImplemented();
    }

    /**
     * @inheritdoc IAppStateVerifier
     * @dev verifies a chain of ICS23 proofs
     * Each computed subroot starting from index 0 must match the value of the next proof (hence chained proofs).
     * The cosmos SDK and ics23 support chained proofs to switch between different proof specs.
     * Custom proof specs are not supported here. Only Iavl and Tendermint or similar proof specs are supported.
     */
    function verifyMembership(bytes32 appHash, bytes memory key, bytes32 value, Ics23Proof calldata proofs)
        public
        view
    {
        console2.log("  verifyMembershipABI 1", gasleft());
        //console2.log("key in verify membership: ", string(key), string(proofs.proof[0].key));
        // first check that the provided proof indeed proves the keys and values.
        if (keccak256(key) != keccak256(proofs.proof[0].key)) {
            revert InvalidProofKey();
        }
        console2.log("  verifyMembershipABI 2", gasleft());
        if (keccak256(abi.encodePacked(value)) != keccak256(proofs.proof[0].value)) revert InvalidProofValue();
        // proofs are chained backwards. First proof in the list (proof[0]) corresponds to the packet proof, meaning
        // that can be checked against the next subroot value (i.e. ibc root). Once the first proof is verified,
        // we can check the second that corresponds to the ibc proof, that is checked against the app hash (app root)
        if (bytes32(proofs.proof[1].value) != _verify(proofs.proof[0])) revert InvalidPacketProof();
        console2.log("  verifyMembershipABI 3", gasleft());
        if (appHash != _verify(proofs.proof[1])) revert InvalidIbcStateProof();
        console2.log("  verifyMembershipABI 4", gasleft());
    }

    function verifyMembershipRLP(bytes32 appHash, bytes memory key, bytes32 value, bytes memory proofs) public view {
        console2.log("  verifyMembershipRLP 1", gasleft());
        RLPReader.RLPItem[] memory opIcs23Proof = RLPReader.readList(proofs);

        console2.log("  verifyMembershipRLP 2", gasleft());
        RLPReader.RLPItem[] memory iavl = RLPReader.readList(opIcs23Proof[0]);
        console2.log("  verifyMembershipRLP 3", gasleft());
        RLPReader.RLPItem[] memory simple = RLPReader.readList(opIcs23Proof[1]);
        console2.log("  verifyMembershipRLP 4", gasleft());

        // first check that the provided proof indeed proves the keys and values.
        bytes memory keyRLP = RLPReader.readBytes(iavl[1]);
        if (keccak256(key) != keccak256(keyRLP)) {
            revert InvalidProofKey();
        }

        console2.log("  verifyMembershipRLP 5", gasleft());
        bytes memory valueRLP = RLPReader.readBytes(iavl[2]);
        if (keccak256(abi.encodePacked(value)) != keccak256(valueRLP)) {
            revert InvalidProofValue();
        }
        console2.log("  verifyMembershipRLP 6", gasleft());
        // proofs are chained backwards. First proof in the list (proof[0]) corresponds to the packet proof, meaning
        // that can be checked against the next subroot value (i.e. ibc root). Once the first proof is verified,
        // we can check the second that corresponds to the ibc proof, that is checked against the app hash (app root)
        if (
            bytes32(RLPReader.readBytes(simple[2]))
                != _verifyRLPOP(keyRLP, valueRLP, RLPReader.readBytes(iavl[3]), RLPReader.readList(iavl[0]))
        ) revert InvalidPacketProof();
        console2.log("  verifyMembershipRLP 7", gasleft());
        if (
            appHash
                != _verifyRLPOP(
                    RLPReader.readBytes(simple[1]),
                    RLPReader.readBytes(simple[2]),
                    RLPReader.readBytes(simple[3]),
                    RLPReader.readList(simple[0])
                )
        ) revert InvalidIbcStateProof();
        console2.log("  verifyMembershipRLP 8", gasleft());
    }

    /**
     * @dev Verifies an ICS23 proof through the root hash based on the provided proof.
     * @dev This code was adapted from the ICS23 membership verification found here:
     * https://github.com/cosmos/ics23/blob/go/v0.10.0/go/ics23.go#L36
     * @param proof The ICS23 proof to be verified.
     * @return computed The computed root hash.
     */
    function _verify(OpIcs23Proof calldata proof) internal view returns (bytes32 computed) {
        console2.log("    verify abi 1b", gasleft());
        bytes32 hashedData = sha256(proof.value);
        console2.log("    verify abi 2b", gasleft());
        computed = sha256(
            abi.encodePacked(
                proof.prefix, _encodeVarint(proof.key.length), proof.key, _encodeVarint(hashedData.length), hashedData
            )
        );
        console2.log("    verify abi 3b", gasleft());

        for (uint256 i = 0; i < proof.path.length; i++) {
            console2.log("    verify abi 4b", gasleft());
            computed = sha256(abi.encodePacked(proof.path[i].prefix, computed, proof.path[i].suffix));
        }
        console2.log("    verify abi 5b", gasleft());
    }

    function _verifyRLP(RLPReader.RLPItem[] memory proof) internal view returns (bytes32 computed) {
        console2.log("    verify RLP 1b", gasleft());
        bytes memory key = RLPReader.readBytes(proof[1]);
        console2.log("   verify RLP 2b", gasleft());
        bytes memory value = RLPReader.readBytes(proof[2]);
        console2.log("   verify RLP 3b", gasleft());
        bytes32 hashedData = sha256(value);
        console2.log("   verify RLP 4b", gasleft());
        computed = sha256(
            abi.encodePacked(
                RLPReader.readBytes(proof[3]),
                _encodeVarint(key.length),
                key,
                _encodeVarint(hashedData.length),
                hashedData
            )
        );
        console2.log("   verify RLP 5b", gasleft());

        RLPReader.RLPItem[] memory paths = RLPReader.readList(proof[0]);
        console2.log("    verify RLP 6b", gasleft());
        for (uint256 i = 0; i < paths.length; i++) {
            console2.log("    verify RLP 7b", gasleft());
            RLPReader.RLPItem[] memory path = RLPReader.readList(paths[i]);
            console2.log("    verify RLP 8b", gasleft());
            computed = sha256(abi.encodePacked(RLPReader.readBytes(path[0]), computed, RLPReader.readBytes(path[1])));
            console2.log("    verify RLP 9b", gasleft());
        }
        console2.log("    verify RLP 10b", gasleft());
    }

    function _verifyRLPOP(bytes memory key, bytes memory value, bytes memory prefix, RLPReader.RLPItem[] memory paths)
        internal
        pure
        returns (bytes32 computed)
    {
        bytes32 hashedData = sha256(value);
        computed = sha256(
            abi.encodePacked(
                prefix,
                _encodeVarint(key.length),
                key,
                _encodeVarint(hashedData.length),
                hashedData
            )
        );

        for (uint256 i = 0; i < paths.length; i++) {
            RLPReader.RLPItem[] memory path = RLPReader.readList(paths[i]);
            computed = sha256(abi.encodePacked(RLPReader.readBytes(path[0]), computed, RLPReader.readBytes(path[1])));
        }
    }

    /**
     * @dev Encodes an integer value into a variable-length integer format.
     * @param value The integer value to be encoded.
     * @return encoded The encoded bytes array.
     */
    function _encodeVarint(uint256 value) internal pure returns (bytes memory encoded) {
        bytes memory result;
        while (value >= 0x80) {
            bytes.concat(result, bytes1(uint8((value & 0x7F) | 0x80)));
            value >>= 7;
        }
        encoded = bytes.concat(result, bytes1(uint8(value)));
    }
}
