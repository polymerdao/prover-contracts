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

import {IAppStateVerifier} from "../interfaces/IAppStateVerifier.sol";
import {Ics23Proof, OpIcs23Proof} from "../libs/ReceiptParser.sol";

/**
 * @title OptimisticProofVerifier
 * @notice Verifies proofs related to Optimistic Rollup state updates
 * @author Polymer Labs
 */
abstract contract AppStateVerifier is IAppStateVerifier {
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
        pure
    {
        // first check that the provided proof indeed proves the keys and values.
        if (keccak256(key) != keccak256(proofs.proof[0].key)) {
            revert InvalidProofKey();
        }
        if (keccak256(abi.encodePacked(value)) != keccak256(proofs.proof[0].value)) revert InvalidProofValue();
        // proofs are chained backwards. First proof in the list (proof[0]) corresponds to the packet proof, meaning
        // that can be checked against the next subroot value (i.e. ibc root). Once the first proof is verified,
        // we can check the second that corresponds to the ibc proof, that is checked against the app hash (app root)
        if (bytes32(proofs.proof[1].value) != _verify(proofs.proof[0])) revert InvalidPacketProof();
        if (appHash != _verify(proofs.proof[1])) revert InvalidIbcStateProof();
    }

    /*
    header: key start (abs) (2B), key end (abs) (2B), value start (abs) (2B), value end (abs) (2B), num paths (1B),
    layer-0: prefix, varint(key.length), key, varint(hash(value).length), hash(value)
    path-n: [header: suffix start (rel) (1B), suffix end (rel) (1B)],  path[n].prefix, path[n].suffix
    */
    function verifyMembershipNew(bytes32 root, bytes memory key, bytes32 value, bytes calldata proof) public pure {
        uint16 keystart = uint16(uint8(proof[0])) << 8 | uint8(proof[1]);
        uint16 keyend = uint16(uint8(proof[2])) << 8 | uint8(proof[3]);
        uint8 numpaths = uint8(proof[4]);
        uint32 offset = keyend + 33;

        require(keystart < keyend);
        require(numpaths > 0);
        require(proof[keyend] == 0x20);

        if (keccak256(key) != keccak256(proof[keystart:keyend])) revert InvalidProofKey();

        if (keccak256(abi.encodePacked(sha256(abi.encodePacked(value)))) != keccak256(proof[keyend + 1:offset])) {
            revert InvalidProofValue();
        }

        bytes32 prehash = sha256(proof[5:offset]);

        for (uint16 i = 0; i < numpaths; i++) {
            uint8 suffixstart = uint8(proof[offset]);
            uint8 suffixend = uint8(proof[offset + 1]);

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
     * @dev Verifies an ICS23 proof through the root hash based on the provided proof.
     * @dev This code was adapted from the ICS23 membership verification found here:
     * https://github.com/cosmos/ics23/blob/go/v0.10.0/go/ics23.go#L36
     * @param proof The ICS23 proof to be verified.
     * @return computed The computed root hash.
     */
    function _verify(OpIcs23Proof calldata proof) internal pure returns (bytes32 computed) {
        bytes32 hashedData = sha256(proof.value);
        computed = sha256(
            abi.encodePacked(
                proof.prefix, _encodeVarint(proof.key.length), proof.key, _encodeVarint(hashedData.length), hashedData
            )
        );

        for (uint256 i = 0; i < proof.path.length; i++) {
            computed = sha256(abi.encodePacked(proof.path[i].prefix, computed, proof.path[i].suffix));
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
