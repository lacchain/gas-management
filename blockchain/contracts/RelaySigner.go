// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relay

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

// RelayABI is the input ABI used to generate the binding from.
const RelayABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"GasLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"name\":\"GasUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"Hashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Relayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"charge\",\"type\":\"uint256\"}],\"name\":\"TransactionRelayed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedFunction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"relayMetaTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"setGasLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Relay is an auto generated Go binding around an Ethereum contract.
type Relay struct {
	RelayCaller     // Read-only binding to the contract
	RelayTransactor // Write-only binding to the contract
	RelayFilterer   // Log filterer for contract events
}

// RelayCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelaySession struct {
	Contract     *Relay            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayCallerSession struct {
	Contract *RelayCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RelayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayTransactorSession struct {
	Contract     *RelayTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayRaw struct {
	Contract *Relay // Generic contract binding to access the raw methods on
}

// RelayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayCallerRaw struct {
	Contract *RelayCaller // Generic read-only contract binding to access the raw methods on
}

// RelayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayTransactorRaw struct {
	Contract *RelayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelay creates a new instance of Relay, bound to a specific deployed contract.
func NewRelay(address common.Address, backend bind.ContractBackend) (*Relay, error) {
	contract, err := bindRelay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relay{RelayCaller: RelayCaller{contract: contract}, RelayTransactor: RelayTransactor{contract: contract}, RelayFilterer: RelayFilterer{contract: contract}}, nil
}

// NewRelayCaller creates a new read-only instance of Relay, bound to a specific deployed contract.
func NewRelayCaller(address common.Address, caller bind.ContractCaller) (*RelayCaller, error) {
	contract, err := bindRelay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayCaller{contract: contract}, nil
}

// NewRelayTransactor creates a new write-only instance of Relay, bound to a specific deployed contract.
func NewRelayTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayTransactor, error) {
	contract, err := bindRelay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayTransactor{contract: contract}, nil
}

// NewRelayFilterer creates a new log filterer instance of Relay, bound to a specific deployed contract.
func NewRelayFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayFilterer, error) {
	contract, err := bindRelay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayFilterer{contract: contract}, nil
}

// bindRelay binds a generic wrapper to an already deployed contract.
func bindRelay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relay *RelayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relay.Contract.RelayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relay *RelayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relay.Contract.RelayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relay *RelayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relay.Contract.RelayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relay *RelayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relay *RelayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relay *RelayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relay.Contract.contract.Transact(opts, method, params...)
}

// GetGasLimit is a free data retrieval call binding the contract method 0x1a93d1c3.
//
// Solidity: function getGasLimit() view returns(uint256)
func (_Relay *RelayCaller) GetGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getGasLimit")
	return *ret0, err
}

// GetGasLimit is a free data retrieval call binding the contract method 0x1a93d1c3.
//
// Solidity: function getGasLimit() view returns(uint256)
func (_Relay *RelaySession) GetGasLimit() (*big.Int, error) {
	return _Relay.Contract.GetGasLimit(&_Relay.CallOpts)
}

// GetGasLimit is a free data retrieval call binding the contract method 0x1a93d1c3.
//
// Solidity: function getGasLimit() view returns(uint256)
func (_Relay *RelayCallerSession) GetGasLimit() (*big.Int, error) {
	return _Relay.Contract.GetGasLimit(&_Relay.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_Relay *RelayCaller) GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getNonce", from)
	return *ret0, err
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_Relay *RelaySession) GetNonce(from common.Address) (*big.Int, error) {
	return _Relay.Contract.GetNonce(&_Relay.CallOpts, from)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_Relay *RelayCallerSession) GetNonce(from common.Address) (*big.Int, error) {
	return _Relay.Contract.GetNonce(&_Relay.CallOpts, from)
}

// RelayMetaTx is a paid mutator transaction binding the contract method 0x84950d4b.
//
// Solidity: function relayMetaTx(address from, address to, bytes encodedFunction, uint256 gasLimit, uint256 nonce, bytes signature) returns(bool success)
func (_Relay *RelayTransactor) RelayMetaTx(opts *bind.TransactOpts, from common.Address, to common.Address, encodedFunction []byte, gasLimit *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "relayMetaTx", from, to, encodedFunction, gasLimit, nonce, signature)
}

// RelayMetaTx is a paid mutator transaction binding the contract method 0x84950d4b.
//
// Solidity: function relayMetaTx(address from, address to, bytes encodedFunction, uint256 gasLimit, uint256 nonce, bytes signature) returns(bool success)
func (_Relay *RelaySession) RelayMetaTx(from common.Address, to common.Address, encodedFunction []byte, gasLimit *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Relay.Contract.RelayMetaTx(&_Relay.TransactOpts, from, to, encodedFunction, gasLimit, nonce, signature)
}

// RelayMetaTx is a paid mutator transaction binding the contract method 0x84950d4b.
//
// Solidity: function relayMetaTx(address from, address to, bytes encodedFunction, uint256 gasLimit, uint256 nonce, bytes signature) returns(bool success)
func (_Relay *RelayTransactorSession) RelayMetaTx(from common.Address, to common.Address, encodedFunction []byte, gasLimit *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Relay.Contract.RelayMetaTx(&_Relay.TransactOpts, from, to, encodedFunction, gasLimit, nonce, signature)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xc7c8459a.
//
// Solidity: function setGasLimit(address to, uint256 gasLimit) returns(bool)
func (_Relay *RelayTransactor) SetGasLimit(opts *bind.TransactOpts, to common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "setGasLimit", to, gasLimit)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xc7c8459a.
//
// Solidity: function setGasLimit(address to, uint256 gasLimit) returns(bool)
func (_Relay *RelaySession) SetGasLimit(to common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _Relay.Contract.SetGasLimit(&_Relay.TransactOpts, to, gasLimit)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xc7c8459a.
//
// Solidity: function setGasLimit(address to, uint256 gasLimit) returns(bool)
func (_Relay *RelayTransactorSession) SetGasLimit(to common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _Relay.Contract.SetGasLimit(&_Relay.TransactOpts, to, gasLimit)
}

// RelayGasLimitIterator is returned from FilterGasLimit and is used to iterate over the raw logs and unpacked data for GasLimit events raised by the Relay contract.
type RelayGasLimitIterator struct {
	Event *RelayGasLimit // Event containing the contract specifics and raw log

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
func (it *RelayGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayGasLimit)
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
		it.Event = new(RelayGasLimit)
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
func (it *RelayGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayGasLimit represents a GasLimit event raised by the Relay contract.
type RelayGasLimit struct {
	GasUsed  *big.Int
	GasLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGasLimit is a free log retrieval operation binding the contract event 0x1bdaae1a8d1138ae7f64124bed20c7c54e1ecc16b7c99aee91a8bdf3ede06405.
//
// Solidity: event GasLimit(uint256 gasUsed, uint256 gasLimit)
func (_Relay *RelayFilterer) FilterGasLimit(opts *bind.FilterOpts) (*RelayGasLimitIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "GasLimit")
	if err != nil {
		return nil, err
	}
	return &RelayGasLimitIterator{contract: _Relay.contract, event: "GasLimit", logs: logs, sub: sub}, nil
}

// WatchGasLimit is a free log subscription operation binding the contract event 0x1bdaae1a8d1138ae7f64124bed20c7c54e1ecc16b7c99aee91a8bdf3ede06405.
//
// Solidity: event GasLimit(uint256 gasUsed, uint256 gasLimit)
func (_Relay *RelayFilterer) WatchGasLimit(opts *bind.WatchOpts, sink chan<- *RelayGasLimit) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "GasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayGasLimit)
				if err := _Relay.contract.UnpackLog(event, "GasLimit", log); err != nil {
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

// ParseGasLimit is a log parse operation binding the contract event 0x1bdaae1a8d1138ae7f64124bed20c7c54e1ecc16b7c99aee91a8bdf3ede06405.
//
// Solidity: event GasLimit(uint256 gasUsed, uint256 gasLimit)
func (_Relay *RelayFilterer) ParseGasLimit(log types.Log) (*RelayGasLimit, error) {
	event := new(RelayGasLimit)
	if err := _Relay.contract.UnpackLog(event, "GasLimit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayGasUsedIterator is returned from FilterGasUsed and is used to iterate over the raw logs and unpacked data for GasUsed events raised by the Relay contract.
type RelayGasUsedIterator struct {
	Event *RelayGasUsed // Event containing the contract specifics and raw log

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
func (it *RelayGasUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayGasUsed)
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
		it.Event = new(RelayGasUsed)
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
func (it *RelayGasUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayGasUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayGasUsed represents a GasUsed event raised by the Relay contract.
type RelayGasUsed struct {
	GasUsed *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGasUsed is a free log retrieval operation binding the contract event 0x275c82ca4f491f0c9b7ffa4f6e3e85613c4a0d2c4f67b64995dc68806f0586e5.
//
// Solidity: event GasUsed(uint256 gasUsed)
func (_Relay *RelayFilterer) FilterGasUsed(opts *bind.FilterOpts) (*RelayGasUsedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "GasUsed")
	if err != nil {
		return nil, err
	}
	return &RelayGasUsedIterator{contract: _Relay.contract, event: "GasUsed", logs: logs, sub: sub}, nil
}

// WatchGasUsed is a free log subscription operation binding the contract event 0x275c82ca4f491f0c9b7ffa4f6e3e85613c4a0d2c4f67b64995dc68806f0586e5.
//
// Solidity: event GasUsed(uint256 gasUsed)
func (_Relay *RelayFilterer) WatchGasUsed(opts *bind.WatchOpts, sink chan<- *RelayGasUsed) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "GasUsed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayGasUsed)
				if err := _Relay.contract.UnpackLog(event, "GasUsed", log); err != nil {
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

// ParseGasUsed is a log parse operation binding the contract event 0x275c82ca4f491f0c9b7ffa4f6e3e85613c4a0d2c4f67b64995dc68806f0586e5.
//
// Solidity: event GasUsed(uint256 gasUsed)
func (_Relay *RelayFilterer) ParseGasUsed(log types.Log) (*RelayGasUsed, error) {
	event := new(RelayGasUsed)
	if err := _Relay.contract.UnpackLog(event, "GasUsed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayHashedIterator is returned from FilterHashed and is used to iterate over the raw logs and unpacked data for Hashed events raised by the Relay contract.
type RelayHashedIterator struct {
	Event *RelayHashed // Event containing the contract specifics and raw log

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
func (it *RelayHashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHashed)
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
		it.Event = new(RelayHashed)
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
func (it *RelayHashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHashed represents a Hashed event raised by the Relay contract.
type RelayHashed struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHashed is a free log retrieval operation binding the contract event 0x2d3a7f4267f4c2aeefe8806fd51e14a3d8154d64d231067e22e1ad59ccf53a00.
//
// Solidity: event Hashed(bytes32 hash)
func (_Relay *RelayFilterer) FilterHashed(opts *bind.FilterOpts) (*RelayHashedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "Hashed")
	if err != nil {
		return nil, err
	}
	return &RelayHashedIterator{contract: _Relay.contract, event: "Hashed", logs: logs, sub: sub}, nil
}

// WatchHashed is a free log subscription operation binding the contract event 0x2d3a7f4267f4c2aeefe8806fd51e14a3d8154d64d231067e22e1ad59ccf53a00.
//
// Solidity: event Hashed(bytes32 hash)
func (_Relay *RelayFilterer) WatchHashed(opts *bind.WatchOpts, sink chan<- *RelayHashed) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "Hashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHashed)
				if err := _Relay.contract.UnpackLog(event, "Hashed", log); err != nil {
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

// ParseHashed is a log parse operation binding the contract event 0x2d3a7f4267f4c2aeefe8806fd51e14a3d8154d64d231067e22e1ad59ccf53a00.
//
// Solidity: event Hashed(bytes32 hash)
func (_Relay *RelayFilterer) ParseHashed(log types.Log) (*RelayHashed, error) {
	event := new(RelayHashed)
	if err := _Relay.contract.UnpackLog(event, "Hashed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayRelayedIterator is returned from FilterRelayed and is used to iterate over the raw logs and unpacked data for Relayed events raised by the Relay contract.
type RelayRelayedIterator struct {
	Event *RelayRelayed // Event containing the contract specifics and raw log

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
func (it *RelayRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayRelayed)
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
		it.Event = new(RelayRelayed)
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
func (it *RelayRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayRelayed represents a Relayed event raised by the Relay contract.
type RelayRelayed struct {
	Sender common.Address
	From   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRelayed is a free log retrieval operation binding the contract event 0x79f72f9dacecfa9af3cfe946364971d0ef4826ffd35451658b283d58a382c20f.
//
// Solidity: event Relayed(address indexed sender, address indexed from)
func (_Relay *RelayFilterer) FilterRelayed(opts *bind.FilterOpts, sender []common.Address, from []common.Address) (*RelayRelayedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Relay.contract.FilterLogs(opts, "Relayed", senderRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &RelayRelayedIterator{contract: _Relay.contract, event: "Relayed", logs: logs, sub: sub}, nil
}

// WatchRelayed is a free log subscription operation binding the contract event 0x79f72f9dacecfa9af3cfe946364971d0ef4826ffd35451658b283d58a382c20f.
//
// Solidity: event Relayed(address indexed sender, address indexed from)
func (_Relay *RelayFilterer) WatchRelayed(opts *bind.WatchOpts, sink chan<- *RelayRelayed, sender []common.Address, from []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Relay.contract.WatchLogs(opts, "Relayed", senderRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayRelayed)
				if err := _Relay.contract.UnpackLog(event, "Relayed", log); err != nil {
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

// ParseRelayed is a log parse operation binding the contract event 0x79f72f9dacecfa9af3cfe946364971d0ef4826ffd35451658b283d58a382c20f.
//
// Solidity: event Relayed(address indexed sender, address indexed from)
func (_Relay *RelayFilterer) ParseRelayed(log types.Log) (*RelayRelayed, error) {
	event := new(RelayRelayed)
	if err := _Relay.contract.UnpackLog(event, "Relayed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayTransactionRelayedIterator is returned from FilterTransactionRelayed and is used to iterate over the raw logs and unpacked data for TransactionRelayed events raised by the Relay contract.
type RelayTransactionRelayedIterator struct {
	Event *RelayTransactionRelayed // Event containing the contract specifics and raw log

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
func (it *RelayTransactionRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayTransactionRelayed)
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
		it.Event = new(RelayTransactionRelayed)
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
func (it *RelayTransactionRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayTransactionRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayTransactionRelayed represents a TransactionRelayed event raised by the Relay contract.
type RelayTransactionRelayed struct {
	Relay    common.Address
	From     common.Address
	To       common.Address
	Selector [4]byte
	Charge   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransactionRelayed is a free log retrieval operation binding the contract event 0x2a8e164a2bd95ffaf59651ebd8e75e4e879d97cab53c11dde9dfdbcceb1e6780.
//
// Solidity: event TransactionRelayed(address indexed relay, address indexed from, address indexed to, bytes4 selector, uint256 charge)
func (_Relay *RelayFilterer) FilterTransactionRelayed(opts *bind.FilterOpts, relay []common.Address, from []common.Address, to []common.Address) (*RelayTransactionRelayedIterator, error) {

	var relayRule []interface{}
	for _, relayItem := range relay {
		relayRule = append(relayRule, relayItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Relay.contract.FilterLogs(opts, "TransactionRelayed", relayRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RelayTransactionRelayedIterator{contract: _Relay.contract, event: "TransactionRelayed", logs: logs, sub: sub}, nil
}

// WatchTransactionRelayed is a free log subscription operation binding the contract event 0x2a8e164a2bd95ffaf59651ebd8e75e4e879d97cab53c11dde9dfdbcceb1e6780.
//
// Solidity: event TransactionRelayed(address indexed relay, address indexed from, address indexed to, bytes4 selector, uint256 charge)
func (_Relay *RelayFilterer) WatchTransactionRelayed(opts *bind.WatchOpts, sink chan<- *RelayTransactionRelayed, relay []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var relayRule []interface{}
	for _, relayItem := range relay {
		relayRule = append(relayRule, relayItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Relay.contract.WatchLogs(opts, "TransactionRelayed", relayRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayTransactionRelayed)
				if err := _Relay.contract.UnpackLog(event, "TransactionRelayed", log); err != nil {
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

// ParseTransactionRelayed is a log parse operation binding the contract event 0x2a8e164a2bd95ffaf59651ebd8e75e4e879d97cab53c11dde9dfdbcceb1e6780.
//
// Solidity: event TransactionRelayed(address indexed relay, address indexed from, address indexed to, bytes4 selector, uint256 charge)
func (_Relay *RelayFilterer) ParseTransactionRelayed(log types.Log) (*RelayTransactionRelayed, error) {
	event := new(RelayTransactionRelayed)
	if err := _Relay.contract.UnpackLog(event, "TransactionRelayed", log); err != nil {
		return nil, err
	}
	return event, nil
}
