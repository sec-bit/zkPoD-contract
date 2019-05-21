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
const PODEXABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ecc_pub_u1_\",\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyerDeposits_\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"unDepositAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"t2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"s_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"t3\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bulletins_\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"size\",\"type\":\"uint64\"},{\"name\":\"s\",\"type\":\"uint64\"},{\"name\":\"n\",\"type\":\"uint64\"},{\"name\":\"sigma_mkl_root\",\"type\":\"uint256\"},{\"name\":\"vrf_meta_digest\",\"type\":\"uint256\"},{\"name\":\"pledge_value\",\"type\":\"uint256\"},{\"name\":\"unDepositAt\",\"type\":\"uint256\"},{\"name\":\"blt_type\",\"type\":\"uint8\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"t1\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_x\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_y\",\"type\":\"uint256\"}],\"name\":\"LogG1Point\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_size\",\"type\":\"uint64\"},{\"name\":\"_s\",\"type\":\"uint64\"},{\"name\":\"_n\",\"type\":\"uint64\"},{\"name\":\"_sigma_mkl_root\",\"type\":\"uint256\"},{\"name\":\"_vrf_meta_digest\",\"type\":\"uint256\"},{\"name\":\"_blt_type\",\"type\":\"uint256\"}],\"name\":\"publish\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bltKey\",\"type\":\"bytes32\"}],\"name\":\"unPublish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"buyerDeposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"buyerUnDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bltKey\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_seed0\",\"type\":\"bytes32\"},{\"name\":\"_sessionId\",\"type\":\"uint256\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_seed2\",\"type\":\"bytes32\"},{\"name\":\"_k_mkl_root\",\"type\":\"bytes32\"},{\"name\":\"_count\",\"type\":\"uint64\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"submitProofBatch1\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_sessionId\",\"type\":\"uint256\"},{\"name\":\"_i\",\"type\":\"uint64\"},{\"name\":\"_j\",\"type\":\"uint64\"},{\"name\":\"_tx\",\"type\":\"uint256\"},{\"name\":\"_ty\",\"type\":\"uint256\"},{\"name\":\"_mkl_path\",\"type\":\"bytes32[]\"},{\"name\":\"_sCnt\",\"type\":\"uint64\"}],\"name\":\"claimBatch1\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_seed0\",\"type\":\"bytes32\"},{\"name\":\"_sCnt\",\"type\":\"uint64\"},{\"name\":\"_sessionId\",\"type\":\"uint256\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_seed2\",\"type\":\"bytes32\"},{\"name\":\"_sigma_vw\",\"type\":\"uint256\"},{\"name\":\"_count\",\"type\":\"uint64\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_rs\",\"type\":\"bytes32[2]\"}],\"name\":\"submitProofBatch2\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_s_d\",\"type\":\"uint256\"},{\"name\":\"_s_x0_lgs\",\"type\":\"uint256\"},{\"name\":\"_sessionId\",\"type\":\"uint256\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_r_u0_x0_lgs\",\"type\":\"uint256[2]\"},{\"name\":\"_r_u0d\",\"type\":\"uint256[2]\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"submitProofBatch3\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_s_r\",\"type\":\"uint256\"},{\"name\":\"_sessionId\",\"type\":\"uint256\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_g_exp_r\",\"type\":\"uint256[2]\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_rs\",\"type\":\"bytes32[2]\"}],\"name\":\"submitProofVRF\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"count\",\"type\":\"uint64\"},{\"name\":\"s\",\"type\":\"uint64\"},{\"name\":\"seed0\",\"type\":\"bytes32\"},{\"name\":\"seed2\",\"type\":\"bytes32\"},{\"name\":\"sigma_vw\",\"type\":\"uint256\"}],\"name\":\"verifyProofBatch2\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_r_u0d\",\"type\":\"uint256[2]\"},{\"name\":\"_r_u0_x0_lgs\",\"type\":\"uint256[2]\"},{\"name\":\"_s_d\",\"type\":\"uint256\"},{\"name\":\"_s_x0_lgs\",\"type\":\"uint256\"}],\"name\":\"verifyProofBatch3\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_g_exp_r\",\"type\":\"uint256[2]\"},{\"name\":\"_s_r\",\"type\":\"uint256\"}],\"name\":\"verifyProofVRF\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"bytes32\"},{\"name\":\"_y\",\"type\":\"bytes32\"},{\"name\":\"_ij\",\"type\":\"uint64\"},{\"name\":\"_ns\",\"type\":\"uint64\"},{\"name\":\"_root\",\"type\":\"bytes32\"},{\"name\":\"_mkl_path\",\"type\":\"bytes32[]\"}],\"name\":\"verifyPath\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"bytes32\"},{\"name\":\"_y\",\"type\":\"bytes32\"}],\"name\":\"hashInSha3\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"bytes32\"},{\"name\":\"_y\",\"type\":\"uint64\"}],\"name\":\"hashInSha3\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"seed\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"chain\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_n\",\"type\":\"uint256\"}],\"name\":\"log2ub\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sessionId\",\"type\":\"uint256\"}],\"name\":\"getRecordBatch1\",\"outputs\":[{\"name\":\"seed0\",\"type\":\"bytes32\"},{\"name\":\"seed2\",\"type\":\"bytes32\"},{\"name\":\"k_mkl_root\",\"type\":\"bytes32\"},{\"name\":\"count\",\"type\":\"uint64\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"expireAt\",\"type\":\"uint256\"},{\"name\":\"submitAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sessionId\",\"type\":\"uint256\"}],\"name\":\"getRecordBatch2\",\"outputs\":[{\"name\":\"seed0\",\"type\":\"bytes32\"},{\"name\":\"submitAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sessionId\",\"type\":\"uint256\"}],\"name\":\"getRecordBatch3\",\"outputs\":[{\"name\":\"d\",\"type\":\"uint256\"},{\"name\":\"x0_lgs\",\"type\":\"uint256\"},{\"name\":\"submitAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"},{\"name\":\"_b\",\"type\":\"address\"},{\"name\":\"_sessionId\",\"type\":\"uint256\"}],\"name\":\"getRecordVRF\",\"outputs\":[{\"name\":\"r\",\"type\":\"uint256\"},{\"name\":\"submitAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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
// Solidity: function bulletins_(bytes32 ) constant returns(address owner, uint64 size, uint64 s, uint64 n, uint256 sigma_mkl_root, uint256 vrf_meta_digest, uint256 pledge_value, uint256 unDepositAt, uint8 blt_type, uint8 status)
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
	Status        uint8
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
		Status        uint8
	})
	out := ret
	err := _PODEX.contract.Call(opts, out, "bulletins_", arg0)
	return *ret, err
}

// Bulletins is a free data retrieval call binding the contract method 0xd9417785.
//
// Solidity: function bulletins_(bytes32 ) constant returns(address owner, uint64 size, uint64 s, uint64 n, uint256 sigma_mkl_root, uint256 vrf_meta_digest, uint256 pledge_value, uint256 unDepositAt, uint8 blt_type, uint8 status)
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
	Status        uint8
}, error) {
	return _PODEX.Contract.Bulletins(&_PODEX.CallOpts, arg0)
}

// Bulletins is a free data retrieval call binding the contract method 0xd9417785.
//
// Solidity: function bulletins_(bytes32 ) constant returns(address owner, uint64 size, uint64 s, uint64 n, uint256 sigma_mkl_root, uint256 vrf_meta_digest, uint256 pledge_value, uint256 unDepositAt, uint8 blt_type, uint8 status)
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
	Status        uint8
}, error) {
	return _PODEX.Contract.Bulletins(&_PODEX.CallOpts, arg0)
}

// BuyerDeposits is a free data retrieval call binding the contract method 0x8ca90a25.
//
// Solidity: function buyerDeposits_(address , address ) constant returns(uint256 value, uint256 unDepositAt, uint8 status)
func (_PODEX *PODEXCaller) BuyerDeposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Value       *big.Int
	UnDepositAt *big.Int
	Status      uint8
}, error) {
	ret := new(struct {
		Value       *big.Int
		UnDepositAt *big.Int
		Status      uint8
	})
	out := ret
	err := _PODEX.contract.Call(opts, out, "buyerDeposits_", arg0, arg1)
	return *ret, err
}

// BuyerDeposits is a free data retrieval call binding the contract method 0x8ca90a25.
//
// Solidity: function buyerDeposits_(address , address ) constant returns(uint256 value, uint256 unDepositAt, uint8 status)
func (_PODEX *PODEXSession) BuyerDeposits(arg0 common.Address, arg1 common.Address) (struct {
	Value       *big.Int
	UnDepositAt *big.Int
	Status      uint8
}, error) {
	return _PODEX.Contract.BuyerDeposits(&_PODEX.CallOpts, arg0, arg1)
}

// BuyerDeposits is a free data retrieval call binding the contract method 0x8ca90a25.
//
// Solidity: function buyerDeposits_(address , address ) constant returns(uint256 value, uint256 unDepositAt, uint8 status)
func (_PODEX *PODEXCallerSession) BuyerDeposits(arg0 common.Address, arg1 common.Address) (struct {
	Value       *big.Int
	UnDepositAt *big.Int
	Status      uint8
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

// EccPubU1 is a free data retrieval call binding the contract method 0x52bfce18.
//
// Solidity: function ecc_pub_u1_(uint256 ) constant returns(uint256 X, uint256 Y)
func (_PODEX *PODEXCaller) EccPubU1(opts *bind.CallOpts, arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	ret := new(struct {
		X *big.Int
		Y *big.Int
	})
	out := ret
	err := _PODEX.contract.Call(opts, out, "ecc_pub_u1_", arg0)
	return *ret, err
}

// EccPubU1 is a free data retrieval call binding the contract method 0x52bfce18.
//
// Solidity: function ecc_pub_u1_(uint256 ) constant returns(uint256 X, uint256 Y)
func (_PODEX *PODEXSession) EccPubU1(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _PODEX.Contract.EccPubU1(&_PODEX.CallOpts, arg0)
}

// EccPubU1 is a free data retrieval call binding the contract method 0x52bfce18.
//
// Solidity: function ecc_pub_u1_(uint256 ) constant returns(uint256 X, uint256 Y)
func (_PODEX *PODEXCallerSession) EccPubU1(arg0 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _PODEX.Contract.EccPubU1(&_PODEX.CallOpts, arg0)
}

// GetRecordBatch1 is a free data retrieval call binding the contract method 0x321d0cba.
//
// Solidity: function getRecordBatch1(address _a, address _b, uint256 _sessionId) constant returns(bytes32 seed0, bytes32 seed2, bytes32 k_mkl_root, uint64 count, uint256 price, uint256 expireAt, uint256 submitAt)
func (_PODEX *PODEXCaller) GetRecordBatch1(opts *bind.CallOpts, _a common.Address, _b common.Address, _sessionId *big.Int) (struct {
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
	err := _PODEX.contract.Call(opts, out, "getRecordBatch1", _a, _b, _sessionId)
	return *ret, err
}

// GetRecordBatch1 is a free data retrieval call binding the contract method 0x321d0cba.
//
// Solidity: function getRecordBatch1(address _a, address _b, uint256 _sessionId) constant returns(bytes32 seed0, bytes32 seed2, bytes32 k_mkl_root, uint64 count, uint256 price, uint256 expireAt, uint256 submitAt)
func (_PODEX *PODEXSession) GetRecordBatch1(_a common.Address, _b common.Address, _sessionId *big.Int) (struct {
	Seed0    [32]byte
	Seed2    [32]byte
	KMklRoot [32]byte
	Count    uint64
	Price    *big.Int
	ExpireAt *big.Int
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordBatch1(&_PODEX.CallOpts, _a, _b, _sessionId)
}

// GetRecordBatch1 is a free data retrieval call binding the contract method 0x321d0cba.
//
// Solidity: function getRecordBatch1(address _a, address _b, uint256 _sessionId) constant returns(bytes32 seed0, bytes32 seed2, bytes32 k_mkl_root, uint64 count, uint256 price, uint256 expireAt, uint256 submitAt)
func (_PODEX *PODEXCallerSession) GetRecordBatch1(_a common.Address, _b common.Address, _sessionId *big.Int) (struct {
	Seed0    [32]byte
	Seed2    [32]byte
	KMklRoot [32]byte
	Count    uint64
	Price    *big.Int
	ExpireAt *big.Int
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordBatch1(&_PODEX.CallOpts, _a, _b, _sessionId)
}

// GetRecordBatch2 is a free data retrieval call binding the contract method 0x0ae429ea.
//
// Solidity: function getRecordBatch2(address _a, address _b, uint256 _sessionId) constant returns(bytes32 seed0, uint256 submitAt)
func (_PODEX *PODEXCaller) GetRecordBatch2(opts *bind.CallOpts, _a common.Address, _b common.Address, _sessionId *big.Int) (struct {
	Seed0    [32]byte
	SubmitAt *big.Int
}, error) {
	ret := new(struct {
		Seed0    [32]byte
		SubmitAt *big.Int
	})
	out := ret
	err := _PODEX.contract.Call(opts, out, "getRecordBatch2", _a, _b, _sessionId)
	return *ret, err
}

// GetRecordBatch2 is a free data retrieval call binding the contract method 0x0ae429ea.
//
// Solidity: function getRecordBatch2(address _a, address _b, uint256 _sessionId) constant returns(bytes32 seed0, uint256 submitAt)
func (_PODEX *PODEXSession) GetRecordBatch2(_a common.Address, _b common.Address, _sessionId *big.Int) (struct {
	Seed0    [32]byte
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordBatch2(&_PODEX.CallOpts, _a, _b, _sessionId)
}

// GetRecordBatch2 is a free data retrieval call binding the contract method 0x0ae429ea.
//
// Solidity: function getRecordBatch2(address _a, address _b, uint256 _sessionId) constant returns(bytes32 seed0, uint256 submitAt)
func (_PODEX *PODEXCallerSession) GetRecordBatch2(_a common.Address, _b common.Address, _sessionId *big.Int) (struct {
	Seed0    [32]byte
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordBatch2(&_PODEX.CallOpts, _a, _b, _sessionId)
}

// GetRecordBatch3 is a free data retrieval call binding the contract method 0x74ef2ae9.
//
// Solidity: function getRecordBatch3(address _a, address _b, uint256 _sessionId) constant returns(uint256 d, uint256 x0_lgs, uint256 submitAt)
func (_PODEX *PODEXCaller) GetRecordBatch3(opts *bind.CallOpts, _a common.Address, _b common.Address, _sessionId *big.Int) (struct {
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
	err := _PODEX.contract.Call(opts, out, "getRecordBatch3", _a, _b, _sessionId)
	return *ret, err
}

// GetRecordBatch3 is a free data retrieval call binding the contract method 0x74ef2ae9.
//
// Solidity: function getRecordBatch3(address _a, address _b, uint256 _sessionId) constant returns(uint256 d, uint256 x0_lgs, uint256 submitAt)
func (_PODEX *PODEXSession) GetRecordBatch3(_a common.Address, _b common.Address, _sessionId *big.Int) (struct {
	D        *big.Int
	X0Lgs    *big.Int
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordBatch3(&_PODEX.CallOpts, _a, _b, _sessionId)
}

// GetRecordBatch3 is a free data retrieval call binding the contract method 0x74ef2ae9.
//
// Solidity: function getRecordBatch3(address _a, address _b, uint256 _sessionId) constant returns(uint256 d, uint256 x0_lgs, uint256 submitAt)
func (_PODEX *PODEXCallerSession) GetRecordBatch3(_a common.Address, _b common.Address, _sessionId *big.Int) (struct {
	D        *big.Int
	X0Lgs    *big.Int
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordBatch3(&_PODEX.CallOpts, _a, _b, _sessionId)
}

// GetRecordVRF is a free data retrieval call binding the contract method 0xb79881e1.
//
// Solidity: function getRecordVRF(address _a, address _b, uint256 _sessionId) constant returns(uint256 r, uint256 submitAt)
func (_PODEX *PODEXCaller) GetRecordVRF(opts *bind.CallOpts, _a common.Address, _b common.Address, _sessionId *big.Int) (struct {
	R        *big.Int
	SubmitAt *big.Int
}, error) {
	ret := new(struct {
		R        *big.Int
		SubmitAt *big.Int
	})
	out := ret
	err := _PODEX.contract.Call(opts, out, "getRecordVRF", _a, _b, _sessionId)
	return *ret, err
}

// GetRecordVRF is a free data retrieval call binding the contract method 0xb79881e1.
//
// Solidity: function getRecordVRF(address _a, address _b, uint256 _sessionId) constant returns(uint256 r, uint256 submitAt)
func (_PODEX *PODEXSession) GetRecordVRF(_a common.Address, _b common.Address, _sessionId *big.Int) (struct {
	R        *big.Int
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordVRF(&_PODEX.CallOpts, _a, _b, _sessionId)
}

// GetRecordVRF is a free data retrieval call binding the contract method 0xb79881e1.
//
// Solidity: function getRecordVRF(address _a, address _b, uint256 _sessionId) constant returns(uint256 r, uint256 submitAt)
func (_PODEX *PODEXCallerSession) GetRecordVRF(_a common.Address, _b common.Address, _sessionId *big.Int) (struct {
	R        *big.Int
	SubmitAt *big.Int
}, error) {
	return _PODEX.Contract.GetRecordVRF(&_PODEX.CallOpts, _a, _b, _sessionId)
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

// T2 is a free data retrieval call binding the contract method 0xbaf2f868.
//
// Solidity: function t2() constant returns(uint256)
func (_PODEX *PODEXCaller) T2(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "t2")
	return *ret0, err
}

// T2 is a free data retrieval call binding the contract method 0xbaf2f868.
//
// Solidity: function t2() constant returns(uint256)
func (_PODEX *PODEXSession) T2() (*big.Int, error) {
	return _PODEX.Contract.T2(&_PODEX.CallOpts)
}

// T2 is a free data retrieval call binding the contract method 0xbaf2f868.
//
// Solidity: function t2() constant returns(uint256)
func (_PODEX *PODEXCallerSession) T2() (*big.Int, error) {
	return _PODEX.Contract.T2(&_PODEX.CallOpts)
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

// VerifyPath is a free data retrieval call binding the contract method 0x30ea0028.
//
// Solidity: function verifyPath(bytes32 _x, bytes32 _y, uint64 _ij, uint64 _ns, bytes32 _root, bytes32[] _mkl_path) constant returns(bool)
func (_PODEX *PODEXCaller) VerifyPath(opts *bind.CallOpts, _x [32]byte, _y [32]byte, _ij uint64, _ns uint64, _root [32]byte, _mkl_path [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "verifyPath", _x, _y, _ij, _ns, _root, _mkl_path)
	return *ret0, err
}

// VerifyPath is a free data retrieval call binding the contract method 0x30ea0028.
//
// Solidity: function verifyPath(bytes32 _x, bytes32 _y, uint64 _ij, uint64 _ns, bytes32 _root, bytes32[] _mkl_path) constant returns(bool)
func (_PODEX *PODEXSession) VerifyPath(_x [32]byte, _y [32]byte, _ij uint64, _ns uint64, _root [32]byte, _mkl_path [][32]byte) (bool, error) {
	return _PODEX.Contract.VerifyPath(&_PODEX.CallOpts, _x, _y, _ij, _ns, _root, _mkl_path)
}

// VerifyPath is a free data retrieval call binding the contract method 0x30ea0028.
//
// Solidity: function verifyPath(bytes32 _x, bytes32 _y, uint64 _ij, uint64 _ns, bytes32 _root, bytes32[] _mkl_path) constant returns(bool)
func (_PODEX *PODEXCallerSession) VerifyPath(_x [32]byte, _y [32]byte, _ij uint64, _ns uint64, _root [32]byte, _mkl_path [][32]byte) (bool, error) {
	return _PODEX.Contract.VerifyPath(&_PODEX.CallOpts, _x, _y, _ij, _ns, _root, _mkl_path)
}

// VerifyProofBatch2 is a free data retrieval call binding the contract method 0xdd3e54cb.
//
// Solidity: function verifyProofBatch2(uint64 count, uint64 s, bytes32 seed0, bytes32 seed2, uint256 sigma_vw) constant returns(bool)
func (_PODEX *PODEXCaller) VerifyProofBatch2(opts *bind.CallOpts, count uint64, s uint64, seed0 [32]byte, seed2 [32]byte, sigma_vw *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "verifyProofBatch2", count, s, seed0, seed2, sigma_vw)
	return *ret0, err
}

// VerifyProofBatch2 is a free data retrieval call binding the contract method 0xdd3e54cb.
//
// Solidity: function verifyProofBatch2(uint64 count, uint64 s, bytes32 seed0, bytes32 seed2, uint256 sigma_vw) constant returns(bool)
func (_PODEX *PODEXSession) VerifyProofBatch2(count uint64, s uint64, seed0 [32]byte, seed2 [32]byte, sigma_vw *big.Int) (bool, error) {
	return _PODEX.Contract.VerifyProofBatch2(&_PODEX.CallOpts, count, s, seed0, seed2, sigma_vw)
}

// VerifyProofBatch2 is a free data retrieval call binding the contract method 0xdd3e54cb.
//
// Solidity: function verifyProofBatch2(uint64 count, uint64 s, bytes32 seed0, bytes32 seed2, uint256 sigma_vw) constant returns(bool)
func (_PODEX *PODEXCallerSession) VerifyProofBatch2(count uint64, s uint64, seed0 [32]byte, seed2 [32]byte, sigma_vw *big.Int) (bool, error) {
	return _PODEX.Contract.VerifyProofBatch2(&_PODEX.CallOpts, count, s, seed0, seed2, sigma_vw)
}

// VerifyProofBatch3 is a free data retrieval call binding the contract method 0x1778d3c6.
//
// Solidity: function verifyProofBatch3(uint256[2] _r_u0d, uint256[2] _r_u0_x0_lgs, uint256 _s_d, uint256 _s_x0_lgs) constant returns(bool)
func (_PODEX *PODEXCaller) VerifyProofBatch3(opts *bind.CallOpts, _r_u0d [2]*big.Int, _r_u0_x0_lgs [2]*big.Int, _s_d *big.Int, _s_x0_lgs *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "verifyProofBatch3", _r_u0d, _r_u0_x0_lgs, _s_d, _s_x0_lgs)
	return *ret0, err
}

// VerifyProofBatch3 is a free data retrieval call binding the contract method 0x1778d3c6.
//
// Solidity: function verifyProofBatch3(uint256[2] _r_u0d, uint256[2] _r_u0_x0_lgs, uint256 _s_d, uint256 _s_x0_lgs) constant returns(bool)
func (_PODEX *PODEXSession) VerifyProofBatch3(_r_u0d [2]*big.Int, _r_u0_x0_lgs [2]*big.Int, _s_d *big.Int, _s_x0_lgs *big.Int) (bool, error) {
	return _PODEX.Contract.VerifyProofBatch3(&_PODEX.CallOpts, _r_u0d, _r_u0_x0_lgs, _s_d, _s_x0_lgs)
}

// VerifyProofBatch3 is a free data retrieval call binding the contract method 0x1778d3c6.
//
// Solidity: function verifyProofBatch3(uint256[2] _r_u0d, uint256[2] _r_u0_x0_lgs, uint256 _s_d, uint256 _s_x0_lgs) constant returns(bool)
func (_PODEX *PODEXCallerSession) VerifyProofBatch3(_r_u0d [2]*big.Int, _r_u0_x0_lgs [2]*big.Int, _s_d *big.Int, _s_x0_lgs *big.Int) (bool, error) {
	return _PODEX.Contract.VerifyProofBatch3(&_PODEX.CallOpts, _r_u0d, _r_u0_x0_lgs, _s_d, _s_x0_lgs)
}

// VerifyProofVRF is a free data retrieval call binding the contract method 0x4352ab9a.
//
// Solidity: function verifyProofVRF(uint256[2] _g_exp_r, uint256 _s_r) constant returns(bool)
func (_PODEX *PODEXCaller) VerifyProofVRF(opts *bind.CallOpts, _g_exp_r [2]*big.Int, _s_r *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PODEX.contract.Call(opts, out, "verifyProofVRF", _g_exp_r, _s_r)
	return *ret0, err
}

// VerifyProofVRF is a free data retrieval call binding the contract method 0x4352ab9a.
//
// Solidity: function verifyProofVRF(uint256[2] _g_exp_r, uint256 _s_r) constant returns(bool)
func (_PODEX *PODEXSession) VerifyProofVRF(_g_exp_r [2]*big.Int, _s_r *big.Int) (bool, error) {
	return _PODEX.Contract.VerifyProofVRF(&_PODEX.CallOpts, _g_exp_r, _s_r)
}

// VerifyProofVRF is a free data retrieval call binding the contract method 0x4352ab9a.
//
// Solidity: function verifyProofVRF(uint256[2] _g_exp_r, uint256 _s_r) constant returns(bool)
func (_PODEX *PODEXCallerSession) VerifyProofVRF(_g_exp_r [2]*big.Int, _s_r *big.Int) (bool, error) {
	return _PODEX.Contract.VerifyProofVRF(&_PODEX.CallOpts, _g_exp_r, _s_r)
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
// Solidity: function claimBatch1(address _a, uint256 _sessionId, uint64 _i, uint64 _j, uint256 _tx, uint256 _ty, bytes32[] _mkl_path, uint64 _sCnt) returns()
func (_PODEX *PODEXTransactor) ClaimBatch1(opts *bind.TransactOpts, _a common.Address, _sessionId *big.Int, _i uint64, _j uint64, _tx *big.Int, _ty *big.Int, _mkl_path [][32]byte, _sCnt uint64) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "claimBatch1", _a, _sessionId, _i, _j, _tx, _ty, _mkl_path, _sCnt)
}

// ClaimBatch1 is a paid mutator transaction binding the contract method 0x0846fa02.
//
// Solidity: function claimBatch1(address _a, uint256 _sessionId, uint64 _i, uint64 _j, uint256 _tx, uint256 _ty, bytes32[] _mkl_path, uint64 _sCnt) returns()
func (_PODEX *PODEXSession) ClaimBatch1(_a common.Address, _sessionId *big.Int, _i uint64, _j uint64, _tx *big.Int, _ty *big.Int, _mkl_path [][32]byte, _sCnt uint64) (*types.Transaction, error) {
	return _PODEX.Contract.ClaimBatch1(&_PODEX.TransactOpts, _a, _sessionId, _i, _j, _tx, _ty, _mkl_path, _sCnt)
}

// ClaimBatch1 is a paid mutator transaction binding the contract method 0x0846fa02.
//
// Solidity: function claimBatch1(address _a, uint256 _sessionId, uint64 _i, uint64 _j, uint256 _tx, uint256 _ty, bytes32[] _mkl_path, uint64 _sCnt) returns()
func (_PODEX *PODEXTransactorSession) ClaimBatch1(_a common.Address, _sessionId *big.Int, _i uint64, _j uint64, _tx *big.Int, _ty *big.Int, _mkl_path [][32]byte, _sCnt uint64) (*types.Transaction, error) {
	return _PODEX.Contract.ClaimBatch1(&_PODEX.TransactOpts, _a, _sessionId, _i, _j, _tx, _ty, _mkl_path, _sCnt)
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

// SubmitProofBatch1 is a paid mutator transaction binding the contract method 0x9d33e64a.
//
// Solidity: function submitProofBatch1(bytes32 _seed0, uint256 _sessionId, address _b, bytes32 _seed2, bytes32 _k_mkl_root, uint64 _count, uint256 _price, uint256 _expireAt, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_PODEX *PODEXTransactor) SubmitProofBatch1(opts *bind.TransactOpts, _seed0 [32]byte, _sessionId *big.Int, _b common.Address, _seed2 [32]byte, _k_mkl_root [32]byte, _count uint64, _price *big.Int, _expireAt *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "submitProofBatch1", _seed0, _sessionId, _b, _seed2, _k_mkl_root, _count, _price, _expireAt, _v, _r, _s)
}

// SubmitProofBatch1 is a paid mutator transaction binding the contract method 0x9d33e64a.
//
// Solidity: function submitProofBatch1(bytes32 _seed0, uint256 _sessionId, address _b, bytes32 _seed2, bytes32 _k_mkl_root, uint64 _count, uint256 _price, uint256 _expireAt, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_PODEX *PODEXSession) SubmitProofBatch1(_seed0 [32]byte, _sessionId *big.Int, _b common.Address, _seed2 [32]byte, _k_mkl_root [32]byte, _count uint64, _price *big.Int, _expireAt *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofBatch1(&_PODEX.TransactOpts, _seed0, _sessionId, _b, _seed2, _k_mkl_root, _count, _price, _expireAt, _v, _r, _s)
}

// SubmitProofBatch1 is a paid mutator transaction binding the contract method 0x9d33e64a.
//
// Solidity: function submitProofBatch1(bytes32 _seed0, uint256 _sessionId, address _b, bytes32 _seed2, bytes32 _k_mkl_root, uint64 _count, uint256 _price, uint256 _expireAt, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_PODEX *PODEXTransactorSession) SubmitProofBatch1(_seed0 [32]byte, _sessionId *big.Int, _b common.Address, _seed2 [32]byte, _k_mkl_root [32]byte, _count uint64, _price *big.Int, _expireAt *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofBatch1(&_PODEX.TransactOpts, _seed0, _sessionId, _b, _seed2, _k_mkl_root, _count, _price, _expireAt, _v, _r, _s)
}

// SubmitProofBatch2 is a paid mutator transaction binding the contract method 0x56182831.
//
// Solidity: function submitProofBatch2(bytes32 _seed0, uint64 _sCnt, uint256 _sessionId, address _b, bytes32 _seed2, uint256 _sigma_vw, uint64 _count, uint256 _price, uint256 _expireAt, uint8 _v, bytes32[2] _rs) returns()
func (_PODEX *PODEXTransactor) SubmitProofBatch2(opts *bind.TransactOpts, _seed0 [32]byte, _sCnt uint64, _sessionId *big.Int, _b common.Address, _seed2 [32]byte, _sigma_vw *big.Int, _count uint64, _price *big.Int, _expireAt *big.Int, _v uint8, _rs [2][32]byte) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "submitProofBatch2", _seed0, _sCnt, _sessionId, _b, _seed2, _sigma_vw, _count, _price, _expireAt, _v, _rs)
}

// SubmitProofBatch2 is a paid mutator transaction binding the contract method 0x56182831.
//
// Solidity: function submitProofBatch2(bytes32 _seed0, uint64 _sCnt, uint256 _sessionId, address _b, bytes32 _seed2, uint256 _sigma_vw, uint64 _count, uint256 _price, uint256 _expireAt, uint8 _v, bytes32[2] _rs) returns()
func (_PODEX *PODEXSession) SubmitProofBatch2(_seed0 [32]byte, _sCnt uint64, _sessionId *big.Int, _b common.Address, _seed2 [32]byte, _sigma_vw *big.Int, _count uint64, _price *big.Int, _expireAt *big.Int, _v uint8, _rs [2][32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofBatch2(&_PODEX.TransactOpts, _seed0, _sCnt, _sessionId, _b, _seed2, _sigma_vw, _count, _price, _expireAt, _v, _rs)
}

// SubmitProofBatch2 is a paid mutator transaction binding the contract method 0x56182831.
//
// Solidity: function submitProofBatch2(bytes32 _seed0, uint64 _sCnt, uint256 _sessionId, address _b, bytes32 _seed2, uint256 _sigma_vw, uint64 _count, uint256 _price, uint256 _expireAt, uint8 _v, bytes32[2] _rs) returns()
func (_PODEX *PODEXTransactorSession) SubmitProofBatch2(_seed0 [32]byte, _sCnt uint64, _sessionId *big.Int, _b common.Address, _seed2 [32]byte, _sigma_vw *big.Int, _count uint64, _price *big.Int, _expireAt *big.Int, _v uint8, _rs [2][32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofBatch2(&_PODEX.TransactOpts, _seed0, _sCnt, _sessionId, _b, _seed2, _sigma_vw, _count, _price, _expireAt, _v, _rs)
}

// SubmitProofBatch3 is a paid mutator transaction binding the contract method 0xbf8c5d28.
//
// Solidity: function submitProofBatch3(uint256 _s_d, uint256 _s_x0_lgs, uint256 _sessionId, address _b, uint256[2] _r_u0_x0_lgs, uint256[2] _r_u0d, uint256 _price, uint256 _expireAt, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_PODEX *PODEXTransactor) SubmitProofBatch3(opts *bind.TransactOpts, _s_d *big.Int, _s_x0_lgs *big.Int, _sessionId *big.Int, _b common.Address, _r_u0_x0_lgs [2]*big.Int, _r_u0d [2]*big.Int, _price *big.Int, _expireAt *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "submitProofBatch3", _s_d, _s_x0_lgs, _sessionId, _b, _r_u0_x0_lgs, _r_u0d, _price, _expireAt, _v, _r, _s)
}

// SubmitProofBatch3 is a paid mutator transaction binding the contract method 0xbf8c5d28.
//
// Solidity: function submitProofBatch3(uint256 _s_d, uint256 _s_x0_lgs, uint256 _sessionId, address _b, uint256[2] _r_u0_x0_lgs, uint256[2] _r_u0d, uint256 _price, uint256 _expireAt, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_PODEX *PODEXSession) SubmitProofBatch3(_s_d *big.Int, _s_x0_lgs *big.Int, _sessionId *big.Int, _b common.Address, _r_u0_x0_lgs [2]*big.Int, _r_u0d [2]*big.Int, _price *big.Int, _expireAt *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofBatch3(&_PODEX.TransactOpts, _s_d, _s_x0_lgs, _sessionId, _b, _r_u0_x0_lgs, _r_u0d, _price, _expireAt, _v, _r, _s)
}

// SubmitProofBatch3 is a paid mutator transaction binding the contract method 0xbf8c5d28.
//
// Solidity: function submitProofBatch3(uint256 _s_d, uint256 _s_x0_lgs, uint256 _sessionId, address _b, uint256[2] _r_u0_x0_lgs, uint256[2] _r_u0d, uint256 _price, uint256 _expireAt, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_PODEX *PODEXTransactorSession) SubmitProofBatch3(_s_d *big.Int, _s_x0_lgs *big.Int, _sessionId *big.Int, _b common.Address, _r_u0_x0_lgs [2]*big.Int, _r_u0d [2]*big.Int, _price *big.Int, _expireAt *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofBatch3(&_PODEX.TransactOpts, _s_d, _s_x0_lgs, _sessionId, _b, _r_u0_x0_lgs, _r_u0d, _price, _expireAt, _v, _r, _s)
}

// SubmitProofVRF is a paid mutator transaction binding the contract method 0x564d0823.
//
// Solidity: function submitProofVRF(uint256 _s_r, uint256 _sessionId, address _b, uint256[2] _g_exp_r, uint256 _price, uint256 _expireAt, uint8 _v, bytes32[2] _rs) returns()
func (_PODEX *PODEXTransactor) SubmitProofVRF(opts *bind.TransactOpts, _s_r *big.Int, _sessionId *big.Int, _b common.Address, _g_exp_r [2]*big.Int, _price *big.Int, _expireAt *big.Int, _v uint8, _rs [2][32]byte) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "submitProofVRF", _s_r, _sessionId, _b, _g_exp_r, _price, _expireAt, _v, _rs)
}

// SubmitProofVRF is a paid mutator transaction binding the contract method 0x564d0823.
//
// Solidity: function submitProofVRF(uint256 _s_r, uint256 _sessionId, address _b, uint256[2] _g_exp_r, uint256 _price, uint256 _expireAt, uint8 _v, bytes32[2] _rs) returns()
func (_PODEX *PODEXSession) SubmitProofVRF(_s_r *big.Int, _sessionId *big.Int, _b common.Address, _g_exp_r [2]*big.Int, _price *big.Int, _expireAt *big.Int, _v uint8, _rs [2][32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofVRF(&_PODEX.TransactOpts, _s_r, _sessionId, _b, _g_exp_r, _price, _expireAt, _v, _rs)
}

// SubmitProofVRF is a paid mutator transaction binding the contract method 0x564d0823.
//
// Solidity: function submitProofVRF(uint256 _s_r, uint256 _sessionId, address _b, uint256[2] _g_exp_r, uint256 _price, uint256 _expireAt, uint8 _v, bytes32[2] _rs) returns()
func (_PODEX *PODEXTransactorSession) SubmitProofVRF(_s_r *big.Int, _sessionId *big.Int, _b common.Address, _g_exp_r [2]*big.Int, _price *big.Int, _expireAt *big.Int, _v uint8, _rs [2][32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.SubmitProofVRF(&_PODEX.TransactOpts, _s_r, _sessionId, _b, _g_exp_r, _price, _expireAt, _v, _rs)
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

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 _bltKey) returns()
func (_PODEX *PODEXTransactor) Withdraw(opts *bind.TransactOpts, _bltKey [32]byte) (*types.Transaction, error) {
	return _PODEX.contract.Transact(opts, "withdraw", _bltKey)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 _bltKey) returns()
func (_PODEX *PODEXSession) Withdraw(_bltKey [32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.Withdraw(&_PODEX.TransactOpts, _bltKey)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 _bltKey) returns()
func (_PODEX *PODEXTransactorSession) Withdraw(_bltKey [32]byte) (*types.Transaction, error) {
	return _PODEX.Contract.Withdraw(&_PODEX.TransactOpts, _bltKey)
}

// PODEXLogG1PointIterator is returned from FilterLogG1Point and is used to iterate over the raw logs and unpacked data for LogG1Point events raised by the PODEX contract.
type PODEXLogG1PointIterator struct {
	Event *PODEXLogG1Point // Event containing the contract specifics and raw log

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
func (it *PODEXLogG1PointIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PODEXLogG1Point)
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
		it.Event = new(PODEXLogG1Point)
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
func (it *PODEXLogG1PointIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PODEXLogG1PointIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PODEXLogG1Point represents a LogG1Point event raised by the PODEX contract.
type PODEXLogG1Point struct {
	X   *big.Int
	Y   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogG1Point is a free log retrieval operation binding the contract event 0xf31d566be682591378c2ef5bb05942bede97b136de9ca8539e20afb34daaf425.
//
// Solidity: event LogG1Point(uint256 _x, uint256 _y)
func (_PODEX *PODEXFilterer) FilterLogG1Point(opts *bind.FilterOpts) (*PODEXLogG1PointIterator, error) {

	logs, sub, err := _PODEX.contract.FilterLogs(opts, "LogG1Point")
	if err != nil {
		return nil, err
	}
	return &PODEXLogG1PointIterator{contract: _PODEX.contract, event: "LogG1Point", logs: logs, sub: sub}, nil
}

// WatchLogG1Point is a free log subscription operation binding the contract event 0xf31d566be682591378c2ef5bb05942bede97b136de9ca8539e20afb34daaf425.
//
// Solidity: event LogG1Point(uint256 _x, uint256 _y)
func (_PODEX *PODEXFilterer) WatchLogG1Point(opts *bind.WatchOpts, sink chan<- *PODEXLogG1Point) (event.Subscription, error) {

	logs, sub, err := _PODEX.contract.WatchLogs(opts, "LogG1Point")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PODEXLogG1Point)
				if err := _PODEX.contract.UnpackLog(event, "LogG1Point", log); err != nil {
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
