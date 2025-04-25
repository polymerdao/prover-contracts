// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../contracts/core/native_fallback/L2/OPStackCannonProver.sol";

/**
 * @title ComputeOPStackCannonProverCreate2Address
 * @notice Computes the expected deterministic address for OPStackCannonProver contract 
 * @dev This script computes and displays the CREATE2 address where the OPStackCannonProver
 *      will be deployed when using the DeployOPStackCannonProverCreate2 script.
 *      This can be used to verify the expected address before actual deployment.
 *      
 *      To run: forge script script/ComputeOPStackCannonProverCreate2Address.s.sol --private-key $PRIVATE_KEY
 */
contract ComputeOPStackCannonProverCreate2Address is Script {
    function run() external {
        // Get the deployer address from the private key
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        
        // Use the same salt as in DeployOPStackCannonProverCreate2.s.sol
        bytes32 salt = keccak256(abi.encodePacked("polymer-opstack-cannon-prover"));
        
        // Compute the initialization code hash
        bytes memory creationCode = type(OPStackCannonProver).creationCode;
        bytes32 initCodeHash = keccak256(creationCode);
        
        // Compute the deterministic address using CREATE2 formula
        bytes32 data = keccak256(
            abi.encodePacked(
                bytes1(0xff),
                deployer,
                salt,
                initCodeHash
            )
        );
        
        // Convert the result to an address (take the last 20 bytes)
        address predictedAddress = address(uint160(uint256(data)));
        
        // Display the information
        console.log("Deployer address:", deployer);
        console.log("Salt (hex):", vm.toString(salt));
        console.log("Init code hash:", vm.toString(initCodeHash));
        console.log("Predicted OPStackCannonProver address:", predictedAddress);
    }
}