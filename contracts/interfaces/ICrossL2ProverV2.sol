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

/**
 * @title ICrossL2Prover
 * @author Polymer Labs
 * @notice A contract that can prove polymers state. Since polymer is an aggregator of many chains' states, this
 * contract can in turn be used to prove any arbitrary events and/or storage on counterparty chains.
 */
interface ICrossL2ProverV2 {
    /**
     * @notice Validate a log on a counterparty chain has been written to polymer state.
     * @param proof: The proof that the hash of a given log has been written to polymer state, returned from the polymer
     * proof api.
     *   // Proof calldata is encoded in the following format:
     *   //  +--------------------------------------------------+
     *   //  |  state root (32 bytes)                           | index: 0:32
     *   //  +--------------------------------------------------+
     *   //  |  signature (65 bytes)                            | index: 32:97
     *   //  +--------------------------------------------------+
     *   //  |  source chain ID (big endian, 4 bytes)           | index: 97:101
     *   //  +--------------------------------------------------+
     *   //  |  polymer height (big endian, 8 bytes)            | index: 101:109
     *   //  +--------------------------------------------------+
     *   //  |  source chain block height (big endian, 8 bytes) | index: 109:117
     *   //  +--------------------------------------------------+
     *  //  |  receipt index (big endian, 2 bytes)             | index: 117:119
     *   //  +--------------------------------------------------+
     *   //  |  event index (1 byte)                            | index: 119
     *   //  +--------------------------------------------------+
     *   //  |  number of topics (1 byte)                       | index: 120
     *   //  +--------------------------------------------------+
     *   //  |  event data end (big endian, 2 bytes)            | index: 121:123
     *   //  +--------------------------------------------------+
     *   //  |  event emitter (contract address) (20 bytes)     | index: 123:143
     *   //  +--------------------------------------------------+
     *   //  |  topics (32 bytes * number of topics)            | index: 143 + 32 * number of topics: eventDatEnd
     *   //  +--------------------------------------------------+
     *   //  |  event data (x bytes)                            | index: 123:eventDataEnd
     *   //  +--------------------------------------------------+
     *   //  |  iavl proof (x bytes)                            | index: eventDataEnd:
     *   //  +--------------------------------------------------+
     * @return chainId The source chainID the event was emitted on. This is not part of the pre-image of the log but is
     * part of the polymer key (thus preventing spoofing attacks)
     * @return emittingContract The address of the contract that emitted the log on the source chain, is part of the
     * pre-image of the log hash committed to polymer.
     * @return topics The topics of the event. First topic is the event signature that can be calculated by
     * Event.selector. The remaining elements in this array are the indexed parameters of the event. This is part of the
     * pre-image of the log hash committed to polymer.
     * @return unindexedData The abi encoded non-indexed parameters of the event, part of the pre-image of the log
     * hash committed to polymer.
     */
    function validateEvent(bytes calldata proof)
        external
        view
        returns (uint32 chainId, address emittingContract, bytes calldata topics, bytes calldata unindexedData);

    /**
     * @notice Parse srcChain, blockNumber, receiptIndex, and eventIndex for a requested proof.
     * @param proof: The proof that the hash of a given log has been written to polymer state, returned from the polymer
     * proof api, same proof as in validateEvent and inspectPolymerState.
     * @return srcChain The chainID of the chain that the receipt was emitted on.
     * @return blockNumber The block number of the block that the receipt was emitted in.
     * @return receiptIndex The index of the receipt within the block.
     * @return logIndex The index of the log within the receipt. Note: this is not the index of the log within the
     * block.
     */
    function inspectLogIdentifier(bytes calldata proof)
        external
        pure
        returns (uint32 srcChain, uint64 blockNumber, uint16 receiptIndex, uint8 logIndex);

    /**
     * Parse polymer state root, height , and signature over height and root pair which can be verified by
     * crypto.pubkey(keccak(polymerStateRoot, polymerHeight))
     */

    /**
     * @notice Parse polymer state root, height, and sequencer signature over height and root pair for visibility &
     * transparency.
     * @param proof: The proof that the hash of a given log has been written to polymer state, returned from the polymer
     * proof api, same proof as in validateEvent and inspectPolymerState.
     *  @return stateRoot The root of the polymer state trie that has been signed by the polymer sequencer key.
     *  @return height The height of the polymer state trie that has been signed by the polymer sequencer key.
     *  @return signature The signature over the keccak(root + height) that can be verified.
     */
    function inspectPolymerState(bytes calldata proof)
        external
        pure
        returns (bytes32 stateRoot, uint64 height, bytes memory signature);

    /**
     * @notice Verify a membership proof that a key and value exists at a given polymer state root
     * @param root The root of the polymer state trie
     * @param key The key to verify
     * @param value The value to verify
     * @param proof The proof of the key and value, encoded as follows:
     *  +------------------------------------------------------------------------------------+
     *  |                                    Header (5 bytes):                                |
     *  +------------------------------------------------------------------------------------+
     *  | Key start (2B) | Key end (2B) | Value start (2B) | Value end (2B) | Num paths (1B)  |
     *  +------------------------------------------------------------------------------------+
     *  |                                Layer 0 prefix (x bytes):                            |
     *  +------------------------------------------------------------------------------------+
     *  | Prefix (var len) | Key length (var len) | Key (var len) |  0x20 , hash(value) (32B) |
     *  +------------------------------------------------------------------------------------+
     *  |                                      Path-n:                                        |
     *  +------------------------------------------------------------------------------------+
     *  | [header: Suffix start (1B), Suffix end (rel) (1B)],  path[n].prefix, path[n].suffix |
     *  +------------------------------------------------------------------------------------+
     */
    function verifyMembership(bytes32 root, bytes memory key, bytes32 value, bytes calldata proof) external pure;

    /**
     * @notice Parse an event from a binary encoded event
     * @param rawEvent The raw event data, the hash of which is committed into polymer if a valid proof exists
     * @param numTopics The number of topics in the event, needed to parse the event correctly
     * @return emittingContract The address of the contract that emitted the event
     * @return topics The topics of the event. First topic is the event signature that can be calculated by
     * Event.selector. The remaining elements in this array are the indexed parameters of the event.
     * @return unindexedData The abi encoded non-indexed parameters of the event.
     */
    function parseEvent(bytes calldata rawEvent, uint8 numTopics)
        external
        pure
        returns (address emittingContract, bytes memory topics, bytes memory unindexedData);
}
