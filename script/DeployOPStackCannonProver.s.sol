// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/core/native_fallback/L2/OPStackCannonProver.sol";

contract DeployOPStackCannonProverScript is Script {
    function run() external returns (address opStackCannonProver) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        vm.startBroadcast(deployerPrivateKey);
        
        // Deploy OPStackCannonProver contract
        OPStackCannonProver proverContract = new OPStackCannonProver();
        
        vm.stopBroadcast();
        
        return address(proverContract);
    }
}