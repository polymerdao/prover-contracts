// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package receiptparser

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

// ReceiptParserMetaData contains all meta data concerning the ReceiptParser contract.
var ReceiptParserMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"bytesToAddr\",\"inputs\":[{\"name\":\"a\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"invalidAddressBytes\",\"inputs\":[]}]",
}

// ReceiptParserABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiptParserMetaData.ABI instead.
var ReceiptParserABI = ReceiptParserMetaData.ABI

// ReceiptParser is an auto generated Go binding around an Ethereum contract.
type ReceiptParser struct {
	ReceiptParserCaller     // Read-only binding to the contract
	ReceiptParserTransactor // Write-only binding to the contract
	ReceiptParserFilterer   // Log filterer for contract events
}

// ReceiptParserCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiptParserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptParserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiptParserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptParserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiptParserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptParserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiptParserSession struct {
	Contract     *ReceiptParser    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiptParserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiptParserCallerSession struct {
	Contract *ReceiptParserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ReceiptParserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiptParserTransactorSession struct {
	Contract     *ReceiptParserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ReceiptParserRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiptParserRaw struct {
	Contract *ReceiptParser // Generic contract binding to access the raw methods on
}

// ReceiptParserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiptParserCallerRaw struct {
	Contract *ReceiptParserCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiptParserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiptParserTransactorRaw struct {
	Contract *ReceiptParserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiptParser creates a new instance of ReceiptParser, bound to a specific deployed contract.
func NewReceiptParser(address common.Address, backend bind.ContractBackend) (*ReceiptParser, error) {
	contract, err := bindReceiptParser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiptParser{ReceiptParserCaller: ReceiptParserCaller{contract: contract}, ReceiptParserTransactor: ReceiptParserTransactor{contract: contract}, ReceiptParserFilterer: ReceiptParserFilterer{contract: contract}}, nil
}

// NewReceiptParserCaller creates a new read-only instance of ReceiptParser, bound to a specific deployed contract.
func NewReceiptParserCaller(address common.Address, caller bind.ContractCaller) (*ReceiptParserCaller, error) {
	contract, err := bindReceiptParser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptParserCaller{contract: contract}, nil
}

// NewReceiptParserTransactor creates a new write-only instance of ReceiptParser, bound to a specific deployed contract.
func NewReceiptParserTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiptParserTransactor, error) {
	contract, err := bindReceiptParser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptParserTransactor{contract: contract}, nil
}

// NewReceiptParserFilterer creates a new log filterer instance of ReceiptParser, bound to a specific deployed contract.
func NewReceiptParserFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiptParserFilterer, error) {
	contract, err := bindReceiptParser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiptParserFilterer{contract: contract}, nil
}

// bindReceiptParser binds a generic wrapper to an already deployed contract.
func bindReceiptParser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReceiptParserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptParser *ReceiptParserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptParser.Contract.ReceiptParserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptParser *ReceiptParserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptParser.Contract.ReceiptParserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptParser *ReceiptParserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptParser.Contract.ReceiptParserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptParser *ReceiptParserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptParser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptParser *ReceiptParserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptParser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptParser *ReceiptParserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptParser.Contract.contract.Transact(opts, method, params...)
}

// BytesToAddr is a free data retrieval call binding the contract method 0x5ef6228c.
//
// Solidity: function bytesToAddr(bytes a) pure returns(address addr)
func (_ReceiptParser *ReceiptParserCaller) BytesToAddr(opts *bind.CallOpts, a []byte) (common.Address, error) {
	var out []interface{}
	err := _ReceiptParser.contract.Call(opts, &out, "bytesToAddr", a)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BytesToAddr is a free data retrieval call binding the contract method 0x5ef6228c.
//
// Solidity: function bytesToAddr(bytes a) pure returns(address addr)
func (_ReceiptParser *ReceiptParserSession) BytesToAddr(a []byte) (common.Address, error) {
	return _ReceiptParser.Contract.BytesToAddr(&_ReceiptParser.CallOpts, a)
}

// BytesToAddr is a free data retrieval call binding the contract method 0x5ef6228c.
//
// Solidity: function bytesToAddr(bytes a) pure returns(address addr)
func (_ReceiptParser *ReceiptParserCallerSession) BytesToAddr(a []byte) (common.Address, error) {
	return _ReceiptParser.Contract.BytesToAddr(&_ReceiptParser.CallOpts, a)
}
