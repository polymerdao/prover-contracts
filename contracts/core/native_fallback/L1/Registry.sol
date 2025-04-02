// SPDX-License-Identifier: Apache-2.0
/*
 * Copyright 2024, Polymer Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity 0.8.15;

import {IRegistry} from "../../../interfaces/IRegistry.sol";
import {L2Configuration, L1Configuration} from "../../../libs/RegistryTypes.sol";

contract Registry is IRegistry {
    address public immutable multiSigOwner;

    bool public immutable onlyAdd;

    mapping(uint256 => bytes32) public l2ChainConfigurationHashMap;

    mapping(uint256 => L2Configuration) public l2ChainConfigurations;

    mapping(uint256 => bytes32) public l1ChainConfigurationHashMap;

    mapping(uint256 => L1Configuration) public l1ChainConfigurations;

    mapping(address => mapping(uint256 => uint256)) public granteeBitmap;

    mapping(uint256 => uint256) public irrevocableChainIDBitmap;

    event NewGrantee(uint256 indexed chainID, address indexed grantee);

    event NewIrrevocableGrantee(uint256 indexed chainID, address indexed grantee);

    event L2ChainConfigurationUpdated(uint256 indexed chainID, bytes32 indexed configHash);

    event L1ChainConfigurationUpdated(uint256 indexed chainID, bytes32 indexed configHash);

    struct InitialL2Configuration {
        uint256 chainID;
        L2Configuration config;
    }

    struct InitialL1Configuration {
        uint256 chainID;
        L1Configuration config;
    }

    modifier onlyMultiSig() {
        require(msg.sender == multiSigOwner, "Not authorized");
        _;
    }

    constructor(
        address _multiSigOwner,
        bool _onlyAdd,
        InitialL2Configuration[] memory _initialL2Configurations,
        InitialL1Configuration[] memory _initialL1Configurations
    ) {
        require(_multiSigOwner != address(0), "Invalid multiSig address");
        multiSigOwner = _multiSigOwner;
        onlyAdd = _onlyAdd;
        for (uint256 i = 0; i < _initialL2Configurations.length; ++i) {
            _setL2ChainConfiguration(_initialL2Configurations[i].chainID, _initialL2Configurations[i].config);
        }
        for (uint256 i = 0; i < _initialL1Configurations.length; ++i) {
            _setL1ChainConfiguration(_initialL1Configurations[i].chainID, _initialL1Configurations[i].config);
        }
    }

    function _setL1ChainConfiguration(uint256 _chainID, L1Configuration memory _config) internal {
        bytes32 _hash = keccak256(abi.encode(_config));
        l1ChainConfigurationHashMap[_chainID] = _hash;
        l1ChainConfigurations[_chainID] = _config;
        emit L1ChainConfigurationUpdated(_chainID, _hash);
    }

    function updateL1ChainConfiguration(uint256 _chainID, L1Configuration calldata _config) external {
        require(_isGrantee(msg.sender, _chainID), "Not authorized");
        require(_isAllowedL1ConfigUpdate(_chainID), "Existing config cannot be updated");
        _setL1ChainConfiguration(_chainID, _config);
    }

    function _setL2ChainConfiguration(uint256 _chainID, L2Configuration memory _config) internal {
        bytes32 _hash = keccak256(abi.encode(_config));
        l2ChainConfigurationHashMap[_chainID] = _hash;
        l2ChainConfigurations[_chainID] = _config;
        emit L2ChainConfigurationUpdated(_chainID, _hash);
    }

    function updateL2ChainConfiguration(uint256 _chainID, L2Configuration calldata _config) external {
        require(_isGrantee(msg.sender, _chainID), "Not authorized");
        require(_isAllowedL2ConfigUpdate(_chainID), "Existing config cannot be updated");
        _setL2ChainConfiguration(_chainID, _config);
    }

    modifier isRevocable(uint256 _chainID) {
        require((irrevocableChainIDBitmap[_chainID / 256] & (1 << (_chainID % 256))) == 0, "ChainID is irrevocable");
        _;
    }

    function grantChainID(address _grantee, uint256 _chainID) external onlyMultiSig isRevocable(_chainID) {
        return _grantChainID(_grantee, _chainID);
    }

    function grantChainIDIrrevocable(address _grantee, uint256 _chainID) external onlyMultiSig isRevocable(_chainID) {
        return _grantChainIDIrrevocable(_grantee, _chainID);
    }

    function _grantChainID(address _grantee, uint256 _chainID) internal isRevocable(_chainID) {
        uint256 _bucket = _chainID / 256;
        uint256 _bit = _chainID % 256;
        granteeBitmap[_grantee][_bucket] |= (1 << _bit);
        emit NewGrantee(_chainID, _grantee);
    }

    function _grantChainIDIrrevocable(address _grantee, uint256 _chainID) internal isRevocable(_chainID) {
        uint256 _bucket = _chainID / 256;
        uint256 _bit = _chainID % 256;
        granteeBitmap[_grantee][_bucket] |= (1 << _bit);
        irrevocableChainIDBitmap[_bucket] |= (1 << _bit);
        emit NewIrrevocableGrantee(_chainID, _grantee);
    }

    function grantChainIDRange(address _grantee, uint256 _startChainID, uint256 _stopChainID) external onlyMultiSig {
        require(_startChainID <= _stopChainID, "Invalid range");
        for (uint256 i = _startChainID; i <= _stopChainID; i++) {
            _grantChainID(_grantee, i);
        }
    }

    function grantChainIDRangeIrrevocable(address _grantee, uint256 _startChainID, uint256 _stopChainID)
        external
        onlyMultiSig
    {
        require(_startChainID <= _stopChainID, "Invalid range");
        for (uint256 i = _startChainID; i <= _stopChainID; i++) {
            _grantChainIDIrrevocable(_grantee, i);
        }
    }

    function _isGrantee(address _grantee, uint256 _chainID) internal view returns (bool) {
        uint256 _bucket = _chainID / 256;
        uint256 _bit = _chainID % 256;
        return (granteeBitmap[_grantee][_bucket] & (1 << _bit)) != 0;
    }

    function _isAllowedL2ConfigUpdate(uint256 _chainID) internal view returns (bool) {
        if (!onlyAdd) {
            return true;
        }
        return l2ChainConfigurationHashMap[_chainID] == bytes32(0);
    }

    function _isAllowedL1ConfigUpdate(uint256 _chainID) internal view returns (bool) {
        if (!onlyAdd) {
            return true;
        }
        return l1ChainConfigurationHashMap[_chainID] == bytes32(0);
    }

    function isGrantee(address _grantee, uint256 _chainID) external view returns (bool) {
        return _isGrantee(_grantee, _chainID);
    }

    function getL2ConfigAddresses(uint256 _chainID) external view returns (address[] memory) {
        return l2ChainConfigurations[_chainID].addresses;
    }

    function getL2ConfigStorageSlots(uint256 _chainID) external view returns (uint256[] memory) {
        return l2ChainConfigurations[_chainID].storageSlots;
    }
}
