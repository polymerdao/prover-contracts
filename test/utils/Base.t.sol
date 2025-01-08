// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import {stdStorage, StdStorage} from "forge-std/Test.sol";
import {Ics23Proof} from "../../contracts/libs/ReceiptParser.sol";

// Base contract for testing Dispatcher
contract Base is Test {
    using stdStorage for StdStorage;

    string rootDir = vm.projectRoot();

    bytes32 PEPTIDE_CHAIN_ID = bytes32(uint256(444));

    // Load a receipt proof specifically for testing receipt proofs for opstack and arbitrum receipts
    function load_receipt(string memory proofPath, string memory receiptPath)
        internal
        returns (bytes memory receiptIdx, bytes[] memory receiptProof, bytes32 receiptRoot, bytes memory receiptRLP)
    {
        (receiptRLP) = vm.parseBytes(vm.readFile(string.concat(rootDir, receiptPath)));

        (, receiptProof, receiptRoot,,, receiptIdx) = abi.decode(
            vm.parseBytes(vm.readFile(string.concat(rootDir, proofPath))),
            (Ics23Proof, bytes[], bytes32, uint256, string, bytes)
        );
    }

    // Load a full proof bytes proof from a file into calldata
    function load_proof(string memory filepath, address lightClient) internal returns (Ics23Proof memory) {
        (bytes32 apphash, Ics23Proof memory proof) =
            abi.decode(vm.parseBytes(vm.readFile(string.concat(rootDir, filepath))), (bytes32, Ics23Proof));

        // this loads the app hash we got from the testing data into the consensus state manager internals
        // at the height it's supposed to go. That is, a block less than where the proof was generated from.
        stdstore.target(lightClient).sig("consensusStates(uint256)").with_key(proof.height - 1).checked_write(apphash);
        // trick the fraud time window check
        vm.warp(block.timestamp + 1);

        return proof;
    }
}
