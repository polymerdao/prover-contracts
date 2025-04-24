// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

import "forge-std/Script.sol";
import "../contracts/examples/TestContract.sol";

contract DeployTestContractScript is Script {
    function run() external returns (address testContract) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        // Get initial value from environment, default to 42
        uint256 initialValue = vm.envOr("TEST_CONTRACT_INITIAL_VALUE", uint256(42));
        
        vm.startBroadcast(deployerPrivateKey);
        
        // Deploy test contract with initial value
        TestContract testContractInstance = new TestContract(initialValue);
        
        vm.stopBroadcast();
        
        return address(testContractInstance);
    }
}