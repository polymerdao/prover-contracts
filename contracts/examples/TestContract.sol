// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.15;

/**
 * @title TestContract
 * @dev A simple contract for testing storage proofs
 */
contract TestContract {
    // This value will be stored at slot 0
    uint256 public value;
    
    /**
     * @dev Set the value in storage
     * @param _value The value to store
     */
    function setValue(uint256 _value) external {
        value = _value;
    }
    
    /**
     * @dev Get the current value
     * @return The stored value
     */
    function getValue() external view returns (uint256) {
        return value;
    }
    
    /**
     * @dev Constructor sets an initial value
     * @param _initialValue Initial value to store
     */
    constructor(uint256 _initialValue) {
        value = _initialValue;
    }
}