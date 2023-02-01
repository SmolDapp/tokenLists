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
)

// UniV3FactoryMetaData contains all meta data concerning the UniV3Factory contract.
var UniV3FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"FeeAmountEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"enableFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniV3FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use UniV3FactoryMetaData.ABI instead.
var UniV3FactoryABI = UniV3FactoryMetaData.ABI

// UniV3Factory is an auto generated Go binding around an Ethereum contract.
type UniV3Factory struct {
	UniV3FactoryCaller     // Read-only binding to the contract
	UniV3FactoryTransactor // Write-only binding to the contract
	UniV3FactoryFilterer   // Log filterer for contract events
}

// UniV3FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniV3FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniV3FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniV3FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniV3FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniV3FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniV3FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniV3FactorySession struct {
	Contract     *UniV3Factory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniV3FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniV3FactoryCallerSession struct {
	Contract *UniV3FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UniV3FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniV3FactoryTransactorSession struct {
	Contract     *UniV3FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UniV3FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniV3FactoryRaw struct {
	Contract *UniV3Factory // Generic contract binding to access the raw methods on
}

// UniV3FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniV3FactoryCallerRaw struct {
	Contract *UniV3FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UniV3FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniV3FactoryTransactorRaw struct {
	Contract *UniV3FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniV3Factory creates a new instance of UniV3Factory, bound to a specific deployed contract.
func NewUniV3Factory(address common.Address, backend bind.ContractBackend) (*UniV3Factory, error) {
	contract, err := bindUniV3Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniV3Factory{UniV3FactoryCaller: UniV3FactoryCaller{contract: contract}, UniV3FactoryTransactor: UniV3FactoryTransactor{contract: contract}, UniV3FactoryFilterer: UniV3FactoryFilterer{contract: contract}}, nil
}

// NewUniV3FactoryCaller creates a new read-only instance of UniV3Factory, bound to a specific deployed contract.
func NewUniV3FactoryCaller(address common.Address, caller bind.ContractCaller) (*UniV3FactoryCaller, error) {
	contract, err := bindUniV3Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniV3FactoryCaller{contract: contract}, nil
}

// NewUniV3FactoryTransactor creates a new write-only instance of UniV3Factory, bound to a specific deployed contract.
func NewUniV3FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UniV3FactoryTransactor, error) {
	contract, err := bindUniV3Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniV3FactoryTransactor{contract: contract}, nil
}

// NewUniV3FactoryFilterer creates a new log filterer instance of UniV3Factory, bound to a specific deployed contract.
func NewUniV3FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UniV3FactoryFilterer, error) {
	contract, err := bindUniV3Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniV3FactoryFilterer{contract: contract}, nil
}

// bindUniV3Factory binds a generic wrapper to an already deployed contract.
func bindUniV3Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniV3FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniV3Factory *UniV3FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniV3Factory.Contract.UniV3FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniV3Factory *UniV3FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniV3Factory.Contract.UniV3FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniV3Factory *UniV3FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniV3Factory.Contract.UniV3FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniV3Factory *UniV3FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniV3Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniV3Factory *UniV3FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniV3Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniV3Factory *UniV3FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniV3Factory.Contract.contract.Transact(opts, method, params...)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_UniV3Factory *UniV3FactoryCaller) FeeAmountTickSpacing(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniV3Factory.contract.Call(opts, &out, "feeAmountTickSpacing", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_UniV3Factory *UniV3FactorySession) FeeAmountTickSpacing(arg0 *big.Int) (*big.Int, error) {
	return _UniV3Factory.Contract.FeeAmountTickSpacing(&_UniV3Factory.CallOpts, arg0)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_UniV3Factory *UniV3FactoryCallerSession) FeeAmountTickSpacing(arg0 *big.Int) (*big.Int, error) {
	return _UniV3Factory.Contract.FeeAmountTickSpacing(&_UniV3Factory.CallOpts, arg0)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_UniV3Factory *UniV3FactoryCaller) GetPool(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniV3Factory.contract.Call(opts, &out, "getPool", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_UniV3Factory *UniV3FactorySession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _UniV3Factory.Contract.GetPool(&_UniV3Factory.CallOpts, arg0, arg1, arg2)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_UniV3Factory *UniV3FactoryCallerSession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _UniV3Factory.Contract.GetPool(&_UniV3Factory.CallOpts, arg0, arg1, arg2)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniV3Factory *UniV3FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniV3Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniV3Factory *UniV3FactorySession) Owner() (common.Address, error) {
	return _UniV3Factory.Contract.Owner(&_UniV3Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniV3Factory *UniV3FactoryCallerSession) Owner() (common.Address, error) {
	return _UniV3Factory.Contract.Owner(&_UniV3Factory.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_UniV3Factory *UniV3FactoryCaller) Parameters(opts *bind.CallOpts) (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	var out []interface{}
	err := _UniV3Factory.contract.Call(opts, &out, "parameters")

	outstruct := new(struct {
		Factory     common.Address
		Token0      common.Address
		Token1      common.Address
		Fee         *big.Int
		TickSpacing *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Factory = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TickSpacing = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_UniV3Factory *UniV3FactorySession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _UniV3Factory.Contract.Parameters(&_UniV3Factory.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_UniV3Factory *UniV3FactoryCallerSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _UniV3Factory.Contract.Parameters(&_UniV3Factory.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_UniV3Factory *UniV3FactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _UniV3Factory.contract.Transact(opts, "createPool", tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_UniV3Factory *UniV3FactorySession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _UniV3Factory.Contract.CreatePool(&_UniV3Factory.TransactOpts, tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_UniV3Factory *UniV3FactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _UniV3Factory.Contract.CreatePool(&_UniV3Factory.TransactOpts, tokenA, tokenB, fee)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_UniV3Factory *UniV3FactoryTransactor) EnableFeeAmount(opts *bind.TransactOpts, fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _UniV3Factory.contract.Transact(opts, "enableFeeAmount", fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_UniV3Factory *UniV3FactorySession) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _UniV3Factory.Contract.EnableFeeAmount(&_UniV3Factory.TransactOpts, fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_UniV3Factory *UniV3FactoryTransactorSession) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _UniV3Factory.Contract.EnableFeeAmount(&_UniV3Factory.TransactOpts, fee, tickSpacing)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_UniV3Factory *UniV3FactoryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _UniV3Factory.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_UniV3Factory *UniV3FactorySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _UniV3Factory.Contract.SetOwner(&_UniV3Factory.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_UniV3Factory *UniV3FactoryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _UniV3Factory.Contract.SetOwner(&_UniV3Factory.TransactOpts, _owner)
}

// UniV3FactoryFeeAmountEnabledIterator is returned from FilterFeeAmountEnabled and is used to iterate over the raw logs and unpacked data for FeeAmountEnabled events raised by the UniV3Factory contract.
type UniV3FactoryFeeAmountEnabledIterator struct {
	Event *UniV3FactoryFeeAmountEnabled // Event containing the contract specifics and raw log

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
func (it *UniV3FactoryFeeAmountEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniV3FactoryFeeAmountEnabled)
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
		it.Event = new(UniV3FactoryFeeAmountEnabled)
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
func (it *UniV3FactoryFeeAmountEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniV3FactoryFeeAmountEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniV3FactoryFeeAmountEnabled represents a FeeAmountEnabled event raised by the UniV3Factory contract.
type UniV3FactoryFeeAmountEnabled struct {
	Fee         *big.Int
	TickSpacing *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeAmountEnabled is a free log retrieval operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_UniV3Factory *UniV3FactoryFilterer) FilterFeeAmountEnabled(opts *bind.FilterOpts, fee []*big.Int, tickSpacing []*big.Int) (*UniV3FactoryFeeAmountEnabledIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _UniV3Factory.contract.FilterLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return &UniV3FactoryFeeAmountEnabledIterator{contract: _UniV3Factory.contract, event: "FeeAmountEnabled", logs: logs, sub: sub}, nil
}

// WatchFeeAmountEnabled is a free log subscription operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_UniV3Factory *UniV3FactoryFilterer) WatchFeeAmountEnabled(opts *bind.WatchOpts, sink chan<- *UniV3FactoryFeeAmountEnabled, fee []*big.Int, tickSpacing []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _UniV3Factory.contract.WatchLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniV3FactoryFeeAmountEnabled)
				if err := _UniV3Factory.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
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

// ParseFeeAmountEnabled is a log parse operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_UniV3Factory *UniV3FactoryFilterer) ParseFeeAmountEnabled(log types.Log) (*UniV3FactoryFeeAmountEnabled, error) {
	event := new(UniV3FactoryFeeAmountEnabled)
	if err := _UniV3Factory.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniV3FactoryOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the UniV3Factory contract.
type UniV3FactoryOwnerChangedIterator struct {
	Event *UniV3FactoryOwnerChanged // Event containing the contract specifics and raw log

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
func (it *UniV3FactoryOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniV3FactoryOwnerChanged)
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
		it.Event = new(UniV3FactoryOwnerChanged)
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
func (it *UniV3FactoryOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniV3FactoryOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniV3FactoryOwnerChanged represents a OwnerChanged event raised by the UniV3Factory contract.
type UniV3FactoryOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_UniV3Factory *UniV3FactoryFilterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*UniV3FactoryOwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniV3Factory.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniV3FactoryOwnerChangedIterator{contract: _UniV3Factory.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_UniV3Factory *UniV3FactoryFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *UniV3FactoryOwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniV3Factory.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniV3FactoryOwnerChanged)
				if err := _UniV3Factory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_UniV3Factory *UniV3FactoryFilterer) ParseOwnerChanged(log types.Log) (*UniV3FactoryOwnerChanged, error) {
	event := new(UniV3FactoryOwnerChanged)
	if err := _UniV3Factory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniV3FactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the UniV3Factory contract.
type UniV3FactoryPoolCreatedIterator struct {
	Event *UniV3FactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *UniV3FactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniV3FactoryPoolCreated)
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
		it.Event = new(UniV3FactoryPoolCreated)
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
func (it *UniV3FactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniV3FactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniV3FactoryPoolCreated represents a PoolCreated event raised by the UniV3Factory contract.
type UniV3FactoryPoolCreated struct {
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Pool        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_UniV3Factory *UniV3FactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address, fee []*big.Int) (*UniV3FactoryPoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _UniV3Factory.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return &UniV3FactoryPoolCreatedIterator{contract: _UniV3Factory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_UniV3Factory *UniV3FactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *UniV3FactoryPoolCreated, token0 []common.Address, token1 []common.Address, fee []*big.Int) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _UniV3Factory.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniV3FactoryPoolCreated)
				if err := _UniV3Factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_UniV3Factory *UniV3FactoryFilterer) ParsePoolCreated(log types.Log) (*UniV3FactoryPoolCreated, error) {
	event := new(UniV3FactoryPoolCreated)
	if err := _UniV3Factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
