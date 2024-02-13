// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Erc20AltMetaData contains all meta data concerning the Erc20Alt contract.
var Erc20AltMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"push\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"bytes32\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"pull\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"symbol_\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}

// Erc20AltABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20AltMetaData.ABI instead.
var Erc20AltABI = Erc20AltMetaData.ABI

// Erc20Alt is an auto generated Go binding around an Ethereum contract.
type Erc20Alt struct {
	Erc20AltCaller     // Read-only binding to the contract
	Erc20AltTransactor // Write-only binding to the contract
	Erc20AltFilterer   // Log filterer for contract events
}

// Erc20AltCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20AltCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20AltTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20AltTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20AltFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20AltFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20AltSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20AltSession struct {
	Contract     *Erc20Alt         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20AltCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20AltCallerSession struct {
	Contract *Erc20AltCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Erc20AltTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20AltTransactorSession struct {
	Contract     *Erc20AltTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Erc20AltRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20AltRaw struct {
	Contract *Erc20Alt // Generic contract binding to access the raw methods on
}

// Erc20AltCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20AltCallerRaw struct {
	Contract *Erc20AltCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20AltTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20AltTransactorRaw struct {
	Contract *Erc20AltTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20Alt creates a new instance of Erc20Alt, bound to a specific deployed contract.
func NewErc20Alt(address common.Address, backend bind.ContractBackend) (*Erc20Alt, error) {
	contract, err := bindErc20Alt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20Alt{Erc20AltCaller: Erc20AltCaller{contract: contract}, Erc20AltTransactor: Erc20AltTransactor{contract: contract}, Erc20AltFilterer: Erc20AltFilterer{contract: contract}}, nil
}

// NewErc20AltCaller creates a new read-only instance of Erc20Alt, bound to a specific deployed contract.
func NewErc20AltCaller(address common.Address, caller bind.ContractCaller) (*Erc20AltCaller, error) {
	contract, err := bindErc20Alt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20AltCaller{contract: contract}, nil
}

// NewErc20AltTransactor creates a new write-only instance of Erc20Alt, bound to a specific deployed contract.
func NewErc20AltTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20AltTransactor, error) {
	contract, err := bindErc20Alt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20AltTransactor{contract: contract}, nil
}

// NewErc20AltFilterer creates a new log filterer instance of Erc20Alt, bound to a specific deployed contract.
func NewErc20AltFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20AltFilterer, error) {
	contract, err := bindErc20Alt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20AltFilterer{contract: contract}, nil
}

// bindErc20Alt binds a generic wrapper to an already deployed contract.
func bindErc20Alt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20AltMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Alt *Erc20AltRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Alt.Contract.Erc20AltCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Alt *Erc20AltRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Erc20AltTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Alt *Erc20AltRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Erc20AltTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Alt *Erc20AltCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Alt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Alt *Erc20AltTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Alt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Alt *Erc20AltTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Alt.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address guy) returns(uint256)
func (_Erc20Alt *Erc20AltCaller) Allowance(opts *bind.CallOpts, src common.Address, guy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Alt.contract.Call(opts, &out, "allowance", src, guy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address guy) returns(uint256)
func (_Erc20Alt *Erc20AltSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _Erc20Alt.Contract.Allowance(&_Erc20Alt.CallOpts, src, guy)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address guy) returns(uint256)
func (_Erc20Alt *Erc20AltCallerSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _Erc20Alt.Contract.Allowance(&_Erc20Alt.CallOpts, src, guy)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() returns(address)
func (_Erc20Alt *Erc20AltCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20Alt.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() returns(address)
func (_Erc20Alt *Erc20AltSession) Authority() (common.Address, error) {
	return _Erc20Alt.Contract.Authority(&_Erc20Alt.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() returns(address)
func (_Erc20Alt *Erc20AltCallerSession) Authority() (common.Address, error) {
	return _Erc20Alt.Contract.Authority(&_Erc20Alt.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address src) returns(uint256)
func (_Erc20Alt *Erc20AltCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Alt.contract.Call(opts, &out, "balanceOf", src)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address src) returns(uint256)
func (_Erc20Alt *Erc20AltSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _Erc20Alt.Contract.BalanceOf(&_Erc20Alt.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address src) returns(uint256)
func (_Erc20Alt *Erc20AltCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _Erc20Alt.Contract.BalanceOf(&_Erc20Alt.CallOpts, src)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256)
func (_Erc20Alt *Erc20AltCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Alt.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256)
func (_Erc20Alt *Erc20AltSession) Decimals() (*big.Int, error) {
	return _Erc20Alt.Contract.Decimals(&_Erc20Alt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256)
func (_Erc20Alt *Erc20AltCallerSession) Decimals() (*big.Int, error) {
	return _Erc20Alt.Contract.Decimals(&_Erc20Alt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(bytes32)
func (_Erc20Alt *Erc20AltCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20Alt.contract.Call(opts, &out, "name")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(bytes32)
func (_Erc20Alt *Erc20AltSession) Name() ([32]byte, error) {
	return _Erc20Alt.Contract.Name(&_Erc20Alt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(bytes32)
func (_Erc20Alt *Erc20AltCallerSession) Name() ([32]byte, error) {
	return _Erc20Alt.Contract.Name(&_Erc20Alt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Erc20Alt *Erc20AltCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20Alt.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Erc20Alt *Erc20AltSession) Owner() (common.Address, error) {
	return _Erc20Alt.Contract.Owner(&_Erc20Alt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Erc20Alt *Erc20AltCallerSession) Owner() (common.Address, error) {
	return _Erc20Alt.Contract.Owner(&_Erc20Alt.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() returns(bool)
func (_Erc20Alt *Erc20AltCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc20Alt.contract.Call(opts, &out, "stopped")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() returns(bool)
func (_Erc20Alt *Erc20AltSession) Stopped() (bool, error) {
	return _Erc20Alt.Contract.Stopped(&_Erc20Alt.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() returns(bool)
func (_Erc20Alt *Erc20AltCallerSession) Stopped() (bool, error) {
	return _Erc20Alt.Contract.Stopped(&_Erc20Alt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(bytes32)
func (_Erc20Alt *Erc20AltCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20Alt.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(bytes32)
func (_Erc20Alt *Erc20AltSession) Symbol() ([32]byte, error) {
	return _Erc20Alt.Contract.Symbol(&_Erc20Alt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(bytes32)
func (_Erc20Alt *Erc20AltCallerSession) Symbol() ([32]byte, error) {
	return _Erc20Alt.Contract.Symbol(&_Erc20Alt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Erc20Alt *Erc20AltCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Alt.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Erc20Alt *Erc20AltSession) TotalSupply() (*big.Int, error) {
	return _Erc20Alt.Contract.TotalSupply(&_Erc20Alt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_Erc20Alt *Erc20AltCallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20Alt.Contract.TotalSupply(&_Erc20Alt.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Erc20Alt *Erc20AltTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Erc20Alt *Erc20AltSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Approve(&_Erc20Alt.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Erc20Alt *Erc20AltTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Approve(&_Erc20Alt.TransactOpts, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(uint128 wad) returns()
func (_Erc20Alt *Erc20AltTransactor) Burn(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.contract.Transact(opts, "burn", wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(uint128 wad) returns()
func (_Erc20Alt *Erc20AltSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Burn(&_Erc20Alt.TransactOpts, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(uint128 wad) returns()
func (_Erc20Alt *Erc20AltTransactorSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Burn(&_Erc20Alt.TransactOpts, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(uint128 wad) returns()
func (_Erc20Alt *Erc20AltTransactor) Mint(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.contract.Transact(opts, "mint", wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(uint128 wad) returns()
func (_Erc20Alt *Erc20AltSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Mint(&_Erc20Alt.TransactOpts, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(uint128 wad) returns()
func (_Erc20Alt *Erc20AltTransactorSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Mint(&_Erc20Alt.TransactOpts, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(address src, uint128 wad) returns(bool)
func (_Erc20Alt *Erc20AltTransactor) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.contract.Transact(opts, "pull", src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(address src, uint128 wad) returns(bool)
func (_Erc20Alt *Erc20AltSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Pull(&_Erc20Alt.TransactOpts, src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(address src, uint128 wad) returns(bool)
func (_Erc20Alt *Erc20AltTransactorSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Pull(&_Erc20Alt.TransactOpts, src, wad)
}

// Push is a paid mutator transaction binding the contract method 0x3452f51d.
//
// Solidity: function push(address dst, uint128 wad) returns(bool)
func (_Erc20Alt *Erc20AltTransactor) Push(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.contract.Transact(opts, "push", dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0x3452f51d.
//
// Solidity: function push(address dst, uint128 wad) returns(bool)
func (_Erc20Alt *Erc20AltSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Push(&_Erc20Alt.TransactOpts, dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0x3452f51d.
//
// Solidity: function push(address dst, uint128 wad) returns(bool)
func (_Erc20Alt *Erc20AltTransactorSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Push(&_Erc20Alt.TransactOpts, dst, wad)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Erc20Alt *Erc20AltTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Erc20Alt.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Erc20Alt *Erc20AltSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Erc20Alt.Contract.SetAuthority(&_Erc20Alt.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_Erc20Alt *Erc20AltTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Erc20Alt.Contract.SetAuthority(&_Erc20Alt.TransactOpts, authority_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(bytes32 name_) returns()
func (_Erc20Alt *Erc20AltTransactor) SetName(opts *bind.TransactOpts, name_ [32]byte) (*types.Transaction, error) {
	return _Erc20Alt.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(bytes32 name_) returns()
func (_Erc20Alt *Erc20AltSession) SetName(name_ [32]byte) (*types.Transaction, error) {
	return _Erc20Alt.Contract.SetName(&_Erc20Alt.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(bytes32 name_) returns()
func (_Erc20Alt *Erc20AltTransactorSession) SetName(name_ [32]byte) (*types.Transaction, error) {
	return _Erc20Alt.Contract.SetName(&_Erc20Alt.TransactOpts, name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Erc20Alt *Erc20AltTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Erc20Alt.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Erc20Alt *Erc20AltSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Erc20Alt.Contract.SetOwner(&_Erc20Alt.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_Erc20Alt *Erc20AltTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Erc20Alt.Contract.SetOwner(&_Erc20Alt.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Erc20Alt *Erc20AltTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Alt.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Erc20Alt *Erc20AltSession) Start() (*types.Transaction, error) {
	return _Erc20Alt.Contract.Start(&_Erc20Alt.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Erc20Alt *Erc20AltTransactorSession) Start() (*types.Transaction, error) {
	return _Erc20Alt.Contract.Start(&_Erc20Alt.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Erc20Alt *Erc20AltTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Alt.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Erc20Alt *Erc20AltSession) Stop() (*types.Transaction, error) {
	return _Erc20Alt.Contract.Stop(&_Erc20Alt.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Erc20Alt *Erc20AltTransactorSession) Stop() (*types.Transaction, error) {
	return _Erc20Alt.Contract.Stop(&_Erc20Alt.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Erc20Alt *Erc20AltTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Erc20Alt *Erc20AltSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Transfer(&_Erc20Alt.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Erc20Alt *Erc20AltTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.Transfer(&_Erc20Alt.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Erc20Alt *Erc20AltTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Erc20Alt *Erc20AltSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.TransferFrom(&_Erc20Alt.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Erc20Alt *Erc20AltTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Erc20Alt.Contract.TransferFrom(&_Erc20Alt.TransactOpts, src, dst, wad)
}

// Erc20AltApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20Alt contract.
type Erc20AltApprovalIterator struct {
	Event *Erc20AltApproval // Event containing the contract specifics and raw log

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
func (it *Erc20AltApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20AltApproval)
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
		it.Event = new(Erc20AltApproval)
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
func (it *Erc20AltApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20AltApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20AltApproval represents a Approval event raised by the Erc20Alt contract.
type Erc20AltApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20Alt *Erc20AltFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20AltApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20Alt.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20AltApprovalIterator{contract: _Erc20Alt.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20Alt *Erc20AltFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20AltApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20Alt.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20AltApproval)
				if err := _Erc20Alt.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20Alt *Erc20AltFilterer) ParseApproval(log types.Log) (*Erc20AltApproval, error) {
	event := new(Erc20AltApproval)
	if err := _Erc20Alt.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20AltLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Erc20Alt contract.
type Erc20AltLogSetAuthorityIterator struct {
	Event *Erc20AltLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *Erc20AltLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20AltLogSetAuthority)
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
		it.Event = new(Erc20AltLogSetAuthority)
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
func (it *Erc20AltLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20AltLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20AltLogSetAuthority represents a LogSetAuthority event raised by the Erc20Alt contract.
type Erc20AltLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Erc20Alt *Erc20AltFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*Erc20AltLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Erc20Alt.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &Erc20AltLogSetAuthorityIterator{contract: _Erc20Alt.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Erc20Alt *Erc20AltFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *Erc20AltLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Erc20Alt.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20AltLogSetAuthority)
				if err := _Erc20Alt.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// ParseLogSetAuthority is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_Erc20Alt *Erc20AltFilterer) ParseLogSetAuthority(log types.Log) (*Erc20AltLogSetAuthority, error) {
	event := new(Erc20AltLogSetAuthority)
	if err := _Erc20Alt.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20AltLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Erc20Alt contract.
type Erc20AltLogSetOwnerIterator struct {
	Event *Erc20AltLogSetOwner // Event containing the contract specifics and raw log

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
func (it *Erc20AltLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20AltLogSetOwner)
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
		it.Event = new(Erc20AltLogSetOwner)
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
func (it *Erc20AltLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20AltLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20AltLogSetOwner represents a LogSetOwner event raised by the Erc20Alt contract.
type Erc20AltLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Erc20Alt *Erc20AltFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*Erc20AltLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Erc20Alt.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &Erc20AltLogSetOwnerIterator{contract: _Erc20Alt.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Erc20Alt *Erc20AltFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *Erc20AltLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Erc20Alt.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20AltLogSetOwner)
				if err := _Erc20Alt.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// ParseLogSetOwner is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_Erc20Alt *Erc20AltFilterer) ParseLogSetOwner(log types.Log) (*Erc20AltLogSetOwner, error) {
	event := new(Erc20AltLogSetOwner)
	if err := _Erc20Alt.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20AltTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20Alt contract.
type Erc20AltTransferIterator struct {
	Event *Erc20AltTransfer // Event containing the contract specifics and raw log

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
func (it *Erc20AltTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20AltTransfer)
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
		it.Event = new(Erc20AltTransfer)
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
func (it *Erc20AltTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20AltTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20AltTransfer represents a Transfer event raised by the Erc20Alt contract.
type Erc20AltTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20Alt *Erc20AltFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20AltTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20Alt.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20AltTransferIterator{contract: _Erc20Alt.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20Alt *Erc20AltFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20AltTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20Alt.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20AltTransfer)
				if err := _Erc20Alt.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20Alt *Erc20AltFilterer) ParseTransfer(log types.Log) (*Erc20AltTransfer, error) {
	event := new(Erc20AltTransfer)
	if err := _Erc20Alt.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
