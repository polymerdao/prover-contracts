// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// L1Configuration is an auto generated low-level Go binding around an user-defined struct.
type L1Configuration struct {
	BlockHashOracle                       common.Address
	SettlementRegistry                    common.Address
	SettlementRegistryL2ConfigMappingSlot *big.Int
	SettlementRegistryL1ConfigMappingSlot *big.Int
}

// L2Configuration is an auto generated low-level Go binding around an user-defined struct.
type L2Configuration struct {
	Prover               common.Address
	Addresses            []common.Address
	StorageSlots         []*big.Int
	VersionNumber        *big.Int
	FinalityDelaySeconds *big.Int
	L2Type               uint8
}

// RegistryInitialL1Configuration is an auto generated low-level Go binding around an user-defined struct.
type RegistryInitialL1Configuration struct {
	ChainID *big.Int
	Config  L1Configuration
}

// RegistryInitialL2Configuration is an auto generated low-level Go binding around an user-defined struct.
type RegistryInitialL2Configuration struct {
	ChainID *big.Int
	Config  L2Configuration
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialL2Configurations\",\"type\":\"tuple[]\",\"internalType\":\"structRegistry.InitialL2Configuration[]\",\"components\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"enumType\"}]}]},{\"name\":\"_initialL1Configurations\",\"type\":\"tuple[]\",\"internalType\":\"structRegistry.InitialL1Configuration[]\",\"components\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structL1Configuration\",\"components\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getL1BlockHashOracle\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getL2ConfigAddresses\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getL2ConfigStorageSlots\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getL2ConfigType\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantChainID\",\"inputs\":[{\"name\":\"_grantee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantChainIDIrrevocable\",\"inputs\":[{\"name\":\"_grantee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantChainIDRange\",\"inputs\":[{\"name\":\"_grantee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_startChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stopChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantChainIDRangeIrrevocable\",\"inputs\":[{\"name\":\"_grantee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_startChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stopChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isIrrevocableGrantee\",\"inputs\":[{\"name\":\"_grantee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRevocableGrantee\",\"inputs\":[{\"name\":\"_grantee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1ChainConfigurationHashMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1ChainConfigurations\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2ChainConfigurationHashMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2ChainConfigurations\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"enumType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeChainID\",\"inputs\":[{\"name\":\"_grantee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateL1ChainConfiguration\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structL1Configuration\",\"components\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateL2ChainConfiguration\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"enumType\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"L1ChainConfigurationUpdated\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"configHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L2ChainConfigurationUpdated\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"configHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidRange\",\"inputs\":[{\"name\":\"startChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stopChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registry *RegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registry *RegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Registry.Contract.DEFAULTADMINROLE(&_Registry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registry *RegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Registry.Contract.DEFAULTADMINROLE(&_Registry.CallOpts)
}

// GetL1BlockHashOracle is a free data retrieval call binding the contract method 0x7fc0ad31.
//
// Solidity: function getL1BlockHashOracle(uint256 _chainID) view returns(address)
func (_Registry *RegistryCaller) GetL1BlockHashOracle(opts *bind.CallOpts, _chainID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getL1BlockHashOracle", _chainID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1BlockHashOracle is a free data retrieval call binding the contract method 0x7fc0ad31.
//
// Solidity: function getL1BlockHashOracle(uint256 _chainID) view returns(address)
func (_Registry *RegistrySession) GetL1BlockHashOracle(_chainID *big.Int) (common.Address, error) {
	return _Registry.Contract.GetL1BlockHashOracle(&_Registry.CallOpts, _chainID)
}

// GetL1BlockHashOracle is a free data retrieval call binding the contract method 0x7fc0ad31.
//
// Solidity: function getL1BlockHashOracle(uint256 _chainID) view returns(address)
func (_Registry *RegistryCallerSession) GetL1BlockHashOracle(_chainID *big.Int) (common.Address, error) {
	return _Registry.Contract.GetL1BlockHashOracle(&_Registry.CallOpts, _chainID)
}

// GetL2ConfigAddresses is a free data retrieval call binding the contract method 0x974ec7fc.
//
// Solidity: function getL2ConfigAddresses(uint256 _chainID) view returns(address[])
func (_Registry *RegistryCaller) GetL2ConfigAddresses(opts *bind.CallOpts, _chainID *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getL2ConfigAddresses", _chainID)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetL2ConfigAddresses is a free data retrieval call binding the contract method 0x974ec7fc.
//
// Solidity: function getL2ConfigAddresses(uint256 _chainID) view returns(address[])
func (_Registry *RegistrySession) GetL2ConfigAddresses(_chainID *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetL2ConfigAddresses(&_Registry.CallOpts, _chainID)
}

// GetL2ConfigAddresses is a free data retrieval call binding the contract method 0x974ec7fc.
//
// Solidity: function getL2ConfigAddresses(uint256 _chainID) view returns(address[])
func (_Registry *RegistryCallerSession) GetL2ConfigAddresses(_chainID *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetL2ConfigAddresses(&_Registry.CallOpts, _chainID)
}

// GetL2ConfigStorageSlots is a free data retrieval call binding the contract method 0xa4f4ae39.
//
// Solidity: function getL2ConfigStorageSlots(uint256 _chainID) view returns(uint256[])
func (_Registry *RegistryCaller) GetL2ConfigStorageSlots(opts *bind.CallOpts, _chainID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getL2ConfigStorageSlots", _chainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetL2ConfigStorageSlots is a free data retrieval call binding the contract method 0xa4f4ae39.
//
// Solidity: function getL2ConfigStorageSlots(uint256 _chainID) view returns(uint256[])
func (_Registry *RegistrySession) GetL2ConfigStorageSlots(_chainID *big.Int) ([]*big.Int, error) {
	return _Registry.Contract.GetL2ConfigStorageSlots(&_Registry.CallOpts, _chainID)
}

// GetL2ConfigStorageSlots is a free data retrieval call binding the contract method 0xa4f4ae39.
//
// Solidity: function getL2ConfigStorageSlots(uint256 _chainID) view returns(uint256[])
func (_Registry *RegistryCallerSession) GetL2ConfigStorageSlots(_chainID *big.Int) ([]*big.Int, error) {
	return _Registry.Contract.GetL2ConfigStorageSlots(&_Registry.CallOpts, _chainID)
}

// GetL2ConfigType is a free data retrieval call binding the contract method 0x5338efd4.
//
// Solidity: function getL2ConfigType(uint256 _chainID) view returns(uint8)
func (_Registry *RegistryCaller) GetL2ConfigType(opts *bind.CallOpts, _chainID *big.Int) (uint8, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getL2ConfigType", _chainID)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetL2ConfigType is a free data retrieval call binding the contract method 0x5338efd4.
//
// Solidity: function getL2ConfigType(uint256 _chainID) view returns(uint8)
func (_Registry *RegistrySession) GetL2ConfigType(_chainID *big.Int) (uint8, error) {
	return _Registry.Contract.GetL2ConfigType(&_Registry.CallOpts, _chainID)
}

// GetL2ConfigType is a free data retrieval call binding the contract method 0x5338efd4.
//
// Solidity: function getL2ConfigType(uint256 _chainID) view returns(uint8)
func (_Registry *RegistryCallerSession) GetL2ConfigType(_chainID *big.Int) (uint8, error) {
	return _Registry.Contract.GetL2ConfigType(&_Registry.CallOpts, _chainID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registry *RegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registry *RegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Registry.Contract.GetRoleAdmin(&_Registry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registry *RegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Registry.Contract.GetRoleAdmin(&_Registry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Registry.Contract.HasRole(&_Registry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Registry.Contract.HasRole(&_Registry.CallOpts, role, account)
}

// IsIrrevocableGrantee is a free data retrieval call binding the contract method 0xbc839f07.
//
// Solidity: function isIrrevocableGrantee(address _grantee, uint256 _chainID) view returns(bool)
func (_Registry *RegistryCaller) IsIrrevocableGrantee(opts *bind.CallOpts, _grantee common.Address, _chainID *big.Int) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isIrrevocableGrantee", _grantee, _chainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIrrevocableGrantee is a free data retrieval call binding the contract method 0xbc839f07.
//
// Solidity: function isIrrevocableGrantee(address _grantee, uint256 _chainID) view returns(bool)
func (_Registry *RegistrySession) IsIrrevocableGrantee(_grantee common.Address, _chainID *big.Int) (bool, error) {
	return _Registry.Contract.IsIrrevocableGrantee(&_Registry.CallOpts, _grantee, _chainID)
}

// IsIrrevocableGrantee is a free data retrieval call binding the contract method 0xbc839f07.
//
// Solidity: function isIrrevocableGrantee(address _grantee, uint256 _chainID) view returns(bool)
func (_Registry *RegistryCallerSession) IsIrrevocableGrantee(_grantee common.Address, _chainID *big.Int) (bool, error) {
	return _Registry.Contract.IsIrrevocableGrantee(&_Registry.CallOpts, _grantee, _chainID)
}

// IsRevocableGrantee is a free data retrieval call binding the contract method 0xadfb4594.
//
// Solidity: function isRevocableGrantee(address _grantee, uint256 _chainID) view returns(bool)
func (_Registry *RegistryCaller) IsRevocableGrantee(opts *bind.CallOpts, _grantee common.Address, _chainID *big.Int) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isRevocableGrantee", _grantee, _chainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRevocableGrantee is a free data retrieval call binding the contract method 0xadfb4594.
//
// Solidity: function isRevocableGrantee(address _grantee, uint256 _chainID) view returns(bool)
func (_Registry *RegistrySession) IsRevocableGrantee(_grantee common.Address, _chainID *big.Int) (bool, error) {
	return _Registry.Contract.IsRevocableGrantee(&_Registry.CallOpts, _grantee, _chainID)
}

// IsRevocableGrantee is a free data retrieval call binding the contract method 0xadfb4594.
//
// Solidity: function isRevocableGrantee(address _grantee, uint256 _chainID) view returns(bool)
func (_Registry *RegistryCallerSession) IsRevocableGrantee(_grantee common.Address, _chainID *big.Int) (bool, error) {
	return _Registry.Contract.IsRevocableGrantee(&_Registry.CallOpts, _grantee, _chainID)
}

// L1ChainConfigurationHashMap is a free data retrieval call binding the contract method 0xd2eca660.
//
// Solidity: function l1ChainConfigurationHashMap(uint256 ) view returns(bytes32)
func (_Registry *RegistryCaller) L1ChainConfigurationHashMap(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "l1ChainConfigurationHashMap", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1ChainConfigurationHashMap is a free data retrieval call binding the contract method 0xd2eca660.
//
// Solidity: function l1ChainConfigurationHashMap(uint256 ) view returns(bytes32)
func (_Registry *RegistrySession) L1ChainConfigurationHashMap(arg0 *big.Int) ([32]byte, error) {
	return _Registry.Contract.L1ChainConfigurationHashMap(&_Registry.CallOpts, arg0)
}

// L1ChainConfigurationHashMap is a free data retrieval call binding the contract method 0xd2eca660.
//
// Solidity: function l1ChainConfigurationHashMap(uint256 ) view returns(bytes32)
func (_Registry *RegistryCallerSession) L1ChainConfigurationHashMap(arg0 *big.Int) ([32]byte, error) {
	return _Registry.Contract.L1ChainConfigurationHashMap(&_Registry.CallOpts, arg0)
}

// L1ChainConfigurations is a free data retrieval call binding the contract method 0xa9a9730a.
//
// Solidity: function l1ChainConfigurations(uint256 ) view returns(address blockHashOracle, address settlementRegistry, uint256 settlementRegistryL2ConfigMappingSlot, uint256 settlementRegistryL1ConfigMappingSlot)
func (_Registry *RegistryCaller) L1ChainConfigurations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockHashOracle                       common.Address
	SettlementRegistry                    common.Address
	SettlementRegistryL2ConfigMappingSlot *big.Int
	SettlementRegistryL1ConfigMappingSlot *big.Int
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "l1ChainConfigurations", arg0)

	outstruct := new(struct {
		BlockHashOracle                       common.Address
		SettlementRegistry                    common.Address
		SettlementRegistryL2ConfigMappingSlot *big.Int
		SettlementRegistryL1ConfigMappingSlot *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockHashOracle = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SettlementRegistry = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.SettlementRegistryL2ConfigMappingSlot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SettlementRegistryL1ConfigMappingSlot = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// L1ChainConfigurations is a free data retrieval call binding the contract method 0xa9a9730a.
//
// Solidity: function l1ChainConfigurations(uint256 ) view returns(address blockHashOracle, address settlementRegistry, uint256 settlementRegistryL2ConfigMappingSlot, uint256 settlementRegistryL1ConfigMappingSlot)
func (_Registry *RegistrySession) L1ChainConfigurations(arg0 *big.Int) (struct {
	BlockHashOracle                       common.Address
	SettlementRegistry                    common.Address
	SettlementRegistryL2ConfigMappingSlot *big.Int
	SettlementRegistryL1ConfigMappingSlot *big.Int
}, error) {
	return _Registry.Contract.L1ChainConfigurations(&_Registry.CallOpts, arg0)
}

// L1ChainConfigurations is a free data retrieval call binding the contract method 0xa9a9730a.
//
// Solidity: function l1ChainConfigurations(uint256 ) view returns(address blockHashOracle, address settlementRegistry, uint256 settlementRegistryL2ConfigMappingSlot, uint256 settlementRegistryL1ConfigMappingSlot)
func (_Registry *RegistryCallerSession) L1ChainConfigurations(arg0 *big.Int) (struct {
	BlockHashOracle                       common.Address
	SettlementRegistry                    common.Address
	SettlementRegistryL2ConfigMappingSlot *big.Int
	SettlementRegistryL1ConfigMappingSlot *big.Int
}, error) {
	return _Registry.Contract.L1ChainConfigurations(&_Registry.CallOpts, arg0)
}

// L2ChainConfigurationHashMap is a free data retrieval call binding the contract method 0xeaef16e1.
//
// Solidity: function l2ChainConfigurationHashMap(uint256 ) view returns(bytes32)
func (_Registry *RegistryCaller) L2ChainConfigurationHashMap(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "l2ChainConfigurationHashMap", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2ChainConfigurationHashMap is a free data retrieval call binding the contract method 0xeaef16e1.
//
// Solidity: function l2ChainConfigurationHashMap(uint256 ) view returns(bytes32)
func (_Registry *RegistrySession) L2ChainConfigurationHashMap(arg0 *big.Int) ([32]byte, error) {
	return _Registry.Contract.L2ChainConfigurationHashMap(&_Registry.CallOpts, arg0)
}

// L2ChainConfigurationHashMap is a free data retrieval call binding the contract method 0xeaef16e1.
//
// Solidity: function l2ChainConfigurationHashMap(uint256 ) view returns(bytes32)
func (_Registry *RegistryCallerSession) L2ChainConfigurationHashMap(arg0 *big.Int) ([32]byte, error) {
	return _Registry.Contract.L2ChainConfigurationHashMap(&_Registry.CallOpts, arg0)
}

// L2ChainConfigurations is a free data retrieval call binding the contract method 0x63281a15.
//
// Solidity: function l2ChainConfigurations(uint256 ) view returns(address prover, uint256 versionNumber, uint256 finalityDelaySeconds, uint8 l2Type)
func (_Registry *RegistryCaller) L2ChainConfigurations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Prover               common.Address
	VersionNumber        *big.Int
	FinalityDelaySeconds *big.Int
	L2Type               uint8
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "l2ChainConfigurations", arg0)

	outstruct := new(struct {
		Prover               common.Address
		VersionNumber        *big.Int
		FinalityDelaySeconds *big.Int
		L2Type               uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Prover = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.VersionNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FinalityDelaySeconds = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.L2Type = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// L2ChainConfigurations is a free data retrieval call binding the contract method 0x63281a15.
//
// Solidity: function l2ChainConfigurations(uint256 ) view returns(address prover, uint256 versionNumber, uint256 finalityDelaySeconds, uint8 l2Type)
func (_Registry *RegistrySession) L2ChainConfigurations(arg0 *big.Int) (struct {
	Prover               common.Address
	VersionNumber        *big.Int
	FinalityDelaySeconds *big.Int
	L2Type               uint8
}, error) {
	return _Registry.Contract.L2ChainConfigurations(&_Registry.CallOpts, arg0)
}

// L2ChainConfigurations is a free data retrieval call binding the contract method 0x63281a15.
//
// Solidity: function l2ChainConfigurations(uint256 ) view returns(address prover, uint256 versionNumber, uint256 finalityDelaySeconds, uint8 l2Type)
func (_Registry *RegistryCallerSession) L2ChainConfigurations(arg0 *big.Int) (struct {
	Prover               common.Address
	VersionNumber        *big.Int
	FinalityDelaySeconds *big.Int
	L2Type               uint8
}, error) {
	return _Registry.Contract.L2ChainConfigurations(&_Registry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registry *RegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registry *RegistrySession) Paused() (bool, error) {
	return _Registry.Contract.Paused(&_Registry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registry *RegistryCallerSession) Paused() (bool, error) {
	return _Registry.Contract.Paused(&_Registry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registry.Contract.SupportsInterface(&_Registry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registry.Contract.SupportsInterface(&_Registry.CallOpts, interfaceId)
}

// GrantChainID is a paid mutator transaction binding the contract method 0x2537bc05.
//
// Solidity: function grantChainID(address _grantee, uint256 _chainID) returns()
func (_Registry *RegistryTransactor) GrantChainID(opts *bind.TransactOpts, _grantee common.Address, _chainID *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "grantChainID", _grantee, _chainID)
}

// GrantChainID is a paid mutator transaction binding the contract method 0x2537bc05.
//
// Solidity: function grantChainID(address _grantee, uint256 _chainID) returns()
func (_Registry *RegistrySession) GrantChainID(_grantee common.Address, _chainID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.GrantChainID(&_Registry.TransactOpts, _grantee, _chainID)
}

// GrantChainID is a paid mutator transaction binding the contract method 0x2537bc05.
//
// Solidity: function grantChainID(address _grantee, uint256 _chainID) returns()
func (_Registry *RegistryTransactorSession) GrantChainID(_grantee common.Address, _chainID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.GrantChainID(&_Registry.TransactOpts, _grantee, _chainID)
}

// GrantChainIDIrrevocable is a paid mutator transaction binding the contract method 0x658ce22d.
//
// Solidity: function grantChainIDIrrevocable(address _grantee, uint256 _chainID) returns()
func (_Registry *RegistryTransactor) GrantChainIDIrrevocable(opts *bind.TransactOpts, _grantee common.Address, _chainID *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "grantChainIDIrrevocable", _grantee, _chainID)
}

// GrantChainIDIrrevocable is a paid mutator transaction binding the contract method 0x658ce22d.
//
// Solidity: function grantChainIDIrrevocable(address _grantee, uint256 _chainID) returns()
func (_Registry *RegistrySession) GrantChainIDIrrevocable(_grantee common.Address, _chainID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.GrantChainIDIrrevocable(&_Registry.TransactOpts, _grantee, _chainID)
}

// GrantChainIDIrrevocable is a paid mutator transaction binding the contract method 0x658ce22d.
//
// Solidity: function grantChainIDIrrevocable(address _grantee, uint256 _chainID) returns()
func (_Registry *RegistryTransactorSession) GrantChainIDIrrevocable(_grantee common.Address, _chainID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.GrantChainIDIrrevocable(&_Registry.TransactOpts, _grantee, _chainID)
}

// GrantChainIDRange is a paid mutator transaction binding the contract method 0xe8b24600.
//
// Solidity: function grantChainIDRange(address _grantee, uint256 _startChainID, uint256 _stopChainID) returns()
func (_Registry *RegistryTransactor) GrantChainIDRange(opts *bind.TransactOpts, _grantee common.Address, _startChainID *big.Int, _stopChainID *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "grantChainIDRange", _grantee, _startChainID, _stopChainID)
}

// GrantChainIDRange is a paid mutator transaction binding the contract method 0xe8b24600.
//
// Solidity: function grantChainIDRange(address _grantee, uint256 _startChainID, uint256 _stopChainID) returns()
func (_Registry *RegistrySession) GrantChainIDRange(_grantee common.Address, _startChainID *big.Int, _stopChainID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.GrantChainIDRange(&_Registry.TransactOpts, _grantee, _startChainID, _stopChainID)
}

// GrantChainIDRange is a paid mutator transaction binding the contract method 0xe8b24600.
//
// Solidity: function grantChainIDRange(address _grantee, uint256 _startChainID, uint256 _stopChainID) returns()
func (_Registry *RegistryTransactorSession) GrantChainIDRange(_grantee common.Address, _startChainID *big.Int, _stopChainID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.GrantChainIDRange(&_Registry.TransactOpts, _grantee, _startChainID, _stopChainID)
}

// GrantChainIDRangeIrrevocable is a paid mutator transaction binding the contract method 0x09d24d68.
//
// Solidity: function grantChainIDRangeIrrevocable(address _grantee, uint256 _startChainID, uint256 _stopChainID) returns()
func (_Registry *RegistryTransactor) GrantChainIDRangeIrrevocable(opts *bind.TransactOpts, _grantee common.Address, _startChainID *big.Int, _stopChainID *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "grantChainIDRangeIrrevocable", _grantee, _startChainID, _stopChainID)
}

// GrantChainIDRangeIrrevocable is a paid mutator transaction binding the contract method 0x09d24d68.
//
// Solidity: function grantChainIDRangeIrrevocable(address _grantee, uint256 _startChainID, uint256 _stopChainID) returns()
func (_Registry *RegistrySession) GrantChainIDRangeIrrevocable(_grantee common.Address, _startChainID *big.Int, _stopChainID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.GrantChainIDRangeIrrevocable(&_Registry.TransactOpts, _grantee, _startChainID, _stopChainID)
}

// GrantChainIDRangeIrrevocable is a paid mutator transaction binding the contract method 0x09d24d68.
//
// Solidity: function grantChainIDRangeIrrevocable(address _grantee, uint256 _startChainID, uint256 _stopChainID) returns()
func (_Registry *RegistryTransactorSession) GrantChainIDRangeIrrevocable(_grantee common.Address, _startChainID *big.Int, _stopChainID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.GrantChainIDRangeIrrevocable(&_Registry.TransactOpts, _grantee, _startChainID, _stopChainID)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.GrantRole(&_Registry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.GrantRole(&_Registry.TransactOpts, role, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Registry *RegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RenounceRole(&_Registry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RenounceRole(&_Registry.TransactOpts, role, account)
}

// RevokeChainID is a paid mutator transaction binding the contract method 0x012ce089.
//
// Solidity: function revokeChainID(address _grantee, uint256 _chainID) returns()
func (_Registry *RegistryTransactor) RevokeChainID(opts *bind.TransactOpts, _grantee common.Address, _chainID *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "revokeChainID", _grantee, _chainID)
}

// RevokeChainID is a paid mutator transaction binding the contract method 0x012ce089.
//
// Solidity: function revokeChainID(address _grantee, uint256 _chainID) returns()
func (_Registry *RegistrySession) RevokeChainID(_grantee common.Address, _chainID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RevokeChainID(&_Registry.TransactOpts, _grantee, _chainID)
}

// RevokeChainID is a paid mutator transaction binding the contract method 0x012ce089.
//
// Solidity: function revokeChainID(address _grantee, uint256 _chainID) returns()
func (_Registry *RegistryTransactorSession) RevokeChainID(_grantee common.Address, _chainID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RevokeChainID(&_Registry.TransactOpts, _grantee, _chainID)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RevokeRole(&_Registry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RevokeRole(&_Registry.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// UpdateL1ChainConfiguration is a paid mutator transaction binding the contract method 0x82c614a9.
//
// Solidity: function updateL1ChainConfiguration(uint256 _chainID, (address,address,uint256,uint256) _config) returns()
func (_Registry *RegistryTransactor) UpdateL1ChainConfiguration(opts *bind.TransactOpts, _chainID *big.Int, _config L1Configuration) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateL1ChainConfiguration", _chainID, _config)
}

// UpdateL1ChainConfiguration is a paid mutator transaction binding the contract method 0x82c614a9.
//
// Solidity: function updateL1ChainConfiguration(uint256 _chainID, (address,address,uint256,uint256) _config) returns()
func (_Registry *RegistrySession) UpdateL1ChainConfiguration(_chainID *big.Int, _config L1Configuration) (*types.Transaction, error) {
	return _Registry.Contract.UpdateL1ChainConfiguration(&_Registry.TransactOpts, _chainID, _config)
}

// UpdateL1ChainConfiguration is a paid mutator transaction binding the contract method 0x82c614a9.
//
// Solidity: function updateL1ChainConfiguration(uint256 _chainID, (address,address,uint256,uint256) _config) returns()
func (_Registry *RegistryTransactorSession) UpdateL1ChainConfiguration(_chainID *big.Int, _config L1Configuration) (*types.Transaction, error) {
	return _Registry.Contract.UpdateL1ChainConfiguration(&_Registry.TransactOpts, _chainID, _config)
}

// UpdateL2ChainConfiguration is a paid mutator transaction binding the contract method 0x53467449.
//
// Solidity: function updateL2ChainConfiguration(uint256 _chainID, (address,address[],uint256[],uint256,uint256,uint8) _config) returns()
func (_Registry *RegistryTransactor) UpdateL2ChainConfiguration(opts *bind.TransactOpts, _chainID *big.Int, _config L2Configuration) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateL2ChainConfiguration", _chainID, _config)
}

// UpdateL2ChainConfiguration is a paid mutator transaction binding the contract method 0x53467449.
//
// Solidity: function updateL2ChainConfiguration(uint256 _chainID, (address,address[],uint256[],uint256,uint256,uint8) _config) returns()
func (_Registry *RegistrySession) UpdateL2ChainConfiguration(_chainID *big.Int, _config L2Configuration) (*types.Transaction, error) {
	return _Registry.Contract.UpdateL2ChainConfiguration(&_Registry.TransactOpts, _chainID, _config)
}

// UpdateL2ChainConfiguration is a paid mutator transaction binding the contract method 0x53467449.
//
// Solidity: function updateL2ChainConfiguration(uint256 _chainID, (address,address[],uint256[],uint256,uint256,uint8) _config) returns()
func (_Registry *RegistryTransactorSession) UpdateL2ChainConfiguration(_chainID *big.Int, _config L2Configuration) (*types.Transaction, error) {
	return _Registry.Contract.UpdateL2ChainConfiguration(&_Registry.TransactOpts, _chainID, _config)
}

// RegistryL1ChainConfigurationUpdatedIterator is returned from FilterL1ChainConfigurationUpdated and is used to iterate over the raw logs and unpacked data for L1ChainConfigurationUpdated events raised by the Registry contract.
type RegistryL1ChainConfigurationUpdatedIterator struct {
	Event *RegistryL1ChainConfigurationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryL1ChainConfigurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryL1ChainConfigurationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryL1ChainConfigurationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryL1ChainConfigurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryL1ChainConfigurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryL1ChainConfigurationUpdated represents a L1ChainConfigurationUpdated event raised by the Registry contract.
type RegistryL1ChainConfigurationUpdated struct {
	ChainID    *big.Int
	ConfigHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterL1ChainConfigurationUpdated is a free log retrieval operation binding the contract event 0xf2cbdfca96a2b93acf77796faef38fe901606474d1dc69fd99796b184b788a9a.
//
// Solidity: event L1ChainConfigurationUpdated(uint256 indexed chainID, bytes32 indexed configHash)
func (_Registry *RegistryFilterer) FilterL1ChainConfigurationUpdated(opts *bind.FilterOpts, chainID []*big.Int, configHash [][32]byte) (*RegistryL1ChainConfigurationUpdatedIterator, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var configHashRule []interface{}
	for _, configHashItem := range configHash {
		configHashRule = append(configHashRule, configHashItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "L1ChainConfigurationUpdated", chainIDRule, configHashRule)
	if err != nil {
		return nil, err
	}
	return &RegistryL1ChainConfigurationUpdatedIterator{contract: _Registry.contract, event: "L1ChainConfigurationUpdated", logs: logs, sub: sub}, nil
}

// WatchL1ChainConfigurationUpdated is a free log subscription operation binding the contract event 0xf2cbdfca96a2b93acf77796faef38fe901606474d1dc69fd99796b184b788a9a.
//
// Solidity: event L1ChainConfigurationUpdated(uint256 indexed chainID, bytes32 indexed configHash)
func (_Registry *RegistryFilterer) WatchL1ChainConfigurationUpdated(opts *bind.WatchOpts, sink chan<- *RegistryL1ChainConfigurationUpdated, chainID []*big.Int, configHash [][32]byte) (event.Subscription, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var configHashRule []interface{}
	for _, configHashItem := range configHash {
		configHashRule = append(configHashRule, configHashItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "L1ChainConfigurationUpdated", chainIDRule, configHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryL1ChainConfigurationUpdated)
				if err := _Registry.contract.UnpackLog(event, "L1ChainConfigurationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseL1ChainConfigurationUpdated is a log parse operation binding the contract event 0xf2cbdfca96a2b93acf77796faef38fe901606474d1dc69fd99796b184b788a9a.
//
// Solidity: event L1ChainConfigurationUpdated(uint256 indexed chainID, bytes32 indexed configHash)
func (_Registry *RegistryFilterer) ParseL1ChainConfigurationUpdated(log types.Log) (*RegistryL1ChainConfigurationUpdated, error) {
	event := new(RegistryL1ChainConfigurationUpdated)
	if err := _Registry.contract.UnpackLog(event, "L1ChainConfigurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryL2ChainConfigurationUpdatedIterator is returned from FilterL2ChainConfigurationUpdated and is used to iterate over the raw logs and unpacked data for L2ChainConfigurationUpdated events raised by the Registry contract.
type RegistryL2ChainConfigurationUpdatedIterator struct {
	Event *RegistryL2ChainConfigurationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryL2ChainConfigurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryL2ChainConfigurationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryL2ChainConfigurationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryL2ChainConfigurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryL2ChainConfigurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryL2ChainConfigurationUpdated represents a L2ChainConfigurationUpdated event raised by the Registry contract.
type RegistryL2ChainConfigurationUpdated struct {
	ChainID    *big.Int
	ConfigHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterL2ChainConfigurationUpdated is a free log retrieval operation binding the contract event 0x5c38bb2a34850964e91618d58976c032cec224be07254f31a9a1c95f24f68c0d.
//
// Solidity: event L2ChainConfigurationUpdated(uint256 indexed chainID, bytes32 indexed configHash)
func (_Registry *RegistryFilterer) FilterL2ChainConfigurationUpdated(opts *bind.FilterOpts, chainID []*big.Int, configHash [][32]byte) (*RegistryL2ChainConfigurationUpdatedIterator, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var configHashRule []interface{}
	for _, configHashItem := range configHash {
		configHashRule = append(configHashRule, configHashItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "L2ChainConfigurationUpdated", chainIDRule, configHashRule)
	if err != nil {
		return nil, err
	}
	return &RegistryL2ChainConfigurationUpdatedIterator{contract: _Registry.contract, event: "L2ChainConfigurationUpdated", logs: logs, sub: sub}, nil
}

// WatchL2ChainConfigurationUpdated is a free log subscription operation binding the contract event 0x5c38bb2a34850964e91618d58976c032cec224be07254f31a9a1c95f24f68c0d.
//
// Solidity: event L2ChainConfigurationUpdated(uint256 indexed chainID, bytes32 indexed configHash)
func (_Registry *RegistryFilterer) WatchL2ChainConfigurationUpdated(opts *bind.WatchOpts, sink chan<- *RegistryL2ChainConfigurationUpdated, chainID []*big.Int, configHash [][32]byte) (event.Subscription, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var configHashRule []interface{}
	for _, configHashItem := range configHash {
		configHashRule = append(configHashRule, configHashItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "L2ChainConfigurationUpdated", chainIDRule, configHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryL2ChainConfigurationUpdated)
				if err := _Registry.contract.UnpackLog(event, "L2ChainConfigurationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseL2ChainConfigurationUpdated is a log parse operation binding the contract event 0x5c38bb2a34850964e91618d58976c032cec224be07254f31a9a1c95f24f68c0d.
//
// Solidity: event L2ChainConfigurationUpdated(uint256 indexed chainID, bytes32 indexed configHash)
func (_Registry *RegistryFilterer) ParseL2ChainConfigurationUpdated(log types.Log) (*RegistryL2ChainConfigurationUpdated, error) {
	event := new(RegistryL2ChainConfigurationUpdated)
	if err := _Registry.contract.UnpackLog(event, "L2ChainConfigurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryOwnershipTransferred, error) {
	event := new(RegistryOwnershipTransferred)
	if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Registry contract.
type RegistryPausedIterator struct {
	Event *RegistryPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryPaused represents a Paused event raised by the Registry contract.
type RegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registry *RegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*RegistryPausedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RegistryPausedIterator{contract: _Registry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registry *RegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RegistryPaused) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryPaused)
				if err := _Registry.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registry *RegistryFilterer) ParsePaused(log types.Log) (*RegistryPaused, error) {
	event := new(RegistryPaused)
	if err := _Registry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Registry contract.
type RegistryRoleAdminChangedIterator struct {
	Event *RegistryRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleAdminChanged represents a RoleAdminChanged event raised by the Registry contract.
type RegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Registry *RegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleAdminChangedIterator{contract: _Registry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Registry *RegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleAdminChanged)
				if err := _Registry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Registry *RegistryFilterer) ParseRoleAdminChanged(log types.Log) (*RegistryRoleAdminChanged, error) {
	event := new(RegistryRoleAdminChanged)
	if err := _Registry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Registry contract.
type RegistryRoleGrantedIterator struct {
	Event *RegistryRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleGranted represents a RoleGranted event raised by the Registry contract.
type RegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleGrantedIterator{contract: _Registry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleGranted)
				if err := _Registry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) ParseRoleGranted(log types.Log) (*RegistryRoleGranted, error) {
	event := new(RegistryRoleGranted)
	if err := _Registry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Registry contract.
type RegistryRoleRevokedIterator struct {
	Event *RegistryRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleRevoked represents a RoleRevoked event raised by the Registry contract.
type RegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleRevokedIterator{contract: _Registry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleRevoked)
				if err := _Registry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) ParseRoleRevoked(log types.Log) (*RegistryRoleRevoked, error) {
	event := new(RegistryRoleRevoked)
	if err := _Registry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Registry contract.
type RegistryUnpausedIterator struct {
	Event *RegistryUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryUnpaused represents a Unpaused event raised by the Registry contract.
type RegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registry *RegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RegistryUnpausedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RegistryUnpausedIterator{contract: _Registry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registry *RegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryUnpaused)
				if err := _Registry.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registry *RegistryFilterer) ParseUnpaused(log types.Log) (*RegistryUnpaused, error) {
	event := new(RegistryUnpaused)
	if err := _Registry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
