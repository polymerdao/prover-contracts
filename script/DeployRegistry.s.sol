// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/core/native_fallback/L1/Registry.sol";
import "../contracts/libs/RegistryTypes.sol";

contract DeployRegistryScript is Script {
    function run() external returns (address registry) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address initialOwner = vm.addr(deployerPrivateKey);
        
        vm.startBroadcast(deployerPrivateKey);

        // Create empty arrays for initial configurations
        Registry.InitialL2Configuration[] memory initialL2Configs = new Registry.InitialL2Configuration[](0);
        Registry.InitialL1Configuration[] memory initialL1Configs = new Registry.InitialL1Configuration[](0);
        
        // Deploy Registry contract
        Registry registryContract = new Registry(
            initialOwner,
            initialL2Configs,
            initialL1Configs
        );
        
        vm.stopBroadcast();

        return address(registryContract);
    }
}