// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package biddingFactory

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

// BiddingFactoryMetaData contains all meta data concerning the BiddingFactory contract.
var BiddingFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getActiveGames\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfActiveGames\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BiddingFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use BiddingFactoryMetaData.ABI instead.
var BiddingFactoryABI = BiddingFactoryMetaData.ABI

// BiddingFactory is an auto generated Go binding around an Ethereum contract.
type BiddingFactory struct {
	BiddingFactoryCaller     // Read-only binding to the contract
	BiddingFactoryTransactor // Write-only binding to the contract
	BiddingFactoryFilterer   // Log filterer for contract events
}

// BiddingFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BiddingFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BiddingFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BiddingFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BiddingFactorySession struct {
	Contract     *BiddingFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BiddingFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BiddingFactoryCallerSession struct {
	Contract *BiddingFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BiddingFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BiddingFactoryTransactorSession struct {
	Contract     *BiddingFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BiddingFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BiddingFactoryRaw struct {
	Contract *BiddingFactory // Generic contract binding to access the raw methods on
}

// BiddingFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BiddingFactoryCallerRaw struct {
	Contract *BiddingFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BiddingFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BiddingFactoryTransactorRaw struct {
	Contract *BiddingFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBiddingFactory creates a new instance of BiddingFactory, bound to a specific deployed contract.
func NewBiddingFactory(address common.Address, backend bind.ContractBackend) (*BiddingFactory, error) {
	contract, err := bindBiddingFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BiddingFactory{BiddingFactoryCaller: BiddingFactoryCaller{contract: contract}, BiddingFactoryTransactor: BiddingFactoryTransactor{contract: contract}, BiddingFactoryFilterer: BiddingFactoryFilterer{contract: contract}}, nil
}

// NewBiddingFactoryCaller creates a new read-only instance of BiddingFactory, bound to a specific deployed contract.
func NewBiddingFactoryCaller(address common.Address, caller bind.ContractCaller) (*BiddingFactoryCaller, error) {
	contract, err := bindBiddingFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BiddingFactoryCaller{contract: contract}, nil
}

// NewBiddingFactoryTransactor creates a new write-only instance of BiddingFactory, bound to a specific deployed contract.
func NewBiddingFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BiddingFactoryTransactor, error) {
	contract, err := bindBiddingFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BiddingFactoryTransactor{contract: contract}, nil
}

// NewBiddingFactoryFilterer creates a new log filterer instance of BiddingFactory, bound to a specific deployed contract.
func NewBiddingFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BiddingFactoryFilterer, error) {
	contract, err := bindBiddingFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BiddingFactoryFilterer{contract: contract}, nil
}

// bindBiddingFactory binds a generic wrapper to an already deployed contract.
func bindBiddingFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BiddingFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BiddingFactory *BiddingFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BiddingFactory.Contract.BiddingFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BiddingFactory *BiddingFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BiddingFactory.Contract.BiddingFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BiddingFactory *BiddingFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BiddingFactory.Contract.BiddingFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BiddingFactory *BiddingFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BiddingFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BiddingFactory *BiddingFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BiddingFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BiddingFactory *BiddingFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BiddingFactory.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BiddingFactory *BiddingFactoryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BiddingFactory.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BiddingFactory *BiddingFactorySession) Admin() (common.Address, error) {
	return _BiddingFactory.Contract.Admin(&_BiddingFactory.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BiddingFactory *BiddingFactoryCallerSession) Admin() (common.Address, error) {
	return _BiddingFactory.Contract.Admin(&_BiddingFactory.CallOpts)
}

// GetActiveGames is a free data retrieval call binding the contract method 0x4fd0fb9e.
//
// Solidity: function getActiveGames(uint256 _id) view returns(address)
func (_BiddingFactory *BiddingFactoryCaller) GetActiveGames(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BiddingFactory.contract.Call(opts, &out, "getActiveGames", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetActiveGames is a free data retrieval call binding the contract method 0x4fd0fb9e.
//
// Solidity: function getActiveGames(uint256 _id) view returns(address)
func (_BiddingFactory *BiddingFactorySession) GetActiveGames(_id *big.Int) (common.Address, error) {
	return _BiddingFactory.Contract.GetActiveGames(&_BiddingFactory.CallOpts, _id)
}

// GetActiveGames is a free data retrieval call binding the contract method 0x4fd0fb9e.
//
// Solidity: function getActiveGames(uint256 _id) view returns(address)
func (_BiddingFactory *BiddingFactoryCallerSession) GetActiveGames(_id *big.Int) (common.Address, error) {
	return _BiddingFactory.Contract.GetActiveGames(&_BiddingFactory.CallOpts, _id)
}

// NumberOfActiveGames is a free data retrieval call binding the contract method 0x5aa0e14e.
//
// Solidity: function numberOfActiveGames() view returns(uint256)
func (_BiddingFactory *BiddingFactoryCaller) NumberOfActiveGames(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BiddingFactory.contract.Call(opts, &out, "numberOfActiveGames")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfActiveGames is a free data retrieval call binding the contract method 0x5aa0e14e.
//
// Solidity: function numberOfActiveGames() view returns(uint256)
func (_BiddingFactory *BiddingFactorySession) NumberOfActiveGames() (*big.Int, error) {
	return _BiddingFactory.Contract.NumberOfActiveGames(&_BiddingFactory.CallOpts)
}

// NumberOfActiveGames is a free data retrieval call binding the contract method 0x5aa0e14e.
//
// Solidity: function numberOfActiveGames() view returns(uint256)
func (_BiddingFactory *BiddingFactoryCallerSession) NumberOfActiveGames() (*big.Int, error) {
	return _BiddingFactory.Contract.NumberOfActiveGames(&_BiddingFactory.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (_BiddingFactory *BiddingFactoryTransactor) Create(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BiddingFactory.contract.Transact(opts, "create")
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (_BiddingFactory *BiddingFactorySession) Create() (*types.Transaction, error) {
	return _BiddingFactory.Contract.Create(&_BiddingFactory.TransactOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns()
func (_BiddingFactory *BiddingFactoryTransactorSession) Create() (*types.Transaction, error) {
	return _BiddingFactory.Contract.Create(&_BiddingFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newAdmin) returns()
func (_BiddingFactory *BiddingFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BiddingFactory.contract.Transact(opts, "transferOwnership", newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newAdmin) returns()
func (_BiddingFactory *BiddingFactorySession) TransferOwnership(newAdmin common.Address) (*types.Transaction, error) {
	return _BiddingFactory.Contract.TransferOwnership(&_BiddingFactory.TransactOpts, newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newAdmin) returns()
func (_BiddingFactory *BiddingFactoryTransactorSession) TransferOwnership(newAdmin common.Address) (*types.Transaction, error) {
	return _BiddingFactory.Contract.TransferOwnership(&_BiddingFactory.TransactOpts, newAdmin)
}
