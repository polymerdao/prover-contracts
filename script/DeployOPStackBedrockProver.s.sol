// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/core/native_fallback/L2/OPStackBedrockProver.sol";

contract DeployOPStackBedrockProverScript is Script {
    function run() external returns (address opStackBedrockProver) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        vm.startBroadcast(deployerPrivateKey);
        
        // Deploy OPStackBedrockProver contract
        OPStackBedrockProver proverContract = new OPStackBedrockProver();
        
        vm.stopBroadcast();
        
        return address(proverContract);
    }
}