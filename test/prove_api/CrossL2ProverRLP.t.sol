// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import {Base} from "./utils/Base.t.sol";
import {CrossL2Prover} from "../../contracts/core/prove_api/CrossL2Prover.sol";
import {SequencerSignatureVerifier} from "../../contracts/core/prove_api/SequencerSignatureVerifier.sol";

import {ISignatureVerifier} from "../../contracts/interfaces/ISignatureVerifier.sol";
import {ICrossL2Prover} from "../../contracts/interfaces/ICrossL2Prover.sol";

// These were two tests which test for some edge cases with RLP encoding which we can keep as regression testing in case
// we change any low-level RLP encoding logic.
contract CrossL2ProverDevnet is Base {
    ICrossL2Prover prover;

    function setUp() public {
        SequencerSignatureVerifier sigVerifier = new SequencerSignatureVerifier(
            address(0x11a72195e668328fEe607a309EfA2C42B2893E1d),
            bytes32(0x0000000000000000000000000000000000000000000000000000000000000385)
        );

        prover = new CrossL2Prover(sigVerifier, "proof_api", 100);
    }

    function test_RLP_1() public {
        bytes memory clientUpdate = load_bytes_from_hex("/test/payload/RLP/clientUpdate1.hex");

        bytes memory proof = load_bytes_from_hex("/test/payload/RLP/proof1.hex");

        (bool success, bytes memory returnData) = address(prover).call(clientUpdate);

        require(success, "Call to client update failed");
        prover.validateReceipt(proof);
        prover.validateEvent(0, proof);
    }

    function test_RLP_2() public {
        bytes memory clientUpdate = load_bytes_from_hex("/test/payload/RLP/clientUpdate2.hex");

        bytes memory proof = load_bytes_from_hex("/test/payload/RLP/proof2.hex");

        (bool success, bytes memory returnData) = address(prover).call(clientUpdate);

        require(success, "Call to client update failed");
        prover.validateReceipt(proof);
        prover.validateEvent(0, proof);
    }
}
