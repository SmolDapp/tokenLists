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

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Lp      common.Address
	Stable  bool
	Token0  common.Address
	Token1  common.Address
	Factory common.Address
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Lp               common.Address
	Symbol           string
	Decimals         uint8
	Stable           bool
	TotalSupply      *big.Int
	Token0           common.Address
	Reserve0         *big.Int
	Claimable0       *big.Int
	Token1           common.Address
	Reserve1         *big.Int
	Claimable1       *big.Int
	Gauge            common.Address
	GaugeTotalSupply *big.Int
	GaugeAlive       bool
	Fee              common.Address
	Bribe            common.Address
	Factory          common.Address
	Emissions        *big.Int
	EmissionsToken   common.Address
	AccountBalance   *big.Int
	AccountEarned    *big.Int
	AccountStaked    *big.Int
	PoolFee          *big.Int
	Token0Fees       *big.Int
	Token1Fees       *big.Int
}

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	TokenAddress   common.Address
	Symbol         string
	Decimals       uint8
	AccountBalance *big.Int
	Listed         bool
}

// Struct3 is an auto generated low-level Go binding around an user-defined struct.
type Struct3 struct {
	Token  common.Address
	Amount *big.Int
}

// Struct5 is an auto generated low-level Go binding around an user-defined struct.
type Struct5 struct {
	VenftId *big.Int
	Lp      common.Address
	Amount  *big.Int
	Token   common.Address
	Fee     common.Address
	Bribe   common.Address
}

// Struct4 is an auto generated low-level Go binding around an user-defined struct.
type Struct4 struct {
	Ts        *big.Int
	Lp        common.Address
	Votes     *big.Int
	Emissions *big.Int
	Bribes    []Struct3
	Fees      []Struct3
}

// VeloSugarV2MetaData contains all meta data concerning the VeloSugarV2 contract.
var VeloSugarV2MetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"},{\"name\":\"_registry\",\"type\":\"address\"},{\"name\":\"_v1_voter\",\"type\":\"address\"},{\"name\":\"_convertor\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"toMigrate\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"stable\",\"type\":\"bool\"},{\"name\":\"total_supply\",\"type\":\"uint256\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"reserve0\",\"type\":\"uint256\"},{\"name\":\"claimable0\",\"type\":\"uint256\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"reserve1\",\"type\":\"uint256\"},{\"name\":\"claimable1\",\"type\":\"uint256\"},{\"name\":\"gauge\",\"type\":\"address\"},{\"name\":\"gauge_total_supply\",\"type\":\"uint256\"},{\"name\":\"gauge_alive\",\"type\":\"bool\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"},{\"name\":\"factory\",\"type\":\"address\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"emissions_token\",\"type\":\"address\"},{\"name\":\"account_balance\",\"type\":\"uint256\"},{\"name\":\"account_earned\",\"type\":\"uint256\"},{\"name\":\"account_staked\",\"type\":\"uint256\"},{\"name\":\"pool_fee\",\"type\":\"uint256\"},{\"name\":\"token0_fees\",\"type\":\"uint256\"},{\"name\":\"token1_fees\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"forSwaps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"stable\",\"type\":\"bool\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"factory\",\"type\":\"address\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokens\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token_address\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"account_balance\",\"type\":\"uint256\"},{\"name\":\"listed\",\"type\":\"bool\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"all\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"stable\",\"type\":\"bool\"},{\"name\":\"total_supply\",\"type\":\"uint256\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"reserve0\",\"type\":\"uint256\"},{\"name\":\"claimable0\",\"type\":\"uint256\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"reserve1\",\"type\":\"uint256\"},{\"name\":\"claimable1\",\"type\":\"uint256\"},{\"name\":\"gauge\",\"type\":\"address\"},{\"name\":\"gauge_total_supply\",\"type\":\"uint256\"},{\"name\":\"gauge_alive\",\"type\":\"bool\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"},{\"name\":\"factory\",\"type\":\"address\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"emissions_token\",\"type\":\"address\"},{\"name\":\"account_balance\",\"type\":\"uint256\"},{\"name\":\"account_earned\",\"type\":\"uint256\"},{\"name\":\"account_staked\",\"type\":\"uint256\"},{\"name\":\"pool_fee\",\"type\":\"uint256\"},{\"name\":\"token0_fees\",\"type\":\"uint256\"},{\"name\":\"token1_fees\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"byIndex\",\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"},{\"name\":\"stable\",\"type\":\"bool\"},{\"name\":\"total_supply\",\"type\":\"uint256\"},{\"name\":\"token0\",\"type\":\"address\"},{\"name\":\"reserve0\",\"type\":\"uint256\"},{\"name\":\"claimable0\",\"type\":\"uint256\"},{\"name\":\"token1\",\"type\":\"address\"},{\"name\":\"reserve1\",\"type\":\"uint256\"},{\"name\":\"claimable1\",\"type\":\"uint256\"},{\"name\":\"gauge\",\"type\":\"address\"},{\"name\":\"gauge_total_supply\",\"type\":\"uint256\"},{\"name\":\"gauge_alive\",\"type\":\"bool\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"},{\"name\":\"factory\",\"type\":\"address\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"emissions_token\",\"type\":\"address\"},{\"name\":\"account_balance\",\"type\":\"uint256\"},{\"name\":\"account_earned\",\"type\":\"uint256\"},{\"name\":\"account_staked\",\"type\":\"uint256\"},{\"name\":\"pool_fee\",\"type\":\"uint256\"},{\"name\":\"token0_fees\",\"type\":\"uint256\"},{\"name\":\"token1_fees\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epochsLatest\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"ts\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"votes\",\"type\":\"uint256\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"bribes\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]},{\"name\":\"fees\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epochsByAddress\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"ts\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"votes\",\"type\":\"uint256\"},{\"name\":\"emissions\",\"type\":\"uint256\"},{\"name\":\"bribes\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]},{\"name\":\"fees\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rewards\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_venft_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"venft_id\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rewardsByAddress\",\"inputs\":[{\"name\":\"_venft_id\",\"type\":\"uint256\"},{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"venft_id\",\"type\":\"uint256\"},{\"name\":\"lp\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"bribe\",\"type\":\"address\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"voter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"convertor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"v1_voter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"v1_factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"v1_token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// VeloSugarV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use VeloSugarV2MetaData.ABI instead.
var VeloSugarV2ABI = VeloSugarV2MetaData.ABI

// VeloSugarV2 is an auto generated Go binding around an Ethereum contract.
type VeloSugarV2 struct {
	VeloSugarV2Caller     // Read-only binding to the contract
	VeloSugarV2Transactor // Write-only binding to the contract
	VeloSugarV2Filterer   // Log filterer for contract events
}

// VeloSugarV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type VeloSugarV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeloSugarV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VeloSugarV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeloSugarV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VeloSugarV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeloSugarV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VeloSugarV2Session struct {
	Contract     *VeloSugarV2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VeloSugarV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VeloSugarV2CallerSession struct {
	Contract *VeloSugarV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VeloSugarV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VeloSugarV2TransactorSession struct {
	Contract     *VeloSugarV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VeloSugarV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type VeloSugarV2Raw struct {
	Contract *VeloSugarV2 // Generic contract binding to access the raw methods on
}

// VeloSugarV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VeloSugarV2CallerRaw struct {
	Contract *VeloSugarV2Caller // Generic read-only contract binding to access the raw methods on
}

// VeloSugarV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VeloSugarV2TransactorRaw struct {
	Contract *VeloSugarV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVeloSugarV2 creates a new instance of VeloSugarV2, bound to a specific deployed contract.
func NewVeloSugarV2(address common.Address, backend bind.ContractBackend) (*VeloSugarV2, error) {
	contract, err := bindVeloSugarV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VeloSugarV2{VeloSugarV2Caller: VeloSugarV2Caller{contract: contract}, VeloSugarV2Transactor: VeloSugarV2Transactor{contract: contract}, VeloSugarV2Filterer: VeloSugarV2Filterer{contract: contract}}, nil
}

// NewVeloSugarV2Caller creates a new read-only instance of VeloSugarV2, bound to a specific deployed contract.
func NewVeloSugarV2Caller(address common.Address, caller bind.ContractCaller) (*VeloSugarV2Caller, error) {
	contract, err := bindVeloSugarV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VeloSugarV2Caller{contract: contract}, nil
}

// NewVeloSugarV2Transactor creates a new write-only instance of VeloSugarV2, bound to a specific deployed contract.
func NewVeloSugarV2Transactor(address common.Address, transactor bind.ContractTransactor) (*VeloSugarV2Transactor, error) {
	contract, err := bindVeloSugarV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VeloSugarV2Transactor{contract: contract}, nil
}

// NewVeloSugarV2Filterer creates a new log filterer instance of VeloSugarV2, bound to a specific deployed contract.
func NewVeloSugarV2Filterer(address common.Address, filterer bind.ContractFilterer) (*VeloSugarV2Filterer, error) {
	contract, err := bindVeloSugarV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VeloSugarV2Filterer{contract: contract}, nil
}

// bindVeloSugarV2 binds a generic wrapper to an already deployed contract.
func bindVeloSugarV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VeloSugarV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VeloSugarV2 *VeloSugarV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VeloSugarV2.Contract.VeloSugarV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VeloSugarV2 *VeloSugarV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeloSugarV2.Contract.VeloSugarV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VeloSugarV2 *VeloSugarV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VeloSugarV2.Contract.VeloSugarV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VeloSugarV2 *VeloSugarV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VeloSugarV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VeloSugarV2 *VeloSugarV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeloSugarV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VeloSugarV2 *VeloSugarV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VeloSugarV2.Contract.contract.Transact(opts, method, params...)
}

// All is a free data retrieval call binding the contract method 0xc0d0bf32.
//
// Solidity: function all(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_VeloSugarV2 *VeloSugarV2Caller) All(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct0, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "all", _limit, _offset, _account)

	if err != nil {
		return *new([]Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct0)).(*[]Struct0)

	return out0, err

}

// All is a free data retrieval call binding the contract method 0xc0d0bf32.
//
// Solidity: function all(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_VeloSugarV2 *VeloSugarV2Session) All(_limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct0, error) {
	return _VeloSugarV2.Contract.All(&_VeloSugarV2.CallOpts, _limit, _offset, _account)
}

// All is a free data retrieval call binding the contract method 0xc0d0bf32.
//
// Solidity: function all(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_VeloSugarV2 *VeloSugarV2CallerSession) All(_limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct0, error) {
	return _VeloSugarV2.Contract.All(&_VeloSugarV2.CallOpts, _limit, _offset, _account)
}

// ByIndex is a free data retrieval call binding the contract method 0xfbb70183.
//
// Solidity: function byIndex(uint256 _index, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_VeloSugarV2 *VeloSugarV2Caller) ByIndex(opts *bind.CallOpts, _index *big.Int, _account common.Address) (Struct0, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "byIndex", _index, _account)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// ByIndex is a free data retrieval call binding the contract method 0xfbb70183.
//
// Solidity: function byIndex(uint256 _index, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_VeloSugarV2 *VeloSugarV2Session) ByIndex(_index *big.Int, _account common.Address) (Struct0, error) {
	return _VeloSugarV2.Contract.ByIndex(&_VeloSugarV2.CallOpts, _index, _account)
}

// ByIndex is a free data retrieval call binding the contract method 0xfbb70183.
//
// Solidity: function byIndex(uint256 _index, address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_VeloSugarV2 *VeloSugarV2CallerSession) ByIndex(_index *big.Int, _account common.Address) (Struct0, error) {
	return _VeloSugarV2.Contract.ByIndex(&_VeloSugarV2.CallOpts, _index, _account)
}

// Convertor is a free data retrieval call binding the contract method 0xb5030306.
//
// Solidity: function convertor() view returns(address)
func (_VeloSugarV2 *VeloSugarV2Caller) Convertor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "convertor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Convertor is a free data retrieval call binding the contract method 0xb5030306.
//
// Solidity: function convertor() view returns(address)
func (_VeloSugarV2 *VeloSugarV2Session) Convertor() (common.Address, error) {
	return _VeloSugarV2.Contract.Convertor(&_VeloSugarV2.CallOpts)
}

// Convertor is a free data retrieval call binding the contract method 0xb5030306.
//
// Solidity: function convertor() view returns(address)
func (_VeloSugarV2 *VeloSugarV2CallerSession) Convertor() (common.Address, error) {
	return _VeloSugarV2.Contract.Convertor(&_VeloSugarV2.CallOpts)
}

// EpochsByAddress is a free data retrieval call binding the contract method 0x8878d06c.
//
// Solidity: function epochsByAddress(uint256 _limit, uint256 _offset, address _address) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugarV2 *VeloSugarV2Caller) EpochsByAddress(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _address common.Address) ([]Struct4, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "epochsByAddress", _limit, _offset, _address)

	if err != nil {
		return *new([]Struct4), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct4)).(*[]Struct4)

	return out0, err

}

// EpochsByAddress is a free data retrieval call binding the contract method 0x8878d06c.
//
// Solidity: function epochsByAddress(uint256 _limit, uint256 _offset, address _address) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugarV2 *VeloSugarV2Session) EpochsByAddress(_limit *big.Int, _offset *big.Int, _address common.Address) ([]Struct4, error) {
	return _VeloSugarV2.Contract.EpochsByAddress(&_VeloSugarV2.CallOpts, _limit, _offset, _address)
}

// EpochsByAddress is a free data retrieval call binding the contract method 0x8878d06c.
//
// Solidity: function epochsByAddress(uint256 _limit, uint256 _offset, address _address) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugarV2 *VeloSugarV2CallerSession) EpochsByAddress(_limit *big.Int, _offset *big.Int, _address common.Address) ([]Struct4, error) {
	return _VeloSugarV2.Contract.EpochsByAddress(&_VeloSugarV2.CallOpts, _limit, _offset, _address)
}

// EpochsLatest is a free data retrieval call binding the contract method 0xd94b9bc6.
//
// Solidity: function epochsLatest(uint256 _limit, uint256 _offset) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugarV2 *VeloSugarV2Caller) EpochsLatest(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int) ([]Struct4, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "epochsLatest", _limit, _offset)

	if err != nil {
		return *new([]Struct4), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct4)).(*[]Struct4)

	return out0, err

}

// EpochsLatest is a free data retrieval call binding the contract method 0xd94b9bc6.
//
// Solidity: function epochsLatest(uint256 _limit, uint256 _offset) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugarV2 *VeloSugarV2Session) EpochsLatest(_limit *big.Int, _offset *big.Int) ([]Struct4, error) {
	return _VeloSugarV2.Contract.EpochsLatest(&_VeloSugarV2.CallOpts, _limit, _offset)
}

// EpochsLatest is a free data retrieval call binding the contract method 0xd94b9bc6.
//
// Solidity: function epochsLatest(uint256 _limit, uint256 _offset) view returns((uint256,address,uint256,uint256,(address,uint256)[],(address,uint256)[])[])
func (_VeloSugarV2 *VeloSugarV2CallerSession) EpochsLatest(_limit *big.Int, _offset *big.Int) ([]Struct4, error) {
	return _VeloSugarV2.Contract.EpochsLatest(&_VeloSugarV2.CallOpts, _limit, _offset)
}

// ForSwaps is a free data retrieval call binding the contract method 0xcc259a0e.
//
// Solidity: function forSwaps() view returns((address,bool,address,address,address)[])
func (_VeloSugarV2 *VeloSugarV2Caller) ForSwaps(opts *bind.CallOpts) ([]Struct1, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "forSwaps")

	if err != nil {
		return *new([]Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct1)).(*[]Struct1)

	return out0, err

}

// ForSwaps is a free data retrieval call binding the contract method 0xcc259a0e.
//
// Solidity: function forSwaps() view returns((address,bool,address,address,address)[])
func (_VeloSugarV2 *VeloSugarV2Session) ForSwaps() ([]Struct1, error) {
	return _VeloSugarV2.Contract.ForSwaps(&_VeloSugarV2.CallOpts)
}

// ForSwaps is a free data retrieval call binding the contract method 0xcc259a0e.
//
// Solidity: function forSwaps() view returns((address,bool,address,address,address)[])
func (_VeloSugarV2 *VeloSugarV2CallerSession) ForSwaps() ([]Struct1, error) {
	return _VeloSugarV2.Contract.ForSwaps(&_VeloSugarV2.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_VeloSugarV2 *VeloSugarV2Caller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_VeloSugarV2 *VeloSugarV2Session) Registry() (common.Address, error) {
	return _VeloSugarV2.Contract.Registry(&_VeloSugarV2.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_VeloSugarV2 *VeloSugarV2CallerSession) Registry() (common.Address, error) {
	return _VeloSugarV2.Contract.Registry(&_VeloSugarV2.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0xa9c57fee.
//
// Solidity: function rewards(uint256 _limit, uint256 _offset, uint256 _venft_id) view returns((uint256,address,uint256,address,address,address)[])
func (_VeloSugarV2 *VeloSugarV2Caller) Rewards(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _venft_id *big.Int) ([]Struct5, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "rewards", _limit, _offset, _venft_id)

	if err != nil {
		return *new([]Struct5), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct5)).(*[]Struct5)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0xa9c57fee.
//
// Solidity: function rewards(uint256 _limit, uint256 _offset, uint256 _venft_id) view returns((uint256,address,uint256,address,address,address)[])
func (_VeloSugarV2 *VeloSugarV2Session) Rewards(_limit *big.Int, _offset *big.Int, _venft_id *big.Int) ([]Struct5, error) {
	return _VeloSugarV2.Contract.Rewards(&_VeloSugarV2.CallOpts, _limit, _offset, _venft_id)
}

// Rewards is a free data retrieval call binding the contract method 0xa9c57fee.
//
// Solidity: function rewards(uint256 _limit, uint256 _offset, uint256 _venft_id) view returns((uint256,address,uint256,address,address,address)[])
func (_VeloSugarV2 *VeloSugarV2CallerSession) Rewards(_limit *big.Int, _offset *big.Int, _venft_id *big.Int) ([]Struct5, error) {
	return _VeloSugarV2.Contract.Rewards(&_VeloSugarV2.CallOpts, _limit, _offset, _venft_id)
}

// RewardsByAddress is a free data retrieval call binding the contract method 0xcd824fb4.
//
// Solidity: function rewardsByAddress(uint256 _venft_id, address _pool) view returns((uint256,address,uint256,address,address,address)[])
func (_VeloSugarV2 *VeloSugarV2Caller) RewardsByAddress(opts *bind.CallOpts, _venft_id *big.Int, _pool common.Address) ([]Struct5, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "rewardsByAddress", _venft_id, _pool)

	if err != nil {
		return *new([]Struct5), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct5)).(*[]Struct5)

	return out0, err

}

// RewardsByAddress is a free data retrieval call binding the contract method 0xcd824fb4.
//
// Solidity: function rewardsByAddress(uint256 _venft_id, address _pool) view returns((uint256,address,uint256,address,address,address)[])
func (_VeloSugarV2 *VeloSugarV2Session) RewardsByAddress(_venft_id *big.Int, _pool common.Address) ([]Struct5, error) {
	return _VeloSugarV2.Contract.RewardsByAddress(&_VeloSugarV2.CallOpts, _venft_id, _pool)
}

// RewardsByAddress is a free data retrieval call binding the contract method 0xcd824fb4.
//
// Solidity: function rewardsByAddress(uint256 _venft_id, address _pool) view returns((uint256,address,uint256,address,address,address)[])
func (_VeloSugarV2 *VeloSugarV2CallerSession) RewardsByAddress(_venft_id *big.Int, _pool common.Address) ([]Struct5, error) {
	return _VeloSugarV2.Contract.RewardsByAddress(&_VeloSugarV2.CallOpts, _venft_id, _pool)
}

// ToMigrate is a free data retrieval call binding the contract method 0xa1aaf452.
//
// Solidity: function toMigrate(address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_VeloSugarV2 *VeloSugarV2Caller) ToMigrate(opts *bind.CallOpts, _account common.Address) ([]Struct0, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "toMigrate", _account)

	if err != nil {
		return *new([]Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct0)).(*[]Struct0)

	return out0, err

}

// ToMigrate is a free data retrieval call binding the contract method 0xa1aaf452.
//
// Solidity: function toMigrate(address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_VeloSugarV2 *VeloSugarV2Session) ToMigrate(_account common.Address) ([]Struct0, error) {
	return _VeloSugarV2.Contract.ToMigrate(&_VeloSugarV2.CallOpts, _account)
}

// ToMigrate is a free data retrieval call binding the contract method 0xa1aaf452.
//
// Solidity: function toMigrate(address _account) view returns((address,string,uint8,bool,uint256,address,uint256,uint256,address,uint256,uint256,address,uint256,bool,address,address,address,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_VeloSugarV2 *VeloSugarV2CallerSession) ToMigrate(_account common.Address) ([]Struct0, error) {
	return _VeloSugarV2.Contract.ToMigrate(&_VeloSugarV2.CallOpts, _account)
}

// Tokens is a free data retrieval call binding the contract method 0x5cc33187.
//
// Solidity: function tokens(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,uint256,bool)[])
func (_VeloSugarV2 *VeloSugarV2Caller) Tokens(opts *bind.CallOpts, _limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct2, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "tokens", _limit, _offset, _account)

	if err != nil {
		return *new([]Struct2), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct2)).(*[]Struct2)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x5cc33187.
//
// Solidity: function tokens(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,uint256,bool)[])
func (_VeloSugarV2 *VeloSugarV2Session) Tokens(_limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct2, error) {
	return _VeloSugarV2.Contract.Tokens(&_VeloSugarV2.CallOpts, _limit, _offset, _account)
}

// Tokens is a free data retrieval call binding the contract method 0x5cc33187.
//
// Solidity: function tokens(uint256 _limit, uint256 _offset, address _account) view returns((address,string,uint8,uint256,bool)[])
func (_VeloSugarV2 *VeloSugarV2CallerSession) Tokens(_limit *big.Int, _offset *big.Int, _account common.Address) ([]Struct2, error) {
	return _VeloSugarV2.Contract.Tokens(&_VeloSugarV2.CallOpts, _limit, _offset, _account)
}

// V1Factory is a free data retrieval call binding the contract method 0x52d51859.
//
// Solidity: function v1_factory() view returns(address)
func (_VeloSugarV2 *VeloSugarV2Caller) V1Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "v1_factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// V1Factory is a free data retrieval call binding the contract method 0x52d51859.
//
// Solidity: function v1_factory() view returns(address)
func (_VeloSugarV2 *VeloSugarV2Session) V1Factory() (common.Address, error) {
	return _VeloSugarV2.Contract.V1Factory(&_VeloSugarV2.CallOpts)
}

// V1Factory is a free data retrieval call binding the contract method 0x52d51859.
//
// Solidity: function v1_factory() view returns(address)
func (_VeloSugarV2 *VeloSugarV2CallerSession) V1Factory() (common.Address, error) {
	return _VeloSugarV2.Contract.V1Factory(&_VeloSugarV2.CallOpts)
}

// V1Token is a free data retrieval call binding the contract method 0x9ea10138.
//
// Solidity: function v1_token() view returns(address)
func (_VeloSugarV2 *VeloSugarV2Caller) V1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "v1_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// V1Token is a free data retrieval call binding the contract method 0x9ea10138.
//
// Solidity: function v1_token() view returns(address)
func (_VeloSugarV2 *VeloSugarV2Session) V1Token() (common.Address, error) {
	return _VeloSugarV2.Contract.V1Token(&_VeloSugarV2.CallOpts)
}

// V1Token is a free data retrieval call binding the contract method 0x9ea10138.
//
// Solidity: function v1_token() view returns(address)
func (_VeloSugarV2 *VeloSugarV2CallerSession) V1Token() (common.Address, error) {
	return _VeloSugarV2.Contract.V1Token(&_VeloSugarV2.CallOpts)
}

// V1Voter is a free data retrieval call binding the contract method 0x2217f996.
//
// Solidity: function v1_voter() view returns(address)
func (_VeloSugarV2 *VeloSugarV2Caller) V1Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "v1_voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// V1Voter is a free data retrieval call binding the contract method 0x2217f996.
//
// Solidity: function v1_voter() view returns(address)
func (_VeloSugarV2 *VeloSugarV2Session) V1Voter() (common.Address, error) {
	return _VeloSugarV2.Contract.V1Voter(&_VeloSugarV2.CallOpts)
}

// V1Voter is a free data retrieval call binding the contract method 0x2217f996.
//
// Solidity: function v1_voter() view returns(address)
func (_VeloSugarV2 *VeloSugarV2CallerSession) V1Voter() (common.Address, error) {
	return _VeloSugarV2.Contract.V1Voter(&_VeloSugarV2.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_VeloSugarV2 *VeloSugarV2Caller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VeloSugarV2.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_VeloSugarV2 *VeloSugarV2Session) Voter() (common.Address, error) {
	return _VeloSugarV2.Contract.Voter(&_VeloSugarV2.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_VeloSugarV2 *VeloSugarV2CallerSession) Voter() (common.Address, error) {
	return _VeloSugarV2.Contract.Voter(&_VeloSugarV2.CallOpts)
}
