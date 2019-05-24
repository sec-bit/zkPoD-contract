// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PODEXABI is the input ABI used to generate the binding from.
const PODEXABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"publicVar_\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyerDeposits_\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"unDepositAt\",\"type\":\"uint256\"},{\"name\":\"pendingCnt\",\"type\":\"uint256\"},{\"name\":\"stat\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"s_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"t3\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bulletins_\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"size\",\"type\":\"uint64\"},{\"name\":\"s\",\"type\":\"uint64\"},{\"name\":\"n\",\"type\":\"uint64\"},{\"name\":\"sigma_mkl_root\",\"type\":\"uint256\"},{\"name\":\"vrf_meta_digest\",\"type\":\"uint256\"},{\"name\":\"pledge_value\",\"type\":\"uint256\"},{\"name\":\"unDepositAt\",\"type\":\"uint256\"},{\"name\":\"blt_type\",\"type\":\"uint8\"},{\"name\":\"stat\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"t1\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"t4\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_publicVar\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_mode\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"OnDeal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_seed0\",\"type\":\"bytes32\"}],\"name\":\"OnBatch1Key\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"OnBatch1Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"OnBatch1Deal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_seed0\",\"type\":\"bytes32\"}],\"name\":\"OnBatch2Deal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_d\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_x0_lgs\",\"type\":\"uint256\"}],\"name\":\"OnBatch3Deal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_a\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_b\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_sid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_r\",\"type\":\"uint256\"}],\"name\":\"OnVRFDeal\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_size\",\"type\":\"uint64\"},{\"name\":\"_s\",\"type\":\"uint64\"},{\"name\":\"_n\",\"type\":\"uint64\"},{\"name\":\"_sigma_mkl_root\",\"type\":\"uint256\"},{\"name\":\"_vrf_meta_digest\",\"type\":\"uint256\"},{\"name\":\"_blt_type\",\"type\":\"uint256\"}],\"name\":\"publish\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bltKey\",\"type\":\"bytes32\"}],\"name\":\"unPublish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"buyerDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"buyerUnDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bltKey\",\"type\":\"bytes32\"}],\"name\":\"withdrawA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawB\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_seed0\",\"type\":\"bytes32\"},{\"name\":\"_sid\",\"type\":\"uint256\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_seed2\",\"type\":\"bytes32\"},{\"name\":\"_k_mkl_root\",\"type\":\"bytes32\"},{\"name\":\"_count\",\"type\":\"uint64\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"submitProofBatch1\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"},{\"name\":\"_i\",\"type\":\"uint64\"},{\"name\":\"_j\",\"type\":\"uint64\"},{\"name\":\"_tx\",\"type\":\"uint256\"},{\"name\":\"_ty\",\"type\":\"uint256\"},{\"name\":\"_mkl_path\",\"type\":\"bytes32[]\"},{\"name\":\"_sCnt\",\"type\":\"uint64\"}],\"name\":\"claimBatch1\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"settleBatch1Deal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_seed0\",\"type\":\"bytes32\"},{\"name\":\"_sCnt\",\"type\":\"uint64\"},{\"name\":\"_sid\",\"type\":\"uint256\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_seed2\",\"type\":\"bytes32\"},{\"name\":\"_sigma_vw\",\"type\":\"uint256\"},{\"name\":\"_count\",\"type\":\"uint64\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"submitProofBatch2\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_s_d\",\"type\":\"uint256\"},{\"name\":\"_s_x0_lgs\",\"type\":\"uint256\"},{\"name\":\"_sid\",\"type\":\"uint256\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_r_u0_x0_lgs\",\"type\":\"uint256[2]\"},{\"name\":\"_r_u0d\",\"type\":\"uint256[2]\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"submitProofBatch3\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_s_r\",\"type\":\"uint256\"},{\"name\":\"_sid\",\"type\":\"uint256\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_g_exp_r\",\"type\":\"uint256[2]\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"submitProofVRF\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"count\",\"type\":\"uint64\"},{\"name\":\"s\",\"type\":\"uint64\"},{\"name\":\"seed0\",\"type\":\"bytes32\"},{\"name\":\"seed2\",\"type\":\"bytes32\"},{\"name\":\"sigma_vw\",\"type\":\"uint256\"}],\"name\":\"vrfyProofBatch2\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_r_u0d\",\"type\":\"uint256[2]\"},{\"name\":\"_r_u0_x0_lgs\",\"type\":\"uint256[2]\"},{\"name\":\"_s_d\",\"type\":\"uint256\"},{\"name\":\"_s_x0_lgs\",\"type\":\"uint256\"}],\"name\":\"vrfyProofBatch3\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_g_exp_r\",\"type\":\"uint256[2]\"},{\"name\":\"_s_r\",\"type\":\"uint256\"}],\"name\":\"vrfyProofVRF\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"bytes32\"},{\"name\":\"_y\",\"type\":\"bytes32\"},{\"name\":\"_ij\",\"type\":\"uint64\"},{\"name\":\"_ns\",\"type\":\"uint64\"},{\"name\":\"_root\",\"type\":\"bytes32\"},{\"name\":\"_mkl_path\",\"type\":\"bytes32[]\"}],\"name\":\"vrfyPath\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"bytes32\"},{\"name\":\"_y\",\"type\":\"bytes32\"}],\"name\":\"hashInSha3\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"bytes32\"},{\"name\":\"_y\",\"type\":\"uint64\"}],\"name\":\"hashInSha3\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"seed\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"chain\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_n\",\"type\":\"uint256\"}],\"name\":\"log2ub\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"getSessionRecord\",\"outputs\":[{\"name\":\"submitAt\",\"type\":\"uint256\"},{\"name\":\"mode\",\"type\":\"uint8\"},{\"name\":\"stat\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"getRecordBatch1\",\"outputs\":[{\"name\":\"seed0\",\"type\":\"bytes32\"},{\"name\":\"seed2\",\"type\":\"bytes32\"},{\"name\":\"k_mkl_root\",\"type\":\"bytes32\"},{\"name\":\"count\",\"type\":\"uint64\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"expireAt\",\"type\":\"uint256\"},{\"name\":\"submitAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"getRecordBatch2\",\"outputs\":[{\"name\":\"seed0\",\"type\":\"bytes32\"},{\"name\":\"submitAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"getRecordBatch3\",\"outputs\":[{\"name\":\"d\",\"type\":\"uint256\"},{\"name\":\"x0_lgs\",\"type\":\"uint256\"},{\"name\":\"submitAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"getRecordVRF\",\"outputs\":[{\"name\":\"r\",\"type\":\"uint256\"},{\"name\":\"submitAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PODEX is an auto generated Go binding around an Ethereum contract.
type PODEX struct {
	PODEXCaller     // Read-only binding to the contract
	PODEXTransactor // Write-only binding to the contract
	PODEXFilterer   // Log filterer for contract events
}

// PODEXCaller is an auto generated read-only Go binding around an Ethereum contract.
type PODEXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PODEXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PODEXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PODEXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PODEXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PODEXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PODEXSession struct {
	Contract     *PODEX            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PODEXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PODEXCallerSession struct {
	Contract *PODEXCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PODEXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PODEXTransactorSession struct {
	Contract     *PODEXTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PODEXRaw is an auto generated low-level Go binding around an Ethereum contract.
type PODEXRaw struct {
	Contract *PODEX // Generic contract binding to access the raw methods on
}

// PODEXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PODEXCallerRaw struct {
	Contract *PODEXCaller // Generic read-only contract binding to access the raw methods on
}

// PODEXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PODEXTransactorRaw struct {
	Contract *PODEXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPODEX creates a new instance of PODEX, bound to a specific deployed contract.
func NewPODEX(address common.Address, backend bind.ContractBackend) (*PODEX, error) {
	contract, err := bindPODEX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PODEX{PODEXCaller: PODEXCaller{contract: contract}, PODEXTransactor: PODEXTransactor{contract: contract}, PODEXFilterer: PODEXFilterer{contract: contract}}, nil
}

// NewPODEXCaller creates a new read-only instance of PODEX, bound to a specific deployed contract.
func NewPODEXCaller(address common.Address, caller bind.ContractCaller) (*PODEXCaller, error) {
	contract, err := bindPODEX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PODEXCaller{contract: contract}, nil
}

// NewPODEXTransactor creates a new write-only instance of PODEX, bound to a specific deployed contract.
func NewPODEXTransactor(address common.Address, transactor bind.ContractTransactor) (*PODEXTransactor, error) {
	contract, err := bindPODEX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PODEXTransactor{contract: contract}, nil
}

// NewPODEXFilterer creates a new log filterer instance of PODEX, bound to a specific deployed contract.
func NewPODEXFilterer(address common.Address, filterer bind.ContractFilterer) (*PODEXFilterer, error) {
	contract, err := bindPODEX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PODEXFilterer{contract: contract}, nil
}

// bindPODEX binds a generic wrapper to an already deployed contract.
func bindPODEX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PODEXABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PODEX *PODEXRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PODEX.Contract.PODEXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PODEX *PODEXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PODEX.Contract.PODEXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PODEX *PODEXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PODEX.Contract.PODEXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PODEX *PODEXCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PODEX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PODEX *PODEXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PODEX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PODEX *PODEXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PODEX.Contract.contract.Transact(opts, method, params...)
}

// Bulletins is a free data retrieval call binding the contract method 0xd9417785.
//
// Solidity: function bulletins_(bytes32 ) constant returns(address owner, uint64 size, uint64 s, uint64 n, uint256 sigma_mkl_root, uint256 vrf_meta_digest, uint256 pledge_value, uint256 unDepositAt, uint8 blt_type, uint8 stat)
func (_PODEX *PODEXCaller) Bulletins(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Owner         common.Address
	Size          uint64
	S             uint64
	N             uint64
	SigmaMklRoot  *big.Int
	VrfMetaDigest *big.Int
	PledgeValue   *big.Int
	UnDepositAt   *big.Int
	BltType       uint8
	Stat          uint8
}, error) {
	ret := new(struct {
		Owner         common.Address
		Size          uint64
		S             uint64
		N             uint64
		SigmaMklRoot  *big.Int
		VrfMetaDigest *big.Int
		PledgeValue   *big.Int
		UnDepositAt   *big.Int
		BltType       uint8
		Stat          uint8
	})
	out := ret
	err := _PODEX.contract.Call(opts, out, "bulletins_", arg0)
	return *ret, err
}

// Bulletins is a free data retrieval call binding the contract method 0xd9417785.
//
// Solidity: function bulletins_(bytes32 ) constant returns(address owner, uint64 size, uint64 s, uint64 n, uint256 sigma_mkl_root, uint256 vrf_meta_digest, uint256 pledge_value, uint256 unDepositAt, uint8 blt_type, uint8 stat)
func (_PODEX *PODEXSession) Bulletins(arg0 [32]byte) (struct {
	Owner         common.Address
	Size          uint64
	S             uint64
	N             uint64
	SigmaMklRoot  *big.Int
	VrfMetaDigest *big.Int
	PledgeValue   *big.Int
	UnDepositAt   *big.Int
	BltType       uint8
	Stat          uint8
}, error) {
	return _PODEX.Contract.Bulletins(&_PODEX.CallOpts, arg0)
}

// Bulletins is a free data retrieval call binding the contract method 0xd9417785.
//
// Solidity: function bulletins_(bytes32 ) constant returns(address owner, uint64 size, uint64 s, uint64 n, uint256 sigma_mkl_root, uint256 vrf_meta_digest, uint256 pledge_value, uint256 unDepositAt, uint8 blt_type, uint8 stat)
func (_PODEX *PODEXCallerSession) Bulletins(arg0 [32]byte) (struct {
	Owner         common.Address
	Size          uint64
	S             uint64
	N             uint64
	SigmaMklRoot  *big.Int
	VrfMetaDigest *big.Int
	PledgeValue   *big.Int
	UnDepositAt   *big.Int
	BltType       uint8
	Stat          uint8
}, error) {
	return _PODEX.Contract.Bulletins(&_PODEX.CallOpts, arg0)
}

// BuyerDeposits is a free data retrieval call binding the contract method 0x8ca90a25.
//
// Solidity: function buyerDeposits_(address , address ) constant returns(uint256 value, uint256 unDepositAt, uint256 pendingCnt, uint8 stat)
func (_PODEX *PODEXCaller) BuyerDeposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Value       *big.Int
	UnDepositAt *big.Int
	PendingCnt  *big.Int
	Stat        uint8
}, error) {
	ret := new(struct {
		Value       *big.Int
		UnDepositAt *big.Int
		PendingCnt  *big.Int
		Stat        uint8
	})
	out := ret
	err := _PODEX.contract.Call(opts, out, "buyerDeposits_", arg0, arg1)
	return *ret, err
}

// BuyerDeposits is a free data retrieval call binding the contract method 0x8ca90a25.
//
// Solidity: function buyerDeposits_(address , address ) constant returns(uint256 value, uint256 unDepositAt, uint256 pendingCnt, uint8 stat)
func (_PODEX *PODEXSession) BuyerDeposits(arg0 common.Address, arg1 common.Address) (struct {
	Value       *big.Int
	UnDepositAt *big.Int
	PendingCnt  *big.Int
	Stat        uint8
}, error) {
	return _PODEX.Contract.BuyerDeposits(&_PODEX.CallOpts, arg0, arg1)
}

// BuyerDeposits is a free data retrieval call binding the contract method 0x8ca90a25.
//
// Solidity: function buyerDeposits_(address , address ) constant returns(uint256 value, uint256 unDepositAt, uint256 pendingCnt, uint8 stat)
func (_PODEX *PODEXCallerSession) BuyerDeposits(arg0 common.Address, arg1 common.Address) (struct {
	Value       *big.Int
	UnDepositAt *big.Int
	PendingCnt  *big.Int
	Stat        uint8
}, error) {
	return _PODEX.Contract.BuyerDeposits(&_PODEX.CallOpts, arg0, arg1)
}

// Chain is a free data retrieval call binding the contract method 0x749e0178.
//
// Solidity: function chain(bytes32 seed, uint64 index) constant returns(uint256)
func (_PODEX *PODEXCaller) Chain(opts *bind.CallOpts, seed [32]byte, index uint64) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "chain", seed, index)
	return *ret0, err
}

// Chain is a free data retrieval call binding the contract method 0x749e0178.
//
// Solidity: function chain(bytes32 seed, uint64 index) constant returns(uint256)
func (_PODEX *PODEXSession) Chain(seed [32]byte, index uint64) (*big.Int, error) {
	return _PODEX.Contract.Chain(&_PODEX.CallOpts, seed, index)
}

// Chain is a free data retrieval call binding the contract method 0x749e0178.
//
// Solidity: function chain(bytes32 seed, uint64 index) constant returns(uint256)
func (_PODEX *PODEXCallerSession) Chain(seed [32]byte, index uint64) (*big.Int, error) {
	return _PODEX.Contract.Chain(&_PODEX.CallOpts, seed, index)
}

// GetRecordBatch1 is a free data retrieval call binding the contract method 0x321d0cba.
//
// Solidity: function getRecordBatch1(address _a, address _b, uint256 _sid) constant returns(bytes32 seed0, bytes32 seed2, bytes32 k_mkl_root, uint64 count, uint256 price, uint256 expireAt, uint256 submitAt)
func (_PODEX *PODEXCaller) GetRecordBatch1(opts *bind.CallOpts, _a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0    [32]byte
	Seed2    [32]byte
	KMklRoot [32]byte
	Count    uint64
	Price    *big.Int
	ExpireAt *big.Int
	SubmitAt *big.Int
}, error) {
	ret := new(struct {
		Seed0    [32]byte
		Seed2    [32]byte
		KMklRoot [32]byte
		Count    uint64
		Price    *big.Int
		ExpireAt *big.Int
		SubmitAt *big.Int
	})
	out := ret
	err := _PODEX.contract.Call(opts, out, "getRecordBatch1", _a, _b, _sid)
	return *ret, err
}

// GetRecordBatch1 is a free data retrieval call binding the contract method 0x321d0cba.
//
// Solidity: function getRecordBatch1(address _a, address _b, uint256 _sid) constant returns(bytes32 seed0, bytes32 seed2, bytes32 k_mkl_root, uint64 count, uint256 price, uint256 expireAt, uint256 submitAt)
func (_PODEX *PODEXSession) GetRecordBatch1(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0    [32]byte
	Seed2    [32]byte
	KMklRoot [32]byte
	Count    uint64
	Price    *big.Int
	ExpireAt *big.Int
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordBatch1(&_PODEX.CallOpts, _a, _b, _sid)
}

// GetRecordBatch1 is a free data retrieval call binding the contract method 0x321d0cba.
//
// Solidity: function getRecordBatch1(address _a, address _b, uint256 _sid) constant returns(bytes32 seed0, bytes32 seed2, bytes32 k_mkl_root, uint64 count, uint256 price, uint256 expireAt, uint256 submitAt)
func (_PODEX *PODEXCallerSession) GetRecordBatch1(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0    [32]byte
	Seed2    [32]byte
	KMklRoot [32]byte
	Count    uint64
	Price    *big.Int
	ExpireAt *big.Int
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordBatch1(&_PODEX.CallOpts, _a, _b, _sid)
}

// GetRecordBatch2 is a free data retrieval call binding the contract method 0x0ae429ea.
//
// Solidity: function getRecordBatch2(address _a, address _b, uint256 _sid) constant returns(bytes32 seed0, uint256 submitAt)
func (_PODEX *PODEXCaller) GetRecordBatch2(opts *bind.CallOpts, _a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0    [32]byte
	SubmitAt *big.Int
}, error) {
	ret := new(struct {
		Seed0    [32]byte
		SubmitAt *big.Int
	})
	out := ret
	err := _PODEX.contract.Call(opts, out, "getRecordBatch2", _a, _b, _sid)
	return *ret, err
}

// GetRecordBatch2 is a free data retrieval call binding the contract method 0x0ae429ea.
//
// Solidity: function getRecordBatch2(address _a, address _b, uint256 _sid) constant returns(bytes32 seed0, uint256 submitAt)
func (_PODEX *PODEXSession) GetRecordBatch2(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0    [32]byte
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordBatch2(&_PODEX.CallOpts, _a, _b, _sid)
}

// GetRecordBatch2 is a free data retrieval call binding the contract method 0x0ae429ea.
//
// Solidity: function getRecordBatch2(address _a, address _b, uint256 _sid) constant returns(bytes32 seed0, uint256 submitAt)
func (_PODEX *PODEXCallerSession) GetRecordBatch2(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	Seed0    [32]byte
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordBatch2(&_PODEX.CallOpts, _a, _b, _sid)
}

// GetRecordBatch3 is a free data retrieval call binding the contract method 0x74ef2ae9.
//
// Solidity: function getRecordBatch3(address _a, address _b, uint256 _sid) constant returns(uint256 d, uint256 x0_lgs, uint256 submitAt)
func (_PODEX *PODEXCaller) GetRecordBatch3(opts *bind.CallOpts, _a common.Address, _b common.Address, _sid *big.Int) (struct {
	D        *big.Int
	X0Lgs    *big.Int
	SubmitAt *big.Int
}, error) {
	ret := new(struct {
		D        *big.Int
		X0Lgs    *big.Int
		SubmitAt *big.Int
	})
	out := ret
	err := _PODEX.contract.Call(opts, out, "getRecordBatch3", _a, _b, _sid)
	return *ret, err
}

// GetRecordBatch3 is a free data retrieval call binding the contract method 0x74ef2ae9.
//
// Solidity: function getRecordBatch3(address _a, address _b, uint256 _sid) constant returns(uint256 d, uint256 x0_lgs, uint256 submitAt)
func (_PODEX *PODEXSession) GetRecordBatch3(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	D        *big.Int
	X0Lgs    *big.Int
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordBatch3(&_PODEX.CallOpts, _a, _b, _sid)
}

// GetRecordBatch3 is a free data retrieval call binding the contract method 0x74ef2ae9.
//
// Solidity: function getRecordBatch3(address _a, address _b, uint256 _sid) constant returns(uint256 d, uint256 x0_lgs, uint256 submitAt)
func (_PODEX *PODEXCallerSession) GetRecordBatch3(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	D        *big.Int
	X0Lgs    *big.Int
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordBatch3(&_PODEX.CallOpts, _a, _b, _sid)
}

// GetRecordVRF is a free data retrieval call binding the contract method 0xb79881e1.
//
// Solidity: function getRecordVRF(address _a, address _b, uint256 _sid) constant returns(uint256 r, uint256 submitAt)
func (_PODEX *PODEXCaller) GetRecordVRF(opts *bind.CallOpts, _a common.Address, _b common.Address, _sid *big.Int) (struct {
	R        *big.Int
	SubmitAt *big.Int
}, error) {
	ret := new(struct {
		R        *big.Int
		SubmitAt *big.Int
	})
	out := ret
	err := _PODEX.contract.Call(opts, out, "getRecordVRF", _a, _b, _sid)
	return *ret, err
}

// GetRecordVRF is a free data retrieval call binding the contract method 0xb79881e1.
//
// Solidity: function getRecordVRF(address _a, address _b, uint256 _sid) constant returns(uint256 r, uint256 submitAt)
func (_PODEX *PODEXSession) GetRecordVRF(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	R        *big.Int
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordVRF(&_PODEX.CallOpts, _a, _b, _sid)
}

// GetRecordVRF is a free data retrieval call binding the contract method 0xb79881e1.
//
// Solidity: function getRecordVRF(address _a, address _b, uint256 _sid) constant returns(uint256 r, uint256 submitAt)
func (_PODEX *PODEXCallerSession) GetRecordVRF(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	R        *big.Int
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordVRF(&_PODEX.CallOpts, _a, _b, _sid)
}

// GetSessionRecord is a free data retrieval call binding the contract method 0xacff2f99.
//
// Solidity: function getSessionRecord(address _a, address _b, uint256 _sid) constant returns(uint256 submitAt, uint8 mode, uint8 stat)
func (_PODEX *PODEXCaller) GetSessionRecord(opts *bind.CallOpts, _a common.Address, _b common.Address, _sid *big.Int) (struct {
	SubmitAt *big.Int
	Mode     uint8
	Stat     uint8
}, error) {
	ret := new(struct {
		SubmitAt *big.Int
		Mode     uint8
		Stat     uint8
	})
	out := ret
	err := _PODEX.contract.Call(opts, out, "getSessionRecord", _a, _b, _sid)
	return *ret, err
}

// GetSessionRecord is a free data retrieval call binding the contract method 0xacff2f99.
//
// Solidity: function getSessionRecord(address _a, address _b, uint256 _sid) constant returns(uint256 submitAt, uint8 mode, uint8 stat)
func (_PODEX *PODEXSession) GetSessionRecord(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	SubmitAt *big.Int
	Mode     uint8
	Stat     uint8
}, error) {
	return _PODEX.Contract.GetSessionRecord(&_PODEX.CallOpts, _a, _b, _sid)
}

// GetSessionRecord is a free data retrieval call binding the contract method 0xacff2f99.
//
// Solidity: function getSessionRecord(address _a, address _b, uint256 _sid) constant returns(uint256 submitAt, uint8 mode, uint8 stat)
func (_PODEX *PODEXCallerSession) GetSessionRecord(_a common.Address, _b common.Address, _sid *big.Int) (struct {
	SubmitAt *big.Int
	Mode     uint8
	Stat     uint8
}, error) {
	return _PODEX.Contract.GetSessionRecord(&_PODEX.CallOpts, _a, _b, _sid)
}

// HashInSha3 is a free data retrieval call binding the contract method 0x1013f453.
//
// Solidity: function hashInSha3(bytes32 _x, uint64 _y) constant returns(bytes32)
func (_PODEX *PODEXCaller) HashInSha3(opts *bind.CallOpts, _x [32]byte, _y uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "hashInSha3", _x, _y)
	return *ret0, err
}

// HashInSha3 is a free data retrieval call binding the contract method 0x1013f453.
//
// Solidity: function hashInSha3(bytes32 _x, uint64 _y) constant returns(bytes32)
func (_PODEX *PODEXSession) HashInSha3(_x [32]byte, _y uint64) ([32]byte, error) {
	return _PODEX.Contract.HashInSha3(&_PODEX.CallOpts, _x, _y)
}

// HashInSha3 is a free data retrieval call binding the contract method 0x1013f453.
//
// Solidity: function hashInSha3(bytes32 _x, uint64 _y) constant returns(bytes32)
func (_PODEX *PODEXCallerSession) HashInSha3(_x [32]byte, _y uint64) ([32]byte, error) {
	return _PODEX.Contract.HashInSha3(&_PODEX.CallOpts, _x, _y)
}

// Log2ub is a free data retrieval call binding the contract method 0x426423c9.
//
// Solidity: function log2ub(uint256 _n) constant returns(uint256)
func (_PODEX *PODEXCaller) Log2ub(opts *bind.CallOpts, _n *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "log2ub", _n)
	return *ret0, err
}

// Log2ub is a free data retrieval call binding the contract method 0x426423c9.
//
// Solidity: function log2ub(uint256 _n) constant returns(uint256)
func (_PODEX *PODEXSession) Log2ub(_n *big.Int) (*big.Int, error) {
	return _PODEX.Contract.Log2ub(&_PODEX.CallOpts, _n)
}

// Log2ub is a free data retrieval call binding the contract method 0x426423c9.
//
// Solidity: function log2ub(uint256 _n) constant returns(uint256)
func (_PODEX *PODEXCallerSession) Log2ub(_n *big.Int) (*big.Int, error) {
	return _PODEX.Contract.Log2ub(&_PODEX.CallOpts, _n)
}

// PublicVar is a free data retrieval call binding the contract method 0x0a6c93e0.
//
// Solidity: function publicVar_() constant returns(address)
func (_PODEX *PODEXCaller) PublicVar(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "publicVar_")
	return *ret0, err
}

// PublicVar is a free data retrieval call binding the contract method 0x0a6c93e0.
//
// Solidity: function publicVar_() constant returns(address)
func (_PODEX *PODEXSession) PublicVar() (common.Address, error) {
	return _PODEX.Contract.PublicVar(&_PODEX.CallOpts)
}

// PublicVar is a free data retrieval call binding the contract method 0x0a6c93e0.
//
// Solidity: function publicVar_() constant returns(address)
func (_PODEX *PODEXCallerSession) PublicVar() (common.Address, error) {
	return _PODEX.Contract.PublicVar(&_PODEX.CallOpts)
}

// S is a free data retrieval call binding the contract method 0xc6b46963.
//
// Solidity: function s_() constant returns(uint64)
func (_PODEX *PODEXCaller) S(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "s_")
	return *ret0, err
}

// S is a free data retrieval call binding the contract method 0xc6b46963.
//
// Solidity: function s_() constant returns(uint64)
func (_PODEX *PODEXSession) S() (uint64, error) {
	return _PODEX.Contract.S(&_PODEX.CallOpts)
}

// S is a free data retrieval call binding the contract method 0xc6b46963.
//
// Solidity: function s_() constant returns(uint64)
func (_PODEX *PODEXCallerSession) S() (uint64, error) {
	return _PODEX.Contract.S(&_PODEX.CallOpts)
}

// T1 is a free data retrieval call binding the contract method 0xfb5343f3.
//
// Solidity: function t1() constant returns(uint256)
func (_PODEX *PODEXCaller) T1(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "t1")
	return *ret0, err
}

// T1 is a free data retrieval call binding the contract method 0xfb5343f3.
//
// Solidity: function t1() constant returns(uint256)
func (_PODEX *PODEXSession) T1() (*big.Int, error) {
	return _PODEX.Contract.T1(&_PODEX.CallOpts)
}

// T1 is a free data retrieval call binding the contract method 0xfb5343f3.
//
// Solidity: function t1() constant returns(uint256)
func (_PODEX *PODEXCallerSession) T1() (*big.Int, error) {
	return _PODEX.Contract.T1(&_PODEX.CallOpts)
}

// T3 is a free data retrieval call binding the contract method 0xcfad78b1.
//
// Solidity: function t3() constant returns(uint256)
func (_PODEX *PODEXCaller) T3(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "t3")
	return *ret0, err
}

// T3 is a free data retrieval call binding the contract method 0xcfad78b1.
//
// Solidity: function t3() constant returns(uint256)
func (_PODEX *PODEXSession) T3() (*big.Int, error) {
	return _PODEX.Contract.T3(&_PODEX.CallOpts)
}

// T3 is a free data retrieval call binding the contract method 0xcfad78b1.
//
// Solidity: function t3() constant returns(uint256)
func (_PODEX *PODEXCallerSession) T3() (*big.Int, error) {
	return _PODEX.Contract.T3(&_PODEX.CallOpts)
}

// T4 is a free data retrieval call binding the contract method 0xfeae062d.
//
// Solidity: function t4() constant returns(uint256)
func (_PODEX *PODEXCaller) T4(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "t4")
	return *ret0, err
}

// T4 is a free data retrieval call binding the contract method 0xfeae062d.
//
// Solidity: function t4() constant returns(uint256)
func (_PODEX *PODEXSession) T4() (*big.Int, error) {
	return _PODEX.Contract.T4(&_PODEX.CallOpts)
}

// T4 is a free data retrieval call binding the contract method 0xfeae062d.
//
// Solidity: function t4() constant returns(uint256)
func (_PODEX *PODEXCallerSession) T4() (*big.Int, error) {
	return _PODEX.Contract.T4(&_PODEX.CallOpts)
}

// VrfyPath is a free data retrieval call binding the contract method 0xd54ee12d.
//
// Solidity: function vrfyPath(bytes32 _x, bytes32 _y, uint64 _ij, uint64 _ns, bytes32 _root, bytes32[] _mkl_path) constant returns(bool)
func (_PODEX *PODEXCaller) VrfyPath(opts *bind.CallOpts, _x [32]byte, _y [32]byte, _ij uint64, _ns uint64, _root [32]byte, _mkl_path [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "vrfyPath", _x, _y, _ij, _ns, _root, _mkl_path)
	return *ret0, err
}

// VrfyPath is a free data retrieval call binding the contract method 0xd54ee12d.
//
// Solidity: function vrfyPath(bytes32 _x, bytes32 _y, uint64 _ij, uint64 _ns, bytes32 _root, bytes32[] _mkl_path) constant returns(bool)
func (_PODEX *PODEXSession) VrfyPath(_x [32]byte, _y [32]byte, _ij uint64, _ns uint64, _root [32]byte, _mkl_path [][32]byte) (bool, error) {
	return _PODEX.Contract.VrfyPath(&_PODEX.CallOpts, _x, _y, _ij, _ns, _root, _mkl_path)
}

// VrfyPath is a free data retrieval call binding the contract method 0xd54ee12d.
//
// Solidity: function vrfyPath(bytes32 _x, bytes32 _y, uint64 _ij, uint64 _ns, bytes32 _root, bytes32[] _mkl_path) constant returns(bool)
func (_PODEX *PODEXCallerSession) VrfyPath(_x [32]byte, _y [32]byte, _ij uint64, _ns uint64, _root [32]byte, _mkl_path [][32]byte) (bool, error) {
	return _PODEX.Contract.VrfyPath(&_PODEX.CallOpts, _x, _y, _ij, _ns, _root, _mkl_path)
}

// VrfyProofBatch2 is a free data retrieval call binding the contract method 0x16ddfc3c.
//
// Solidity: function vrfyProofBatch2(uint64 count, uint64 s, bytes32 seed0, bytes32 seed2, uint256 sigma_vw) constant returns(bool)
func (_PODEX *PODEXCaller) VrfyProofBatch2(opts *bind.CallOpts, count uint64, s uint64, seed0 [32]byte, seed2 [32]byte, sigma_vw *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "vrfyProofBatch2", count, s, seed0, seed2, sigma_vw)
	return *ret0, err
}

// VrfyProofBatch2 is a free data retrieval call binding the contract method 0x16ddfc3c.
//
// Solidity: function vrfyProofBatch2(uint64 count, uint64 s, bytes32 seed0, bytes32 seed2, uint256 sigma_vw) constant returns(bool)
func (_PODEX *PODEXSession) VrfyProofBatch2(count uint64, s uint64, seed0 [32]byte, seed2 [32]byte, sigma_vw *big.Int) (bool, error) {
	return _PODEX.Contract.VrfyProofBatch2(&_PODEX.CallOpts, count, s, seed0, seed2, sigma_vw)
}

// VrfyProofBatch2 is a free data retrieval call binding the contract method 0x16ddfc3c.
//
// Solidity: function vrfyProofBatch2(uint64 count, uint64 s, bytes32 seed0, bytes32 seed2, uint256 sigma_vw) constant returns(bool)
func (_PODEX *PODEXCallerSession) VrfyProofBatch2(count uint64, s uint64, seed0 [32]byte, seed2 [32]byte, sigma_vw *big.Int) (bool, error) {
	return _PODEX.Contract.VrfyProofBatch2(&_PODEX.CallOpts, count, s, seed0, seed2, sigma_vw)
}

// VrfyProofBatch3 is a free data retrieval call binding the contract method 0xaa7e4af9.
//
// Solidity: function vrfyProofBatch3(uint256[2] _r_u0d, uint256[2] _r_u0_x0_lgs, uint256 _s_d, uint256 _s_x0_lgs) constant returns(bool)
func (_PODEX *PODEXCaller) VrfyProofBatch3(opts *bind.CallOpts, _r_u0d [2]*big.Int, _r_u0_x0_lgs [2]*big.Int, _s_d *big.Int, _s_x0_lgs *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "vrfyProofBatch3", _r_u0d, _r_u0_x0_lgs, _s_d, _s_x0_lgs)
	return *ret0, err
}

// VrfyProofBatch3 is a free data retrieval call binding the contract method 0xaa7e4af9.
//
// Solidity: function vrfyProofBatch3(uint256[2] _r_u0d, uint256[2] _r_u0_x0_lgs, uint256 _s_d, uint256 _s_x0_lgs) constant returns(bool)
func (_PODEX *PODEXSession) VrfyProofBatch3(_r_u0d [2]*big.Int, _r_u0_x0_lgs [2]*big.Int, _s_d *big.Int, _s_x0_lgs *big.Int) (bool, error) {
	return _PODEX.Contract.VrfyProofBatch3(&_PODEX.CallOpts, _r_u0d, _r_u0_x0_lgs, _s_d, _s_x0_lgs)
}

// VrfyProofBatch3 is a free data retrieval call binding the contract method 0xaa7e4af9.
//
// Solidity: function vrfyProofBatch3(uint256[2] _r_u0d, uint256[2] _r_u0_x0_lgs, uint256 _s_d, uint256 _s_x0_lgs) constant returns(bool)
func (_PODEX *PODEXCallerSession) VrfyProofBatch3(_r_u0d [2]*big.Int, _r_u0_x0_lgs [2]*big.Int, _s_d *big.Int, _s_x0_lgs *big.Int) (bool, error) {
	return _PODEX.Contract.VrfyProofBatch3(&_PODEX.CallOpts, _r_u0d, _r_u0_x0_lgs, _s_d, _s_x0_lgs)
}

// VrfyProofVRF is a free data retrieval call binding the contract method 0xaeb8f1fa.
//
// Solidity: function vrfyProofVRF(uint256[2] _g_exp_r, uint256 _s_r) constant returns(bool)
func (_PODEX *PODEXCaller) VrfyProofVRF(opts *bind.CallOpts, _g_exp_r [2]*big.Int, _s_r *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "vrfyProofVRF", _g_exp_r, _s_r)
	return *ret0, err
}

// VrfyProofVRF is a free data retrieval call binding the contract method 0xaeb8f1fa.
//
// Solidity: function vrfyProofVRF(uint256[2] _g_exp_r, uint256 _s_r) constant returns(bool)
func (_PODEX *PODEXSession) VrfyProofVRF(_g_exp_r [2]*big.Int, _s_r *big.Int) (bool, error) {
	return _PODEX.Contract.VrfyProofVRF(&_PODEX.CallOpts, _g_exp_r, _s_r)
}

// VrfyProofVRF is a free data retrieval call binding the contract method 0xaeb8f1fa.
//
// Solidity: function vrfyProofVRF(uint256[2] _g_exp_r, uint256 _s_r) constant returns(bool)
func (_PODEX *PODEXCallerSession) VrfyProofVRF(_g_exp_r [2]*big.Int, _s_r *big.Int) (bool, error) {
	return _PODEX.Contract.VrfyProofVRF(&_PODEX.CallOpts, _g_exp_r, _s_r)
}

// BuyerDeposit is a paid mutator transaction binding the contract method 0x7ffdf46b.
//
// Solidity: function buyerDeposit(address _to) returns()
func (_PODEX *PODEXTransactor) BuyerDeposit(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "buyerDeposit", _to)
}

// BuyerDeposit is a paid mutator transaction binding the contract method 0x7ffdf46b.
//
// Solidity: function buyerDeposit(address _to) returns()
func (_PODEX *PODEXSession) BuyerDeposit(_to common.Address) (*types.Transaction, error) {
	return _PODEX.Contract.BuyerDeposit(&_PODEX.TransactOpts, _to)
}

// BuyerDeposit is a paid mutator transaction binding the contract method 0x7ffdf46b.
//
// Solidity: function buyerDeposit(address _to) returns()
func (_PODEX *PODEXTransactorSession) BuyerDeposit(_to common.Address) (*types.Transaction, error) {
	return _PODEX.Contract.BuyerDeposit(&_PODEX.TransactOpts, _to)
}

// BuyerUnDeposit is a paid mutator transaction binding the contract method 0x87faeac1.
//
// Solidity: function buyerUnDeposit(address _to) returns()
func (_PODEX *PODEXTransactor) BuyerUnDeposit(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "buyerUnDeposit", _to)
}

// BuyerUnDeposit is a paid mutator transaction binding the contract method 0x87faeac1.
//
// Solidity: function buyerUnDeposit(address _to) returns()
func (_PODEX *PODEXSession) BuyerUnDeposit(_to common.Address) (*types.Transaction, error) {
	return _PODEX.Contract.BuyerUnDeposit(&_PODEX.TransactOpts, _to)
}

// BuyerUnDeposit is a paid mutator transaction binding the contract method 0x87faeac1.
//
// Solidity: function buyerUnDeposit(address _to) returns()
func (_PODEX *PODEXTransactorSession) BuyerUnDeposit(_to common.Address) (*types.Transaction, error) {
	return _PODEX.Contract.BuyerUnDeposit(&_PODEX.TransactOpts, _to)
}

// ClaimBatch1 is a paid mutator transaction binding the contract method 0x0846fa02.
//
// Solidity: function claimBatch1(address _a, uint256 _sid, uint64 _i, uint64 _j, uint256 _tx, uint256 _ty, bytes32[] _mkl_path, uint64 _sCnt) returns()
func (_PODEX *PODEXTransactor) ClaimBatch1(opts *bind.TransactOpts, _a common.Address, _sid *big.Int, _i uint64, _j uint64, _tx *big.Int, _ty *big.Int, _mkl_path [][32]byte, _sCnt uint64) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "claimBatch1", _a, _sid, _i, _j, _tx, _ty, _mkl_path, _sCnt)
}

// ClaimBatch1 is a paid mutator transaction binding the contract method 0x0846fa02.
//
// Solidity: function claimBatch1(address _a, uint256 _sid, uint64 _i, uint64 _j, uint256 _tx, uint256 _ty, bytes32[] _mkl_path, uint64 _sCnt) returns()
func (_PODEX *PODEXSession) ClaimBatch1(_a common.Address, _sid *big.Int, _i uint64, _j uint64, _tx *big.Int, _ty *big.Int, _mkl_path [][32]byte, _sCnt uint64) (*types.Transaction, error) {
	return _PODEX.Contract.ClaimBatch1(&_PODEX.TransactOpts, _a, _sid, _i, _j, _tx, _ty, _mkl_path, _sCnt)
}

// ClaimBatch1 is a paid mutator transaction binding the contract method 0x0846fa02.
//
// Solidity: function claimBatch1(address _a, uint256 _sid, uint64 _i, uint64 _j, uint256 _tx, uint256 _ty, bytes32[] _mkl_path, uint64 _sCnt) returns()
func (_PODEX *PODEXTransactorSession) ClaimBatch1(_a common.Address, _sid *big.Int, _i uint64, _j uint64, _tx *big.Int, _ty *big.Int, _mkl_path [][32]byte, _sCnt uint64) (*types.Transaction, error) {
	return _PODEX.Contract.ClaimBatch1(&_PODEX.TransactOpts, _a, _sid, _i, _j, _tx, _ty, _mkl_path, _sCnt)
}

// Publish is a paid mutator transaction binding the contract method 0x56e25912.
//
// Solidity: function publish(uint64 _size, uint64 _s, uint64 _n, uint256 _sigma_mkl_root, uint256 _vrf_meta_digest, uint256 _blt_type) returns()
func (_PODEX *PODEXTransactor) Publish(opts *bind.TransactOpts, _size uint64, _s uint64, _n uint64, _sigma_mkl_root *big.Int, _vrf_meta_digest *big.Int, _blt_type *big.Int) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "publish", _size, _s, _n, _sigma_mkl_root, _vrf_meta_digest, _blt_type)
}

// Publish is a paid mutator transaction binding the contract method 0x56e25912.
//
// Solidity: function publish(uint64 _size, uint64 _s, uint64 _n, uint256 _sigma_mkl_root, uint256 _vrf_meta_digest, uint256 _blt_type) returns()
func (_PODEX *PODEXSession) Publish(_size uint64, _s uint64, _n uint64, _sigma_mkl_root *big.Int, _vrf_meta_digest *big.Int, _blt_type *big.Int) (*types.Transaction, error) {
	return _PODEX.Contract.Publish(&_PODEX.TransactOpts, _size, _s, _n, _sigma_mkl_root, _vrf_meta_digest, _blt_type)
}

// Publish is a paid mutator transaction binding the contract method 0x56e25912.
//
// Solidity: function publish(uint64 _size, uint64 _s, uint64 _n, uint256 _sigma_mkl_root, uint256 _vrf_meta_digest, uint256 _blt_type) returns()
func (_PODEX *PODEXTransactorSession) Publish(_size uint64, _s uint64, _n uint64, _sigma_mkl_root *big.Int, _vrf_meta_digest *big.Int, _blt_type *big.Int) (*types.Transaction, error) {
	return _PODEX.Contract.Publish(&_PODEX.TransactOpts, _size, _s, _n, _sigma_mkl_root, _vrf_meta_digest, _blt_type)
}

// SettleBatch1Deal is a paid mutator transaction binding the contract method 0xdf480a3e.
//
// Solidity: function settleBatch1Deal(address _a, address _b, uint256 _sid) returns()
func (_PODEX *PODEXTransactor) SettleBatch1Deal(opts *bind.TransactOpts, _a common.Address, _b common.Address, _sid *big.Int) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "settleBatch1Deal", _a, _b, _sid)
}

// SettleBatch1Deal is a paid mutator transaction binding the contract method 0xdf480a3e.
//
// Solidity: function settleBatch1Deal(address _a, address _b, uint256 _sid) returns()
func (_PODEX *PODEXSession) SettleBatch1Deal(_a common.Address, _b common.Address, _sid *big.Int) (*types.Transaction, error) {
	return _PODEX.Contract.SettleBatch1Deal(&_PODEX.TransactOpts, _a, _b, _sid)
}

// SettleBatch1Deal is a paid mutator transaction binding the contract method 0xdf480a3e.
//
// Solidity: function settleBatch1Deal(address _a, address _b, uint256 _sid) returns()
func (_PODEX *PODEXTransactorSession) SettleBatch1Deal(_a common.Address, _b common.Address, _sid *big.Int) (*types.Transaction, error) {
	return _PODEX.Contract.SettleBatch1Deal(&_PODEX.TransactOpts, _a, _b, _sid)
}

// SubmitProofBatch1 is a paid mutator transaction binding the contract method 0xbd3aa96d.
//
// Solidity: function submitProofBatch1(bytes32 _seed0, uint256 _sid, address _b, bytes32 _seed2, bytes32 _k_mkl_root, uint64 _count, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_PODEX *PODEXTransactor) SubmitProofBatch1(opts *bind.TransactOpts, _seed0 [32]byte, _sid *big.Int, _b common.Address, _seed2 [32]byte, _k_mkl_root [32]byte, _count uint64, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "submitProofBatch1", _seed0, _sid, _b, _seed2, _k_mkl_root, _count, _price, _expireAt, _sig)
}

// SubmitProofBatch1 is a paid mutator transaction binding the contract method 0xbd3aa96d.
//
// Solidity: function submitProofBatch1(bytes32 _seed0, uint256 _sid, address _b, bytes32 _seed2, bytes32 _k_mkl_root, uint64 _count, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_PODEX *PODEXSession) SubmitProofBatch1(_seed0 [32]byte, _sid *big.Int, _b common.Address, _seed2 [32]byte, _k_mkl_root [32]byte, _count uint64, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofBatch1(&_PODEX.TransactOpts, _seed0, _sid, _b, _seed2, _k_mkl_root, _count, _price, _expireAt, _sig)
}

// SubmitProofBatch1 is a paid mutator transaction binding the contract method 0xbd3aa96d.
//
// Solidity: function submitProofBatch1(bytes32 _seed0, uint256 _sid, address _b, bytes32 _seed2, bytes32 _k_mkl_root, uint64 _count, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_PODEX *PODEXTransactorSession) SubmitProofBatch1(_seed0 [32]byte, _sid *big.Int, _b common.Address, _seed2 [32]byte, _k_mkl_root [32]byte, _count uint64, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofBatch1(&_PODEX.TransactOpts, _seed0, _sid, _b, _seed2, _k_mkl_root, _count, _price, _expireAt, _sig)
}

// SubmitProofBatch2 is a paid mutator transaction binding the contract method 0x0bddf632.
//
// Solidity: function submitProofBatch2(bytes32 _seed0, uint64 _sCnt, uint256 _sid, address _b, bytes32 _seed2, uint256 _sigma_vw, uint64 _count, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_PODEX *PODEXTransactor) SubmitProofBatch2(opts *bind.TransactOpts, _seed0 [32]byte, _sCnt uint64, _sid *big.Int, _b common.Address, _seed2 [32]byte, _sigma_vw *big.Int, _count uint64, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "submitProofBatch2", _seed0, _sCnt, _sid, _b, _seed2, _sigma_vw, _count, _price, _expireAt, _sig)
}

// SubmitProofBatch2 is a paid mutator transaction binding the contract method 0x0bddf632.
//
// Solidity: function submitProofBatch2(bytes32 _seed0, uint64 _sCnt, uint256 _sid, address _b, bytes32 _seed2, uint256 _sigma_vw, uint64 _count, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_PODEX *PODEXSession) SubmitProofBatch2(_seed0 [32]byte, _sCnt uint64, _sid *big.Int, _b common.Address, _seed2 [32]byte, _sigma_vw *big.Int, _count uint64, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofBatch2(&_PODEX.TransactOpts, _seed0, _sCnt, _sid, _b, _seed2, _sigma_vw, _count, _price, _expireAt, _sig)
}

// SubmitProofBatch2 is a paid mutator transaction binding the contract method 0x0bddf632.
//
// Solidity: function submitProofBatch2(bytes32 _seed0, uint64 _sCnt, uint256 _sid, address _b, bytes32 _seed2, uint256 _sigma_vw, uint64 _count, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_PODEX *PODEXTransactorSession) SubmitProofBatch2(_seed0 [32]byte, _sCnt uint64, _sid *big.Int, _b common.Address, _seed2 [32]byte, _sigma_vw *big.Int, _count uint64, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofBatch2(&_PODEX.TransactOpts, _seed0, _sCnt, _sid, _b, _seed2, _sigma_vw, _count, _price, _expireAt, _sig)
}

// SubmitProofBatch3 is a paid mutator transaction binding the contract method 0x21d13e3a.
//
// Solidity: function submitProofBatch3(uint256 _s_d, uint256 _s_x0_lgs, uint256 _sid, address _b, uint256[2] _r_u0_x0_lgs, uint256[2] _r_u0d, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_PODEX *PODEXTransactor) SubmitProofBatch3(opts *bind.TransactOpts, _s_d *big.Int, _s_x0_lgs *big.Int, _sid *big.Int, _b common.Address, _r_u0_x0_lgs [2]*big.Int, _r_u0d [2]*big.Int, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "submitProofBatch3", _s_d, _s_x0_lgs, _sid, _b, _r_u0_x0_lgs, _r_u0d, _price, _expireAt, _sig)
}

// SubmitProofBatch3 is a paid mutator transaction binding the contract method 0x21d13e3a.
//
// Solidity: function submitProofBatch3(uint256 _s_d, uint256 _s_x0_lgs, uint256 _sid, address _b, uint256[2] _r_u0_x0_lgs, uint256[2] _r_u0d, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_PODEX *PODEXSession) SubmitProofBatch3(_s_d *big.Int, _s_x0_lgs *big.Int, _sid *big.Int, _b common.Address, _r_u0_x0_lgs [2]*big.Int, _r_u0d [2]*big.Int, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofBatch3(&_PODEX.TransactOpts, _s_d, _s_x0_lgs, _sid, _b, _r_u0_x0_lgs, _r_u0d, _price, _expireAt, _sig)
}

// SubmitProofBatch3 is a paid mutator transaction binding the contract method 0x21d13e3a.
//
// Solidity: function submitProofBatch3(uint256 _s_d, uint256 _s_x0_lgs, uint256 _sid, address _b, uint256[2] _r_u0_x0_lgs, uint256[2] _r_u0d, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_PODEX *PODEXTransactorSession) SubmitProofBatch3(_s_d *big.Int, _s_x0_lgs *big.Int, _sid *big.Int, _b common.Address, _r_u0_x0_lgs [2]*big.Int, _r_u0d [2]*big.Int, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofBatch3(&_PODEX.TransactOpts, _s_d, _s_x0_lgs, _sid, _b, _r_u0_x0_lgs, _r_u0d, _price, _expireAt, _sig)
}

// SubmitProofVRF is a paid mutator transaction binding the contract method 0x50f9b445.
//
// Solidity: function submitProofVRF(uint256 _s_r, uint256 _sid, address _b, uint256[2] _g_exp_r, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_PODEX *PODEXTransactor) SubmitProofVRF(opts *bind.TransactOpts, _s_r *big.Int, _sid *big.Int, _b common.Address, _g_exp_r [2]*big.Int, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "submitProofVRF", _s_r, _sid, _b, _g_exp_r, _price, _expireAt, _sig)
}

// SubmitProofVRF is a paid mutator transaction binding the contract method 0x50f9b445.
//
// Solidity: function submitProofVRF(uint256 _s_r, uint256 _sid, address _b, uint256[2] _g_exp_r, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_PODEX *PODEXSession) SubmitProofVRF(_s_r *big.Int, _sid *big.Int, _b common.Address, _g_exp_r [2]*big.Int, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofVRF(&_PODEX.TransactOpts, _s_r, _sid, _b, _g_exp_r, _price, _expireAt, _sig)
}

// SubmitProofVRF is a paid mutator transaction binding the contract method 0x50f9b445.
//
// Solidity: function submitProofVRF(uint256 _s_r, uint256 _sid, address _b, uint256[2] _g_exp_r, uint256 _price, uint256 _expireAt, bytes _sig) returns()
func (_PODEX *PODEXTransactorSession) SubmitProofVRF(_s_r *big.Int, _sid *big.Int, _b common.Address, _g_exp_r [2]*big.Int, _price *big.Int, _expireAt *big.Int, _sig []byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofVRF(&_PODEX.TransactOpts, _s_r, _sid, _b, _g_exp_r, _price, _expireAt, _sig)
}

// UnPublish is a paid mutator transaction binding the contract method 0x52bb0780.
//
// Solidity: function unPublish(bytes32 _bltKey) returns()
func (_PODEX *PODEXTransactor) UnPublish(opts *bind.TransactOpts, _bltKey [32]byte) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "unPublish", _bltKey)
}

// UnPublish is a paid mutator transaction binding the contract method 0x52bb0780.
//
// Solidity: function unPublish(bytes32 _bltKey) returns()
func (_PODEX *PODEXSession) UnPublish(_bltKey [32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.UnPublish(&_PODEX.TransactOpts, _bltKey)
}

// UnPublish is a paid mutator transaction binding the contract method 0x52bb0780.
//
// Solidity: function unPublish(bytes32 _bltKey) returns()
func (_PODEX *PODEXTransactorSession) UnPublish(_bltKey [32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.UnPublish(&_PODEX.TransactOpts, _bltKey)
}

// WithdrawA is a paid mutator transaction binding the contract method 0x9ecd2944.
//
// Solidity: function withdrawA(bytes32 _bltKey) returns()
func (_PODEX *PODEXTransactor) WithdrawA(opts *bind.TransactOpts, _bltKey [32]byte) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "withdrawA", _bltKey)
}

// WithdrawA is a paid mutator transaction binding the contract method 0x9ecd2944.
//
// Solidity: function withdrawA(bytes32 _bltKey) returns()
func (_PODEX *PODEXSession) WithdrawA(_bltKey [32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.WithdrawA(&_PODEX.TransactOpts, _bltKey)
}

// WithdrawA is a paid mutator transaction binding the contract method 0x9ecd2944.
//
// Solidity: function withdrawA(bytes32 _bltKey) returns()
func (_PODEX *PODEXTransactorSession) WithdrawA(_bltKey [32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.WithdrawA(&_PODEX.TransactOpts, _bltKey)
}

// WithdrawB is a paid mutator transaction binding the contract method 0x2927251f.
//
// Solidity: function withdrawB(address _to) returns()
func (_PODEX *PODEXTransactor) WithdrawB(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "withdrawB", _to)
}

// WithdrawB is a paid mutator transaction binding the contract method 0x2927251f.
//
// Solidity: function withdrawB(address _to) returns()
func (_PODEX *PODEXSession) WithdrawB(_to common.Address) (*types.Transaction, error) {
	return _PODEX.Contract.WithdrawB(&_PODEX.TransactOpts, _to)
}

// WithdrawB is a paid mutator transaction binding the contract method 0x2927251f.
//
// Solidity: function withdrawB(address _to) returns()
func (_PODEX *PODEXTransactorSession) WithdrawB(_to common.Address) (*types.Transaction, error) {
	return _PODEX.Contract.WithdrawB(&_PODEX.TransactOpts, _to)
}

// PODEXOnBatch1ClaimIterator is returned from FilterOnBatch1Claim and is used to iterate over the raw logs and unpacked data for OnBatch1Claim events raised by the PODEX contract.
type PODEXOnBatch1ClaimIterator struct {
	Event *PODEXOnBatch1Claim // Event containing the contract specifics and raw log

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
func (it *PODEXOnBatch1ClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PODEXOnBatch1Claim)
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
		it.Event = new(PODEXOnBatch1Claim)
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
func (it *PODEXOnBatch1ClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PODEXOnBatch1ClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PODEXOnBatch1Claim represents a OnBatch1Claim event raised by the PODEX contract.
type PODEXOnBatch1Claim struct {
	A   common.Address
	B   common.Address
	Sid *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOnBatch1Claim is a free log retrieval operation binding the contract event 0x2502be5e06bef8815dce6e5cd8a4716834c23245c72cfeba602441f05744e6f4.
//
// Solidity: event OnBatch1Claim(address indexed _a, address indexed _b, uint256 indexed _sid)
func (_PODEX *PODEXFilterer) FilterOnBatch1Claim(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*PODEXOnBatch1ClaimIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.FilterLogs(opts, "OnBatch1Claim", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &PODEXOnBatch1ClaimIterator{contract: _PODEX.contract, event: "OnBatch1Claim", logs: logs, sub: sub}, nil
}

// WatchOnBatch1Claim is a free log subscription operation binding the contract event 0x2502be5e06bef8815dce6e5cd8a4716834c23245c72cfeba602441f05744e6f4.
//
// Solidity: event OnBatch1Claim(address indexed _a, address indexed _b, uint256 indexed _sid)
func (_PODEX *PODEXFilterer) WatchOnBatch1Claim(opts *bind.WatchOpts, sink chan<- *PODEXOnBatch1Claim, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.WatchLogs(opts, "OnBatch1Claim", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PODEXOnBatch1Claim)
				if err := _PODEX.contract.UnpackLog(event, "OnBatch1Claim", log); err != nil {
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

// PODEXOnBatch1DealIterator is returned from FilterOnBatch1Deal and is used to iterate over the raw logs and unpacked data for OnBatch1Deal events raised by the PODEX contract.
type PODEXOnBatch1DealIterator struct {
	Event *PODEXOnBatch1Deal // Event containing the contract specifics and raw log

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
func (it *PODEXOnBatch1DealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PODEXOnBatch1Deal)
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
		it.Event = new(PODEXOnBatch1Deal)
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
func (it *PODEXOnBatch1DealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PODEXOnBatch1DealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PODEXOnBatch1Deal represents a OnBatch1Deal event raised by the PODEX contract.
type PODEXOnBatch1Deal struct {
	A     common.Address
	B     common.Address
	Sid   *big.Int
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOnBatch1Deal is a free log retrieval operation binding the contract event 0x97ba31aaf3572239af72fddb099e107414e93b2801a058b873dfc7ecaa777a47.
//
// Solidity: event OnBatch1Deal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _price)
func (_PODEX *PODEXFilterer) FilterOnBatch1Deal(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*PODEXOnBatch1DealIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.FilterLogs(opts, "OnBatch1Deal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &PODEXOnBatch1DealIterator{contract: _PODEX.contract, event: "OnBatch1Deal", logs: logs, sub: sub}, nil
}

// WatchOnBatch1Deal is a free log subscription operation binding the contract event 0x97ba31aaf3572239af72fddb099e107414e93b2801a058b873dfc7ecaa777a47.
//
// Solidity: event OnBatch1Deal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _price)
func (_PODEX *PODEXFilterer) WatchOnBatch1Deal(opts *bind.WatchOpts, sink chan<- *PODEXOnBatch1Deal, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.WatchLogs(opts, "OnBatch1Deal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PODEXOnBatch1Deal)
				if err := _PODEX.contract.UnpackLog(event, "OnBatch1Deal", log); err != nil {
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

// PODEXOnBatch1KeyIterator is returned from FilterOnBatch1Key and is used to iterate over the raw logs and unpacked data for OnBatch1Key events raised by the PODEX contract.
type PODEXOnBatch1KeyIterator struct {
	Event *PODEXOnBatch1Key // Event containing the contract specifics and raw log

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
func (it *PODEXOnBatch1KeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PODEXOnBatch1Key)
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
		it.Event = new(PODEXOnBatch1Key)
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
func (it *PODEXOnBatch1KeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PODEXOnBatch1KeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PODEXOnBatch1Key represents a OnBatch1Key event raised by the PODEX contract.
type PODEXOnBatch1Key struct {
	A     common.Address
	B     common.Address
	Sid   *big.Int
	Seed0 [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOnBatch1Key is a free log retrieval operation binding the contract event 0x08068d0fcd10dfa2c800de1d7ec458dcaab4a39626af415baf1d694f6a404484.
//
// Solidity: event OnBatch1Key(address indexed _a, address indexed _b, uint256 indexed _sid, bytes32 _seed0)
func (_PODEX *PODEXFilterer) FilterOnBatch1Key(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*PODEXOnBatch1KeyIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.FilterLogs(opts, "OnBatch1Key", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &PODEXOnBatch1KeyIterator{contract: _PODEX.contract, event: "OnBatch1Key", logs: logs, sub: sub}, nil
}

// WatchOnBatch1Key is a free log subscription operation binding the contract event 0x08068d0fcd10dfa2c800de1d7ec458dcaab4a39626af415baf1d694f6a404484.
//
// Solidity: event OnBatch1Key(address indexed _a, address indexed _b, uint256 indexed _sid, bytes32 _seed0)
func (_PODEX *PODEXFilterer) WatchOnBatch1Key(opts *bind.WatchOpts, sink chan<- *PODEXOnBatch1Key, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.WatchLogs(opts, "OnBatch1Key", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PODEXOnBatch1Key)
				if err := _PODEX.contract.UnpackLog(event, "OnBatch1Key", log); err != nil {
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

// PODEXOnBatch2DealIterator is returned from FilterOnBatch2Deal and is used to iterate over the raw logs and unpacked data for OnBatch2Deal events raised by the PODEX contract.
type PODEXOnBatch2DealIterator struct {
	Event *PODEXOnBatch2Deal // Event containing the contract specifics and raw log

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
func (it *PODEXOnBatch2DealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PODEXOnBatch2Deal)
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
		it.Event = new(PODEXOnBatch2Deal)
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
func (it *PODEXOnBatch2DealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PODEXOnBatch2DealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PODEXOnBatch2Deal represents a OnBatch2Deal event raised by the PODEX contract.
type PODEXOnBatch2Deal struct {
	A     common.Address
	B     common.Address
	Sid   *big.Int
	Seed0 [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOnBatch2Deal is a free log retrieval operation binding the contract event 0xc8dff79bd96b1390c89f4aa1d31095efeafc211fb5bb0bf27d1b557abcab4920.
//
// Solidity: event OnBatch2Deal(address indexed _a, address indexed _b, uint256 indexed _sid, bytes32 _seed0)
func (_PODEX *PODEXFilterer) FilterOnBatch2Deal(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*PODEXOnBatch2DealIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.FilterLogs(opts, "OnBatch2Deal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &PODEXOnBatch2DealIterator{contract: _PODEX.contract, event: "OnBatch2Deal", logs: logs, sub: sub}, nil
}

// WatchOnBatch2Deal is a free log subscription operation binding the contract event 0xc8dff79bd96b1390c89f4aa1d31095efeafc211fb5bb0bf27d1b557abcab4920.
//
// Solidity: event OnBatch2Deal(address indexed _a, address indexed _b, uint256 indexed _sid, bytes32 _seed0)
func (_PODEX *PODEXFilterer) WatchOnBatch2Deal(opts *bind.WatchOpts, sink chan<- *PODEXOnBatch2Deal, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.WatchLogs(opts, "OnBatch2Deal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PODEXOnBatch2Deal)
				if err := _PODEX.contract.UnpackLog(event, "OnBatch2Deal", log); err != nil {
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

// PODEXOnBatch3DealIterator is returned from FilterOnBatch3Deal and is used to iterate over the raw logs and unpacked data for OnBatch3Deal events raised by the PODEX contract.
type PODEXOnBatch3DealIterator struct {
	Event *PODEXOnBatch3Deal // Event containing the contract specifics and raw log

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
func (it *PODEXOnBatch3DealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PODEXOnBatch3Deal)
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
		it.Event = new(PODEXOnBatch3Deal)
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
func (it *PODEXOnBatch3DealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PODEXOnBatch3DealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PODEXOnBatch3Deal represents a OnBatch3Deal event raised by the PODEX contract.
type PODEXOnBatch3Deal struct {
	A     common.Address
	B     common.Address
	Sid   *big.Int
	D     *big.Int
	X0Lgs *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOnBatch3Deal is a free log retrieval operation binding the contract event 0xd07f399bfee8ccf0cbabea8dd2888393732aeac7ad3aca575931ef4b0d7d66e3.
//
// Solidity: event OnBatch3Deal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _d, uint256 _x0_lgs)
func (_PODEX *PODEXFilterer) FilterOnBatch3Deal(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*PODEXOnBatch3DealIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.FilterLogs(opts, "OnBatch3Deal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &PODEXOnBatch3DealIterator{contract: _PODEX.contract, event: "OnBatch3Deal", logs: logs, sub: sub}, nil
}

// WatchOnBatch3Deal is a free log subscription operation binding the contract event 0xd07f399bfee8ccf0cbabea8dd2888393732aeac7ad3aca575931ef4b0d7d66e3.
//
// Solidity: event OnBatch3Deal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _d, uint256 _x0_lgs)
func (_PODEX *PODEXFilterer) WatchOnBatch3Deal(opts *bind.WatchOpts, sink chan<- *PODEXOnBatch3Deal, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.WatchLogs(opts, "OnBatch3Deal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PODEXOnBatch3Deal)
				if err := _PODEX.contract.UnpackLog(event, "OnBatch3Deal", log); err != nil {
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

// PODEXOnDealIterator is returned from FilterOnDeal and is used to iterate over the raw logs and unpacked data for OnDeal events raised by the PODEX contract.
type PODEXOnDealIterator struct {
	Event *PODEXOnDeal // Event containing the contract specifics and raw log

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
func (it *PODEXOnDealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PODEXOnDeal)
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
		it.Event = new(PODEXOnDeal)
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
func (it *PODEXOnDealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PODEXOnDealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PODEXOnDeal represents a OnDeal event raised by the PODEX contract.
type PODEXOnDeal struct {
	A     common.Address
	B     common.Address
	Sid   *big.Int
	Mode  uint8
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOnDeal is a free log retrieval operation binding the contract event 0x37a5d6eac0a1d706c2eeea2c832256ebf3253a7b42d5a6cedd60e67f4195449b.
//
// Solidity: event OnDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint8 _mode, uint256 _price)
func (_PODEX *PODEXFilterer) FilterOnDeal(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*PODEXOnDealIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.FilterLogs(opts, "OnDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &PODEXOnDealIterator{contract: _PODEX.contract, event: "OnDeal", logs: logs, sub: sub}, nil
}

// WatchOnDeal is a free log subscription operation binding the contract event 0x37a5d6eac0a1d706c2eeea2c832256ebf3253a7b42d5a6cedd60e67f4195449b.
//
// Solidity: event OnDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint8 _mode, uint256 _price)
func (_PODEX *PODEXFilterer) WatchOnDeal(opts *bind.WatchOpts, sink chan<- *PODEXOnDeal, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.WatchLogs(opts, "OnDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PODEXOnDeal)
				if err := _PODEX.contract.UnpackLog(event, "OnDeal", log); err != nil {
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

// PODEXOnVRFDealIterator is returned from FilterOnVRFDeal and is used to iterate over the raw logs and unpacked data for OnVRFDeal events raised by the PODEX contract.
type PODEXOnVRFDealIterator struct {
	Event *PODEXOnVRFDeal // Event containing the contract specifics and raw log

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
func (it *PODEXOnVRFDealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PODEXOnVRFDeal)
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
		it.Event = new(PODEXOnVRFDeal)
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
func (it *PODEXOnVRFDealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PODEXOnVRFDealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PODEXOnVRFDeal represents a OnVRFDeal event raised by the PODEX contract.
type PODEXOnVRFDeal struct {
	A   common.Address
	B   common.Address
	Sid *big.Int
	R   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOnVRFDeal is a free log retrieval operation binding the contract event 0x0cf813cbe764040d8f6a8bd61c65524fefa1b75e383c20d3673881e921173a52.
//
// Solidity: event OnVRFDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _r)
func (_PODEX *PODEXFilterer) FilterOnVRFDeal(opts *bind.FilterOpts, _a []common.Address, _b []common.Address, _sid []*big.Int) (*PODEXOnVRFDealIterator, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.FilterLogs(opts, "OnVRFDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return &PODEXOnVRFDealIterator{contract: _PODEX.contract, event: "OnVRFDeal", logs: logs, sub: sub}, nil
}

// WatchOnVRFDeal is a free log subscription operation binding the contract event 0x0cf813cbe764040d8f6a8bd61c65524fefa1b75e383c20d3673881e921173a52.
//
// Solidity: event OnVRFDeal(address indexed _a, address indexed _b, uint256 indexed _sid, uint256 _r)
func (_PODEX *PODEXFilterer) WatchOnVRFDeal(opts *bind.WatchOpts, sink chan<- *PODEXOnVRFDeal, _a []common.Address, _b []common.Address, _sid []*big.Int) (event.Subscription, error) {

	var _aRule []interface{}
	for _, _aItem := range _a {
		_aRule = append(_aRule, _aItem)
	}
	var _bRule []interface{}
	for _, _bItem := range _b {
		_bRule = append(_bRule, _bItem)
	}
	var _sidRule []interface{}
	for _, _sidItem := range _sid {
		_sidRule = append(_sidRule, _sidItem)
	}

	logs, sub, err := _PODEX.contract.WatchLogs(opts, "OnVRFDeal", _aRule, _bRule, _sidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PODEXOnVRFDeal)
				if err := _PODEX.contract.UnpackLog(event, "OnVRFDeal", log); err != nil {
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
