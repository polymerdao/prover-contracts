// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import {console2} from "forge-std/Test.sol";
import {CrossL2Prover} from "../contracts/core/CrossL2Prover.sol";
import {SequencerSignatureVerifier} from "../contracts/core/SequencerSignatureVerifier.sol";

// This suite can be used to reproduce any contract-related issues from our live devnet or testnet versions.
// This is useful for debugging because it enables things like being able to directly modify contracts easily and
// injecting things like foundry's console2.log and run the forge test commands with  the verose (-vvvv) option

// To Use:
// 1.) Find the calldata for the contract client update to debug and copy it into the clientUpdate var in the setup()
// 2.) Get a proof api response for the corresponding client update and copy it into the receiptproof var in setup() (
// note: you might need to convert from base 64 to hex)
// 3.) run the test depending on wether it's a devnet or testnet issue:
//      - for devnet:
// forge test --match-test test_devnet_repro -vvvv
//      - for testnet:
// forge test --match-test test_testnet_repro -vvvv

contract ContractDebugReproTest is Test {
    bytes clientUpdate; // the contract client update which will update the peptide hash that the verify receipt proof
        // can use
    bytes receiptProof; // Used to test any reciept proof. The hex encoded version of what is returned from a completed
        // proof api job.

    // Values currently deployed for devnet
    address devnetSigner = 0x11a72195e668328fEe607a309EfA2C42B2893E1d;
    bytes32 devnetPeptideClientId = bytes32(0x0000000000000000000000000000000000000000000000000000000000000385);

    // Values currently deployed for testnet
    address testnetSigner = 0xf029563823C84983AdC20fEA201889b566cAa005;
    bytes32 testnetPeptideClientId = bytes32(0x0000000000000000000000000000000000000000000000000000000001bc1bc1);

    function setUp() public {
        // FILL VALUES OUT HERE
        clientUpdate = hex"";
        receiptProof = hex"";
    }

    function test_devnet_repro() public {
        vm.skip(true); // Comment this out when running tests!
        SequencerSignatureVerifier sigVerifier = new SequencerSignatureVerifier(devnetSigner, devnetPeptideClientId);

        CrossL2Prover crossProver = new CrossL2Prover(sigVerifier, "proof_api");
        // Do contract client update
        (bool success, bytes memory returnData) = address(crossProver).call(clientUpdate);
        // Do receipt call using proof
        crossProver.validateReceipt(receiptProof);
    }

    function test_testnet_repro() public {
        vm.skip(true); // Comment this out when running tests!

        // Comment below out
        SequencerSignatureVerifier sigVerifier = new SequencerSignatureVerifier(testnetSigner, testnetPeptideClientId);

        CrossL2Prover crossProver = new CrossL2Prover(sigVerifier, "proof_api");
        // Do contract client update
        (bool success, bytes memory returnData) = address(crossProver).call(clientUpdate);
        // Do receipt call using proof
        crossProver.validateReceipt(receiptProof);
    }
}
