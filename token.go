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
)

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountCurrency\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountTokens\",\"type\":\"uint256\"}],\"name\":\"AutoLiquify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"ContractSwapEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEAD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_hasLiqBeenAdded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_ratios\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"marketing\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"charity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"operations\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"burn\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"totalSwap\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_taxRates\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"buyFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"sellFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"transferFee\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_taxWallets\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"marketing\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"charity\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"operations\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractSwapEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dexRouter\",\"outputs\":[{\"internalType\":\"contractIRouter02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableProtectionsLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxTX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxWallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromLimits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromProtection\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBuyTaxes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxRoundtripTax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSellTaxes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTransferTaxes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTxLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxWalletLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"multiSendTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"piContractSwapsEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"piSwapPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protectionsLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeBlacklisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeSniper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"swapEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"priceImpactSwapEnabled\",\"type\":\"bool\"}],\"name\":\"setContractSwapEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setExcludedFromFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setExcludedFromLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setExcludedFromProtection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"setGasPriceLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"}],\"name\":\"setInitializer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"divisor\",\"type\":\"uint256\"}],\"name\":\"setMaxTxPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"divisor\",\"type\":\"uint256\"}],\"name\":\"setMaxWalletSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceImpactSwapPercent\",\"type\":\"uint256\"}],\"name\":\"setPriceImpactSwapAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_antiSnipe\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_antiGas\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_antiBlock\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_algo\",\"type\":\"bool\"}],\"name\":\"setProtectionSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"marketing\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"operations\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"charity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"burn\",\"type\":\"uint16\"}],\"name\":\"setRatios\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"thresholdPercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"thresholdDivisor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountDivisor\",\"type\":\"uint256\"}],\"name\":\"setSwapSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"buyFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"sellFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"transferFee\",\"type\":\"uint16\"}],\"name\":\"setTaxes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setTaxesAutomatic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"marketing\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"operations\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"charity\",\"type\":\"address\"}],\"name\":\"setWallets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sweepContingency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxesAutomatic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// DEAD is a free data retrieval call binding the contract method 0x03fd2a45.
//
// Solidity: function DEAD() view returns(address)
func (_Token *TokenCaller) DEAD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "DEAD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEAD is a free data retrieval call binding the contract method 0x03fd2a45.
//
// Solidity: function DEAD() view returns(address)
func (_Token *TokenSession) DEAD() (common.Address, error) {
	return _Token.Contract.DEAD(&_Token.CallOpts)
}

// DEAD is a free data retrieval call binding the contract method 0x03fd2a45.
//
// Solidity: function DEAD() view returns(address)
func (_Token *TokenCallerSession) DEAD() (common.Address, error) {
	return _Token.Contract.DEAD(&_Token.CallOpts)
}

// ZERO is a free data retrieval call binding the contract method 0x58fa63ca.
//
// Solidity: function ZERO() view returns(address)
func (_Token *TokenCaller) ZERO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "ZERO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZERO is a free data retrieval call binding the contract method 0x58fa63ca.
//
// Solidity: function ZERO() view returns(address)
func (_Token *TokenSession) ZERO() (common.Address, error) {
	return _Token.Contract.ZERO(&_Token.CallOpts)
}

// ZERO is a free data retrieval call binding the contract method 0x58fa63ca.
//
// Solidity: function ZERO() view returns(address)
func (_Token *TokenCallerSession) ZERO() (common.Address, error) {
	return _Token.Contract.ZERO(&_Token.CallOpts)
}

// HasLiqBeenAdded is a free data retrieval call binding the contract method 0x50a8e016.
//
// Solidity: function _hasLiqBeenAdded() view returns(bool)
func (_Token *TokenCaller) HasLiqBeenAdded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "_hasLiqBeenAdded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasLiqBeenAdded is a free data retrieval call binding the contract method 0x50a8e016.
//
// Solidity: function _hasLiqBeenAdded() view returns(bool)
func (_Token *TokenSession) HasLiqBeenAdded() (bool, error) {
	return _Token.Contract.HasLiqBeenAdded(&_Token.CallOpts)
}

// HasLiqBeenAdded is a free data retrieval call binding the contract method 0x50a8e016.
//
// Solidity: function _hasLiqBeenAdded() view returns(bool)
func (_Token *TokenCallerSession) HasLiqBeenAdded() (bool, error) {
	return _Token.Contract.HasLiqBeenAdded(&_Token.CallOpts)
}

// Ratios is a free data retrieval call binding the contract method 0xcf847706.
//
// Solidity: function _ratios() view returns(uint16 marketing, uint16 charity, uint16 operations, uint16 burn, uint16 totalSwap)
func (_Token *TokenCaller) Ratios(opts *bind.CallOpts) (struct {
	Marketing  uint16
	Charity    uint16
	Operations uint16
	Burn       uint16
	TotalSwap  uint16
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "_ratios")

	outstruct := new(struct {
		Marketing  uint16
		Charity    uint16
		Operations uint16
		Burn       uint16
		TotalSwap  uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Marketing = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Charity = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.Operations = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.Burn = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.TotalSwap = *abi.ConvertType(out[4], new(uint16)).(*uint16)

	return *outstruct, err

}

// Ratios is a free data retrieval call binding the contract method 0xcf847706.
//
// Solidity: function _ratios() view returns(uint16 marketing, uint16 charity, uint16 operations, uint16 burn, uint16 totalSwap)
func (_Token *TokenSession) Ratios() (struct {
	Marketing  uint16
	Charity    uint16
	Operations uint16
	Burn       uint16
	TotalSwap  uint16
}, error) {
	return _Token.Contract.Ratios(&_Token.CallOpts)
}

// Ratios is a free data retrieval call binding the contract method 0xcf847706.
//
// Solidity: function _ratios() view returns(uint16 marketing, uint16 charity, uint16 operations, uint16 burn, uint16 totalSwap)
func (_Token *TokenCallerSession) Ratios() (struct {
	Marketing  uint16
	Charity    uint16
	Operations uint16
	Burn       uint16
	TotalSwap  uint16
}, error) {
	return _Token.Contract.Ratios(&_Token.CallOpts)
}

// TaxRates is a free data retrieval call binding the contract method 0x069d955f.
//
// Solidity: function _taxRates() view returns(uint16 buyFee, uint16 sellFee, uint16 transferFee)
func (_Token *TokenCaller) TaxRates(opts *bind.CallOpts) (struct {
	BuyFee      uint16
	SellFee     uint16
	TransferFee uint16
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "_taxRates")

	outstruct := new(struct {
		BuyFee      uint16
		SellFee     uint16
		TransferFee uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BuyFee = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.SellFee = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.TransferFee = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// TaxRates is a free data retrieval call binding the contract method 0x069d955f.
//
// Solidity: function _taxRates() view returns(uint16 buyFee, uint16 sellFee, uint16 transferFee)
func (_Token *TokenSession) TaxRates() (struct {
	BuyFee      uint16
	SellFee     uint16
	TransferFee uint16
}, error) {
	return _Token.Contract.TaxRates(&_Token.CallOpts)
}

// TaxRates is a free data retrieval call binding the contract method 0x069d955f.
//
// Solidity: function _taxRates() view returns(uint16 buyFee, uint16 sellFee, uint16 transferFee)
func (_Token *TokenCallerSession) TaxRates() (struct {
	BuyFee      uint16
	SellFee     uint16
	TransferFee uint16
}, error) {
	return _Token.Contract.TaxRates(&_Token.CallOpts)
}

// TaxWallets is a free data retrieval call binding the contract method 0xf94aa1b4.
//
// Solidity: function _taxWallets() view returns(address marketing, address charity, address operations)
func (_Token *TokenCaller) TaxWallets(opts *bind.CallOpts) (struct {
	Marketing  common.Address
	Charity    common.Address
	Operations common.Address
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "_taxWallets")

	outstruct := new(struct {
		Marketing  common.Address
		Charity    common.Address
		Operations common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Marketing = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Charity = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Operations = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// TaxWallets is a free data retrieval call binding the contract method 0xf94aa1b4.
//
// Solidity: function _taxWallets() view returns(address marketing, address charity, address operations)
func (_Token *TokenSession) TaxWallets() (struct {
	Marketing  common.Address
	Charity    common.Address
	Operations common.Address
}, error) {
	return _Token.Contract.TaxWallets(&_Token.CallOpts)
}

// TaxWallets is a free data retrieval call binding the contract method 0xf94aa1b4.
//
// Solidity: function _taxWallets() view returns(address marketing, address charity, address operations)
func (_Token *TokenCallerSession) TaxWallets() (struct {
	Marketing  common.Address
	Charity    common.Address
	Operations common.Address
}, error) {
	return _Token.Contract.TaxWallets(&_Token.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_Token *TokenSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_Token *TokenCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, holder, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Token *TokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, account)
}

// BurnEnabled is a free data retrieval call binding the contract method 0x5dc96d16.
//
// Solidity: function burnEnabled() view returns(bool)
func (_Token *TokenCaller) BurnEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "burnEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnEnabled is a free data retrieval call binding the contract method 0x5dc96d16.
//
// Solidity: function burnEnabled() view returns(bool)
func (_Token *TokenSession) BurnEnabled() (bool, error) {
	return _Token.Contract.BurnEnabled(&_Token.CallOpts)
}

// BurnEnabled is a free data retrieval call binding the contract method 0x5dc96d16.
//
// Solidity: function burnEnabled() view returns(bool)
func (_Token *TokenCallerSession) BurnEnabled() (bool, error) {
	return _Token.Contract.BurnEnabled(&_Token.CallOpts)
}

// ContractSwapEnabled is a free data retrieval call binding the contract method 0xfdb78c0e.
//
// Solidity: function contractSwapEnabled() view returns(bool)
func (_Token *TokenCaller) ContractSwapEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "contractSwapEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractSwapEnabled is a free data retrieval call binding the contract method 0xfdb78c0e.
//
// Solidity: function contractSwapEnabled() view returns(bool)
func (_Token *TokenSession) ContractSwapEnabled() (bool, error) {
	return _Token.Contract.ContractSwapEnabled(&_Token.CallOpts)
}

// ContractSwapEnabled is a free data retrieval call binding the contract method 0xfdb78c0e.
//
// Solidity: function contractSwapEnabled() view returns(bool)
func (_Token *TokenCallerSession) ContractSwapEnabled() (bool, error) {
	return _Token.Contract.ContractSwapEnabled(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// DexRouter is a free data retrieval call binding the contract method 0x0758d924.
//
// Solidity: function dexRouter() view returns(address)
func (_Token *TokenCaller) DexRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "dexRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DexRouter is a free data retrieval call binding the contract method 0x0758d924.
//
// Solidity: function dexRouter() view returns(address)
func (_Token *TokenSession) DexRouter() (common.Address, error) {
	return _Token.Contract.DexRouter(&_Token.CallOpts)
}

// DexRouter is a free data retrieval call binding the contract method 0x0758d924.
//
// Solidity: function dexRouter() view returns(address)
func (_Token *TokenCallerSession) DexRouter() (common.Address, error) {
	return _Token.Contract.DexRouter(&_Token.CallOpts)
}

// GetMaxTX is a free data retrieval call binding the contract method 0x6ebd0078.
//
// Solidity: function getMaxTX() view returns(uint256)
func (_Token *TokenCaller) GetMaxTX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getMaxTX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxTX is a free data retrieval call binding the contract method 0x6ebd0078.
//
// Solidity: function getMaxTX() view returns(uint256)
func (_Token *TokenSession) GetMaxTX() (*big.Int, error) {
	return _Token.Contract.GetMaxTX(&_Token.CallOpts)
}

// GetMaxTX is a free data retrieval call binding the contract method 0x6ebd0078.
//
// Solidity: function getMaxTX() view returns(uint256)
func (_Token *TokenCallerSession) GetMaxTX() (*big.Int, error) {
	return _Token.Contract.GetMaxTX(&_Token.CallOpts)
}

// GetMaxWallet is a free data retrieval call binding the contract method 0x0fa604e4.
//
// Solidity: function getMaxWallet() view returns(uint256)
func (_Token *TokenCaller) GetMaxWallet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getMaxWallet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxWallet is a free data retrieval call binding the contract method 0x0fa604e4.
//
// Solidity: function getMaxWallet() view returns(uint256)
func (_Token *TokenSession) GetMaxWallet() (*big.Int, error) {
	return _Token.Contract.GetMaxWallet(&_Token.CallOpts)
}

// GetMaxWallet is a free data retrieval call binding the contract method 0x0fa604e4.
//
// Solidity: function getMaxWallet() view returns(uint256)
func (_Token *TokenCallerSession) GetMaxWallet() (*big.Int, error) {
	return _Token.Contract.GetMaxWallet(&_Token.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Token *TokenCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Token *TokenSession) GetOwner() (common.Address, error) {
	return _Token.Contract.GetOwner(&_Token.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Token *TokenCallerSession) GetOwner() (common.Address, error) {
	return _Token.Contract.GetOwner(&_Token.CallOpts)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_Token *TokenCaller) IsBlacklisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isBlacklisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_Token *TokenSession) IsBlacklisted(account common.Address) (bool, error) {
	return _Token.Contract.IsBlacklisted(&_Token.CallOpts, account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_Token *TokenCallerSession) IsBlacklisted(account common.Address) (bool, error) {
	return _Token.Contract.IsBlacklisted(&_Token.CallOpts, account)
}

// IsExcludedFromFees is a free data retrieval call binding the contract method 0x4fbee193.
//
// Solidity: function isExcludedFromFees(address account) view returns(bool)
func (_Token *TokenCaller) IsExcludedFromFees(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isExcludedFromFees", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromFees is a free data retrieval call binding the contract method 0x4fbee193.
//
// Solidity: function isExcludedFromFees(address account) view returns(bool)
func (_Token *TokenSession) IsExcludedFromFees(account common.Address) (bool, error) {
	return _Token.Contract.IsExcludedFromFees(&_Token.CallOpts, account)
}

// IsExcludedFromFees is a free data retrieval call binding the contract method 0x4fbee193.
//
// Solidity: function isExcludedFromFees(address account) view returns(bool)
func (_Token *TokenCallerSession) IsExcludedFromFees(account common.Address) (bool, error) {
	return _Token.Contract.IsExcludedFromFees(&_Token.CallOpts, account)
}

// IsExcludedFromLimits is a free data retrieval call binding the contract method 0x5cce86cd.
//
// Solidity: function isExcludedFromLimits(address account) view returns(bool)
func (_Token *TokenCaller) IsExcludedFromLimits(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isExcludedFromLimits", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromLimits is a free data retrieval call binding the contract method 0x5cce86cd.
//
// Solidity: function isExcludedFromLimits(address account) view returns(bool)
func (_Token *TokenSession) IsExcludedFromLimits(account common.Address) (bool, error) {
	return _Token.Contract.IsExcludedFromLimits(&_Token.CallOpts, account)
}

// IsExcludedFromLimits is a free data retrieval call binding the contract method 0x5cce86cd.
//
// Solidity: function isExcludedFromLimits(address account) view returns(bool)
func (_Token *TokenCallerSession) IsExcludedFromLimits(account common.Address) (bool, error) {
	return _Token.Contract.IsExcludedFromLimits(&_Token.CallOpts, account)
}

// IsExcludedFromProtection is a free data retrieval call binding the contract method 0x0dcbcf1c.
//
// Solidity: function isExcludedFromProtection(address account) view returns(bool)
func (_Token *TokenCaller) IsExcludedFromProtection(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isExcludedFromProtection", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromProtection is a free data retrieval call binding the contract method 0x0dcbcf1c.
//
// Solidity: function isExcludedFromProtection(address account) view returns(bool)
func (_Token *TokenSession) IsExcludedFromProtection(account common.Address) (bool, error) {
	return _Token.Contract.IsExcludedFromProtection(&_Token.CallOpts, account)
}

// IsExcludedFromProtection is a free data retrieval call binding the contract method 0x0dcbcf1c.
//
// Solidity: function isExcludedFromProtection(address account) view returns(bool)
func (_Token *TokenCallerSession) IsExcludedFromProtection(account common.Address) (bool, error) {
	return _Token.Contract.IsExcludedFromProtection(&_Token.CallOpts, account)
}

// LpPair is a free data retrieval call binding the contract method 0x452ed4f1.
//
// Solidity: function lpPair() view returns(address)
func (_Token *TokenCaller) LpPair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "lpPair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpPair is a free data retrieval call binding the contract method 0x452ed4f1.
//
// Solidity: function lpPair() view returns(address)
func (_Token *TokenSession) LpPair() (common.Address, error) {
	return _Token.Contract.LpPair(&_Token.CallOpts)
}

// LpPair is a free data retrieval call binding the contract method 0x452ed4f1.
//
// Solidity: function lpPair() view returns(address)
func (_Token *TokenCallerSession) LpPair() (common.Address, error) {
	return _Token.Contract.LpPair(&_Token.CallOpts)
}

// MaxBuyTaxes is a free data retrieval call binding the contract method 0x2b28fc7a.
//
// Solidity: function maxBuyTaxes() view returns(uint256)
func (_Token *TokenCaller) MaxBuyTaxes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "maxBuyTaxes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBuyTaxes is a free data retrieval call binding the contract method 0x2b28fc7a.
//
// Solidity: function maxBuyTaxes() view returns(uint256)
func (_Token *TokenSession) MaxBuyTaxes() (*big.Int, error) {
	return _Token.Contract.MaxBuyTaxes(&_Token.CallOpts)
}

// MaxBuyTaxes is a free data retrieval call binding the contract method 0x2b28fc7a.
//
// Solidity: function maxBuyTaxes() view returns(uint256)
func (_Token *TokenCallerSession) MaxBuyTaxes() (*big.Int, error) {
	return _Token.Contract.MaxBuyTaxes(&_Token.CallOpts)
}

// MaxRoundtripTax is a free data retrieval call binding the contract method 0x46ea7ac8.
//
// Solidity: function maxRoundtripTax() view returns(uint256)
func (_Token *TokenCaller) MaxRoundtripTax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "maxRoundtripTax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRoundtripTax is a free data retrieval call binding the contract method 0x46ea7ac8.
//
// Solidity: function maxRoundtripTax() view returns(uint256)
func (_Token *TokenSession) MaxRoundtripTax() (*big.Int, error) {
	return _Token.Contract.MaxRoundtripTax(&_Token.CallOpts)
}

// MaxRoundtripTax is a free data retrieval call binding the contract method 0x46ea7ac8.
//
// Solidity: function maxRoundtripTax() view returns(uint256)
func (_Token *TokenCallerSession) MaxRoundtripTax() (*big.Int, error) {
	return _Token.Contract.MaxRoundtripTax(&_Token.CallOpts)
}

// MaxSellTaxes is a free data retrieval call binding the contract method 0xb3d514fb.
//
// Solidity: function maxSellTaxes() view returns(uint256)
func (_Token *TokenCaller) MaxSellTaxes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "maxSellTaxes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSellTaxes is a free data retrieval call binding the contract method 0xb3d514fb.
//
// Solidity: function maxSellTaxes() view returns(uint256)
func (_Token *TokenSession) MaxSellTaxes() (*big.Int, error) {
	return _Token.Contract.MaxSellTaxes(&_Token.CallOpts)
}

// MaxSellTaxes is a free data retrieval call binding the contract method 0xb3d514fb.
//
// Solidity: function maxSellTaxes() view returns(uint256)
func (_Token *TokenCallerSession) MaxSellTaxes() (*big.Int, error) {
	return _Token.Contract.MaxSellTaxes(&_Token.CallOpts)
}

// MaxTransferTaxes is a free data retrieval call binding the contract method 0xb1b08f71.
//
// Solidity: function maxTransferTaxes() view returns(uint256)
func (_Token *TokenCaller) MaxTransferTaxes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "maxTransferTaxes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTransferTaxes is a free data retrieval call binding the contract method 0xb1b08f71.
//
// Solidity: function maxTransferTaxes() view returns(uint256)
func (_Token *TokenSession) MaxTransferTaxes() (*big.Int, error) {
	return _Token.Contract.MaxTransferTaxes(&_Token.CallOpts)
}

// MaxTransferTaxes is a free data retrieval call binding the contract method 0xb1b08f71.
//
// Solidity: function maxTransferTaxes() view returns(uint256)
func (_Token *TokenCallerSession) MaxTransferTaxes() (*big.Int, error) {
	return _Token.Contract.MaxTransferTaxes(&_Token.CallOpts)
}

// MaxTxLocked is a free data retrieval call binding the contract method 0x9d6ee83a.
//
// Solidity: function maxTxLocked() view returns(bool)
func (_Token *TokenCaller) MaxTxLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "maxTxLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MaxTxLocked is a free data retrieval call binding the contract method 0x9d6ee83a.
//
// Solidity: function maxTxLocked() view returns(bool)
func (_Token *TokenSession) MaxTxLocked() (bool, error) {
	return _Token.Contract.MaxTxLocked(&_Token.CallOpts)
}

// MaxTxLocked is a free data retrieval call binding the contract method 0x9d6ee83a.
//
// Solidity: function maxTxLocked() view returns(bool)
func (_Token *TokenCallerSession) MaxTxLocked() (bool, error) {
	return _Token.Contract.MaxTxLocked(&_Token.CallOpts)
}

// MaxWalletLocked is a free data retrieval call binding the contract method 0xf25df6fc.
//
// Solidity: function maxWalletLocked() view returns(bool)
func (_Token *TokenCaller) MaxWalletLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "maxWalletLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MaxWalletLocked is a free data retrieval call binding the contract method 0xf25df6fc.
//
// Solidity: function maxWalletLocked() view returns(bool)
func (_Token *TokenSession) MaxWalletLocked() (bool, error) {
	return _Token.Contract.MaxWalletLocked(&_Token.CallOpts)
}

// MaxWalletLocked is a free data retrieval call binding the contract method 0xf25df6fc.
//
// Solidity: function maxWalletLocked() view returns(bool)
func (_Token *TokenCallerSession) MaxWalletLocked() (bool, error) {
	return _Token.Contract.MaxWalletLocked(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// PiContractSwapsEnabled is a free data retrieval call binding the contract method 0xb7df8b36.
//
// Solidity: function piContractSwapsEnabled() view returns(bool)
func (_Token *TokenCaller) PiContractSwapsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "piContractSwapsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PiContractSwapsEnabled is a free data retrieval call binding the contract method 0xb7df8b36.
//
// Solidity: function piContractSwapsEnabled() view returns(bool)
func (_Token *TokenSession) PiContractSwapsEnabled() (bool, error) {
	return _Token.Contract.PiContractSwapsEnabled(&_Token.CallOpts)
}

// PiContractSwapsEnabled is a free data retrieval call binding the contract method 0xb7df8b36.
//
// Solidity: function piContractSwapsEnabled() view returns(bool)
func (_Token *TokenCallerSession) PiContractSwapsEnabled() (bool, error) {
	return _Token.Contract.PiContractSwapsEnabled(&_Token.CallOpts)
}

// PiSwapPercent is a free data retrieval call binding the contract method 0x28577751.
//
// Solidity: function piSwapPercent() view returns(uint256)
func (_Token *TokenCaller) PiSwapPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "piSwapPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PiSwapPercent is a free data retrieval call binding the contract method 0x28577751.
//
// Solidity: function piSwapPercent() view returns(uint256)
func (_Token *TokenSession) PiSwapPercent() (*big.Int, error) {
	return _Token.Contract.PiSwapPercent(&_Token.CallOpts)
}

// PiSwapPercent is a free data retrieval call binding the contract method 0x28577751.
//
// Solidity: function piSwapPercent() view returns(uint256)
func (_Token *TokenCallerSession) PiSwapPercent() (*big.Int, error) {
	return _Token.Contract.PiSwapPercent(&_Token.CallOpts)
}

// ProtectionsLocked is a free data retrieval call binding the contract method 0x81c3215c.
//
// Solidity: function protectionsLocked() view returns(bool)
func (_Token *TokenCaller) ProtectionsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "protectionsLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProtectionsLocked is a free data retrieval call binding the contract method 0x81c3215c.
//
// Solidity: function protectionsLocked() view returns(bool)
func (_Token *TokenSession) ProtectionsLocked() (bool, error) {
	return _Token.Contract.ProtectionsLocked(&_Token.CallOpts)
}

// ProtectionsLocked is a free data retrieval call binding the contract method 0x81c3215c.
//
// Solidity: function protectionsLocked() view returns(bool)
func (_Token *TokenCallerSession) ProtectionsLocked() (bool, error) {
	return _Token.Contract.ProtectionsLocked(&_Token.CallOpts)
}

// SwapAmount is a free data retrieval call binding the contract method 0x2e8fa821.
//
// Solidity: function swapAmount() view returns(uint256)
func (_Token *TokenCaller) SwapAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "swapAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapAmount is a free data retrieval call binding the contract method 0x2e8fa821.
//
// Solidity: function swapAmount() view returns(uint256)
func (_Token *TokenSession) SwapAmount() (*big.Int, error) {
	return _Token.Contract.SwapAmount(&_Token.CallOpts)
}

// SwapAmount is a free data retrieval call binding the contract method 0x2e8fa821.
//
// Solidity: function swapAmount() view returns(uint256)
func (_Token *TokenCallerSession) SwapAmount() (*big.Int, error) {
	return _Token.Contract.SwapAmount(&_Token.CallOpts)
}

// SwapThreshold is a free data retrieval call binding the contract method 0x0445b667.
//
// Solidity: function swapThreshold() view returns(uint256)
func (_Token *TokenCaller) SwapThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "swapThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapThreshold is a free data retrieval call binding the contract method 0x0445b667.
//
// Solidity: function swapThreshold() view returns(uint256)
func (_Token *TokenSession) SwapThreshold() (*big.Int, error) {
	return _Token.Contract.SwapThreshold(&_Token.CallOpts)
}

// SwapThreshold is a free data retrieval call binding the contract method 0x0445b667.
//
// Solidity: function swapThreshold() view returns(uint256)
func (_Token *TokenCallerSession) SwapThreshold() (*big.Int, error) {
	return _Token.Contract.SwapThreshold(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TaxesAutomatic is a free data retrieval call binding the contract method 0xae57b3ec.
//
// Solidity: function taxesAutomatic() view returns(bool)
func (_Token *TokenCaller) TaxesAutomatic(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "taxesAutomatic")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaxesAutomatic is a free data retrieval call binding the contract method 0xae57b3ec.
//
// Solidity: function taxesAutomatic() view returns(bool)
func (_Token *TokenSession) TaxesAutomatic() (bool, error) {
	return _Token.Contract.TaxesAutomatic(&_Token.CallOpts)
}

// TaxesAutomatic is a free data retrieval call binding the contract method 0xae57b3ec.
//
// Solidity: function taxesAutomatic() view returns(bool)
func (_Token *TokenCallerSession) TaxesAutomatic() (bool, error) {
	return _Token.Contract.TaxesAutomatic(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TradingEnabled is a free data retrieval call binding the contract method 0x4ada218b.
//
// Solidity: function tradingEnabled() view returns(bool)
func (_Token *TokenCaller) TradingEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "tradingEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TradingEnabled is a free data retrieval call binding the contract method 0x4ada218b.
//
// Solidity: function tradingEnabled() view returns(bool)
func (_Token *TokenSession) TradingEnabled() (bool, error) {
	return _Token.Contract.TradingEnabled(&_Token.CallOpts)
}

// TradingEnabled is a free data retrieval call binding the contract method 0x4ada218b.
//
// Solidity: function tradingEnabled() view returns(bool)
func (_Token *TokenCallerSession) TradingEnabled() (bool, error) {
	return _Token.Contract.TradingEnabled(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, amount)
}

// DisableProtectionsLock is a paid mutator transaction binding the contract method 0xb3da99e2.
//
// Solidity: function disableProtectionsLock() returns()
func (_Token *TokenTransactor) DisableProtectionsLock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "disableProtectionsLock")
}

// DisableProtectionsLock is a paid mutator transaction binding the contract method 0xb3da99e2.
//
// Solidity: function disableProtectionsLock() returns()
func (_Token *TokenSession) DisableProtectionsLock() (*types.Transaction, error) {
	return _Token.Contract.DisableProtectionsLock(&_Token.TransactOpts)
}

// DisableProtectionsLock is a paid mutator transaction binding the contract method 0xb3da99e2.
//
// Solidity: function disableProtectionsLock() returns()
func (_Token *TokenTransactorSession) DisableProtectionsLock() (*types.Transaction, error) {
	return _Token.Contract.DisableProtectionsLock(&_Token.TransactOpts)
}

// EnableTrading is a paid mutator transaction binding the contract method 0x8a8c523c.
//
// Solidity: function enableTrading() returns()
func (_Token *TokenTransactor) EnableTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "enableTrading")
}

// EnableTrading is a paid mutator transaction binding the contract method 0x8a8c523c.
//
// Solidity: function enableTrading() returns()
func (_Token *TokenSession) EnableTrading() (*types.Transaction, error) {
	return _Token.Contract.EnableTrading(&_Token.TransactOpts)
}

// EnableTrading is a paid mutator transaction binding the contract method 0x8a8c523c.
//
// Solidity: function enableTrading() returns()
func (_Token *TokenTransactorSession) EnableTrading() (*types.Transaction, error) {
	return _Token.Contract.EnableTrading(&_Token.TransactOpts)
}

// MultiSendTokens is a paid mutator transaction binding the contract method 0x2610eaca.
//
// Solidity: function multiSendTokens(address[] accounts, uint256[] amounts) returns()
func (_Token *TokenTransactor) MultiSendTokens(opts *bind.TransactOpts, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "multiSendTokens", accounts, amounts)
}

// MultiSendTokens is a paid mutator transaction binding the contract method 0x2610eaca.
//
// Solidity: function multiSendTokens(address[] accounts, uint256[] amounts) returns()
func (_Token *TokenSession) MultiSendTokens(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.MultiSendTokens(&_Token.TransactOpts, accounts, amounts)
}

// MultiSendTokens is a paid mutator transaction binding the contract method 0x2610eaca.
//
// Solidity: function multiSendTokens(address[] accounts, uint256[] amounts) returns()
func (_Token *TokenTransactorSession) MultiSendTokens(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.MultiSendTokens(&_Token.TransactOpts, accounts, amounts)
}

// RemoveBlacklisted is a paid mutator transaction binding the contract method 0xc6a276c2.
//
// Solidity: function removeBlacklisted(address account) returns()
func (_Token *TokenTransactor) RemoveBlacklisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeBlacklisted", account)
}

// RemoveBlacklisted is a paid mutator transaction binding the contract method 0xc6a276c2.
//
// Solidity: function removeBlacklisted(address account) returns()
func (_Token *TokenSession) RemoveBlacklisted(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveBlacklisted(&_Token.TransactOpts, account)
}

// RemoveBlacklisted is a paid mutator transaction binding the contract method 0xc6a276c2.
//
// Solidity: function removeBlacklisted(address account) returns()
func (_Token *TokenTransactorSession) RemoveBlacklisted(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveBlacklisted(&_Token.TransactOpts, account)
}

// RemoveSniper is a paid mutator transaction binding the contract method 0x33251a0b.
//
// Solidity: function removeSniper(address account) returns()
func (_Token *TokenTransactor) RemoveSniper(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeSniper", account)
}

// RemoveSniper is a paid mutator transaction binding the contract method 0x33251a0b.
//
// Solidity: function removeSniper(address account) returns()
func (_Token *TokenSession) RemoveSniper(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveSniper(&_Token.TransactOpts, account)
}

// RemoveSniper is a paid mutator transaction binding the contract method 0x33251a0b.
//
// Solidity: function removeSniper(address account) returns()
func (_Token *TokenTransactorSession) RemoveSniper(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveSniper(&_Token.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// SetContractSwapEnabled is a paid mutator transaction binding the contract method 0xbfc2fc35.
//
// Solidity: function setContractSwapEnabled(bool swapEnabled, bool priceImpactSwapEnabled) returns()
func (_Token *TokenTransactor) SetContractSwapEnabled(opts *bind.TransactOpts, swapEnabled bool, priceImpactSwapEnabled bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setContractSwapEnabled", swapEnabled, priceImpactSwapEnabled)
}

// SetContractSwapEnabled is a paid mutator transaction binding the contract method 0xbfc2fc35.
//
// Solidity: function setContractSwapEnabled(bool swapEnabled, bool priceImpactSwapEnabled) returns()
func (_Token *TokenSession) SetContractSwapEnabled(swapEnabled bool, priceImpactSwapEnabled bool) (*types.Transaction, error) {
	return _Token.Contract.SetContractSwapEnabled(&_Token.TransactOpts, swapEnabled, priceImpactSwapEnabled)
}

// SetContractSwapEnabled is a paid mutator transaction binding the contract method 0xbfc2fc35.
//
// Solidity: function setContractSwapEnabled(bool swapEnabled, bool priceImpactSwapEnabled) returns()
func (_Token *TokenTransactorSession) SetContractSwapEnabled(swapEnabled bool, priceImpactSwapEnabled bool) (*types.Transaction, error) {
	return _Token.Contract.SetContractSwapEnabled(&_Token.TransactOpts, swapEnabled, priceImpactSwapEnabled)
}

// SetExcludedFromFees is a paid mutator transaction binding the contract method 0x590ffdce.
//
// Solidity: function setExcludedFromFees(address account, bool enabled) returns()
func (_Token *TokenTransactor) SetExcludedFromFees(opts *bind.TransactOpts, account common.Address, enabled bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setExcludedFromFees", account, enabled)
}

// SetExcludedFromFees is a paid mutator transaction binding the contract method 0x590ffdce.
//
// Solidity: function setExcludedFromFees(address account, bool enabled) returns()
func (_Token *TokenSession) SetExcludedFromFees(account common.Address, enabled bool) (*types.Transaction, error) {
	return _Token.Contract.SetExcludedFromFees(&_Token.TransactOpts, account, enabled)
}

// SetExcludedFromFees is a paid mutator transaction binding the contract method 0x590ffdce.
//
// Solidity: function setExcludedFromFees(address account, bool enabled) returns()
func (_Token *TokenTransactorSession) SetExcludedFromFees(account common.Address, enabled bool) (*types.Transaction, error) {
	return _Token.Contract.SetExcludedFromFees(&_Token.TransactOpts, account, enabled)
}

// SetExcludedFromLimits is a paid mutator transaction binding the contract method 0x36fddb04.
//
// Solidity: function setExcludedFromLimits(address account, bool enabled) returns()
func (_Token *TokenTransactor) SetExcludedFromLimits(opts *bind.TransactOpts, account common.Address, enabled bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setExcludedFromLimits", account, enabled)
}

// SetExcludedFromLimits is a paid mutator transaction binding the contract method 0x36fddb04.
//
// Solidity: function setExcludedFromLimits(address account, bool enabled) returns()
func (_Token *TokenSession) SetExcludedFromLimits(account common.Address, enabled bool) (*types.Transaction, error) {
	return _Token.Contract.SetExcludedFromLimits(&_Token.TransactOpts, account, enabled)
}

// SetExcludedFromLimits is a paid mutator transaction binding the contract method 0x36fddb04.
//
// Solidity: function setExcludedFromLimits(address account, bool enabled) returns()
func (_Token *TokenTransactorSession) SetExcludedFromLimits(account common.Address, enabled bool) (*types.Transaction, error) {
	return _Token.Contract.SetExcludedFromLimits(&_Token.TransactOpts, account, enabled)
}

// SetExcludedFromProtection is a paid mutator transaction binding the contract method 0xfeadde9c.
//
// Solidity: function setExcludedFromProtection(address account, bool enabled) returns()
func (_Token *TokenTransactor) SetExcludedFromProtection(opts *bind.TransactOpts, account common.Address, enabled bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setExcludedFromProtection", account, enabled)
}

// SetExcludedFromProtection is a paid mutator transaction binding the contract method 0xfeadde9c.
//
// Solidity: function setExcludedFromProtection(address account, bool enabled) returns()
func (_Token *TokenSession) SetExcludedFromProtection(account common.Address, enabled bool) (*types.Transaction, error) {
	return _Token.Contract.SetExcludedFromProtection(&_Token.TransactOpts, account, enabled)
}

// SetExcludedFromProtection is a paid mutator transaction binding the contract method 0xfeadde9c.
//
// Solidity: function setExcludedFromProtection(address account, bool enabled) returns()
func (_Token *TokenTransactorSession) SetExcludedFromProtection(account common.Address, enabled bool) (*types.Transaction, error) {
	return _Token.Contract.SetExcludedFromProtection(&_Token.TransactOpts, account, enabled)
}

// SetGasPriceLimit is a paid mutator transaction binding the contract method 0x09231602.
//
// Solidity: function setGasPriceLimit(uint256 gas) returns()
func (_Token *TokenTransactor) SetGasPriceLimit(opts *bind.TransactOpts, gas *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setGasPriceLimit", gas)
}

// SetGasPriceLimit is a paid mutator transaction binding the contract method 0x09231602.
//
// Solidity: function setGasPriceLimit(uint256 gas) returns()
func (_Token *TokenSession) SetGasPriceLimit(gas *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetGasPriceLimit(&_Token.TransactOpts, gas)
}

// SetGasPriceLimit is a paid mutator transaction binding the contract method 0x09231602.
//
// Solidity: function setGasPriceLimit(uint256 gas) returns()
func (_Token *TokenTransactorSession) SetGasPriceLimit(gas *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetGasPriceLimit(&_Token.TransactOpts, gas)
}

// SetInitializer is a paid mutator transaction binding the contract method 0x5c24b074.
//
// Solidity: function setInitializer(address initializer) returns()
func (_Token *TokenTransactor) SetInitializer(opts *bind.TransactOpts, initializer common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setInitializer", initializer)
}

// SetInitializer is a paid mutator transaction binding the contract method 0x5c24b074.
//
// Solidity: function setInitializer(address initializer) returns()
func (_Token *TokenSession) SetInitializer(initializer common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetInitializer(&_Token.TransactOpts, initializer)
}

// SetInitializer is a paid mutator transaction binding the contract method 0x5c24b074.
//
// Solidity: function setInitializer(address initializer) returns()
func (_Token *TokenTransactorSession) SetInitializer(initializer common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetInitializer(&_Token.TransactOpts, initializer)
}

// SetMaxTxPercent is a paid mutator transaction binding the contract method 0x3f3cf56c.
//
// Solidity: function setMaxTxPercent(uint256 percent, uint256 divisor) returns()
func (_Token *TokenTransactor) SetMaxTxPercent(opts *bind.TransactOpts, percent *big.Int, divisor *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setMaxTxPercent", percent, divisor)
}

// SetMaxTxPercent is a paid mutator transaction binding the contract method 0x3f3cf56c.
//
// Solidity: function setMaxTxPercent(uint256 percent, uint256 divisor) returns()
func (_Token *TokenSession) SetMaxTxPercent(percent *big.Int, divisor *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetMaxTxPercent(&_Token.TransactOpts, percent, divisor)
}

// SetMaxTxPercent is a paid mutator transaction binding the contract method 0x3f3cf56c.
//
// Solidity: function setMaxTxPercent(uint256 percent, uint256 divisor) returns()
func (_Token *TokenTransactorSession) SetMaxTxPercent(percent *big.Int, divisor *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetMaxTxPercent(&_Token.TransactOpts, percent, divisor)
}

// SetMaxWalletSize is a paid mutator transaction binding the contract method 0x26003957.
//
// Solidity: function setMaxWalletSize(uint256 percent, uint256 divisor) returns()
func (_Token *TokenTransactor) SetMaxWalletSize(opts *bind.TransactOpts, percent *big.Int, divisor *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setMaxWalletSize", percent, divisor)
}

// SetMaxWalletSize is a paid mutator transaction binding the contract method 0x26003957.
//
// Solidity: function setMaxWalletSize(uint256 percent, uint256 divisor) returns()
func (_Token *TokenSession) SetMaxWalletSize(percent *big.Int, divisor *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetMaxWalletSize(&_Token.TransactOpts, percent, divisor)
}

// SetMaxWalletSize is a paid mutator transaction binding the contract method 0x26003957.
//
// Solidity: function setMaxWalletSize(uint256 percent, uint256 divisor) returns()
func (_Token *TokenTransactorSession) SetMaxWalletSize(percent *big.Int, divisor *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetMaxWalletSize(&_Token.TransactOpts, percent, divisor)
}

// SetPriceImpactSwapAmount is a paid mutator transaction binding the contract method 0x4e718e48.
//
// Solidity: function setPriceImpactSwapAmount(uint256 priceImpactSwapPercent) returns()
func (_Token *TokenTransactor) SetPriceImpactSwapAmount(opts *bind.TransactOpts, priceImpactSwapPercent *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setPriceImpactSwapAmount", priceImpactSwapPercent)
}

// SetPriceImpactSwapAmount is a paid mutator transaction binding the contract method 0x4e718e48.
//
// Solidity: function setPriceImpactSwapAmount(uint256 priceImpactSwapPercent) returns()
func (_Token *TokenSession) SetPriceImpactSwapAmount(priceImpactSwapPercent *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetPriceImpactSwapAmount(&_Token.TransactOpts, priceImpactSwapPercent)
}

// SetPriceImpactSwapAmount is a paid mutator transaction binding the contract method 0x4e718e48.
//
// Solidity: function setPriceImpactSwapAmount(uint256 priceImpactSwapPercent) returns()
func (_Token *TokenTransactorSession) SetPriceImpactSwapAmount(priceImpactSwapPercent *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetPriceImpactSwapAmount(&_Token.TransactOpts, priceImpactSwapPercent)
}

// SetProtectionSettings is a paid mutator transaction binding the contract method 0x29dd8798.
//
// Solidity: function setProtectionSettings(bool _antiSnipe, bool _antiGas, bool _antiBlock, bool _algo) returns()
func (_Token *TokenTransactor) SetProtectionSettings(opts *bind.TransactOpts, _antiSnipe bool, _antiGas bool, _antiBlock bool, _algo bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setProtectionSettings", _antiSnipe, _antiGas, _antiBlock, _algo)
}

// SetProtectionSettings is a paid mutator transaction binding the contract method 0x29dd8798.
//
// Solidity: function setProtectionSettings(bool _antiSnipe, bool _antiGas, bool _antiBlock, bool _algo) returns()
func (_Token *TokenSession) SetProtectionSettings(_antiSnipe bool, _antiGas bool, _antiBlock bool, _algo bool) (*types.Transaction, error) {
	return _Token.Contract.SetProtectionSettings(&_Token.TransactOpts, _antiSnipe, _antiGas, _antiBlock, _algo)
}

// SetProtectionSettings is a paid mutator transaction binding the contract method 0x29dd8798.
//
// Solidity: function setProtectionSettings(bool _antiSnipe, bool _antiGas, bool _antiBlock, bool _algo) returns()
func (_Token *TokenTransactorSession) SetProtectionSettings(_antiSnipe bool, _antiGas bool, _antiBlock bool, _algo bool) (*types.Transaction, error) {
	return _Token.Contract.SetProtectionSettings(&_Token.TransactOpts, _antiSnipe, _antiGas, _antiBlock, _algo)
}

// SetRatios is a paid mutator transaction binding the contract method 0x0712d165.
//
// Solidity: function setRatios(uint16 marketing, uint16 operations, uint16 charity, uint16 burn) returns()
func (_Token *TokenTransactor) SetRatios(opts *bind.TransactOpts, marketing uint16, operations uint16, charity uint16, burn uint16) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setRatios", marketing, operations, charity, burn)
}

// SetRatios is a paid mutator transaction binding the contract method 0x0712d165.
//
// Solidity: function setRatios(uint16 marketing, uint16 operations, uint16 charity, uint16 burn) returns()
func (_Token *TokenSession) SetRatios(marketing uint16, operations uint16, charity uint16, burn uint16) (*types.Transaction, error) {
	return _Token.Contract.SetRatios(&_Token.TransactOpts, marketing, operations, charity, burn)
}

// SetRatios is a paid mutator transaction binding the contract method 0x0712d165.
//
// Solidity: function setRatios(uint16 marketing, uint16 operations, uint16 charity, uint16 burn) returns()
func (_Token *TokenTransactorSession) SetRatios(marketing uint16, operations uint16, charity uint16, burn uint16) (*types.Transaction, error) {
	return _Token.Contract.SetRatios(&_Token.TransactOpts, marketing, operations, charity, burn)
}

// SetSwapSettings is a paid mutator transaction binding the contract method 0xfb78680d.
//
// Solidity: function setSwapSettings(uint256 thresholdPercent, uint256 thresholdDivisor, uint256 amountPercent, uint256 amountDivisor) returns()
func (_Token *TokenTransactor) SetSwapSettings(opts *bind.TransactOpts, thresholdPercent *big.Int, thresholdDivisor *big.Int, amountPercent *big.Int, amountDivisor *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setSwapSettings", thresholdPercent, thresholdDivisor, amountPercent, amountDivisor)
}

// SetSwapSettings is a paid mutator transaction binding the contract method 0xfb78680d.
//
// Solidity: function setSwapSettings(uint256 thresholdPercent, uint256 thresholdDivisor, uint256 amountPercent, uint256 amountDivisor) returns()
func (_Token *TokenSession) SetSwapSettings(thresholdPercent *big.Int, thresholdDivisor *big.Int, amountPercent *big.Int, amountDivisor *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetSwapSettings(&_Token.TransactOpts, thresholdPercent, thresholdDivisor, amountPercent, amountDivisor)
}

// SetSwapSettings is a paid mutator transaction binding the contract method 0xfb78680d.
//
// Solidity: function setSwapSettings(uint256 thresholdPercent, uint256 thresholdDivisor, uint256 amountPercent, uint256 amountDivisor) returns()
func (_Token *TokenTransactorSession) SetSwapSettings(thresholdPercent *big.Int, thresholdDivisor *big.Int, amountPercent *big.Int, amountDivisor *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetSwapSettings(&_Token.TransactOpts, thresholdPercent, thresholdDivisor, amountPercent, amountDivisor)
}

// SetTaxes is a paid mutator transaction binding the contract method 0x32cde664.
//
// Solidity: function setTaxes(uint16 buyFee, uint16 sellFee, uint16 transferFee) returns()
func (_Token *TokenTransactor) SetTaxes(opts *bind.TransactOpts, buyFee uint16, sellFee uint16, transferFee uint16) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setTaxes", buyFee, sellFee, transferFee)
}

// SetTaxes is a paid mutator transaction binding the contract method 0x32cde664.
//
// Solidity: function setTaxes(uint16 buyFee, uint16 sellFee, uint16 transferFee) returns()
func (_Token *TokenSession) SetTaxes(buyFee uint16, sellFee uint16, transferFee uint16) (*types.Transaction, error) {
	return _Token.Contract.SetTaxes(&_Token.TransactOpts, buyFee, sellFee, transferFee)
}

// SetTaxes is a paid mutator transaction binding the contract method 0x32cde664.
//
// Solidity: function setTaxes(uint16 buyFee, uint16 sellFee, uint16 transferFee) returns()
func (_Token *TokenTransactorSession) SetTaxes(buyFee uint16, sellFee uint16, transferFee uint16) (*types.Transaction, error) {
	return _Token.Contract.SetTaxes(&_Token.TransactOpts, buyFee, sellFee, transferFee)
}

// SetTaxesAutomatic is a paid mutator transaction binding the contract method 0x3be05999.
//
// Solidity: function setTaxesAutomatic(bool enabled) returns()
func (_Token *TokenTransactor) SetTaxesAutomatic(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setTaxesAutomatic", enabled)
}

// SetTaxesAutomatic is a paid mutator transaction binding the contract method 0x3be05999.
//
// Solidity: function setTaxesAutomatic(bool enabled) returns()
func (_Token *TokenSession) SetTaxesAutomatic(enabled bool) (*types.Transaction, error) {
	return _Token.Contract.SetTaxesAutomatic(&_Token.TransactOpts, enabled)
}

// SetTaxesAutomatic is a paid mutator transaction binding the contract method 0x3be05999.
//
// Solidity: function setTaxesAutomatic(bool enabled) returns()
func (_Token *TokenTransactorSession) SetTaxesAutomatic(enabled bool) (*types.Transaction, error) {
	return _Token.Contract.SetTaxesAutomatic(&_Token.TransactOpts, enabled)
}

// SetWallets is a paid mutator transaction binding the contract method 0x75cb1bd1.
//
// Solidity: function setWallets(address marketing, address operations, address charity) returns()
func (_Token *TokenTransactor) SetWallets(opts *bind.TransactOpts, marketing common.Address, operations common.Address, charity common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setWallets", marketing, operations, charity)
}

// SetWallets is a paid mutator transaction binding the contract method 0x75cb1bd1.
//
// Solidity: function setWallets(address marketing, address operations, address charity) returns()
func (_Token *TokenSession) SetWallets(marketing common.Address, operations common.Address, charity common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetWallets(&_Token.TransactOpts, marketing, operations, charity)
}

// SetWallets is a paid mutator transaction binding the contract method 0x75cb1bd1.
//
// Solidity: function setWallets(address marketing, address operations, address charity) returns()
func (_Token *TokenTransactorSession) SetWallets(marketing common.Address, operations common.Address, charity common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetWallets(&_Token.TransactOpts, marketing, operations, charity)
}

// SweepContingency is a paid mutator transaction binding the contract method 0xee5d9c2d.
//
// Solidity: function sweepContingency() returns()
func (_Token *TokenTransactor) SweepContingency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "sweepContingency")
}

// SweepContingency is a paid mutator transaction binding the contract method 0xee5d9c2d.
//
// Solidity: function sweepContingency() returns()
func (_Token *TokenSession) SweepContingency() (*types.Transaction, error) {
	return _Token.Contract.SweepContingency(&_Token.TransactOpts)
}

// SweepContingency is a paid mutator transaction binding the contract method 0xee5d9c2d.
//
// Solidity: function sweepContingency() returns()
func (_Token *TokenTransactorSession) SweepContingency() (*types.Transaction, error) {
	return _Token.Contract.SweepContingency(&_Token.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, sender, recipient, amount)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address newOwner) returns()
func (_Token *TokenTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address newOwner) returns()
func (_Token *TokenSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwner(&_Token.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address newOwner) returns()
func (_Token *TokenTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwner(&_Token.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenSession) Receive() (*types.Transaction, error) {
	return _Token.Contract.Receive(&_Token.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenTransactorSession) Receive() (*types.Transaction, error) {
	return _Token.Contract.Receive(&_Token.TransactOpts)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenAutoLiquifyIterator is returned from FilterAutoLiquify and is used to iterate over the raw logs and unpacked data for AutoLiquify events raised by the Token contract.
type TokenAutoLiquifyIterator struct {
	Event *TokenAutoLiquify // Event containing the contract specifics and raw log

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
func (it *TokenAutoLiquifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAutoLiquify)
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
		it.Event = new(TokenAutoLiquify)
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
func (it *TokenAutoLiquifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAutoLiquifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAutoLiquify represents a AutoLiquify event raised by the Token contract.
type TokenAutoLiquify struct {
	AmountCurrency *big.Int
	AmountTokens   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAutoLiquify is a free log retrieval operation binding the contract event 0x424db2872186fa7e7afa7a5e902ed3b49a2ef19c2f5431e672462495dd6b4506.
//
// Solidity: event AutoLiquify(uint256 amountCurrency, uint256 amountTokens)
func (_Token *TokenFilterer) FilterAutoLiquify(opts *bind.FilterOpts) (*TokenAutoLiquifyIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "AutoLiquify")
	if err != nil {
		return nil, err
	}
	return &TokenAutoLiquifyIterator{contract: _Token.contract, event: "AutoLiquify", logs: logs, sub: sub}, nil
}

// WatchAutoLiquify is a free log subscription operation binding the contract event 0x424db2872186fa7e7afa7a5e902ed3b49a2ef19c2f5431e672462495dd6b4506.
//
// Solidity: event AutoLiquify(uint256 amountCurrency, uint256 amountTokens)
func (_Token *TokenFilterer) WatchAutoLiquify(opts *bind.WatchOpts, sink chan<- *TokenAutoLiquify) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "AutoLiquify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAutoLiquify)
				if err := _Token.contract.UnpackLog(event, "AutoLiquify", log); err != nil {
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

// ParseAutoLiquify is a log parse operation binding the contract event 0x424db2872186fa7e7afa7a5e902ed3b49a2ef19c2f5431e672462495dd6b4506.
//
// Solidity: event AutoLiquify(uint256 amountCurrency, uint256 amountTokens)
func (_Token *TokenFilterer) ParseAutoLiquify(log types.Log) (*TokenAutoLiquify, error) {
	event := new(TokenAutoLiquify)
	if err := _Token.contract.UnpackLog(event, "AutoLiquify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenContractSwapEnabledUpdatedIterator is returned from FilterContractSwapEnabledUpdated and is used to iterate over the raw logs and unpacked data for ContractSwapEnabledUpdated events raised by the Token contract.
type TokenContractSwapEnabledUpdatedIterator struct {
	Event *TokenContractSwapEnabledUpdated // Event containing the contract specifics and raw log

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
func (it *TokenContractSwapEnabledUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenContractSwapEnabledUpdated)
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
		it.Event = new(TokenContractSwapEnabledUpdated)
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
func (it *TokenContractSwapEnabledUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenContractSwapEnabledUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenContractSwapEnabledUpdated represents a ContractSwapEnabledUpdated event raised by the Token contract.
type TokenContractSwapEnabledUpdated struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractSwapEnabledUpdated is a free log retrieval operation binding the contract event 0x7b0a47d3b0234280b6c9213c5bbff44c8b6001bea7770b3950280f9141053257.
//
// Solidity: event ContractSwapEnabledUpdated(bool enabled)
func (_Token *TokenFilterer) FilterContractSwapEnabledUpdated(opts *bind.FilterOpts) (*TokenContractSwapEnabledUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "ContractSwapEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenContractSwapEnabledUpdatedIterator{contract: _Token.contract, event: "ContractSwapEnabledUpdated", logs: logs, sub: sub}, nil
}

// WatchContractSwapEnabledUpdated is a free log subscription operation binding the contract event 0x7b0a47d3b0234280b6c9213c5bbff44c8b6001bea7770b3950280f9141053257.
//
// Solidity: event ContractSwapEnabledUpdated(bool enabled)
func (_Token *TokenFilterer) WatchContractSwapEnabledUpdated(opts *bind.WatchOpts, sink chan<- *TokenContractSwapEnabledUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "ContractSwapEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenContractSwapEnabledUpdated)
				if err := _Token.contract.UnpackLog(event, "ContractSwapEnabledUpdated", log); err != nil {
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

// ParseContractSwapEnabledUpdated is a log parse operation binding the contract event 0x7b0a47d3b0234280b6c9213c5bbff44c8b6001bea7770b3950280f9141053257.
//
// Solidity: event ContractSwapEnabledUpdated(bool enabled)
func (_Token *TokenFilterer) ParseContractSwapEnabledUpdated(log types.Log) (*TokenContractSwapEnabledUpdated, error) {
	event := new(TokenContractSwapEnabledUpdated)
	if err := _Token.contract.UnpackLog(event, "ContractSwapEnabledUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token contract.
type TokenOwnershipTransferredIterator struct {
	Event *TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnershipTransferred)
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
		it.Event = new(TokenOwnershipTransferred)
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
func (it *TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnershipTransferred represents a OwnershipTransferred event raised by the Token contract.
type TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnershipTransferredIterator{contract: _Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnershipTransferred)
				if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Token *TokenFilterer) ParseOwnershipTransferred(log types.Log) (*TokenOwnershipTransferred, error) {
	event := new(TokenOwnershipTransferred)
	if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
