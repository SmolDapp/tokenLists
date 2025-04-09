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

// ITokenRegistryToken is an auto generated low-level Go binding around an user-defined struct.
type ITokenRegistryToken struct {
	ContractAddress common.Address
	Name            string
	Symbol          string
	Decimals        uint8
	LogoURI         string
	Status          uint8
	Metadata        []MetadataValue
}

// MetadataInput is an auto generated low-level Go binding around an user-defined struct.
type MetadataInput struct {
	Field string
	Value string
}

// MetadataValue is an auto generated low-level Go binding around an user-defined struct.
type MetadataValue struct {
	Field    string
	Value    string
	IsActive bool
}

// TokenRegistryMetaData contains all meta data concerning the TokenRegistry contract.
var TokenRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokentroller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenMetadata\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"TokenApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"TokenRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"TokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTokentroller\",\"type\":\"address\"}],\"name\":\"TokentrollerUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structMetadataInput[]\",\"name\":\"metadata\",\"type\":\"tuple[]\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"approveToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getToken\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"logoURI\",\"type\":\"string\"},{\"internalType\":\"enumTokenStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structMetadataValue[]\",\"name\":\"metadata\",\"type\":\"tuple[]\"}],\"internalType\":\"structITokenRegistry.Token\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"includeMetadata\",\"type\":\"bool\"}],\"name\":\"getToken\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"logoURI\",\"type\":\"string\"},{\"internalType\":\"enumTokenStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structMetadataValue[]\",\"name\":\"metadata\",\"type\":\"tuple[]\"}],\"internalType\":\"structITokenRegistry.Token\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"approved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rejected\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"contractAddresses\",\"type\":\"address[]\"}],\"name\":\"getTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"logoURI\",\"type\":\"string\"},{\"internalType\":\"enumTokenStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structMetadataValue[]\",\"name\":\"metadata\",\"type\":\"tuple[]\"}],\"internalType\":\"structITokenRegistry.Token[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"contractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"includeMetadata\",\"type\":\"bool\"}],\"name\":\"getTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"logoURI\",\"type\":\"string\"},{\"internalType\":\"enumTokenStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structMetadataValue[]\",\"name\":\"metadata\",\"type\":\"tuple[]\"}],\"internalType\":\"structITokenRegistry.Token[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"enumTokenStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"includeMetadata\",\"type\":\"bool\"}],\"name\":\"listTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"logoURI\",\"type\":\"string\"},{\"internalType\":\"enumTokenStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structMetadataValue[]\",\"name\":\"metadata\",\"type\":\"tuple[]\"}],\"internalType\":\"structITokenRegistry.Token[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"enumTokenStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"listTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"logoURI\",\"type\":\"string\"},{\"internalType\":\"enumTokenStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structMetadataValue[]\",\"name\":\"metadata\",\"type\":\"tuple[]\"}],\"internalType\":\"structITokenRegistry.Token[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"rejectToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMetadata\",\"outputs\":[{\"internalType\":\"contractITokenMetadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"tokenStatus\",\"outputs\":[{\"internalType\":\"enumTokenStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokentroller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTokentroller\",\"type\":\"address\"}],\"name\":\"updateTokentroller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenRegistryMetaData.ABI instead.
var TokenRegistryABI = TokenRegistryMetaData.ABI

// TokenRegistry is an auto generated Go binding around an Ethereum contract.
type TokenRegistry struct {
	TokenRegistryCaller     // Read-only binding to the contract
	TokenRegistryTransactor // Write-only binding to the contract
	TokenRegistryFilterer   // Log filterer for contract events
}

// TokenRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenRegistrySession struct {
	Contract     *TokenRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenRegistryCallerSession struct {
	Contract *TokenRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenRegistryTransactorSession struct {
	Contract     *TokenRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRegistryRaw struct {
	Contract *TokenRegistry // Generic contract binding to access the raw methods on
}

// TokenRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenRegistryCallerRaw struct {
	Contract *TokenRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenRegistryTransactorRaw struct {
	Contract *TokenRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenRegistry creates a new instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistry(address common.Address, backend bind.ContractBackend) (*TokenRegistry, error) {
	contract, err := bindTokenRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRegistry{TokenRegistryCaller: TokenRegistryCaller{contract: contract}, TokenRegistryTransactor: TokenRegistryTransactor{contract: contract}, TokenRegistryFilterer: TokenRegistryFilterer{contract: contract}}, nil
}

// NewTokenRegistryCaller creates a new read-only instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryCaller(address common.Address, caller bind.ContractCaller) (*TokenRegistryCaller, error) {
	contract, err := bindTokenRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryCaller{contract: contract}, nil
}

// NewTokenRegistryTransactor creates a new write-only instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRegistryTransactor, error) {
	contract, err := bindTokenRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTransactor{contract: contract}, nil
}

// NewTokenRegistryFilterer creates a new log filterer instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenRegistryFilterer, error) {
	contract, err := bindTokenRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryFilterer{contract: contract}, nil
}

// bindTokenRegistry binds a generic wrapper to an already deployed contract.
func bindTokenRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRegistry *TokenRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRegistry.Contract.TokenRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRegistry *TokenRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TokenRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRegistry *TokenRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TokenRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRegistry *TokenRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRegistry *TokenRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRegistry *TokenRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address contractAddress) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[]))
func (_TokenRegistry *TokenRegistryCaller) GetToken(opts *bind.CallOpts, contractAddress common.Address) (ITokenRegistryToken, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "getToken", contractAddress)

	if err != nil {
		return *new(ITokenRegistryToken), err
	}

	out0 := *abi.ConvertType(out[0], new(ITokenRegistryToken)).(*ITokenRegistryToken)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address contractAddress) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[]))
func (_TokenRegistry *TokenRegistrySession) GetToken(contractAddress common.Address) (ITokenRegistryToken, error) {
	return _TokenRegistry.Contract.GetToken(&_TokenRegistry.CallOpts, contractAddress)
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address contractAddress) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[]))
func (_TokenRegistry *TokenRegistryCallerSession) GetToken(contractAddress common.Address) (ITokenRegistryToken, error) {
	return _TokenRegistry.Contract.GetToken(&_TokenRegistry.CallOpts, contractAddress)
}

// GetToken0 is a free data retrieval call binding the contract method 0xf80eb03c.
//
// Solidity: function getToken(address contractAddress, bool includeMetadata) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[]))
func (_TokenRegistry *TokenRegistryCaller) GetToken0(opts *bind.CallOpts, contractAddress common.Address, includeMetadata bool) (ITokenRegistryToken, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "getToken0", contractAddress, includeMetadata)

	if err != nil {
		return *new(ITokenRegistryToken), err
	}

	out0 := *abi.ConvertType(out[0], new(ITokenRegistryToken)).(*ITokenRegistryToken)

	return out0, err

}

// GetToken0 is a free data retrieval call binding the contract method 0xf80eb03c.
//
// Solidity: function getToken(address contractAddress, bool includeMetadata) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[]))
func (_TokenRegistry *TokenRegistrySession) GetToken0(contractAddress common.Address, includeMetadata bool) (ITokenRegistryToken, error) {
	return _TokenRegistry.Contract.GetToken0(&_TokenRegistry.CallOpts, contractAddress, includeMetadata)
}

// GetToken0 is a free data retrieval call binding the contract method 0xf80eb03c.
//
// Solidity: function getToken(address contractAddress, bool includeMetadata) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[]))
func (_TokenRegistry *TokenRegistryCallerSession) GetToken0(contractAddress common.Address, includeMetadata bool) (ITokenRegistryToken, error) {
	return _TokenRegistry.Contract.GetToken0(&_TokenRegistry.CallOpts, contractAddress, includeMetadata)
}

// GetTokenCounts is a free data retrieval call binding the contract method 0xcee19474.
//
// Solidity: function getTokenCounts() view returns(uint256 pending, uint256 approved, uint256 rejected)
func (_TokenRegistry *TokenRegistryCaller) GetTokenCounts(opts *bind.CallOpts) (struct {
	Pending  *big.Int
	Approved *big.Int
	Rejected *big.Int
}, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "getTokenCounts")

	outstruct := new(struct {
		Pending  *big.Int
		Approved *big.Int
		Rejected *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pending = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Approved = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Rejected = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenCounts is a free data retrieval call binding the contract method 0xcee19474.
//
// Solidity: function getTokenCounts() view returns(uint256 pending, uint256 approved, uint256 rejected)
func (_TokenRegistry *TokenRegistrySession) GetTokenCounts() (struct {
	Pending  *big.Int
	Approved *big.Int
	Rejected *big.Int
}, error) {
	return _TokenRegistry.Contract.GetTokenCounts(&_TokenRegistry.CallOpts)
}

// GetTokenCounts is a free data retrieval call binding the contract method 0xcee19474.
//
// Solidity: function getTokenCounts() view returns(uint256 pending, uint256 approved, uint256 rejected)
func (_TokenRegistry *TokenRegistryCallerSession) GetTokenCounts() (struct {
	Pending  *big.Int
	Approved *big.Int
	Rejected *big.Int
}, error) {
	return _TokenRegistry.Contract.GetTokenCounts(&_TokenRegistry.CallOpts)
}

// GetTokens is a free data retrieval call binding the contract method 0x273a94aa.
//
// Solidity: function getTokens(address[] contractAddresses) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[])[])
func (_TokenRegistry *TokenRegistryCaller) GetTokens(opts *bind.CallOpts, contractAddresses []common.Address) ([]ITokenRegistryToken, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "getTokens", contractAddresses)

	if err != nil {
		return *new([]ITokenRegistryToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITokenRegistryToken)).(*[]ITokenRegistryToken)

	return out0, err

}

// GetTokens is a free data retrieval call binding the contract method 0x273a94aa.
//
// Solidity: function getTokens(address[] contractAddresses) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[])[])
func (_TokenRegistry *TokenRegistrySession) GetTokens(contractAddresses []common.Address) ([]ITokenRegistryToken, error) {
	return _TokenRegistry.Contract.GetTokens(&_TokenRegistry.CallOpts, contractAddresses)
}

// GetTokens is a free data retrieval call binding the contract method 0x273a94aa.
//
// Solidity: function getTokens(address[] contractAddresses) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[])[])
func (_TokenRegistry *TokenRegistryCallerSession) GetTokens(contractAddresses []common.Address) ([]ITokenRegistryToken, error) {
	return _TokenRegistry.Contract.GetTokens(&_TokenRegistry.CallOpts, contractAddresses)
}

// GetTokens0 is a free data retrieval call binding the contract method 0xb8691614.
//
// Solidity: function getTokens(address[] contractAddresses, bool includeMetadata) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[])[])
func (_TokenRegistry *TokenRegistryCaller) GetTokens0(opts *bind.CallOpts, contractAddresses []common.Address, includeMetadata bool) ([]ITokenRegistryToken, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "getTokens0", contractAddresses, includeMetadata)

	if err != nil {
		return *new([]ITokenRegistryToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITokenRegistryToken)).(*[]ITokenRegistryToken)

	return out0, err

}

// GetTokens0 is a free data retrieval call binding the contract method 0xb8691614.
//
// Solidity: function getTokens(address[] contractAddresses, bool includeMetadata) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[])[])
func (_TokenRegistry *TokenRegistrySession) GetTokens0(contractAddresses []common.Address, includeMetadata bool) ([]ITokenRegistryToken, error) {
	return _TokenRegistry.Contract.GetTokens0(&_TokenRegistry.CallOpts, contractAddresses, includeMetadata)
}

// GetTokens0 is a free data retrieval call binding the contract method 0xb8691614.
//
// Solidity: function getTokens(address[] contractAddresses, bool includeMetadata) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[])[])
func (_TokenRegistry *TokenRegistryCallerSession) GetTokens0(contractAddresses []common.Address, includeMetadata bool) ([]ITokenRegistryToken, error) {
	return _TokenRegistry.Contract.GetTokens0(&_TokenRegistry.CallOpts, contractAddresses, includeMetadata)
}

// ListTokens is a free data retrieval call binding the contract method 0xa9941816.
//
// Solidity: function listTokens(uint256 offset, uint256 limit, uint8 status, bool includeMetadata) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[])[] tokens, uint256 total)
func (_TokenRegistry *TokenRegistryCaller) ListTokens(opts *bind.CallOpts, offset *big.Int, limit *big.Int, status uint8, includeMetadata bool) (struct {
	Tokens []ITokenRegistryToken
	Total  *big.Int
}, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "listTokens", offset, limit, status, includeMetadata)

	outstruct := new(struct {
		Tokens []ITokenRegistryToken
		Total  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]ITokenRegistryToken)).(*[]ITokenRegistryToken)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ListTokens is a free data retrieval call binding the contract method 0xa9941816.
//
// Solidity: function listTokens(uint256 offset, uint256 limit, uint8 status, bool includeMetadata) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[])[] tokens, uint256 total)
func (_TokenRegistry *TokenRegistrySession) ListTokens(offset *big.Int, limit *big.Int, status uint8, includeMetadata bool) (struct {
	Tokens []ITokenRegistryToken
	Total  *big.Int
}, error) {
	return _TokenRegistry.Contract.ListTokens(&_TokenRegistry.CallOpts, offset, limit, status, includeMetadata)
}

// ListTokens is a free data retrieval call binding the contract method 0xa9941816.
//
// Solidity: function listTokens(uint256 offset, uint256 limit, uint8 status, bool includeMetadata) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[])[] tokens, uint256 total)
func (_TokenRegistry *TokenRegistryCallerSession) ListTokens(offset *big.Int, limit *big.Int, status uint8, includeMetadata bool) (struct {
	Tokens []ITokenRegistryToken
	Total  *big.Int
}, error) {
	return _TokenRegistry.Contract.ListTokens(&_TokenRegistry.CallOpts, offset, limit, status, includeMetadata)
}

// ListTokens0 is a free data retrieval call binding the contract method 0xe5c4be1d.
//
// Solidity: function listTokens(uint256 offset, uint256 limit, uint8 status) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[])[] tokens, uint256 total)
func (_TokenRegistry *TokenRegistryCaller) ListTokens0(opts *bind.CallOpts, offset *big.Int, limit *big.Int, status uint8) (struct {
	Tokens []ITokenRegistryToken
	Total  *big.Int
}, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "listTokens0", offset, limit, status)

	outstruct := new(struct {
		Tokens []ITokenRegistryToken
		Total  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]ITokenRegistryToken)).(*[]ITokenRegistryToken)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ListTokens0 is a free data retrieval call binding the contract method 0xe5c4be1d.
//
// Solidity: function listTokens(uint256 offset, uint256 limit, uint8 status) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[])[] tokens, uint256 total)
func (_TokenRegistry *TokenRegistrySession) ListTokens0(offset *big.Int, limit *big.Int, status uint8) (struct {
	Tokens []ITokenRegistryToken
	Total  *big.Int
}, error) {
	return _TokenRegistry.Contract.ListTokens0(&_TokenRegistry.CallOpts, offset, limit, status)
}

// ListTokens0 is a free data retrieval call binding the contract method 0xe5c4be1d.
//
// Solidity: function listTokens(uint256 offset, uint256 limit, uint8 status) view returns((address,string,string,uint8,string,uint8,(string,string,bool)[])[] tokens, uint256 total)
func (_TokenRegistry *TokenRegistryCallerSession) ListTokens0(offset *big.Int, limit *big.Int, status uint8) (struct {
	Tokens []ITokenRegistryToken
	Total  *big.Int
}, error) {
	return _TokenRegistry.Contract.ListTokens0(&_TokenRegistry.CallOpts, offset, limit, status)
}

// TokenMetadata is a free data retrieval call binding the contract method 0xf5b2cad7.
//
// Solidity: function tokenMetadata() view returns(address)
func (_TokenRegistry *TokenRegistryCaller) TokenMetadata(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "tokenMetadata")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMetadata is a free data retrieval call binding the contract method 0xf5b2cad7.
//
// Solidity: function tokenMetadata() view returns(address)
func (_TokenRegistry *TokenRegistrySession) TokenMetadata() (common.Address, error) {
	return _TokenRegistry.Contract.TokenMetadata(&_TokenRegistry.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0xf5b2cad7.
//
// Solidity: function tokenMetadata() view returns(address)
func (_TokenRegistry *TokenRegistryCallerSession) TokenMetadata() (common.Address, error) {
	return _TokenRegistry.Contract.TokenMetadata(&_TokenRegistry.CallOpts)
}

// TokenStatus is a free data retrieval call binding the contract method 0x0acac942.
//
// Solidity: function tokenStatus(address contractAddress) view returns(uint8)
func (_TokenRegistry *TokenRegistryCaller) TokenStatus(opts *bind.CallOpts, contractAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "tokenStatus", contractAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenStatus is a free data retrieval call binding the contract method 0x0acac942.
//
// Solidity: function tokenStatus(address contractAddress) view returns(uint8)
func (_TokenRegistry *TokenRegistrySession) TokenStatus(contractAddress common.Address) (uint8, error) {
	return _TokenRegistry.Contract.TokenStatus(&_TokenRegistry.CallOpts, contractAddress)
}

// TokenStatus is a free data retrieval call binding the contract method 0x0acac942.
//
// Solidity: function tokenStatus(address contractAddress) view returns(uint8)
func (_TokenRegistry *TokenRegistryCallerSession) TokenStatus(contractAddress common.Address) (uint8, error) {
	return _TokenRegistry.Contract.TokenStatus(&_TokenRegistry.CallOpts, contractAddress)
}

// Tokentroller is a free data retrieval call binding the contract method 0x53403983.
//
// Solidity: function tokentroller() view returns(address)
func (_TokenRegistry *TokenRegistryCaller) Tokentroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "tokentroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokentroller is a free data retrieval call binding the contract method 0x53403983.
//
// Solidity: function tokentroller() view returns(address)
func (_TokenRegistry *TokenRegistrySession) Tokentroller() (common.Address, error) {
	return _TokenRegistry.Contract.Tokentroller(&_TokenRegistry.CallOpts)
}

// Tokentroller is a free data retrieval call binding the contract method 0x53403983.
//
// Solidity: function tokentroller() view returns(address)
func (_TokenRegistry *TokenRegistryCallerSession) Tokentroller() (common.Address, error) {
	return _TokenRegistry.Contract.Tokentroller(&_TokenRegistry.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0xeeddbd97.
//
// Solidity: function addToken(address contractAddress, (string,string)[] metadata) returns()
func (_TokenRegistry *TokenRegistryTransactor) AddToken(opts *bind.TransactOpts, contractAddress common.Address, metadata []MetadataInput) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "addToken", contractAddress, metadata)
}

// AddToken is a paid mutator transaction binding the contract method 0xeeddbd97.
//
// Solidity: function addToken(address contractAddress, (string,string)[] metadata) returns()
func (_TokenRegistry *TokenRegistrySession) AddToken(contractAddress common.Address, metadata []MetadataInput) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddToken(&_TokenRegistry.TransactOpts, contractAddress, metadata)
}

// AddToken is a paid mutator transaction binding the contract method 0xeeddbd97.
//
// Solidity: function addToken(address contractAddress, (string,string)[] metadata) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) AddToken(contractAddress common.Address, metadata []MetadataInput) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddToken(&_TokenRegistry.TransactOpts, contractAddress, metadata)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x80b2edd8.
//
// Solidity: function approveToken(address contractAddress) returns()
func (_TokenRegistry *TokenRegistryTransactor) ApproveToken(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "approveToken", contractAddress)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x80b2edd8.
//
// Solidity: function approveToken(address contractAddress) returns()
func (_TokenRegistry *TokenRegistrySession) ApproveToken(contractAddress common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.ApproveToken(&_TokenRegistry.TransactOpts, contractAddress)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x80b2edd8.
//
// Solidity: function approveToken(address contractAddress) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) ApproveToken(contractAddress common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.ApproveToken(&_TokenRegistry.TransactOpts, contractAddress)
}

// RejectToken is a paid mutator transaction binding the contract method 0x93c56d54.
//
// Solidity: function rejectToken(address contractAddress, string reason) returns()
func (_TokenRegistry *TokenRegistryTransactor) RejectToken(opts *bind.TransactOpts, contractAddress common.Address, reason string) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "rejectToken", contractAddress, reason)
}

// RejectToken is a paid mutator transaction binding the contract method 0x93c56d54.
//
// Solidity: function rejectToken(address contractAddress, string reason) returns()
func (_TokenRegistry *TokenRegistrySession) RejectToken(contractAddress common.Address, reason string) (*types.Transaction, error) {
	return _TokenRegistry.Contract.RejectToken(&_TokenRegistry.TransactOpts, contractAddress, reason)
}

// RejectToken is a paid mutator transaction binding the contract method 0x93c56d54.
//
// Solidity: function rejectToken(address contractAddress, string reason) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) RejectToken(contractAddress common.Address, reason string) (*types.Transaction, error) {
	return _TokenRegistry.Contract.RejectToken(&_TokenRegistry.TransactOpts, contractAddress, reason)
}

// UpdateTokentroller is a paid mutator transaction binding the contract method 0xf177a13a.
//
// Solidity: function updateTokentroller(address newTokentroller) returns()
func (_TokenRegistry *TokenRegistryTransactor) UpdateTokentroller(opts *bind.TransactOpts, newTokentroller common.Address) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "updateTokentroller", newTokentroller)
}

// UpdateTokentroller is a paid mutator transaction binding the contract method 0xf177a13a.
//
// Solidity: function updateTokentroller(address newTokentroller) returns()
func (_TokenRegistry *TokenRegistrySession) UpdateTokentroller(newTokentroller common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.UpdateTokentroller(&_TokenRegistry.TransactOpts, newTokentroller)
}

// UpdateTokentroller is a paid mutator transaction binding the contract method 0xf177a13a.
//
// Solidity: function updateTokentroller(address newTokentroller) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) UpdateTokentroller(newTokentroller common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.UpdateTokentroller(&_TokenRegistry.TransactOpts, newTokentroller)
}

// TokenRegistryTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the TokenRegistry contract.
type TokenRegistryTokenAddedIterator struct {
	Event *TokenRegistryTokenAdded // Event containing the contract specifics and raw log

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
func (it *TokenRegistryTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryTokenAdded)
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
		it.Event = new(TokenRegistryTokenAdded)
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
func (it *TokenRegistryTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryTokenAdded represents a TokenAdded event raised by the TokenRegistry contract.
type TokenRegistryTokenAdded struct {
	ContractAddress common.Address
	Submitter       common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0xdffbd9ded1c09446f09377de547142dcce7dc541c8b0b028142b1eba7026b9e7.
//
// Solidity: event TokenAdded(address indexed contractAddress, address indexed submitter)
func (_TokenRegistry *TokenRegistryFilterer) FilterTokenAdded(opts *bind.FilterOpts, contractAddress []common.Address, submitter []common.Address) (*TokenRegistryTokenAddedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "TokenAdded", contractAddressRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTokenAddedIterator{contract: _TokenRegistry.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0xdffbd9ded1c09446f09377de547142dcce7dc541c8b0b028142b1eba7026b9e7.
//
// Solidity: event TokenAdded(address indexed contractAddress, address indexed submitter)
func (_TokenRegistry *TokenRegistryFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *TokenRegistryTokenAdded, contractAddress []common.Address, submitter []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "TokenAdded", contractAddressRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryTokenAdded)
				if err := _TokenRegistry.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0xdffbd9ded1c09446f09377de547142dcce7dc541c8b0b028142b1eba7026b9e7.
//
// Solidity: event TokenAdded(address indexed contractAddress, address indexed submitter)
func (_TokenRegistry *TokenRegistryFilterer) ParseTokenAdded(log types.Log) (*TokenRegistryTokenAdded, error) {
	event := new(TokenRegistryTokenAdded)
	if err := _TokenRegistry.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRegistryTokenApprovedIterator is returned from FilterTokenApproved and is used to iterate over the raw logs and unpacked data for TokenApproved events raised by the TokenRegistry contract.
type TokenRegistryTokenApprovedIterator struct {
	Event *TokenRegistryTokenApproved // Event containing the contract specifics and raw log

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
func (it *TokenRegistryTokenApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryTokenApproved)
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
		it.Event = new(TokenRegistryTokenApproved)
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
func (it *TokenRegistryTokenApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryTokenApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryTokenApproved represents a TokenApproved event raised by the TokenRegistry contract.
type TokenRegistryTokenApproved struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenApproved is a free log retrieval operation binding the contract event 0x5f5916d70d5479c1795a9d461360dfa5c673bc37904c8ab4fcbdc970b9e90f3d.
//
// Solidity: event TokenApproved(address indexed contractAddress)
func (_TokenRegistry *TokenRegistryFilterer) FilterTokenApproved(opts *bind.FilterOpts, contractAddress []common.Address) (*TokenRegistryTokenApprovedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "TokenApproved", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTokenApprovedIterator{contract: _TokenRegistry.contract, event: "TokenApproved", logs: logs, sub: sub}, nil
}

// WatchTokenApproved is a free log subscription operation binding the contract event 0x5f5916d70d5479c1795a9d461360dfa5c673bc37904c8ab4fcbdc970b9e90f3d.
//
// Solidity: event TokenApproved(address indexed contractAddress)
func (_TokenRegistry *TokenRegistryFilterer) WatchTokenApproved(opts *bind.WatchOpts, sink chan<- *TokenRegistryTokenApproved, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "TokenApproved", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryTokenApproved)
				if err := _TokenRegistry.contract.UnpackLog(event, "TokenApproved", log); err != nil {
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

// ParseTokenApproved is a log parse operation binding the contract event 0x5f5916d70d5479c1795a9d461360dfa5c673bc37904c8ab4fcbdc970b9e90f3d.
//
// Solidity: event TokenApproved(address indexed contractAddress)
func (_TokenRegistry *TokenRegistryFilterer) ParseTokenApproved(log types.Log) (*TokenRegistryTokenApproved, error) {
	event := new(TokenRegistryTokenApproved)
	if err := _TokenRegistry.contract.UnpackLog(event, "TokenApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRegistryTokenRejectedIterator is returned from FilterTokenRejected and is used to iterate over the raw logs and unpacked data for TokenRejected events raised by the TokenRegistry contract.
type TokenRegistryTokenRejectedIterator struct {
	Event *TokenRegistryTokenRejected // Event containing the contract specifics and raw log

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
func (it *TokenRegistryTokenRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryTokenRejected)
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
		it.Event = new(TokenRegistryTokenRejected)
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
func (it *TokenRegistryTokenRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryTokenRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryTokenRejected represents a TokenRejected event raised by the TokenRegistry contract.
type TokenRegistryTokenRejected struct {
	ContractAddress common.Address
	Reason          string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenRejected is a free log retrieval operation binding the contract event 0xfb4e86fc32b67bf889887f2e33859d9d6ed16aa324570d71171d77a7ee8114e5.
//
// Solidity: event TokenRejected(address indexed contractAddress, string reason)
func (_TokenRegistry *TokenRegistryFilterer) FilterTokenRejected(opts *bind.FilterOpts, contractAddress []common.Address) (*TokenRegistryTokenRejectedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "TokenRejected", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTokenRejectedIterator{contract: _TokenRegistry.contract, event: "TokenRejected", logs: logs, sub: sub}, nil
}

// WatchTokenRejected is a free log subscription operation binding the contract event 0xfb4e86fc32b67bf889887f2e33859d9d6ed16aa324570d71171d77a7ee8114e5.
//
// Solidity: event TokenRejected(address indexed contractAddress, string reason)
func (_TokenRegistry *TokenRegistryFilterer) WatchTokenRejected(opts *bind.WatchOpts, sink chan<- *TokenRegistryTokenRejected, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "TokenRejected", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryTokenRejected)
				if err := _TokenRegistry.contract.UnpackLog(event, "TokenRejected", log); err != nil {
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

// ParseTokenRejected is a log parse operation binding the contract event 0xfb4e86fc32b67bf889887f2e33859d9d6ed16aa324570d71171d77a7ee8114e5.
//
// Solidity: event TokenRejected(address indexed contractAddress, string reason)
func (_TokenRegistry *TokenRegistryFilterer) ParseTokenRejected(log types.Log) (*TokenRegistryTokenRejected, error) {
	event := new(TokenRegistryTokenRejected)
	if err := _TokenRegistry.contract.UnpackLog(event, "TokenRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRegistryTokenUpdatedIterator is returned from FilterTokenUpdated and is used to iterate over the raw logs and unpacked data for TokenUpdated events raised by the TokenRegistry contract.
type TokenRegistryTokenUpdatedIterator struct {
	Event *TokenRegistryTokenUpdated // Event containing the contract specifics and raw log

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
func (it *TokenRegistryTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryTokenUpdated)
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
		it.Event = new(TokenRegistryTokenUpdated)
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
func (it *TokenRegistryTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryTokenUpdated represents a TokenUpdated event raised by the TokenRegistry contract.
type TokenRegistryTokenUpdated struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenUpdated is a free log retrieval operation binding the contract event 0x5ba6b30cd4b2f9e9e67f4feb9b9df10d5da3b057598e6901b217b7d590345e30.
//
// Solidity: event TokenUpdated(address indexed contractAddress)
func (_TokenRegistry *TokenRegistryFilterer) FilterTokenUpdated(opts *bind.FilterOpts, contractAddress []common.Address) (*TokenRegistryTokenUpdatedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "TokenUpdated", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTokenUpdatedIterator{contract: _TokenRegistry.contract, event: "TokenUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenUpdated is a free log subscription operation binding the contract event 0x5ba6b30cd4b2f9e9e67f4feb9b9df10d5da3b057598e6901b217b7d590345e30.
//
// Solidity: event TokenUpdated(address indexed contractAddress)
func (_TokenRegistry *TokenRegistryFilterer) WatchTokenUpdated(opts *bind.WatchOpts, sink chan<- *TokenRegistryTokenUpdated, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "TokenUpdated", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryTokenUpdated)
				if err := _TokenRegistry.contract.UnpackLog(event, "TokenUpdated", log); err != nil {
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

// ParseTokenUpdated is a log parse operation binding the contract event 0x5ba6b30cd4b2f9e9e67f4feb9b9df10d5da3b057598e6901b217b7d590345e30.
//
// Solidity: event TokenUpdated(address indexed contractAddress)
func (_TokenRegistry *TokenRegistryFilterer) ParseTokenUpdated(log types.Log) (*TokenRegistryTokenUpdated, error) {
	event := new(TokenRegistryTokenUpdated)
	if err := _TokenRegistry.contract.UnpackLog(event, "TokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRegistryTokentrollerUpdatedIterator is returned from FilterTokentrollerUpdated and is used to iterate over the raw logs and unpacked data for TokentrollerUpdated events raised by the TokenRegistry contract.
type TokenRegistryTokentrollerUpdatedIterator struct {
	Event *TokenRegistryTokentrollerUpdated // Event containing the contract specifics and raw log

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
func (it *TokenRegistryTokentrollerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryTokentrollerUpdated)
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
		it.Event = new(TokenRegistryTokentrollerUpdated)
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
func (it *TokenRegistryTokentrollerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryTokentrollerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryTokentrollerUpdated represents a TokentrollerUpdated event raised by the TokenRegistry contract.
type TokenRegistryTokentrollerUpdated struct {
	NewTokentroller common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokentrollerUpdated is a free log retrieval operation binding the contract event 0x24628d813f923f48992318b49019f3219233a50975fe9694bf526acffe4330ca.
//
// Solidity: event TokentrollerUpdated(address indexed newTokentroller)
func (_TokenRegistry *TokenRegistryFilterer) FilterTokentrollerUpdated(opts *bind.FilterOpts, newTokentroller []common.Address) (*TokenRegistryTokentrollerUpdatedIterator, error) {

	var newTokentrollerRule []interface{}
	for _, newTokentrollerItem := range newTokentroller {
		newTokentrollerRule = append(newTokentrollerRule, newTokentrollerItem)
	}

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "TokentrollerUpdated", newTokentrollerRule)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTokentrollerUpdatedIterator{contract: _TokenRegistry.contract, event: "TokentrollerUpdated", logs: logs, sub: sub}, nil
}

// WatchTokentrollerUpdated is a free log subscription operation binding the contract event 0x24628d813f923f48992318b49019f3219233a50975fe9694bf526acffe4330ca.
//
// Solidity: event TokentrollerUpdated(address indexed newTokentroller)
func (_TokenRegistry *TokenRegistryFilterer) WatchTokentrollerUpdated(opts *bind.WatchOpts, sink chan<- *TokenRegistryTokentrollerUpdated, newTokentroller []common.Address) (event.Subscription, error) {

	var newTokentrollerRule []interface{}
	for _, newTokentrollerItem := range newTokentroller {
		newTokentrollerRule = append(newTokentrollerRule, newTokentrollerItem)
	}

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "TokentrollerUpdated", newTokentrollerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryTokentrollerUpdated)
				if err := _TokenRegistry.contract.UnpackLog(event, "TokentrollerUpdated", log); err != nil {
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

// ParseTokentrollerUpdated is a log parse operation binding the contract event 0x24628d813f923f48992318b49019f3219233a50975fe9694bf526acffe4330ca.
//
// Solidity: event TokentrollerUpdated(address indexed newTokentroller)
func (_TokenRegistry *TokenRegistryFilterer) ParseTokentrollerUpdated(log types.Log) (*TokenRegistryTokentrollerUpdated, error) {
	event := new(TokenRegistryTokentrollerUpdated)
	if err := _TokenRegistry.contract.UnpackLog(event, "TokentrollerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
