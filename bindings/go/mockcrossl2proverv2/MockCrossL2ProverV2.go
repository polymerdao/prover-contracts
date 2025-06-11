// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockcrossl2proverv2

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

// MockCrossL2ProverV2MetaData contains all meta data concerning the MockCrossL2ProverV2 contract.
var MockCrossL2ProverV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"clientType_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sequencer_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LIGHT_CLIENT_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumLightClientType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SEQUENCER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clientType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generateAndEmitProof\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"topics\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generateAndSendProof\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"topics\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"validatorContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generateMockProof\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numTopics\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"topics_\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"unindexedData_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"inspectLogIdentifier\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"srcChain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiptIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"logIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"inspectPolymerState\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseEvent\",\"inputs\":[{\"name\":\"rawEvent\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"numTopics\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"emittingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"topics\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"unindexedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"ping\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateEvent\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"emittingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"topics\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"unindexedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateSolLogs\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"programID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"logMessages\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"Ping\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProofGenerated\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidProofRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSequencerSignature\",\"inputs\":[]}]",
}

// MockCrossL2ProverV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MockCrossL2ProverV2MetaData.ABI instead.
var MockCrossL2ProverV2ABI = MockCrossL2ProverV2MetaData.ABI

// MockCrossL2ProverV2 is an auto generated Go binding around an Ethereum contract.
type MockCrossL2ProverV2 struct {
	MockCrossL2ProverV2Caller     // Read-only binding to the contract
	MockCrossL2ProverV2Transactor // Write-only binding to the contract
	MockCrossL2ProverV2Filterer   // Log filterer for contract events
}

// MockCrossL2ProverV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MockCrossL2ProverV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockCrossL2ProverV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MockCrossL2ProverV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockCrossL2ProverV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockCrossL2ProverV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockCrossL2ProverV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockCrossL2ProverV2Session struct {
	Contract     *MockCrossL2ProverV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MockCrossL2ProverV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockCrossL2ProverV2CallerSession struct {
	Contract *MockCrossL2ProverV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MockCrossL2ProverV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockCrossL2ProverV2TransactorSession struct {
	Contract     *MockCrossL2ProverV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MockCrossL2ProverV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MockCrossL2ProverV2Raw struct {
	Contract *MockCrossL2ProverV2 // Generic contract binding to access the raw methods on
}

// MockCrossL2ProverV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockCrossL2ProverV2CallerRaw struct {
	Contract *MockCrossL2ProverV2Caller // Generic read-only contract binding to access the raw methods on
}

// MockCrossL2ProverV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockCrossL2ProverV2TransactorRaw struct {
	Contract *MockCrossL2ProverV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMockCrossL2ProverV2 creates a new instance of MockCrossL2ProverV2, bound to a specific deployed contract.
func NewMockCrossL2ProverV2(address common.Address, backend bind.ContractBackend) (*MockCrossL2ProverV2, error) {
	contract, err := bindMockCrossL2ProverV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockCrossL2ProverV2{MockCrossL2ProverV2Caller: MockCrossL2ProverV2Caller{contract: contract}, MockCrossL2ProverV2Transactor: MockCrossL2ProverV2Transactor{contract: contract}, MockCrossL2ProverV2Filterer: MockCrossL2ProverV2Filterer{contract: contract}}, nil
}

// NewMockCrossL2ProverV2Caller creates a new read-only instance of MockCrossL2ProverV2, bound to a specific deployed contract.
func NewMockCrossL2ProverV2Caller(address common.Address, caller bind.ContractCaller) (*MockCrossL2ProverV2Caller, error) {
	contract, err := bindMockCrossL2ProverV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockCrossL2ProverV2Caller{contract: contract}, nil
}

// NewMockCrossL2ProverV2Transactor creates a new write-only instance of MockCrossL2ProverV2, bound to a specific deployed contract.
func NewMockCrossL2ProverV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MockCrossL2ProverV2Transactor, error) {
	contract, err := bindMockCrossL2ProverV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockCrossL2ProverV2Transactor{contract: contract}, nil
}

// NewMockCrossL2ProverV2Filterer creates a new log filterer instance of MockCrossL2ProverV2, bound to a specific deployed contract.
func NewMockCrossL2ProverV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MockCrossL2ProverV2Filterer, error) {
	contract, err := bindMockCrossL2ProverV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockCrossL2ProverV2Filterer{contract: contract}, nil
}

// bindMockCrossL2ProverV2 binds a generic wrapper to an already deployed contract.
func bindMockCrossL2ProverV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockCrossL2ProverV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockCrossL2ProverV2.Contract.MockCrossL2ProverV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockCrossL2ProverV2.Contract.MockCrossL2ProverV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockCrossL2ProverV2.Contract.MockCrossL2ProverV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockCrossL2ProverV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockCrossL2ProverV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockCrossL2ProverV2.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(bytes32)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Caller) CHAINID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockCrossL2ProverV2.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(bytes32)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) CHAINID() ([32]byte, error) {
	return _MockCrossL2ProverV2.Contract.CHAINID(&_MockCrossL2ProverV2.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(bytes32)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2CallerSession) CHAINID() ([32]byte, error) {
	return _MockCrossL2ProverV2.Contract.CHAINID(&_MockCrossL2ProverV2.CallOpts)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Caller) LIGHTCLIENTTYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockCrossL2ProverV2.contract.Call(opts, &out, "LIGHT_CLIENT_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) LIGHTCLIENTTYPE() (uint8, error) {
	return _MockCrossL2ProverV2.Contract.LIGHTCLIENTTYPE(&_MockCrossL2ProverV2.CallOpts)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2CallerSession) LIGHTCLIENTTYPE() (uint8, error) {
	return _MockCrossL2ProverV2.Contract.LIGHTCLIENTTYPE(&_MockCrossL2ProverV2.CallOpts)
}

// SEQUENCER is a free data retrieval call binding the contract method 0x75fd4ca9.
//
// Solidity: function SEQUENCER() view returns(address)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Caller) SEQUENCER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockCrossL2ProverV2.contract.Call(opts, &out, "SEQUENCER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SEQUENCER is a free data retrieval call binding the contract method 0x75fd4ca9.
//
// Solidity: function SEQUENCER() view returns(address)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) SEQUENCER() (common.Address, error) {
	return _MockCrossL2ProverV2.Contract.SEQUENCER(&_MockCrossL2ProverV2.CallOpts)
}

// SEQUENCER is a free data retrieval call binding the contract method 0x75fd4ca9.
//
// Solidity: function SEQUENCER() view returns(address)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2CallerSession) SEQUENCER() (common.Address, error) {
	return _MockCrossL2ProverV2.Contract.SEQUENCER(&_MockCrossL2ProverV2.CallOpts)
}

// ClientType is a free data retrieval call binding the contract method 0xb3768f0d.
//
// Solidity: function clientType() view returns(string)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Caller) ClientType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockCrossL2ProverV2.contract.Call(opts, &out, "clientType")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ClientType is a free data retrieval call binding the contract method 0xb3768f0d.
//
// Solidity: function clientType() view returns(string)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) ClientType() (string, error) {
	return _MockCrossL2ProverV2.Contract.ClientType(&_MockCrossL2ProverV2.CallOpts)
}

// ClientType is a free data retrieval call binding the contract method 0xb3768f0d.
//
// Solidity: function clientType() view returns(string)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2CallerSession) ClientType() (string, error) {
	return _MockCrossL2ProverV2.Contract.ClientType(&_MockCrossL2ProverV2.CallOpts)
}

// GenerateMockProof is a free data retrieval call binding the contract method 0x4a0dad13.
//
// Solidity: function generateMockProof(uint32 chainId_, uint8 numTopics, address emitter, bytes32[] topics_, bytes unindexedData_) pure returns(bytes)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Caller) GenerateMockProof(opts *bind.CallOpts, chainId_ uint32, numTopics uint8, emitter common.Address, topics_ [][32]byte, unindexedData_ []byte) ([]byte, error) {
	var out []interface{}
	err := _MockCrossL2ProverV2.contract.Call(opts, &out, "generateMockProof", chainId_, numTopics, emitter, topics_, unindexedData_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateMockProof is a free data retrieval call binding the contract method 0x4a0dad13.
//
// Solidity: function generateMockProof(uint32 chainId_, uint8 numTopics, address emitter, bytes32[] topics_, bytes unindexedData_) pure returns(bytes)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) GenerateMockProof(chainId_ uint32, numTopics uint8, emitter common.Address, topics_ [][32]byte, unindexedData_ []byte) ([]byte, error) {
	return _MockCrossL2ProverV2.Contract.GenerateMockProof(&_MockCrossL2ProverV2.CallOpts, chainId_, numTopics, emitter, topics_, unindexedData_)
}

// GenerateMockProof is a free data retrieval call binding the contract method 0x4a0dad13.
//
// Solidity: function generateMockProof(uint32 chainId_, uint8 numTopics, address emitter, bytes32[] topics_, bytes unindexedData_) pure returns(bytes)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2CallerSession) GenerateMockProof(chainId_ uint32, numTopics uint8, emitter common.Address, topics_ [][32]byte, unindexedData_ []byte) ([]byte, error) {
	return _MockCrossL2ProverV2.Contract.GenerateMockProof(&_MockCrossL2ProverV2.CallOpts, chainId_, numTopics, emitter, topics_, unindexedData_)
}

// InspectLogIdentifier is a free data retrieval call binding the contract method 0x9e79c3f0.
//
// Solidity: function inspectLogIdentifier(bytes proof) pure returns(uint32 srcChain, uint64 blockNumber, uint32 receiptIndex, uint32 logIndex)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Caller) InspectLogIdentifier(opts *bind.CallOpts, proof []byte) (struct {
	SrcChain     uint32
	BlockNumber  uint64
	ReceiptIndex uint32
	LogIndex     uint32
}, error) {
	var out []interface{}
	err := _MockCrossL2ProverV2.contract.Call(opts, &out, "inspectLogIdentifier", proof)

	outstruct := new(struct {
		SrcChain     uint32
		BlockNumber  uint64
		ReceiptIndex uint32
		LogIndex     uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SrcChain = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ReceiptIndex = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.LogIndex = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// InspectLogIdentifier is a free data retrieval call binding the contract method 0x9e79c3f0.
//
// Solidity: function inspectLogIdentifier(bytes proof) pure returns(uint32 srcChain, uint64 blockNumber, uint32 receiptIndex, uint32 logIndex)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) InspectLogIdentifier(proof []byte) (struct {
	SrcChain     uint32
	BlockNumber  uint64
	ReceiptIndex uint32
	LogIndex     uint32
}, error) {
	return _MockCrossL2ProverV2.Contract.InspectLogIdentifier(&_MockCrossL2ProverV2.CallOpts, proof)
}

// InspectLogIdentifier is a free data retrieval call binding the contract method 0x9e79c3f0.
//
// Solidity: function inspectLogIdentifier(bytes proof) pure returns(uint32 srcChain, uint64 blockNumber, uint32 receiptIndex, uint32 logIndex)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2CallerSession) InspectLogIdentifier(proof []byte) (struct {
	SrcChain     uint32
	BlockNumber  uint64
	ReceiptIndex uint32
	LogIndex     uint32
}, error) {
	return _MockCrossL2ProverV2.Contract.InspectLogIdentifier(&_MockCrossL2ProverV2.CallOpts, proof)
}

// InspectPolymerState is a free data retrieval call binding the contract method 0xd9615898.
//
// Solidity: function inspectPolymerState(bytes proof) pure returns(bytes32 stateRoot, uint64 height, bytes signature)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Caller) InspectPolymerState(opts *bind.CallOpts, proof []byte) (struct {
	StateRoot [32]byte
	Height    uint64
	Signature []byte
}, error) {
	var out []interface{}
	err := _MockCrossL2ProverV2.contract.Call(opts, &out, "inspectPolymerState", proof)

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
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) InspectPolymerState(proof []byte) (struct {
	StateRoot [32]byte
	Height    uint64
	Signature []byte
}, error) {
	return _MockCrossL2ProverV2.Contract.InspectPolymerState(&_MockCrossL2ProverV2.CallOpts, proof)
}

// InspectPolymerState is a free data retrieval call binding the contract method 0xd9615898.
//
// Solidity: function inspectPolymerState(bytes proof) pure returns(bytes32 stateRoot, uint64 height, bytes signature)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2CallerSession) InspectPolymerState(proof []byte) (struct {
	StateRoot [32]byte
	Height    uint64
	Signature []byte
}, error) {
	return _MockCrossL2ProverV2.Contract.InspectPolymerState(&_MockCrossL2ProverV2.CallOpts, proof)
}

// ParseEvent is a free data retrieval call binding the contract method 0x0d15d928.
//
// Solidity: function parseEvent(bytes rawEvent, uint8 numTopics) pure returns(address emittingContract, bytes topics, bytes unindexedData)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Caller) ParseEvent(opts *bind.CallOpts, rawEvent []byte, numTopics uint8) (struct {
	EmittingContract common.Address
	Topics           []byte
	UnindexedData    []byte
}, error) {
	var out []interface{}
	err := _MockCrossL2ProverV2.contract.Call(opts, &out, "parseEvent", rawEvent, numTopics)

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
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) ParseEvent(rawEvent []byte, numTopics uint8) (struct {
	EmittingContract common.Address
	Topics           []byte
	UnindexedData    []byte
}, error) {
	return _MockCrossL2ProverV2.Contract.ParseEvent(&_MockCrossL2ProverV2.CallOpts, rawEvent, numTopics)
}

// ParseEvent is a free data retrieval call binding the contract method 0x0d15d928.
//
// Solidity: function parseEvent(bytes rawEvent, uint8 numTopics) pure returns(address emittingContract, bytes topics, bytes unindexedData)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2CallerSession) ParseEvent(rawEvent []byte, numTopics uint8) (struct {
	EmittingContract common.Address
	Topics           []byte
	UnindexedData    []byte
}, error) {
	return _MockCrossL2ProverV2.Contract.ParseEvent(&_MockCrossL2ProverV2.CallOpts, rawEvent, numTopics)
}

// ValidateEvent is a free data retrieval call binding the contract method 0x0b9ca54a.
//
// Solidity: function validateEvent(bytes proof) view returns(uint32 chainId, address emittingContract, bytes topics, bytes unindexedData)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Caller) ValidateEvent(opts *bind.CallOpts, proof []byte) (struct {
	ChainId          uint32
	EmittingContract common.Address
	Topics           []byte
	UnindexedData    []byte
}, error) {
	var out []interface{}
	err := _MockCrossL2ProverV2.contract.Call(opts, &out, "validateEvent", proof)

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
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) ValidateEvent(proof []byte) (struct {
	ChainId          uint32
	EmittingContract common.Address
	Topics           []byte
	UnindexedData    []byte
}, error) {
	return _MockCrossL2ProverV2.Contract.ValidateEvent(&_MockCrossL2ProverV2.CallOpts, proof)
}

// ValidateEvent is a free data retrieval call binding the contract method 0x0b9ca54a.
//
// Solidity: function validateEvent(bytes proof) view returns(uint32 chainId, address emittingContract, bytes topics, bytes unindexedData)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2CallerSession) ValidateEvent(proof []byte) (struct {
	ChainId          uint32
	EmittingContract common.Address
	Topics           []byte
	UnindexedData    []byte
}, error) {
	return _MockCrossL2ProverV2.Contract.ValidateEvent(&_MockCrossL2ProverV2.CallOpts, proof)
}

// ValidateSolLogs is a free data retrieval call binding the contract method 0xd73a8ad6.
//
// Solidity: function validateSolLogs(bytes proof) view returns(uint32 chainId, bytes32 programID, string[] logMessages)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Caller) ValidateSolLogs(opts *bind.CallOpts, proof []byte) (struct {
	ChainId     uint32
	ProgramID   [32]byte
	LogMessages []string
}, error) {
	var out []interface{}
	err := _MockCrossL2ProverV2.contract.Call(opts, &out, "validateSolLogs", proof)

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
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) ValidateSolLogs(proof []byte) (struct {
	ChainId     uint32
	ProgramID   [32]byte
	LogMessages []string
}, error) {
	return _MockCrossL2ProverV2.Contract.ValidateSolLogs(&_MockCrossL2ProverV2.CallOpts, proof)
}

// ValidateSolLogs is a free data retrieval call binding the contract method 0xd73a8ad6.
//
// Solidity: function validateSolLogs(bytes proof) view returns(uint32 chainId, bytes32 programID, string[] logMessages)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2CallerSession) ValidateSolLogs(proof []byte) (struct {
	ChainId     uint32
	ProgramID   [32]byte
	LogMessages []string
}, error) {
	return _MockCrossL2ProverV2.Contract.ValidateSolLogs(&_MockCrossL2ProverV2.CallOpts, proof)
}

// VerifyMembership is a free data retrieval call binding the contract method 0x861f2843.
//
// Solidity: function verifyMembership(bytes32 root, bytes key, bytes32 value, bytes proof) pure returns()
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Caller) VerifyMembership(opts *bind.CallOpts, root [32]byte, key []byte, value [32]byte, proof []byte) error {
	var out []interface{}
	err := _MockCrossL2ProverV2.contract.Call(opts, &out, "verifyMembership", root, key, value, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyMembership is a free data retrieval call binding the contract method 0x861f2843.
//
// Solidity: function verifyMembership(bytes32 root, bytes key, bytes32 value, bytes proof) pure returns()
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) VerifyMembership(root [32]byte, key []byte, value [32]byte, proof []byte) error {
	return _MockCrossL2ProverV2.Contract.VerifyMembership(&_MockCrossL2ProverV2.CallOpts, root, key, value, proof)
}

// VerifyMembership is a free data retrieval call binding the contract method 0x861f2843.
//
// Solidity: function verifyMembership(bytes32 root, bytes key, bytes32 value, bytes proof) pure returns()
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2CallerSession) VerifyMembership(root [32]byte, key []byte, value [32]byte, proof []byte) error {
	return _MockCrossL2ProverV2.Contract.VerifyMembership(&_MockCrossL2ProverV2.CallOpts, root, key, value, proof)
}

// GenerateAndEmitProof is a paid mutator transaction binding the contract method 0xafc489ee.
//
// Solidity: function generateAndEmitProof(uint32 chainId_, address emitter, bytes32[] topics, bytes data) returns(bytes)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Transactor) GenerateAndEmitProof(opts *bind.TransactOpts, chainId_ uint32, emitter common.Address, topics [][32]byte, data []byte) (*types.Transaction, error) {
	return _MockCrossL2ProverV2.contract.Transact(opts, "generateAndEmitProof", chainId_, emitter, topics, data)
}

// GenerateAndEmitProof is a paid mutator transaction binding the contract method 0xafc489ee.
//
// Solidity: function generateAndEmitProof(uint32 chainId_, address emitter, bytes32[] topics, bytes data) returns(bytes)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) GenerateAndEmitProof(chainId_ uint32, emitter common.Address, topics [][32]byte, data []byte) (*types.Transaction, error) {
	return _MockCrossL2ProverV2.Contract.GenerateAndEmitProof(&_MockCrossL2ProverV2.TransactOpts, chainId_, emitter, topics, data)
}

// GenerateAndEmitProof is a paid mutator transaction binding the contract method 0xafc489ee.
//
// Solidity: function generateAndEmitProof(uint32 chainId_, address emitter, bytes32[] topics, bytes data) returns(bytes)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2TransactorSession) GenerateAndEmitProof(chainId_ uint32, emitter common.Address, topics [][32]byte, data []byte) (*types.Transaction, error) {
	return _MockCrossL2ProverV2.Contract.GenerateAndEmitProof(&_MockCrossL2ProverV2.TransactOpts, chainId_, emitter, topics, data)
}

// GenerateAndSendProof is a paid mutator transaction binding the contract method 0xc4a0f968.
//
// Solidity: function generateAndSendProof(uint32 chainId_, address emitter, bytes32[] topics, bytes data, address validatorContract) returns(bytes)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Transactor) GenerateAndSendProof(opts *bind.TransactOpts, chainId_ uint32, emitter common.Address, topics [][32]byte, data []byte, validatorContract common.Address) (*types.Transaction, error) {
	return _MockCrossL2ProverV2.contract.Transact(opts, "generateAndSendProof", chainId_, emitter, topics, data, validatorContract)
}

// GenerateAndSendProof is a paid mutator transaction binding the contract method 0xc4a0f968.
//
// Solidity: function generateAndSendProof(uint32 chainId_, address emitter, bytes32[] topics, bytes data, address validatorContract) returns(bytes)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) GenerateAndSendProof(chainId_ uint32, emitter common.Address, topics [][32]byte, data []byte, validatorContract common.Address) (*types.Transaction, error) {
	return _MockCrossL2ProverV2.Contract.GenerateAndSendProof(&_MockCrossL2ProverV2.TransactOpts, chainId_, emitter, topics, data, validatorContract)
}

// GenerateAndSendProof is a paid mutator transaction binding the contract method 0xc4a0f968.
//
// Solidity: function generateAndSendProof(uint32 chainId_, address emitter, bytes32[] topics, bytes data, address validatorContract) returns(bytes)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2TransactorSession) GenerateAndSendProof(chainId_ uint32, emitter common.Address, topics [][32]byte, data []byte, validatorContract common.Address) (*types.Transaction, error) {
	return _MockCrossL2ProverV2.Contract.GenerateAndSendProof(&_MockCrossL2ProverV2.TransactOpts, chainId_, emitter, topics, data, validatorContract)
}

// Ping is a paid mutator transaction binding the contract method 0x5c36b186.
//
// Solidity: function ping() returns()
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Transactor) Ping(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockCrossL2ProverV2.contract.Transact(opts, "ping")
}

// Ping is a paid mutator transaction binding the contract method 0x5c36b186.
//
// Solidity: function ping() returns()
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Session) Ping() (*types.Transaction, error) {
	return _MockCrossL2ProverV2.Contract.Ping(&_MockCrossL2ProverV2.TransactOpts)
}

// Ping is a paid mutator transaction binding the contract method 0x5c36b186.
//
// Solidity: function ping() returns()
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2TransactorSession) Ping() (*types.Transaction, error) {
	return _MockCrossL2ProverV2.Contract.Ping(&_MockCrossL2ProverV2.TransactOpts)
}

// MockCrossL2ProverV2PingIterator is returned from FilterPing and is used to iterate over the raw logs and unpacked data for Ping events raised by the MockCrossL2ProverV2 contract.
type MockCrossL2ProverV2PingIterator struct {
	Event *MockCrossL2ProverV2Ping // Event containing the contract specifics and raw log

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
func (it *MockCrossL2ProverV2PingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockCrossL2ProverV2Ping)
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
		it.Event = new(MockCrossL2ProverV2Ping)
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
func (it *MockCrossL2ProverV2PingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockCrossL2ProverV2PingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockCrossL2ProverV2Ping represents a Ping event raised by the MockCrossL2ProverV2 contract.
type MockCrossL2ProverV2Ping struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPing is a free log retrieval operation binding the contract event 0xca6e822df923f741dfe968d15d80a18abd25bd1e748bcb9ad81fea5bbb7386af.
//
// Solidity: event Ping()
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Filterer) FilterPing(opts *bind.FilterOpts) (*MockCrossL2ProverV2PingIterator, error) {

	logs, sub, err := _MockCrossL2ProverV2.contract.FilterLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return &MockCrossL2ProverV2PingIterator{contract: _MockCrossL2ProverV2.contract, event: "Ping", logs: logs, sub: sub}, nil
}

// WatchPing is a free log subscription operation binding the contract event 0xca6e822df923f741dfe968d15d80a18abd25bd1e748bcb9ad81fea5bbb7386af.
//
// Solidity: event Ping()
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Filterer) WatchPing(opts *bind.WatchOpts, sink chan<- *MockCrossL2ProverV2Ping) (event.Subscription, error) {

	logs, sub, err := _MockCrossL2ProverV2.contract.WatchLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockCrossL2ProverV2Ping)
				if err := _MockCrossL2ProverV2.contract.UnpackLog(event, "Ping", log); err != nil {
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

// ParsePing is a log parse operation binding the contract event 0xca6e822df923f741dfe968d15d80a18abd25bd1e748bcb9ad81fea5bbb7386af.
//
// Solidity: event Ping()
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Filterer) ParsePing(log types.Log) (*MockCrossL2ProverV2Ping, error) {
	event := new(MockCrossL2ProverV2Ping)
	if err := _MockCrossL2ProverV2.contract.UnpackLog(event, "Ping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockCrossL2ProverV2ProofGeneratedIterator is returned from FilterProofGenerated and is used to iterate over the raw logs and unpacked data for ProofGenerated events raised by the MockCrossL2ProverV2 contract.
type MockCrossL2ProverV2ProofGeneratedIterator struct {
	Event *MockCrossL2ProverV2ProofGenerated // Event containing the contract specifics and raw log

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
func (it *MockCrossL2ProverV2ProofGeneratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockCrossL2ProverV2ProofGenerated)
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
		it.Event = new(MockCrossL2ProverV2ProofGenerated)
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
func (it *MockCrossL2ProverV2ProofGeneratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockCrossL2ProverV2ProofGeneratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockCrossL2ProverV2ProofGenerated represents a ProofGenerated event raised by the MockCrossL2ProverV2 contract.
type MockCrossL2ProverV2ProofGenerated struct {
	Proof []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProofGenerated is a free log retrieval operation binding the contract event 0x4844f9b37904a6125b8c8c6df00b455fb018fac712949a4144c2f9adddb49822.
//
// Solidity: event ProofGenerated(bytes proof)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Filterer) FilterProofGenerated(opts *bind.FilterOpts) (*MockCrossL2ProverV2ProofGeneratedIterator, error) {

	logs, sub, err := _MockCrossL2ProverV2.contract.FilterLogs(opts, "ProofGenerated")
	if err != nil {
		return nil, err
	}
	return &MockCrossL2ProverV2ProofGeneratedIterator{contract: _MockCrossL2ProverV2.contract, event: "ProofGenerated", logs: logs, sub: sub}, nil
}

// WatchProofGenerated is a free log subscription operation binding the contract event 0x4844f9b37904a6125b8c8c6df00b455fb018fac712949a4144c2f9adddb49822.
//
// Solidity: event ProofGenerated(bytes proof)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Filterer) WatchProofGenerated(opts *bind.WatchOpts, sink chan<- *MockCrossL2ProverV2ProofGenerated) (event.Subscription, error) {

	logs, sub, err := _MockCrossL2ProverV2.contract.WatchLogs(opts, "ProofGenerated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockCrossL2ProverV2ProofGenerated)
				if err := _MockCrossL2ProverV2.contract.UnpackLog(event, "ProofGenerated", log); err != nil {
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

// ParseProofGenerated is a log parse operation binding the contract event 0x4844f9b37904a6125b8c8c6df00b455fb018fac712949a4144c2f9adddb49822.
//
// Solidity: event ProofGenerated(bytes proof)
func (_MockCrossL2ProverV2 *MockCrossL2ProverV2Filterer) ParseProofGenerated(log types.Log) (*MockCrossL2ProverV2ProofGenerated, error) {
	event := new(MockCrossL2ProverV2ProofGenerated)
	if err := _MockCrossL2ProverV2.contract.UnpackLog(event, "ProofGenerated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
