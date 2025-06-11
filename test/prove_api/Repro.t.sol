// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import {console2} from "forge-std/Test.sol";
import {CrossL2ProverV2} from "../../contracts/core/prove_api/CrossL2ProverV2.sol";
import {SequencerSignatureVerifier} from "../../contracts/core/prove_api/SequencerSignatureVerifier.sol";

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
    bytes proof; // Used to test any event proof.

    // Values currently deployed for devnet
    address devnetSigner = 0x11a72195e668328fEe607a309EfA2C42B2893E1d;
    bytes32 devnetPeptideClientId = bytes32(0x0000000000000000000000000000000000000000000000000000000000000385);

    // Values currently deployed for testnet
    address testnetSigner = 0xf029563823C84983AdC20fEA201889b566cAa005;
    bytes32 testnetPeptideClientId = bytes32(0x0000000000000000000000000000000000000000000000000000000001bc1bc1);

    function setUp() public {
        // FILL VALUES OUT HERE
        proof = hex"";
    }

    function test_devnet_repro() public {
        vm.skip(true); // Comment this out when running tests!
        SequencerSignatureVerifier sigVerifier = new SequencerSignatureVerifier(devnetSigner, devnetPeptideClientId);

        CrossL2ProverV2 crossProver = new CrossL2ProverV2("proof_api", devnetSigner, devnetPeptideClientId);
        // Do event call using proof
        crossProver.validateEvent(proof);
    }

    function test_testnet_repro() public {
        vm.skip(true); // Comment this out when running tests!

        // Comment below out
        SequencerSignatureVerifier sigVerifier = new SequencerSignatureVerifier(testnetSigner, testnetPeptideClientId);

        CrossL2ProverV2 crossProver = new CrossL2ProverV2("proof_api", testnetSigner, testnetPeptideClientId);
        // Do event call using proof
        crossProver.validateEvent(proof);
    }
}
