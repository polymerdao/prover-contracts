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

import {MerkleTrie} from "optimism/libraries/trie/MerkleTrie.sol";
import {ReceiptParser} from "../../libs/ReceiptParser.sol";
import {AppStateVerifier} from "../../base/AppStateVerifier.sol";
import {ICrossL2Prover} from "../../interfaces/ICrossL2Prover.sol";
import {Ics23Proof} from "../../libs/ReceiptParser.sol";
import {ISignatureVerifier} from "../../interfaces/ISignatureVerifier.sol";
import {LightClientType} from "../../interfaces/IClientUpdates.sol";

contract CrossL2Prover is AppStateVerifier, ICrossL2Prover {
    LightClientType public constant LIGHT_CLIENT_TYPE = LightClientType.SequencerLightClient; // Stored as a constant
        // for cheap on-chain use

    string public clientType;
    uint256 public immutable RING_BUFFER_LENGTH; // Length of which we mod heights in storage slots

    uint256 public latestHeight; // Used to prevent from griefing attacks that update to old heights

    ISignatureVerifier public immutable verifier;

    // Store peptide apphashes for a given height
    mapping(uint256 => uint256) public peptideAppHashes;

    error CannotUpdateToOlderHeight();

    constructor(ISignatureVerifier verifier_, string memory clientType_, uint256 ringBufferLength_) {
        verifier = verifier_;
        clientType = clientType_;
        RING_BUFFER_LENGTH = ringBufferLength_;
    }

    /**
     * @dev Adds an appHash to the internal store, after verifying the client update proof associated with the light
     * client implementation.
     * @param peptideAppHash App hash (state root) to be verified
     * @param proof An array of bytes that contain the l1blockhash and the sequencer's signature. The first 32 bytes of
     * this argument should be the l1BlockHash, and the remaining bytes should be the sequencer signature which attests
     * to the peptide AppHash
     * for that l1BlockHash
     */
    function updateClient(bytes calldata proof, uint256 peptideHeight, uint256 peptideAppHash) external {
        _updateClient(proof, peptideHeight, peptideAppHash);
    }

    function validateEvent(uint256 logIndex, bytes calldata proof)
        external
        view
        virtual
        returns (string memory chainId, address emittingContract, bytes[] memory topics, bytes memory unindexedData)
    {
        bytes memory receiptRLP;
        (chainId, receiptRLP) = validateReceipt(proof);
        (emittingContract, topics, unindexedData) = ReceiptParser.parseLog(logIndex, receiptRLP);
    }

    function getState(uint256 height) external view virtual returns (uint256) {
        return _getPeptideAppHash(height);
    }

    /**
     * @inheritdoc ICrossL2Prover
     */
    function validateReceipt(bytes calldata proof) public view virtual returns (string memory, bytes memory) {
        (
            Ics23Proof memory peptideAppProof,
            bytes[] memory receiptMMPTProof,
            bytes32 receiptRoot,
            uint64 eventHeight,
            string memory srcChainId,
            bytes memory receiptIndex
        ) = abi.decode(proof, (Ics23Proof, bytes[], bytes32, uint64, string, bytes));
        // Before we can trust the receipt root, we first need to verify that the receipt root is indeed stored
        // on peptide at the given clientID and height.

        // VerifyMembership verifies the receipt root  through an ics23 proof of peptide state that attests that the
        // given eventHeight has the receipt root at the peptide height
        this.verifyMembership(
            bytes32(_getPeptideAppHash(peptideAppProof.height - 1)), // a proof generated at height H can only be
                // verified against state root (app hash) from block H - 1. this means the relayer must have updated the
                // contract with the app hash from the previous block and that is why we use proof.height - 1 here.
            ReceiptParser.receiptRootKey(srcChainId, clientType, eventHeight),
            receiptRoot,
            peptideAppProof
        );

        // Now that verifyMembership passed, we can now trust the receiptRoot.
        // Now we just simply have to prove that raw receipt is indeed part of the receipt root at the given receipt
        // index.
        // This is done through a Merkle proof.

        return (srcChainId, MerkleTrie.get(receiptIndex, receiptMMPTProof, receiptRoot));
    }

    function _updateClient(bytes calldata proof, uint256 peptideHeight, uint256 peptideAppHash) internal {
        if (peptideHeight <= latestHeight) {
            revert CannotUpdateToOlderHeight();
        }

        verifier.verifyStateUpdate(peptideHeight, bytes32(peptideAppHash), bytes32(proof[:32]), proof[32:]);
        peptideAppHashes[peptideHeight % RING_BUFFER_LENGTH] = peptideAppHash;
        latestHeight = peptideHeight;
    }

    function _getPeptideAppHash(uint256 _height) internal view returns (uint256) {
        return peptideAppHashes[_height % RING_BUFFER_LENGTH];
    }
}
