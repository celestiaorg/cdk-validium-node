// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package icdkvalidium

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

// ICDKValidiumBatchData is an auto generated low-level Go binding around an user-defined struct.
type ICDKValidiumBatchData struct {
	TransactionsHash   [32]byte
	GlobalExitRoot     [32]byte
	Timestamp          uint64
	MinForcedTimestamp uint64
}

// ICDKValidiumForcedBatchData is an auto generated low-level Go binding around an user-defined struct.
type ICDKValidiumForcedBatchData struct {
	Transactions       []byte
	GlobalExitRoot     [32]byte
	MinForcedTimestamp uint64
}

// ICDKValidiumInitializePackedParameters is an auto generated low-level Go binding around an user-defined struct.
type ICDKValidiumInitializePackedParameters struct {
	Admin                    common.Address
	TrustedSequencer         common.Address
	PendingStateTimeout      uint64
	TrustedAggregator        common.Address
	TrustedAggregatorTimeout uint64
}

// ICDKValidiumPendingState is an auto generated low-level Go binding around an user-defined struct.
type ICDKValidiumPendingState struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}

// ICDKValidiumSequencedBatchData is an auto generated low-level Go binding around an user-defined struct.
type ICDKValidiumSequencedBatchData struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}

// IcdkvalidiumMetaData contains all meta data concerning the Icdkvalidium contract.
var IcdkvalidiumMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ActivateForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"ConsolidatePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"OverridePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"provedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProveNonDeterministicPendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"SetMultiplierBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"SetPendingStateTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"SetTrustedAggregatorTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"SetVerifyBatchTimeTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"UpdateZkEVMVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequencedBatchNum\",\"type\":\"uint64\"}],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"}],\"name\":\"checkStateRootInsidePrime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"consolidatePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maticAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forkID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"trustedAggregatorTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structICDKValidium.InitializePackedParameters\",\"name\":\"initializePackedParameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"genesisRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isForcedBatchDisallowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"isPendingStateConsolidable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingState\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingStateConsolidated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matic\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierBatchFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"overridePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStateTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingStateTransitions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structICDKValidium.PendingState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"proveNonDeterministicPendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifierRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structICDKValidium.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structICDKValidium.ForcedBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"sequencedBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"internalType\":\"structICDKValidium.SequencedBatchData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"setMultiplierBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"setPendingStateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"setTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"setTrustedAggregatorTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"setVerifyBatchTimeTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregatorTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyBatchTimeTarget\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IcdkvalidiumABI is the input ABI used to generate the binding from.
// Deprecated: Use IcdkvalidiumMetaData.ABI instead.
var IcdkvalidiumABI = IcdkvalidiumMetaData.ABI

// Icdkvalidium is an auto generated Go binding around an Ethereum contract.
type Icdkvalidium struct {
	IcdkvalidiumCaller     // Read-only binding to the contract
	IcdkvalidiumTransactor // Write-only binding to the contract
	IcdkvalidiumFilterer   // Log filterer for contract events
}

// IcdkvalidiumCaller is an auto generated read-only Go binding around an Ethereum contract.
type IcdkvalidiumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IcdkvalidiumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IcdkvalidiumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IcdkvalidiumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IcdkvalidiumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IcdkvalidiumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IcdkvalidiumSession struct {
	Contract     *Icdkvalidium     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IcdkvalidiumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IcdkvalidiumCallerSession struct {
	Contract *IcdkvalidiumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IcdkvalidiumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IcdkvalidiumTransactorSession struct {
	Contract     *IcdkvalidiumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IcdkvalidiumRaw is an auto generated low-level Go binding around an Ethereum contract.
type IcdkvalidiumRaw struct {
	Contract *Icdkvalidium // Generic contract binding to access the raw methods on
}

// IcdkvalidiumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IcdkvalidiumCallerRaw struct {
	Contract *IcdkvalidiumCaller // Generic read-only contract binding to access the raw methods on
}

// IcdkvalidiumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IcdkvalidiumTransactorRaw struct {
	Contract *IcdkvalidiumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIcdkvalidium creates a new instance of Icdkvalidium, bound to a specific deployed contract.
func NewIcdkvalidium(address common.Address, backend bind.ContractBackend) (*Icdkvalidium, error) {
	contract, err := bindIcdkvalidium(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Icdkvalidium{IcdkvalidiumCaller: IcdkvalidiumCaller{contract: contract}, IcdkvalidiumTransactor: IcdkvalidiumTransactor{contract: contract}, IcdkvalidiumFilterer: IcdkvalidiumFilterer{contract: contract}}, nil
}

// NewIcdkvalidiumCaller creates a new read-only instance of Icdkvalidium, bound to a specific deployed contract.
func NewIcdkvalidiumCaller(address common.Address, caller bind.ContractCaller) (*IcdkvalidiumCaller, error) {
	contract, err := bindIcdkvalidium(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumCaller{contract: contract}, nil
}

// NewIcdkvalidiumTransactor creates a new write-only instance of Icdkvalidium, bound to a specific deployed contract.
func NewIcdkvalidiumTransactor(address common.Address, transactor bind.ContractTransactor) (*IcdkvalidiumTransactor, error) {
	contract, err := bindIcdkvalidium(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumTransactor{contract: contract}, nil
}

// NewIcdkvalidiumFilterer creates a new log filterer instance of Icdkvalidium, bound to a specific deployed contract.
func NewIcdkvalidiumFilterer(address common.Address, filterer bind.ContractFilterer) (*IcdkvalidiumFilterer, error) {
	contract, err := bindIcdkvalidium(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumFilterer{contract: contract}, nil
}

// bindIcdkvalidium binds a generic wrapper to an already deployed contract.
func bindIcdkvalidium(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IcdkvalidiumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Icdkvalidium *IcdkvalidiumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Icdkvalidium.Contract.IcdkvalidiumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Icdkvalidium *IcdkvalidiumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.IcdkvalidiumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Icdkvalidium *IcdkvalidiumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.IcdkvalidiumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Icdkvalidium *IcdkvalidiumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Icdkvalidium.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Icdkvalidium *IcdkvalidiumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Icdkvalidium *IcdkvalidiumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Icdkvalidium *IcdkvalidiumSession) Admin() (common.Address, error) {
	return _Icdkvalidium.Contract.Admin(&_Icdkvalidium.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCallerSession) Admin() (common.Address, error) {
	return _Icdkvalidium.Contract.Admin(&_Icdkvalidium.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Icdkvalidium *IcdkvalidiumCaller) BatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "batchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Icdkvalidium *IcdkvalidiumSession) BatchFee() (*big.Int, error) {
	return _Icdkvalidium.Contract.BatchFee(&_Icdkvalidium.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Icdkvalidium *IcdkvalidiumCallerSession) BatchFee() (*big.Int, error) {
	return _Icdkvalidium.Contract.BatchFee(&_Icdkvalidium.CallOpts)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Icdkvalidium *IcdkvalidiumCaller) BatchNumToStateRoot(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "batchNumToStateRoot", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Icdkvalidium *IcdkvalidiumSession) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Icdkvalidium.Contract.BatchNumToStateRoot(&_Icdkvalidium.CallOpts, arg0)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Icdkvalidium *IcdkvalidiumCallerSession) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Icdkvalidium.Contract.BatchNumToStateRoot(&_Icdkvalidium.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Icdkvalidium *IcdkvalidiumSession) BridgeAddress() (common.Address, error) {
	return _Icdkvalidium.Contract.BridgeAddress(&_Icdkvalidium.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCallerSession) BridgeAddress() (common.Address, error) {
	return _Icdkvalidium.Contract.BridgeAddress(&_Icdkvalidium.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Icdkvalidium *IcdkvalidiumCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Icdkvalidium *IcdkvalidiumSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Icdkvalidium.Contract.CalculateRewardPerBatch(&_Icdkvalidium.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Icdkvalidium *IcdkvalidiumCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Icdkvalidium.Contract.CalculateRewardPerBatch(&_Icdkvalidium.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) ChainID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) ChainID() (uint64, error) {
	return _Icdkvalidium.Contract.ChainID(&_Icdkvalidium.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) ChainID() (uint64, error) {
	return _Icdkvalidium.Contract.ChainID(&_Icdkvalidium.CallOpts)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Icdkvalidium *IcdkvalidiumCaller) CheckStateRootInsidePrime(opts *bind.CallOpts, newStateRoot *big.Int) (bool, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "checkStateRootInsidePrime", newStateRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Icdkvalidium *IcdkvalidiumSession) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Icdkvalidium.Contract.CheckStateRootInsidePrime(&_Icdkvalidium.CallOpts, newStateRoot)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Icdkvalidium *IcdkvalidiumCallerSession) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Icdkvalidium.Contract.CheckStateRootInsidePrime(&_Icdkvalidium.CallOpts, newStateRoot)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) ForceBatchTimeout() (uint64, error) {
	return _Icdkvalidium.Contract.ForceBatchTimeout(&_Icdkvalidium.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Icdkvalidium.Contract.ForceBatchTimeout(&_Icdkvalidium.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Icdkvalidium *IcdkvalidiumCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Icdkvalidium *IcdkvalidiumSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Icdkvalidium.Contract.ForcedBatches(&_Icdkvalidium.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Icdkvalidium *IcdkvalidiumCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Icdkvalidium.Contract.ForcedBatches(&_Icdkvalidium.CallOpts, arg0)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) ForkID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "forkID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) ForkID() (uint64, error) {
	return _Icdkvalidium.Contract.ForkID(&_Icdkvalidium.CallOpts)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) ForkID() (uint64, error) {
	return _Icdkvalidium.Contract.ForkID(&_Icdkvalidium.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Icdkvalidium *IcdkvalidiumCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Icdkvalidium *IcdkvalidiumSession) GetForcedBatchFee() (*big.Int, error) {
	return _Icdkvalidium.Contract.GetForcedBatchFee(&_Icdkvalidium.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Icdkvalidium *IcdkvalidiumCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Icdkvalidium.Contract.GetForcedBatchFee(&_Icdkvalidium.CallOpts)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Icdkvalidium *IcdkvalidiumCaller) GetInputSnarkBytes(opts *bind.CallOpts, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "getInputSnarkBytes", initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Icdkvalidium *IcdkvalidiumSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Icdkvalidium.Contract.GetInputSnarkBytes(&_Icdkvalidium.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Icdkvalidium *IcdkvalidiumCallerSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Icdkvalidium.Contract.GetInputSnarkBytes(&_Icdkvalidium.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) GetLastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "getLastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) GetLastVerifiedBatch() (uint64, error) {
	return _Icdkvalidium.Contract.GetLastVerifiedBatch(&_Icdkvalidium.CallOpts)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) GetLastVerifiedBatch() (uint64, error) {
	return _Icdkvalidium.Contract.GetLastVerifiedBatch(&_Icdkvalidium.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Icdkvalidium *IcdkvalidiumSession) GlobalExitRootManager() (common.Address, error) {
	return _Icdkvalidium.Contract.GlobalExitRootManager(&_Icdkvalidium.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Icdkvalidium.Contract.GlobalExitRootManager(&_Icdkvalidium.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Icdkvalidium *IcdkvalidiumCaller) IsForcedBatchDisallowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "isForcedBatchDisallowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Icdkvalidium *IcdkvalidiumSession) IsForcedBatchDisallowed() (bool, error) {
	return _Icdkvalidium.Contract.IsForcedBatchDisallowed(&_Icdkvalidium.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Icdkvalidium *IcdkvalidiumCallerSession) IsForcedBatchDisallowed() (bool, error) {
	return _Icdkvalidium.Contract.IsForcedBatchDisallowed(&_Icdkvalidium.CallOpts)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Icdkvalidium *IcdkvalidiumCaller) IsPendingStateConsolidable(opts *bind.CallOpts, pendingStateNum uint64) (bool, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "isPendingStateConsolidable", pendingStateNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Icdkvalidium *IcdkvalidiumSession) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Icdkvalidium.Contract.IsPendingStateConsolidable(&_Icdkvalidium.CallOpts, pendingStateNum)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Icdkvalidium *IcdkvalidiumCallerSession) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Icdkvalidium.Contract.IsPendingStateConsolidable(&_Icdkvalidium.CallOpts, pendingStateNum)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) LastBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "lastBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) LastBatchSequenced() (uint64, error) {
	return _Icdkvalidium.Contract.LastBatchSequenced(&_Icdkvalidium.CallOpts)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) LastBatchSequenced() (uint64, error) {
	return _Icdkvalidium.Contract.LastBatchSequenced(&_Icdkvalidium.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) LastForceBatch() (uint64, error) {
	return _Icdkvalidium.Contract.LastForceBatch(&_Icdkvalidium.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) LastForceBatch() (uint64, error) {
	return _Icdkvalidium.Contract.LastForceBatch(&_Icdkvalidium.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) LastForceBatchSequenced() (uint64, error) {
	return _Icdkvalidium.Contract.LastForceBatchSequenced(&_Icdkvalidium.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Icdkvalidium.Contract.LastForceBatchSequenced(&_Icdkvalidium.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) LastPendingState(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "lastPendingState")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) LastPendingState() (uint64, error) {
	return _Icdkvalidium.Contract.LastPendingState(&_Icdkvalidium.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) LastPendingState() (uint64, error) {
	return _Icdkvalidium.Contract.LastPendingState(&_Icdkvalidium.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) LastPendingStateConsolidated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "lastPendingStateConsolidated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) LastPendingStateConsolidated() (uint64, error) {
	return _Icdkvalidium.Contract.LastPendingStateConsolidated(&_Icdkvalidium.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) LastPendingStateConsolidated() (uint64, error) {
	return _Icdkvalidium.Contract.LastPendingStateConsolidated(&_Icdkvalidium.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) LastTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "lastTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) LastTimestamp() (uint64, error) {
	return _Icdkvalidium.Contract.LastTimestamp(&_Icdkvalidium.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) LastTimestamp() (uint64, error) {
	return _Icdkvalidium.Contract.LastTimestamp(&_Icdkvalidium.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) LastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "lastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) LastVerifiedBatch() (uint64, error) {
	return _Icdkvalidium.Contract.LastVerifiedBatch(&_Icdkvalidium.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) LastVerifiedBatch() (uint64, error) {
	return _Icdkvalidium.Contract.LastVerifiedBatch(&_Icdkvalidium.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCaller) Matic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "matic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Icdkvalidium *IcdkvalidiumSession) Matic() (common.Address, error) {
	return _Icdkvalidium.Contract.Matic(&_Icdkvalidium.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCallerSession) Matic() (common.Address, error) {
	return _Icdkvalidium.Contract.Matic(&_Icdkvalidium.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Icdkvalidium *IcdkvalidiumCaller) MultiplierBatchFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "multiplierBatchFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Icdkvalidium *IcdkvalidiumSession) MultiplierBatchFee() (uint16, error) {
	return _Icdkvalidium.Contract.MultiplierBatchFee(&_Icdkvalidium.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Icdkvalidium *IcdkvalidiumCallerSession) MultiplierBatchFee() (uint16, error) {
	return _Icdkvalidium.Contract.MultiplierBatchFee(&_Icdkvalidium.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Icdkvalidium *IcdkvalidiumCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Icdkvalidium *IcdkvalidiumSession) NetworkName() (string, error) {
	return _Icdkvalidium.Contract.NetworkName(&_Icdkvalidium.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Icdkvalidium *IcdkvalidiumCallerSession) NetworkName() (string, error) {
	return _Icdkvalidium.Contract.NetworkName(&_Icdkvalidium.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Icdkvalidium *IcdkvalidiumSession) PendingAdmin() (common.Address, error) {
	return _Icdkvalidium.Contract.PendingAdmin(&_Icdkvalidium.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCallerSession) PendingAdmin() (common.Address, error) {
	return _Icdkvalidium.Contract.PendingAdmin(&_Icdkvalidium.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) PendingStateTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "pendingStateTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) PendingStateTimeout() (uint64, error) {
	return _Icdkvalidium.Contract.PendingStateTimeout(&_Icdkvalidium.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) PendingStateTimeout() (uint64, error) {
	return _Icdkvalidium.Contract.PendingStateTimeout(&_Icdkvalidium.CallOpts)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns((uint64,uint64,bytes32,bytes32))
func (_Icdkvalidium *IcdkvalidiumCaller) PendingStateTransitions(opts *bind.CallOpts, arg0 *big.Int) (ICDKValidiumPendingState, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "pendingStateTransitions", arg0)

	if err != nil {
		return *new(ICDKValidiumPendingState), err
	}

	out0 := *abi.ConvertType(out[0], new(ICDKValidiumPendingState)).(*ICDKValidiumPendingState)

	return out0, err

}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns((uint64,uint64,bytes32,bytes32))
func (_Icdkvalidium *IcdkvalidiumSession) PendingStateTransitions(arg0 *big.Int) (ICDKValidiumPendingState, error) {
	return _Icdkvalidium.Contract.PendingStateTransitions(&_Icdkvalidium.CallOpts, arg0)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns((uint64,uint64,bytes32,bytes32))
func (_Icdkvalidium *IcdkvalidiumCallerSession) PendingStateTransitions(arg0 *big.Int) (ICDKValidiumPendingState, error) {
	return _Icdkvalidium.Contract.PendingStateTransitions(&_Icdkvalidium.CallOpts, arg0)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCaller) RollupVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "rollupVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Icdkvalidium *IcdkvalidiumSession) RollupVerifier() (common.Address, error) {
	return _Icdkvalidium.Contract.RollupVerifier(&_Icdkvalidium.CallOpts)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCallerSession) RollupVerifier() (common.Address, error) {
	return _Icdkvalidium.Contract.RollupVerifier(&_Icdkvalidium.CallOpts)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns((bytes32,uint64,uint64))
func (_Icdkvalidium *IcdkvalidiumCaller) SequencedBatches(opts *bind.CallOpts, arg0 uint64) (ICDKValidiumSequencedBatchData, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "sequencedBatches", arg0)

	if err != nil {
		return *new(ICDKValidiumSequencedBatchData), err
	}

	out0 := *abi.ConvertType(out[0], new(ICDKValidiumSequencedBatchData)).(*ICDKValidiumSequencedBatchData)

	return out0, err

}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns((bytes32,uint64,uint64))
func (_Icdkvalidium *IcdkvalidiumSession) SequencedBatches(arg0 uint64) (ICDKValidiumSequencedBatchData, error) {
	return _Icdkvalidium.Contract.SequencedBatches(&_Icdkvalidium.CallOpts, arg0)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns((bytes32,uint64,uint64))
func (_Icdkvalidium *IcdkvalidiumCallerSession) SequencedBatches(arg0 uint64) (ICDKValidiumSequencedBatchData, error) {
	return _Icdkvalidium.Contract.SequencedBatches(&_Icdkvalidium.CallOpts, arg0)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCaller) TrustedAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "trustedAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Icdkvalidium *IcdkvalidiumSession) TrustedAggregator() (common.Address, error) {
	return _Icdkvalidium.Contract.TrustedAggregator(&_Icdkvalidium.CallOpts)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCallerSession) TrustedAggregator() (common.Address, error) {
	return _Icdkvalidium.Contract.TrustedAggregator(&_Icdkvalidium.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) TrustedAggregatorTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "trustedAggregatorTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Icdkvalidium.Contract.TrustedAggregatorTimeout(&_Icdkvalidium.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Icdkvalidium.Contract.TrustedAggregatorTimeout(&_Icdkvalidium.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Icdkvalidium *IcdkvalidiumSession) TrustedSequencer() (common.Address, error) {
	return _Icdkvalidium.Contract.TrustedSequencer(&_Icdkvalidium.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Icdkvalidium *IcdkvalidiumCallerSession) TrustedSequencer() (common.Address, error) {
	return _Icdkvalidium.Contract.TrustedSequencer(&_Icdkvalidium.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Icdkvalidium *IcdkvalidiumCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Icdkvalidium *IcdkvalidiumSession) TrustedSequencerURL() (string, error) {
	return _Icdkvalidium.Contract.TrustedSequencerURL(&_Icdkvalidium.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Icdkvalidium *IcdkvalidiumCallerSession) TrustedSequencerURL() (string, error) {
	return _Icdkvalidium.Contract.TrustedSequencerURL(&_Icdkvalidium.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCaller) VerifyBatchTimeTarget(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Icdkvalidium.contract.Call(opts, &out, "verifyBatchTimeTarget")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Icdkvalidium.Contract.VerifyBatchTimeTarget(&_Icdkvalidium.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Icdkvalidium *IcdkvalidiumCallerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Icdkvalidium.Contract.VerifyBatchTimeTarget(&_Icdkvalidium.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Icdkvalidium *IcdkvalidiumSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Icdkvalidium.Contract.AcceptAdminRole(&_Icdkvalidium.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Icdkvalidium.Contract.AcceptAdminRole(&_Icdkvalidium.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) ActivateEmergencyState(opts *bind.TransactOpts, sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "activateEmergencyState", sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Icdkvalidium *IcdkvalidiumSession) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.ActivateEmergencyState(&_Icdkvalidium.TransactOpts, sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.ActivateEmergencyState(&_Icdkvalidium.TransactOpts, sequencedBatchNum)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) ActivateForceBatches(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "activateForceBatches")
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Icdkvalidium *IcdkvalidiumSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Icdkvalidium.Contract.ActivateForceBatches(&_Icdkvalidium.TransactOpts)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Icdkvalidium.Contract.ActivateForceBatches(&_Icdkvalidium.TransactOpts)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) ConsolidatePendingState(opts *bind.TransactOpts, pendingStateNum uint64) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "consolidatePendingState", pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Icdkvalidium *IcdkvalidiumSession) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.ConsolidatePendingState(&_Icdkvalidium.TransactOpts, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.ConsolidatePendingState(&_Icdkvalidium.TransactOpts, pendingStateNum)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Icdkvalidium *IcdkvalidiumSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Icdkvalidium.Contract.DeactivateEmergencyState(&_Icdkvalidium.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Icdkvalidium.Contract.DeactivateEmergencyState(&_Icdkvalidium.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "forceBatch", transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Icdkvalidium *IcdkvalidiumSession) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.ForceBatch(&_Icdkvalidium.TransactOpts, transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.ForceBatch(&_Icdkvalidium.TransactOpts, transactions, maticAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) Initialize(opts *bind.TransactOpts, initializePackedParameters ICDKValidiumInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "initialize", initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Icdkvalidium *IcdkvalidiumSession) Initialize(initializePackedParameters ICDKValidiumInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.Initialize(&_Icdkvalidium.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) Initialize(initializePackedParameters ICDKValidiumInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.Initialize(&_Icdkvalidium.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) OverridePendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "overridePendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Icdkvalidium *IcdkvalidiumSession) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.OverridePendingState(&_Icdkvalidium.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.OverridePendingState(&_Icdkvalidium.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) ProveNonDeterministicPendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "proveNonDeterministicPendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Icdkvalidium *IcdkvalidiumSession) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.ProveNonDeterministicPendingState(&_Icdkvalidium.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.ProveNonDeterministicPendingState(&_Icdkvalidium.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x438a5399.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes message) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) SequenceBatches(opts *bind.TransactOpts, batches []ICDKValidiumBatchData, l2Coinbase common.Address, message []byte) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "sequenceBatches", batches, l2Coinbase, message)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x438a5399.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes message) returns()
func (_Icdkvalidium *IcdkvalidiumSession) SequenceBatches(batches []ICDKValidiumBatchData, l2Coinbase common.Address, message []byte) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SequenceBatches(&_Icdkvalidium.TransactOpts, batches, l2Coinbase, message)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x438a5399.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes message) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) SequenceBatches(batches []ICDKValidiumBatchData, l2Coinbase common.Address, message []byte) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SequenceBatches(&_Icdkvalidium.TransactOpts, batches, l2Coinbase, message)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []ICDKValidiumForcedBatchData) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Icdkvalidium *IcdkvalidiumSession) SequenceForceBatches(batches []ICDKValidiumForcedBatchData) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SequenceForceBatches(&_Icdkvalidium.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) SequenceForceBatches(batches []ICDKValidiumForcedBatchData) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SequenceForceBatches(&_Icdkvalidium.TransactOpts, batches)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Icdkvalidium *IcdkvalidiumSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetForceBatchTimeout(&_Icdkvalidium.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetForceBatchTimeout(&_Icdkvalidium.TransactOpts, newforceBatchTimeout)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) SetMultiplierBatchFee(opts *bind.TransactOpts, newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "setMultiplierBatchFee", newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Icdkvalidium *IcdkvalidiumSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetMultiplierBatchFee(&_Icdkvalidium.TransactOpts, newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetMultiplierBatchFee(&_Icdkvalidium.TransactOpts, newMultiplierBatchFee)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) SetPendingStateTimeout(opts *bind.TransactOpts, newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "setPendingStateTimeout", newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Icdkvalidium *IcdkvalidiumSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetPendingStateTimeout(&_Icdkvalidium.TransactOpts, newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetPendingStateTimeout(&_Icdkvalidium.TransactOpts, newPendingStateTimeout)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) SetTrustedAggregator(opts *bind.TransactOpts, newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "setTrustedAggregator", newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Icdkvalidium *IcdkvalidiumSession) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetTrustedAggregator(&_Icdkvalidium.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetTrustedAggregator(&_Icdkvalidium.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) SetTrustedAggregatorTimeout(opts *bind.TransactOpts, newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "setTrustedAggregatorTimeout", newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Icdkvalidium *IcdkvalidiumSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetTrustedAggregatorTimeout(&_Icdkvalidium.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetTrustedAggregatorTimeout(&_Icdkvalidium.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Icdkvalidium *IcdkvalidiumSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetTrustedSequencer(&_Icdkvalidium.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetTrustedSequencer(&_Icdkvalidium.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Icdkvalidium *IcdkvalidiumSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetTrustedSequencerURL(&_Icdkvalidium.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetTrustedSequencerURL(&_Icdkvalidium.TransactOpts, newTrustedSequencerURL)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) SetVerifyBatchTimeTarget(opts *bind.TransactOpts, newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "setVerifyBatchTimeTarget", newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Icdkvalidium *IcdkvalidiumSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetVerifyBatchTimeTarget(&_Icdkvalidium.TransactOpts, newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.SetVerifyBatchTimeTarget(&_Icdkvalidium.TransactOpts, newVerifyBatchTimeTarget)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Icdkvalidium *IcdkvalidiumSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.TransferAdminRole(&_Icdkvalidium.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.TransferAdminRole(&_Icdkvalidium.TransactOpts, newPendingAdmin)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) VerifyBatches(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "verifyBatches", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Icdkvalidium *IcdkvalidiumSession) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.VerifyBatches(&_Icdkvalidium.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.VerifyBatches(&_Icdkvalidium.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Icdkvalidium *IcdkvalidiumTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Icdkvalidium.contract.Transact(opts, "verifyBatchesTrustedAggregator", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Icdkvalidium *IcdkvalidiumSession) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.VerifyBatchesTrustedAggregator(&_Icdkvalidium.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Icdkvalidium *IcdkvalidiumTransactorSession) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Icdkvalidium.Contract.VerifyBatchesTrustedAggregator(&_Icdkvalidium.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// IcdkvalidiumAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Icdkvalidium contract.
type IcdkvalidiumAcceptAdminRoleIterator struct {
	Event *IcdkvalidiumAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumAcceptAdminRole)
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
		it.Event = new(IcdkvalidiumAcceptAdminRole)
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
func (it *IcdkvalidiumAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumAcceptAdminRole represents a AcceptAdminRole event raised by the Icdkvalidium contract.
type IcdkvalidiumAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*IcdkvalidiumAcceptAdminRoleIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumAcceptAdminRoleIterator{contract: _Icdkvalidium.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumAcceptAdminRole)
				if err := _Icdkvalidium.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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

// ParseAcceptAdminRole is a log parse operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseAcceptAdminRole(log types.Log) (*IcdkvalidiumAcceptAdminRole, error) {
	event := new(IcdkvalidiumAcceptAdminRole)
	if err := _Icdkvalidium.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumActivateForceBatchesIterator is returned from FilterActivateForceBatches and is used to iterate over the raw logs and unpacked data for ActivateForceBatches events raised by the Icdkvalidium contract.
type IcdkvalidiumActivateForceBatchesIterator struct {
	Event *IcdkvalidiumActivateForceBatches // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumActivateForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumActivateForceBatches)
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
		it.Event = new(IcdkvalidiumActivateForceBatches)
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
func (it *IcdkvalidiumActivateForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumActivateForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumActivateForceBatches represents a ActivateForceBatches event raised by the Icdkvalidium contract.
type IcdkvalidiumActivateForceBatches struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActivateForceBatches is a free log retrieval operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterActivateForceBatches(opts *bind.FilterOpts) (*IcdkvalidiumActivateForceBatchesIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumActivateForceBatchesIterator{contract: _Icdkvalidium.contract, event: "ActivateForceBatches", logs: logs, sub: sub}, nil
}

// WatchActivateForceBatches is a free log subscription operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchActivateForceBatches(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumActivateForceBatches) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumActivateForceBatches)
				if err := _Icdkvalidium.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
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

// ParseActivateForceBatches is a log parse operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseActivateForceBatches(log types.Log) (*IcdkvalidiumActivateForceBatches, error) {
	event := new(IcdkvalidiumActivateForceBatches)
	if err := _Icdkvalidium.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumConsolidatePendingStateIterator is returned from FilterConsolidatePendingState and is used to iterate over the raw logs and unpacked data for ConsolidatePendingState events raised by the Icdkvalidium contract.
type IcdkvalidiumConsolidatePendingStateIterator struct {
	Event *IcdkvalidiumConsolidatePendingState // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumConsolidatePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumConsolidatePendingState)
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
		it.Event = new(IcdkvalidiumConsolidatePendingState)
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
func (it *IcdkvalidiumConsolidatePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumConsolidatePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumConsolidatePendingState represents a ConsolidatePendingState event raised by the Icdkvalidium contract.
type IcdkvalidiumConsolidatePendingState struct {
	NumBatch        uint64
	StateRoot       [32]byte
	PendingStateNum uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterConsolidatePendingState is a free log retrieval operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterConsolidatePendingState(opts *bind.FilterOpts, numBatch []uint64, pendingStateNum []uint64) (*IcdkvalidiumConsolidatePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumConsolidatePendingStateIterator{contract: _Icdkvalidium.contract, event: "ConsolidatePendingState", logs: logs, sub: sub}, nil
}

// WatchConsolidatePendingState is a free log subscription operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchConsolidatePendingState(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumConsolidatePendingState, numBatch []uint64, pendingStateNum []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumConsolidatePendingState)
				if err := _Icdkvalidium.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
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

// ParseConsolidatePendingState is a log parse operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseConsolidatePendingState(log types.Log) (*IcdkvalidiumConsolidatePendingState, error) {
	event := new(IcdkvalidiumConsolidatePendingState)
	if err := _Icdkvalidium.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Icdkvalidium contract.
type IcdkvalidiumForceBatchIterator struct {
	Event *IcdkvalidiumForceBatch // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumForceBatch)
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
		it.Event = new(IcdkvalidiumForceBatch)
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
func (it *IcdkvalidiumForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumForceBatch represents a ForceBatch event raised by the Icdkvalidium contract.
type IcdkvalidiumForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*IcdkvalidiumForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumForceBatchIterator{contract: _Icdkvalidium.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumForceBatch)
				if err := _Icdkvalidium.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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

// ParseForceBatch is a log parse operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseForceBatch(log types.Log) (*IcdkvalidiumForceBatch, error) {
	event := new(IcdkvalidiumForceBatch)
	if err := _Icdkvalidium.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumOverridePendingStateIterator is returned from FilterOverridePendingState and is used to iterate over the raw logs and unpacked data for OverridePendingState events raised by the Icdkvalidium contract.
type IcdkvalidiumOverridePendingStateIterator struct {
	Event *IcdkvalidiumOverridePendingState // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumOverridePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumOverridePendingState)
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
		it.Event = new(IcdkvalidiumOverridePendingState)
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
func (it *IcdkvalidiumOverridePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumOverridePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumOverridePendingState represents a OverridePendingState event raised by the Icdkvalidium contract.
type IcdkvalidiumOverridePendingState struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOverridePendingState is a free log retrieval operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterOverridePendingState(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*IcdkvalidiumOverridePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumOverridePendingStateIterator{contract: _Icdkvalidium.contract, event: "OverridePendingState", logs: logs, sub: sub}, nil
}

// WatchOverridePendingState is a free log subscription operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchOverridePendingState(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumOverridePendingState, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumOverridePendingState)
				if err := _Icdkvalidium.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
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

// ParseOverridePendingState is a log parse operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseOverridePendingState(log types.Log) (*IcdkvalidiumOverridePendingState, error) {
	event := new(IcdkvalidiumOverridePendingState)
	if err := _Icdkvalidium.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumProveNonDeterministicPendingStateIterator is returned from FilterProveNonDeterministicPendingState and is used to iterate over the raw logs and unpacked data for ProveNonDeterministicPendingState events raised by the Icdkvalidium contract.
type IcdkvalidiumProveNonDeterministicPendingStateIterator struct {
	Event *IcdkvalidiumProveNonDeterministicPendingState // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumProveNonDeterministicPendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumProveNonDeterministicPendingState)
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
		it.Event = new(IcdkvalidiumProveNonDeterministicPendingState)
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
func (it *IcdkvalidiumProveNonDeterministicPendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumProveNonDeterministicPendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumProveNonDeterministicPendingState represents a ProveNonDeterministicPendingState event raised by the Icdkvalidium contract.
type IcdkvalidiumProveNonDeterministicPendingState struct {
	StoredStateRoot [32]byte
	ProvedStateRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProveNonDeterministicPendingState is a free log retrieval operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterProveNonDeterministicPendingState(opts *bind.FilterOpts) (*IcdkvalidiumProveNonDeterministicPendingStateIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumProveNonDeterministicPendingStateIterator{contract: _Icdkvalidium.contract, event: "ProveNonDeterministicPendingState", logs: logs, sub: sub}, nil
}

// WatchProveNonDeterministicPendingState is a free log subscription operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchProveNonDeterministicPendingState(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumProveNonDeterministicPendingState) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumProveNonDeterministicPendingState)
				if err := _Icdkvalidium.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
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

// ParseProveNonDeterministicPendingState is a log parse operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseProveNonDeterministicPendingState(log types.Log) (*IcdkvalidiumProveNonDeterministicPendingState, error) {
	event := new(IcdkvalidiumProveNonDeterministicPendingState)
	if err := _Icdkvalidium.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Icdkvalidium contract.
type IcdkvalidiumSequenceBatchesIterator struct {
	Event *IcdkvalidiumSequenceBatches // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumSequenceBatches)
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
		it.Event = new(IcdkvalidiumSequenceBatches)
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
func (it *IcdkvalidiumSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumSequenceBatches represents a SequenceBatches event raised by the Icdkvalidium contract.
type IcdkvalidiumSequenceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*IcdkvalidiumSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumSequenceBatchesIterator{contract: _Icdkvalidium.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumSequenceBatches)
				if err := _Icdkvalidium.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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

// ParseSequenceBatches is a log parse operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseSequenceBatches(log types.Log) (*IcdkvalidiumSequenceBatches, error) {
	event := new(IcdkvalidiumSequenceBatches)
	if err := _Icdkvalidium.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Icdkvalidium contract.
type IcdkvalidiumSequenceForceBatchesIterator struct {
	Event *IcdkvalidiumSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumSequenceForceBatches)
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
		it.Event = new(IcdkvalidiumSequenceForceBatches)
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
func (it *IcdkvalidiumSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumSequenceForceBatches represents a SequenceForceBatches event raised by the Icdkvalidium contract.
type IcdkvalidiumSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*IcdkvalidiumSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumSequenceForceBatchesIterator{contract: _Icdkvalidium.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumSequenceForceBatches)
				if err := _Icdkvalidium.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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

// ParseSequenceForceBatches is a log parse operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseSequenceForceBatches(log types.Log) (*IcdkvalidiumSequenceForceBatches, error) {
	event := new(IcdkvalidiumSequenceForceBatches)
	if err := _Icdkvalidium.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Icdkvalidium contract.
type IcdkvalidiumSetForceBatchTimeoutIterator struct {
	Event *IcdkvalidiumSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumSetForceBatchTimeout)
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
		it.Event = new(IcdkvalidiumSetForceBatchTimeout)
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
func (it *IcdkvalidiumSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Icdkvalidium contract.
type IcdkvalidiumSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*IcdkvalidiumSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumSetForceBatchTimeoutIterator{contract: _Icdkvalidium.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumSetForceBatchTimeout)
				if err := _Icdkvalidium.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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

// ParseSetForceBatchTimeout is a log parse operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseSetForceBatchTimeout(log types.Log) (*IcdkvalidiumSetForceBatchTimeout, error) {
	event := new(IcdkvalidiumSetForceBatchTimeout)
	if err := _Icdkvalidium.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumSetMultiplierBatchFeeIterator is returned from FilterSetMultiplierBatchFee and is used to iterate over the raw logs and unpacked data for SetMultiplierBatchFee events raised by the Icdkvalidium contract.
type IcdkvalidiumSetMultiplierBatchFeeIterator struct {
	Event *IcdkvalidiumSetMultiplierBatchFee // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumSetMultiplierBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumSetMultiplierBatchFee)
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
		it.Event = new(IcdkvalidiumSetMultiplierBatchFee)
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
func (it *IcdkvalidiumSetMultiplierBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumSetMultiplierBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumSetMultiplierBatchFee represents a SetMultiplierBatchFee event raised by the Icdkvalidium contract.
type IcdkvalidiumSetMultiplierBatchFee struct {
	NewMultiplierBatchFee uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMultiplierBatchFee is a free log retrieval operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterSetMultiplierBatchFee(opts *bind.FilterOpts) (*IcdkvalidiumSetMultiplierBatchFeeIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumSetMultiplierBatchFeeIterator{contract: _Icdkvalidium.contract, event: "SetMultiplierBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetMultiplierBatchFee is a free log subscription operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchSetMultiplierBatchFee(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumSetMultiplierBatchFee) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumSetMultiplierBatchFee)
				if err := _Icdkvalidium.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
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

// ParseSetMultiplierBatchFee is a log parse operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseSetMultiplierBatchFee(log types.Log) (*IcdkvalidiumSetMultiplierBatchFee, error) {
	event := new(IcdkvalidiumSetMultiplierBatchFee)
	if err := _Icdkvalidium.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumSetPendingStateTimeoutIterator is returned from FilterSetPendingStateTimeout and is used to iterate over the raw logs and unpacked data for SetPendingStateTimeout events raised by the Icdkvalidium contract.
type IcdkvalidiumSetPendingStateTimeoutIterator struct {
	Event *IcdkvalidiumSetPendingStateTimeout // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumSetPendingStateTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumSetPendingStateTimeout)
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
		it.Event = new(IcdkvalidiumSetPendingStateTimeout)
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
func (it *IcdkvalidiumSetPendingStateTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumSetPendingStateTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumSetPendingStateTimeout represents a SetPendingStateTimeout event raised by the Icdkvalidium contract.
type IcdkvalidiumSetPendingStateTimeout struct {
	NewPendingStateTimeout uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetPendingStateTimeout is a free log retrieval operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterSetPendingStateTimeout(opts *bind.FilterOpts) (*IcdkvalidiumSetPendingStateTimeoutIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumSetPendingStateTimeoutIterator{contract: _Icdkvalidium.contract, event: "SetPendingStateTimeout", logs: logs, sub: sub}, nil
}

// WatchSetPendingStateTimeout is a free log subscription operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchSetPendingStateTimeout(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumSetPendingStateTimeout) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumSetPendingStateTimeout)
				if err := _Icdkvalidium.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
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

// ParseSetPendingStateTimeout is a log parse operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseSetPendingStateTimeout(log types.Log) (*IcdkvalidiumSetPendingStateTimeout, error) {
	event := new(IcdkvalidiumSetPendingStateTimeout)
	if err := _Icdkvalidium.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Icdkvalidium contract.
type IcdkvalidiumSetTrustedAggregatorIterator struct {
	Event *IcdkvalidiumSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumSetTrustedAggregator)
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
		it.Event = new(IcdkvalidiumSetTrustedAggregator)
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
func (it *IcdkvalidiumSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumSetTrustedAggregator represents a SetTrustedAggregator event raised by the Icdkvalidium contract.
type IcdkvalidiumSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*IcdkvalidiumSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumSetTrustedAggregatorIterator{contract: _Icdkvalidium.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumSetTrustedAggregator)
				if err := _Icdkvalidium.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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

// ParseSetTrustedAggregator is a log parse operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseSetTrustedAggregator(log types.Log) (*IcdkvalidiumSetTrustedAggregator, error) {
	event := new(IcdkvalidiumSetTrustedAggregator)
	if err := _Icdkvalidium.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumSetTrustedAggregatorTimeoutIterator is returned from FilterSetTrustedAggregatorTimeout and is used to iterate over the raw logs and unpacked data for SetTrustedAggregatorTimeout events raised by the Icdkvalidium contract.
type IcdkvalidiumSetTrustedAggregatorTimeoutIterator struct {
	Event *IcdkvalidiumSetTrustedAggregatorTimeout // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumSetTrustedAggregatorTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumSetTrustedAggregatorTimeout)
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
		it.Event = new(IcdkvalidiumSetTrustedAggregatorTimeout)
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
func (it *IcdkvalidiumSetTrustedAggregatorTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumSetTrustedAggregatorTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumSetTrustedAggregatorTimeout represents a SetTrustedAggregatorTimeout event raised by the Icdkvalidium contract.
type IcdkvalidiumSetTrustedAggregatorTimeout struct {
	NewTrustedAggregatorTimeout uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregatorTimeout is a free log retrieval operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterSetTrustedAggregatorTimeout(opts *bind.FilterOpts) (*IcdkvalidiumSetTrustedAggregatorTimeoutIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumSetTrustedAggregatorTimeoutIterator{contract: _Icdkvalidium.contract, event: "SetTrustedAggregatorTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregatorTimeout is a free log subscription operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchSetTrustedAggregatorTimeout(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumSetTrustedAggregatorTimeout) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumSetTrustedAggregatorTimeout)
				if err := _Icdkvalidium.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
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

// ParseSetTrustedAggregatorTimeout is a log parse operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseSetTrustedAggregatorTimeout(log types.Log) (*IcdkvalidiumSetTrustedAggregatorTimeout, error) {
	event := new(IcdkvalidiumSetTrustedAggregatorTimeout)
	if err := _Icdkvalidium.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Icdkvalidium contract.
type IcdkvalidiumSetTrustedSequencerIterator struct {
	Event *IcdkvalidiumSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumSetTrustedSequencer)
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
		it.Event = new(IcdkvalidiumSetTrustedSequencer)
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
func (it *IcdkvalidiumSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumSetTrustedSequencer represents a SetTrustedSequencer event raised by the Icdkvalidium contract.
type IcdkvalidiumSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*IcdkvalidiumSetTrustedSequencerIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumSetTrustedSequencerIterator{contract: _Icdkvalidium.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumSetTrustedSequencer)
				if err := _Icdkvalidium.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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

// ParseSetTrustedSequencer is a log parse operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseSetTrustedSequencer(log types.Log) (*IcdkvalidiumSetTrustedSequencer, error) {
	event := new(IcdkvalidiumSetTrustedSequencer)
	if err := _Icdkvalidium.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Icdkvalidium contract.
type IcdkvalidiumSetTrustedSequencerURLIterator struct {
	Event *IcdkvalidiumSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumSetTrustedSequencerURL)
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
		it.Event = new(IcdkvalidiumSetTrustedSequencerURL)
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
func (it *IcdkvalidiumSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Icdkvalidium contract.
type IcdkvalidiumSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*IcdkvalidiumSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumSetTrustedSequencerURLIterator{contract: _Icdkvalidium.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumSetTrustedSequencerURL)
				if err := _Icdkvalidium.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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

// ParseSetTrustedSequencerURL is a log parse operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseSetTrustedSequencerURL(log types.Log) (*IcdkvalidiumSetTrustedSequencerURL, error) {
	event := new(IcdkvalidiumSetTrustedSequencerURL)
	if err := _Icdkvalidium.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumSetVerifyBatchTimeTargetIterator is returned from FilterSetVerifyBatchTimeTarget and is used to iterate over the raw logs and unpacked data for SetVerifyBatchTimeTarget events raised by the Icdkvalidium contract.
type IcdkvalidiumSetVerifyBatchTimeTargetIterator struct {
	Event *IcdkvalidiumSetVerifyBatchTimeTarget // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumSetVerifyBatchTimeTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumSetVerifyBatchTimeTarget)
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
		it.Event = new(IcdkvalidiumSetVerifyBatchTimeTarget)
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
func (it *IcdkvalidiumSetVerifyBatchTimeTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumSetVerifyBatchTimeTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumSetVerifyBatchTimeTarget represents a SetVerifyBatchTimeTarget event raised by the Icdkvalidium contract.
type IcdkvalidiumSetVerifyBatchTimeTarget struct {
	NewVerifyBatchTimeTarget uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetVerifyBatchTimeTarget is a free log retrieval operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterSetVerifyBatchTimeTarget(opts *bind.FilterOpts) (*IcdkvalidiumSetVerifyBatchTimeTargetIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumSetVerifyBatchTimeTargetIterator{contract: _Icdkvalidium.contract, event: "SetVerifyBatchTimeTarget", logs: logs, sub: sub}, nil
}

// WatchSetVerifyBatchTimeTarget is a free log subscription operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchSetVerifyBatchTimeTarget(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumSetVerifyBatchTimeTarget) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumSetVerifyBatchTimeTarget)
				if err := _Icdkvalidium.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
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

// ParseSetVerifyBatchTimeTarget is a log parse operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseSetVerifyBatchTimeTarget(log types.Log) (*IcdkvalidiumSetVerifyBatchTimeTarget, error) {
	event := new(IcdkvalidiumSetVerifyBatchTimeTarget)
	if err := _Icdkvalidium.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Icdkvalidium contract.
type IcdkvalidiumTransferAdminRoleIterator struct {
	Event *IcdkvalidiumTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumTransferAdminRole)
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
		it.Event = new(IcdkvalidiumTransferAdminRole)
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
func (it *IcdkvalidiumTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumTransferAdminRole represents a TransferAdminRole event raised by the Icdkvalidium contract.
type IcdkvalidiumTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*IcdkvalidiumTransferAdminRoleIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumTransferAdminRoleIterator{contract: _Icdkvalidium.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumTransferAdminRole)
				if err := _Icdkvalidium.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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

// ParseTransferAdminRole is a log parse operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseTransferAdminRole(log types.Log) (*IcdkvalidiumTransferAdminRole, error) {
	event := new(IcdkvalidiumTransferAdminRole)
	if err := _Icdkvalidium.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumUpdateZkEVMVersionIterator is returned from FilterUpdateZkEVMVersion and is used to iterate over the raw logs and unpacked data for UpdateZkEVMVersion events raised by the Icdkvalidium contract.
type IcdkvalidiumUpdateZkEVMVersionIterator struct {
	Event *IcdkvalidiumUpdateZkEVMVersion // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumUpdateZkEVMVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumUpdateZkEVMVersion)
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
		it.Event = new(IcdkvalidiumUpdateZkEVMVersion)
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
func (it *IcdkvalidiumUpdateZkEVMVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumUpdateZkEVMVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumUpdateZkEVMVersion represents a UpdateZkEVMVersion event raised by the Icdkvalidium contract.
type IcdkvalidiumUpdateZkEVMVersion struct {
	NumBatch uint64
	ForkID   uint64
	Version  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateZkEVMVersion is a free log retrieval operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterUpdateZkEVMVersion(opts *bind.FilterOpts) (*IcdkvalidiumUpdateZkEVMVersionIterator, error) {

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "UpdateZkEVMVersion")
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumUpdateZkEVMVersionIterator{contract: _Icdkvalidium.contract, event: "UpdateZkEVMVersion", logs: logs, sub: sub}, nil
}

// WatchUpdateZkEVMVersion is a free log subscription operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchUpdateZkEVMVersion(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumUpdateZkEVMVersion) (event.Subscription, error) {

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "UpdateZkEVMVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumUpdateZkEVMVersion)
				if err := _Icdkvalidium.contract.UnpackLog(event, "UpdateZkEVMVersion", log); err != nil {
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

// ParseUpdateZkEVMVersion is a log parse operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseUpdateZkEVMVersion(log types.Log) (*IcdkvalidiumUpdateZkEVMVersion, error) {
	event := new(IcdkvalidiumUpdateZkEVMVersion)
	if err := _Icdkvalidium.contract.UnpackLog(event, "UpdateZkEVMVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Icdkvalidium contract.
type IcdkvalidiumVerifyBatchesIterator struct {
	Event *IcdkvalidiumVerifyBatches // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumVerifyBatches)
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
		it.Event = new(IcdkvalidiumVerifyBatches)
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
func (it *IcdkvalidiumVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumVerifyBatches represents a VerifyBatches event raised by the Icdkvalidium contract.
type IcdkvalidiumVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*IcdkvalidiumVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumVerifyBatchesIterator{contract: _Icdkvalidium.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumVerifyBatches)
				if err := _Icdkvalidium.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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

// ParseVerifyBatches is a log parse operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseVerifyBatches(log types.Log) (*IcdkvalidiumVerifyBatches, error) {
	event := new(IcdkvalidiumVerifyBatches)
	if err := _Icdkvalidium.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IcdkvalidiumVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Icdkvalidium contract.
type IcdkvalidiumVerifyBatchesTrustedAggregatorIterator struct {
	Event *IcdkvalidiumVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *IcdkvalidiumVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IcdkvalidiumVerifyBatchesTrustedAggregator)
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
		it.Event = new(IcdkvalidiumVerifyBatchesTrustedAggregator)
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
func (it *IcdkvalidiumVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IcdkvalidiumVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IcdkvalidiumVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Icdkvalidium contract.
type IcdkvalidiumVerifyBatchesTrustedAggregator struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatchesTrustedAggregator is a free log retrieval operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Icdkvalidium *IcdkvalidiumFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*IcdkvalidiumVerifyBatchesTrustedAggregatorIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Icdkvalidium.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &IcdkvalidiumVerifyBatchesTrustedAggregatorIterator{contract: _Icdkvalidium.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Icdkvalidium *IcdkvalidiumFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *IcdkvalidiumVerifyBatchesTrustedAggregator, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Icdkvalidium.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IcdkvalidiumVerifyBatchesTrustedAggregator)
				if err := _Icdkvalidium.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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

// ParseVerifyBatchesTrustedAggregator is a log parse operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Icdkvalidium *IcdkvalidiumFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*IcdkvalidiumVerifyBatchesTrustedAggregator, error) {
	event := new(IcdkvalidiumVerifyBatchesTrustedAggregator)
	if err := _Icdkvalidium.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
