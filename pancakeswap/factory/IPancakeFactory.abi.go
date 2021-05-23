// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PancakeswapABI is the input ABI used to generate the binding from.
const PancakeswapABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Pancakeswap is an auto generated Go binding around an Ethereum contract.
type Pancakeswap struct {
	PancakeswapCaller     // Read-only binding to the contract
	PancakeswapTransactor // Write-only binding to the contract
	PancakeswapFilterer   // Log filterer for contract events
}

// PancakeswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeswapSession struct {
	Contract     *Pancakeswap      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeswapCallerSession struct {
	Contract *PancakeswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PancakeswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeswapTransactorSession struct {
	Contract     *PancakeswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PancakeswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeswapRaw struct {
	Contract *Pancakeswap // Generic contract binding to access the raw methods on
}

// PancakeswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeswapCallerRaw struct {
	Contract *PancakeswapCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeswapTransactorRaw struct {
	Contract *PancakeswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeswap creates a new instance of Pancakeswap, bound to a specific deployed contract.
func NewPancakeswap(address common.Address, backend bind.ContractBackend) (*Pancakeswap, error) {
	contract, err := bindPancakeswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pancakeswap{PancakeswapCaller: PancakeswapCaller{contract: contract}, PancakeswapTransactor: PancakeswapTransactor{contract: contract}, PancakeswapFilterer: PancakeswapFilterer{contract: contract}}, nil
}

// NewPancakeswapCaller creates a new read-only instance of Pancakeswap, bound to a specific deployed contract.
func NewPancakeswapCaller(address common.Address, caller bind.ContractCaller) (*PancakeswapCaller, error) {
	contract, err := bindPancakeswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeswapCaller{contract: contract}, nil
}

// NewPancakeswapTransactor creates a new write-only instance of Pancakeswap, bound to a specific deployed contract.
func NewPancakeswapTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeswapTransactor, error) {
	contract, err := bindPancakeswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeswapTransactor{contract: contract}, nil
}

// NewPancakeswapFilterer creates a new log filterer instance of Pancakeswap, bound to a specific deployed contract.
func NewPancakeswapFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeswapFilterer, error) {
	contract, err := bindPancakeswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeswapFilterer{contract: contract}, nil
}

// bindPancakeswap binds a generic wrapper to an already deployed contract.
func bindPancakeswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PancakeswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pancakeswap *PancakeswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pancakeswap.Contract.PancakeswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pancakeswap *PancakeswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancakeswap.Contract.PancakeswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pancakeswap *PancakeswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pancakeswap.Contract.PancakeswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pancakeswap *PancakeswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pancakeswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pancakeswap *PancakeswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancakeswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pancakeswap *PancakeswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pancakeswap.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_Pancakeswap *PancakeswapCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Pancakeswap.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_Pancakeswap *PancakeswapSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Pancakeswap.Contract.AllPairs(&_Pancakeswap.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_Pancakeswap *PancakeswapCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Pancakeswap.Contract.AllPairs(&_Pancakeswap.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Pancakeswap *PancakeswapCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancakeswap.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Pancakeswap *PancakeswapSession) AllPairsLength() (*big.Int, error) {
	return _Pancakeswap.Contract.AllPairsLength(&_Pancakeswap.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Pancakeswap *PancakeswapCallerSession) AllPairsLength() (*big.Int, error) {
	return _Pancakeswap.Contract.AllPairsLength(&_Pancakeswap.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Pancakeswap *PancakeswapCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancakeswap.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Pancakeswap *PancakeswapSession) FeeTo() (common.Address, error) {
	return _Pancakeswap.Contract.FeeTo(&_Pancakeswap.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Pancakeswap *PancakeswapCallerSession) FeeTo() (common.Address, error) {
	return _Pancakeswap.Contract.FeeTo(&_Pancakeswap.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Pancakeswap *PancakeswapCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancakeswap.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Pancakeswap *PancakeswapSession) FeeToSetter() (common.Address, error) {
	return _Pancakeswap.Contract.FeeToSetter(&_Pancakeswap.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Pancakeswap *PancakeswapCallerSession) FeeToSetter() (common.Address, error) {
	return _Pancakeswap.Contract.FeeToSetter(&_Pancakeswap.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Pancakeswap *PancakeswapCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _Pancakeswap.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Pancakeswap *PancakeswapSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Pancakeswap.Contract.GetPair(&_Pancakeswap.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Pancakeswap *PancakeswapCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Pancakeswap.Contract.GetPair(&_Pancakeswap.CallOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Pancakeswap *PancakeswapTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Pancakeswap.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Pancakeswap *PancakeswapSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Pancakeswap.Contract.CreatePair(&_Pancakeswap.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Pancakeswap *PancakeswapTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Pancakeswap.Contract.CreatePair(&_Pancakeswap.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_Pancakeswap *PancakeswapTransactor) SetFeeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Pancakeswap.contract.Transact(opts, "setFeeTo", arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_Pancakeswap *PancakeswapSession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _Pancakeswap.Contract.SetFeeTo(&_Pancakeswap.TransactOpts, arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_Pancakeswap *PancakeswapTransactorSession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _Pancakeswap.Contract.SetFeeTo(&_Pancakeswap.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_Pancakeswap *PancakeswapTransactor) SetFeeToSetter(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Pancakeswap.contract.Transact(opts, "setFeeToSetter", arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_Pancakeswap *PancakeswapSession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _Pancakeswap.Contract.SetFeeToSetter(&_Pancakeswap.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_Pancakeswap *PancakeswapTransactorSession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _Pancakeswap.Contract.SetFeeToSetter(&_Pancakeswap.TransactOpts, arg0)
}

// PancakeswapPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Pancakeswap contract.
type PancakeswapPairCreatedIterator struct {
	Event *PancakeswapPairCreated // Event containing the contract specifics and raw log

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
func (it *PancakeswapPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeswapPairCreated)
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
		it.Event = new(PancakeswapPairCreated)
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
func (it *PancakeswapPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeswapPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeswapPairCreated represents a PairCreated event raised by the Pancakeswap contract.
type PancakeswapPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Pancakeswap *PancakeswapFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*PancakeswapPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Pancakeswap.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &PancakeswapPairCreatedIterator{contract: _Pancakeswap.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Pancakeswap *PancakeswapFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *PancakeswapPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Pancakeswap.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeswapPairCreated)
				if err := _Pancakeswap.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Pancakeswap *PancakeswapFilterer) ParsePairCreated(log types.Log) (*PancakeswapPairCreated, error) {
	event := new(PancakeswapPairCreated)
	if err := _Pancakeswap.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
