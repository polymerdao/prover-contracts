// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/core/native_fallback/L1/Registry.sol";
import "../contracts/core/native_fallback/L2/NativeProver.sol";
import "../contracts/core/native_fallback/L2/OPStackCannonProver.sol";
import "../contracts/core/native_fallback/L2/OPStackBedrockProver.sol";
import "../contracts/examples/TestContract.sol";
import "../contracts/libs/RegistryTypes.sol";

/**
 * @title ComputeCreate2Addresses
 * @notice Computes the expected CREATE2 addresses for all contracts used in the system
 * @dev This script doesn't deploy any contracts, it just computes and outputs the expected
 *      addresses based on the same salts and deployer address used in the deployment scripts.
 */
contract ComputeCreate2Addresses is Script {
    function run() external {
        console.log("===== Computing expected CREATE2 addresses =====");
        console.log("These addresses can be used to verify deterministic deployments");
        
        // Get deployer address (we'll use the first account from anvil/hardhat by default)
        uint256 deployerPrivateKey = vm.envOr("PRIVATE_KEY", uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80));
        address deployer = vm.addr(deployerPrivateKey);
        console.log("Deployer address:", deployer);
        
        // Compute Registry address
        bytes32 registrySalt = keccak256(abi.encodePacked("polymer-registry"));
        address owner = deployer; // Owner is the same as deployer
        
        // Create empty arrays for initial configurations
        Registry.InitialL2Configuration[] memory initialL2Configs = new Registry.InitialL2Configuration[](0);
        Registry.InitialL1Configuration[] memory initialL1Configs = new Registry.InitialL1Configuration[](0);
        
        // Compute init code with constructor args
        bytes memory registryInitCode = abi.encodePacked(
            type(Registry).creationCode,
            abi.encode(owner, initialL2Configs, initialL1Configs)
        );
        
        address registryAddr = computeCreate2Address(
            deployer,
            registrySalt,
            registryInitCode
        );
        console.log("Registry CREATE2 address:", registryAddr);
        
        // Compute OPStackCannonProver address
        bytes32 cannonSalt = keccak256(abi.encodePacked("polymer-opstack-cannon-prover"));
        address cannonProverAddr = computeCreate2Address(
            deployer,
            cannonSalt,
            type(OPStackCannonProver).creationCode
        );
        console.log("OPStackCannonProver CREATE2 address:", cannonProverAddr);
        
        // Compute OPStackBedrockProver address
        bytes32 bedrockSalt = keccak256(abi.encodePacked("polymer-opstack-bedrock-prover"));
        address bedrockProverAddr = computeCreate2Address(
            deployer,
            bedrockSalt,
            type(OPStackBedrockProver).creationCode
        );
        console.log("OPStackBedrockProver CREATE2 address:", bedrockProverAddr);
        
        // Compute NativeProver address with standard params
        // Note: This is more complex due to constructor args
        uint256 l1ChainId = 1234;
        address blockHashOracle = 0x4200000000000000000000000000000000000015;
        address settlementRegistry = 0x5FbDB2315678afecb367f032d93F642f64180aa3; // Example address
        
        // Create salt for NativeProver
        bytes32 nativeSalt = keccak256(abi.encodePacked(
            "polymer-native-prover",
            l1ChainId,
            blockHashOracle,
            settlementRegistry
        ));
        
        // Create L1Configuration struct
        L1Configuration memory l1Config = L1Configuration({
            blockHashOracle: blockHashOracle,
            settlementRegistry: settlementRegistry,
            settlementRegistryL2ConfigMappingSlot: 1,
            settlementRegistryL1ConfigMappingSlot: 2,
            settlementBlocksDelay: 10
        });
        
        // Create empty array for initial L2 configurations
        NativeProver.InitialL2Configuration[] memory initialL2NativeConfigs = new NativeProver.InitialL2Configuration[](0);
        
        // Compute init code with constructor args
        bytes memory nativeProverInitCode = abi.encodePacked(
            type(NativeProver).creationCode,
            abi.encode(l1ChainId, l1Config, initialL2Configs)
        );
        
        address nativeProverAddr = computeCreate2Address(
            deployer,
            nativeSalt,
            nativeProverInitCode
        );
        console.log("NativeProver CREATE2 address with standard params:", nativeProverAddr);
        
        // Compute TestContract address with standard value
        uint256 testContractValue = 123456789;
        bytes32 testSalt = keccak256(abi.encodePacked("polymer-test-contract", testContractValue));
        
        bytes memory testContractInitCode = abi.encodePacked(
            type(TestContract).creationCode,
            abi.encode(testContractValue)
        );
        
        address testContractAddr = computeCreate2Address(
            deployer,
            testSalt,
            testContractInitCode
        );
        console.log("TestContract CREATE2 address with value", testContractValue, ":", testContractAddr);
    }
    
    function computeCreate2Address(
        address deployer,
        bytes32 salt,
        bytes memory bytecode
    ) internal pure returns (address) {
        bytes32 hash = keccak256(
            abi.encodePacked(
                bytes1(0xff),
                deployer,
                salt,
                keccak256(bytecode)
            )
        );
        
        return address(uint160(uint256(hash)));
    }
}