// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface IL1Block {
    /// @notice The latest L1 blockhash.
    function hash() external view returns (bytes32);
}
