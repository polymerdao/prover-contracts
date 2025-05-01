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
	L2Type               uint8
}

// NativeProverInitialL2Configuration is an auto generated low-level Go binding around an user-defined struct.
type NativeProverInitialL2Configuration struct {
	ChainID *big.Int
	Config  L2Configuration
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1ChainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_initialL2Configurations\",\"type\":\"tuple[]\",\"internalType\":\"structNativeProver.InitialL2Configuration[]\",\"components\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"enumType\"}]}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"L1_CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"L1_CONFIGURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementBlocksDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configureAndProve\",\"inputs\":[{\"name\":\"_updateArgs\",\"type\":\"tuple\",\"internalType\":\"structUpdateL2ConfigArgs\",\"components\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"enumType\"}]},{\"name\":\"l1StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"rlpEncodedRegistryAccountData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"l1RegistryProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"name\":\"_proveArgs\",\"type\":\"tuple\",\"internalType\":\"structProveScalarArgs\",\"components\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_rlpEncodedL1Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_rlpEncodedL2Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_settledStateProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l2StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_rlpEncodedContractAccount\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l2AccountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2ChainConfigurations\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"enumType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"packGameID\",\"inputs\":[{\"name\":\"_gameType\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_gameProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"gameId_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"proveL1Native\",\"inputs\":[{\"name\":\"_args\",\"type\":\"tuple\",\"internalType\":\"structProveL1ScalarArgs\",\"components\":[{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l1WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_rlpEncodedL1Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l1StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_rlpEncodedContractAccount\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l1AccountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveNative\",\"inputs\":[{\"name\":\"_args\",\"type\":\"tuple\",\"internalType\":\"structProveScalarArgs\",\"components\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_rlpEncodedL1Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_rlpEncodedL2Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_settledStateProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l2StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_rlpEncodedContractAccount\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l2AccountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveSettledState\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_rlpEncodedL2Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l1WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proveSettlementLayerState\",\"inputs\":[{\"name\":\"_rlpEncodedL1Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proveStorage\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_val\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"proveStorageValue\",\"inputs\":[{\"name\":\"_args\",\"type\":\"tuple\",\"internalType\":\"structProveScalarArgs\",\"components\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_l2StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_rlpEncodedContractState\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l2AccountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provenStates\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rlpEncodeDataLibList\",\"inputs\":[{\"name\":\"_dataList\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"setInitialL1Config\",\"inputs\":[{\"name\":\"_l1Configuration\",\"type\":\"tuple\",\"internalType\":\"structL1Configuration\",\"components\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementBlocksDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stateProvers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISettledStateProver\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAndProve\",\"inputs\":[{\"name\":\"_updateArgs\",\"type\":\"tuple\",\"internalType\":\"structUpdateL2ConfigArgs\",\"components\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"enumType\"}]},{\"name\":\"l1StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"rlpEncodedRegistryAccountData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"l1RegistryProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"name\":\"_proveArgs\",\"type\":\"tuple\",\"internalType\":\"structProveScalarArgs\",\"components\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_rlpEncodedL1Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_rlpEncodedL2Header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_settledStateProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l2StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_rlpEncodedContractAccount\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l2AccountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"storingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storageValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateL1ChainConfiguration\",\"inputs\":[{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structL1Configuration\",\"components\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementBlocksDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_l1StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_rlpEncodedRegistryAccountData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l1RegistryProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_l1WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateL2ChainConfiguration\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"enumType\"}]},{\"name\":\"_l1StorageProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_rlpEncodedRegistryAccountData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l1RegistryProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_l1WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"L1WorldStateProven\",\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"_L1WorldStateRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L2WorldStateProven\",\"inputs\":[{\"name\":\"_destinationChainID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"_L2WorldStateRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DestinationChainStateRootNotProved\",\"inputs\":[{\"name\":\"_blockProofStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"IncorrectContractStorageRoot\",\"inputs\":[{\"name\":\"_contractStorageRoot\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidAccountProof\",\"inputs\":[{\"name\":\"_address\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidL1ConfigurationProof\",\"inputs\":[{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structL1Configuration\",\"components\":[{\"name\":\"blockHashOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementBlocksDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"settlementRegistryL2ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settlementRegistryL1ConfigMappingSlot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"type\":\"error\",\"name\":\"InvalidL2ConfigurationProof\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structL2Configuration\",\"components\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"storageSlots\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"finalityDelaySeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Type\",\"type\":\"uint8\",\"internalType\":\"enumType\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedBlock\",\"inputs\":[{\"name\":\"_expectedBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_calculatedBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidSettledStateProof\",\"inputs\":[{\"name\":\"_chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidStorageProof\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_val\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidStorageProof\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_val\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"NeedLaterBlock\",\"inputs\":[{\"name\":\"_inputBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nextProvableBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OutdatedBlock\",\"inputs\":[{\"name\":\"_inputBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_latestBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SettlementChainStateRootNotProven\",\"inputs\":[{\"name\":\"_blockProofStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_l1WorldStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
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
// Solidity: function L1_CONFIGURATION() view returns(address blockHashOracle, uint256 settlementBlocksDelay, address settlementRegistry, uint256 settlementRegistryL2ConfigMappingSlot, uint256 settlementRegistryL1ConfigMappingSlot)
func (_NativeProver *NativeProverCaller) L1CONFIGURATION(opts *bind.CallOpts) (struct {
	BlockHashOracle                       common.Address
	SettlementBlocksDelay                 *big.Int
	SettlementRegistry                    common.Address
	SettlementRegistryL2ConfigMappingSlot *big.Int
	SettlementRegistryL1ConfigMappingSlot *big.Int
}, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "L1_CONFIGURATION")

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

// L1CONFIGURATION is a free data retrieval call binding the contract method 0xbff89934.
//
// Solidity: function L1_CONFIGURATION() view returns(address blockHashOracle, uint256 settlementBlocksDelay, address settlementRegistry, uint256 settlementRegistryL2ConfigMappingSlot, uint256 settlementRegistryL1ConfigMappingSlot)
func (_NativeProver *NativeProverSession) L1CONFIGURATION() (struct {
	BlockHashOracle                       common.Address
	SettlementBlocksDelay                 *big.Int
	SettlementRegistry                    common.Address
	SettlementRegistryL2ConfigMappingSlot *big.Int
	SettlementRegistryL1ConfigMappingSlot *big.Int
}, error) {
	return _NativeProver.Contract.L1CONFIGURATION(&_NativeProver.CallOpts)
}

// L1CONFIGURATION is a free data retrieval call binding the contract method 0xbff89934.
//
// Solidity: function L1_CONFIGURATION() view returns(address blockHashOracle, uint256 settlementBlocksDelay, address settlementRegistry, uint256 settlementRegistryL2ConfigMappingSlot, uint256 settlementRegistryL1ConfigMappingSlot)
func (_NativeProver *NativeProverCallerSession) L1CONFIGURATION() (struct {
	BlockHashOracle                       common.Address
	SettlementBlocksDelay                 *big.Int
	SettlementRegistry                    common.Address
	SettlementRegistryL2ConfigMappingSlot *big.Int
	SettlementRegistryL1ConfigMappingSlot *big.Int
}, error) {
	return _NativeProver.Contract.L1CONFIGURATION(&_NativeProver.CallOpts)
}

// ConfigureAndProve is a free data retrieval call binding the contract method 0x3c873bb2.
//
// Solidity: function configureAndProve(((address,address[],uint256[],uint256,uint256,uint8),bytes[],bytes,bytes[]) _updateArgs, (uint256,address,bytes32,bytes32,bytes32) _proveArgs, bytes _rlpEncodedL1Header, bytes _rlpEncodedL2Header, bytes _settledStateProof, bytes[] _l2StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l2AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverCaller) ConfigureAndProve(opts *bind.CallOpts, _updateArgs UpdateL2ConfigArgs, _proveArgs ProveScalarArgs, _rlpEncodedL1Header []byte, _rlpEncodedL2Header []byte, _settledStateProof []byte, _l2StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l2AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "configureAndProve", _updateArgs, _proveArgs, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof)

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

// ConfigureAndProve is a free data retrieval call binding the contract method 0x3c873bb2.
//
// Solidity: function configureAndProve(((address,address[],uint256[],uint256,uint256,uint8),bytes[],bytes,bytes[]) _updateArgs, (uint256,address,bytes32,bytes32,bytes32) _proveArgs, bytes _rlpEncodedL1Header, bytes _rlpEncodedL2Header, bytes _settledStateProof, bytes[] _l2StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l2AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverSession) ConfigureAndProve(_updateArgs UpdateL2ConfigArgs, _proveArgs ProveScalarArgs, _rlpEncodedL1Header []byte, _rlpEncodedL2Header []byte, _settledStateProof []byte, _l2StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l2AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	return _NativeProver.Contract.ConfigureAndProve(&_NativeProver.CallOpts, _updateArgs, _proveArgs, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof)
}

// ConfigureAndProve is a free data retrieval call binding the contract method 0x3c873bb2.
//
// Solidity: function configureAndProve(((address,address[],uint256[],uint256,uint256,uint8),bytes[],bytes,bytes[]) _updateArgs, (uint256,address,bytes32,bytes32,bytes32) _proveArgs, bytes _rlpEncodedL1Header, bytes _rlpEncodedL2Header, bytes _settledStateProof, bytes[] _l2StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l2AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverCallerSession) ConfigureAndProve(_updateArgs UpdateL2ConfigArgs, _proveArgs ProveScalarArgs, _rlpEncodedL1Header []byte, _rlpEncodedL2Header []byte, _settledStateProof []byte, _l2StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l2AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	return _NativeProver.Contract.ConfigureAndProve(&_NativeProver.CallOpts, _updateArgs, _proveArgs, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof)
}

// L2ChainConfigurations is a free data retrieval call binding the contract method 0x63281a15.
//
// Solidity: function l2ChainConfigurations(uint256 ) view returns(address prover, uint256 versionNumber, uint256 finalityDelaySeconds, uint8 l2Type)
func (_NativeProver *NativeProverCaller) L2ChainConfigurations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Prover               common.Address
	VersionNumber        *big.Int
	FinalityDelaySeconds *big.Int
	L2Type               uint8
}, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "l2ChainConfigurations", arg0)

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
func (_NativeProver *NativeProverSession) L2ChainConfigurations(arg0 *big.Int) (struct {
	Prover               common.Address
	VersionNumber        *big.Int
	FinalityDelaySeconds *big.Int
	L2Type               uint8
}, error) {
	return _NativeProver.Contract.L2ChainConfigurations(&_NativeProver.CallOpts, arg0)
}

// L2ChainConfigurations is a free data retrieval call binding the contract method 0x63281a15.
//
// Solidity: function l2ChainConfigurations(uint256 ) view returns(address prover, uint256 versionNumber, uint256 finalityDelaySeconds, uint8 l2Type)
func (_NativeProver *NativeProverCallerSession) L2ChainConfigurations(arg0 *big.Int) (struct {
	Prover               common.Address
	VersionNumber        *big.Int
	FinalityDelaySeconds *big.Int
	L2Type               uint8
}, error) {
	return _NativeProver.Contract.L2ChainConfigurations(&_NativeProver.CallOpts, arg0)
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

// PackGameID is a free data retrieval call binding the contract method 0x90c12ec4.
//
// Solidity: function packGameID(uint32 _gameType, uint64 _timestamp, address _gameProxy) pure returns(bytes32 gameId_)
func (_NativeProver *NativeProverCaller) PackGameID(opts *bind.CallOpts, _gameType uint32, _timestamp uint64, _gameProxy common.Address) ([32]byte, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "packGameID", _gameType, _timestamp, _gameProxy)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PackGameID is a free data retrieval call binding the contract method 0x90c12ec4.
//
// Solidity: function packGameID(uint32 _gameType, uint64 _timestamp, address _gameProxy) pure returns(bytes32 gameId_)
func (_NativeProver *NativeProverSession) PackGameID(_gameType uint32, _timestamp uint64, _gameProxy common.Address) ([32]byte, error) {
	return _NativeProver.Contract.PackGameID(&_NativeProver.CallOpts, _gameType, _timestamp, _gameProxy)
}

// PackGameID is a free data retrieval call binding the contract method 0x90c12ec4.
//
// Solidity: function packGameID(uint32 _gameType, uint64 _timestamp, address _gameProxy) pure returns(bytes32 gameId_)
func (_NativeProver *NativeProverCallerSession) PackGameID(_gameType uint32, _timestamp uint64, _gameProxy common.Address) ([32]byte, error) {
	return _NativeProver.Contract.PackGameID(&_NativeProver.CallOpts, _gameType, _timestamp, _gameProxy)
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

// ProveNative is a free data retrieval call binding the contract method 0x37608649.
//
// Solidity: function proveNative((uint256,address,bytes32,bytes32,bytes32) _args, bytes _rlpEncodedL1Header, bytes _rlpEncodedL2Header, bytes _settledStateProof, bytes[] _l2StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l2AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverCaller) ProveNative(opts *bind.CallOpts, _args ProveScalarArgs, _rlpEncodedL1Header []byte, _rlpEncodedL2Header []byte, _settledStateProof []byte, _l2StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l2AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "proveNative", _args, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof)

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

// ProveNative is a free data retrieval call binding the contract method 0x37608649.
//
// Solidity: function proveNative((uint256,address,bytes32,bytes32,bytes32) _args, bytes _rlpEncodedL1Header, bytes _rlpEncodedL2Header, bytes _settledStateProof, bytes[] _l2StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l2AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverSession) ProveNative(_args ProveScalarArgs, _rlpEncodedL1Header []byte, _rlpEncodedL2Header []byte, _settledStateProof []byte, _l2StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l2AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	return _NativeProver.Contract.ProveNative(&_NativeProver.CallOpts, _args, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof)
}

// ProveNative is a free data retrieval call binding the contract method 0x37608649.
//
// Solidity: function proveNative((uint256,address,bytes32,bytes32,bytes32) _args, bytes _rlpEncodedL1Header, bytes _rlpEncodedL2Header, bytes _settledStateProof, bytes[] _l2StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l2AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverCallerSession) ProveNative(_args ProveScalarArgs, _rlpEncodedL1Header []byte, _rlpEncodedL2Header []byte, _settledStateProof []byte, _l2StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l2AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	return _NativeProver.Contract.ProveNative(&_NativeProver.CallOpts, _args, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof)
}

// ProveStorage is a free data retrieval call binding the contract method 0x848e9abb.
//
// Solidity: function proveStorage(bytes _key, bytes _val, bytes[] _proof, bytes32 _root) pure returns()
func (_NativeProver *NativeProverCaller) ProveStorage(opts *bind.CallOpts, _key []byte, _val []byte, _proof [][]byte, _root [32]byte) error {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "proveStorage", _key, _val, _proof, _root)

	if err != nil {
		return err
	}

	return err

}

// ProveStorage is a free data retrieval call binding the contract method 0x848e9abb.
//
// Solidity: function proveStorage(bytes _key, bytes _val, bytes[] _proof, bytes32 _root) pure returns()
func (_NativeProver *NativeProverSession) ProveStorage(_key []byte, _val []byte, _proof [][]byte, _root [32]byte) error {
	return _NativeProver.Contract.ProveStorage(&_NativeProver.CallOpts, _key, _val, _proof, _root)
}

// ProveStorage is a free data retrieval call binding the contract method 0x848e9abb.
//
// Solidity: function proveStorage(bytes _key, bytes _val, bytes[] _proof, bytes32 _root) pure returns()
func (_NativeProver *NativeProverCallerSession) ProveStorage(_key []byte, _val []byte, _proof [][]byte, _root [32]byte) error {
	return _NativeProver.Contract.ProveStorage(&_NativeProver.CallOpts, _key, _val, _proof, _root)
}

// ProveStorageValue is a free data retrieval call binding the contract method 0x7db46d26.
//
// Solidity: function proveStorageValue((uint256,address,bytes32,bytes32,bytes32) _args, bytes[] _l2StorageProof, bytes _rlpEncodedContractState, bytes[] _l2AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverCaller) ProveStorageValue(opts *bind.CallOpts, _args ProveScalarArgs, _l2StorageProof [][]byte, _rlpEncodedContractState []byte, _l2AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "proveStorageValue", _args, _l2StorageProof, _rlpEncodedContractState, _l2AccountProof)

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

// ProveStorageValue is a free data retrieval call binding the contract method 0x7db46d26.
//
// Solidity: function proveStorageValue((uint256,address,bytes32,bytes32,bytes32) _args, bytes[] _l2StorageProof, bytes _rlpEncodedContractState, bytes[] _l2AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverSession) ProveStorageValue(_args ProveScalarArgs, _l2StorageProof [][]byte, _rlpEncodedContractState []byte, _l2AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	return _NativeProver.Contract.ProveStorageValue(&_NativeProver.CallOpts, _args, _l2StorageProof, _rlpEncodedContractState, _l2AccountProof)
}

// ProveStorageValue is a free data retrieval call binding the contract method 0x7db46d26.
//
// Solidity: function proveStorageValue((uint256,address,bytes32,bytes32,bytes32) _args, bytes[] _l2StorageProof, bytes _rlpEncodedContractState, bytes[] _l2AccountProof) view returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverCallerSession) ProveStorageValue(_args ProveScalarArgs, _l2StorageProof [][]byte, _rlpEncodedContractState []byte, _l2AccountProof [][]byte) (struct {
	ChainID         *big.Int
	StoringContract common.Address
	StorageSlot     [32]byte
	StorageValue    [32]byte
}, error) {
	return _NativeProver.Contract.ProveStorageValue(&_NativeProver.CallOpts, _args, _l2StorageProof, _rlpEncodedContractState, _l2AccountProof)
}

// ProvenStates is a free data retrieval call binding the contract method 0x3bb79618.
//
// Solidity: function provenStates(uint256 ) view returns(uint256 blockNumber, bytes32 blockHash, bytes32 stateRoot)
func (_NativeProver *NativeProverCaller) ProvenStates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	StateRoot   [32]byte
}, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "provenStates", arg0)

	outstruct := new(struct {
		BlockNumber *big.Int
		BlockHash   [32]byte
		StateRoot   [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.StateRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// ProvenStates is a free data retrieval call binding the contract method 0x3bb79618.
//
// Solidity: function provenStates(uint256 ) view returns(uint256 blockNumber, bytes32 blockHash, bytes32 stateRoot)
func (_NativeProver *NativeProverSession) ProvenStates(arg0 *big.Int) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	StateRoot   [32]byte
}, error) {
	return _NativeProver.Contract.ProvenStates(&_NativeProver.CallOpts, arg0)
}

// ProvenStates is a free data retrieval call binding the contract method 0x3bb79618.
//
// Solidity: function provenStates(uint256 ) view returns(uint256 blockNumber, bytes32 blockHash, bytes32 stateRoot)
func (_NativeProver *NativeProverCallerSession) ProvenStates(arg0 *big.Int) (struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	StateRoot   [32]byte
}, error) {
	return _NativeProver.Contract.ProvenStates(&_NativeProver.CallOpts, arg0)
}

// RlpEncodeDataLibList is a free data retrieval call binding the contract method 0xb9d8ac26.
//
// Solidity: function rlpEncodeDataLibList(bytes[] _dataList) pure returns(bytes)
func (_NativeProver *NativeProverCaller) RlpEncodeDataLibList(opts *bind.CallOpts, _dataList [][]byte) ([]byte, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "rlpEncodeDataLibList", _dataList)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// RlpEncodeDataLibList is a free data retrieval call binding the contract method 0xb9d8ac26.
//
// Solidity: function rlpEncodeDataLibList(bytes[] _dataList) pure returns(bytes)
func (_NativeProver *NativeProverSession) RlpEncodeDataLibList(_dataList [][]byte) ([]byte, error) {
	return _NativeProver.Contract.RlpEncodeDataLibList(&_NativeProver.CallOpts, _dataList)
}

// RlpEncodeDataLibList is a free data retrieval call binding the contract method 0xb9d8ac26.
//
// Solidity: function rlpEncodeDataLibList(bytes[] _dataList) pure returns(bytes)
func (_NativeProver *NativeProverCallerSession) RlpEncodeDataLibList(_dataList [][]byte) ([]byte, error) {
	return _NativeProver.Contract.RlpEncodeDataLibList(&_NativeProver.CallOpts, _dataList)
}

// StateProvers is a free data retrieval call binding the contract method 0x1b4d8105.
//
// Solidity: function stateProvers(uint256 ) view returns(address)
func (_NativeProver *NativeProverCaller) StateProvers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NativeProver.contract.Call(opts, &out, "stateProvers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateProvers is a free data retrieval call binding the contract method 0x1b4d8105.
//
// Solidity: function stateProvers(uint256 ) view returns(address)
func (_NativeProver *NativeProverSession) StateProvers(arg0 *big.Int) (common.Address, error) {
	return _NativeProver.Contract.StateProvers(&_NativeProver.CallOpts, arg0)
}

// StateProvers is a free data retrieval call binding the contract method 0x1b4d8105.
//
// Solidity: function stateProvers(uint256 ) view returns(address)
func (_NativeProver *NativeProverCallerSession) StateProvers(arg0 *big.Int) (common.Address, error) {
	return _NativeProver.Contract.StateProvers(&_NativeProver.CallOpts, arg0)
}

// ProveSettledState is a paid mutator transaction binding the contract method 0xa7c69e2d.
//
// Solidity: function proveSettledState(uint256 _chainID, bytes32 _l2WorldStateRoot, bytes _rlpEncodedL2Header, bytes32 _l1WorldStateRoot, bytes _proof) returns()
func (_NativeProver *NativeProverTransactor) ProveSettledState(opts *bind.TransactOpts, _chainID *big.Int, _l2WorldStateRoot [32]byte, _rlpEncodedL2Header []byte, _l1WorldStateRoot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _NativeProver.contract.Transact(opts, "proveSettledState", _chainID, _l2WorldStateRoot, _rlpEncodedL2Header, _l1WorldStateRoot, _proof)
}

// ProveSettledState is a paid mutator transaction binding the contract method 0xa7c69e2d.
//
// Solidity: function proveSettledState(uint256 _chainID, bytes32 _l2WorldStateRoot, bytes _rlpEncodedL2Header, bytes32 _l1WorldStateRoot, bytes _proof) returns()
func (_NativeProver *NativeProverSession) ProveSettledState(_chainID *big.Int, _l2WorldStateRoot [32]byte, _rlpEncodedL2Header []byte, _l1WorldStateRoot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _NativeProver.Contract.ProveSettledState(&_NativeProver.TransactOpts, _chainID, _l2WorldStateRoot, _rlpEncodedL2Header, _l1WorldStateRoot, _proof)
}

// ProveSettledState is a paid mutator transaction binding the contract method 0xa7c69e2d.
//
// Solidity: function proveSettledState(uint256 _chainID, bytes32 _l2WorldStateRoot, bytes _rlpEncodedL2Header, bytes32 _l1WorldStateRoot, bytes _proof) returns()
func (_NativeProver *NativeProverTransactorSession) ProveSettledState(_chainID *big.Int, _l2WorldStateRoot [32]byte, _rlpEncodedL2Header []byte, _l1WorldStateRoot [32]byte, _proof []byte) (*types.Transaction, error) {
	return _NativeProver.Contract.ProveSettledState(&_NativeProver.TransactOpts, _chainID, _l2WorldStateRoot, _rlpEncodedL2Header, _l1WorldStateRoot, _proof)
}

// ProveSettlementLayerState is a paid mutator transaction binding the contract method 0x90cba459.
//
// Solidity: function proveSettlementLayerState(bytes _rlpEncodedL1Header) returns()
func (_NativeProver *NativeProverTransactor) ProveSettlementLayerState(opts *bind.TransactOpts, _rlpEncodedL1Header []byte) (*types.Transaction, error) {
	return _NativeProver.contract.Transact(opts, "proveSettlementLayerState", _rlpEncodedL1Header)
}

// ProveSettlementLayerState is a paid mutator transaction binding the contract method 0x90cba459.
//
// Solidity: function proveSettlementLayerState(bytes _rlpEncodedL1Header) returns()
func (_NativeProver *NativeProverSession) ProveSettlementLayerState(_rlpEncodedL1Header []byte) (*types.Transaction, error) {
	return _NativeProver.Contract.ProveSettlementLayerState(&_NativeProver.TransactOpts, _rlpEncodedL1Header)
}

// ProveSettlementLayerState is a paid mutator transaction binding the contract method 0x90cba459.
//
// Solidity: function proveSettlementLayerState(bytes _rlpEncodedL1Header) returns()
func (_NativeProver *NativeProverTransactorSession) ProveSettlementLayerState(_rlpEncodedL1Header []byte) (*types.Transaction, error) {
	return _NativeProver.Contract.ProveSettlementLayerState(&_NativeProver.TransactOpts, _rlpEncodedL1Header)
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

// SetInitialL1Config is a paid mutator transaction binding the contract method 0xe820de97.
//
// Solidity: function setInitialL1Config((address,uint256,address,uint256,uint256) _l1Configuration) returns()
func (_NativeProver *NativeProverTransactor) SetInitialL1Config(opts *bind.TransactOpts, _l1Configuration L1Configuration) (*types.Transaction, error) {
	return _NativeProver.contract.Transact(opts, "setInitialL1Config", _l1Configuration)
}

// SetInitialL1Config is a paid mutator transaction binding the contract method 0xe820de97.
//
// Solidity: function setInitialL1Config((address,uint256,address,uint256,uint256) _l1Configuration) returns()
func (_NativeProver *NativeProverSession) SetInitialL1Config(_l1Configuration L1Configuration) (*types.Transaction, error) {
	return _NativeProver.Contract.SetInitialL1Config(&_NativeProver.TransactOpts, _l1Configuration)
}

// SetInitialL1Config is a paid mutator transaction binding the contract method 0xe820de97.
//
// Solidity: function setInitialL1Config((address,uint256,address,uint256,uint256) _l1Configuration) returns()
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

// UpdateAndProve is a paid mutator transaction binding the contract method 0xc1ed98af.
//
// Solidity: function updateAndProve(((address,address[],uint256[],uint256,uint256,uint8),bytes[],bytes,bytes[]) _updateArgs, (uint256,address,bytes32,bytes32,bytes32) _proveArgs, bytes _rlpEncodedL1Header, bytes _rlpEncodedL2Header, bytes _settledStateProof, bytes[] _l2StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l2AccountProof) returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverTransactor) UpdateAndProve(opts *bind.TransactOpts, _updateArgs UpdateL2ConfigArgs, _proveArgs ProveScalarArgs, _rlpEncodedL1Header []byte, _rlpEncodedL2Header []byte, _settledStateProof []byte, _l2StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l2AccountProof [][]byte) (*types.Transaction, error) {
	return _NativeProver.contract.Transact(opts, "updateAndProve", _updateArgs, _proveArgs, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof)
}

// UpdateAndProve is a paid mutator transaction binding the contract method 0xc1ed98af.
//
// Solidity: function updateAndProve(((address,address[],uint256[],uint256,uint256,uint8),bytes[],bytes,bytes[]) _updateArgs, (uint256,address,bytes32,bytes32,bytes32) _proveArgs, bytes _rlpEncodedL1Header, bytes _rlpEncodedL2Header, bytes _settledStateProof, bytes[] _l2StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l2AccountProof) returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverSession) UpdateAndProve(_updateArgs UpdateL2ConfigArgs, _proveArgs ProveScalarArgs, _rlpEncodedL1Header []byte, _rlpEncodedL2Header []byte, _settledStateProof []byte, _l2StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l2AccountProof [][]byte) (*types.Transaction, error) {
	return _NativeProver.Contract.UpdateAndProve(&_NativeProver.TransactOpts, _updateArgs, _proveArgs, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof)
}

// UpdateAndProve is a paid mutator transaction binding the contract method 0xc1ed98af.
//
// Solidity: function updateAndProve(((address,address[],uint256[],uint256,uint256,uint8),bytes[],bytes,bytes[]) _updateArgs, (uint256,address,bytes32,bytes32,bytes32) _proveArgs, bytes _rlpEncodedL1Header, bytes _rlpEncodedL2Header, bytes _settledStateProof, bytes[] _l2StorageProof, bytes _rlpEncodedContractAccount, bytes[] _l2AccountProof) returns(uint256 chainID, address storingContract, bytes32 storageSlot, bytes32 storageValue)
func (_NativeProver *NativeProverTransactorSession) UpdateAndProve(_updateArgs UpdateL2ConfigArgs, _proveArgs ProveScalarArgs, _rlpEncodedL1Header []byte, _rlpEncodedL2Header []byte, _settledStateProof []byte, _l2StorageProof [][]byte, _rlpEncodedContractAccount []byte, _l2AccountProof [][]byte) (*types.Transaction, error) {
	return _NativeProver.Contract.UpdateAndProve(&_NativeProver.TransactOpts, _updateArgs, _proveArgs, _rlpEncodedL1Header, _rlpEncodedL2Header, _settledStateProof, _l2StorageProof, _rlpEncodedContractAccount, _l2AccountProof)
}

// UpdateL1ChainConfiguration is a paid mutator transaction binding the contract method 0xdbcd2a5b.
//
// Solidity: function updateL1ChainConfiguration((address,uint256,address,uint256,uint256) _config, bytes[] _l1StorageProof, bytes _rlpEncodedRegistryAccountData, bytes[] _l1RegistryProof, bytes32 _l1WorldStateRoot) returns()
func (_NativeProver *NativeProverTransactor) UpdateL1ChainConfiguration(opts *bind.TransactOpts, _config L1Configuration, _l1StorageProof [][]byte, _rlpEncodedRegistryAccountData []byte, _l1RegistryProof [][]byte, _l1WorldStateRoot [32]byte) (*types.Transaction, error) {
	return _NativeProver.contract.Transact(opts, "updateL1ChainConfiguration", _config, _l1StorageProof, _rlpEncodedRegistryAccountData, _l1RegistryProof, _l1WorldStateRoot)
}

// UpdateL1ChainConfiguration is a paid mutator transaction binding the contract method 0xdbcd2a5b.
//
// Solidity: function updateL1ChainConfiguration((address,uint256,address,uint256,uint256) _config, bytes[] _l1StorageProof, bytes _rlpEncodedRegistryAccountData, bytes[] _l1RegistryProof, bytes32 _l1WorldStateRoot) returns()
func (_NativeProver *NativeProverSession) UpdateL1ChainConfiguration(_config L1Configuration, _l1StorageProof [][]byte, _rlpEncodedRegistryAccountData []byte, _l1RegistryProof [][]byte, _l1WorldStateRoot [32]byte) (*types.Transaction, error) {
	return _NativeProver.Contract.UpdateL1ChainConfiguration(&_NativeProver.TransactOpts, _config, _l1StorageProof, _rlpEncodedRegistryAccountData, _l1RegistryProof, _l1WorldStateRoot)
}

// UpdateL1ChainConfiguration is a paid mutator transaction binding the contract method 0xdbcd2a5b.
//
// Solidity: function updateL1ChainConfiguration((address,uint256,address,uint256,uint256) _config, bytes[] _l1StorageProof, bytes _rlpEncodedRegistryAccountData, bytes[] _l1RegistryProof, bytes32 _l1WorldStateRoot) returns()
func (_NativeProver *NativeProverTransactorSession) UpdateL1ChainConfiguration(_config L1Configuration, _l1StorageProof [][]byte, _rlpEncodedRegistryAccountData []byte, _l1RegistryProof [][]byte, _l1WorldStateRoot [32]byte) (*types.Transaction, error) {
	return _NativeProver.Contract.UpdateL1ChainConfiguration(&_NativeProver.TransactOpts, _config, _l1StorageProof, _rlpEncodedRegistryAccountData, _l1RegistryProof, _l1WorldStateRoot)
}

// UpdateL2ChainConfiguration is a paid mutator transaction binding the contract method 0x2ac41ed1.
//
// Solidity: function updateL2ChainConfiguration(uint256 _chainID, (address,address[],uint256[],uint256,uint256,uint8) _config, bytes[] _l1StorageProof, bytes _rlpEncodedRegistryAccountData, bytes[] _l1RegistryProof, bytes32 _l1WorldStateRoot) returns()
func (_NativeProver *NativeProverTransactor) UpdateL2ChainConfiguration(opts *bind.TransactOpts, _chainID *big.Int, _config L2Configuration, _l1StorageProof [][]byte, _rlpEncodedRegistryAccountData []byte, _l1RegistryProof [][]byte, _l1WorldStateRoot [32]byte) (*types.Transaction, error) {
	return _NativeProver.contract.Transact(opts, "updateL2ChainConfiguration", _chainID, _config, _l1StorageProof, _rlpEncodedRegistryAccountData, _l1RegistryProof, _l1WorldStateRoot)
}

// UpdateL2ChainConfiguration is a paid mutator transaction binding the contract method 0x2ac41ed1.
//
// Solidity: function updateL2ChainConfiguration(uint256 _chainID, (address,address[],uint256[],uint256,uint256,uint8) _config, bytes[] _l1StorageProof, bytes _rlpEncodedRegistryAccountData, bytes[] _l1RegistryProof, bytes32 _l1WorldStateRoot) returns()
func (_NativeProver *NativeProverSession) UpdateL2ChainConfiguration(_chainID *big.Int, _config L2Configuration, _l1StorageProof [][]byte, _rlpEncodedRegistryAccountData []byte, _l1RegistryProof [][]byte, _l1WorldStateRoot [32]byte) (*types.Transaction, error) {
	return _NativeProver.Contract.UpdateL2ChainConfiguration(&_NativeProver.TransactOpts, _chainID, _config, _l1StorageProof, _rlpEncodedRegistryAccountData, _l1RegistryProof, _l1WorldStateRoot)
}

// UpdateL2ChainConfiguration is a paid mutator transaction binding the contract method 0x2ac41ed1.
//
// Solidity: function updateL2ChainConfiguration(uint256 _chainID, (address,address[],uint256[],uint256,uint256,uint8) _config, bytes[] _l1StorageProof, bytes _rlpEncodedRegistryAccountData, bytes[] _l1RegistryProof, bytes32 _l1WorldStateRoot) returns()
func (_NativeProver *NativeProverTransactorSession) UpdateL2ChainConfiguration(_chainID *big.Int, _config L2Configuration, _l1StorageProof [][]byte, _rlpEncodedRegistryAccountData []byte, _l1RegistryProof [][]byte, _l1WorldStateRoot [32]byte) (*types.Transaction, error) {
	return _NativeProver.Contract.UpdateL2ChainConfiguration(&_NativeProver.TransactOpts, _chainID, _config, _l1StorageProof, _rlpEncodedRegistryAccountData, _l1RegistryProof, _l1WorldStateRoot)
}

// NativeProverL1WorldStateProvenIterator is returned from FilterL1WorldStateProven and is used to iterate over the raw logs and unpacked data for L1WorldStateProven events raised by the NativeProver contract.
type NativeProverL1WorldStateProvenIterator struct {
	Event *NativeProverL1WorldStateProven // Event containing the contract specifics and raw log

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
func (it *NativeProverL1WorldStateProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeProverL1WorldStateProven)
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
		it.Event = new(NativeProverL1WorldStateProven)
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
func (it *NativeProverL1WorldStateProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeProverL1WorldStateProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeProverL1WorldStateProven represents a L1WorldStateProven event raised by the NativeProver contract.
type NativeProverL1WorldStateProven struct {
	BlockNumber      *big.Int
	L1WorldStateRoot [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterL1WorldStateProven is a free log retrieval operation binding the contract event 0xe7c32f66db8f545c384c1cc1ab188c2cc54fb6c2207477d3b141104ce9529dce.
//
// Solidity: event L1WorldStateProven(uint256 indexed _blockNumber, bytes32 _L1WorldStateRoot)
func (_NativeProver *NativeProverFilterer) FilterL1WorldStateProven(opts *bind.FilterOpts, _blockNumber []*big.Int) (*NativeProverL1WorldStateProvenIterator, error) {

	var _blockNumberRule []interface{}
	for _, _blockNumberItem := range _blockNumber {
		_blockNumberRule = append(_blockNumberRule, _blockNumberItem)
	}

	logs, sub, err := _NativeProver.contract.FilterLogs(opts, "L1WorldStateProven", _blockNumberRule)
	if err != nil {
		return nil, err
	}
	return &NativeProverL1WorldStateProvenIterator{contract: _NativeProver.contract, event: "L1WorldStateProven", logs: logs, sub: sub}, nil
}

// WatchL1WorldStateProven is a free log subscription operation binding the contract event 0xe7c32f66db8f545c384c1cc1ab188c2cc54fb6c2207477d3b141104ce9529dce.
//
// Solidity: event L1WorldStateProven(uint256 indexed _blockNumber, bytes32 _L1WorldStateRoot)
func (_NativeProver *NativeProverFilterer) WatchL1WorldStateProven(opts *bind.WatchOpts, sink chan<- *NativeProverL1WorldStateProven, _blockNumber []*big.Int) (event.Subscription, error) {

	var _blockNumberRule []interface{}
	for _, _blockNumberItem := range _blockNumber {
		_blockNumberRule = append(_blockNumberRule, _blockNumberItem)
	}

	logs, sub, err := _NativeProver.contract.WatchLogs(opts, "L1WorldStateProven", _blockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeProverL1WorldStateProven)
				if err := _NativeProver.contract.UnpackLog(event, "L1WorldStateProven", log); err != nil {
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

// ParseL1WorldStateProven is a log parse operation binding the contract event 0xe7c32f66db8f545c384c1cc1ab188c2cc54fb6c2207477d3b141104ce9529dce.
//
// Solidity: event L1WorldStateProven(uint256 indexed _blockNumber, bytes32 _L1WorldStateRoot)
func (_NativeProver *NativeProverFilterer) ParseL1WorldStateProven(log types.Log) (*NativeProverL1WorldStateProven, error) {
	event := new(NativeProverL1WorldStateProven)
	if err := _NativeProver.contract.UnpackLog(event, "L1WorldStateProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeProverL2WorldStateProvenIterator is returned from FilterL2WorldStateProven and is used to iterate over the raw logs and unpacked data for L2WorldStateProven events raised by the NativeProver contract.
type NativeProverL2WorldStateProvenIterator struct {
	Event *NativeProverL2WorldStateProven // Event containing the contract specifics and raw log

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
func (it *NativeProverL2WorldStateProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeProverL2WorldStateProven)
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
		it.Event = new(NativeProverL2WorldStateProven)
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
func (it *NativeProverL2WorldStateProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeProverL2WorldStateProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeProverL2WorldStateProven represents a L2WorldStateProven event raised by the NativeProver contract.
type NativeProverL2WorldStateProven struct {
	DestinationChainID *big.Int
	BlockNumber        *big.Int
	L2WorldStateRoot   [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterL2WorldStateProven is a free log retrieval operation binding the contract event 0x1b2eb3bd40a31077e07e65f6e39a64744149e65c101fc943f28024f9928448a9.
//
// Solidity: event L2WorldStateProven(uint256 indexed _destinationChainID, uint256 indexed _blockNumber, bytes32 _L2WorldStateRoot)
func (_NativeProver *NativeProverFilterer) FilterL2WorldStateProven(opts *bind.FilterOpts, _destinationChainID []*big.Int, _blockNumber []*big.Int) (*NativeProverL2WorldStateProvenIterator, error) {

	var _destinationChainIDRule []interface{}
	for _, _destinationChainIDItem := range _destinationChainID {
		_destinationChainIDRule = append(_destinationChainIDRule, _destinationChainIDItem)
	}
	var _blockNumberRule []interface{}
	for _, _blockNumberItem := range _blockNumber {
		_blockNumberRule = append(_blockNumberRule, _blockNumberItem)
	}

	logs, sub, err := _NativeProver.contract.FilterLogs(opts, "L2WorldStateProven", _destinationChainIDRule, _blockNumberRule)
	if err != nil {
		return nil, err
	}
	return &NativeProverL2WorldStateProvenIterator{contract: _NativeProver.contract, event: "L2WorldStateProven", logs: logs, sub: sub}, nil
}

// WatchL2WorldStateProven is a free log subscription operation binding the contract event 0x1b2eb3bd40a31077e07e65f6e39a64744149e65c101fc943f28024f9928448a9.
//
// Solidity: event L2WorldStateProven(uint256 indexed _destinationChainID, uint256 indexed _blockNumber, bytes32 _L2WorldStateRoot)
func (_NativeProver *NativeProverFilterer) WatchL2WorldStateProven(opts *bind.WatchOpts, sink chan<- *NativeProverL2WorldStateProven, _destinationChainID []*big.Int, _blockNumber []*big.Int) (event.Subscription, error) {

	var _destinationChainIDRule []interface{}
	for _, _destinationChainIDItem := range _destinationChainID {
		_destinationChainIDRule = append(_destinationChainIDRule, _destinationChainIDItem)
	}
	var _blockNumberRule []interface{}
	for _, _blockNumberItem := range _blockNumber {
		_blockNumberRule = append(_blockNumberRule, _blockNumberItem)
	}

	logs, sub, err := _NativeProver.contract.WatchLogs(opts, "L2WorldStateProven", _destinationChainIDRule, _blockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeProverL2WorldStateProven)
				if err := _NativeProver.contract.UnpackLog(event, "L2WorldStateProven", log); err != nil {
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

// ParseL2WorldStateProven is a log parse operation binding the contract event 0x1b2eb3bd40a31077e07e65f6e39a64744149e65c101fc943f28024f9928448a9.
//
// Solidity: event L2WorldStateProven(uint256 indexed _destinationChainID, uint256 indexed _blockNumber, bytes32 _L2WorldStateRoot)
func (_NativeProver *NativeProverFilterer) ParseL2WorldStateProven(log types.Log) (*NativeProverL2WorldStateProven, error) {
	event := new(NativeProverL2WorldStateProven)
	if err := _NativeProver.contract.UnpackLog(event, "L2WorldStateProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
