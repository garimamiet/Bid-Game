// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bidding

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

// BiddingMetaData contains all meta data concerning the Bidding contract.
var BiddingMetaData = &bind.MetaData{
<<<<<<< HEAD
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"bidStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"commissionWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidValue\",\"type\":\"uint256\"}],\"name\":\"newBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"rewardClaimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COMMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkWinner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readyToStart\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardCollected\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBidding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
=======
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"bidStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"commissionWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidValue\",\"type\":\"uint256\"}],\"name\":\"newBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"rewardDistributed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COMMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readyToStart\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardCollected\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBidding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
}

// BiddingABI is the input ABI used to generate the binding from.
// Deprecated: Use BiddingMetaData.ABI instead.
var BiddingABI = BiddingMetaData.ABI

// Bidding is an auto generated Go binding around an Ethereum contract.
type Bidding struct {
	BiddingCaller     // Read-only binding to the contract
	BiddingTransactor // Write-only binding to the contract
	BiddingFilterer   // Log filterer for contract events
}

// BiddingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BiddingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BiddingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BiddingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BiddingSession struct {
	Contract     *Bidding          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BiddingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BiddingCallerSession struct {
	Contract *BiddingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BiddingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BiddingTransactorSession struct {
	Contract     *BiddingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BiddingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BiddingRaw struct {
	Contract *Bidding // Generic contract binding to access the raw methods on
}

// BiddingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BiddingCallerRaw struct {
	Contract *BiddingCaller // Generic read-only contract binding to access the raw methods on
}

// BiddingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BiddingTransactorRaw struct {
	Contract *BiddingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBidding creates a new instance of Bidding, bound to a specific deployed contract.
func NewBidding(address common.Address, backend bind.ContractBackend) (*Bidding, error) {
	contract, err := bindBidding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bidding{BiddingCaller: BiddingCaller{contract: contract}, BiddingTransactor: BiddingTransactor{contract: contract}, BiddingFilterer: BiddingFilterer{contract: contract}}, nil
}

// NewBiddingCaller creates a new read-only instance of Bidding, bound to a specific deployed contract.
func NewBiddingCaller(address common.Address, caller bind.ContractCaller) (*BiddingCaller, error) {
	contract, err := bindBidding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BiddingCaller{contract: contract}, nil
}

// NewBiddingTransactor creates a new write-only instance of Bidding, bound to a specific deployed contract.
func NewBiddingTransactor(address common.Address, transactor bind.ContractTransactor) (*BiddingTransactor, error) {
	contract, err := bindBidding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BiddingTransactor{contract: contract}, nil
}

// NewBiddingFilterer creates a new log filterer instance of Bidding, bound to a specific deployed contract.
func NewBiddingFilterer(address common.Address, filterer bind.ContractFilterer) (*BiddingFilterer, error) {
	contract, err := bindBidding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BiddingFilterer{contract: contract}, nil
}

// bindBidding binds a generic wrapper to an already deployed contract.
func bindBidding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BiddingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bidding *BiddingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bidding.Contract.BiddingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bidding *BiddingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bidding.Contract.BiddingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bidding *BiddingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bidding.Contract.BiddingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bidding *BiddingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bidding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bidding *BiddingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bidding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bidding *BiddingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bidding.Contract.contract.Transact(opts, method, params...)
}

// COMMISSION is a free data retrieval call binding the contract method 0x562df3d5.
//
// Solidity: function COMMISSION() view returns(uint256)
func (_Bidding *BiddingCaller) COMMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bidding.contract.Call(opts, &out, "COMMISSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMISSION is a free data retrieval call binding the contract method 0x562df3d5.
//
// Solidity: function COMMISSION() view returns(uint256)
func (_Bidding *BiddingSession) COMMISSION() (*big.Int, error) {
	return _Bidding.Contract.COMMISSION(&_Bidding.CallOpts)
}

// COMMISSION is a free data retrieval call binding the contract method 0x562df3d5.
//
// Solidity: function COMMISSION() view returns(uint256)
func (_Bidding *BiddingCallerSession) COMMISSION() (*big.Int, error) {
	return _Bidding.Contract.COMMISSION(&_Bidding.CallOpts)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_Bidding *BiddingCaller) DURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bidding.contract.Call(opts, &out, "DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_Bidding *BiddingSession) DURATION() (*big.Int, error) {
	return _Bidding.Contract.DURATION(&_Bidding.CallOpts)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_Bidding *BiddingCallerSession) DURATION() (*big.Int, error) {
	return _Bidding.Contract.DURATION(&_Bidding.CallOpts)
}

// EXTENSION is a free data retrieval call binding the contract method 0x46f13619.
//
// Solidity: function EXTENSION() view returns(uint256)
func (_Bidding *BiddingCaller) EXTENSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bidding.contract.Call(opts, &out, "EXTENSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXTENSION is a free data retrieval call binding the contract method 0x46f13619.
//
// Solidity: function EXTENSION() view returns(uint256)
func (_Bidding *BiddingSession) EXTENSION() (*big.Int, error) {
	return _Bidding.Contract.EXTENSION(&_Bidding.CallOpts)
}

// EXTENSION is a free data retrieval call binding the contract method 0x46f13619.
//
// Solidity: function EXTENSION() view returns(uint256)
func (_Bidding *BiddingCallerSession) EXTENSION() (*big.Int, error) {
	return _Bidding.Contract.EXTENSION(&_Bidding.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Bidding *BiddingCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bidding.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Bidding *BiddingSession) Admin() (common.Address, error) {
	return _Bidding.Contract.Admin(&_Bidding.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Bidding *BiddingCallerSession) Admin() (common.Address, error) {
	return _Bidding.Contract.Admin(&_Bidding.CallOpts)
}

<<<<<<< HEAD
// CheckWinner is a free data retrieval call binding the contract method 0xad38867e.
//
// Solidity: function checkWinner() view returns(address)
func (_Bidding *BiddingCaller) CheckWinner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bidding.contract.Call(opts, &out, "checkWinner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckWinner is a free data retrieval call binding the contract method 0xad38867e.
//
// Solidity: function checkWinner() view returns(address)
func (_Bidding *BiddingSession) CheckWinner() (common.Address, error) {
	return _Bidding.Contract.CheckWinner(&_Bidding.CallOpts)
}

// CheckWinner is a free data retrieval call binding the contract method 0xad38867e.
//
// Solidity: function checkWinner() view returns(address)
func (_Bidding *BiddingCallerSession) CheckWinner() (common.Address, error) {
	return _Bidding.Contract.CheckWinner(&_Bidding.CallOpts)
}

=======
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_Bidding *BiddingCaller) Expiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bidding.contract.Call(opts, &out, "expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_Bidding *BiddingSession) Expiry() (*big.Int, error) {
	return _Bidding.Contract.Expiry(&_Bidding.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_Bidding *BiddingCallerSession) Expiry() (*big.Int, error) {
	return _Bidding.Contract.Expiry(&_Bidding.CallOpts)
}

// LastBid is a free data retrieval call binding the contract method 0x5e3d3957.
//
// Solidity: function lastBid() view returns(uint256)
func (_Bidding *BiddingCaller) LastBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bidding.contract.Call(opts, &out, "lastBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBid is a free data retrieval call binding the contract method 0x5e3d3957.
//
// Solidity: function lastBid() view returns(uint256)
func (_Bidding *BiddingSession) LastBid() (*big.Int, error) {
	return _Bidding.Contract.LastBid(&_Bidding.CallOpts)
}

// LastBid is a free data retrieval call binding the contract method 0x5e3d3957.
//
// Solidity: function lastBid() view returns(uint256)
func (_Bidding *BiddingCallerSession) LastBid() (*big.Int, error) {
	return _Bidding.Contract.LastBid(&_Bidding.CallOpts)
}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_Bidding *BiddingCaller) LastBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bidding.contract.Call(opts, &out, "lastBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_Bidding *BiddingSession) LastBidder() (common.Address, error) {
	return _Bidding.Contract.LastBidder(&_Bidding.CallOpts)
}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_Bidding *BiddingCallerSession) LastBidder() (common.Address, error) {
	return _Bidding.Contract.LastBidder(&_Bidding.CallOpts)
}

// ReadyToStart is a free data retrieval call binding the contract method 0x304b39a4.
//
// Solidity: function readyToStart() view returns(bool)
func (_Bidding *BiddingCaller) ReadyToStart(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bidding.contract.Call(opts, &out, "readyToStart")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReadyToStart is a free data retrieval call binding the contract method 0x304b39a4.
//
// Solidity: function readyToStart() view returns(bool)
func (_Bidding *BiddingSession) ReadyToStart() (bool, error) {
	return _Bidding.Contract.ReadyToStart(&_Bidding.CallOpts)
}

// ReadyToStart is a free data retrieval call binding the contract method 0x304b39a4.
//
// Solidity: function readyToStart() view returns(bool)
func (_Bidding *BiddingCallerSession) ReadyToStart() (bool, error) {
	return _Bidding.Contract.ReadyToStart(&_Bidding.CallOpts)
}

// RewardCollected is a free data retrieval call binding the contract method 0xb2ea78a8.
//
// Solidity: function rewardCollected() view returns(uint256)
func (_Bidding *BiddingCaller) RewardCollected(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bidding.contract.Call(opts, &out, "rewardCollected")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardCollected is a free data retrieval call binding the contract method 0xb2ea78a8.
//
// Solidity: function rewardCollected() view returns(uint256)
func (_Bidding *BiddingSession) RewardCollected() (*big.Int, error) {
	return _Bidding.Contract.RewardCollected(&_Bidding.CallOpts)
}

// RewardCollected is a free data retrieval call binding the contract method 0xb2ea78a8.
//
// Solidity: function rewardCollected() view returns(uint256)
func (_Bidding *BiddingCallerSession) RewardCollected() (*big.Int, error) {
	return _Bidding.Contract.RewardCollected(&_Bidding.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() view returns(uint256)
func (_Bidding *BiddingCaller) Start(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bidding.contract.Call(opts, &out, "start")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() view returns(uint256)
func (_Bidding *BiddingSession) Start() (*big.Int, error) {
	return _Bidding.Contract.Start(&_Bidding.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() view returns(uint256)
func (_Bidding *BiddingCallerSession) Start() (*big.Int, error) {
	return _Bidding.Contract.Start(&_Bidding.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Bidding *BiddingTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bidding.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Bidding *BiddingSession) Bid() (*types.Transaction, error) {
	return _Bidding.Contract.Bid(&_Bidding.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Bidding *BiddingTransactorSession) Bid() (*types.Transaction, error) {
	return _Bidding.Contract.Bid(&_Bidding.TransactOpts)
}

<<<<<<< HEAD
// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Bidding *BiddingTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bidding.contract.Transact(opts, "claimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Bidding *BiddingSession) ClaimReward() (*types.Transaction, error) {
	return _Bidding.Contract.ClaimReward(&_Bidding.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Bidding *BiddingTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _Bidding.Contract.ClaimReward(&_Bidding.TransactOpts)
=======
// DistributeReward is a paid mutator transaction binding the contract method 0x8f73c5ae.
//
// Solidity: function distributeReward() returns()
func (_Bidding *BiddingTransactor) DistributeReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bidding.contract.Transact(opts, "distributeReward")
}

// DistributeReward is a paid mutator transaction binding the contract method 0x8f73c5ae.
//
// Solidity: function distributeReward() returns()
func (_Bidding *BiddingSession) DistributeReward() (*types.Transaction, error) {
	return _Bidding.Contract.DistributeReward(&_Bidding.TransactOpts)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x8f73c5ae.
//
// Solidity: function distributeReward() returns()
func (_Bidding *BiddingTransactorSession) DistributeReward() (*types.Transaction, error) {
	return _Bidding.Contract.DistributeReward(&_Bidding.TransactOpts)
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
}

// StartBidding is a paid mutator transaction binding the contract method 0x7d7b39e4.
//
// Solidity: function startBidding() returns()
func (_Bidding *BiddingTransactor) StartBidding(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bidding.contract.Transact(opts, "startBidding")
}

// StartBidding is a paid mutator transaction binding the contract method 0x7d7b39e4.
//
// Solidity: function startBidding() returns()
func (_Bidding *BiddingSession) StartBidding() (*types.Transaction, error) {
	return _Bidding.Contract.StartBidding(&_Bidding.TransactOpts)
}

// StartBidding is a paid mutator transaction binding the contract method 0x7d7b39e4.
//
// Solidity: function startBidding() returns()
func (_Bidding *BiddingTransactorSession) StartBidding() (*types.Transaction, error) {
	return _Bidding.Contract.StartBidding(&_Bidding.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newAdmin) returns()
func (_Bidding *BiddingTransactor) TransferOwnership(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Bidding.contract.Transact(opts, "transferOwnership", newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newAdmin) returns()
func (_Bidding *BiddingSession) TransferOwnership(newAdmin common.Address) (*types.Transaction, error) {
	return _Bidding.Contract.TransferOwnership(&_Bidding.TransactOpts, newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newAdmin) returns()
func (_Bidding *BiddingTransactorSession) TransferOwnership(newAdmin common.Address) (*types.Transaction, error) {
	return _Bidding.Contract.TransferOwnership(&_Bidding.TransactOpts, newAdmin)
}

// WithdrawCommission is a paid mutator transaction binding the contract method 0x3e25e837.
//
// Solidity: function withdrawCommission() returns()
func (_Bidding *BiddingTransactor) WithdrawCommission(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bidding.contract.Transact(opts, "withdrawCommission")
}

// WithdrawCommission is a paid mutator transaction binding the contract method 0x3e25e837.
//
// Solidity: function withdrawCommission() returns()
func (_Bidding *BiddingSession) WithdrawCommission() (*types.Transaction, error) {
	return _Bidding.Contract.WithdrawCommission(&_Bidding.TransactOpts)
}

// WithdrawCommission is a paid mutator transaction binding the contract method 0x3e25e837.
//
// Solidity: function withdrawCommission() returns()
func (_Bidding *BiddingTransactorSession) WithdrawCommission() (*types.Transaction, error) {
	return _Bidding.Contract.WithdrawCommission(&_Bidding.TransactOpts)
}

// BiddingBidStartedIterator is returned from FilterBidStarted and is used to iterate over the raw logs and unpacked data for BidStarted events raised by the Bidding contract.
type BiddingBidStartedIterator struct {
	Event *BiddingBidStarted // Event containing the contract specifics and raw log

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
func (it *BiddingBidStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingBidStarted)
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
		it.Event = new(BiddingBidStarted)
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
func (it *BiddingBidStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingBidStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingBidStarted represents a BidStarted event raised by the Bidding contract.
type BiddingBidStarted struct {
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBidStarted is a free log retrieval operation binding the contract event 0xd135e0bd5bdb81c2daf4dbc71d27206a802fefb752227d9c455a44d1d2225a3e.
//
// Solidity: event bidStarted(uint256 startTime)
func (_Bidding *BiddingFilterer) FilterBidStarted(opts *bind.FilterOpts) (*BiddingBidStartedIterator, error) {

	logs, sub, err := _Bidding.contract.FilterLogs(opts, "bidStarted")
	if err != nil {
		return nil, err
	}
	return &BiddingBidStartedIterator{contract: _Bidding.contract, event: "bidStarted", logs: logs, sub: sub}, nil
}

// WatchBidStarted is a free log subscription operation binding the contract event 0xd135e0bd5bdb81c2daf4dbc71d27206a802fefb752227d9c455a44d1d2225a3e.
//
// Solidity: event bidStarted(uint256 startTime)
func (_Bidding *BiddingFilterer) WatchBidStarted(opts *bind.WatchOpts, sink chan<- *BiddingBidStarted) (event.Subscription, error) {

	logs, sub, err := _Bidding.contract.WatchLogs(opts, "bidStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingBidStarted)
				if err := _Bidding.contract.UnpackLog(event, "bidStarted", log); err != nil {
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

// ParseBidStarted is a log parse operation binding the contract event 0xd135e0bd5bdb81c2daf4dbc71d27206a802fefb752227d9c455a44d1d2225a3e.
//
// Solidity: event bidStarted(uint256 startTime)
func (_Bidding *BiddingFilterer) ParseBidStarted(log types.Log) (*BiddingBidStarted, error) {
	event := new(BiddingBidStarted)
	if err := _Bidding.contract.UnpackLog(event, "bidStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BiddingCommissionWithdrawnIterator is returned from FilterCommissionWithdrawn and is used to iterate over the raw logs and unpacked data for CommissionWithdrawn events raised by the Bidding contract.
type BiddingCommissionWithdrawnIterator struct {
	Event *BiddingCommissionWithdrawn // Event containing the contract specifics and raw log

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
func (it *BiddingCommissionWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingCommissionWithdrawn)
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
		it.Event = new(BiddingCommissionWithdrawn)
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
func (it *BiddingCommissionWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingCommissionWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingCommissionWithdrawn represents a CommissionWithdrawn event raised by the Bidding contract.
type BiddingCommissionWithdrawn struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCommissionWithdrawn is a free log retrieval operation binding the contract event 0x69b10852677cc2903b52fece19ad362b19e82627b217979720053501a31c8243.
//
// Solidity: event commissionWithdrawn()
func (_Bidding *BiddingFilterer) FilterCommissionWithdrawn(opts *bind.FilterOpts) (*BiddingCommissionWithdrawnIterator, error) {

	logs, sub, err := _Bidding.contract.FilterLogs(opts, "commissionWithdrawn")
	if err != nil {
		return nil, err
	}
	return &BiddingCommissionWithdrawnIterator{contract: _Bidding.contract, event: "commissionWithdrawn", logs: logs, sub: sub}, nil
}

// WatchCommissionWithdrawn is a free log subscription operation binding the contract event 0x69b10852677cc2903b52fece19ad362b19e82627b217979720053501a31c8243.
//
// Solidity: event commissionWithdrawn()
func (_Bidding *BiddingFilterer) WatchCommissionWithdrawn(opts *bind.WatchOpts, sink chan<- *BiddingCommissionWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Bidding.contract.WatchLogs(opts, "commissionWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingCommissionWithdrawn)
				if err := _Bidding.contract.UnpackLog(event, "commissionWithdrawn", log); err != nil {
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

// ParseCommissionWithdrawn is a log parse operation binding the contract event 0x69b10852677cc2903b52fece19ad362b19e82627b217979720053501a31c8243.
//
// Solidity: event commissionWithdrawn()
func (_Bidding *BiddingFilterer) ParseCommissionWithdrawn(log types.Log) (*BiddingCommissionWithdrawn, error) {
	event := new(BiddingCommissionWithdrawn)
	if err := _Bidding.contract.UnpackLog(event, "commissionWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BiddingNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the Bidding contract.
type BiddingNewBidIterator struct {
	Event *BiddingNewBid // Event containing the contract specifics and raw log

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
func (it *BiddingNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingNewBid)
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
		it.Event = new(BiddingNewBid)
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
func (it *BiddingNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingNewBid represents a NewBid event raised by the Bidding contract.
type BiddingNewBid struct {
	User     common.Address
	BidValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0x7b6bc441ac62da4ede87e2fafbf3885745a385ec42ec7611d5d8094e7b850be1.
//
// Solidity: event newBid(address user, uint256 bidValue)
func (_Bidding *BiddingFilterer) FilterNewBid(opts *bind.FilterOpts) (*BiddingNewBidIterator, error) {

	logs, sub, err := _Bidding.contract.FilterLogs(opts, "newBid")
	if err != nil {
		return nil, err
	}
	return &BiddingNewBidIterator{contract: _Bidding.contract, event: "newBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0x7b6bc441ac62da4ede87e2fafbf3885745a385ec42ec7611d5d8094e7b850be1.
//
// Solidity: event newBid(address user, uint256 bidValue)
func (_Bidding *BiddingFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *BiddingNewBid) (event.Subscription, error) {

	logs, sub, err := _Bidding.contract.WatchLogs(opts, "newBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingNewBid)
				if err := _Bidding.contract.UnpackLog(event, "newBid", log); err != nil {
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

// ParseNewBid is a log parse operation binding the contract event 0x7b6bc441ac62da4ede87e2fafbf3885745a385ec42ec7611d5d8094e7b850be1.
//
// Solidity: event newBid(address user, uint256 bidValue)
func (_Bidding *BiddingFilterer) ParseNewBid(log types.Log) (*BiddingNewBid, error) {
	event := new(BiddingNewBid)
	if err := _Bidding.contract.UnpackLog(event, "newBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

<<<<<<< HEAD
// BiddingRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Bidding contract.
type BiddingRewardClaimedIterator struct {
	Event *BiddingRewardClaimed // Event containing the contract specifics and raw log
=======
// BiddingRewardDistributedIterator is returned from FilterRewardDistributed and is used to iterate over the raw logs and unpacked data for RewardDistributed events raised by the Bidding contract.
type BiddingRewardDistributedIterator struct {
	Event *BiddingRewardDistributed // Event containing the contract specifics and raw log
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878

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
<<<<<<< HEAD
func (it *BiddingRewardClaimedIterator) Next() bool {
=======
func (it *BiddingRewardDistributedIterator) Next() bool {
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
<<<<<<< HEAD
			it.Event = new(BiddingRewardClaimed)
=======
			it.Event = new(BiddingRewardDistributed)
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
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
<<<<<<< HEAD
		it.Event = new(BiddingRewardClaimed)
=======
		it.Event = new(BiddingRewardDistributed)
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
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
<<<<<<< HEAD
func (it *BiddingRewardClaimedIterator) Error() error {
=======
func (it *BiddingRewardDistributedIterator) Error() error {
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
<<<<<<< HEAD
func (it *BiddingRewardClaimedIterator) Close() error {
=======
func (it *BiddingRewardDistributedIterator) Close() error {
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
	it.sub.Unsubscribe()
	return nil
}

<<<<<<< HEAD
// BiddingRewardClaimed represents a RewardClaimed event raised by the Bidding contract.
type BiddingRewardClaimed struct {
=======
// BiddingRewardDistributed represents a RewardDistributed event raised by the Bidding contract.
type BiddingRewardDistributed struct {
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
	Winner common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

<<<<<<< HEAD
// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x83b78188b13346b2ffb484da70d42ee27de7fbf9f2bd8045269e10ed643ccd76.
//
// Solidity: event rewardClaimed(address winner, uint256 reward)
func (_Bidding *BiddingFilterer) FilterRewardClaimed(opts *bind.FilterOpts) (*BiddingRewardClaimedIterator, error) {

	logs, sub, err := _Bidding.contract.FilterLogs(opts, "rewardClaimed")
	if err != nil {
		return nil, err
	}
	return &BiddingRewardClaimedIterator{contract: _Bidding.contract, event: "rewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x83b78188b13346b2ffb484da70d42ee27de7fbf9f2bd8045269e10ed643ccd76.
//
// Solidity: event rewardClaimed(address winner, uint256 reward)
func (_Bidding *BiddingFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *BiddingRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _Bidding.contract.WatchLogs(opts, "rewardClaimed")
=======
// FilterRewardDistributed is a free log retrieval operation binding the contract event 0xefa60ee4747682bd8de7a57958e41685ad2342d7be0f49cba9f08bac04d5d30d.
//
// Solidity: event rewardDistributed(address winner, uint256 reward)
func (_Bidding *BiddingFilterer) FilterRewardDistributed(opts *bind.FilterOpts) (*BiddingRewardDistributedIterator, error) {

	logs, sub, err := _Bidding.contract.FilterLogs(opts, "rewardDistributed")
	if err != nil {
		return nil, err
	}
	return &BiddingRewardDistributedIterator{contract: _Bidding.contract, event: "rewardDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardDistributed is a free log subscription operation binding the contract event 0xefa60ee4747682bd8de7a57958e41685ad2342d7be0f49cba9f08bac04d5d30d.
//
// Solidity: event rewardDistributed(address winner, uint256 reward)
func (_Bidding *BiddingFilterer) WatchRewardDistributed(opts *bind.WatchOpts, sink chan<- *BiddingRewardDistributed) (event.Subscription, error) {

	logs, sub, err := _Bidding.contract.WatchLogs(opts, "rewardDistributed")
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
<<<<<<< HEAD
				event := new(BiddingRewardClaimed)
				if err := _Bidding.contract.UnpackLog(event, "rewardClaimed", log); err != nil {
=======
				event := new(BiddingRewardDistributed)
				if err := _Bidding.contract.UnpackLog(event, "rewardDistributed", log); err != nil {
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
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

<<<<<<< HEAD
// ParseRewardClaimed is a log parse operation binding the contract event 0x83b78188b13346b2ffb484da70d42ee27de7fbf9f2bd8045269e10ed643ccd76.
//
// Solidity: event rewardClaimed(address winner, uint256 reward)
func (_Bidding *BiddingFilterer) ParseRewardClaimed(log types.Log) (*BiddingRewardClaimed, error) {
	event := new(BiddingRewardClaimed)
	if err := _Bidding.contract.UnpackLog(event, "rewardClaimed", log); err != nil {
=======
// ParseRewardDistributed is a log parse operation binding the contract event 0xefa60ee4747682bd8de7a57958e41685ad2342d7be0f49cba9f08bac04d5d30d.
//
// Solidity: event rewardDistributed(address winner, uint256 reward)
func (_Bidding *BiddingFilterer) ParseRewardDistributed(log types.Log) (*BiddingRewardDistributed, error) {
	event := new(BiddingRewardDistributed)
	if err := _Bidding.contract.UnpackLog(event, "rewardDistributed", log); err != nil {
>>>>>>> d7fc605da95ee08a2f11e4ed0c66aee8b42ab878
		return nil, err
	}
	event.Raw = log
	return event, nil
}
