// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// AETHERMetaData contains all meta data concerning the AETHER contract.
var AETHERMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"FileCIDRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"registerFileCID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getFileCID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AETHERABI is the input ABI used to generate the binding from.
// Deprecated: Use AETHERMetaData.ABI instead.
var AETHERABI = AETHERMetaData.ABI

// AETHER is an auto generated Go binding around an Ethereum contract.
type AETHER struct {
	AETHERCaller     // Read-only binding to the contract
	AETHERTransactor // Write-only binding to the contract
	AETHERFilterer   // Log filterer for contract events
}

// AETHERCaller is an auto generated read-only Go binding around an Ethereum contract.
type AETHERCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AETHERTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AETHERTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AETHERFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AETHERFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AETHERSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AETHERSession struct {
	Contract     *AETHER           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AETHERCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AETHERCallerSession struct {
	Contract *AETHERCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AETHERTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AETHERTransactorSession struct {
	Contract     *AETHERTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AETHERRaw is an auto generated low-level Go binding around an Ethereum contract.
type AETHERRaw struct {
	Contract *AETHER // Generic contract binding to access the raw methods on
}

// AETHERCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AETHERCallerRaw struct {
	Contract *AETHERCaller // Generic read-only contract binding to access the raw methods on
}

// AETHERTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AETHERTransactorRaw struct {
	Contract *AETHERTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAETHER creates a new instance of AETHER, bound to a specific deployed contract.
func NewAETHER(address common.Address, backend bind.ContractBackend) (*AETHER, error) {
	contract, err := bindAETHER(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AETHER{AETHERCaller: AETHERCaller{contract: contract}, AETHERTransactor: AETHERTransactor{contract: contract}, AETHERFilterer: AETHERFilterer{contract: contract}}, nil
}

// NewAETHERCaller creates a new read-only instance of AETHER, bound to a specific deployed contract.
func NewAETHERCaller(address common.Address, caller bind.ContractCaller) (*AETHERCaller, error) {
	contract, err := bindAETHER(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AETHERCaller{contract: contract}, nil
}

// NewAETHERTransactor creates a new write-only instance of AETHER, bound to a specific deployed contract.
func NewAETHERTransactor(address common.Address, transactor bind.ContractTransactor) (*AETHERTransactor, error) {
	contract, err := bindAETHER(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AETHERTransactor{contract: contract}, nil
}

// NewAETHERFilterer creates a new log filterer instance of AETHER, bound to a specific deployed contract.
func NewAETHERFilterer(address common.Address, filterer bind.ContractFilterer) (*AETHERFilterer, error) {
	contract, err := bindAETHER(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AETHERFilterer{contract: contract}, nil
}

// bindAETHER binds a generic wrapper to an already deployed contract.
func bindAETHER(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AETHERMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AETHER *AETHERRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AETHER.Contract.AETHERCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AETHER *AETHERRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AETHER.Contract.AETHERTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AETHER *AETHERRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AETHER.Contract.AETHERTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AETHER *AETHERCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AETHER.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AETHER *AETHERTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AETHER.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AETHER *AETHERTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AETHER.Contract.contract.Transact(opts, method, params...)
}

// GetFileCID is a free data retrieval call binding the contract method 0xa8ecc290.
//
// Solidity: function getFileCID(string fileName, address user) view returns(string)
func (_AETHER *AETHERCaller) GetFileCID(opts *bind.CallOpts, fileName string, user common.Address) (string, error) {
	var out []interface{}
	err := _AETHER.contract.Call(opts, &out, "getFileCID", fileName, user)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFileCID is a free data retrieval call binding the contract method 0xa8ecc290.
//
// Solidity: function getFileCID(string fileName, address user) view returns(string)
func (_AETHER *AETHERSession) GetFileCID(fileName string, user common.Address) (string, error) {
	return _AETHER.Contract.GetFileCID(&_AETHER.CallOpts, fileName, user)
}

// GetFileCID is a free data retrieval call binding the contract method 0xa8ecc290.
//
// Solidity: function getFileCID(string fileName, address user) view returns(string)
func (_AETHER *AETHERCallerSession) GetFileCID(fileName string, user common.Address) (string, error) {
	return _AETHER.Contract.GetFileCID(&_AETHER.CallOpts, fileName, user)
}

// GetUserAddress is a free data retrieval call binding the contract method 0x4985e85c.
//
// Solidity: function getUserAddress(string name) view returns(address)
func (_AETHER *AETHERCaller) GetUserAddress(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _AETHER.contract.Call(opts, &out, "getUserAddress", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserAddress is a free data retrieval call binding the contract method 0x4985e85c.
//
// Solidity: function getUserAddress(string name) view returns(address)
func (_AETHER *AETHERSession) GetUserAddress(name string) (common.Address, error) {
	return _AETHER.Contract.GetUserAddress(&_AETHER.CallOpts, name)
}

// GetUserAddress is a free data retrieval call binding the contract method 0x4985e85c.
//
// Solidity: function getUserAddress(string name) view returns(address)
func (_AETHER *AETHERCallerSession) GetUserAddress(name string) (common.Address, error) {
	return _AETHER.Contract.GetUserAddress(&_AETHER.CallOpts, name)
}

// GetUserName is a free data retrieval call binding the contract method 0xd84f55ee.
//
// Solidity: function getUserName(address user) view returns(string)
func (_AETHER *AETHERCaller) GetUserName(opts *bind.CallOpts, user common.Address) (string, error) {
	var out []interface{}
	err := _AETHER.contract.Call(opts, &out, "getUserName", user)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUserName is a free data retrieval call binding the contract method 0xd84f55ee.
//
// Solidity: function getUserName(address user) view returns(string)
func (_AETHER *AETHERSession) GetUserName(user common.Address) (string, error) {
	return _AETHER.Contract.GetUserName(&_AETHER.CallOpts, user)
}

// GetUserName is a free data retrieval call binding the contract method 0xd84f55ee.
//
// Solidity: function getUserName(address user) view returns(string)
func (_AETHER *AETHERCallerSession) GetUserName(user common.Address) (string, error) {
	return _AETHER.Contract.GetUserName(&_AETHER.CallOpts, user)
}

// RegisterFileCID is a paid mutator transaction binding the contract method 0xb4550079.
//
// Solidity: function registerFileCID(string fileName, string cid) returns()
func (_AETHER *AETHERTransactor) RegisterFileCID(opts *bind.TransactOpts, fileName string, cid string) (*types.Transaction, error) {
	return _AETHER.contract.Transact(opts, "registerFileCID", fileName, cid)
}

// RegisterFileCID is a paid mutator transaction binding the contract method 0xb4550079.
//
// Solidity: function registerFileCID(string fileName, string cid) returns()
func (_AETHER *AETHERSession) RegisterFileCID(fileName string, cid string) (*types.Transaction, error) {
	return _AETHER.Contract.RegisterFileCID(&_AETHER.TransactOpts, fileName, cid)
}

// RegisterFileCID is a paid mutator transaction binding the contract method 0xb4550079.
//
// Solidity: function registerFileCID(string fileName, string cid) returns()
func (_AETHER *AETHERTransactorSession) RegisterFileCID(fileName string, cid string) (*types.Transaction, error) {
	return _AETHER.Contract.RegisterFileCID(&_AETHER.TransactOpts, fileName, cid)
}

// RegisterName is a paid mutator transaction binding the contract method 0x0830602b.
//
// Solidity: function registerName(string name) returns()
func (_AETHER *AETHERTransactor) RegisterName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _AETHER.contract.Transact(opts, "registerName", name)
}

// RegisterName is a paid mutator transaction binding the contract method 0x0830602b.
//
// Solidity: function registerName(string name) returns()
func (_AETHER *AETHERSession) RegisterName(name string) (*types.Transaction, error) {
	return _AETHER.Contract.RegisterName(&_AETHER.TransactOpts, name)
}

// RegisterName is a paid mutator transaction binding the contract method 0x0830602b.
//
// Solidity: function registerName(string name) returns()
func (_AETHER *AETHERTransactorSession) RegisterName(name string) (*types.Transaction, error) {
	return _AETHER.Contract.RegisterName(&_AETHER.TransactOpts, name)
}

// AETHERFileCIDRegisteredIterator is returned from FilterFileCIDRegistered and is used to iterate over the raw logs and unpacked data for FileCIDRegistered events raised by the AETHER contract.
type AETHERFileCIDRegisteredIterator struct {
	Event *AETHERFileCIDRegistered // Event containing the contract specifics and raw log

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
func (it *AETHERFileCIDRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AETHERFileCIDRegistered)
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
		it.Event = new(AETHERFileCIDRegistered)
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
func (it *AETHERFileCIDRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AETHERFileCIDRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AETHERFileCIDRegistered represents a FileCIDRegistered event raised by the AETHER contract.
type AETHERFileCIDRegistered struct {
	User     common.Address
	FileName string
	Cid      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFileCIDRegistered is a free log retrieval operation binding the contract event 0xeec69f53240a97723077f725d7a08f2605d098ea27a6d2f0296953dcf2a06d6a.
//
// Solidity: event FileCIDRegistered(address indexed user, string fileName, string cid)
func (_AETHER *AETHERFilterer) FilterFileCIDRegistered(opts *bind.FilterOpts, user []common.Address) (*AETHERFileCIDRegisteredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AETHER.contract.FilterLogs(opts, "FileCIDRegistered", userRule)
	if err != nil {
		return nil, err
	}
	return &AETHERFileCIDRegisteredIterator{contract: _AETHER.contract, event: "FileCIDRegistered", logs: logs, sub: sub}, nil
}

// WatchFileCIDRegistered is a free log subscription operation binding the contract event 0xeec69f53240a97723077f725d7a08f2605d098ea27a6d2f0296953dcf2a06d6a.
//
// Solidity: event FileCIDRegistered(address indexed user, string fileName, string cid)
func (_AETHER *AETHERFilterer) WatchFileCIDRegistered(opts *bind.WatchOpts, sink chan<- *AETHERFileCIDRegistered, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AETHER.contract.WatchLogs(opts, "FileCIDRegistered", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AETHERFileCIDRegistered)
				if err := _AETHER.contract.UnpackLog(event, "FileCIDRegistered", log); err != nil {
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

// ParseFileCIDRegistered is a log parse operation binding the contract event 0xeec69f53240a97723077f725d7a08f2605d098ea27a6d2f0296953dcf2a06d6a.
//
// Solidity: event FileCIDRegistered(address indexed user, string fileName, string cid)
func (_AETHER *AETHERFilterer) ParseFileCIDRegistered(log types.Log) (*AETHERFileCIDRegistered, error) {
	event := new(AETHERFileCIDRegistered)
	if err := _AETHER.contract.UnpackLog(event, "FileCIDRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AETHERNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the AETHER contract.
type AETHERNameRegisteredIterator struct {
	Event *AETHERNameRegistered // Event containing the contract specifics and raw log

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
func (it *AETHERNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AETHERNameRegistered)
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
		it.Event = new(AETHERNameRegistered)
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
func (it *AETHERNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AETHERNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AETHERNameRegistered represents a NameRegistered event raised by the AETHER contract.
type AETHERNameRegistered struct {
	User common.Address
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0x56a1a2764a85c689329fbc04f299b64690a26b1dc5ac82e903cf8ad805cd5396.
//
// Solidity: event NameRegistered(address indexed user, string name)
func (_AETHER *AETHERFilterer) FilterNameRegistered(opts *bind.FilterOpts, user []common.Address) (*AETHERNameRegisteredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AETHER.contract.FilterLogs(opts, "NameRegistered", userRule)
	if err != nil {
		return nil, err
	}
	return &AETHERNameRegisteredIterator{contract: _AETHER.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0x56a1a2764a85c689329fbc04f299b64690a26b1dc5ac82e903cf8ad805cd5396.
//
// Solidity: event NameRegistered(address indexed user, string name)
func (_AETHER *AETHERFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *AETHERNameRegistered, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AETHER.contract.WatchLogs(opts, "NameRegistered", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AETHERNameRegistered)
				if err := _AETHER.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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

// ParseNameRegistered is a log parse operation binding the contract event 0x56a1a2764a85c689329fbc04f299b64690a26b1dc5ac82e903cf8ad805cd5396.
//
// Solidity: event NameRegistered(address indexed user, string name)
func (_AETHER *AETHERFilterer) ParseNameRegistered(log types.Log) (*AETHERNameRegistered, error) {
	event := new(AETHERNameRegistered)
	if err := _AETHER.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
