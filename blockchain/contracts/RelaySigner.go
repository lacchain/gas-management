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
const RelayABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_blocksFrequency\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractDeployed\",\"type\":\"address\"}],\"name\":\"ContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsedLastBlocks\",\"type\":\"uint256\"}],\"name\":\"GasUsedByTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Relayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"charge\",\"type\":\"uint256\"}],\"name\":\"TransactionRelayed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNode\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"}],\"name\":\"deleteNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_byteCode\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"deployMetaTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"deployedAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasUsedLastBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedFunction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"relayMetaTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_blocksFrequency\",\"type\":\"uint8\"}],\"name\":\"setBlocksFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGasUsed\",\"type\":\"uint256\"}],\"name\":\"setGasUsedLastBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// GetGasUsedLastBlocks is a free data retrieval call binding the contract method 0xd03ce2db.
//
// Solidity: function getGasUsedLastBlocks() view returns(uint256)
func (_Relay *RelayCaller) GetGasUsedLastBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getGasUsedLastBlocks")
	return *ret0, err
}

// GetGasUsedLastBlocks is a free data retrieval call binding the contract method 0xd03ce2db.
//
// Solidity: function getGasUsedLastBlocks() view returns(uint256)
func (_Relay *RelaySession) GetGasUsedLastBlocks() (*big.Int, error) {
	return _Relay.Contract.GetGasUsedLastBlocks(&_Relay.CallOpts)
}

// GetGasUsedLastBlocks is a free data retrieval call binding the contract method 0xd03ce2db.
//
// Solidity: function getGasUsedLastBlocks() view returns(uint256)
func (_Relay *RelayCallerSession) GetGasUsedLastBlocks() (*big.Int, error) {
	return _Relay.Contract.GetGasUsedLastBlocks(&_Relay.CallOpts)
}

// GetNodes is a free data retrieval call binding the contract method 0xe29581aa.
//
// Solidity: function getNodes() view returns(uint256)
func (_Relay *RelayCaller) GetNodes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getNodes")
	return *ret0, err
}

// GetNodes is a free data retrieval call binding the contract method 0xe29581aa.
//
// Solidity: function getNodes() view returns(uint256)
func (_Relay *RelaySession) GetNodes() (*big.Int, error) {
	return _Relay.Contract.GetNodes(&_Relay.CallOpts)
}

// GetNodes is a free data retrieval call binding the contract method 0xe29581aa.
//
// Solidity: function getNodes() view returns(uint256)
func (_Relay *RelayCallerSession) GetNodes() (*big.Int, error) {
	return _Relay.Contract.GetNodes(&_Relay.CallOpts)
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

// AddNode is a paid mutator transaction binding the contract method 0x9d95f1cc.
//
// Solidity: function addNode(address newNode) returns()
func (_Relay *RelayTransactor) AddNode(opts *bind.TransactOpts, newNode common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "addNode", newNode)
}

// AddNode is a paid mutator transaction binding the contract method 0x9d95f1cc.
//
// Solidity: function addNode(address newNode) returns()
func (_Relay *RelaySession) AddNode(newNode common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddNode(&_Relay.TransactOpts, newNode)
}

// AddNode is a paid mutator transaction binding the contract method 0x9d95f1cc.
//
// Solidity: function addNode(address newNode) returns()
func (_Relay *RelayTransactorSession) AddNode(newNode common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddNode(&_Relay.TransactOpts, newNode)
}

// DeleteNode is a paid mutator transaction binding the contract method 0x61bacc6d.
//
// Solidity: function deleteNode(uint16 index) returns()
func (_Relay *RelayTransactor) DeleteNode(opts *bind.TransactOpts, index uint16) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "deleteNode", index)
}

// DeleteNode is a paid mutator transaction binding the contract method 0x61bacc6d.
//
// Solidity: function deleteNode(uint16 index) returns()
func (_Relay *RelaySession) DeleteNode(index uint16) (*types.Transaction, error) {
	return _Relay.Contract.DeleteNode(&_Relay.TransactOpts, index)
}

// DeleteNode is a paid mutator transaction binding the contract method 0x61bacc6d.
//
// Solidity: function deleteNode(uint16 index) returns()
func (_Relay *RelayTransactorSession) DeleteNode(index uint16) (*types.Transaction, error) {
	return _Relay.Contract.DeleteNode(&_Relay.TransactOpts, index)
}

// DeployMetaTx is a paid mutator transaction binding the contract method 0x8fad625a.
//
// Solidity: function deployMetaTx(address from, bytes _byteCode, uint256 gasLimit, uint256 nonce, bytes signature) returns(bool success, address deployedAddress)
func (_Relay *RelayTransactor) DeployMetaTx(opts *bind.TransactOpts, from common.Address, _byteCode []byte, gasLimit *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "deployMetaTx", from, _byteCode, gasLimit, nonce, signature)
}

// DeployMetaTx is a paid mutator transaction binding the contract method 0x8fad625a.
//
// Solidity: function deployMetaTx(address from, bytes _byteCode, uint256 gasLimit, uint256 nonce, bytes signature) returns(bool success, address deployedAddress)
func (_Relay *RelaySession) DeployMetaTx(from common.Address, _byteCode []byte, gasLimit *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Relay.Contract.DeployMetaTx(&_Relay.TransactOpts, from, _byteCode, gasLimit, nonce, signature)
}

// DeployMetaTx is a paid mutator transaction binding the contract method 0x8fad625a.
//
// Solidity: function deployMetaTx(address from, bytes _byteCode, uint256 gasLimit, uint256 nonce, bytes signature) returns(bool success, address deployedAddress)
func (_Relay *RelayTransactorSession) DeployMetaTx(from common.Address, _byteCode []byte, gasLimit *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Relay.Contract.DeployMetaTx(&_Relay.TransactOpts, from, _byteCode, gasLimit, nonce, signature)
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

// SetBlocksFrequency is a paid mutator transaction binding the contract method 0xc98f8d98.
//
// Solidity: function setBlocksFrequency(uint8 _blocksFrequency) returns()
func (_Relay *RelayTransactor) SetBlocksFrequency(opts *bind.TransactOpts, _blocksFrequency uint8) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "setBlocksFrequency", _blocksFrequency)
}

// SetBlocksFrequency is a paid mutator transaction binding the contract method 0xc98f8d98.
//
// Solidity: function setBlocksFrequency(uint8 _blocksFrequency) returns()
func (_Relay *RelaySession) SetBlocksFrequency(_blocksFrequency uint8) (*types.Transaction, error) {
	return _Relay.Contract.SetBlocksFrequency(&_Relay.TransactOpts, _blocksFrequency)
}

// SetBlocksFrequency is a paid mutator transaction binding the contract method 0xc98f8d98.
//
// Solidity: function setBlocksFrequency(uint8 _blocksFrequency) returns()
func (_Relay *RelayTransactorSession) SetBlocksFrequency(_blocksFrequency uint8) (*types.Transaction, error) {
	return _Relay.Contract.SetBlocksFrequency(&_Relay.TransactOpts, _blocksFrequency)
}

// SetGasUsedLastBlocks is a paid mutator transaction binding the contract method 0x1aa4de53.
//
// Solidity: function setGasUsedLastBlocks(uint256 newGasUsed) returns()
func (_Relay *RelayTransactor) SetGasUsedLastBlocks(opts *bind.TransactOpts, newGasUsed *big.Int) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "setGasUsedLastBlocks", newGasUsed)
}

// SetGasUsedLastBlocks is a paid mutator transaction binding the contract method 0x1aa4de53.
//
// Solidity: function setGasUsedLastBlocks(uint256 newGasUsed) returns()
func (_Relay *RelaySession) SetGasUsedLastBlocks(newGasUsed *big.Int) (*types.Transaction, error) {
	return _Relay.Contract.SetGasUsedLastBlocks(&_Relay.TransactOpts, newGasUsed)
}

// SetGasUsedLastBlocks is a paid mutator transaction binding the contract method 0x1aa4de53.
//
// Solidity: function setGasUsedLastBlocks(uint256 newGasUsed) returns()
func (_Relay *RelayTransactorSession) SetGasUsedLastBlocks(newGasUsed *big.Int) (*types.Transaction, error) {
	return _Relay.Contract.SetGasUsedLastBlocks(&_Relay.TransactOpts, newGasUsed)
}

// RelayContractDeployedIterator is returned from FilterContractDeployed and is used to iterate over the raw logs and unpacked data for ContractDeployed events raised by the Relay contract.
type RelayContractDeployedIterator struct {
	Event *RelayContractDeployed // Event containing the contract specifics and raw log

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
func (it *RelayContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayContractDeployed)
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
		it.Event = new(RelayContractDeployed)
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
func (it *RelayContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayContractDeployed represents a ContractDeployed event raised by the Relay contract.
type RelayContractDeployed struct {
	ContractDeployed common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterContractDeployed is a free log retrieval operation binding the contract event 0x8ffcdc15a283d706d38281f500270d8b5a656918f555de0913d7455e3e6bc1bf.
//
// Solidity: event ContractDeployed(address contractDeployed)
func (_Relay *RelayFilterer) FilterContractDeployed(opts *bind.FilterOpts) (*RelayContractDeployedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "ContractDeployed")
	if err != nil {
		return nil, err
	}
	return &RelayContractDeployedIterator{contract: _Relay.contract, event: "ContractDeployed", logs: logs, sub: sub}, nil
}

// WatchContractDeployed is a free log subscription operation binding the contract event 0x8ffcdc15a283d706d38281f500270d8b5a656918f555de0913d7455e3e6bc1bf.
//
// Solidity: event ContractDeployed(address contractDeployed)
func (_Relay *RelayFilterer) WatchContractDeployed(opts *bind.WatchOpts, sink chan<- *RelayContractDeployed) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "ContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayContractDeployed)
				if err := _Relay.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
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

// ParseContractDeployed is a log parse operation binding the contract event 0x8ffcdc15a283d706d38281f500270d8b5a656918f555de0913d7455e3e6bc1bf.
//
// Solidity: event ContractDeployed(address contractDeployed)
func (_Relay *RelayFilterer) ParseContractDeployed(log types.Log) (*RelayContractDeployed, error) {
	event := new(RelayContractDeployed)
	if err := _Relay.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayGasUsedByTransactionIterator is returned from FilterGasUsedByTransaction and is used to iterate over the raw logs and unpacked data for GasUsedByTransaction events raised by the Relay contract.
type RelayGasUsedByTransactionIterator struct {
	Event *RelayGasUsedByTransaction // Event containing the contract specifics and raw log

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
func (it *RelayGasUsedByTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayGasUsedByTransaction)
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
		it.Event = new(RelayGasUsedByTransaction)
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
func (it *RelayGasUsedByTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayGasUsedByTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayGasUsedByTransaction represents a GasUsedByTransaction event raised by the Relay contract.
type RelayGasUsedByTransaction struct {
	BlockNumber       *big.Int
	GasUsed           *big.Int
	GasLimit          *big.Int
	GasUsedLastBlocks *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGasUsedByTransaction is a free log retrieval operation binding the contract event 0x0a605bd90e43d7f86ea737757e9c53b9c9169430e9115de003b4631de735eb07.
//
// Solidity: event GasUsedByTransaction(uint256 blockNumber, uint256 gasUsed, uint256 gasLimit, uint256 gasUsedLastBlocks)
func (_Relay *RelayFilterer) FilterGasUsedByTransaction(opts *bind.FilterOpts) (*RelayGasUsedByTransactionIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "GasUsedByTransaction")
	if err != nil {
		return nil, err
	}
	return &RelayGasUsedByTransactionIterator{contract: _Relay.contract, event: "GasUsedByTransaction", logs: logs, sub: sub}, nil
}

// WatchGasUsedByTransaction is a free log subscription operation binding the contract event 0x0a605bd90e43d7f86ea737757e9c53b9c9169430e9115de003b4631de735eb07.
//
// Solidity: event GasUsedByTransaction(uint256 blockNumber, uint256 gasUsed, uint256 gasLimit, uint256 gasUsedLastBlocks)
func (_Relay *RelayFilterer) WatchGasUsedByTransaction(opts *bind.WatchOpts, sink chan<- *RelayGasUsedByTransaction) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "GasUsedByTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayGasUsedByTransaction)
				if err := _Relay.contract.UnpackLog(event, "GasUsedByTransaction", log); err != nil {
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

// ParseGasUsedByTransaction is a log parse operation binding the contract event 0x0a605bd90e43d7f86ea737757e9c53b9c9169430e9115de003b4631de735eb07.
//
// Solidity: event GasUsedByTransaction(uint256 blockNumber, uint256 gasUsed, uint256 gasLimit, uint256 gasUsedLastBlocks)
func (_Relay *RelayFilterer) ParseGasUsedByTransaction(log types.Log) (*RelayGasUsedByTransaction, error) {
	event := new(RelayGasUsedByTransaction)
	if err := _Relay.contract.UnpackLog(event, "GasUsedByTransaction", log); err != nil {
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
