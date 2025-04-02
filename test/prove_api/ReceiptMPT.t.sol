// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {console2} from "forge-std/Test.sol";
import {Base} from "./utils/Base.t.sol";
import {MerkleTrie} from "optimism/libraries/trie/MerkleTrie.sol";

// Testing receipt proofs are as the relayer encodes them.
contract ReceiptMPT is Base {
    function setUp() public {}

    function test_optimism_receipt() public {
        (bytes memory receiptIdx, bytes[] memory receiptProof, bytes32 receiptRoot, bytes memory receiptRLP) =
            load_receipt("/test/prove_api/payload/op-proof.hex", "/test/prove_api/payload/op-receipt.hex");
        assertEq(receiptRLP, MerkleTrie.get(receiptIdx, receiptProof, receiptRoot));
    }

    function test_arbitrum_receipt() public {
        (bytes memory receiptIdx, bytes[] memory receiptProof, bytes32 receiptRoot, bytes memory receiptRLP) =
            load_receipt("/test/prove_api/payload/arb-proof.hex", "/test/prove_api/payload/arb-receipt.hex");
        assertEq(receiptRLP, MerkleTrie.get(receiptIdx, receiptProof, receiptRoot));
    }
}
