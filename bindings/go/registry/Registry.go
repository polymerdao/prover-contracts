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
	SettlementBlocksDelay                 *big.Int
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_multiSigOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_onlyAdd\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_initialL2Configurations\",\"type\":\"tuple[]\",\"internalType\":\"structRegistry.InitialL2Configuration[]\",\"components\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"_initialL1Configurations\",\"type\":\"tuple[]\",\"internalType\":\"structRegistry.InitialL1Configuration[]\",\"components\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structL1Configuration\",\"components\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementBlocksDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getL2ConfigAddresses\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getL2ConfigStorageSlots\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantChainID\",\"inputs\":[{\"name\":\"_grantee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantChainIDIrrevocable\",\"inputs\":[{\"name\":\"_grantee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantChainIDRange\",\"inputs\":[{\"name\":\"_grantee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_startChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stopChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantChainIDRangeIrrevocable\",\"inputs\":[{\"name\":\"_grantee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_startChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stopChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"granteeBitmap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"irrevocableChainIDBitmap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isGrantee\",\"inputs\":[{\"name\":\"_grantee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1ChainConfigurationHashMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1ChainConfigurations\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementBlocksDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2ChainConfigurationHashMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2ChainConfigurations\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"multiSigOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onlyAdd\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateL1ChainConfiguration\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structL1Configuration\",\"components\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementBlocksDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateL2ChainConfiguration\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"L1ChainConfigurationUpdated\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"configHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L2ChainConfigurationUpdated\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"configHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewGrantee\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"grantee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewIrrevocableGrantee\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"grantee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
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

// GranteeBitmap is a free data retrieval call binding the contract method 0x50257fb1.
//
// Solidity: function granteeBitmap(address , uint256 ) view returns(uint256)
func (_Registry *RegistryCaller) GranteeBitmap(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "granteeBitmap", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GranteeBitmap is a free data retrieval call binding the contract method 0x50257fb1.
//
// Solidity: function granteeBitmap(address , uint256 ) view returns(uint256)
func (_Registry *RegistrySession) GranteeBitmap(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Registry.Contract.GranteeBitmap(&_Registry.CallOpts, arg0, arg1)
}

// GranteeBitmap is a free data retrieval call binding the contract method 0x50257fb1.
//
// Solidity: function granteeBitmap(address , uint256 ) view returns(uint256)
func (_Registry *RegistryCallerSession) GranteeBitmap(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Registry.Contract.GranteeBitmap(&_Registry.CallOpts, arg0, arg1)
}

// IrrevocableChainIDBitmap is a free data retrieval call binding the contract method 0xea4d2585.
//
// Solidity: function irrevocableChainIDBitmap(uint256 ) view returns(uint256)
func (_Registry *RegistryCaller) IrrevocableChainIDBitmap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "irrevocableChainIDBitmap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IrrevocableChainIDBitmap is a free data retrieval call binding the contract method 0xea4d2585.
//
// Solidity: function irrevocableChainIDBitmap(uint256 ) view returns(uint256)
func (_Registry *RegistrySession) IrrevocableChainIDBitmap(arg0 *big.Int) (*big.Int, error) {
	return _Registry.Contract.IrrevocableChainIDBitmap(&_Registry.CallOpts, arg0)
}

// IrrevocableChainIDBitmap is a free data retrieval call binding the contract method 0xea4d2585.
//
// Solidity: function irrevocableChainIDBitmap(uint256 ) view returns(uint256)
func (_Registry *RegistryCallerSession) IrrevocableChainIDBitmap(arg0 *big.Int) (*big.Int, error) {
	return _Registry.Contract.IrrevocableChainIDBitmap(&_Registry.CallOpts, arg0)
}

// IsGrantee is a free data retrieval call binding the contract method 0xe465a320.
//
// Solidity: function isGrantee(address _grantee, uint256 _chainID) view returns(bool)
func (_Registry *RegistryCaller) IsGrantee(opts *bind.CallOpts, _grantee common.Address, _chainID *big.Int) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isGrantee", _grantee, _chainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGrantee is a free data retrieval call binding the contract method 0xe465a320.
//
// Solidity: function isGrantee(address _grantee, uint256 _chainID) view returns(bool)
func (_Registry *RegistrySession) IsGrantee(_grantee common.Address, _chainID *big.Int) (bool, error) {
	return _Registry.Contract.IsGrantee(&_Registry.CallOpts, _grantee, _chainID)
}

// IsGrantee is a free data retrieval call binding the contract method 0xe465a320.
//
// Solidity: function isGrantee(address _grantee, uint256 _chainID) view returns(bool)
func (_Registry *RegistryCallerSession) IsGrantee(_grantee common.Address, _chainID *big.Int) (bool, error) {
	return _Registry.Contract.IsGrantee(&_Registry.CallOpts, _grantee, _chainID)
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
// Solidity: function l1ChainConfigurations(uint256 ) view returns(address blockHashOracle, uint256 settlementBlocksDelay, address settlementRegistry, uint256 settlementRegistryL2ConfigMappingSlot, uint256 settlementRegistryL1ConfigMappingSlot)
func (_Registry *RegistryCaller) L1ChainConfigurations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockHashOracle                       common.Address
	SettlementBlocksDelay                 *big.Int
	SettlementRegistry                    common.Address
	SettlementRegistryL2ConfigMappingSlot *big.Int
	SettlementRegistryL1ConfigMappingSlot *big.Int
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "l1ChainConfigurations", arg0)

	outstruct := new(struct {
		BlockHashOracle                       common.Address
		SettlementBlocksDelay                 *big.Int
		SettlementRegistry                    common.Address
		SettlementRegistryL2ConfigMappingSlot *big.Int
		SettlementRegistryL1ConfigMappingSlot *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockHashOracle = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SettlementBlocksDelay = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SettlementRegistry = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.SettlementRegistryL2ConfigMappingSlot = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SettlementRegistryL1ConfigMappingSlot = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// L1ChainConfigurations is a free data retrieval call binding the contract method 0xa9a9730a.
//
// Solidity: function l1ChainConfigurations(uint256 ) view returns(address blockHashOracle, uint256 settlementBlocksDelay, address settlementRegistry, uint256 settlementRegistryL2ConfigMappingSlot, uint256 settlementRegistryL1ConfigMappingSlot)
func (_Registry *RegistrySession) L1ChainConfigurations(arg0 *big.Int) (struct {
	BlockHashOracle                       common.Address
	SettlementBlocksDelay                 *big.Int
	SettlementRegistry                    common.Address
	SettlementRegistryL2ConfigMappingSlot *big.Int
	SettlementRegistryL1ConfigMappingSlot *big.Int
}, error) {
	return _Registry.Contract.L1ChainConfigurations(&_Registry.CallOpts, arg0)
}

// L1ChainConfigurations is a free data retrieval call binding the contract method 0xa9a9730a.
//
// Solidity: function l1ChainConfigurations(uint256 ) view returns(address blockHashOracle, uint256 settlementBlocksDelay, address settlementRegistry, uint256 settlementRegistryL2ConfigMappingSlot, uint256 settlementRegistryL1ConfigMappingSlot)
func (_Registry *RegistryCallerSession) L1ChainConfigurations(arg0 *big.Int) (struct {
	BlockHashOracle                       common.Address
	SettlementBlocksDelay                 *big.Int
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
// Solidity: function l2ChainConfigurations(uint256 ) view returns(address prover, uint256 versionNumber, uint256 finalityDelaySeconds)
func (_Registry *RegistryCaller) L2ChainConfigurations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Prover               common.Address
	VersionNumber        *big.Int
	FinalityDelaySeconds *big.Int
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "l2ChainConfigurations", arg0)

	outstruct := new(struct {
		Prover               common.Address
		VersionNumber        *big.Int
		FinalityDelaySeconds *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Prover = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.VersionNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FinalityDelaySeconds = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// L2ChainConfigurations is a free data retrieval call binding the contract method 0x63281a15.
//
// Solidity: function l2ChainConfigurations(uint256 ) view returns(address prover, uint256 versionNumber, uint256 finalityDelaySeconds)
func (_Registry *RegistrySession) L2ChainConfigurations(arg0 *big.Int) (struct {
	Prover               common.Address
	VersionNumber        *big.Int
	FinalityDelaySeconds *big.Int
}, error) {
	return _Registry.Contract.L2ChainConfigurations(&_Registry.CallOpts, arg0)
}

// L2ChainConfigurations is a free data retrieval call binding the contract method 0x63281a15.
//
// Solidity: function l2ChainConfigurations(uint256 ) view returns(address prover, uint256 versionNumber, uint256 finalityDelaySeconds)
func (_Registry *RegistryCallerSession) L2ChainConfigurations(arg0 *big.Int) (struct {
	Prover               common.Address
	VersionNumber        *big.Int
	FinalityDelaySeconds *big.Int
}, error) {
	return _Registry.Contract.L2ChainConfigurations(&_Registry.CallOpts, arg0)
}

// MultiSigOwner is a free data retrieval call binding the contract method 0x0329dd62.
//
// Solidity: function multiSigOwner() view returns(address)
func (_Registry *RegistryCaller) MultiSigOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "multiSigOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MultiSigOwner is a free data retrieval call binding the contract method 0x0329dd62.
//
// Solidity: function multiSigOwner() view returns(address)
func (_Registry *RegistrySession) MultiSigOwner() (common.Address, error) {
	return _Registry.Contract.MultiSigOwner(&_Registry.CallOpts)
}

// MultiSigOwner is a free data retrieval call binding the contract method 0x0329dd62.
//
// Solidity: function multiSigOwner() view returns(address)
func (_Registry *RegistryCallerSession) MultiSigOwner() (common.Address, error) {
	return _Registry.Contract.MultiSigOwner(&_Registry.CallOpts)
}

// OnlyAdd is a free data retrieval call binding the contract method 0x00d67e79.
//
// Solidity: function onlyAdd() view returns(bool)
func (_Registry *RegistryCaller) OnlyAdd(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "onlyAdd")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlyAdd is a free data retrieval call binding the contract method 0x00d67e79.
//
// Solidity: function onlyAdd() view returns(bool)
func (_Registry *RegistrySession) OnlyAdd() (bool, error) {
	return _Registry.Contract.OnlyAdd(&_Registry.CallOpts)
}

// OnlyAdd is a free data retrieval call binding the contract method 0x00d67e79.
//
// Solidity: function onlyAdd() view returns(bool)
func (_Registry *RegistryCallerSession) OnlyAdd() (bool, error) {
	return _Registry.Contract.OnlyAdd(&_Registry.CallOpts)
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

// UpdateL1ChainConfiguration is a paid mutator transaction binding the contract method 0xa6a24191.
//
// Solidity: function updateL1ChainConfiguration(uint256 _chainID, (address,uint256,address,uint256,uint256) _config) returns()
func (_Registry *RegistryTransactor) UpdateL1ChainConfiguration(opts *bind.TransactOpts, _chainID *big.Int, _config L1Configuration) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateL1ChainConfiguration", _chainID, _config)
}

// UpdateL1ChainConfiguration is a paid mutator transaction binding the contract method 0xa6a24191.
//
// Solidity: function updateL1ChainConfiguration(uint256 _chainID, (address,uint256,address,uint256,uint256) _config) returns()
func (_Registry *RegistrySession) UpdateL1ChainConfiguration(_chainID *big.Int, _config L1Configuration) (*types.Transaction, error) {
	return _Registry.Contract.UpdateL1ChainConfiguration(&_Registry.TransactOpts, _chainID, _config)
}

// UpdateL1ChainConfiguration is a paid mutator transaction binding the contract method 0xa6a24191.
//
// Solidity: function updateL1ChainConfiguration(uint256 _chainID, (address,uint256,address,uint256,uint256) _config) returns()
func (_Registry *RegistryTransactorSession) UpdateL1ChainConfiguration(_chainID *big.Int, _config L1Configuration) (*types.Transaction, error) {
	return _Registry.Contract.UpdateL1ChainConfiguration(&_Registry.TransactOpts, _chainID, _config)
}

// UpdateL2ChainConfiguration is a paid mutator transaction binding the contract method 0x3b6b4975.
//
// Solidity: function updateL2ChainConfiguration(uint256 _chainID, (address,address[],uint256[],uint256,uint256) _config) returns()
func (_Registry *RegistryTransactor) UpdateL2ChainConfiguration(opts *bind.TransactOpts, _chainID *big.Int, _config L2Configuration) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateL2ChainConfiguration", _chainID, _config)
}

// UpdateL2ChainConfiguration is a paid mutator transaction binding the contract method 0x3b6b4975.
//
// Solidity: function updateL2ChainConfiguration(uint256 _chainID, (address,address[],uint256[],uint256,uint256) _config) returns()
func (_Registry *RegistrySession) UpdateL2ChainConfiguration(_chainID *big.Int, _config L2Configuration) (*types.Transaction, error) {
	return _Registry.Contract.UpdateL2ChainConfiguration(&_Registry.TransactOpts, _chainID, _config)
}

// UpdateL2ChainConfiguration is a paid mutator transaction binding the contract method 0x3b6b4975.
//
// Solidity: function updateL2ChainConfiguration(uint256 _chainID, (address,address[],uint256[],uint256,uint256) _config) returns()
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

// RegistryNewGranteeIterator is returned from FilterNewGrantee and is used to iterate over the raw logs and unpacked data for NewGrantee events raised by the Registry contract.
type RegistryNewGranteeIterator struct {
	Event *RegistryNewGrantee // Event containing the contract specifics and raw log

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
func (it *RegistryNewGranteeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryNewGrantee)
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
		it.Event = new(RegistryNewGrantee)
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
func (it *RegistryNewGranteeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryNewGranteeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryNewGrantee represents a NewGrantee event raised by the Registry contract.
type RegistryNewGrantee struct {
	ChainID *big.Int
	Grantee common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewGrantee is a free log retrieval operation binding the contract event 0x344ff5c631f29d0320d431db3da3961d2fc92a536f4f08c488792583d4a97971.
//
// Solidity: event NewGrantee(uint256 indexed chainID, address indexed grantee)
func (_Registry *RegistryFilterer) FilterNewGrantee(opts *bind.FilterOpts, chainID []*big.Int, grantee []common.Address) (*RegistryNewGranteeIterator, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "NewGrantee", chainIDRule, granteeRule)
	if err != nil {
		return nil, err
	}
	return &RegistryNewGranteeIterator{contract: _Registry.contract, event: "NewGrantee", logs: logs, sub: sub}, nil
}

// WatchNewGrantee is a free log subscription operation binding the contract event 0x344ff5c631f29d0320d431db3da3961d2fc92a536f4f08c488792583d4a97971.
//
// Solidity: event NewGrantee(uint256 indexed chainID, address indexed grantee)
func (_Registry *RegistryFilterer) WatchNewGrantee(opts *bind.WatchOpts, sink chan<- *RegistryNewGrantee, chainID []*big.Int, grantee []common.Address) (event.Subscription, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "NewGrantee", chainIDRule, granteeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryNewGrantee)
				if err := _Registry.contract.UnpackLog(event, "NewGrantee", log); err != nil {
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

// ParseNewGrantee is a log parse operation binding the contract event 0x344ff5c631f29d0320d431db3da3961d2fc92a536f4f08c488792583d4a97971.
//
// Solidity: event NewGrantee(uint256 indexed chainID, address indexed grantee)
func (_Registry *RegistryFilterer) ParseNewGrantee(log types.Log) (*RegistryNewGrantee, error) {
	event := new(RegistryNewGrantee)
	if err := _Registry.contract.UnpackLog(event, "NewGrantee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryNewIrrevocableGranteeIterator is returned from FilterNewIrrevocableGrantee and is used to iterate over the raw logs and unpacked data for NewIrrevocableGrantee events raised by the Registry contract.
type RegistryNewIrrevocableGranteeIterator struct {
	Event *RegistryNewIrrevocableGrantee // Event containing the contract specifics and raw log

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
func (it *RegistryNewIrrevocableGranteeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryNewIrrevocableGrantee)
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
		it.Event = new(RegistryNewIrrevocableGrantee)
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
func (it *RegistryNewIrrevocableGranteeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryNewIrrevocableGranteeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryNewIrrevocableGrantee represents a NewIrrevocableGrantee event raised by the Registry contract.
type RegistryNewIrrevocableGrantee struct {
	ChainID *big.Int
	Grantee common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewIrrevocableGrantee is a free log retrieval operation binding the contract event 0xa4285119a5eeb9e389ae4510c20caae9fea8065eb4ea177201f20bfba3713b03.
//
// Solidity: event NewIrrevocableGrantee(uint256 indexed chainID, address indexed grantee)
func (_Registry *RegistryFilterer) FilterNewIrrevocableGrantee(opts *bind.FilterOpts, chainID []*big.Int, grantee []common.Address) (*RegistryNewIrrevocableGranteeIterator, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "NewIrrevocableGrantee", chainIDRule, granteeRule)
	if err != nil {
		return nil, err
	}
	return &RegistryNewIrrevocableGranteeIterator{contract: _Registry.contract, event: "NewIrrevocableGrantee", logs: logs, sub: sub}, nil
}

// WatchNewIrrevocableGrantee is a free log subscription operation binding the contract event 0xa4285119a5eeb9e389ae4510c20caae9fea8065eb4ea177201f20bfba3713b03.
//
// Solidity: event NewIrrevocableGrantee(uint256 indexed chainID, address indexed grantee)
func (_Registry *RegistryFilterer) WatchNewIrrevocableGrantee(opts *bind.WatchOpts, sink chan<- *RegistryNewIrrevocableGrantee, chainID []*big.Int, grantee []common.Address) (event.Subscription, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "NewIrrevocableGrantee", chainIDRule, granteeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryNewIrrevocableGrantee)
				if err := _Registry.contract.UnpackLog(event, "NewIrrevocableGrantee", log); err != nil {
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

// ParseNewIrrevocableGrantee is a log parse operation binding the contract event 0xa4285119a5eeb9e389ae4510c20caae9fea8065eb4ea177201f20bfba3713b03.
//
// Solidity: event NewIrrevocableGrantee(uint256 indexed chainID, address indexed grantee)
func (_Registry *RegistryFilterer) ParseNewIrrevocableGrantee(log types.Log) (*RegistryNewIrrevocableGrantee, error) {
	event := new(RegistryNewIrrevocableGrantee)
	if err := _Registry.contract.UnpackLog(event, "NewIrrevocableGrantee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
