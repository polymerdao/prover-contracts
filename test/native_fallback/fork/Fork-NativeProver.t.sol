// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import {Test} from "forge-std/Test.sol";
import "../../../contracts/core/native_fallback/L2/NativeProver.sol";
import "../../../contracts/libs/RegistryTypes.sol";
import "../../../script/DeployRegistry.s.sol";

contract DeployNativeProverScript is DeployRegistryScript, Test {
    NativeProver nativeProver;
    bool shouldRunTest;

    function setUp() public {
        shouldRunTest = vm.envOr("FORK_TEST", false);
        if (!shouldRunTest) {
            return;
        }
        vm.createSelectFork(vm.envString("BASE_RPC_URL"));
        uint256 chainId = vm.envUint("CHAIN_ID");
        address settlementRegistry = vm.envAddress("SETTLEMENT_REGISTRY");

        // Create an L1Configuration
        L1Configuration memory l1Config = L1Configuration({
            blockHashOracle: blockHashOracle,
            settlementRegistry: settlementRegistry,
            settlementRegistryL2ConfigMappingSlot: uint256(2),
            settlementRegistryL1ConfigMappingSlot: _STARTING_L1_MAPPING_SLOT
        });

        L2Configuration memory baseL2Config = L2Configuration({
            prover: baseProver,
            addresses: addresses(0xd6E6dBf4F7EA0ac412fD8b65ED297e64BB7a06E1),
            storageSlots: storageSlots(),
            versionNumber: 0,
            finalityDelaySeconds: 0,
            l2Type: Type.OPStackCannon
        });

        L2Configuration memory opL2Config = L2Configuration({
            prover: opProver,
            addresses: addresses(0x05F9613aDB30026FFd634f38e5C4dFd30a197Fa1),
            storageSlots: storageSlots(),
            versionNumber: 0,
            finalityDelaySeconds: 0,
            l2Type: Type.OPStackCannon
        });

        vm.startBroadcast(deployerPrivateKey);

        // // Deploy NativeProver contract
        nativeProver = new NativeProver(0x81b5c1c0343ff0087C04F543D0f36dC4745b999F, ethChainId);
        nativeProver.setInitialL1Config(l1Config);

        vm.stopBroadcast();
        console2.log("nativeProver: ", address(nativeProver));
    }

    function test_integration_proof() public {
        vm.skip(!shouldRunTest);
        bytes memory proof = vm.envBytes("PROOF");

        (bool success, bytes memory returnData) = address(nativeProver).call(proof);
        require(success, "Proof execution failed");

        console2.log("returned data");
        console2.logBytes(returnData);
        // Log the returned data
        if (returnData.length > 0) {
            console2.log("Returned data length:", returnData.length);
        }
    }

    function mock_l1_hash_oracle() internal {
        bytes32 l1BlockHash = vm.envBytes32("BlockHash");
    }
}
