// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativeprover

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

// ProveL1ScalarArgs is an auto generated low-level Go binding around an user-defined struct.
type ProveL1ScalarArgs struct {
	ContractAddr     common.Address
	StorageSlot      [32]byte
	StorageValue     [32]byte
	L1WorldStateRoot [32]byte
}

// ProveScalarArgs is an auto generated low-level Go binding around an user-defined struct.
type ProveScalarArgs struct {
	ChainID          *big.Int
	ContractAddr     common.Address
	StorageSlot      [32]byte
	StorageValue     [32]byte
	L2WorldStateRoot [32]byte
}

// UpdateL2ConfigArgs is an auto generated low-level Go binding around an user-defined struct.
type UpdateL2ConfigArgs struct {
	Config                        L2Configuration
	L1StorageProof                [][]byte
	RlpEncodedRegistryAccountData []byte
	L1RegistryProof               [][]byte
}

// NativeProverMetaData contains all meta data concerning the NativeProver contract.
var NativeProverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1ChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"L1_CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"L1_CONFIGURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveL1Native\",\"inputs\":[{\"name\":\"_args\",\"type\":\"tuple\",\"internalType\":\"structProveL1ScalarArgs\",\"components\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l1WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_rlpEncodedL1Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l1StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_rlpEncodedContractAccount\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l1AccountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveL2Native\",\"inputs\":[{\"name\":\"_updateArgs\",\"type\":\"tuple\",\"internalType\":\"structUpdateL2ConfigArgs\",\"components\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"enumType\"}]},{\"name\":\"l1StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"rlpEncodedRegistryAccountData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"l1RegistryProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"name\":\"_proveArgs\",\"type\":\"tuple\",\"internalType\":\"structProveScalarArgs\",\"components\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_rlpEncodedL1Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_rlpEncodedL2Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_settledStateProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l2StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_rlpEncodedContractAccount\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l2AccountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInitialL1Config\",\"inputs\":[{\"name\":\"_l1Configuration\",\"type\":\"tuple\",\"internalType\":\"structL1Configuration\",\"components\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateL1ChainConfiguration\",\"inputs\":[{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structL1Configuration\",\"components\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_l1StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_rlpEncodedRegistryAccountData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l1RegistryProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_rlpEncodedL1Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"IncorrectContractStorageRoot\",\"inputs\":[{\"name\":\"_contractStorageRoot\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountProof\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidL1ConfigurationProof\",\"inputs\":[{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structL1Configuration\",\"components\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"type\":\"error\",\"name\":\"InvalidL2ConfigurationProof\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"enumType\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedBlock\",\"inputs\":[{\"name\":\"_expectedBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_calculatedBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidSettledStateProof\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidStorageProof\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_val\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SettlementChainStateRootNotProven\",\"inputs\":[{\"name\":\"_blockProofStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_l1WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// NativeProverABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeProverMetaData.ABI instead.
var NativeProverABI = NativeProverMetaData.ABI

// NativeProver is an auto generated Go binding around an Ethereum contract.
type NativeProver struct {
	NativeProverCaller     // Read-only binding to the contract
	NativeProverTransactor // Write-only binding to the contract
	NativeProverFilterer   // Log filterer for contract events
}

// NativeProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeProverSession struct {
	Contract     *NativeProver     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeProverCallerSession struct {
	Contract *NativeProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NativeProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeProverTransactorSession struct {
	Contract     *NativeProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NativeProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeProverRaw struct {
	Contract *NativeProver // Generic contract binding to access the raw methods on
}

// NativeProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeProverCallerRaw struct {
	Contract *NativeProverCaller // Generic read-only contract binding to access the raw methods on
}

// NativeProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeProverTransactorRaw struct {
	Contract *NativeProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeProver creates a new instance of NativeProver, bound to a specific deployed contract.
func NewNativeProver(address common.Address, backend bind.ContractBackend) (*NativeProver, error) {
	contract, err := bindNativeProver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeProver{NativeProverCaller: NativeProverCaller{contract: contract}, NativeProverTransactor: NativeProverTransactor{contract: contract}, NativeProverFilterer: NativeProverFilterer{contract: contract}}, nil
}

// NewNativeProverCaller creates a new read-only instance of NativeProver, bound to a specific deployed contract.
func NewNativeProverCaller(address common.Address, caller bind.ContractCaller) (*NativeProverCaller, error) {
	contract, err := bindNativeProver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeProverCaller{contract: contract}, nil
}

// NewNativeProverTransactor creates a new write-only instance of NativeProver, bound to a specific deployed contract.
func NewNativeProverTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeProverTransactor, error) {
	contract, err := bindNativeProver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeProverTransactor{contract: contract}, nil
}

// NewNativeProverFilterer creates a new log filterer instance of NativeProver, bound to a specific deployed contract.
func NewNativeProverFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeProverFilterer, error) {
	contract, err := bindNativeProver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeProverFilterer{contract: contract}, nil
}

// bindNativeProver binds a generic wrapper to an already deployed contract.
func bindNativeProver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeProverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeProver *NativeProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeProver.Contract.NativeProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeProver *NativeProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeProver.Contract.NativeProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeProver *NativeProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeProver.Contract.NativeProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeProver *NativeProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeProver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeProver *NativeProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeProver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeProver *NativeProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeProver.Contract.contract.Transact(opts, method, params...)
}

// L1CHAINID is a free data retrieval call binding the contract method 0x2f90b184.
//
// Solidity: function L1_CHAIN_ID() view returns(uint256)
func (_NativeProver *NativeProverCaller) L1CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "L1_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1CHAINID is a free data retrieval call binding the contract method 0x2f90b184.
//
// Solidity: function L1_CHAIN_ID() view returns(uint256)
func (_NativeProver *NativeProverSession) L1CHAINID() (*big.Int, error) {
	return _NativeProver.Contract.L1CHAINID(&_NativeProver.CallOpts)
}

// L1CHAINID is a free data retrieval call binding the contract method 0x2f90b184.
//
// Solidity: function L1_CHAIN_ID() view returns(uint256)
func (_NativeProver *NativeProverCallerSession) L1CHAINID() (*big.Int, error) {
	return _NativeProver.Contract.L1CHAINID(&_NativeProver.CallOpts)
}

// L1CONFIGURATION is a free data retrieval call binding the contract method 0xbff89934.
//
// Solidity: function L1_CONFIGURATION() view returns(address blockHashOracle, address settlementRegistry, uint256 settlementRegistryL2ConfigMappingSlot, uint256 settlementRegistryL1ConfigMappingSlot)
func (_NativeProver *NativeProverCaller) L1CONFIGURATION(opts *bind.CallOpts) (struct {
	BlockHashOracle                       common.Address
	SettlementRegistry                    common.Address
	SettlementRegistryL2ConfigMappingSlot *big.Int
	SettlementRegistryL1ConfigMappingSlot *big.Int
}, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "L1_CONFIGURATION")

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

// L1CONFIGURATION is a free data retrieval call binding the contract method 0xbff89934.
//
// Solidity: function L1_CONFIGURATION() view returns(address blockHashOracle, address settlementRegistry, uint256 settlementRegistryL2ConfigMappingSlot, uint256 settlementRegistryL1ConfigMappingSlot)
func (_NativeProver *NativeProverSession) L1CONFIGURATION() (struct {
	BlockHashOracle                       common.Address
	SettlementRegistry                    common.Address
	SettlementRegistryL2ConfigMappingSlot *big.Int
	SettlementRegistryL1ConfigMappingSlot *big.Int
}, error) {
	return _NativeProver.Contract.L1CONFIGURATION(&_NativeProver.CallOpts)
}

// L1CONFIGURATION is a free data retrieval call binding the contract method 0xbff89934.
//
// Solidity: function L1_CONFIGURATION() view returns(address blockHashOracle, address settlementRegistry, uint256 settlementRegistryL2ConfigMappingSlot, uint256 settlementRegistryL1ConfigMappingSlot)
func (_NativeProver *NativeProverCallerSession) L1CONFIGURATION() (struct {
	BlockHashOracle                       common.Address
	SettlementRegistry                    common.Address
	SettlementRegistryL2ConfigMappingSlot *big.Int
	SettlementRegistryL1ConfigMappingSlot *big.Int
}, error) {
	return _NativeProver.Contract.L1CONFIGURATION(&_NativeProver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeProver *NativeProverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeProver *NativeProverSession) Owner() (common.Address, error) {
	return _NativeProver.Contract.Owner(&_NativeProver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeProver *NativeProverCallerSession) Owner() (common.Address, error) {
	return _NativeProver.Contract.Owner(&_NativeProver.CallOpts)
}

// ProveL1Native is a free data retrieval call binding the contract method 0x7675c451.
//
// Solidity: function proveL1Native((address,bytes32,bytes32,bytes32) _args, bytes _rlpEncodedL1Header, bytes[] _l1StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l1AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverCaller) ProveL1Native(opts *bind.CallOpts, _args ProveL1ScalarArgs, _rlpEncodedL1Header []byte, _l1StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l1AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "proveL1Native", _args, _rlpEncodedL1Header, _l1StorageProof, _rlpEncodedContractAccount, _l1AccountProof)

	outstruct := new(struct {
		ChainID         *big.Int
		StoringContract common.Address
		StorageSlot     [32]byte
		StorageValue    [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StoringContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StorageSlot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StorageValue = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// ProveL1Native is a free data retrieval call binding the contract method 0x7675c451.
//
// Solidity: function proveL1Native((address,bytes32,bytes32,bytes32) _args, bytes _rlpEncodedL1Header, bytes[] _l1StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l1AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverSession) ProveL1Native(_args ProveL1ScalarArgs, _rlpEncodedL1Header []byte, _l1StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l1AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	return _NativeProver.Contract.ProveL1Native(&_NativeProver.CallOpts, _args, _rlpEncodedL1Header, _l1StorageProof, _rlpEncodedContractAccount, _l1AccountProof)
}

// ProveL1Native is a free data retrieval call binding the contract method 0x7675c451.
//
// Solidity: function proveL1Native((address,bytes32,bytes32,bytes32) _args, bytes _rlpEncodedL1Header, bytes[] _l1StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l1AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverCallerSession) ProveL1Native(_args ProveL1ScalarArgs, _rlpEncodedL1Header []byte, _l1StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l1AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	return _NativeProver.Contract.ProveL1Native(&_NativeProver.CallOpts, _args, _rlpEncodedL1Header, _l1StorageProof, _rlpEncodedContractAccount, _l1AccountProof)
}

// ProveL2Native is a free data retrieval call binding the contract method 0x629d8c96.
//
// Solidity: function proveL2Native(((address,address[],uint256[],uint256,uint256,uint8),bytes[],bytes,bytes[]) _updateArgs, (uint256,address,bytes32,bytes32,bytes32) _proveArgs, bytes _rlpEncodedL1Header, bytes _rlpEncodedL2Header, bytes _settledStateProof, bytes[] _l2StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l2AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverCaller) ProveL2Native(opts *bind.CallOpts, _updateArgs UpdateL2ConfigArgs, _proveArgs ProveScalarArgs, _rlpEncodedL1Header []byte, _rlpEncodedL2Header []byte, _settledStateProof []byte, _l2StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l2AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "proveL2Native", _updateArgs, _proveArgs, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof)

	outstruct := new(struct {
		ChainID         *big.Int
		StoringContract common.Address
		StorageSlot     [32]byte
		StorageValue    [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StoringContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StorageSlot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StorageValue = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// ProveL2Native is a free data retrieval call binding the contract method 0x629d8c96.
//
// Solidity: function proveL2Native(((address,address[],uint256[],uint256,uint256,uint8),bytes[],bytes,bytes[]) _updateArgs, (uint256,address,bytes32,bytes32,bytes32) _proveArgs, bytes _rlpEncodedL1Header, bytes _rlpEncodedL2Header, bytes _settledStateProof, bytes[] _l2StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l2AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverSession) ProveL2Native(_updateArgs UpdateL2ConfigArgs, _proveArgs ProveScalarArgs, _rlpEncodedL1Header []byte, _rlpEncodedL2Header []byte, _settledStateProof []byte, _l2StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l2AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	return _NativeProver.Contract.ProveL2Native(&_NativeProver.CallOpts, _updateArgs, _proveArgs, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof)
}

// ProveL2Native is a free data retrieval call binding the contract method 0x629d8c96.
//
// Solidity: function proveL2Native(((address,address[],uint256[],uint256,uint256,uint8),bytes[],bytes,bytes[]) _updateArgs, (uint256,address,bytes32,bytes32,bytes32) _proveArgs, bytes _rlpEncodedL1Header, bytes _rlpEncodedL2Header, bytes _settledStateProof, bytes[] _l2StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l2AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverCallerSession) ProveL2Native(_updateArgs UpdateL2ConfigArgs, _proveArgs ProveScalarArgs, _rlpEncodedL1Header []byte, _rlpEncodedL2Header []byte, _settledStateProof []byte, _l2StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l2AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	return _NativeProver.Contract.ProveL2Native(&_NativeProver.CallOpts, _updateArgs, _proveArgs, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeProver *NativeProverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeProver.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeProver *NativeProverSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeProver.Contract.RenounceOwnership(&_NativeProver.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeProver *NativeProverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeProver.Contract.RenounceOwnership(&_NativeProver.TransactOpts)
}

// SetInitialL1Config is a paid mutator transaction binding the contract method 0xd1156991.
//
// Solidity: function setInitialL1Config((address,address,uint256,uint256) _l1Configuration) returns()
func (_NativeProver *NativeProverTransactor) SetInitialL1Config(opts *bind.TransactOpts, _l1Configuration L1Configuration) (*types.Transaction, error) {
	return _NativeProver.contract.Transact(opts, "setInitialL1Config", _l1Configuration)
}

// SetInitialL1Config is a paid mutator transaction binding the contract method 0xd1156991.
//
// Solidity: function setInitialL1Config((address,address,uint256,uint256) _l1Configuration) returns()
func (_NativeProver *NativeProverSession) SetInitialL1Config(_l1Configuration L1Configuration) (*types.Transaction, error) {
	return _NativeProver.Contract.SetInitialL1Config(&_NativeProver.TransactOpts, _l1Configuration)
}

// SetInitialL1Config is a paid mutator transaction binding the contract method 0xd1156991.
//
// Solidity: function setInitialL1Config((address,address,uint256,uint256) _l1Configuration) returns()
func (_NativeProver *NativeProverTransactorSession) SetInitialL1Config(_l1Configuration L1Configuration) (*types.Transaction, error) {
	return _NativeProver.Contract.SetInitialL1Config(&_NativeProver.TransactOpts, _l1Configuration)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeProver *NativeProverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeProver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeProver *NativeProverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeProver.Contract.TransferOwnership(&_NativeProver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeProver *NativeProverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeProver.Contract.TransferOwnership(&_NativeProver.TransactOpts, newOwner)
}

// UpdateL1ChainConfiguration is a paid mutator transaction binding the contract method 0x4dd2b168.
//
// Solidity: function updateL1ChainConfiguration((address,address,uint256,uint256) _config, bytes[] _l1StorageProof, bytes _rlpEncodedRegistryAccountData, bytes[] _l1RegistryProof, bytes _rlpEncodedL1Header) returns()
func (_NativeProver *NativeProverTransactor) UpdateL1ChainConfiguration(opts *bind.TransactOpts, _config L1Configuration, _l1StorageProof [][]byte, _rlpEncodedRegistryAccountData []byte, _l1RegistryProof [][]byte, _rlpEncodedL1Header []byte) (*types.Transaction, error) {
	return _NativeProver.contract.Transact(opts, "updateL1ChainConfiguration", _config, _l1StorageProof, _rlpEncodedRegistryAccountData, _l1RegistryProof, _rlpEncodedL1Header)
}

// UpdateL1ChainConfiguration is a paid mutator transaction binding the contract method 0x4dd2b168.
//
// Solidity: function updateL1ChainConfiguration((address,address,uint256,uint256) _config, bytes[] _l1StorageProof, bytes _rlpEncodedRegistryAccountData, bytes[] _l1RegistryProof, bytes _rlpEncodedL1Header) returns()
func (_NativeProver *NativeProverSession) UpdateL1ChainConfiguration(_config L1Configuration, _l1StorageProof [][]byte, _rlpEncodedRegistryAccountData []byte, _l1RegistryProof [][]byte, _rlpEncodedL1Header []byte) (*types.Transaction, error) {
	return _NativeProver.Contract.UpdateL1ChainConfiguration(&_NativeProver.TransactOpts, _config, _l1StorageProof, _rlpEncodedRegistryAccountData, _l1RegistryProof, _rlpEncodedL1Header)
}

// UpdateL1ChainConfiguration is a paid mutator transaction binding the contract method 0x4dd2b168.
//
// Solidity: function updateL1ChainConfiguration((address,address,uint256,uint256) _config, bytes[] _l1StorageProof, bytes _rlpEncodedRegistryAccountData, bytes[] _l1RegistryProof, bytes _rlpEncodedL1Header) returns()
func (_NativeProver *NativeProverTransactorSession) UpdateL1ChainConfiguration(_config L1Configuration, _l1StorageProof [][]byte, _rlpEncodedRegistryAccountData []byte, _l1RegistryProof [][]byte, _rlpEncodedL1Header []byte) (*types.Transaction, error) {
	return _NativeProver.Contract.UpdateL1ChainConfiguration(&_NativeProver.TransactOpts, _config, _l1StorageProof, _rlpEncodedRegistryAccountData, _l1RegistryProof, _rlpEncodedL1Header)
}

// NativeProverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeProver contract.
type NativeProverOwnershipTransferredIterator struct {
	Event *NativeProverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NativeProverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeProverOwnershipTransferred)
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
		it.Event = new(NativeProverOwnershipTransferred)
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
func (it *NativeProverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeProverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeProverOwnershipTransferred represents a OwnershipTransferred event raised by the NativeProver contract.
type NativeProverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeProver *NativeProverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeProverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeProver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeProverOwnershipTransferredIterator{contract: _NativeProver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeProver *NativeProverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeProverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeProver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeProverOwnershipTransferred)
				if err := _NativeProver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NativeProver *NativeProverFilterer) ParseOwnershipTransferred(log types.Log) (*NativeProverOwnershipTransferred, error) {
	event := new(NativeProverOwnershipTransferred)
	if err := _NativeProver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
