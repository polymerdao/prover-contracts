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

import {BitMaps} from "@openzeppelin/contracts/utils/structs/BitMaps.sol";
import {IRegistry} from "../../../interfaces/IRegistry.sol";
import {L2Configuration, L1Configuration, Type} from "../../../libs/RegistryTypes.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Pausable} from "@openzeppelin/contracts/security/Pausable.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

contract Registry is IRegistry, Ownable, Pausable, AccessControl {
    mapping(uint256 => bytes32) public l2ChainConfigurationHashMap;

    mapping(uint256 => L2Configuration) public l2ChainConfigurations;

    mapping(uint256 => bytes32) public l1ChainConfigurationHashMap;

    mapping(uint256 => L1Configuration) public l1ChainConfigurations;

    BitMaps.BitMap internal irrevocableChainIDBitmap;

    bytes32 private constant CHAIN_ROLE_PREFIX = keccak256("CHAIN_ROLE");

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

    error InvalidRange(uint256 startChainID, uint256 stopChainID);

    constructor(
        address _initialOwner,
        InitialL2Configuration[] memory _initialL2Configurations,
        InitialL1Configuration[] memory _initialL1Configurations
    ) {
        if (_initialOwner == address(0)) {
            revert("Ownable: new owner is the zero address");
        }
        _transferOwnership(_initialOwner);

        // Grant the default admin role to the owner
        _setupRole(DEFAULT_ADMIN_ROLE, _initialOwner);

        // Initialize configurations
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
        _setL2ChainConfiguration(_chainID, _config);
    }

    modifier isRevocable(uint256 _chainID) {
        require(BitMaps.get(irrevocableChainIDBitmap, _chainID) == false, "ChainID is irrevocable");
        _;
    }

    /**
     * @dev Get chain-specific role identifier based on chain ID
     */
    function _getChainRole(uint256 _chainID) internal pure returns (bytes32) {
        return keccak256(abi.encode(CHAIN_ROLE_PREFIX, _chainID));
    }

    function grantChainID(address _grantee, uint256 _chainID) external onlyOwner isRevocable(_chainID) {
        return _grantChainID(_grantee, _chainID);
    }

    function grantChainIDIrrevocable(address _grantee, uint256 _chainID) external onlyOwner isRevocable(_chainID) {
        return _grantChainIDIrrevocable(_grantee, _chainID);
    }

    function _grantChainID(address _grantee, uint256 _chainID) internal isRevocable(_chainID) {
        bytes32 role = _getChainRole(_chainID);
        _grantRole(role, _grantee);
    }

    function _grantChainIDIrrevocable(address _grantee, uint256 _chainID) internal isRevocable(_chainID) {
        BitMaps.set(irrevocableChainIDBitmap, _chainID);
        bytes32 role = _getChainRole(_chainID);
        _grantRole(role, _grantee);
        emit NewIrrevocableGrantee(_chainID, _grantee);
    }

    function grantChainIDRange(address _grantee, uint256 _startChainID, uint256 _stopChainID) external onlyOwner {
        if (_startChainID > _stopChainID) {
            revert InvalidRange(_startChainID, _stopChainID);
        }
        for (uint256 i = _startChainID; i <= _stopChainID; i++) {
            _grantChainID(_grantee, i);
        }
    }

    function grantChainIDRangeIrrevocable(address _grantee, uint256 _startChainID, uint256 _stopChainID)
        external
        onlyOwner
    {
        if (_startChainID > _stopChainID) {
            revert InvalidRange(_startChainID, _stopChainID);
        }
        for (uint256 i = _startChainID; i <= _stopChainID; i++) {
            _grantChainIDIrrevocable(_grantee, i);
        }
    }

    function _isGrantee(address _grantee, uint256 _chainID) internal view returns (bool) {
        bytes32 role = _getChainRole(_chainID);
        return hasRole(role, _grantee);
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

    function getL2ConfigType(uint256 _chainID) external view returns (Type) {
        return l2ChainConfigurations[_chainID].l2Type;
    }

    function getL1BlockHashOracle(uint256 _chainID) external view returns (address) {
        return l1ChainConfigurations[_chainID].blockHashOracle;
    }
}
