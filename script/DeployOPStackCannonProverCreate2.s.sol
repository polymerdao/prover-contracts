// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/core/native_fallback/L2/OPStackCannonProver.sol";

/**
 * @title DeployOPStackCannonProverCreate2Script
 * @notice Deploys the OPStackCannonProver contract with CREATE2 for deterministic addresses
 * @dev This script deploys the OPStackCannonProver contract using CREATE2 with a fixed salt.
 *      When this script is run on multiple L2 chains with the same:
 *      1. Bytecode (same contract, compiler version, and settings)
 *      2. Deployer address (same private key)
 *      3. Salt value (hardcoded in this script)
 *      The contract will be deployed to the same address on all chains.
 *
 *      CREATE2 address is determined by: keccak256(0xff ++ deployer ++ salt ++ keccak256(init_code))
 *      For more info, see: https://book.getfoundry.sh/guides/deterministic-deployments-using-create2
 */
contract DeployOPStackCannonProverCreate2Script is Script {
    function run() external returns (address opStackCannonProver) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        // Use a fixed salt to ensure consistent deployment addresses across all chains
        // This salt value should not be changed once in production
        bytes32 salt = keccak256(abi.encodePacked("polymer-opstack-cannon-prover"));
        
        console.log("Using CREATE2 salt:", vm.toString(salt));
        
        vm.startBroadcast(deployerPrivateKey);
        
        // Deploy with CREATE2 using the salt parameter
        OPStackCannonProver proverContract = new OPStackCannonProver{salt: salt}();
        
        console.log("Deployed OPStackCannonProver using CREATE2 at address:", address(proverContract));
        
        vm.stopBroadcast();
        
        return address(proverContract);
    }
}