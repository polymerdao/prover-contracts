// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import {Test} from "forge-std/Test.sol";
import {Base} from "../prove_api/utils/Base.t.sol";

import {OPStackBedrockProver} from "../../contracts/core/native_fallback/L2/OPStackBedrockProver.sol";
import {NativeProver} from "../../contracts/core/native_fallback/L2/NativeProver.sol";
import {UpdateL2ConfigArgs, L2Configuration} from "../../contracts/libs/RegistryTypes.sol";
import {INativeProver} from "../../contracts/interfaces/INativeProver.sol";
import {L2Configuration, Type, ProveScalarArgs} from "../../contracts/libs/RegistryTypes.sol";
import {RLPReader} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPReader.sol";
import {RLPWriter} from "@eth-optimism/contracts-bedrock/src/libraries/rlp/RLPWriter.sol";

// Test a unit test generated from forked op, base, and eth networks.
contract OPStackBedrockProverTest is Base {
    OPStackBedrockProver public prover;

    INativeProver public nativeProver = INativeProver(makeAddr("native-prover"));

    // Test a proof with op as the src chain, and base as the dest chain with proof generated from a
    function test_native_proof() public {
        bytes memory bytecode = load_str_from_hex("/test/native_fallback/payload/base-prover-bytecode.hex");


        vm.etch(address(nativeProver), bytecode);

        bytes memory proof = load_str_from_hex("/test/native_fallback/payload/native-proof.hex");

        (
            UpdateL2ConfigArgs memory _updateArgs,
            ProveScalarArgs memory _proveArgs,
            bytes memory _rlpEncodedL1Header,
            bytes memory _rlpEncodedL2Header,
            bytes memory _settledStateProof,
            bytes[] memory _l2StorageProof,
            bytes memory _rlpEncodedContractAccount,
            bytes[] memory _l2AccountProof
        ) = abi.decode(proof, (UpdateL2ConfigArgs, ProveScalarArgs, bytes, bytes, bytes, bytes[], bytes, bytes[]));

        // nativeProver.updateAndProve(
        //     _updateArgs,
        //     _proveArgs,
        //     _rlpEncodedL1Header,
        //     _rlpEncodedL2Header,
        //     _settledStateProof,
        //     _l2StorageProof,
        //     _rlpEncodedContractAccount,
        //     _l2AccountProof
        // );
    }
}
