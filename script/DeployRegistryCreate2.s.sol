// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/core/native_fallback/L1/Registry.sol";
import "../contracts/libs/RegistryTypes.sol";

/**
 * @title DeployRegistryCreate2Script
 * @notice Deploys the Registry contract with CREATE2 for deterministic addresses
 * @dev This script deploys the Registry contract using CREATE2 with a fixed salt.
 *      This is useful for having a consistent Registry address across multiple L1 chains.
 */
contract DeployRegistryCreate2Script is Script {
    function run() external returns (address registry) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address initialOwner = vm.addr(deployerPrivateKey);
        
        // Use a fixed salt to ensure consistent deployment addresses
        // In production, this salt value should be carefully chosen and never changed
        bytes32 salt = keccak256(abi.encodePacked("polymer-registry"));
        
        console.log("Using CREATE2 salt:", vm.toString(salt));
        console.log("Initial owner:", initialOwner);
        
        vm.startBroadcast(deployerPrivateKey);

        // Create empty arrays for initial configurations
        Registry.InitialL2Configuration[] memory initialL2Configs = new Registry.InitialL2Configuration[](0);
        Registry.InitialL1Configuration[] memory initialL1Configs = new Registry.InitialL1Configuration[](0);
        
        // Deploy Registry contract using CREATE2
        Registry registryContract = new Registry{salt: salt}(
            initialOwner,
            initialL2Configs,
            initialL1Configs
        );
        
        console.log("Deployed Registry using CREATE2 at address:", address(registryContract));
        
        vm.stopBroadcast();

        return address(registryContract);
    }
}