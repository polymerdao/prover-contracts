// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package opstackcannonprover

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

// OPStackCannonProverMetaData contains all meta data concerning the OPStackCannonProver contract.
var OPStackCannonProverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"proveSettledState\",\"inputs\":[{\"name\":\"_chainConfig\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"enumType\"}]},{\"name\":\"_l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_rlpEncodedL2Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l1WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"FaultDisputeGameUnresolved\",\"inputs\":[{\"name\":\"_gameStatus\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"IncorrectDisputeGameFactoryStateRoot\",\"inputs\":[{\"name\":\"_disputeGameFactoryStateRoot\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountProof\",\"inputs\":[{\"name\":\"_address\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidCannonProof\",\"inputs\":[{\"name\":\"_l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_l1WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedBlock\",\"inputs\":[{\"name\":\"_expectedBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_calculatedBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidStorageProof\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_val\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidStorageProof\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_val\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// OPStackCannonProverABI is the input ABI used to generate the binding from.
// Deprecated: Use OPStackCannonProverMetaData.ABI instead.
var OPStackCannonProverABI = OPStackCannonProverMetaData.ABI

// OPStackCannonProver is an auto generated Go binding around an Ethereum contract.
type OPStackCannonProver struct {
	OPStackCannonProverCaller     // Read-only binding to the contract
	OPStackCannonProverTransactor // Write-only binding to the contract
	OPStackCannonProverFilterer   // Log filterer for contract events
}

// OPStackCannonProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type OPStackCannonProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPStackCannonProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OPStackCannonProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPStackCannonProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OPStackCannonProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPStackCannonProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OPStackCannonProverSession struct {
	Contract     *OPStackCannonProver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OPStackCannonProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OPStackCannonProverCallerSession struct {
	Contract *OPStackCannonProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OPStackCannonProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OPStackCannonProverTransactorSession struct {
	Contract     *OPStackCannonProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OPStackCannonProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type OPStackCannonProverRaw struct {
	Contract *OPStackCannonProver // Generic contract binding to access the raw methods on
}

// OPStackCannonProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OPStackCannonProverCallerRaw struct {
	Contract *OPStackCannonProverCaller // Generic read-only contract binding to access the raw methods on
}

// OPStackCannonProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OPStackCannonProverTransactorRaw struct {
	Contract *OPStackCannonProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOPStackCannonProver creates a new instance of OPStackCannonProver, bound to a specific deployed contract.
func NewOPStackCannonProver(address common.Address, backend bind.ContractBackend) (*OPStackCannonProver, error) {
	contract, err := bindOPStackCannonProver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OPStackCannonProver{OPStackCannonProverCaller: OPStackCannonProverCaller{contract: contract}, OPStackCannonProverTransactor: OPStackCannonProverTransactor{contract: contract}, OPStackCannonProverFilterer: OPStackCannonProverFilterer{contract: contract}}, nil
}

// NewOPStackCannonProverCaller creates a new read-only instance of OPStackCannonProver, bound to a specific deployed contract.
func NewOPStackCannonProverCaller(address common.Address, caller bind.ContractCaller) (*OPStackCannonProverCaller, error) {
	contract, err := bindOPStackCannonProver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OPStackCannonProverCaller{contract: contract}, nil
}

// NewOPStackCannonProverTransactor creates a new write-only instance of OPStackCannonProver, bound to a specific deployed contract.
func NewOPStackCannonProverTransactor(address common.Address, transactor bind.ContractTransactor) (*OPStackCannonProverTransactor, error) {
	contract, err := bindOPStackCannonProver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OPStackCannonProverTransactor{contract: contract}, nil
}

// NewOPStackCannonProverFilterer creates a new log filterer instance of OPStackCannonProver, bound to a specific deployed contract.
func NewOPStackCannonProverFilterer(address common.Address, filterer bind.ContractFilterer) (*OPStackCannonProverFilterer, error) {
	contract, err := bindOPStackCannonProver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OPStackCannonProverFilterer{contract: contract}, nil
}

// bindOPStackCannonProver binds a generic wrapper to an already deployed contract.
func bindOPStackCannonProver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OPStackCannonProverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPStackCannonProver *OPStackCannonProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPStackCannonProver.Contract.OPStackCannonProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPStackCannonProver *OPStackCannonProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPStackCannonProver.Contract.OPStackCannonProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPStackCannonProver *OPStackCannonProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPStackCannonProver.Contract.OPStackCannonProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPStackCannonProver *OPStackCannonProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPStackCannonProver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPStackCannonProver *OPStackCannonProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPStackCannonProver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPStackCannonProver *OPStackCannonProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPStackCannonProver.Contract.contract.Transact(opts, method, params...)
}

// ProveSettledState is a free data retrieval call binding the contract method 0x52aba3d8.
//
// Solidity: function proveSettledState((address,address[],uint256[],uint256,uint256,uint8) _chainConfig, bytes32 _l2WorldStateRoot, bytes _rlpEncodedL2Header, bytes32 _l1WorldStateRoot, bytes _proof) pure returns(bool)
func (_OPStackCannonProver *OPStackCannonProverCaller) ProveSettledState(opts *bind.CallOpts, _chainConfig L2Configuration, _l2WorldStateRoot [32]byte, _rlpEncodedL2Header []byte, _l1WorldStateRoot [32]byte, _proof []byte) (bool, error) {
	var out []interface{}
	err := _OPStackCannonProver.contract.Call(opts, &out, "proveSettledState", _chainConfig, _l2WorldStateRoot, _rlpEncodedL2Header, _l1WorldStateRoot, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveSettledState is a free data retrieval call binding the contract method 0x52aba3d8.
//
// Solidity: function proveSettledState((address,address[],uint256[],uint256,uint256,uint8) _chainConfig, bytes32 _l2WorldStateRoot, bytes _rlpEncodedL2Header, bytes32 _l1WorldStateRoot, bytes _proof) pure returns(bool)
func (_OPStackCannonProver *OPStackCannonProverSession) ProveSettledState(_chainConfig L2Configuration, _l2WorldStateRoot [32]byte, _rlpEncodedL2Header []byte, _l1WorldStateRoot [32]byte, _proof []byte) (bool, error) {
	return _OPStackCannonProver.Contract.ProveSettledState(&_OPStackCannonProver.CallOpts, _chainConfig, _l2WorldStateRoot, _rlpEncodedL2Header, _l1WorldStateRoot, _proof)
}

// ProveSettledState is a free data retrieval call binding the contract method 0x52aba3d8.
//
// Solidity: function proveSettledState((address,address[],uint256[],uint256,uint256,uint8) _chainConfig, bytes32 _l2WorldStateRoot, bytes _rlpEncodedL2Header, bytes32 _l1WorldStateRoot, bytes _proof) pure returns(bool)
func (_OPStackCannonProver *OPStackCannonProverCallerSession) ProveSettledState(_chainConfig L2Configuration, _l2WorldStateRoot [32]byte, _rlpEncodedL2Header []byte, _l1WorldStateRoot [32]byte, _proof []byte) (bool, error) {
	return _OPStackCannonProver.Contract.ProveSettledState(&_OPStackCannonProver.CallOpts, _chainConfig, _l2WorldStateRoot, _rlpEncodedL2Header, _l1WorldStateRoot, _proof)
}
