// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package opstackbedrockprover

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

// L2Configuration is an auto generated low-level Go binding around an user-defined struct.
type L2Configuration struct {
	Prover               common.Address
	Addresses            []common.Address
	StorageSlots         []*big.Int
	VersionNumber        *big.Int
	FinalityDelaySeconds *big.Int
	L2Type               uint8
}

// OPStackBedrockProverMetaData contains all meta data concerning the OPStackBedrockProver contract.
var OPStackBedrockProverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"proveSettledState\",\"inputs\":[{\"name\":\"_chainConfig\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"_l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_rlpEncodedL2Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l1WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"BlockBeforeFinalityPeriod\",\"inputs\":[{\"name\":\"_blockTimeStamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_finalityDelayTimeStamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"IncorrectOutputOracleStorageRoot\",\"inputs\":[{\"name\":\"_outputOracleStateRoot\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountProof\",\"inputs\":[{\"name\":\"_address\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidBedrockProof\",\"inputs\":[{\"name\":\"_l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_l1WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidStorageProof\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_val\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// OPStackBedrockProverABI is the input ABI used to generate the binding from.
// Deprecated: Use OPStackBedrockProverMetaData.ABI instead.
var OPStackBedrockProverABI = OPStackBedrockProverMetaData.ABI

// OPStackBedrockProver is an auto generated Go binding around an Ethereum contract.
type OPStackBedrockProver struct {
	OPStackBedrockProverCaller     // Read-only binding to the contract
	OPStackBedrockProverTransactor // Write-only binding to the contract
	OPStackBedrockProverFilterer   // Log filterer for contract events
}

// OPStackBedrockProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type OPStackBedrockProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPStackBedrockProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OPStackBedrockProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPStackBedrockProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OPStackBedrockProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPStackBedrockProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OPStackBedrockProverSession struct {
	Contract     *OPStackBedrockProver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OPStackBedrockProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OPStackBedrockProverCallerSession struct {
	Contract *OPStackBedrockProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OPStackBedrockProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OPStackBedrockProverTransactorSession struct {
	Contract     *OPStackBedrockProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OPStackBedrockProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type OPStackBedrockProverRaw struct {
	Contract *OPStackBedrockProver // Generic contract binding to access the raw methods on
}

// OPStackBedrockProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OPStackBedrockProverCallerRaw struct {
	Contract *OPStackBedrockProverCaller // Generic read-only contract binding to access the raw methods on
}

// OPStackBedrockProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OPStackBedrockProverTransactorRaw struct {
	Contract *OPStackBedrockProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOPStackBedrockProver creates a new instance of OPStackBedrockProver, bound to a specific deployed contract.
func NewOPStackBedrockProver(address common.Address, backend bind.ContractBackend) (*OPStackBedrockProver, error) {
	contract, err := bindOPStackBedrockProver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OPStackBedrockProver{OPStackBedrockProverCaller: OPStackBedrockProverCaller{contract: contract}, OPStackBedrockProverTransactor: OPStackBedrockProverTransactor{contract: contract}, OPStackBedrockProverFilterer: OPStackBedrockProverFilterer{contract: contract}}, nil
}

// NewOPStackBedrockProverCaller creates a new read-only instance of OPStackBedrockProver, bound to a specific deployed contract.
func NewOPStackBedrockProverCaller(address common.Address, caller bind.ContractCaller) (*OPStackBedrockProverCaller, error) {
	contract, err := bindOPStackBedrockProver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OPStackBedrockProverCaller{contract: contract}, nil
}

// NewOPStackBedrockProverTransactor creates a new write-only instance of OPStackBedrockProver, bound to a specific deployed contract.
func NewOPStackBedrockProverTransactor(address common.Address, transactor bind.ContractTransactor) (*OPStackBedrockProverTransactor, error) {
	contract, err := bindOPStackBedrockProver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OPStackBedrockProverTransactor{contract: contract}, nil
}

// NewOPStackBedrockProverFilterer creates a new log filterer instance of OPStackBedrockProver, bound to a specific deployed contract.
func NewOPStackBedrockProverFilterer(address common.Address, filterer bind.ContractFilterer) (*OPStackBedrockProverFilterer, error) {
	contract, err := bindOPStackBedrockProver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OPStackBedrockProverFilterer{contract: contract}, nil
}

// bindOPStackBedrockProver binds a generic wrapper to an already deployed contract.
func bindOPStackBedrockProver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OPStackBedrockProverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPStackBedrockProver *OPStackBedrockProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPStackBedrockProver.Contract.OPStackBedrockProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPStackBedrockProver *OPStackBedrockProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPStackBedrockProver.Contract.OPStackBedrockProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPStackBedrockProver *OPStackBedrockProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPStackBedrockProver.Contract.OPStackBedrockProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPStackBedrockProver *OPStackBedrockProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPStackBedrockProver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPStackBedrockProver *OPStackBedrockProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPStackBedrockProver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPStackBedrockProver *OPStackBedrockProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPStackBedrockProver.Contract.contract.Transact(opts, method, params...)
}

// ProveSettledState is a free data retrieval call binding the contract method 0x52aba3d8.
//
// Solidity: function proveSettledState((address,address[],uint256[],uint256,uint256,uint8) _chainConfig, bytes32 _l2WorldStateRoot, bytes _rlpEncodedL2Header, bytes32 _l1WorldStateRoot, bytes _proof) view returns(bool)
func (_OPStackBedrockProver *OPStackBedrockProverCaller) ProveSettledState(opts *bind.CallOpts, _chainConfig L2Configuration, _l2WorldStateRoot [32]byte, _rlpEncodedL2Header []byte, _l1WorldStateRoot [32]byte, _proof []byte) (bool, error) {
	var out []interface{}
	err := _OPStackBedrockProver.contract.Call(opts, &out, "proveSettledState", _chainConfig, _l2WorldStateRoot, _rlpEncodedL2Header, _l1WorldStateRoot, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveSettledState is a free data retrieval call binding the contract method 0x52aba3d8.
//
// Solidity: function proveSettledState((address,address[],uint256[],uint256,uint256,uint8) _chainConfig, bytes32 _l2WorldStateRoot, bytes _rlpEncodedL2Header, bytes32 _l1WorldStateRoot, bytes _proof) view returns(bool)
func (_OPStackBedrockProver *OPStackBedrockProverSession) ProveSettledState(_chainConfig L2Configuration, _l2WorldStateRoot [32]byte, _rlpEncodedL2Header []byte, _l1WorldStateRoot [32]byte, _proof []byte) (bool, error) {
	return _OPStackBedrockProver.Contract.ProveSettledState(&_OPStackBedrockProver.CallOpts, _chainConfig, _l2WorldStateRoot, _rlpEncodedL2Header, _l1WorldStateRoot, _proof)
}

// ProveSettledState is a free data retrieval call binding the contract method 0x52aba3d8.
//
// Solidity: function proveSettledState((address,address[],uint256[],uint256,uint256,uint8) _chainConfig, bytes32 _l2WorldStateRoot, bytes _rlpEncodedL2Header, bytes32 _l1WorldStateRoot, bytes _proof) view returns(bool)
func (_OPStackBedrockProver *OPStackBedrockProverCallerSession) ProveSettledState(_chainConfig L2Configuration, _l2WorldStateRoot [32]byte, _rlpEncodedL2Header []byte, _l1WorldStateRoot [32]byte, _proof []byte) (bool, error) {
	return _OPStackBedrockProver.Contract.ProveSettledState(&_OPStackBedrockProver.CallOpts, _chainConfig, _l2WorldStateRoot, _rlpEncodedL2Header, _l1WorldStateRoot, _proof)
}
