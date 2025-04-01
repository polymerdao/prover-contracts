// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import {SigningBase} from "./utils/Signing.base.t.sol";
import {CrossL2Prover} from "../../contracts/core/prove_api/CrossL2Prover.sol";
import {SequencerSignatureVerifier} from "../../contracts/core/prove_api/SequencerSignatureVerifier.sol";
import {Ics23Proof, OpIcs23Proof} from "../../contracts/libs/ReceiptParser.sol";
import {ISignatureVerifier} from "../../contracts/interfaces/ISignatureVerifier.sol";
import {ICrossL2Prover} from "../../contracts/interfaces/ICrossL2Prover.sol";
import {IAppStateVerifier} from "../../contracts/interfaces/IAppStateVerifier.sol";

contract CrossL2ProverClientUpdateTest is SigningBase {
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

        // Fill ring buffer to test initialization. Don't start at 0 since client update for height 0 would revert.
        for (uint256 height = crossProver.RING_BUFFER_LENGTH(); height < crossProver.RING_BUFFER_LENGTH() * 2; height++)
        {
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
            uint256 height = crossProver.RING_BUFFER_LENGTH() * 2;
            height < crossProver.RING_BUFFER_LENGTH() * 11;
            height++
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

    // Test that updating a client with an old
    function test_Revert_clientUpdate_oldState() public {
        uint256 recentHeight = 200;
        bytes32 recentPeptideApphash = keccak256(bytes.concat("peptide", abi.encodePacked(recentHeight)));
        bytes32 recentL1BlockHash = keccak256(bytes.concat("l1", abi.encodePacked(recentHeight)));
        bytes32 recentHashToSign = keccak256(
            bytes.concat(
                bytes32(0),
                PEPTIDE_CHAIN_ID,
                keccak256(bytes.concat(bytes32(recentHeight), recentPeptideApphash, recentL1BlockHash))
            )
        );

        // Normal flow
        sign_and_update(recentL1BlockHash, recentPeptideApphash, recentHeight, recentHashToSign);

        uint256 oldHeight = 100;
        bytes32 oldPeptideApphash = keccak256(bytes.concat("peptide", abi.encodePacked(oldHeight)));
        bytes32 oldL1BlockHash = keccak256(bytes.concat("l1", abi.encodePacked(oldHeight)));
        bytes32 oldHashToSign = keccak256(
            bytes.concat(
                bytes32(0),
                PEPTIDE_CHAIN_ID,
                keccak256(bytes.concat(bytes32(oldHeight), oldPeptideApphash, oldL1BlockHash))
            )
        );

        // Expect revert for updating to an old height
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(sequencerPkey, oldHashToSign);
        bytes memory signature = abi.encodePacked(r, s, v); // Annoyingly, v, r, s are in a different order than those

        vm.expectRevert(CrossL2Prover.CannotUpdateToOlderHeight.selector);
        crossProver.updateClient(bytes.concat(oldL1BlockHash, signature), oldHeight, uint256(oldPeptideApphash));

        // Both heights should give the same state
        assertEq(crossProver.getState(recentHeight), uint256(recentPeptideApphash));
        assertEq(crossProver.getState(oldHeight), uint256(recentPeptideApphash));
    }

    function sign_and_update(bytes32 l1BlockHash, bytes32 peptideAppHash, uint256 height, bytes32 hashToSign)
        internal
    {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(sequencerPkey, hashToSign);
        bytes memory signature = abi.encodePacked(r, s, v); // Annoyingly, v, r, s are in a different order than those
        crossProver.updateClient(bytes.concat(l1BlockHash, signature), height, uint256(peptideAppHash));
    }
}
