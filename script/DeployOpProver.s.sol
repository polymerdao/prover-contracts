// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/core/native_fallback/L2/OPStackCannonProver.sol";

contract DeployNativeProverScript is Script {
    function run() external returns (address nativeProver) {
        vm.startBroadcast();
        console2.log(
            "OpstackCannonProver: ", address(new OPStackCannonProver{salt: keccak256(abi.encode(uint256(123)))}())
        );
        vm.stopBroadcast();
    }
}
