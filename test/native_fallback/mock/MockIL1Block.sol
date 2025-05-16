// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.28;

import {IL1Block} from "../../../contracts/interfaces/IL1Block.sol";

contract MockIL1Block is IL1Block {
    bytes32 private mockHash;

    function setBlockHash(bytes32 _hash) external {
        mockHash = _hash;
    }

    function hash() external view returns (bytes32) {
        return mockHash;
    }

    function version() external pure returns (string memory) {
        return "1.0.0";
    }
}
