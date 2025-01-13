// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../../contracts/core/SequencerSignatureVerifier.sol";
import {RLPWriter} from "optimism/libraries/rlp/RLPWriter.sol";
import {L1Header} from "../../contracts/interfaces/IAppStateVerifier.sol";

contract SigningBase is Test {
    using stdJson for string;

    string rootDir = vm.projectRoot();

    bytes32 l1BlockHash;
    bytes32 peptideAppHash;
    uint256 public peptideBlockNumber;
    uint256 public sequencerPkey;
    uint256 public notSequencerPkey;
    address public sequencer;
    bytes32 hashToSign;
    bytes32 domain; // Domain will be empty so we can leave it as initialized to default 0x 32 bytes

    SequencerSignatureVerifier public sigVerifier;
    bytes32 PEPTIDE_CHAIN_ID = bytes32(uint256(444));

    L1Header childl1Block; // Child block, represents the l1 origin of dest chain when peptide catches up to ancestor L1
        // block
    L1Header ancestorL1Block; // Ancestor block, represents the l1 origin of dest chain when peptide wants to submit a

    // client update but is behind, so we need to checkpoint this block

    constructor() {
        (sequencer, sequencerPkey) = makeAddrAndKey("alice");
        (, notSequencerPkey) = makeAddrAndKey(unicode"bob😈");
        sigVerifier = new SequencerSignatureVerifier(sequencer, PEPTIDE_CHAIN_ID);

        // generate the channel_proof.hex file with the following command:
        // cd test-data-generator && go run ./cmd/ --type l1 > ../test/payload/l1_block_0x4df537.hex
        // this is the "rlp" half-encoded header that would be sent by the relayer. this was produced
        // by the test-data-generator tool.
        L1Header memory l1header = abi.decode(
            vm.parseBytes(vm.readFile(string.concat(rootDir, "/test/payload/l1_block_0x4df537.hex"))), (L1Header)
        );

        l1BlockHash = keccak256(RLPWriter.writeList(l1header.header)); // Blockhash that will be signed by sequencer

        peptideBlockNumber = 101;

        // this happens to be the polymer height when the L2OO was updated with the output proposal
        // we are using in the test
        string memory l2BlockJson = vm.readFile(string.concat(rootDir, "/test/payload/l2_block_0x4b0.json"));
        peptideAppHash = abi.decode(l2BlockJson.parseRaw(".result.stateRoot"), (bytes32));

        bytes32 payloadHash = keccak256(abi.encodePacked(peptideBlockNumber, peptideAppHash, l1BlockHash));
        hashToSign = keccak256(bytes.concat(domain, PEPTIDE_CHAIN_ID, payloadHash));
    }

    // Read a specific bytes32 json property from a json file
    function readBytes32FromJson(string memory fileName, string memory property) public view returns (bytes32) {
        string memory blockJson = vm.readFile(string.concat(rootDir, fileName));
        return abi.decode(blockJson.parseRaw(property), (bytes32));
    }

    // Read a specific uint64 json property from a json file
    function readUint64FromJson(string memory fileName, string memory property) public view returns (uint64) {
        string memory blockJson = vm.readFile(string.concat(rootDir, fileName));
        return abi.decode(blockJson.parseRaw(property), (uint64));
    }
}
