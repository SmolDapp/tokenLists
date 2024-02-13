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

// AjnaPoolFactoryMetaData contains all meta data concerning the AjnaPoolFactory contract.
var AjnaPoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ajna_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CreateFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DecimalsNotCompliant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeployQuoteCollateralSameToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeployWithZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"name\":\"PoolAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolInterestRateInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"subsetHash_\",\"type\":\"bytes32\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC20_NON_SUBSET_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ajna\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"interestRate_\",\"type\":\"uint256\"}],\"name\":\"deployPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deployedPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deployedPoolsList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeployedPoolsList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfDeployedPools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"contractERC20Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AjnaPoolFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AjnaPoolFactoryMetaData.ABI instead.
var AjnaPoolFactoryABI = AjnaPoolFactoryMetaData.ABI

// AjnaPoolFactory is an auto generated Go binding around an Ethereum contract.
type AjnaPoolFactory struct {
	AjnaPoolFactoryCaller     // Read-only binding to the contract
	AjnaPoolFactoryTransactor // Write-only binding to the contract
	AjnaPoolFactoryFilterer   // Log filterer for contract events
}

// AjnaPoolFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AjnaPoolFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AjnaPoolFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AjnaPoolFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AjnaPoolFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AjnaPoolFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AjnaPoolFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AjnaPoolFactorySession struct {
	Contract     *AjnaPoolFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AjnaPoolFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AjnaPoolFactoryCallerSession struct {
	Contract *AjnaPoolFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AjnaPoolFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AjnaPoolFactoryTransactorSession struct {
	Contract     *AjnaPoolFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AjnaPoolFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AjnaPoolFactoryRaw struct {
	Contract *AjnaPoolFactory // Generic contract binding to access the raw methods on
}

// AjnaPoolFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AjnaPoolFactoryCallerRaw struct {
	Contract *AjnaPoolFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AjnaPoolFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AjnaPoolFactoryTransactorRaw struct {
	Contract *AjnaPoolFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAjnaPoolFactory creates a new instance of AjnaPoolFactory, bound to a specific deployed contract.
func NewAjnaPoolFactory(address common.Address, backend bind.ContractBackend) (*AjnaPoolFactory, error) {
	contract, err := bindAjnaPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolFactory{AjnaPoolFactoryCaller: AjnaPoolFactoryCaller{contract: contract}, AjnaPoolFactoryTransactor: AjnaPoolFactoryTransactor{contract: contract}, AjnaPoolFactoryFilterer: AjnaPoolFactoryFilterer{contract: contract}}, nil
}

// NewAjnaPoolFactoryCaller creates a new read-only instance of AjnaPoolFactory, bound to a specific deployed contract.
func NewAjnaPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*AjnaPoolFactoryCaller, error) {
	contract, err := bindAjnaPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolFactoryCaller{contract: contract}, nil
}

// NewAjnaPoolFactoryTransactor creates a new write-only instance of AjnaPoolFactory, bound to a specific deployed contract.
func NewAjnaPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AjnaPoolFactoryTransactor, error) {
	contract, err := bindAjnaPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolFactoryTransactor{contract: contract}, nil
}

// NewAjnaPoolFactoryFilterer creates a new log filterer instance of AjnaPoolFactory, bound to a specific deployed contract.
func NewAjnaPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AjnaPoolFactoryFilterer, error) {
	contract, err := bindAjnaPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AjnaPoolFactoryFilterer{contract: contract}, nil
}

// bindAjnaPoolFactory binds a generic wrapper to an already deployed contract.
func bindAjnaPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AjnaPoolFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AjnaPoolFactory *AjnaPoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AjnaPoolFactory.Contract.AjnaPoolFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AjnaPoolFactory *AjnaPoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AjnaPoolFactory.Contract.AjnaPoolFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AjnaPoolFactory *AjnaPoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AjnaPoolFactory.Contract.AjnaPoolFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AjnaPoolFactory *AjnaPoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AjnaPoolFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AjnaPoolFactory *AjnaPoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AjnaPoolFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AjnaPoolFactory *AjnaPoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AjnaPoolFactory.Contract.contract.Transact(opts, method, params...)
}

// ERC20NONSUBSETHASH is a free data retrieval call binding the contract method 0xc38dc7fc.
//
// Solidity: function ERC20_NON_SUBSET_HASH() view returns(bytes32)
func (_AjnaPoolFactory *AjnaPoolFactoryCaller) ERC20NONSUBSETHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AjnaPoolFactory.contract.Call(opts, &out, "ERC20_NON_SUBSET_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20NONSUBSETHASH is a free data retrieval call binding the contract method 0xc38dc7fc.
//
// Solidity: function ERC20_NON_SUBSET_HASH() view returns(bytes32)
func (_AjnaPoolFactory *AjnaPoolFactorySession) ERC20NONSUBSETHASH() ([32]byte, error) {
	return _AjnaPoolFactory.Contract.ERC20NONSUBSETHASH(&_AjnaPoolFactory.CallOpts)
}

// ERC20NONSUBSETHASH is a free data retrieval call binding the contract method 0xc38dc7fc.
//
// Solidity: function ERC20_NON_SUBSET_HASH() view returns(bytes32)
func (_AjnaPoolFactory *AjnaPoolFactoryCallerSession) ERC20NONSUBSETHASH() ([32]byte, error) {
	return _AjnaPoolFactory.Contract.ERC20NONSUBSETHASH(&_AjnaPoolFactory.CallOpts)
}

// MAXRATE is a free data retrieval call binding the contract method 0xc24dbebd.
//
// Solidity: function MAX_RATE() view returns(uint256)
func (_AjnaPoolFactory *AjnaPoolFactoryCaller) MAXRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPoolFactory.contract.Call(opts, &out, "MAX_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXRATE is a free data retrieval call binding the contract method 0xc24dbebd.
//
// Solidity: function MAX_RATE() view returns(uint256)
func (_AjnaPoolFactory *AjnaPoolFactorySession) MAXRATE() (*big.Int, error) {
	return _AjnaPoolFactory.Contract.MAXRATE(&_AjnaPoolFactory.CallOpts)
}

// MAXRATE is a free data retrieval call binding the contract method 0xc24dbebd.
//
// Solidity: function MAX_RATE() view returns(uint256)
func (_AjnaPoolFactory *AjnaPoolFactoryCallerSession) MAXRATE() (*big.Int, error) {
	return _AjnaPoolFactory.Contract.MAXRATE(&_AjnaPoolFactory.CallOpts)
}

// MINRATE is a free data retrieval call binding the contract method 0xd819bfef.
//
// Solidity: function MIN_RATE() view returns(uint256)
func (_AjnaPoolFactory *AjnaPoolFactoryCaller) MINRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPoolFactory.contract.Call(opts, &out, "MIN_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINRATE is a free data retrieval call binding the contract method 0xd819bfef.
//
// Solidity: function MIN_RATE() view returns(uint256)
func (_AjnaPoolFactory *AjnaPoolFactorySession) MINRATE() (*big.Int, error) {
	return _AjnaPoolFactory.Contract.MINRATE(&_AjnaPoolFactory.CallOpts)
}

// MINRATE is a free data retrieval call binding the contract method 0xd819bfef.
//
// Solidity: function MIN_RATE() view returns(uint256)
func (_AjnaPoolFactory *AjnaPoolFactoryCallerSession) MINRATE() (*big.Int, error) {
	return _AjnaPoolFactory.Contract.MINRATE(&_AjnaPoolFactory.CallOpts)
}

// Ajna is a free data retrieval call binding the contract method 0xbb6da0dd.
//
// Solidity: function ajna() view returns(address)
func (_AjnaPoolFactory *AjnaPoolFactoryCaller) Ajna(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AjnaPoolFactory.contract.Call(opts, &out, "ajna")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ajna is a free data retrieval call binding the contract method 0xbb6da0dd.
//
// Solidity: function ajna() view returns(address)
func (_AjnaPoolFactory *AjnaPoolFactorySession) Ajna() (common.Address, error) {
	return _AjnaPoolFactory.Contract.Ajna(&_AjnaPoolFactory.CallOpts)
}

// Ajna is a free data retrieval call binding the contract method 0xbb6da0dd.
//
// Solidity: function ajna() view returns(address)
func (_AjnaPoolFactory *AjnaPoolFactoryCallerSession) Ajna() (common.Address, error) {
	return _AjnaPoolFactory.Contract.Ajna(&_AjnaPoolFactory.CallOpts)
}

// DeployedPools is a free data retrieval call binding the contract method 0x7f165b0b.
//
// Solidity: function deployedPools(bytes32 , address , address ) view returns(address)
func (_AjnaPoolFactory *AjnaPoolFactoryCaller) DeployedPools(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address, arg2 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AjnaPoolFactory.contract.Call(opts, &out, "deployedPools", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeployedPools is a free data retrieval call binding the contract method 0x7f165b0b.
//
// Solidity: function deployedPools(bytes32 , address , address ) view returns(address)
func (_AjnaPoolFactory *AjnaPoolFactorySession) DeployedPools(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (common.Address, error) {
	return _AjnaPoolFactory.Contract.DeployedPools(&_AjnaPoolFactory.CallOpts, arg0, arg1, arg2)
}

// DeployedPools is a free data retrieval call binding the contract method 0x7f165b0b.
//
// Solidity: function deployedPools(bytes32 , address , address ) view returns(address)
func (_AjnaPoolFactory *AjnaPoolFactoryCallerSession) DeployedPools(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (common.Address, error) {
	return _AjnaPoolFactory.Contract.DeployedPools(&_AjnaPoolFactory.CallOpts, arg0, arg1, arg2)
}

// DeployedPoolsList is a free data retrieval call binding the contract method 0xa387245c.
//
// Solidity: function deployedPoolsList(uint256 ) view returns(address)
func (_AjnaPoolFactory *AjnaPoolFactoryCaller) DeployedPoolsList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AjnaPoolFactory.contract.Call(opts, &out, "deployedPoolsList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeployedPoolsList is a free data retrieval call binding the contract method 0xa387245c.
//
// Solidity: function deployedPoolsList(uint256 ) view returns(address)
func (_AjnaPoolFactory *AjnaPoolFactorySession) DeployedPoolsList(arg0 *big.Int) (common.Address, error) {
	return _AjnaPoolFactory.Contract.DeployedPoolsList(&_AjnaPoolFactory.CallOpts, arg0)
}

// DeployedPoolsList is a free data retrieval call binding the contract method 0xa387245c.
//
// Solidity: function deployedPoolsList(uint256 ) view returns(address)
func (_AjnaPoolFactory *AjnaPoolFactoryCallerSession) DeployedPoolsList(arg0 *big.Int) (common.Address, error) {
	return _AjnaPoolFactory.Contract.DeployedPoolsList(&_AjnaPoolFactory.CallOpts, arg0)
}

// GetDeployedPoolsList is a free data retrieval call binding the contract method 0x2b6983af.
//
// Solidity: function getDeployedPoolsList() view returns(address[])
func (_AjnaPoolFactory *AjnaPoolFactoryCaller) GetDeployedPoolsList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AjnaPoolFactory.contract.Call(opts, &out, "getDeployedPoolsList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeployedPoolsList is a free data retrieval call binding the contract method 0x2b6983af.
//
// Solidity: function getDeployedPoolsList() view returns(address[])
func (_AjnaPoolFactory *AjnaPoolFactorySession) GetDeployedPoolsList() ([]common.Address, error) {
	return _AjnaPoolFactory.Contract.GetDeployedPoolsList(&_AjnaPoolFactory.CallOpts)
}

// GetDeployedPoolsList is a free data retrieval call binding the contract method 0x2b6983af.
//
// Solidity: function getDeployedPoolsList() view returns(address[])
func (_AjnaPoolFactory *AjnaPoolFactoryCallerSession) GetDeployedPoolsList() ([]common.Address, error) {
	return _AjnaPoolFactory.Contract.GetDeployedPoolsList(&_AjnaPoolFactory.CallOpts)
}

// GetNumberOfDeployedPools is a free data retrieval call binding the contract method 0xb3d4cfa4.
//
// Solidity: function getNumberOfDeployedPools() view returns(uint256)
func (_AjnaPoolFactory *AjnaPoolFactoryCaller) GetNumberOfDeployedPools(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AjnaPoolFactory.contract.Call(opts, &out, "getNumberOfDeployedPools")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfDeployedPools is a free data retrieval call binding the contract method 0xb3d4cfa4.
//
// Solidity: function getNumberOfDeployedPools() view returns(uint256)
func (_AjnaPoolFactory *AjnaPoolFactorySession) GetNumberOfDeployedPools() (*big.Int, error) {
	return _AjnaPoolFactory.Contract.GetNumberOfDeployedPools(&_AjnaPoolFactory.CallOpts)
}

// GetNumberOfDeployedPools is a free data retrieval call binding the contract method 0xb3d4cfa4.
//
// Solidity: function getNumberOfDeployedPools() view returns(uint256)
func (_AjnaPoolFactory *AjnaPoolFactoryCallerSession) GetNumberOfDeployedPools() (*big.Int, error) {
	return _AjnaPoolFactory.Contract.GetNumberOfDeployedPools(&_AjnaPoolFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_AjnaPoolFactory *AjnaPoolFactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AjnaPoolFactory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_AjnaPoolFactory *AjnaPoolFactorySession) Implementation() (common.Address, error) {
	return _AjnaPoolFactory.Contract.Implementation(&_AjnaPoolFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_AjnaPoolFactory *AjnaPoolFactoryCallerSession) Implementation() (common.Address, error) {
	return _AjnaPoolFactory.Contract.Implementation(&_AjnaPoolFactory.CallOpts)
}

// DeployPool is a paid mutator transaction binding the contract method 0xa3232bf3.
//
// Solidity: function deployPool(address collateral_, address quote_, uint256 interestRate_) returns(address pool_)
func (_AjnaPoolFactory *AjnaPoolFactoryTransactor) DeployPool(opts *bind.TransactOpts, collateral_ common.Address, quote_ common.Address, interestRate_ *big.Int) (*types.Transaction, error) {
	return _AjnaPoolFactory.contract.Transact(opts, "deployPool", collateral_, quote_, interestRate_)
}

// DeployPool is a paid mutator transaction binding the contract method 0xa3232bf3.
//
// Solidity: function deployPool(address collateral_, address quote_, uint256 interestRate_) returns(address pool_)
func (_AjnaPoolFactory *AjnaPoolFactorySession) DeployPool(collateral_ common.Address, quote_ common.Address, interestRate_ *big.Int) (*types.Transaction, error) {
	return _AjnaPoolFactory.Contract.DeployPool(&_AjnaPoolFactory.TransactOpts, collateral_, quote_, interestRate_)
}

// DeployPool is a paid mutator transaction binding the contract method 0xa3232bf3.
//
// Solidity: function deployPool(address collateral_, address quote_, uint256 interestRate_) returns(address pool_)
func (_AjnaPoolFactory *AjnaPoolFactoryTransactorSession) DeployPool(collateral_ common.Address, quote_ common.Address, interestRate_ *big.Int) (*types.Transaction, error) {
	return _AjnaPoolFactory.Contract.DeployPool(&_AjnaPoolFactory.TransactOpts, collateral_, quote_, interestRate_)
}

// AjnaPoolFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the AjnaPoolFactory contract.
type AjnaPoolFactoryPoolCreatedIterator struct {
	Event *AjnaPoolFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *AjnaPoolFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AjnaPoolFactoryPoolCreated)
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
		it.Event = new(AjnaPoolFactoryPoolCreated)
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
func (it *AjnaPoolFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AjnaPoolFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AjnaPoolFactoryPoolCreated represents a PoolCreated event raised by the AjnaPoolFactory contract.
type AjnaPoolFactoryPoolCreated struct {
	Pool       common.Address
	SubsetHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xee1fe091a5213b321c2662b35c0b7cd0d35d10dbcab52b3c9b8768983c67bce3.
//
// Solidity: event PoolCreated(address pool_, bytes32 subsetHash_)
func (_AjnaPoolFactory *AjnaPoolFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts) (*AjnaPoolFactoryPoolCreatedIterator, error) {

	logs, sub, err := _AjnaPoolFactory.contract.FilterLogs(opts, "PoolCreated")
	if err != nil {
		return nil, err
	}
	return &AjnaPoolFactoryPoolCreatedIterator{contract: _AjnaPoolFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xee1fe091a5213b321c2662b35c0b7cd0d35d10dbcab52b3c9b8768983c67bce3.
//
// Solidity: event PoolCreated(address pool_, bytes32 subsetHash_)
func (_AjnaPoolFactory *AjnaPoolFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *AjnaPoolFactoryPoolCreated) (event.Subscription, error) {

	logs, sub, err := _AjnaPoolFactory.contract.WatchLogs(opts, "PoolCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AjnaPoolFactoryPoolCreated)
				if err := _AjnaPoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0xee1fe091a5213b321c2662b35c0b7cd0d35d10dbcab52b3c9b8768983c67bce3.
//
// Solidity: event PoolCreated(address pool_, bytes32 subsetHash_)
func (_AjnaPoolFactory *AjnaPoolFactoryFilterer) ParsePoolCreated(log types.Log) (*AjnaPoolFactoryPoolCreated, error) {
	event := new(AjnaPoolFactoryPoolCreated)
	if err := _AjnaPoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
