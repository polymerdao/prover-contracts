// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crossl2proverv2

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

// CrossL2ProverV2MetaData contains all meta data concerning the CrossL2ProverV2 contract.
var CrossL2ProverV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"clientType_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sequencer_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LIGHT_CLIENT_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumLightClientType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SEQUENCER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clientType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inspectLogIdentifier\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"srcChain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiptIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"logIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"inspectPolymerState\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseEvent\",\"inputs\":[{\"name\":\"rawEvent\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"numTopics\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"emittingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"topics\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"unindexedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"validateEvent\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"emittingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"topics\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"unindexedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateSolLogs\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"programID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"logMessages\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"InvalidProofKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSequencerSignature\",\"inputs\":[]}]",
}

// CrossL2ProverV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossL2ProverV2MetaData.ABI instead.
var CrossL2ProverV2ABI = CrossL2ProverV2MetaData.ABI

// CrossL2ProverV2 is an auto generated Go binding around an Ethereum contract.
type CrossL2ProverV2 struct {
	CrossL2ProverV2Caller     // Read-only binding to the contract
	CrossL2ProverV2Transactor // Write-only binding to the contract
	CrossL2ProverV2Filterer   // Log filterer for contract events
}

// CrossL2ProverV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type CrossL2ProverV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossL2ProverV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossL2ProverV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossL2ProverV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossL2ProverV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossL2ProverV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossL2ProverV2Session struct {
	Contract     *CrossL2ProverV2  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossL2ProverV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossL2ProverV2CallerSession struct {
	Contract *CrossL2ProverV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CrossL2ProverV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossL2ProverV2TransactorSession struct {
	Contract     *CrossL2ProverV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CrossL2ProverV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type CrossL2ProverV2Raw struct {
	Contract *CrossL2ProverV2 // Generic contract binding to access the raw methods on
}

// CrossL2ProverV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossL2ProverV2CallerRaw struct {
	Contract *CrossL2ProverV2Caller // Generic read-only contract binding to access the raw methods on
}

// CrossL2ProverV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossL2ProverV2TransactorRaw struct {
	Contract *CrossL2ProverV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossL2ProverV2 creates a new instance of CrossL2ProverV2, bound to a specific deployed contract.
func NewCrossL2ProverV2(address common.Address, backend bind.ContractBackend) (*CrossL2ProverV2, error) {
	contract, err := bindCrossL2ProverV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossL2ProverV2{CrossL2ProverV2Caller: CrossL2ProverV2Caller{contract: contract}, CrossL2ProverV2Transactor: CrossL2ProverV2Transactor{contract: contract}, CrossL2ProverV2Filterer: CrossL2ProverV2Filterer{contract: contract}}, nil
}

// NewCrossL2ProverV2Caller creates a new read-only instance of CrossL2ProverV2, bound to a specific deployed contract.
func NewCrossL2ProverV2Caller(address common.Address, caller bind.ContractCaller) (*CrossL2ProverV2Caller, error) {
	contract, err := bindCrossL2ProverV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossL2ProverV2Caller{contract: contract}, nil
}

// NewCrossL2ProverV2Transactor creates a new write-only instance of CrossL2ProverV2, bound to a specific deployed contract.
func NewCrossL2ProverV2Transactor(address common.Address, transactor bind.ContractTransactor) (*CrossL2ProverV2Transactor, error) {
	contract, err := bindCrossL2ProverV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossL2ProverV2Transactor{contract: contract}, nil
}

// NewCrossL2ProverV2Filterer creates a new log filterer instance of CrossL2ProverV2, bound to a specific deployed contract.
func NewCrossL2ProverV2Filterer(address common.Address, filterer bind.ContractFilterer) (*CrossL2ProverV2Filterer, error) {
	contract, err := bindCrossL2ProverV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossL2ProverV2Filterer{contract: contract}, nil
}

// bindCrossL2ProverV2 binds a generic wrapper to an already deployed contract.
func bindCrossL2ProverV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossL2ProverV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossL2ProverV2 *CrossL2ProverV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossL2ProverV2.Contract.CrossL2ProverV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossL2ProverV2 *CrossL2ProverV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossL2ProverV2.Contract.CrossL2ProverV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossL2ProverV2 *CrossL2ProverV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossL2ProverV2.Contract.CrossL2ProverV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossL2ProverV2 *CrossL2ProverV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossL2ProverV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossL2ProverV2 *CrossL2ProverV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossL2ProverV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossL2ProverV2 *CrossL2ProverV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossL2ProverV2.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(bytes32)
func (_CrossL2ProverV2 *CrossL2ProverV2Caller) CHAINID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossL2ProverV2.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(bytes32)
func (_CrossL2ProverV2 *CrossL2ProverV2Session) CHAINID() ([32]byte, error) {
	return _CrossL2ProverV2.Contract.CHAINID(&_CrossL2ProverV2.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(bytes32)
func (_CrossL2ProverV2 *CrossL2ProverV2CallerSession) CHAINID() ([32]byte, error) {
	return _CrossL2ProverV2.Contract.CHAINID(&_CrossL2ProverV2.CallOpts)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_CrossL2ProverV2 *CrossL2ProverV2Caller) LIGHTCLIENTTYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossL2ProverV2.contract.Call(opts, &out, "LIGHT_CLIENT_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_CrossL2ProverV2 *CrossL2ProverV2Session) LIGHTCLIENTTYPE() (uint8, error) {
	return _CrossL2ProverV2.Contract.LIGHTCLIENTTYPE(&_CrossL2ProverV2.CallOpts)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_CrossL2ProverV2 *CrossL2ProverV2CallerSession) LIGHTCLIENTTYPE() (uint8, error) {
	return _CrossL2ProverV2.Contract.LIGHTCLIENTTYPE(&_CrossL2ProverV2.CallOpts)
}

// SEQUENCER is a free data retrieval call binding the contract method 0x75fd4ca9.
//
// Solidity: function SEQUENCER() view returns(address)
func (_CrossL2ProverV2 *CrossL2ProverV2Caller) SEQUENCER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossL2ProverV2.contract.Call(opts, &out, "SEQUENCER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SEQUENCER is a free data retrieval call binding the contract method 0x75fd4ca9.
//
// Solidity: function SEQUENCER() view returns(address)
func (_CrossL2ProverV2 *CrossL2ProverV2Session) SEQUENCER() (common.Address, error) {
	return _CrossL2ProverV2.Contract.SEQUENCER(&_CrossL2ProverV2.CallOpts)
}

// SEQUENCER is a free data retrieval call binding the contract method 0x75fd4ca9.
//
// Solidity: function SEQUENCER() view returns(address)
func (_CrossL2ProverV2 *CrossL2ProverV2CallerSession) SEQUENCER() (common.Address, error) {
	return _CrossL2ProverV2.Contract.SEQUENCER(&_CrossL2ProverV2.CallOpts)
}

// ClientType is a free data retrieval call binding the contract method 0xb3768f0d.
//
// Solidity: function clientType() view returns(string)
func (_CrossL2ProverV2 *CrossL2ProverV2Caller) ClientType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossL2ProverV2.contract.Call(opts, &out, "clientType")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ClientType is a free data retrieval call binding the contract method 0xb3768f0d.
//
// Solidity: function clientType() view returns(string)
func (_CrossL2ProverV2 *CrossL2ProverV2Session) ClientType() (string, error) {
	return _CrossL2ProverV2.Contract.ClientType(&_CrossL2ProverV2.CallOpts)
}

// ClientType is a free data retrieval call binding the contract method 0xb3768f0d.
//
// Solidity: function clientType() view returns(string)
func (_CrossL2ProverV2 *CrossL2ProverV2CallerSession) ClientType() (string, error) {
	return _CrossL2ProverV2.Contract.ClientType(&_CrossL2ProverV2.CallOpts)
}

// InspectLogIdentifier is a free data retrieval call binding the contract method 0x9e79c3f0.
//
// Solidity: function inspectLogIdentifier(bytes proof) pure returns(uint32 srcChain, uint64 blockNumber, uint16 receiptIndex, uint8 logIndex)
func (_CrossL2ProverV2 *CrossL2ProverV2Caller) InspectLogIdentifier(opts *bind.CallOpts, proof []byte) (struct {
	SrcChain     uint32
	BlockNumber  uint64
	ReceiptIndex uint16
	LogIndex     uint8
}, error) {
	var out []interface{}
	err := _CrossL2ProverV2.contract.Call(opts, &out, "inspectLogIdentifier", proof)

	outstruct := new(struct {
		SrcChain     uint32
		BlockNumber  uint64
		ReceiptIndex uint16
		LogIndex     uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SrcChain = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ReceiptIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.LogIndex = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// InspectLogIdentifier is a free data retrieval call binding the contract method 0x9e79c3f0.
//
// Solidity: function inspectLogIdentifier(bytes proof) pure returns(uint32 srcChain, uint64 blockNumber, uint16 receiptIndex, uint8 logIndex)
func (_CrossL2ProverV2 *CrossL2ProverV2Session) InspectLogIdentifier(proof []byte) (struct {
	SrcChain     uint32
	BlockNumber  uint64
	ReceiptIndex uint16
	LogIndex     uint8
}, error) {
	return _CrossL2ProverV2.Contract.InspectLogIdentifier(&_CrossL2ProverV2.CallOpts, proof)
}

// InspectLogIdentifier is a free data retrieval call binding the contract method 0x9e79c3f0.
//
// Solidity: function inspectLogIdentifier(bytes proof) pure returns(uint32 srcChain, uint64 blockNumber, uint16 receiptIndex, uint8 logIndex)
func (_CrossL2ProverV2 *CrossL2ProverV2CallerSession) InspectLogIdentifier(proof []byte) (struct {
	SrcChain     uint32
	BlockNumber  uint64
	ReceiptIndex uint16
	LogIndex     uint8
}, error) {
	return _CrossL2ProverV2.Contract.InspectLogIdentifier(&_CrossL2ProverV2.CallOpts, proof)
}

// InspectPolymerState is a free data retrieval call binding the contract method 0xd9615898.
//
// Solidity: function inspectPolymerState(bytes proof) pure returns(bytes32 stateRoot, uint64 height, bytes signature)
func (_CrossL2ProverV2 *CrossL2ProverV2Caller) InspectPolymerState(opts *bind.CallOpts, proof []byte) (struct {
	StateRoot [32]byte
	Height    uint64
	Signature []byte
}, error) {
	var out []interface{}
	err := _CrossL2ProverV2.contract.Call(opts, &out, "inspectPolymerState", proof)

	outstruct := new(struct {
		StateRoot [32]byte
		Height    uint64
		Signature []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StateRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Height = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Signature = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// InspectPolymerState is a free data retrieval call binding the contract method 0xd9615898.
//
// Solidity: function inspectPolymerState(bytes proof) pure returns(bytes32 stateRoot, uint64 height, bytes signature)
func (_CrossL2ProverV2 *CrossL2ProverV2Session) InspectPolymerState(proof []byte) (struct {
	StateRoot [32]byte
	Height    uint64
	Signature []byte
}, error) {
	return _CrossL2ProverV2.Contract.InspectPolymerState(&_CrossL2ProverV2.CallOpts, proof)
}

// InspectPolymerState is a free data retrieval call binding the contract method 0xd9615898.
//
// Solidity: function inspectPolymerState(bytes proof) pure returns(bytes32 stateRoot, uint64 height, bytes signature)
func (_CrossL2ProverV2 *CrossL2ProverV2CallerSession) InspectPolymerState(proof []byte) (struct {
	StateRoot [32]byte
	Height    uint64
	Signature []byte
}, error) {
	return _CrossL2ProverV2.Contract.InspectPolymerState(&_CrossL2ProverV2.CallOpts, proof)
}

// ParseEvent is a free data retrieval call binding the contract method 0x0d15d928.
//
// Solidity: function parseEvent(bytes rawEvent, uint8 numTopics) pure returns(address emittingContract, bytes topics, bytes unindexedData)
func (_CrossL2ProverV2 *CrossL2ProverV2Caller) ParseEvent(opts *bind.CallOpts, rawEvent []byte, numTopics uint8) (struct {
	EmittingContract common.Address
	Topics           []byte
	UnindexedData    []byte
}, error) {
	var out []interface{}
	err := _CrossL2ProverV2.contract.Call(opts, &out, "parseEvent", rawEvent, numTopics)

	outstruct := new(struct {
		EmittingContract common.Address
		Topics           []byte
		UnindexedData    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EmittingContract = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Topics = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.UnindexedData = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ParseEvent is a free data retrieval call binding the contract method 0x0d15d928.
//
// Solidity: function parseEvent(bytes rawEvent, uint8 numTopics) pure returns(address emittingContract, bytes topics, bytes unindexedData)
func (_CrossL2ProverV2 *CrossL2ProverV2Session) ParseEvent(rawEvent []byte, numTopics uint8) (struct {
	EmittingContract common.Address
	Topics           []byte
	UnindexedData    []byte
}, error) {
	return _CrossL2ProverV2.Contract.ParseEvent(&_CrossL2ProverV2.CallOpts, rawEvent, numTopics)
}

// ParseEvent is a free data retrieval call binding the contract method 0x0d15d928.
//
// Solidity: function parseEvent(bytes rawEvent, uint8 numTopics) pure returns(address emittingContract, bytes topics, bytes unindexedData)
func (_CrossL2ProverV2 *CrossL2ProverV2CallerSession) ParseEvent(rawEvent []byte, numTopics uint8) (struct {
	EmittingContract common.Address
	Topics           []byte
	UnindexedData    []byte
}, error) {
	return _CrossL2ProverV2.Contract.ParseEvent(&_CrossL2ProverV2.CallOpts, rawEvent, numTopics)
}

// ValidateEvent is a free data retrieval call binding the contract method 0x0b9ca54a.
//
// Solidity: function validateEvent(bytes proof) view returns(uint32 chainId, address emittingContract, bytes topics, bytes unindexedData)
func (_CrossL2ProverV2 *CrossL2ProverV2Caller) ValidateEvent(opts *bind.CallOpts, proof []byte) (struct {
	ChainId          uint32
	EmittingContract common.Address
	Topics           []byte
	UnindexedData    []byte
}, error) {
	var out []interface{}
	err := _CrossL2ProverV2.contract.Call(opts, &out, "validateEvent", proof)

	outstruct := new(struct {
		ChainId          uint32
		EmittingContract common.Address
		Topics           []byte
		UnindexedData    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.EmittingContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Topics = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.UnindexedData = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ValidateEvent is a free data retrieval call binding the contract method 0x0b9ca54a.
//
// Solidity: function validateEvent(bytes proof) view returns(uint32 chainId, address emittingContract, bytes topics, bytes unindexedData)
func (_CrossL2ProverV2 *CrossL2ProverV2Session) ValidateEvent(proof []byte) (struct {
	ChainId          uint32
	EmittingContract common.Address
	Topics           []byte
	UnindexedData    []byte
}, error) {
	return _CrossL2ProverV2.Contract.ValidateEvent(&_CrossL2ProverV2.CallOpts, proof)
}

// ValidateEvent is a free data retrieval call binding the contract method 0x0b9ca54a.
//
// Solidity: function validateEvent(bytes proof) view returns(uint32 chainId, address emittingContract, bytes topics, bytes unindexedData)
func (_CrossL2ProverV2 *CrossL2ProverV2CallerSession) ValidateEvent(proof []byte) (struct {
	ChainId          uint32
	EmittingContract common.Address
	Topics           []byte
	UnindexedData    []byte
}, error) {
	return _CrossL2ProverV2.Contract.ValidateEvent(&_CrossL2ProverV2.CallOpts, proof)
}

// ValidateSolLogs is a free data retrieval call binding the contract method 0xd73a8ad6.
//
// Solidity: function validateSolLogs(bytes proof) view returns(uint32 chainId, bytes32 programID, string[] logMessages)
func (_CrossL2ProverV2 *CrossL2ProverV2Caller) ValidateSolLogs(opts *bind.CallOpts, proof []byte) (struct {
	ChainId     uint32
	ProgramID   [32]byte
	LogMessages []string
}, error) {
	var out []interface{}
	err := _CrossL2ProverV2.contract.Call(opts, &out, "validateSolLogs", proof)

	outstruct := new(struct {
		ChainId     uint32
		ProgramID   [32]byte
		LogMessages []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ProgramID = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.LogMessages = *abi.ConvertType(out[2], new([]string)).(*[]string)

	return *outstruct, err

}

// ValidateSolLogs is a free data retrieval call binding the contract method 0xd73a8ad6.
//
// Solidity: function validateSolLogs(bytes proof) view returns(uint32 chainId, bytes32 programID, string[] logMessages)
func (_CrossL2ProverV2 *CrossL2ProverV2Session) ValidateSolLogs(proof []byte) (struct {
	ChainId     uint32
	ProgramID   [32]byte
	LogMessages []string
}, error) {
	return _CrossL2ProverV2.Contract.ValidateSolLogs(&_CrossL2ProverV2.CallOpts, proof)
}

// ValidateSolLogs is a free data retrieval call binding the contract method 0xd73a8ad6.
//
// Solidity: function validateSolLogs(bytes proof) view returns(uint32 chainId, bytes32 programID, string[] logMessages)
func (_CrossL2ProverV2 *CrossL2ProverV2CallerSession) ValidateSolLogs(proof []byte) (struct {
	ChainId     uint32
	ProgramID   [32]byte
	LogMessages []string
}, error) {
	return _CrossL2ProverV2.Contract.ValidateSolLogs(&_CrossL2ProverV2.CallOpts, proof)
}

// VerifyMembership is a free data retrieval call binding the contract method 0x861f2843.
//
// Solidity: function verifyMembership(bytes32 root, bytes key, bytes32 value, bytes proof) pure returns()
func (_CrossL2ProverV2 *CrossL2ProverV2Caller) VerifyMembership(opts *bind.CallOpts, root [32]byte, key []byte, value [32]byte, proof []byte) error {
	var out []interface{}
	err := _CrossL2ProverV2.contract.Call(opts, &out, "verifyMembership", root, key, value, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyMembership is a free data retrieval call binding the contract method 0x861f2843.
//
// Solidity: function verifyMembership(bytes32 root, bytes key, bytes32 value, bytes proof) pure returns()
func (_CrossL2ProverV2 *CrossL2ProverV2Session) VerifyMembership(root [32]byte, key []byte, value [32]byte, proof []byte) error {
	return _CrossL2ProverV2.Contract.VerifyMembership(&_CrossL2ProverV2.CallOpts, root, key, value, proof)
}

// VerifyMembership is a free data retrieval call binding the contract method 0x861f2843.
//
// Solidity: function verifyMembership(bytes32 root, bytes key, bytes32 value, bytes proof) pure returns()
func (_CrossL2ProverV2 *CrossL2ProverV2CallerSession) VerifyMembership(root [32]byte, key []byte, value [32]byte, proof []byte) error {
	return _CrossL2ProverV2.Contract.VerifyMembership(&_CrossL2ProverV2.CallOpts, root, key, value, proof)
}
