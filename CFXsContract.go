// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// CFXsContractOutputCFXsData is an auto generated low-level Go binding around an user-defined struct.
type CFXsContractOutputCFXsData struct {
	Owner  common.Address
	Amount *big.Int
	Data   string
}

// CFXsContractTransaction is an auto generated low-level Go binding around an user-defined struct.
type CFXsContractTransaction struct {
	Inputs  []*big.Int
	Outputs []CFXsContractOutputCFXsData
}

// CFXsContractMetaData contains all meta data concerning the CFXsContract contract.
var CFXsContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"CFXsCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CFXsDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"CFXsEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"CFXsId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"etherAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"locktime\",\"type\":\"uint256\"}],\"name\":\"CFXsLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"CFXsId\",\"type\":\"uint256\"}],\"name\":\"CFXsUnlocked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CFXsCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CFXss\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CreateCFXs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"CFXsId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DangerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"LockedCFXs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_ether\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"locktime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"CFXsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ether\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"locktime\",\"type\":\"uint256\"}],\"name\":\"LockingScript\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"CFXsId\",\"type\":\"uint256\"}],\"name\":\"OwnerUnlockingScript\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"CFXsId\",\"type\":\"uint256\"}],\"name\":\"UnlockingScript\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getLockStates\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"CFXsId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"inscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"inputs\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"internalType\":\"structCFXsContract.OutputCFXsData[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"}],\"internalType\":\"structCFXsContract.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"processTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// CFXsContractABI is the input ABI used to generate the binding from.
// Deprecated: Use CFXsContractMetaData.ABI instead.
var CFXsContractABI = CFXsContractMetaData.ABI

// CFXsContract is an auto generated Go binding around an Ethereum contract.
type CFXsContract struct {
	CFXsContractCaller     // Read-only binding to the contract
	CFXsContractTransactor // Write-only binding to the contract
	CFXsContractFilterer   // Log filterer for contract events
}

// CFXsContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CFXsContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CFXsContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CFXsContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CFXsContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CFXsContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CFXsContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CFXsContractSession struct {
	Contract     *CFXsContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CFXsContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CFXsContractCallerSession struct {
	Contract *CFXsContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CFXsContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CFXsContractTransactorSession struct {
	Contract     *CFXsContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CFXsContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CFXsContractRaw struct {
	Contract *CFXsContract // Generic contract binding to access the raw methods on
}

// CFXsContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CFXsContractCallerRaw struct {
	Contract *CFXsContractCaller // Generic read-only contract binding to access the raw methods on
}

// CFXsContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CFXsContractTransactorRaw struct {
	Contract *CFXsContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCFXsContract creates a new instance of CFXsContract, bound to a specific deployed contract.
func NewCFXsContract(address common.Address, backend bind.ContractBackend) (*CFXsContract, error) {
	contract, err := bindCFXsContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CFXsContract{CFXsContractCaller: CFXsContractCaller{contract: contract}, CFXsContractTransactor: CFXsContractTransactor{contract: contract}, CFXsContractFilterer: CFXsContractFilterer{contract: contract}}, nil
}

// NewCFXsContractCaller creates a new read-only instance of CFXsContract, bound to a specific deployed contract.
func NewCFXsContractCaller(address common.Address, caller bind.ContractCaller) (*CFXsContractCaller, error) {
	contract, err := bindCFXsContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CFXsContractCaller{contract: contract}, nil
}

// NewCFXsContractTransactor creates a new write-only instance of CFXsContract, bound to a specific deployed contract.
func NewCFXsContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CFXsContractTransactor, error) {
	contract, err := bindCFXsContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CFXsContractTransactor{contract: contract}, nil
}

// NewCFXsContractFilterer creates a new log filterer instance of CFXsContract, bound to a specific deployed contract.
func NewCFXsContractFilterer(address common.Address, filterer bind.ContractFilterer) (*CFXsContractFilterer, error) {
	contract, err := bindCFXsContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CFXsContractFilterer{contract: contract}, nil
}

// bindCFXsContract binds a generic wrapper to an already deployed contract.
func bindCFXsContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CFXsContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CFXsContract *CFXsContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CFXsContract.Contract.CFXsContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CFXsContract *CFXsContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CFXsContract.Contract.CFXsContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CFXsContract *CFXsContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CFXsContract.Contract.CFXsContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CFXsContract *CFXsContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CFXsContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CFXsContract *CFXsContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CFXsContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CFXsContract *CFXsContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CFXsContract.Contract.contract.Transact(opts, method, params...)
}

// CFXsCounter is a free data retrieval call binding the contract method 0xb5c81b42.
//
// Solidity: function CFXsCounter() view returns(uint256)
func (_CFXsContract *CFXsContractCaller) CFXsCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CFXsContract.contract.Call(opts, &out, "CFXsCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CFXsCounter is a free data retrieval call binding the contract method 0xb5c81b42.
//
// Solidity: function CFXsCounter() view returns(uint256)
func (_CFXsContract *CFXsContractSession) CFXsCounter() (*big.Int, error) {
	return _CFXsContract.Contract.CFXsCounter(&_CFXsContract.CallOpts)
}

// CFXsCounter is a free data retrieval call binding the contract method 0xb5c81b42.
//
// Solidity: function CFXsCounter() view returns(uint256)
func (_CFXsContract *CFXsContractCallerSession) CFXsCounter() (*big.Int, error) {
	return _CFXsContract.Contract.CFXsCounter(&_CFXsContract.CallOpts)
}

// CFXss is a free data retrieval call binding the contract method 0xd06d13ba.
//
// Solidity: function CFXss(uint256 ) view returns(uint256 id, address owner, uint256 amount, string data)
func (_CFXsContract *CFXsContractCaller) CFXss(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id     *big.Int
	Owner  common.Address
	Amount *big.Int
	Data   string
}, error) {
	var out []interface{}
	err := _CFXsContract.contract.Call(opts, &out, "CFXss", arg0)

	outstruct := new(struct {
		Id     *big.Int
		Owner  common.Address
		Amount *big.Int
		Data   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// CFXss is a free data retrieval call binding the contract method 0xd06d13ba.
//
// Solidity: function CFXss(uint256 ) view returns(uint256 id, address owner, uint256 amount, string data)
func (_CFXsContract *CFXsContractSession) CFXss(arg0 *big.Int) (struct {
	Id     *big.Int
	Owner  common.Address
	Amount *big.Int
	Data   string
}, error) {
	return _CFXsContract.Contract.CFXss(&_CFXsContract.CallOpts, arg0)
}

// CFXss is a free data retrieval call binding the contract method 0xd06d13ba.
//
// Solidity: function CFXss(uint256 ) view returns(uint256 id, address owner, uint256 amount, string data)
func (_CFXsContract *CFXsContractCallerSession) CFXss(arg0 *big.Int) (struct {
	Id     *big.Int
	Owner  common.Address
	Amount *big.Int
	Data   string
}, error) {
	return _CFXsContract.Contract.CFXss(&_CFXsContract.CallOpts, arg0)
}

// LockedCFXs is a free data retrieval call binding the contract method 0x27139fe5.
//
// Solidity: function LockedCFXs(uint256 ) view returns(uint256 _ether, uint256 locktime)
func (_CFXsContract *CFXsContractCaller) LockedCFXs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Ether    *big.Int
	Locktime *big.Int
}, error) {
	var out []interface{}
	err := _CFXsContract.contract.Call(opts, &out, "LockedCFXs", arg0)

	outstruct := new(struct {
		Ether    *big.Int
		Locktime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ether = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Locktime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LockedCFXs is a free data retrieval call binding the contract method 0x27139fe5.
//
// Solidity: function LockedCFXs(uint256 ) view returns(uint256 _ether, uint256 locktime)
func (_CFXsContract *CFXsContractSession) LockedCFXs(arg0 *big.Int) (struct {
	Ether    *big.Int
	Locktime *big.Int
}, error) {
	return _CFXsContract.Contract.LockedCFXs(&_CFXsContract.CallOpts, arg0)
}

// LockedCFXs is a free data retrieval call binding the contract method 0x27139fe5.
//
// Solidity: function LockedCFXs(uint256 ) view returns(uint256 _ether, uint256 locktime)
func (_CFXsContract *CFXsContractCallerSession) LockedCFXs(arg0 *big.Int) (struct {
	Ether    *big.Int
	Locktime *big.Int
}, error) {
	return _CFXsContract.Contract.LockedCFXs(&_CFXsContract.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_CFXsContract *CFXsContractCaller) BalanceOf(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CFXsContract.contract.Call(opts, &out, "balanceOf", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_CFXsContract *CFXsContractSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _CFXsContract.Contract.BalanceOf(&_CFXsContract.CallOpts, _addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_CFXsContract *CFXsContractCallerSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _CFXsContract.Contract.BalanceOf(&_CFXsContract.CallOpts, _addr)
}

// GetLockStates is a free data retrieval call binding the contract method 0x49c1742f.
//
// Solidity: function getLockStates(uint256 _id) view returns(bool)
func (_CFXsContract *CFXsContractCaller) GetLockStates(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _CFXsContract.contract.Call(opts, &out, "getLockStates", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetLockStates is a free data retrieval call binding the contract method 0x49c1742f.
//
// Solidity: function getLockStates(uint256 _id) view returns(bool)
func (_CFXsContract *CFXsContractSession) GetLockStates(_id *big.Int) (bool, error) {
	return _CFXsContract.Contract.GetLockStates(&_CFXsContract.CallOpts, _id)
}

// GetLockStates is a free data retrieval call binding the contract method 0x49c1742f.
//
// Solidity: function getLockStates(uint256 _id) view returns(bool)
func (_CFXsContract *CFXsContractCallerSession) GetLockStates(_id *big.Int) (bool, error) {
	return _CFXsContract.Contract.GetLockStates(&_CFXsContract.CallOpts, _id)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CFXsContract *CFXsContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CFXsContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CFXsContract *CFXsContractSession) TotalSupply() (*big.Int, error) {
	return _CFXsContract.Contract.TotalSupply(&_CFXsContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CFXsContract *CFXsContractCallerSession) TotalSupply() (*big.Int, error) {
	return _CFXsContract.Contract.TotalSupply(&_CFXsContract.CallOpts)
}

// CreateCFXs is a paid mutator transaction binding the contract method 0x7d80e2ca.
//
// Solidity: function CreateCFXs() returns()
func (_CFXsContract *CFXsContractTransactor) CreateCFXs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CFXsContract.contract.Transact(opts, "CreateCFXs")
}

// CreateCFXs is a paid mutator transaction binding the contract method 0x7d80e2ca.
//
// Solidity: function CreateCFXs() returns()
func (_CFXsContract *CFXsContractSession) CreateCFXs() (*types.Transaction, error) {
	return _CFXsContract.Contract.CreateCFXs(&_CFXsContract.TransactOpts)
}

// CreateCFXs is a paid mutator transaction binding the contract method 0x7d80e2ca.
//
// Solidity: function CreateCFXs() returns()
func (_CFXsContract *CFXsContractTransactorSession) CreateCFXs() (*types.Transaction, error) {
	return _CFXsContract.Contract.CreateCFXs(&_CFXsContract.TransactOpts)
}

// DangerTransfer is a paid mutator transaction binding the contract method 0xc365a4cc.
//
// Solidity: function DangerTransfer(uint256 CFXsId, address _to, uint256 _amount) returns()
func (_CFXsContract *CFXsContractTransactor) DangerTransfer(opts *bind.TransactOpts, CFXsId *big.Int, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CFXsContract.contract.Transact(opts, "DangerTransfer", CFXsId, _to, _amount)
}

// DangerTransfer is a paid mutator transaction binding the contract method 0xc365a4cc.
//
// Solidity: function DangerTransfer(uint256 CFXsId, address _to, uint256 _amount) returns()
func (_CFXsContract *CFXsContractSession) DangerTransfer(CFXsId *big.Int, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CFXsContract.Contract.DangerTransfer(&_CFXsContract.TransactOpts, CFXsId, _to, _amount)
}

// DangerTransfer is a paid mutator transaction binding the contract method 0xc365a4cc.
//
// Solidity: function DangerTransfer(uint256 CFXsId, address _to, uint256 _amount) returns()
func (_CFXsContract *CFXsContractTransactorSession) DangerTransfer(CFXsId *big.Int, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _CFXsContract.Contract.DangerTransfer(&_CFXsContract.TransactOpts, CFXsId, _to, _amount)
}

// LockingScript is a paid mutator transaction binding the contract method 0x6068ef1e.
//
// Solidity: function LockingScript(uint256 CFXsId, uint256 _ether, uint256 locktime) returns()
func (_CFXsContract *CFXsContractTransactor) LockingScript(opts *bind.TransactOpts, CFXsId *big.Int, _ether *big.Int, locktime *big.Int) (*types.Transaction, error) {
	return _CFXsContract.contract.Transact(opts, "LockingScript", CFXsId, _ether, locktime)
}

// LockingScript is a paid mutator transaction binding the contract method 0x6068ef1e.
//
// Solidity: function LockingScript(uint256 CFXsId, uint256 _ether, uint256 locktime) returns()
func (_CFXsContract *CFXsContractSession) LockingScript(CFXsId *big.Int, _ether *big.Int, locktime *big.Int) (*types.Transaction, error) {
	return _CFXsContract.Contract.LockingScript(&_CFXsContract.TransactOpts, CFXsId, _ether, locktime)
}

// LockingScript is a paid mutator transaction binding the contract method 0x6068ef1e.
//
// Solidity: function LockingScript(uint256 CFXsId, uint256 _ether, uint256 locktime) returns()
func (_CFXsContract *CFXsContractTransactorSession) LockingScript(CFXsId *big.Int, _ether *big.Int, locktime *big.Int) (*types.Transaction, error) {
	return _CFXsContract.Contract.LockingScript(&_CFXsContract.TransactOpts, CFXsId, _ether, locktime)
}

// OwnerUnlockingScript is a paid mutator transaction binding the contract method 0x61e56c2c.
//
// Solidity: function OwnerUnlockingScript(uint256 CFXsId) returns()
func (_CFXsContract *CFXsContractTransactor) OwnerUnlockingScript(opts *bind.TransactOpts, CFXsId *big.Int) (*types.Transaction, error) {
	return _CFXsContract.contract.Transact(opts, "OwnerUnlockingScript", CFXsId)
}

// OwnerUnlockingScript is a paid mutator transaction binding the contract method 0x61e56c2c.
//
// Solidity: function OwnerUnlockingScript(uint256 CFXsId) returns()
func (_CFXsContract *CFXsContractSession) OwnerUnlockingScript(CFXsId *big.Int) (*types.Transaction, error) {
	return _CFXsContract.Contract.OwnerUnlockingScript(&_CFXsContract.TransactOpts, CFXsId)
}

// OwnerUnlockingScript is a paid mutator transaction binding the contract method 0x61e56c2c.
//
// Solidity: function OwnerUnlockingScript(uint256 CFXsId) returns()
func (_CFXsContract *CFXsContractTransactorSession) OwnerUnlockingScript(CFXsId *big.Int) (*types.Transaction, error) {
	return _CFXsContract.Contract.OwnerUnlockingScript(&_CFXsContract.TransactOpts, CFXsId)
}

// UnlockingScript is a paid mutator transaction binding the contract method 0xb27ad94e.
//
// Solidity: function UnlockingScript(uint256 CFXsId) payable returns()
func (_CFXsContract *CFXsContractTransactor) UnlockingScript(opts *bind.TransactOpts, CFXsId *big.Int) (*types.Transaction, error) {
	return _CFXsContract.contract.Transact(opts, "UnlockingScript", CFXsId)
}

// UnlockingScript is a paid mutator transaction binding the contract method 0xb27ad94e.
//
// Solidity: function UnlockingScript(uint256 CFXsId) payable returns()
func (_CFXsContract *CFXsContractSession) UnlockingScript(CFXsId *big.Int) (*types.Transaction, error) {
	return _CFXsContract.Contract.UnlockingScript(&_CFXsContract.TransactOpts, CFXsId)
}

// UnlockingScript is a paid mutator transaction binding the contract method 0xb27ad94e.
//
// Solidity: function UnlockingScript(uint256 CFXsId) payable returns()
func (_CFXsContract *CFXsContractTransactorSession) UnlockingScript(CFXsId *big.Int) (*types.Transaction, error) {
	return _CFXsContract.Contract.UnlockingScript(&_CFXsContract.TransactOpts, CFXsId)
}

// Inscribe is a paid mutator transaction binding the contract method 0xed619c20.
//
// Solidity: function inscribe(uint256 CFXsId, string _data) returns()
func (_CFXsContract *CFXsContractTransactor) Inscribe(opts *bind.TransactOpts, CFXsId *big.Int, _data string) (*types.Transaction, error) {
	return _CFXsContract.contract.Transact(opts, "inscribe", CFXsId, _data)
}

// Inscribe is a paid mutator transaction binding the contract method 0xed619c20.
//
// Solidity: function inscribe(uint256 CFXsId, string _data) returns()
func (_CFXsContract *CFXsContractSession) Inscribe(CFXsId *big.Int, _data string) (*types.Transaction, error) {
	return _CFXsContract.Contract.Inscribe(&_CFXsContract.TransactOpts, CFXsId, _data)
}

// Inscribe is a paid mutator transaction binding the contract method 0xed619c20.
//
// Solidity: function inscribe(uint256 CFXsId, string _data) returns()
func (_CFXsContract *CFXsContractTransactorSession) Inscribe(CFXsId *big.Int, _data string) (*types.Transaction, error) {
	return _CFXsContract.Contract.Inscribe(&_CFXsContract.TransactOpts, CFXsId, _data)
}

// ProcessTransaction is a paid mutator transaction binding the contract method 0x5054ff53.
//
// Solidity: function processTransaction((uint256[],(address,uint256,string)[]) _tx) returns()
func (_CFXsContract *CFXsContractTransactor) ProcessTransaction(opts *bind.TransactOpts, _tx CFXsContractTransaction) (*types.Transaction, error) {
	return _CFXsContract.contract.Transact(opts, "processTransaction", _tx)
}

// ProcessTransaction is a paid mutator transaction binding the contract method 0x5054ff53.
//
// Solidity: function processTransaction((uint256[],(address,uint256,string)[]) _tx) returns()
func (_CFXsContract *CFXsContractSession) ProcessTransaction(_tx CFXsContractTransaction) (*types.Transaction, error) {
	return _CFXsContract.Contract.ProcessTransaction(&_CFXsContract.TransactOpts, _tx)
}

// ProcessTransaction is a paid mutator transaction binding the contract method 0x5054ff53.
//
// Solidity: function processTransaction((uint256[],(address,uint256,string)[]) _tx) returns()
func (_CFXsContract *CFXsContractTransactorSession) ProcessTransaction(_tx CFXsContractTransaction) (*types.Transaction, error) {
	return _CFXsContract.Contract.ProcessTransaction(&_CFXsContract.TransactOpts, _tx)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CFXsContract *CFXsContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CFXsContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CFXsContract *CFXsContractSession) Receive() (*types.Transaction, error) {
	return _CFXsContract.Contract.Receive(&_CFXsContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CFXsContract *CFXsContractTransactorSession) Receive() (*types.Transaction, error) {
	return _CFXsContract.Contract.Receive(&_CFXsContract.TransactOpts)
}

// CFXsContractCFXsCreatedIterator is returned from FilterCFXsCreated and is used to iterate over the raw logs and unpacked data for CFXsCreated events raised by the CFXsContract contract.
type CFXsContractCFXsCreatedIterator struct {
	Event *CFXsContractCFXsCreated // Event containing the contract specifics and raw log

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
func (it *CFXsContractCFXsCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CFXsContractCFXsCreated)
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
		it.Event = new(CFXsContractCFXsCreated)
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
func (it *CFXsContractCFXsCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CFXsContractCFXsCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CFXsContractCFXsCreated represents a CFXsCreated event raised by the CFXsContract contract.
type CFXsContractCFXsCreated struct {
	Id     *big.Int
	To     common.Address
	Amount *big.Int
	Data   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCFXsCreated is a free log retrieval operation binding the contract event 0xc619c1c30b367378bda467d6bebdd6f994ab4a6e274528b8707fa97c4752ca48.
//
// Solidity: event CFXsCreated(uint256 id, address to, uint256 amount, string data)
func (_CFXsContract *CFXsContractFilterer) FilterCFXsCreated(opts *bind.FilterOpts) (*CFXsContractCFXsCreatedIterator, error) {

	logs, sub, err := _CFXsContract.contract.FilterLogs(opts, "CFXsCreated")
	if err != nil {
		return nil, err
	}
	return &CFXsContractCFXsCreatedIterator{contract: _CFXsContract.contract, event: "CFXsCreated", logs: logs, sub: sub}, nil
}

// WatchCFXsCreated is a free log subscription operation binding the contract event 0xc619c1c30b367378bda467d6bebdd6f994ab4a6e274528b8707fa97c4752ca48.
//
// Solidity: event CFXsCreated(uint256 id, address to, uint256 amount, string data)
func (_CFXsContract *CFXsContractFilterer) WatchCFXsCreated(opts *bind.WatchOpts, sink chan<- *CFXsContractCFXsCreated) (event.Subscription, error) {

	logs, sub, err := _CFXsContract.contract.WatchLogs(opts, "CFXsCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CFXsContractCFXsCreated)
				if err := _CFXsContract.contract.UnpackLog(event, "CFXsCreated", log); err != nil {
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

// ParseCFXsCreated is a log parse operation binding the contract event 0xc619c1c30b367378bda467d6bebdd6f994ab4a6e274528b8707fa97c4752ca48.
//
// Solidity: event CFXsCreated(uint256 id, address to, uint256 amount, string data)
func (_CFXsContract *CFXsContractFilterer) ParseCFXsCreated(log types.Log) (*CFXsContractCFXsCreated, error) {
	event := new(CFXsContractCFXsCreated)
	if err := _CFXsContract.contract.UnpackLog(event, "CFXsCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CFXsContractCFXsDeletedIterator is returned from FilterCFXsDeleted and is used to iterate over the raw logs and unpacked data for CFXsDeleted events raised by the CFXsContract contract.
type CFXsContractCFXsDeletedIterator struct {
	Event *CFXsContractCFXsDeleted // Event containing the contract specifics and raw log

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
func (it *CFXsContractCFXsDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CFXsContractCFXsDeleted)
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
		it.Event = new(CFXsContractCFXsDeleted)
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
func (it *CFXsContractCFXsDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CFXsContractCFXsDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CFXsContractCFXsDeleted represents a CFXsDeleted event raised by the CFXsContract contract.
type CFXsContractCFXsDeleted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCFXsDeleted is a free log retrieval operation binding the contract event 0x54e9d3d6bf8951e4535558a5850039b2eb9776c9a1fa240f1c51d2ee046f4ac7.
//
// Solidity: event CFXsDeleted(uint256 id)
func (_CFXsContract *CFXsContractFilterer) FilterCFXsDeleted(opts *bind.FilterOpts) (*CFXsContractCFXsDeletedIterator, error) {

	logs, sub, err := _CFXsContract.contract.FilterLogs(opts, "CFXsDeleted")
	if err != nil {
		return nil, err
	}
	return &CFXsContractCFXsDeletedIterator{contract: _CFXsContract.contract, event: "CFXsDeleted", logs: logs, sub: sub}, nil
}

// WatchCFXsDeleted is a free log subscription operation binding the contract event 0x54e9d3d6bf8951e4535558a5850039b2eb9776c9a1fa240f1c51d2ee046f4ac7.
//
// Solidity: event CFXsDeleted(uint256 id)
func (_CFXsContract *CFXsContractFilterer) WatchCFXsDeleted(opts *bind.WatchOpts, sink chan<- *CFXsContractCFXsDeleted) (event.Subscription, error) {

	logs, sub, err := _CFXsContract.contract.WatchLogs(opts, "CFXsDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CFXsContractCFXsDeleted)
				if err := _CFXsContract.contract.UnpackLog(event, "CFXsDeleted", log); err != nil {
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

// ParseCFXsDeleted is a log parse operation binding the contract event 0x54e9d3d6bf8951e4535558a5850039b2eb9776c9a1fa240f1c51d2ee046f4ac7.
//
// Solidity: event CFXsDeleted(uint256 id)
func (_CFXsContract *CFXsContractFilterer) ParseCFXsDeleted(log types.Log) (*CFXsContractCFXsDeleted, error) {
	event := new(CFXsContractCFXsDeleted)
	if err := _CFXsContract.contract.UnpackLog(event, "CFXsDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CFXsContractCFXsEventIterator is returned from FilterCFXsEvent and is used to iterate over the raw logs and unpacked data for CFXsEvent events raised by the CFXsContract contract.
type CFXsContractCFXsEventIterator struct {
	Event *CFXsContractCFXsEvent // Event containing the contract specifics and raw log

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
func (it *CFXsContractCFXsEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CFXsContractCFXsEvent)
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
		it.Event = new(CFXsContractCFXsEvent)
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
func (it *CFXsContractCFXsEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CFXsContractCFXsEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CFXsContractCFXsEvent represents a CFXsEvent event raised by the CFXsContract contract.
type CFXsContractCFXsEvent struct {
	Id     *big.Int
	To     common.Address
	Amount *big.Int
	Data   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCFXsEvent is a free log retrieval operation binding the contract event 0x83c5a3ed95d27f7cb0f9946efaa434a0cf237615a04272bc2bf92d7eb3cf9761.
//
// Solidity: event CFXsEvent(uint256 id, address to, uint256 amount, string data)
func (_CFXsContract *CFXsContractFilterer) FilterCFXsEvent(opts *bind.FilterOpts) (*CFXsContractCFXsEventIterator, error) {

	logs, sub, err := _CFXsContract.contract.FilterLogs(opts, "CFXsEvent")
	if err != nil {
		return nil, err
	}
	return &CFXsContractCFXsEventIterator{contract: _CFXsContract.contract, event: "CFXsEvent", logs: logs, sub: sub}, nil
}

// WatchCFXsEvent is a free log subscription operation binding the contract event 0x83c5a3ed95d27f7cb0f9946efaa434a0cf237615a04272bc2bf92d7eb3cf9761.
//
// Solidity: event CFXsEvent(uint256 id, address to, uint256 amount, string data)
func (_CFXsContract *CFXsContractFilterer) WatchCFXsEvent(opts *bind.WatchOpts, sink chan<- *CFXsContractCFXsEvent) (event.Subscription, error) {

	logs, sub, err := _CFXsContract.contract.WatchLogs(opts, "CFXsEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CFXsContractCFXsEvent)
				if err := _CFXsContract.contract.UnpackLog(event, "CFXsEvent", log); err != nil {
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

// ParseCFXsEvent is a log parse operation binding the contract event 0x83c5a3ed95d27f7cb0f9946efaa434a0cf237615a04272bc2bf92d7eb3cf9761.
//
// Solidity: event CFXsEvent(uint256 id, address to, uint256 amount, string data)
func (_CFXsContract *CFXsContractFilterer) ParseCFXsEvent(log types.Log) (*CFXsContractCFXsEvent, error) {
	event := new(CFXsContractCFXsEvent)
	if err := _CFXsContract.contract.UnpackLog(event, "CFXsEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CFXsContractCFXsLockedIterator is returned from FilterCFXsLocked and is used to iterate over the raw logs and unpacked data for CFXsLocked events raised by the CFXsContract contract.
type CFXsContractCFXsLockedIterator struct {
	Event *CFXsContractCFXsLocked // Event containing the contract specifics and raw log

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
func (it *CFXsContractCFXsLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CFXsContractCFXsLocked)
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
		it.Event = new(CFXsContractCFXsLocked)
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
func (it *CFXsContractCFXsLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CFXsContractCFXsLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CFXsContractCFXsLocked represents a CFXsLocked event raised by the CFXsContract contract.
type CFXsContractCFXsLocked struct {
	CFXsId      *big.Int
	EtherAmount *big.Int
	Locktime    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCFXsLocked is a free log retrieval operation binding the contract event 0xa461494b4bb5b609c6502d7cd54f6572aa434d20cd7717a2b0ee197abae9eec9.
//
// Solidity: event CFXsLocked(uint256 indexed CFXsId, uint256 etherAmount, uint256 locktime)
func (_CFXsContract *CFXsContractFilterer) FilterCFXsLocked(opts *bind.FilterOpts, CFXsId []*big.Int) (*CFXsContractCFXsLockedIterator, error) {

	var CFXsIdRule []interface{}
	for _, CFXsIdItem := range CFXsId {
		CFXsIdRule = append(CFXsIdRule, CFXsIdItem)
	}

	logs, sub, err := _CFXsContract.contract.FilterLogs(opts, "CFXsLocked", CFXsIdRule)
	if err != nil {
		return nil, err
	}
	return &CFXsContractCFXsLockedIterator{contract: _CFXsContract.contract, event: "CFXsLocked", logs: logs, sub: sub}, nil
}

// WatchCFXsLocked is a free log subscription operation binding the contract event 0xa461494b4bb5b609c6502d7cd54f6572aa434d20cd7717a2b0ee197abae9eec9.
//
// Solidity: event CFXsLocked(uint256 indexed CFXsId, uint256 etherAmount, uint256 locktime)
func (_CFXsContract *CFXsContractFilterer) WatchCFXsLocked(opts *bind.WatchOpts, sink chan<- *CFXsContractCFXsLocked, CFXsId []*big.Int) (event.Subscription, error) {

	var CFXsIdRule []interface{}
	for _, CFXsIdItem := range CFXsId {
		CFXsIdRule = append(CFXsIdRule, CFXsIdItem)
	}

	logs, sub, err := _CFXsContract.contract.WatchLogs(opts, "CFXsLocked", CFXsIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CFXsContractCFXsLocked)
				if err := _CFXsContract.contract.UnpackLog(event, "CFXsLocked", log); err != nil {
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

// ParseCFXsLocked is a log parse operation binding the contract event 0xa461494b4bb5b609c6502d7cd54f6572aa434d20cd7717a2b0ee197abae9eec9.
//
// Solidity: event CFXsLocked(uint256 indexed CFXsId, uint256 etherAmount, uint256 locktime)
func (_CFXsContract *CFXsContractFilterer) ParseCFXsLocked(log types.Log) (*CFXsContractCFXsLocked, error) {
	event := new(CFXsContractCFXsLocked)
	if err := _CFXsContract.contract.UnpackLog(event, "CFXsLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CFXsContractCFXsUnlockedIterator is returned from FilterCFXsUnlocked and is used to iterate over the raw logs and unpacked data for CFXsUnlocked events raised by the CFXsContract contract.
type CFXsContractCFXsUnlockedIterator struct {
	Event *CFXsContractCFXsUnlocked // Event containing the contract specifics and raw log

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
func (it *CFXsContractCFXsUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CFXsContractCFXsUnlocked)
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
		it.Event = new(CFXsContractCFXsUnlocked)
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
func (it *CFXsContractCFXsUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CFXsContractCFXsUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CFXsContractCFXsUnlocked represents a CFXsUnlocked event raised by the CFXsContract contract.
type CFXsContractCFXsUnlocked struct {
	CFXsId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCFXsUnlocked is a free log retrieval operation binding the contract event 0x837ebae0d82e9b2e4a3296822f99dfb8a51ff0fb2b7cbde01ce97aad8ea9fde2.
//
// Solidity: event CFXsUnlocked(uint256 indexed CFXsId)
func (_CFXsContract *CFXsContractFilterer) FilterCFXsUnlocked(opts *bind.FilterOpts, CFXsId []*big.Int) (*CFXsContractCFXsUnlockedIterator, error) {

	var CFXsIdRule []interface{}
	for _, CFXsIdItem := range CFXsId {
		CFXsIdRule = append(CFXsIdRule, CFXsIdItem)
	}

	logs, sub, err := _CFXsContract.contract.FilterLogs(opts, "CFXsUnlocked", CFXsIdRule)
	if err != nil {
		return nil, err
	}
	return &CFXsContractCFXsUnlockedIterator{contract: _CFXsContract.contract, event: "CFXsUnlocked", logs: logs, sub: sub}, nil
}

// WatchCFXsUnlocked is a free log subscription operation binding the contract event 0x837ebae0d82e9b2e4a3296822f99dfb8a51ff0fb2b7cbde01ce97aad8ea9fde2.
//
// Solidity: event CFXsUnlocked(uint256 indexed CFXsId)
func (_CFXsContract *CFXsContractFilterer) WatchCFXsUnlocked(opts *bind.WatchOpts, sink chan<- *CFXsContractCFXsUnlocked, CFXsId []*big.Int) (event.Subscription, error) {

	var CFXsIdRule []interface{}
	for _, CFXsIdItem := range CFXsId {
		CFXsIdRule = append(CFXsIdRule, CFXsIdItem)
	}

	logs, sub, err := _CFXsContract.contract.WatchLogs(opts, "CFXsUnlocked", CFXsIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CFXsContractCFXsUnlocked)
				if err := _CFXsContract.contract.UnpackLog(event, "CFXsUnlocked", log); err != nil {
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

// ParseCFXsUnlocked is a log parse operation binding the contract event 0x837ebae0d82e9b2e4a3296822f99dfb8a51ff0fb2b7cbde01ce97aad8ea9fde2.
//
// Solidity: event CFXsUnlocked(uint256 indexed CFXsId)
func (_CFXsContract *CFXsContractFilterer) ParseCFXsUnlocked(log types.Log) (*CFXsContractCFXsUnlocked, error) {
	event := new(CFXsContractCFXsUnlocked)
	if err := _CFXsContract.contract.UnpackLog(event, "CFXsUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
