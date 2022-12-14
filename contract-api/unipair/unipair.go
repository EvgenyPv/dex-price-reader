// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unipair

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

// UnipairMetaData contains all meta data concerning the Unipair contract.
var UnipairMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UnipairABI is the input ABI used to generate the binding from.
// Deprecated: Use UnipairMetaData.ABI instead.
var UnipairABI = UnipairMetaData.ABI

// Unipair is an auto generated Go binding around an Ethereum contract.
type Unipair struct {
	UnipairCaller     // Read-only binding to the contract
	UnipairTransactor // Write-only binding to the contract
	UnipairFilterer   // Log filterer for contract events
}

// UnipairCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnipairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnipairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnipairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnipairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnipairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnipairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnipairSession struct {
	Contract     *Unipair          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnipairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnipairCallerSession struct {
	Contract *UnipairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// UnipairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnipairTransactorSession struct {
	Contract     *UnipairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UnipairRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnipairRaw struct {
	Contract *Unipair // Generic contract binding to access the raw methods on
}

// UnipairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnipairCallerRaw struct {
	Contract *UnipairCaller // Generic read-only contract binding to access the raw methods on
}

// UnipairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnipairTransactorRaw struct {
	Contract *UnipairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnipair creates a new instance of Unipair, bound to a specific deployed contract.
func NewUnipair(address common.Address, backend bind.ContractBackend) (*Unipair, error) {
	contract, err := bindUnipair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Unipair{UnipairCaller: UnipairCaller{contract: contract}, UnipairTransactor: UnipairTransactor{contract: contract}, UnipairFilterer: UnipairFilterer{contract: contract}}, nil
}

// NewUnipairCaller creates a new read-only instance of Unipair, bound to a specific deployed contract.
func NewUnipairCaller(address common.Address, caller bind.ContractCaller) (*UnipairCaller, error) {
	contract, err := bindUnipair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnipairCaller{contract: contract}, nil
}

// NewUnipairTransactor creates a new write-only instance of Unipair, bound to a specific deployed contract.
func NewUnipairTransactor(address common.Address, transactor bind.ContractTransactor) (*UnipairTransactor, error) {
	contract, err := bindUnipair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnipairTransactor{contract: contract}, nil
}

// NewUnipairFilterer creates a new log filterer instance of Unipair, bound to a specific deployed contract.
func NewUnipairFilterer(address common.Address, filterer bind.ContractFilterer) (*UnipairFilterer, error) {
	contract, err := bindUnipair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnipairFilterer{contract: contract}, nil
}

// bindUnipair binds a generic wrapper to an already deployed contract.
func bindUnipair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnipairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unipair *UnipairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unipair.Contract.UnipairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unipair *UnipairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unipair.Contract.UnipairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unipair *UnipairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unipair.Contract.UnipairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unipair *UnipairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unipair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unipair *UnipairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unipair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unipair *UnipairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unipair.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Unipair *UnipairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Unipair *UnipairSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Unipair.Contract.DOMAINSEPARATOR(&_Unipair.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Unipair *UnipairCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Unipair.Contract.DOMAINSEPARATOR(&_Unipair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Unipair *UnipairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Unipair *UnipairSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Unipair.Contract.MINIMUMLIQUIDITY(&_Unipair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Unipair *UnipairCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Unipair.Contract.MINIMUMLIQUIDITY(&_Unipair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Unipair *UnipairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Unipair *UnipairSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Unipair.Contract.PERMITTYPEHASH(&_Unipair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Unipair *UnipairCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Unipair.Contract.PERMITTYPEHASH(&_Unipair.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Unipair *UnipairCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Unipair *UnipairSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Unipair.Contract.Allowance(&_Unipair.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Unipair *UnipairCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Unipair.Contract.Allowance(&_Unipair.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Unipair *UnipairCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Unipair *UnipairSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Unipair.Contract.BalanceOf(&_Unipair.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Unipair *UnipairCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Unipair.Contract.BalanceOf(&_Unipair.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Unipair *UnipairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Unipair *UnipairSession) Decimals() (uint8, error) {
	return _Unipair.Contract.Decimals(&_Unipair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Unipair *UnipairCallerSession) Decimals() (uint8, error) {
	return _Unipair.Contract.Decimals(&_Unipair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Unipair *UnipairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Unipair *UnipairSession) Factory() (common.Address, error) {
	return _Unipair.Contract.Factory(&_Unipair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Unipair *UnipairCallerSession) Factory() (common.Address, error) {
	return _Unipair.Contract.Factory(&_Unipair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Unipair *UnipairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Unipair *UnipairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Unipair.Contract.GetReserves(&_Unipair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Unipair *UnipairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Unipair.Contract.GetReserves(&_Unipair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Unipair *UnipairCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Unipair *UnipairSession) KLast() (*big.Int, error) {
	return _Unipair.Contract.KLast(&_Unipair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Unipair *UnipairCallerSession) KLast() (*big.Int, error) {
	return _Unipair.Contract.KLast(&_Unipair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Unipair *UnipairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Unipair *UnipairSession) Name() (string, error) {
	return _Unipair.Contract.Name(&_Unipair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Unipair *UnipairCallerSession) Name() (string, error) {
	return _Unipair.Contract.Name(&_Unipair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Unipair *UnipairCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Unipair *UnipairSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Unipair.Contract.Nonces(&_Unipair.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Unipair *UnipairCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Unipair.Contract.Nonces(&_Unipair.CallOpts, arg0)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Unipair *UnipairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Unipair *UnipairSession) Price0CumulativeLast() (*big.Int, error) {
	return _Unipair.Contract.Price0CumulativeLast(&_Unipair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Unipair *UnipairCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _Unipair.Contract.Price0CumulativeLast(&_Unipair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Unipair *UnipairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Unipair *UnipairSession) Price1CumulativeLast() (*big.Int, error) {
	return _Unipair.Contract.Price1CumulativeLast(&_Unipair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Unipair *UnipairCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _Unipair.Contract.Price1CumulativeLast(&_Unipair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Unipair *UnipairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Unipair *UnipairSession) Symbol() (string, error) {
	return _Unipair.Contract.Symbol(&_Unipair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Unipair *UnipairCallerSession) Symbol() (string, error) {
	return _Unipair.Contract.Symbol(&_Unipair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Unipair *UnipairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Unipair *UnipairSession) Token0() (common.Address, error) {
	return _Unipair.Contract.Token0(&_Unipair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Unipair *UnipairCallerSession) Token0() (common.Address, error) {
	return _Unipair.Contract.Token0(&_Unipair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Unipair *UnipairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Unipair *UnipairSession) Token1() (common.Address, error) {
	return _Unipair.Contract.Token1(&_Unipair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Unipair *UnipairCallerSession) Token1() (common.Address, error) {
	return _Unipair.Contract.Token1(&_Unipair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Unipair *UnipairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unipair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Unipair *UnipairSession) TotalSupply() (*big.Int, error) {
	return _Unipair.Contract.TotalSupply(&_Unipair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Unipair *UnipairCallerSession) TotalSupply() (*big.Int, error) {
	return _Unipair.Contract.TotalSupply(&_Unipair.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Unipair *UnipairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Unipair.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Unipair *UnipairSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Unipair.Contract.Approve(&_Unipair.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Unipair *UnipairTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Unipair.Contract.Approve(&_Unipair.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Unipair *UnipairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Unipair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Unipair *UnipairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Unipair.Contract.Burn(&_Unipair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Unipair *UnipairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Unipair.Contract.Burn(&_Unipair.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Unipair *UnipairTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _Unipair.contract.Transact(opts, "initialize", _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Unipair *UnipairSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _Unipair.Contract.Initialize(&_Unipair.TransactOpts, _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Unipair *UnipairTransactorSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _Unipair.Contract.Initialize(&_Unipair.TransactOpts, _token0, _token1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Unipair *UnipairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Unipair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Unipair *UnipairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Unipair.Contract.Mint(&_Unipair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Unipair *UnipairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Unipair.Contract.Mint(&_Unipair.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Unipair *UnipairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Unipair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Unipair *UnipairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Unipair.Contract.Permit(&_Unipair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Unipair *UnipairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Unipair.Contract.Permit(&_Unipair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Unipair *UnipairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Unipair.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Unipair *UnipairSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Unipair.Contract.Skim(&_Unipair.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Unipair *UnipairTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Unipair.Contract.Skim(&_Unipair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Unipair *UnipairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Unipair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Unipair *UnipairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Unipair.Contract.Swap(&_Unipair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Unipair *UnipairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Unipair.Contract.Swap(&_Unipair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Unipair *UnipairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unipair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Unipair *UnipairSession) Sync() (*types.Transaction, error) {
	return _Unipair.Contract.Sync(&_Unipair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Unipair *UnipairTransactorSession) Sync() (*types.Transaction, error) {
	return _Unipair.Contract.Sync(&_Unipair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Unipair *UnipairTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Unipair.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Unipair *UnipairSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Unipair.Contract.Transfer(&_Unipair.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Unipair *UnipairTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Unipair.Contract.Transfer(&_Unipair.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Unipair *UnipairTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Unipair.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Unipair *UnipairSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Unipair.Contract.TransferFrom(&_Unipair.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Unipair *UnipairTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Unipair.Contract.TransferFrom(&_Unipair.TransactOpts, from, to, value)
}

// UnipairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Unipair contract.
type UnipairApprovalIterator struct {
	Event *UnipairApproval // Event containing the contract specifics and raw log

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
func (it *UnipairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnipairApproval)
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
		it.Event = new(UnipairApproval)
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
func (it *UnipairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnipairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnipairApproval represents a Approval event raised by the Unipair contract.
type UnipairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Unipair *UnipairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UnipairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Unipair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UnipairApprovalIterator{contract: _Unipair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Unipair *UnipairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UnipairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Unipair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnipairApproval)
				if err := _Unipair.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Unipair *UnipairFilterer) ParseApproval(log types.Log) (*UnipairApproval, error) {
	event := new(UnipairApproval)
	if err := _Unipair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnipairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Unipair contract.
type UnipairBurnIterator struct {
	Event *UnipairBurn // Event containing the contract specifics and raw log

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
func (it *UnipairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnipairBurn)
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
		it.Event = new(UnipairBurn)
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
func (it *UnipairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnipairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnipairBurn represents a Burn event raised by the Unipair contract.
type UnipairBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Unipair *UnipairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*UnipairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unipair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UnipairBurnIterator{contract: _Unipair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Unipair *UnipairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *UnipairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unipair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnipairBurn)
				if err := _Unipair.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Unipair *UnipairFilterer) ParseBurn(log types.Log) (*UnipairBurn, error) {
	event := new(UnipairBurn)
	if err := _Unipair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnipairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Unipair contract.
type UnipairMintIterator struct {
	Event *UnipairMint // Event containing the contract specifics and raw log

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
func (it *UnipairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnipairMint)
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
		it.Event = new(UnipairMint)
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
func (it *UnipairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnipairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnipairMint represents a Mint event raised by the Unipair contract.
type UnipairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Unipair *UnipairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*UnipairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Unipair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &UnipairMintIterator{contract: _Unipair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Unipair *UnipairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *UnipairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Unipair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnipairMint)
				if err := _Unipair.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Unipair *UnipairFilterer) ParseMint(log types.Log) (*UnipairMint, error) {
	event := new(UnipairMint)
	if err := _Unipair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnipairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Unipair contract.
type UnipairSwapIterator struct {
	Event *UnipairSwap // Event containing the contract specifics and raw log

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
func (it *UnipairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnipairSwap)
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
		it.Event = new(UnipairSwap)
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
func (it *UnipairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnipairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnipairSwap represents a Swap event raised by the Unipair contract.
type UnipairSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Unipair *UnipairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*UnipairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unipair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UnipairSwapIterator{contract: _Unipair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Unipair *UnipairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *UnipairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unipair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnipairSwap)
				if err := _Unipair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Unipair *UnipairFilterer) ParseSwap(log types.Log) (*UnipairSwap, error) {
	event := new(UnipairSwap)
	if err := _Unipair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnipairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Unipair contract.
type UnipairSyncIterator struct {
	Event *UnipairSync // Event containing the contract specifics and raw log

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
func (it *UnipairSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnipairSync)
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
		it.Event = new(UnipairSync)
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
func (it *UnipairSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnipairSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnipairSync represents a Sync event raised by the Unipair contract.
type UnipairSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Unipair *UnipairFilterer) FilterSync(opts *bind.FilterOpts) (*UnipairSyncIterator, error) {

	logs, sub, err := _Unipair.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &UnipairSyncIterator{contract: _Unipair.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Unipair *UnipairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *UnipairSync) (event.Subscription, error) {

	logs, sub, err := _Unipair.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnipairSync)
				if err := _Unipair.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Unipair *UnipairFilterer) ParseSync(log types.Log) (*UnipairSync, error) {
	event := new(UnipairSync)
	if err := _Unipair.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnipairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Unipair contract.
type UnipairTransferIterator struct {
	Event *UnipairTransfer // Event containing the contract specifics and raw log

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
func (it *UnipairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnipairTransfer)
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
		it.Event = new(UnipairTransfer)
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
func (it *UnipairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnipairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnipairTransfer represents a Transfer event raised by the Unipair contract.
type UnipairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Unipair *UnipairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UnipairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unipair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UnipairTransferIterator{contract: _Unipair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Unipair *UnipairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UnipairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Unipair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnipairTransfer)
				if err := _Unipair.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Unipair *UnipairFilterer) ParseTransfer(log types.Log) (*UnipairTransfer, error) {
	event := new(UnipairTransfer)
	if err := _Unipair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
