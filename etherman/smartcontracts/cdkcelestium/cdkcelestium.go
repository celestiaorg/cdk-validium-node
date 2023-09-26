// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cdkcelestium

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

// CDKCelestiumBatchData is an auto generated low-level Go binding around an user-defined struct.
type CDKCelestiumBatchData struct {
	TransactionsHash   [32]byte
	GlobalExitRoot     [32]byte
	Timestamp          uint64
	MinForcedTimestamp uint64
}

// CDKCelestiumForcedBatchData is an auto generated low-level Go binding around an user-defined struct.
type CDKCelestiumForcedBatchData struct {
	Transactions       []byte
	GlobalExitRoot     [32]byte
	MinForcedTimestamp uint64
}

// CDKCelestiumInitializePackedParameters is an auto generated low-level Go binding around an user-defined struct.
type CDKCelestiumInitializePackedParameters struct {
	Admin                    common.Address
	TrustedSequencer         common.Address
	PendingStateTimeout      uint64
	TrustedAggregator        common.Address
	TrustedAggregatorTimeout uint64
}

// CdkcelestiumMetaData contains all meta data concerning the Cdkcelestium contract.
var CdkcelestiumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_matic\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"_rollupVerifier\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_forkID\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ActivateForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"ConsolidatePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"OverridePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"provedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProveNonDeterministicPendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"SetMultiplierBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"SetPendingStateTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"SetTrustedAggregatorTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"SetVerifyBatchTimeTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"UpdateZkEVMVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequencedBatchNum\",\"type\":\"uint64\"}],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"}],\"name\":\"checkStateRootInsidePrime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"consolidatePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maticAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forkID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"trustedAggregatorTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structCDKCelestium.InitializePackedParameters\",\"name\":\"initializePackedParameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"genesisRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isForcedBatchDisallowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"isPendingStateConsolidable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingState\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingStateConsolidated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matic\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierBatchFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"overridePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStateTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingStateTransitions\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"proveNonDeterministicPendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifierRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structCDKCelestium.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structCDKCelestium.ForcedBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"sequencedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"setMultiplierBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"setPendingStateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"setTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"setTrustedAggregatorTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"setVerifyBatchTimeTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregatorTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyBatchTimeTarget\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b5060405162006058380380620060588339810160408190526200003591620000a5565b6001600160a01b0395861660c05293851660805291841660a05290921660e0526001600160401b0391821661010052166101205262000131565b6001600160a01b03811681146200008557600080fd5b50565b80516001600160401b0381168114620000a057600080fd5b919050565b60008060008060008060c08789031215620000bf57600080fd5b8651620000cc816200006f565b6020880151909650620000df816200006f565b6040880151909550620000f2816200006f565b606088015190945062000105816200006f565b9250620001156080880162000088565b91506200012560a0880162000088565b90509295509295509295565b60805160a05160c05160e0516101005161012051615e59620001ff6000396000818161069601528181610dec01526131410152600081816108030152610dc20152600081816107c901528181611b16015281816137da0152614c5601526000818161096f01528181610f5f015281816111300152818161174901528181612135015281816139c201526148b3015260008181610a1c0152818161407f01526144d70152600081816108bf01528181611ae40152818161262601528181613996015261416d0152615e596000f3fe608060405234801561001057600080fd5b50600436106103ba5760003560e01c8063841b24d7116101f4578063c754c7ed1161011a578063e7a7ed02116100ad578063f14916d61161007c578063f14916d614610a7e578063f2fde38b14610a91578063f851a44014610aa4578063f8b823e414610ac457600080fd5b8063e7a7ed02146109e7578063e8bf92ed14610a17578063eaeb077b14610a3e578063ed6b010414610a5157600080fd5b8063d2e129f9116100e9578063d2e129f914610991578063d8d1091b146109a4578063d939b315146109b7578063dbc16976146109df57600080fd5b8063c754c7ed146108fc578063c89e42df14610928578063cfa8ed471461093b578063d02103ca1461096a57600080fd5b8063a3c573eb11610192578063b4d63f5811610161578063b4d63f5814610853578063b6b0b097146108ba578063ba58ae39146108e1578063c0ed84e0146108f457600080fd5b8063a3c573eb146107c4578063ada8f919146107eb578063adc879e9146107fe578063afd23cbe1461082557600080fd5b806399f5634e116101ce57806399f5634e146107835780639aa972a31461078b5780639c9f3dfe1461079e578063a066215c146107b157600080fd5b8063841b24d71461072d5780638c3d73011461075d5780638da5cb5b1461076557600080fd5b8063458c0477116102e4578063621dd411116102775780637215541a116102465780637215541a1461066a5780637fcb36531461067d578063831c7ead14610691578063837a4738146106b857600080fd5b8063621dd4111461061c5780636b8616ce1461062f5780636ff512cc1461064f578063715018a61461066257600080fd5b80635392c5e0116102b35780635392c5e0146105d6578063542028d5146106045780635ec919581461060c578063604691691461061457600080fd5b8063458c04771461057c5780634a1a89a7146105905780634a910e6a146105b05780634e487706146105c357600080fd5b8063298789831161035c578063394218e91161032b578063394218e91461050e578063423fa85614610521578063438a539914610541578063456052671461055457600080fd5b806329878983146104a95780632b0006fa146104d55780632c1f816a146104e8578063383b3be8146104fb57600080fd5b80631816b7e5116103985780631816b7e51461042857806319d8ac611461043d578063220d789914610451578063267822471461046457600080fd5b80630a0d9fbe146103bf578063107bf28c146103f657806315064c961461040b575b600080fd5b606f546103d890610100900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b6103fe610acd565b6040516103ed9190615282565b606f546104189060ff1681565b60405190151581526020016103ed565b61043b61043636600461529c565b610b5b565b005b6073546103d89067ffffffffffffffff1681565b6103fe61045f3660046152d8565b610c73565b607b546104849073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016103ed565b6074546104849068010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b61043b6104e336600461533d565b610e4a565b61043b6104f63660046153a5565b61101a565b61041861050936600461541f565b611228565b61043b61051c36600461541f565b61127e565b6073546103d89068010000000000000000900467ffffffffffffffff1681565b61043b61054f3660046154a7565b611402565b6073546103d890700100000000000000000000000000000000900467ffffffffffffffff1681565b6079546103d89067ffffffffffffffff1681565b6079546103d89068010000000000000000900467ffffffffffffffff1681565b61043b6105be36600461541f565b611bd7565b61043b6105d136600461541f565b611c8a565b6105f66105e436600461541f565b60756020526000908152604090205481565b6040519081526020016103ed565b6103fe611e0e565b61043b611e1b565b6105f6611f1b565b61043b61062a36600461533d565b611f31565b6105f661063d36600461541f565b60716020526000908152604090205481565b61043b61065d366004615559565b6122b9565b61043b61238e565b61043b61067836600461541f565b6123a2565b6074546103d89067ffffffffffffffff1681565b6103d87f000000000000000000000000000000000000000000000000000000000000000081565b6107016106c6366004615574565b60786020526000908152604090208054600182015460029092015467ffffffffffffffff808316936801000000000000000090930416919084565b6040805167ffffffffffffffff95861681529490931660208501529183015260608201526080016103ed565b6079546103d8907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b61043b612512565b60335473ffffffffffffffffffffffffffffffffffffffff16610484565b6105f66125de565b61043b6107993660046153a5565b612737565b61043b6107ac36600461541f565b6127e8565b61043b6107bf36600461541f565b612964565b6104847f000000000000000000000000000000000000000000000000000000000000000081565b61043b6107f9366004615559565b612a6a565b6103d87f000000000000000000000000000000000000000000000000000000000000000081565b606f54610840906901000000000000000000900461ffff1681565b60405161ffff90911681526020016103ed565b61089461086136600461541f565b6072602052600090815260409020805460019091015467ffffffffffffffff808216916801000000000000000090041683565b6040805193845267ffffffffffffffff92831660208501529116908201526060016103ed565b6104847f000000000000000000000000000000000000000000000000000000000000000081565b6104186108ef366004615574565b612b2e565b6103d8612bb8565b607b546103d89074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b61043b610936366004615670565b612c0d565b606f54610484906b010000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b6104847f000000000000000000000000000000000000000000000000000000000000000081565b61043b61099f3660046156a5565b612c9a565b61043b6109b2366004615758565b6131e5565b6079546103d890700100000000000000000000000000000000900467ffffffffffffffff1681565b61043b613787565b6073546103d8907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b6104847f000000000000000000000000000000000000000000000000000000000000000081565b61043b610a4c3660046157cd565b613860565b607b54610418907c0100000000000000000000000000000000000000000000000000000000900460ff1681565b61043b610a8c366004615559565b613c56565b61043b610a9f366004615559565b613d28565b607a546104849073ffffffffffffffffffffffffffffffffffffffff1681565b6105f660705481565b60778054610ada90615819565b80601f0160208091040260200160405190810160405280929190818152602001828054610b0690615819565b8015610b535780601f10610b2857610100808354040283529160200191610b53565b820191906000526020600020905b815481529060010190602001808311610b3657829003601f168201915b505050505081565b607a5473ffffffffffffffffffffffffffffffffffffffff163314610bac576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88161ffff161080610bc557506103ff8161ffff16115b15610bfc576040517f4c2533c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffff0000ffffffffffffffffff16690100000000000000000061ffff8416908102919091179091556040519081527f7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5906020015b60405180910390a150565b67ffffffffffffffff8086166000818152607260205260408082205493881682529020546060929115801590610ca7575081155b15610cde576040517f6818c29e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610d15576040517f66385b5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d1e84612b2e565b610d54576040517f176b913c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080517fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003360601b166020820152603481019690965260548601929092527fffffffffffffffff00000000000000000000000000000000000000000000000060c098891b811660748701527f0000000000000000000000000000000000000000000000000000000000000000891b8116607c8701527f0000000000000000000000000000000000000000000000000000000000000000891b81166084870152608c86019490945260ac85015260cc840194909452509290931b90911660ec830152805180830360d401815260f4909201905290565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314610ea7576040517fbbcbbc0500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610eb5868686868686613ddc565b607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff86811691821790925560009081526075602052604090208390556079541615610f3057607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b158015610fb857600080fd5b505af1158015610fcc573d6000803e3d6000fd5b505060405184815233925067ffffffffffffffff871691507fcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe906020015b60405180910390a3505050505050565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314611077576040517fbbcbbc0500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611086878787878787876141a0565b607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169182179092556000908152607560205260409020839055607954161561110157607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b15801561118957600080fd5b505af115801561119d573d6000803e3d6000fd5b50506079805477ffffffffffffffffffffffffffffffffffffffffffffffff167a093a800000000000000000000000000000000000000000000000001790555050604051828152339067ffffffffffffffff8616907fcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf729060200160405180910390a350505050505050565b60795467ffffffffffffffff8281166000908152607860205260408120549092429261126c927001000000000000000000000000000000009092048116911661589b565b67ffffffffffffffff16111592915050565b607a5473ffffffffffffffffffffffffffffffffffffffff1633146112cf576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115611316576040517f1d06e87900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff166113855760795467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690821610611385576040517f401636df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6079805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527f1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a190602001610c68565b606f5460ff161561143f576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f546b010000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff16331461149f576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8360008190036114db576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115611517576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff6801000000000000000082048116600081815260726020526040812054838516949293700100000000000000000000000000000000909304909216919082905b868110156119305760008c8c8381811061157f5761157f6158c3565b90506080020180360381019061159591906158f2565b606081015190915067ffffffffffffffff161561170657846115b681615963565b955050600081600001518260200151836060015160405160200161161293929190928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff891660009081526071909352912054909150811461169b576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8087166000908152607160205260408082209190915560608401519084015190821691161015611700576040517f7f7ab87200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50611804565b6020810151158015906117cd575060208101516040517f257b363200000000000000000000000000000000000000000000000000000000815260048101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063257b3632906024016020604051808303816000875af11580156117a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117cb919061598a565b155b15611804576040517f73bd668d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8667ffffffffffffffff16816040015167ffffffffffffffff161080611837575042816040015167ffffffffffffffff16115b1561186e576040517fea82791600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b838160000151826020015183604001518e6040516020016118fd9594939291909485526020850193909352604084019190915260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060808401919091521b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166068820152607c0190565b60405160208183030381529060405280519060200120935080604001519650508080611928906159a3565b915050611563565b5061193b868561589b565b60735490945067ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690841611156119a4576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006119b082856159db565b6119c49067ffffffffffffffff16886159fc565b604080516060810182528581524267ffffffffffffffff908116602080840191825260738054680100000000000000009081900485168688019081528d861660008181526072909552979093209551865592516001909501805492519585167fffffffffffffffffffffffffffffffff000000000000000000000000000000009384161795851684029590951790945583548c8416911617930292909217905590915082811690851614611aba57607380547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8716021790555b611b0c333083607054611acd9190615a0f565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169291906145da565b611b146146bc565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611b7c57600080fd5b505af1158015611b90573d6000803e3d6000fd5b505060405167ffffffffffffffff881692507f303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce9150600090a2505050505050505050505050565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314611c7e57606f5460ff1615611c3f576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c4881611228565b611c7e576040517f0ce9e4a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c8781614769565b50565b607a5473ffffffffffffffffffffffffffffffffffffffff163314611cdb576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115611d22576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff16611d8d57607b5467ffffffffffffffff74010000000000000000000000000000000000000000909104811690821610611d8d576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b90602001610c68565b60768054610ada90615819565b607a5473ffffffffffffffffffffffffffffffffffffffff163314611e6c576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b547c0100000000000000000000000000000000000000000000000000000000900460ff16611ec8576040517ff6ba91a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690556040517f854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f90600090a1565b60006070546064611f2c9190615a0f565b905090565b606f5460ff1615611f6e576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795467ffffffffffffffff8581166000908152607260205260409020600101544292611fbb9278010000000000000000000000000000000000000000000000009091048116911661589b565b67ffffffffffffffff161115611ffd576040517f8a0704d300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e861200a86866159db565b67ffffffffffffffff16111561204c576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61205a868686868686613ddc565b6120638461497c565b607954700100000000000000000000000000000000900467ffffffffffffffff166000036121ab57607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169182179092556000908152607560205260409020839055607954161561210657607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b15801561218e57600080fd5b505af11580156121a2573d6000803e3d6000fd5b5050505061227b565b6121b36146bc565b6079805467ffffffffffffffff169060006121cd83615963565b825467ffffffffffffffff9182166101009390930a92830292820219169190911790915560408051608081018252428316815287831660208083019182528284018981526060840189815260795487166000908152607890935294909120925183549251861668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169516949094171781559151600183015551600290910155505b604051828152339067ffffffffffffffff8616907f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f59669060200161100a565b607a5473ffffffffffffffffffffffffffffffffffffffff16331461230a576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fff0000000000000000000000000000000000000000ffffffffffffffffffffff166b01000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610c68565b612396614b5c565b6123a06000614bdd565b565b60335473ffffffffffffffffffffffffffffffffffffffff16331461250a5760006123cb612bb8565b90508067ffffffffffffffff168267ffffffffffffffff161161241a576040517f812a372d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff6801000000000000000090910481169083161180612460575067ffffffffffffffff80831660009081526072602052604090206001015416155b15612497576040517f98c5c01400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80831660009081526072602052604090206001015442916124c69162093a80911661589b565b67ffffffffffffffff161115612508576040517fd257555a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b611c87614c54565b607b5473ffffffffffffffffffffffffffffffffffffffff163314612563576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b54607a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600090819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561266d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612691919061598a565b9050600061269d612bb8565b60735467ffffffffffffffff6801000000000000000082048116916126f591700100000000000000000000000000000000820481169178010000000000000000000000000000000000000000000000009004166159db565b6126ff919061589b565b61270991906159db565b67ffffffffffffffff169050806000036127265760009250505090565b6127308183615a55565b9250505090565b606f5460ff1615612774576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612783878787878787876141a0565b67ffffffffffffffff84166000908152607560209081526040918290205482519081529081018490527f1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010910160405180910390a16127df614c54565b50505050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612839576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115612880576040517fcc96507000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff166128e75760795467ffffffffffffffff7001000000000000000000000000000000009091048116908216106128e7576040517f48a05a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607980547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c7590602001610c68565b607a5473ffffffffffffffffffffffffffffffffffffffff1633146129b5576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620151808167ffffffffffffffff1611156129fc576040517fe067dfe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff1661010067ffffffffffffffff8416908102919091179091556040519081527f1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c2890602001610c68565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612abb576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610c68565b600067ffffffff0000000167ffffffffffffffff8316108015612b66575067ffffffff00000001604083901c67ffffffffffffffff16105b8015612b87575067ffffffff00000001608083901c67ffffffffffffffff16105b8015612b9e575067ffffffff0000000160c083901c105b15612bab57506001919050565b506000919050565b919050565b60795460009067ffffffffffffffff1615612bfc575060795467ffffffffffffffff9081166000908152607860205260409020546801000000000000000090041690565b5060745467ffffffffffffffff1690565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612c5e576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6076612c6a8282615ab7565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610c689190615282565b600054610100900460ff1615808015612cba5750600054600160ff909116105b80612cd45750303b158015612cd4575060005460ff166001145b612d65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015612dc357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b612dd06020880188615559565b607a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055612e256040880160208901615559565b606f805473ffffffffffffffffffffffffffffffffffffffff929092166b010000000000000000000000027fff0000000000000000000000000000000000000000ffffffffffffffffffffff909216919091179055612e8a6080880160608901615559565b6074805473ffffffffffffffffffffffffffffffffffffffff9290921668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff9092169190911790556000805260756020527ff9e3fbf150b7a0077118526f473c53cb4734f166167e2c6213e3567dd390b4ad8690556076612f158682615ab7565b506077612f228582615ab7565b5062093a80612f376060890160408a0161541f565b67ffffffffffffffff161115612f79576040517fcc96507000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f89606088016040890161541f565b6079805467ffffffffffffffff92909216700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff90921691909117905562093a80612feb60a0890160808a0161541f565b67ffffffffffffffff16111561302d576040517f1d06e87900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61303d60a088016080890161541f565b6079805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff939093169290920291909117905567016345785d8a0000607055606f80547fffffffffffffffffffffffffffffffffffffffffff00000000000000000000ff166a03ea000000000000070800179055607b80547fffffff000000000000000000ffffffffffffffffffffffffffffffffffffffff167c010000000000069780000000000000000000000000000000000000000017905561311c614cdc565b7fed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd660007f000000000000000000000000000000000000000000000000000000000000000085856040516131729493929190615c1a565b60405180910390a180156127df57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050505050565b607b547c0100000000000000000000000000000000000000000000000000000000900460ff1615613242576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff161561327f576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060008190036132bb576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156132f7576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff78010000000000000000000000000000000000000000000000008204811691613342918491700100000000000000000000000000000000900416615c52565b111561337a576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff680100000000000000008204811660008181526072602052604081205491937001000000000000000000000000000000009004909216915b848110156136245760008787838181106133da576133da6158c3565b90506020028101906133ec9190615c65565b6133f590615ca3565b90508361340181615963565b825180516020918201208185015160408087015190519499509194506000936134639386939101928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89166000908152607190935291205490915081146134ec576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff86166000908152607160205260408120556135116001896159fc565b84036135805742607b60149054906101000a900467ffffffffffffffff16846040015161353e919061589b565b67ffffffffffffffff161115613580576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020838101516040805192830188905282018490526060808301919091524260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016608083015233901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c01604051602081830303815290604052805190602001209450505050808061361c906159a3565b9150506133be565b5061362f848461589b565b6073805467ffffffffffffffff4281167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009092168217808455604080516060810182528781526020808201958652680100000000000000009384900485168284019081528589166000818152607290935284832093518455965160019390930180549151871686027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090921693871693909317179091558554938916700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff938602939093167fffffffffffffffff00000000000000000000000000000000ffffffffffffffff90941693909317919091179093559151929550917f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a49190a2505050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff1633146137d8576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dbc169766040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561384057600080fd5b505af1158015613854573d6000803e3d6000fd5b505050506123a0614d7c565b607b547c0100000000000000000000000000000000000000000000000000000000900460ff16156138bd576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff16156138fa576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000613904611f1b565b905081811115613940576040517f4732fdb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61138883111561397c576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6139be73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163330846145da565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613a2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a4f919061598a565b60738054919250780100000000000000000000000000000000000000000000000090910467ffffffffffffffff16906018613a8983615963565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508484604051613ac0929190615d33565b60408051918290038220602083015281018290527fffffffffffffffff0000000000000000000000000000000000000000000000004260c01b166060820152606801604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815291815281516020928301206073547801000000000000000000000000000000000000000000000000900467ffffffffffffffff1660009081526071909352912055323303613bf057607354604080518381523360208201526060918101829052600091810191909152780100000000000000000000000000000000000000000000000090910467ffffffffffffffff16907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319060800160405180910390a2613c4f565b607360189054906101000a900467ffffffffffffffff1667ffffffffffffffff167ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc93182338888604051613c469493929190615d43565b60405180910390a25b5050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314613ca7576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607480547fffffffff0000000000000000000000000000000000000000ffffffffffffffff166801000000000000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca90602001610c68565b613d30614b5c565b73ffffffffffffffffffffffffffffffffffffffff8116613dd3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401612d5c565b611c8781614bdd565b600080613de7612bb8565b905067ffffffffffffffff881615613eb75760795467ffffffffffffffff9081169089161115613e43576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8089166000908152607860205260409020600281015481549094509091898116680100000000000000009092041614613eb1576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50613f58565b67ffffffffffffffff8716600090815260756020526040902054915081613f0a576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168767ffffffffffffffff161115613f58576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168667ffffffffffffffff1611613fa5576040517fb9b18f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000613fb48888888689610c73565b905060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001600283604051613fe99190615d79565b602060405180830381855afa158015614006573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190614029919061598a565b6140339190615d8b565b6040805160208101825282815290517f9121da8a00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691639121da8a916140b591899190600401615d9f565b602060405180830381865afa1580156140d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906140f69190615dda565b61412c576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6141943361413a858b6159db565b67ffffffffffffffff1661414c6125de565b6141569190615a0f565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190614e0b565b50505050505050505050565b600067ffffffffffffffff88161561426e5760795467ffffffffffffffff90811690891611156141fc576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff8088166000908152607860205260409020600281015481549092888116680100000000000000009092041614614268576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061430a565b5067ffffffffffffffff8516600090815260756020526040902054806142c0576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60745467ffffffffffffffff908116908716111561430a576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795467ffffffffffffffff908116908816118061433c57508767ffffffffffffffff168767ffffffffffffffff1611155b80614363575060795467ffffffffffffffff68010000000000000000909104811690881611155b1561439a576040517fbfa7079f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8781166000908152607860205260409020546801000000000000000090048116908616146143fd576040517f32a2a77f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061440c8787878588610c73565b905060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016002836040516144419190615d79565b602060405180830381855afa15801561445e573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190614481919061598a565b61448b9190615d8b565b6040805160208101825282815290517f9121da8a00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691639121da8a9161450d91889190600401615d9f565b602060405180830381865afa15801561452a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061454e9190615dda565b614584576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8916600090815260786020526040902060020154859003614194576040517fa47276bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526146b69085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152614e66565b50505050565b60795467ffffffffffffffff6801000000000000000082048116911611156123a0576079546000906147059068010000000000000000900467ffffffffffffffff16600161589b565b905061471081611228565b15611c875760795460009060029061473390849067ffffffffffffffff166159db565b61473d9190615dfc565b614747908361589b565b905061475281611228565b156147645761476081614769565b5050565b614760825b60795467ffffffffffffffff6801000000000000000090910481169082161115806147a3575060795467ffffffffffffffff908116908216115b156147da576040517fd086b70b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff818116600081815260786020908152604080832080546074805468010000000000000000928390049098167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000090981688179055600282015487865260759094529382902092909255607980547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff169390940292909217909255600182015490517f33d6247d00000000000000000000000000000000000000000000000000000000815260048101919091529091907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b15801561490c57600080fd5b505af1158015614920573d6000803e3d6000fd5b505050508267ffffffffffffffff168167ffffffffffffffff167f328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e846002015460405161496f91815260200190565b60405180910390a3505050565b6000614986612bb8565b90508160008061499684846159db565b606f5467ffffffffffffffff91821692506000916149ba91610100900416426159fc565b90505b8467ffffffffffffffff168467ffffffffffffffff1614614a455767ffffffffffffffff80851660009081526072602052604090206001810154909116821015614a2357600181015468010000000000000000900467ffffffffffffffff169450614a3f565b614a2d86866159db565b67ffffffffffffffff16935050614a45565b506149bd565b6000614a5184846159fc565b905083811015614aa857808403600c8111614a6c5780614a6f565b600c5b9050806103e80a81606f60099054906101000a900461ffff1661ffff160a6070540281614a9e57614a9e615a26565b0460705550614b18565b838103600c8111614ab95780614abc565b600c5b90506000816103e80a82606f60099054906101000a900461ffff1661ffff160a670de0b6b3a76400000281614af357614af3615a26565b04905080607054670de0b6b3a76400000281614b1157614b11615a26565b0460705550505b683635c9adc5dea000006070541115614b3d57683635c9adc5dea000006070556127df565b633b9aca0060705410156127df57633b9aca0060705550505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146123a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401612d5c565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632072f6c56040518163ffffffff1660e01b8152600401600060405180830381600087803b158015614cbc57600080fd5b505af1158015614cd0573d6000803e3d6000fd5b505050506123a0614f72565b600054610100900460ff16614d73576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401612d5c565b6123a033614bdd565b606f5460ff16614db8576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b390600090a1565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052614e619084907fa9059cbb0000000000000000000000000000000000000000000000000000000090606401614634565b505050565b6000614ec8826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166150059092919063ffffffff16565b805190915015614e615780806020019051810190614ee69190615dda565b614e61576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401612d5c565b606f5460ff1615614faf576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a549790600090a1565b6060615014848460008561501c565b949350505050565b6060824710156150ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401612d5c565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516150d79190615d79565b60006040518083038185875af1925050503d8060008114615114576040519150601f19603f3d011682016040523d82523d6000602084013e615119565b606091505b509150915061512a87838387615135565b979650505050505050565b606083156151cb5782516000036151c45773ffffffffffffffffffffffffffffffffffffffff85163b6151c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401612d5c565b5081615014565b61501483838151156151e05781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d5c9190615282565b60005b8381101561522f578181015183820152602001615217565b50506000910152565b60008151808452615250816020860160208601615214565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006152956020830184615238565b9392505050565b6000602082840312156152ae57600080fd5b813561ffff8116811461529557600080fd5b803567ffffffffffffffff81168114612bb357600080fd5b600080600080600060a086880312156152f057600080fd5b6152f9866152c0565b9450615307602087016152c0565b94979496505050506040830135926060810135926080909101359150565b80610300810183101561533757600080fd5b92915050565b6000806000806000806103a0878903121561535757600080fd5b615360876152c0565b955061536e602088016152c0565b945061537c604088016152c0565b935060608701359250608087013591506153998860a08901615325565b90509295509295509295565b60008060008060008060006103c0888a0312156153c157600080fd5b6153ca886152c0565b96506153d8602089016152c0565b95506153e6604089016152c0565b94506153f4606089016152c0565b93506080880135925060a088013591506154118960c08a01615325565b905092959891949750929550565b60006020828403121561543157600080fd5b615295826152c0565b803573ffffffffffffffffffffffffffffffffffffffff81168114612bb357600080fd5b60008083601f84011261547057600080fd5b50813567ffffffffffffffff81111561548857600080fd5b6020830191508360208285010111156154a057600080fd5b9250929050565b6000806000806000606086880312156154bf57600080fd5b853567ffffffffffffffff808211156154d757600080fd5b818801915088601f8301126154eb57600080fd5b8135818111156154fa57600080fd5b8960208260071b850101111561550f57600080fd5b602083019750809650506155256020890161543a565b9450604088013591508082111561553b57600080fd5b506155488882890161545e565b969995985093965092949392505050565b60006020828403121561556b57600080fd5b6152958261543a565b60006020828403121561558657600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600067ffffffffffffffff808411156155d7576155d761558d565b604051601f85017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561561d5761561d61558d565b8160405280935085815286868601111561563657600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011261566157600080fd5b615295838335602085016155bc565b60006020828403121561568257600080fd5b813567ffffffffffffffff81111561569957600080fd5b61501484828501615650565b6000806000806000808688036101208112156156c057600080fd5b60a08112156156ce57600080fd5b5086955060a0870135945060c087013567ffffffffffffffff808211156156f457600080fd5b6157008a838b01615650565b955060e089013591508082111561571657600080fd5b6157228a838b01615650565b945061010089013591508082111561573957600080fd5b5061574689828a0161545e565b979a9699509497509295939492505050565b6000806020838503121561576b57600080fd5b823567ffffffffffffffff8082111561578357600080fd5b818501915085601f83011261579757600080fd5b8135818111156157a657600080fd5b8660208260051b85010111156157bb57600080fd5b60209290920196919550909350505050565b6000806000604084860312156157e257600080fd5b833567ffffffffffffffff8111156157f957600080fd5b6158058682870161545e565b909790965060209590950135949350505050565b600181811c9082168061582d57607f821691505b602082108103615866577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff8181168382160190808211156158bc576158bc61586c565b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006080828403121561590457600080fd5b6040516080810181811067ffffffffffffffff821117156159275761592761558d565b80604052508235815260208301356020820152615946604084016152c0565b6040820152615957606084016152c0565b60608201529392505050565b600067ffffffffffffffff8083168181036159805761598061586c565b6001019392505050565b60006020828403121561599c57600080fd5b5051919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036159d4576159d461586c565b5060010190565b67ffffffffffffffff8281168282160390808211156158bc576158bc61586c565b818103818111156153375761533761586c565b80820281158282048414176153375761533761586c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082615a6457615a64615a26565b500490565b601f821115614e6157600081815260208120601f850160051c81016020861015615a905750805b601f850160051c820191505b81811015615aaf57828155600101615a9c565b505050505050565b815167ffffffffffffffff811115615ad157615ad161558d565b615ae581615adf8454615819565b84615a69565b602080601f831160018114615b385760008415615b025750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555615aaf565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015615b8557888601518255948401946001909101908401615b66565b5085821015615bc157878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600067ffffffffffffffff808716835280861660208401525060606040830152615c48606083018486615bd1565b9695505050505050565b808201808211156153375761533761586c565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112615c9957600080fd5b9190910192915050565b600060608236031215615cb557600080fd5b6040516060810167ffffffffffffffff8282108183111715615cd957615cd961558d565b816040528435915080821115615cee57600080fd5b50830136601f820112615d0057600080fd5b615d0f368235602084016155bc565b82525060208301356020820152615d28604084016152c0565b604082015292915050565b8183823760009101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201526000615c48606083018486615bd1565b60008251615c99818460208701615214565b600082615d9a57615d9a615a26565b500690565b61032081016103008085843782018360005b6001811015615dd0578151835260209283019290910190600101615db1565b5050509392505050565b600060208284031215615dec57600080fd5b8151801515811461529557600080fd5b600067ffffffffffffffff80841680615e1757615e17615a26565b9216919091049291505056fea264697066735822122019e5e2f1b888da74cdab2300a524e95fe22e43a92b12a1dfcd5bbb6d0174a4ec64736f6c63430008140033",
}

// CdkcelestiumABI is the input ABI used to generate the binding from.
// Deprecated: Use CdkcelestiumMetaData.ABI instead.
var CdkcelestiumABI = CdkcelestiumMetaData.ABI

// CdkcelestiumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CdkcelestiumMetaData.Bin instead.
var CdkcelestiumBin = CdkcelestiumMetaData.Bin

// DeployCdkcelestium deploys a new Ethereum contract, binding an instance of Cdkcelestium to it.
func DeployCdkcelestium(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _matic common.Address, _rollupVerifier common.Address, _bridgeAddress common.Address, _chainID uint64, _forkID uint64) (common.Address, *types.Transaction, *Cdkcelestium, error) {
	parsed, err := CdkcelestiumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CdkcelestiumBin), backend, _globalExitRootManager, _matic, _rollupVerifier, _bridgeAddress, _chainID, _forkID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cdkcelestium{CdkcelestiumCaller: CdkcelestiumCaller{contract: contract}, CdkcelestiumTransactor: CdkcelestiumTransactor{contract: contract}, CdkcelestiumFilterer: CdkcelestiumFilterer{contract: contract}}, nil
}

// Cdkcelestium is an auto generated Go binding around an Ethereum contract.
type Cdkcelestium struct {
	CdkcelestiumCaller     // Read-only binding to the contract
	CdkcelestiumTransactor // Write-only binding to the contract
	CdkcelestiumFilterer   // Log filterer for contract events
}

// CdkcelestiumCaller is an auto generated read-only Go binding around an Ethereum contract.
type CdkcelestiumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdkcelestiumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CdkcelestiumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdkcelestiumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CdkcelestiumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdkcelestiumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CdkcelestiumSession struct {
	Contract     *Cdkcelestium     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CdkcelestiumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CdkcelestiumCallerSession struct {
	Contract *CdkcelestiumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CdkcelestiumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CdkcelestiumTransactorSession struct {
	Contract     *CdkcelestiumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CdkcelestiumRaw is an auto generated low-level Go binding around an Ethereum contract.
type CdkcelestiumRaw struct {
	Contract *Cdkcelestium // Generic contract binding to access the raw methods on
}

// CdkcelestiumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CdkcelestiumCallerRaw struct {
	Contract *CdkcelestiumCaller // Generic read-only contract binding to access the raw methods on
}

// CdkcelestiumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CdkcelestiumTransactorRaw struct {
	Contract *CdkcelestiumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCdkcelestium creates a new instance of Cdkcelestium, bound to a specific deployed contract.
func NewCdkcelestium(address common.Address, backend bind.ContractBackend) (*Cdkcelestium, error) {
	contract, err := bindCdkcelestium(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cdkcelestium{CdkcelestiumCaller: CdkcelestiumCaller{contract: contract}, CdkcelestiumTransactor: CdkcelestiumTransactor{contract: contract}, CdkcelestiumFilterer: CdkcelestiumFilterer{contract: contract}}, nil
}

// NewCdkcelestiumCaller creates a new read-only instance of Cdkcelestium, bound to a specific deployed contract.
func NewCdkcelestiumCaller(address common.Address, caller bind.ContractCaller) (*CdkcelestiumCaller, error) {
	contract, err := bindCdkcelestium(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumCaller{contract: contract}, nil
}

// NewCdkcelestiumTransactor creates a new write-only instance of Cdkcelestium, bound to a specific deployed contract.
func NewCdkcelestiumTransactor(address common.Address, transactor bind.ContractTransactor) (*CdkcelestiumTransactor, error) {
	contract, err := bindCdkcelestium(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumTransactor{contract: contract}, nil
}

// NewCdkcelestiumFilterer creates a new log filterer instance of Cdkcelestium, bound to a specific deployed contract.
func NewCdkcelestiumFilterer(address common.Address, filterer bind.ContractFilterer) (*CdkcelestiumFilterer, error) {
	contract, err := bindCdkcelestium(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumFilterer{contract: contract}, nil
}

// bindCdkcelestium binds a generic wrapper to an already deployed contract.
func bindCdkcelestium(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CdkcelestiumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cdkcelestium *CdkcelestiumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cdkcelestium.Contract.CdkcelestiumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cdkcelestium *CdkcelestiumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.CdkcelestiumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cdkcelestium *CdkcelestiumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.CdkcelestiumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cdkcelestium *CdkcelestiumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cdkcelestium.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cdkcelestium *CdkcelestiumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cdkcelestium *CdkcelestiumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Cdkcelestium *CdkcelestiumCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Cdkcelestium *CdkcelestiumSession) Admin() (common.Address, error) {
	return _Cdkcelestium.Contract.Admin(&_Cdkcelestium.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Cdkcelestium *CdkcelestiumCallerSession) Admin() (common.Address, error) {
	return _Cdkcelestium.Contract.Admin(&_Cdkcelestium.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Cdkcelestium *CdkcelestiumCaller) BatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "batchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Cdkcelestium *CdkcelestiumSession) BatchFee() (*big.Int, error) {
	return _Cdkcelestium.Contract.BatchFee(&_Cdkcelestium.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Cdkcelestium *CdkcelestiumCallerSession) BatchFee() (*big.Int, error) {
	return _Cdkcelestium.Contract.BatchFee(&_Cdkcelestium.CallOpts)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Cdkcelestium *CdkcelestiumCaller) BatchNumToStateRoot(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "batchNumToStateRoot", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Cdkcelestium *CdkcelestiumSession) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Cdkcelestium.Contract.BatchNumToStateRoot(&_Cdkcelestium.CallOpts, arg0)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Cdkcelestium *CdkcelestiumCallerSession) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Cdkcelestium.Contract.BatchNumToStateRoot(&_Cdkcelestium.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Cdkcelestium *CdkcelestiumCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Cdkcelestium *CdkcelestiumSession) BridgeAddress() (common.Address, error) {
	return _Cdkcelestium.Contract.BridgeAddress(&_Cdkcelestium.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Cdkcelestium *CdkcelestiumCallerSession) BridgeAddress() (common.Address, error) {
	return _Cdkcelestium.Contract.BridgeAddress(&_Cdkcelestium.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Cdkcelestium *CdkcelestiumCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Cdkcelestium *CdkcelestiumSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Cdkcelestium.Contract.CalculateRewardPerBatch(&_Cdkcelestium.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Cdkcelestium *CdkcelestiumCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Cdkcelestium.Contract.CalculateRewardPerBatch(&_Cdkcelestium.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) ChainID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) ChainID() (uint64, error) {
	return _Cdkcelestium.Contract.ChainID(&_Cdkcelestium.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) ChainID() (uint64, error) {
	return _Cdkcelestium.Contract.ChainID(&_Cdkcelestium.CallOpts)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Cdkcelestium *CdkcelestiumCaller) CheckStateRootInsidePrime(opts *bind.CallOpts, newStateRoot *big.Int) (bool, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "checkStateRootInsidePrime", newStateRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Cdkcelestium *CdkcelestiumSession) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Cdkcelestium.Contract.CheckStateRootInsidePrime(&_Cdkcelestium.CallOpts, newStateRoot)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Cdkcelestium *CdkcelestiumCallerSession) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Cdkcelestium.Contract.CheckStateRootInsidePrime(&_Cdkcelestium.CallOpts, newStateRoot)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) ForceBatchTimeout() (uint64, error) {
	return _Cdkcelestium.Contract.ForceBatchTimeout(&_Cdkcelestium.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Cdkcelestium.Contract.ForceBatchTimeout(&_Cdkcelestium.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Cdkcelestium *CdkcelestiumCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Cdkcelestium *CdkcelestiumSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Cdkcelestium.Contract.ForcedBatches(&_Cdkcelestium.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Cdkcelestium *CdkcelestiumCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Cdkcelestium.Contract.ForcedBatches(&_Cdkcelestium.CallOpts, arg0)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) ForkID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "forkID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) ForkID() (uint64, error) {
	return _Cdkcelestium.Contract.ForkID(&_Cdkcelestium.CallOpts)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) ForkID() (uint64, error) {
	return _Cdkcelestium.Contract.ForkID(&_Cdkcelestium.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Cdkcelestium *CdkcelestiumCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Cdkcelestium *CdkcelestiumSession) GetForcedBatchFee() (*big.Int, error) {
	return _Cdkcelestium.Contract.GetForcedBatchFee(&_Cdkcelestium.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Cdkcelestium *CdkcelestiumCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Cdkcelestium.Contract.GetForcedBatchFee(&_Cdkcelestium.CallOpts)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Cdkcelestium *CdkcelestiumCaller) GetInputSnarkBytes(opts *bind.CallOpts, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "getInputSnarkBytes", initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Cdkcelestium *CdkcelestiumSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Cdkcelestium.Contract.GetInputSnarkBytes(&_Cdkcelestium.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Cdkcelestium *CdkcelestiumCallerSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Cdkcelestium.Contract.GetInputSnarkBytes(&_Cdkcelestium.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) GetLastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "getLastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) GetLastVerifiedBatch() (uint64, error) {
	return _Cdkcelestium.Contract.GetLastVerifiedBatch(&_Cdkcelestium.CallOpts)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) GetLastVerifiedBatch() (uint64, error) {
	return _Cdkcelestium.Contract.GetLastVerifiedBatch(&_Cdkcelestium.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Cdkcelestium *CdkcelestiumCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Cdkcelestium *CdkcelestiumSession) GlobalExitRootManager() (common.Address, error) {
	return _Cdkcelestium.Contract.GlobalExitRootManager(&_Cdkcelestium.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Cdkcelestium *CdkcelestiumCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Cdkcelestium.Contract.GlobalExitRootManager(&_Cdkcelestium.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Cdkcelestium *CdkcelestiumCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Cdkcelestium *CdkcelestiumSession) IsEmergencyState() (bool, error) {
	return _Cdkcelestium.Contract.IsEmergencyState(&_Cdkcelestium.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Cdkcelestium *CdkcelestiumCallerSession) IsEmergencyState() (bool, error) {
	return _Cdkcelestium.Contract.IsEmergencyState(&_Cdkcelestium.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Cdkcelestium *CdkcelestiumCaller) IsForcedBatchDisallowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "isForcedBatchDisallowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Cdkcelestium *CdkcelestiumSession) IsForcedBatchDisallowed() (bool, error) {
	return _Cdkcelestium.Contract.IsForcedBatchDisallowed(&_Cdkcelestium.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Cdkcelestium *CdkcelestiumCallerSession) IsForcedBatchDisallowed() (bool, error) {
	return _Cdkcelestium.Contract.IsForcedBatchDisallowed(&_Cdkcelestium.CallOpts)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Cdkcelestium *CdkcelestiumCaller) IsPendingStateConsolidable(opts *bind.CallOpts, pendingStateNum uint64) (bool, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "isPendingStateConsolidable", pendingStateNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Cdkcelestium *CdkcelestiumSession) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Cdkcelestium.Contract.IsPendingStateConsolidable(&_Cdkcelestium.CallOpts, pendingStateNum)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Cdkcelestium *CdkcelestiumCallerSession) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Cdkcelestium.Contract.IsPendingStateConsolidable(&_Cdkcelestium.CallOpts, pendingStateNum)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) LastBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "lastBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) LastBatchSequenced() (uint64, error) {
	return _Cdkcelestium.Contract.LastBatchSequenced(&_Cdkcelestium.CallOpts)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) LastBatchSequenced() (uint64, error) {
	return _Cdkcelestium.Contract.LastBatchSequenced(&_Cdkcelestium.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) LastForceBatch() (uint64, error) {
	return _Cdkcelestium.Contract.LastForceBatch(&_Cdkcelestium.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) LastForceBatch() (uint64, error) {
	return _Cdkcelestium.Contract.LastForceBatch(&_Cdkcelestium.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) LastForceBatchSequenced() (uint64, error) {
	return _Cdkcelestium.Contract.LastForceBatchSequenced(&_Cdkcelestium.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Cdkcelestium.Contract.LastForceBatchSequenced(&_Cdkcelestium.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) LastPendingState(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "lastPendingState")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) LastPendingState() (uint64, error) {
	return _Cdkcelestium.Contract.LastPendingState(&_Cdkcelestium.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) LastPendingState() (uint64, error) {
	return _Cdkcelestium.Contract.LastPendingState(&_Cdkcelestium.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) LastPendingStateConsolidated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "lastPendingStateConsolidated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) LastPendingStateConsolidated() (uint64, error) {
	return _Cdkcelestium.Contract.LastPendingStateConsolidated(&_Cdkcelestium.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) LastPendingStateConsolidated() (uint64, error) {
	return _Cdkcelestium.Contract.LastPendingStateConsolidated(&_Cdkcelestium.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) LastTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "lastTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) LastTimestamp() (uint64, error) {
	return _Cdkcelestium.Contract.LastTimestamp(&_Cdkcelestium.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) LastTimestamp() (uint64, error) {
	return _Cdkcelestium.Contract.LastTimestamp(&_Cdkcelestium.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) LastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "lastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) LastVerifiedBatch() (uint64, error) {
	return _Cdkcelestium.Contract.LastVerifiedBatch(&_Cdkcelestium.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) LastVerifiedBatch() (uint64, error) {
	return _Cdkcelestium.Contract.LastVerifiedBatch(&_Cdkcelestium.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Cdkcelestium *CdkcelestiumCaller) Matic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "matic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Cdkcelestium *CdkcelestiumSession) Matic() (common.Address, error) {
	return _Cdkcelestium.Contract.Matic(&_Cdkcelestium.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Cdkcelestium *CdkcelestiumCallerSession) Matic() (common.Address, error) {
	return _Cdkcelestium.Contract.Matic(&_Cdkcelestium.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Cdkcelestium *CdkcelestiumCaller) MultiplierBatchFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "multiplierBatchFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Cdkcelestium *CdkcelestiumSession) MultiplierBatchFee() (uint16, error) {
	return _Cdkcelestium.Contract.MultiplierBatchFee(&_Cdkcelestium.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Cdkcelestium *CdkcelestiumCallerSession) MultiplierBatchFee() (uint16, error) {
	return _Cdkcelestium.Contract.MultiplierBatchFee(&_Cdkcelestium.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Cdkcelestium *CdkcelestiumCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Cdkcelestium *CdkcelestiumSession) NetworkName() (string, error) {
	return _Cdkcelestium.Contract.NetworkName(&_Cdkcelestium.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Cdkcelestium *CdkcelestiumCallerSession) NetworkName() (string, error) {
	return _Cdkcelestium.Contract.NetworkName(&_Cdkcelestium.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cdkcelestium *CdkcelestiumCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cdkcelestium *CdkcelestiumSession) Owner() (common.Address, error) {
	return _Cdkcelestium.Contract.Owner(&_Cdkcelestium.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cdkcelestium *CdkcelestiumCallerSession) Owner() (common.Address, error) {
	return _Cdkcelestium.Contract.Owner(&_Cdkcelestium.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Cdkcelestium *CdkcelestiumCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Cdkcelestium *CdkcelestiumSession) PendingAdmin() (common.Address, error) {
	return _Cdkcelestium.Contract.PendingAdmin(&_Cdkcelestium.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Cdkcelestium *CdkcelestiumCallerSession) PendingAdmin() (common.Address, error) {
	return _Cdkcelestium.Contract.PendingAdmin(&_Cdkcelestium.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) PendingStateTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "pendingStateTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) PendingStateTimeout() (uint64, error) {
	return _Cdkcelestium.Contract.PendingStateTimeout(&_Cdkcelestium.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) PendingStateTimeout() (uint64, error) {
	return _Cdkcelestium.Contract.PendingStateTimeout(&_Cdkcelestium.CallOpts)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Cdkcelestium *CdkcelestiumCaller) PendingStateTransitions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "pendingStateTransitions", arg0)

	outstruct := new(struct {
		Timestamp         uint64
		LastVerifiedBatch uint64
		ExitRoot          [32]byte
		StateRoot         [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.LastVerifiedBatch = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ExitRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StateRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Cdkcelestium *CdkcelestiumSession) PendingStateTransitions(arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	return _Cdkcelestium.Contract.PendingStateTransitions(&_Cdkcelestium.CallOpts, arg0)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Cdkcelestium *CdkcelestiumCallerSession) PendingStateTransitions(arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	return _Cdkcelestium.Contract.PendingStateTransitions(&_Cdkcelestium.CallOpts, arg0)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Cdkcelestium *CdkcelestiumCaller) RollupVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "rollupVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Cdkcelestium *CdkcelestiumSession) RollupVerifier() (common.Address, error) {
	return _Cdkcelestium.Contract.RollupVerifier(&_Cdkcelestium.CallOpts)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Cdkcelestium *CdkcelestiumCallerSession) RollupVerifier() (common.Address, error) {
	return _Cdkcelestium.Contract.RollupVerifier(&_Cdkcelestium.CallOpts)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Cdkcelestium *CdkcelestiumCaller) SequencedBatches(opts *bind.CallOpts, arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "sequencedBatches", arg0)

	outstruct := new(struct {
		AccInputHash               [32]byte
		SequencedTimestamp         uint64
		PreviousLastBatchSequenced uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccInputHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.SequencedTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.PreviousLastBatchSequenced = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Cdkcelestium *CdkcelestiumSession) SequencedBatches(arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	return _Cdkcelestium.Contract.SequencedBatches(&_Cdkcelestium.CallOpts, arg0)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Cdkcelestium *CdkcelestiumCallerSession) SequencedBatches(arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	return _Cdkcelestium.Contract.SequencedBatches(&_Cdkcelestium.CallOpts, arg0)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Cdkcelestium *CdkcelestiumCaller) TrustedAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "trustedAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Cdkcelestium *CdkcelestiumSession) TrustedAggregator() (common.Address, error) {
	return _Cdkcelestium.Contract.TrustedAggregator(&_Cdkcelestium.CallOpts)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Cdkcelestium *CdkcelestiumCallerSession) TrustedAggregator() (common.Address, error) {
	return _Cdkcelestium.Contract.TrustedAggregator(&_Cdkcelestium.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) TrustedAggregatorTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "trustedAggregatorTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Cdkcelestium.Contract.TrustedAggregatorTimeout(&_Cdkcelestium.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Cdkcelestium.Contract.TrustedAggregatorTimeout(&_Cdkcelestium.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Cdkcelestium *CdkcelestiumCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Cdkcelestium *CdkcelestiumSession) TrustedSequencer() (common.Address, error) {
	return _Cdkcelestium.Contract.TrustedSequencer(&_Cdkcelestium.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Cdkcelestium *CdkcelestiumCallerSession) TrustedSequencer() (common.Address, error) {
	return _Cdkcelestium.Contract.TrustedSequencer(&_Cdkcelestium.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Cdkcelestium *CdkcelestiumCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Cdkcelestium *CdkcelestiumSession) TrustedSequencerURL() (string, error) {
	return _Cdkcelestium.Contract.TrustedSequencerURL(&_Cdkcelestium.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Cdkcelestium *CdkcelestiumCallerSession) TrustedSequencerURL() (string, error) {
	return _Cdkcelestium.Contract.TrustedSequencerURL(&_Cdkcelestium.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCaller) VerifyBatchTimeTarget(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cdkcelestium.contract.Call(opts, &out, "verifyBatchTimeTarget")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Cdkcelestium.Contract.VerifyBatchTimeTarget(&_Cdkcelestium.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Cdkcelestium *CdkcelestiumCallerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Cdkcelestium.Contract.VerifyBatchTimeTarget(&_Cdkcelestium.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Cdkcelestium *CdkcelestiumTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Cdkcelestium *CdkcelestiumSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Cdkcelestium.Contract.AcceptAdminRole(&_Cdkcelestium.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Cdkcelestium.Contract.AcceptAdminRole(&_Cdkcelestium.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) ActivateEmergencyState(opts *bind.TransactOpts, sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "activateEmergencyState", sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Cdkcelestium *CdkcelestiumSession) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.ActivateEmergencyState(&_Cdkcelestium.TransactOpts, sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.ActivateEmergencyState(&_Cdkcelestium.TransactOpts, sequencedBatchNum)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Cdkcelestium *CdkcelestiumTransactor) ActivateForceBatches(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "activateForceBatches")
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Cdkcelestium *CdkcelestiumSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Cdkcelestium.Contract.ActivateForceBatches(&_Cdkcelestium.TransactOpts)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Cdkcelestium.Contract.ActivateForceBatches(&_Cdkcelestium.TransactOpts)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) ConsolidatePendingState(opts *bind.TransactOpts, pendingStateNum uint64) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "consolidatePendingState", pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Cdkcelestium *CdkcelestiumSession) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.ConsolidatePendingState(&_Cdkcelestium.TransactOpts, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.ConsolidatePendingState(&_Cdkcelestium.TransactOpts, pendingStateNum)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Cdkcelestium *CdkcelestiumTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Cdkcelestium *CdkcelestiumSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Cdkcelestium.Contract.DeactivateEmergencyState(&_Cdkcelestium.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Cdkcelestium.Contract.DeactivateEmergencyState(&_Cdkcelestium.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "forceBatch", transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Cdkcelestium *CdkcelestiumSession) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.ForceBatch(&_Cdkcelestium.TransactOpts, transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.ForceBatch(&_Cdkcelestium.TransactOpts, transactions, maticAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) Initialize(opts *bind.TransactOpts, initializePackedParameters CDKCelestiumInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "initialize", initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Cdkcelestium *CdkcelestiumSession) Initialize(initializePackedParameters CDKCelestiumInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.Initialize(&_Cdkcelestium.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) Initialize(initializePackedParameters CDKCelestiumInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.Initialize(&_Cdkcelestium.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) OverridePendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "overridePendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Cdkcelestium *CdkcelestiumSession) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.OverridePendingState(&_Cdkcelestium.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.OverridePendingState(&_Cdkcelestium.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) ProveNonDeterministicPendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "proveNonDeterministicPendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Cdkcelestium *CdkcelestiumSession) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.ProveNonDeterministicPendingState(&_Cdkcelestium.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.ProveNonDeterministicPendingState(&_Cdkcelestium.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cdkcelestium *CdkcelestiumTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cdkcelestium *CdkcelestiumSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cdkcelestium.Contract.RenounceOwnership(&_Cdkcelestium.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cdkcelestium.Contract.RenounceOwnership(&_Cdkcelestium.TransactOpts)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x438a5399.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes _message) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) SequenceBatches(opts *bind.TransactOpts, batches []CDKCelestiumBatchData, l2Coinbase common.Address, _message []byte) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "sequenceBatches", batches, l2Coinbase, _message)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x438a5399.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes _message) returns()
func (_Cdkcelestium *CdkcelestiumSession) SequenceBatches(batches []CDKCelestiumBatchData, l2Coinbase common.Address, _message []byte) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SequenceBatches(&_Cdkcelestium.TransactOpts, batches, l2Coinbase, _message)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x438a5399.
//
// Solidity: function sequenceBatches((bytes32,bytes32,uint64,uint64)[] batches, address l2Coinbase, bytes _message) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) SequenceBatches(batches []CDKCelestiumBatchData, l2Coinbase common.Address, _message []byte) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SequenceBatches(&_Cdkcelestium.TransactOpts, batches, l2Coinbase, _message)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []CDKCelestiumForcedBatchData) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Cdkcelestium *CdkcelestiumSession) SequenceForceBatches(batches []CDKCelestiumForcedBatchData) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SequenceForceBatches(&_Cdkcelestium.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) SequenceForceBatches(batches []CDKCelestiumForcedBatchData) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SequenceForceBatches(&_Cdkcelestium.TransactOpts, batches)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Cdkcelestium *CdkcelestiumSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetForceBatchTimeout(&_Cdkcelestium.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetForceBatchTimeout(&_Cdkcelestium.TransactOpts, newforceBatchTimeout)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) SetMultiplierBatchFee(opts *bind.TransactOpts, newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "setMultiplierBatchFee", newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Cdkcelestium *CdkcelestiumSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetMultiplierBatchFee(&_Cdkcelestium.TransactOpts, newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetMultiplierBatchFee(&_Cdkcelestium.TransactOpts, newMultiplierBatchFee)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) SetPendingStateTimeout(opts *bind.TransactOpts, newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "setPendingStateTimeout", newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Cdkcelestium *CdkcelestiumSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetPendingStateTimeout(&_Cdkcelestium.TransactOpts, newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetPendingStateTimeout(&_Cdkcelestium.TransactOpts, newPendingStateTimeout)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) SetTrustedAggregator(opts *bind.TransactOpts, newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "setTrustedAggregator", newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Cdkcelestium *CdkcelestiumSession) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetTrustedAggregator(&_Cdkcelestium.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetTrustedAggregator(&_Cdkcelestium.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) SetTrustedAggregatorTimeout(opts *bind.TransactOpts, newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "setTrustedAggregatorTimeout", newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Cdkcelestium *CdkcelestiumSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetTrustedAggregatorTimeout(&_Cdkcelestium.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetTrustedAggregatorTimeout(&_Cdkcelestium.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Cdkcelestium *CdkcelestiumSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetTrustedSequencer(&_Cdkcelestium.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetTrustedSequencer(&_Cdkcelestium.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Cdkcelestium *CdkcelestiumSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetTrustedSequencerURL(&_Cdkcelestium.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetTrustedSequencerURL(&_Cdkcelestium.TransactOpts, newTrustedSequencerURL)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) SetVerifyBatchTimeTarget(opts *bind.TransactOpts, newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "setVerifyBatchTimeTarget", newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Cdkcelestium *CdkcelestiumSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetVerifyBatchTimeTarget(&_Cdkcelestium.TransactOpts, newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.SetVerifyBatchTimeTarget(&_Cdkcelestium.TransactOpts, newVerifyBatchTimeTarget)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Cdkcelestium *CdkcelestiumSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.TransferAdminRole(&_Cdkcelestium.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.TransferAdminRole(&_Cdkcelestium.TransactOpts, newPendingAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cdkcelestium *CdkcelestiumSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.TransferOwnership(&_Cdkcelestium.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.TransferOwnership(&_Cdkcelestium.TransactOpts, newOwner)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) VerifyBatches(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "verifyBatches", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Cdkcelestium *CdkcelestiumSession) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.VerifyBatches(&_Cdkcelestium.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.VerifyBatches(&_Cdkcelestium.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Cdkcelestium *CdkcelestiumTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Cdkcelestium.contract.Transact(opts, "verifyBatchesTrustedAggregator", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Cdkcelestium *CdkcelestiumSession) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.VerifyBatchesTrustedAggregator(&_Cdkcelestium.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Cdkcelestium *CdkcelestiumTransactorSession) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Cdkcelestium.Contract.VerifyBatchesTrustedAggregator(&_Cdkcelestium.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// CdkcelestiumAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Cdkcelestium contract.
type CdkcelestiumAcceptAdminRoleIterator struct {
	Event *CdkcelestiumAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumAcceptAdminRole)
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
		it.Event = new(CdkcelestiumAcceptAdminRole)
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
func (it *CdkcelestiumAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumAcceptAdminRole represents a AcceptAdminRole event raised by the Cdkcelestium contract.
type CdkcelestiumAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*CdkcelestiumAcceptAdminRoleIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumAcceptAdminRoleIterator{contract: _Cdkcelestium.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *CdkcelestiumAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumAcceptAdminRole)
				if err := _Cdkcelestium.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseAcceptAdminRole(log types.Log) (*CdkcelestiumAcceptAdminRole, error) {
	event := new(CdkcelestiumAcceptAdminRole)
	if err := _Cdkcelestium.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumActivateForceBatchesIterator is returned from FilterActivateForceBatches and is used to iterate over the raw logs and unpacked data for ActivateForceBatches events raised by the Cdkcelestium contract.
type CdkcelestiumActivateForceBatchesIterator struct {
	Event *CdkcelestiumActivateForceBatches // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumActivateForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumActivateForceBatches)
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
		it.Event = new(CdkcelestiumActivateForceBatches)
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
func (it *CdkcelestiumActivateForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumActivateForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumActivateForceBatches represents a ActivateForceBatches event raised by the Cdkcelestium contract.
type CdkcelestiumActivateForceBatches struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActivateForceBatches is a free log retrieval operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Cdkcelestium *CdkcelestiumFilterer) FilterActivateForceBatches(opts *bind.FilterOpts) (*CdkcelestiumActivateForceBatchesIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumActivateForceBatchesIterator{contract: _Cdkcelestium.contract, event: "ActivateForceBatches", logs: logs, sub: sub}, nil
}

// WatchActivateForceBatches is a free log subscription operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Cdkcelestium *CdkcelestiumFilterer) WatchActivateForceBatches(opts *bind.WatchOpts, sink chan<- *CdkcelestiumActivateForceBatches) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumActivateForceBatches)
				if err := _Cdkcelestium.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseActivateForceBatches(log types.Log) (*CdkcelestiumActivateForceBatches, error) {
	event := new(CdkcelestiumActivateForceBatches)
	if err := _Cdkcelestium.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumConsolidatePendingStateIterator is returned from FilterConsolidatePendingState and is used to iterate over the raw logs and unpacked data for ConsolidatePendingState events raised by the Cdkcelestium contract.
type CdkcelestiumConsolidatePendingStateIterator struct {
	Event *CdkcelestiumConsolidatePendingState // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumConsolidatePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumConsolidatePendingState)
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
		it.Event = new(CdkcelestiumConsolidatePendingState)
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
func (it *CdkcelestiumConsolidatePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumConsolidatePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumConsolidatePendingState represents a ConsolidatePendingState event raised by the Cdkcelestium contract.
type CdkcelestiumConsolidatePendingState struct {
	NumBatch        uint64
	StateRoot       [32]byte
	PendingStateNum uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterConsolidatePendingState is a free log retrieval operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterConsolidatePendingState(opts *bind.FilterOpts, numBatch []uint64, pendingStateNum []uint64) (*CdkcelestiumConsolidatePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumConsolidatePendingStateIterator{contract: _Cdkcelestium.contract, event: "ConsolidatePendingState", logs: logs, sub: sub}, nil
}

// WatchConsolidatePendingState is a free log subscription operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchConsolidatePendingState(opts *bind.WatchOpts, sink chan<- *CdkcelestiumConsolidatePendingState, numBatch []uint64, pendingStateNum []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumConsolidatePendingState)
				if err := _Cdkcelestium.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseConsolidatePendingState(log types.Log) (*CdkcelestiumConsolidatePendingState, error) {
	event := new(CdkcelestiumConsolidatePendingState)
	if err := _Cdkcelestium.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Cdkcelestium contract.
type CdkcelestiumEmergencyStateActivatedIterator struct {
	Event *CdkcelestiumEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumEmergencyStateActivated)
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
		it.Event = new(CdkcelestiumEmergencyStateActivated)
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
func (it *CdkcelestiumEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumEmergencyStateActivated represents a EmergencyStateActivated event raised by the Cdkcelestium contract.
type CdkcelestiumEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Cdkcelestium *CdkcelestiumFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*CdkcelestiumEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumEmergencyStateActivatedIterator{contract: _Cdkcelestium.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Cdkcelestium *CdkcelestiumFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *CdkcelestiumEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumEmergencyStateActivated)
				if err := _Cdkcelestium.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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

// ParseEmergencyStateActivated is a log parse operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Cdkcelestium *CdkcelestiumFilterer) ParseEmergencyStateActivated(log types.Log) (*CdkcelestiumEmergencyStateActivated, error) {
	event := new(CdkcelestiumEmergencyStateActivated)
	if err := _Cdkcelestium.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Cdkcelestium contract.
type CdkcelestiumEmergencyStateDeactivatedIterator struct {
	Event *CdkcelestiumEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumEmergencyStateDeactivated)
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
		it.Event = new(CdkcelestiumEmergencyStateDeactivated)
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
func (it *CdkcelestiumEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Cdkcelestium contract.
type CdkcelestiumEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Cdkcelestium *CdkcelestiumFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*CdkcelestiumEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumEmergencyStateDeactivatedIterator{contract: _Cdkcelestium.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Cdkcelestium *CdkcelestiumFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *CdkcelestiumEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumEmergencyStateDeactivated)
				if err := _Cdkcelestium.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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

// ParseEmergencyStateDeactivated is a log parse operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Cdkcelestium *CdkcelestiumFilterer) ParseEmergencyStateDeactivated(log types.Log) (*CdkcelestiumEmergencyStateDeactivated, error) {
	event := new(CdkcelestiumEmergencyStateDeactivated)
	if err := _Cdkcelestium.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Cdkcelestium contract.
type CdkcelestiumForceBatchIterator struct {
	Event *CdkcelestiumForceBatch // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumForceBatch)
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
		it.Event = new(CdkcelestiumForceBatch)
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
func (it *CdkcelestiumForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumForceBatch represents a ForceBatch event raised by the Cdkcelestium contract.
type CdkcelestiumForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*CdkcelestiumForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumForceBatchIterator{contract: _Cdkcelestium.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *CdkcelestiumForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumForceBatch)
				if err := _Cdkcelestium.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseForceBatch(log types.Log) (*CdkcelestiumForceBatch, error) {
	event := new(CdkcelestiumForceBatch)
	if err := _Cdkcelestium.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Cdkcelestium contract.
type CdkcelestiumInitializedIterator struct {
	Event *CdkcelestiumInitialized // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumInitialized)
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
		it.Event = new(CdkcelestiumInitialized)
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
func (it *CdkcelestiumInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumInitialized represents a Initialized event raised by the Cdkcelestium contract.
type CdkcelestiumInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterInitialized(opts *bind.FilterOpts) (*CdkcelestiumInitializedIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumInitializedIterator{contract: _Cdkcelestium.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CdkcelestiumInitialized) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumInitialized)
				if err := _Cdkcelestium.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cdkcelestium *CdkcelestiumFilterer) ParseInitialized(log types.Log) (*CdkcelestiumInitialized, error) {
	event := new(CdkcelestiumInitialized)
	if err := _Cdkcelestium.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumOverridePendingStateIterator is returned from FilterOverridePendingState and is used to iterate over the raw logs and unpacked data for OverridePendingState events raised by the Cdkcelestium contract.
type CdkcelestiumOverridePendingStateIterator struct {
	Event *CdkcelestiumOverridePendingState // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumOverridePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumOverridePendingState)
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
		it.Event = new(CdkcelestiumOverridePendingState)
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
func (it *CdkcelestiumOverridePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumOverridePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumOverridePendingState represents a OverridePendingState event raised by the Cdkcelestium contract.
type CdkcelestiumOverridePendingState struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOverridePendingState is a free log retrieval operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterOverridePendingState(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*CdkcelestiumOverridePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumOverridePendingStateIterator{contract: _Cdkcelestium.contract, event: "OverridePendingState", logs: logs, sub: sub}, nil
}

// WatchOverridePendingState is a free log subscription operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchOverridePendingState(opts *bind.WatchOpts, sink chan<- *CdkcelestiumOverridePendingState, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumOverridePendingState)
				if err := _Cdkcelestium.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseOverridePendingState(log types.Log) (*CdkcelestiumOverridePendingState, error) {
	event := new(CdkcelestiumOverridePendingState)
	if err := _Cdkcelestium.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cdkcelestium contract.
type CdkcelestiumOwnershipTransferredIterator struct {
	Event *CdkcelestiumOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumOwnershipTransferred)
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
		it.Event = new(CdkcelestiumOwnershipTransferred)
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
func (it *CdkcelestiumOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumOwnershipTransferred represents a OwnershipTransferred event raised by the Cdkcelestium contract.
type CdkcelestiumOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CdkcelestiumOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumOwnershipTransferredIterator{contract: _Cdkcelestium.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CdkcelestiumOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumOwnershipTransferred)
				if err := _Cdkcelestium.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseOwnershipTransferred(log types.Log) (*CdkcelestiumOwnershipTransferred, error) {
	event := new(CdkcelestiumOwnershipTransferred)
	if err := _Cdkcelestium.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumProveNonDeterministicPendingStateIterator is returned from FilterProveNonDeterministicPendingState and is used to iterate over the raw logs and unpacked data for ProveNonDeterministicPendingState events raised by the Cdkcelestium contract.
type CdkcelestiumProveNonDeterministicPendingStateIterator struct {
	Event *CdkcelestiumProveNonDeterministicPendingState // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumProveNonDeterministicPendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumProveNonDeterministicPendingState)
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
		it.Event = new(CdkcelestiumProveNonDeterministicPendingState)
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
func (it *CdkcelestiumProveNonDeterministicPendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumProveNonDeterministicPendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumProveNonDeterministicPendingState represents a ProveNonDeterministicPendingState event raised by the Cdkcelestium contract.
type CdkcelestiumProveNonDeterministicPendingState struct {
	StoredStateRoot [32]byte
	ProvedStateRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProveNonDeterministicPendingState is a free log retrieval operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterProveNonDeterministicPendingState(opts *bind.FilterOpts) (*CdkcelestiumProveNonDeterministicPendingStateIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumProveNonDeterministicPendingStateIterator{contract: _Cdkcelestium.contract, event: "ProveNonDeterministicPendingState", logs: logs, sub: sub}, nil
}

// WatchProveNonDeterministicPendingState is a free log subscription operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchProveNonDeterministicPendingState(opts *bind.WatchOpts, sink chan<- *CdkcelestiumProveNonDeterministicPendingState) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumProveNonDeterministicPendingState)
				if err := _Cdkcelestium.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseProveNonDeterministicPendingState(log types.Log) (*CdkcelestiumProveNonDeterministicPendingState, error) {
	event := new(CdkcelestiumProveNonDeterministicPendingState)
	if err := _Cdkcelestium.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Cdkcelestium contract.
type CdkcelestiumSequenceBatchesIterator struct {
	Event *CdkcelestiumSequenceBatches // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumSequenceBatches)
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
		it.Event = new(CdkcelestiumSequenceBatches)
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
func (it *CdkcelestiumSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumSequenceBatches represents a SequenceBatches event raised by the Cdkcelestium contract.
type CdkcelestiumSequenceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*CdkcelestiumSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumSequenceBatchesIterator{contract: _Cdkcelestium.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *CdkcelestiumSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumSequenceBatches)
				if err := _Cdkcelestium.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseSequenceBatches(log types.Log) (*CdkcelestiumSequenceBatches, error) {
	event := new(CdkcelestiumSequenceBatches)
	if err := _Cdkcelestium.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Cdkcelestium contract.
type CdkcelestiumSequenceForceBatchesIterator struct {
	Event *CdkcelestiumSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumSequenceForceBatches)
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
		it.Event = new(CdkcelestiumSequenceForceBatches)
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
func (it *CdkcelestiumSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumSequenceForceBatches represents a SequenceForceBatches event raised by the Cdkcelestium contract.
type CdkcelestiumSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*CdkcelestiumSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumSequenceForceBatchesIterator{contract: _Cdkcelestium.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *CdkcelestiumSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumSequenceForceBatches)
				if err := _Cdkcelestium.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseSequenceForceBatches(log types.Log) (*CdkcelestiumSequenceForceBatches, error) {
	event := new(CdkcelestiumSequenceForceBatches)
	if err := _Cdkcelestium.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Cdkcelestium contract.
type CdkcelestiumSetForceBatchTimeoutIterator struct {
	Event *CdkcelestiumSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumSetForceBatchTimeout)
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
		it.Event = new(CdkcelestiumSetForceBatchTimeout)
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
func (it *CdkcelestiumSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Cdkcelestium contract.
type CdkcelestiumSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*CdkcelestiumSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumSetForceBatchTimeoutIterator{contract: _Cdkcelestium.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *CdkcelestiumSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumSetForceBatchTimeout)
				if err := _Cdkcelestium.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseSetForceBatchTimeout(log types.Log) (*CdkcelestiumSetForceBatchTimeout, error) {
	event := new(CdkcelestiumSetForceBatchTimeout)
	if err := _Cdkcelestium.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumSetMultiplierBatchFeeIterator is returned from FilterSetMultiplierBatchFee and is used to iterate over the raw logs and unpacked data for SetMultiplierBatchFee events raised by the Cdkcelestium contract.
type CdkcelestiumSetMultiplierBatchFeeIterator struct {
	Event *CdkcelestiumSetMultiplierBatchFee // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumSetMultiplierBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumSetMultiplierBatchFee)
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
		it.Event = new(CdkcelestiumSetMultiplierBatchFee)
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
func (it *CdkcelestiumSetMultiplierBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumSetMultiplierBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumSetMultiplierBatchFee represents a SetMultiplierBatchFee event raised by the Cdkcelestium contract.
type CdkcelestiumSetMultiplierBatchFee struct {
	NewMultiplierBatchFee uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMultiplierBatchFee is a free log retrieval operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterSetMultiplierBatchFee(opts *bind.FilterOpts) (*CdkcelestiumSetMultiplierBatchFeeIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumSetMultiplierBatchFeeIterator{contract: _Cdkcelestium.contract, event: "SetMultiplierBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetMultiplierBatchFee is a free log subscription operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchSetMultiplierBatchFee(opts *bind.WatchOpts, sink chan<- *CdkcelestiumSetMultiplierBatchFee) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumSetMultiplierBatchFee)
				if err := _Cdkcelestium.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseSetMultiplierBatchFee(log types.Log) (*CdkcelestiumSetMultiplierBatchFee, error) {
	event := new(CdkcelestiumSetMultiplierBatchFee)
	if err := _Cdkcelestium.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumSetPendingStateTimeoutIterator is returned from FilterSetPendingStateTimeout and is used to iterate over the raw logs and unpacked data for SetPendingStateTimeout events raised by the Cdkcelestium contract.
type CdkcelestiumSetPendingStateTimeoutIterator struct {
	Event *CdkcelestiumSetPendingStateTimeout // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumSetPendingStateTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumSetPendingStateTimeout)
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
		it.Event = new(CdkcelestiumSetPendingStateTimeout)
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
func (it *CdkcelestiumSetPendingStateTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumSetPendingStateTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumSetPendingStateTimeout represents a SetPendingStateTimeout event raised by the Cdkcelestium contract.
type CdkcelestiumSetPendingStateTimeout struct {
	NewPendingStateTimeout uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetPendingStateTimeout is a free log retrieval operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterSetPendingStateTimeout(opts *bind.FilterOpts) (*CdkcelestiumSetPendingStateTimeoutIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumSetPendingStateTimeoutIterator{contract: _Cdkcelestium.contract, event: "SetPendingStateTimeout", logs: logs, sub: sub}, nil
}

// WatchSetPendingStateTimeout is a free log subscription operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchSetPendingStateTimeout(opts *bind.WatchOpts, sink chan<- *CdkcelestiumSetPendingStateTimeout) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumSetPendingStateTimeout)
				if err := _Cdkcelestium.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseSetPendingStateTimeout(log types.Log) (*CdkcelestiumSetPendingStateTimeout, error) {
	event := new(CdkcelestiumSetPendingStateTimeout)
	if err := _Cdkcelestium.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Cdkcelestium contract.
type CdkcelestiumSetTrustedAggregatorIterator struct {
	Event *CdkcelestiumSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumSetTrustedAggregator)
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
		it.Event = new(CdkcelestiumSetTrustedAggregator)
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
func (it *CdkcelestiumSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumSetTrustedAggregator represents a SetTrustedAggregator event raised by the Cdkcelestium contract.
type CdkcelestiumSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*CdkcelestiumSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumSetTrustedAggregatorIterator{contract: _Cdkcelestium.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *CdkcelestiumSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumSetTrustedAggregator)
				if err := _Cdkcelestium.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseSetTrustedAggregator(log types.Log) (*CdkcelestiumSetTrustedAggregator, error) {
	event := new(CdkcelestiumSetTrustedAggregator)
	if err := _Cdkcelestium.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumSetTrustedAggregatorTimeoutIterator is returned from FilterSetTrustedAggregatorTimeout and is used to iterate over the raw logs and unpacked data for SetTrustedAggregatorTimeout events raised by the Cdkcelestium contract.
type CdkcelestiumSetTrustedAggregatorTimeoutIterator struct {
	Event *CdkcelestiumSetTrustedAggregatorTimeout // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumSetTrustedAggregatorTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumSetTrustedAggregatorTimeout)
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
		it.Event = new(CdkcelestiumSetTrustedAggregatorTimeout)
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
func (it *CdkcelestiumSetTrustedAggregatorTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumSetTrustedAggregatorTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumSetTrustedAggregatorTimeout represents a SetTrustedAggregatorTimeout event raised by the Cdkcelestium contract.
type CdkcelestiumSetTrustedAggregatorTimeout struct {
	NewTrustedAggregatorTimeout uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregatorTimeout is a free log retrieval operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterSetTrustedAggregatorTimeout(opts *bind.FilterOpts) (*CdkcelestiumSetTrustedAggregatorTimeoutIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumSetTrustedAggregatorTimeoutIterator{contract: _Cdkcelestium.contract, event: "SetTrustedAggregatorTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregatorTimeout is a free log subscription operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchSetTrustedAggregatorTimeout(opts *bind.WatchOpts, sink chan<- *CdkcelestiumSetTrustedAggregatorTimeout) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumSetTrustedAggregatorTimeout)
				if err := _Cdkcelestium.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseSetTrustedAggregatorTimeout(log types.Log) (*CdkcelestiumSetTrustedAggregatorTimeout, error) {
	event := new(CdkcelestiumSetTrustedAggregatorTimeout)
	if err := _Cdkcelestium.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Cdkcelestium contract.
type CdkcelestiumSetTrustedSequencerIterator struct {
	Event *CdkcelestiumSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumSetTrustedSequencer)
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
		it.Event = new(CdkcelestiumSetTrustedSequencer)
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
func (it *CdkcelestiumSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumSetTrustedSequencer represents a SetTrustedSequencer event raised by the Cdkcelestium contract.
type CdkcelestiumSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*CdkcelestiumSetTrustedSequencerIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumSetTrustedSequencerIterator{contract: _Cdkcelestium.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *CdkcelestiumSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumSetTrustedSequencer)
				if err := _Cdkcelestium.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseSetTrustedSequencer(log types.Log) (*CdkcelestiumSetTrustedSequencer, error) {
	event := new(CdkcelestiumSetTrustedSequencer)
	if err := _Cdkcelestium.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Cdkcelestium contract.
type CdkcelestiumSetTrustedSequencerURLIterator struct {
	Event *CdkcelestiumSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumSetTrustedSequencerURL)
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
		it.Event = new(CdkcelestiumSetTrustedSequencerURL)
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
func (it *CdkcelestiumSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Cdkcelestium contract.
type CdkcelestiumSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*CdkcelestiumSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumSetTrustedSequencerURLIterator{contract: _Cdkcelestium.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *CdkcelestiumSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumSetTrustedSequencerURL)
				if err := _Cdkcelestium.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseSetTrustedSequencerURL(log types.Log) (*CdkcelestiumSetTrustedSequencerURL, error) {
	event := new(CdkcelestiumSetTrustedSequencerURL)
	if err := _Cdkcelestium.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumSetVerifyBatchTimeTargetIterator is returned from FilterSetVerifyBatchTimeTarget and is used to iterate over the raw logs and unpacked data for SetVerifyBatchTimeTarget events raised by the Cdkcelestium contract.
type CdkcelestiumSetVerifyBatchTimeTargetIterator struct {
	Event *CdkcelestiumSetVerifyBatchTimeTarget // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumSetVerifyBatchTimeTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumSetVerifyBatchTimeTarget)
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
		it.Event = new(CdkcelestiumSetVerifyBatchTimeTarget)
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
func (it *CdkcelestiumSetVerifyBatchTimeTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumSetVerifyBatchTimeTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumSetVerifyBatchTimeTarget represents a SetVerifyBatchTimeTarget event raised by the Cdkcelestium contract.
type CdkcelestiumSetVerifyBatchTimeTarget struct {
	NewVerifyBatchTimeTarget uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetVerifyBatchTimeTarget is a free log retrieval operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterSetVerifyBatchTimeTarget(opts *bind.FilterOpts) (*CdkcelestiumSetVerifyBatchTimeTargetIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumSetVerifyBatchTimeTargetIterator{contract: _Cdkcelestium.contract, event: "SetVerifyBatchTimeTarget", logs: logs, sub: sub}, nil
}

// WatchSetVerifyBatchTimeTarget is a free log subscription operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchSetVerifyBatchTimeTarget(opts *bind.WatchOpts, sink chan<- *CdkcelestiumSetVerifyBatchTimeTarget) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumSetVerifyBatchTimeTarget)
				if err := _Cdkcelestium.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseSetVerifyBatchTimeTarget(log types.Log) (*CdkcelestiumSetVerifyBatchTimeTarget, error) {
	event := new(CdkcelestiumSetVerifyBatchTimeTarget)
	if err := _Cdkcelestium.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Cdkcelestium contract.
type CdkcelestiumTransferAdminRoleIterator struct {
	Event *CdkcelestiumTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumTransferAdminRole)
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
		it.Event = new(CdkcelestiumTransferAdminRole)
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
func (it *CdkcelestiumTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumTransferAdminRole represents a TransferAdminRole event raised by the Cdkcelestium contract.
type CdkcelestiumTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*CdkcelestiumTransferAdminRoleIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumTransferAdminRoleIterator{contract: _Cdkcelestium.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *CdkcelestiumTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumTransferAdminRole)
				if err := _Cdkcelestium.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseTransferAdminRole(log types.Log) (*CdkcelestiumTransferAdminRole, error) {
	event := new(CdkcelestiumTransferAdminRole)
	if err := _Cdkcelestium.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumUpdateZkEVMVersionIterator is returned from FilterUpdateZkEVMVersion and is used to iterate over the raw logs and unpacked data for UpdateZkEVMVersion events raised by the Cdkcelestium contract.
type CdkcelestiumUpdateZkEVMVersionIterator struct {
	Event *CdkcelestiumUpdateZkEVMVersion // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumUpdateZkEVMVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumUpdateZkEVMVersion)
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
		it.Event = new(CdkcelestiumUpdateZkEVMVersion)
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
func (it *CdkcelestiumUpdateZkEVMVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumUpdateZkEVMVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumUpdateZkEVMVersion represents a UpdateZkEVMVersion event raised by the Cdkcelestium contract.
type CdkcelestiumUpdateZkEVMVersion struct {
	NumBatch uint64
	ForkID   uint64
	Version  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateZkEVMVersion is a free log retrieval operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterUpdateZkEVMVersion(opts *bind.FilterOpts) (*CdkcelestiumUpdateZkEVMVersionIterator, error) {

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "UpdateZkEVMVersion")
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumUpdateZkEVMVersionIterator{contract: _Cdkcelestium.contract, event: "UpdateZkEVMVersion", logs: logs, sub: sub}, nil
}

// WatchUpdateZkEVMVersion is a free log subscription operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchUpdateZkEVMVersion(opts *bind.WatchOpts, sink chan<- *CdkcelestiumUpdateZkEVMVersion) (event.Subscription, error) {

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "UpdateZkEVMVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumUpdateZkEVMVersion)
				if err := _Cdkcelestium.contract.UnpackLog(event, "UpdateZkEVMVersion", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseUpdateZkEVMVersion(log types.Log) (*CdkcelestiumUpdateZkEVMVersion, error) {
	event := new(CdkcelestiumUpdateZkEVMVersion)
	if err := _Cdkcelestium.contract.UnpackLog(event, "UpdateZkEVMVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Cdkcelestium contract.
type CdkcelestiumVerifyBatchesIterator struct {
	Event *CdkcelestiumVerifyBatches // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumVerifyBatches)
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
		it.Event = new(CdkcelestiumVerifyBatches)
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
func (it *CdkcelestiumVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumVerifyBatches represents a VerifyBatches event raised by the Cdkcelestium contract.
type CdkcelestiumVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*CdkcelestiumVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumVerifyBatchesIterator{contract: _Cdkcelestium.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *CdkcelestiumVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumVerifyBatches)
				if err := _Cdkcelestium.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseVerifyBatches(log types.Log) (*CdkcelestiumVerifyBatches, error) {
	event := new(CdkcelestiumVerifyBatches)
	if err := _Cdkcelestium.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkcelestiumVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Cdkcelestium contract.
type CdkcelestiumVerifyBatchesTrustedAggregatorIterator struct {
	Event *CdkcelestiumVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *CdkcelestiumVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkcelestiumVerifyBatchesTrustedAggregator)
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
		it.Event = new(CdkcelestiumVerifyBatchesTrustedAggregator)
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
func (it *CdkcelestiumVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkcelestiumVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkcelestiumVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Cdkcelestium contract.
type CdkcelestiumVerifyBatchesTrustedAggregator struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatchesTrustedAggregator is a free log retrieval operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Cdkcelestium *CdkcelestiumFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*CdkcelestiumVerifyBatchesTrustedAggregatorIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Cdkcelestium.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &CdkcelestiumVerifyBatchesTrustedAggregatorIterator{contract: _Cdkcelestium.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Cdkcelestium *CdkcelestiumFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *CdkcelestiumVerifyBatchesTrustedAggregator, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Cdkcelestium.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkcelestiumVerifyBatchesTrustedAggregator)
				if err := _Cdkcelestium.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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
func (_Cdkcelestium *CdkcelestiumFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*CdkcelestiumVerifyBatchesTrustedAggregator, error) {
	event := new(CdkcelestiumVerifyBatchesTrustedAggregator)
	if err := _Cdkcelestium.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
