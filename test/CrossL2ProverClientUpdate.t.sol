// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import {SigningBase} from "test/utils/Signing.base.t.sol";
import {CrossL2Prover} from "contracts/core/CrossL2Prover.sol";
import {SequencerSignatureVerifier} from "contracts/core/SequencerSignatureVerifier.sol";
import {Ics23Proof, OpIcs23Proof} from "contracts/libs/ReceiptParser.sol";
import {ISignatureVerifier} from "contracts/interfaces/ISignatureVerifier.sol";
import {ICrossL2Prover} from "contracts/interfaces/ICrossL2Prover.sol";
import {IAppStateVerifier} from "contracts/interfaces/IAppStateVerifier.sol";

contract CrossL2ProverTest is SigningBase {
    using stdJson for string;

    CrossL2Prover crossProver;

    function setUp() public {
        crossProver = new CrossL2Prover(sigVerifier, "proof_api", 100);
    }

    // Happy path for CrossEventProver.updateClient()
    function test_clientUpdate_success() public {
        // Ics23Proof memory iavlProof = abi.decode(peptideAppProofBytes, (Ics23Proof));

        uint256 height = 1;
        bytes32 peptideAppHash = keccak256(bytes.concat("peptide", abi.encodePacked(height)));
        bytes32 l1BlockHash = keccak256(bytes.concat("l1", abi.encodePacked(height)));
        bytes32 hashToSign = keccak256(
            bytes.concat(
                bytes32(0), PEPTIDE_CHAIN_ID, keccak256(bytes.concat(bytes32(height), peptideAppHash, l1BlockHash))
            )
        );

        sign_and_update(l1BlockHash, peptideAppHash, height, hashToSign);
        assertEq(crossProver.getState(height), uint256(peptideAppHash));
    }

    // Test ring buffer specific behavior

    // Happy path for CrossEventProver.updateClient()
    function test_clientUpdate_success_ring_buffer() public {
        // Ics23Proof memory iavlProof = abi.decode(peptideAppProofBytes, (Ics23Proof));

        // Fill ring buffer to test initialization
        for (uint256 height; height < crossProver.RING_BUFFER_LENGTH(); height++) {
            bytes32 peptideAppHash = keccak256(bytes.concat("peptide", abi.encodePacked(height)));
            bytes32 l1BlockHash = keccak256(bytes.concat("l1", abi.encodePacked(height)));
            bytes32 hashToSign = keccak256(
                bytes.concat(
                    bytes32(0), PEPTIDE_CHAIN_ID, keccak256(bytes.concat(bytes32(height), peptideAppHash, l1BlockHash))
                )
            );
            sign_and_update(l1BlockHash, peptideAppHash, height, hashToSign);
            assertEq(crossProver.getState(height), uint256(peptideAppHash));
        }

        // Now we overwrite ring buffer 10x
        for (
            uint256 height = crossProver.RING_BUFFER_LENGTH(); height < crossProver.RING_BUFFER_LENGTH() * 10; height++
        ) {
            bytes32 lastPeptideHash =
                keccak256(bytes.concat("peptide", abi.encodePacked(height - crossProver.RING_BUFFER_LENGTH())));
            assertEq(crossProver.getState(height), uint256(lastPeptideHash));

            bytes32 peptideAppHash = keccak256(bytes.concat("peptide", abi.encodePacked(height)));
            bytes32 l1BlockHash = keccak256(bytes.concat("l1", abi.encodePacked(height)));
            bytes32 hashToSign = keccak256(
                bytes.concat(
                    bytes32(0), PEPTIDE_CHAIN_ID, keccak256(bytes.concat(bytes32(height), peptideAppHash, l1BlockHash))
                )
            );
            sign_and_update(l1BlockHash, peptideAppHash, height, hashToSign);
            assertEq(crossProver.getState(height), uint256(peptideAppHash));
        }
    }

    function sign_and_update(bytes32 l1BlockHash, bytes32 peptideAppHash, uint256 height, bytes32 hashToSign)
        internal
    {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(sequencerPkey, hashToSign);
        bytes memory signature = abi.encodePacked(r, s, v); // Annoyingly, v, r, s are in a different order than those
        crossProver.updateClient(bytes.concat(l1BlockHash, signature), height, uint256(peptideAppHash));
    }
}
