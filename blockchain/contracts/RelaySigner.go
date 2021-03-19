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
const RelayABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_blocksFrequency\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_accountIngress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"AccountIngressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIRelayHub.ErrorCode\",\"name\":\"errorCode\",\"type\":\"uint8\"}],\"name\":\"BadTransactionSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"blocksFrequency\",\"type\":\"uint8\"}],\"name\":\"BlockFrequencyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractDeployed\",\"type\":\"address\"}],\"name\":\"ContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"countExceeded\",\"type\":\"uint8\"}],\"name\":\"GasLimitExceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsedLastBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"averageLastBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGasLimit\",\"type\":\"uint256\"}],\"name\":\"GasLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsedLastBlocks\",\"type\":\"uint256\"}],\"name\":\"GasUsedByTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsedRelayHub\",\"type\":\"uint256\"}],\"name\":\"GasUsedRelayHubChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxGasBlockLimit\",\"type\":\"uint256\"}],\"name\":\"MaxGasBlockLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newNode\",\"type\":\"address\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"NodeBlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldNode\",\"type\":\"address\"}],\"name\":\"NodeDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"decodedFunction\",\"type\":\"bytes\"}],\"name\":\"Parameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"Recalculated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Relayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"name\":\"TransactionRelayed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNode\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"deleteNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasUsedLastBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accountIngress\",\"type\":\"address\"}],\"name\":\"setAccounIngress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_blocksFrequency\",\"type\":\"uint8\"}],\"name\":\"setBlocksFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGasUsed\",\"type\":\"uint256\"}],\"name\":\"setGasUsedLastBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasUsedRelayHub\",\"type\":\"uint256\"}],\"name\":\"setGasUsedRelayHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxGasBlockLimit\",\"type\":\"uint256\"}],\"name\":\"setMaxGasBlockLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signingData\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"relayMetaTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signingData\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"deployMetaTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"deployedAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMsgSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"name\":\"increaseGasUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Relay *RelayCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "DEFAULT_ADMIN_ROLE")
	return *ret0, err
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Relay *RelaySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Relay.Contract.DEFAULTADMINROLE(&_Relay.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Relay *RelayCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Relay.Contract.DEFAULTADMINROLE(&_Relay.CallOpts)
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

// GetMsgSender is a free data retrieval call binding the contract method 0x7a6ce2e1.
//
// Solidity: function getMsgSender() view returns(address)
func (_Relay *RelayCaller) GetMsgSender(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getMsgSender")
	return *ret0, err
}

// GetMsgSender is a free data retrieval call binding the contract method 0x7a6ce2e1.
//
// Solidity: function getMsgSender() view returns(address)
func (_Relay *RelaySession) GetMsgSender() (common.Address, error) {
	return _Relay.Contract.GetMsgSender(&_Relay.CallOpts)
}

// GetMsgSender is a free data retrieval call binding the contract method 0x7a6ce2e1.
//
// Solidity: function getMsgSender() view returns(address)
func (_Relay *RelayCallerSession) GetMsgSender() (common.Address, error) {
	return _Relay.Contract.GetMsgSender(&_Relay.CallOpts)
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

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Relay *RelayCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getRoleAdmin", role)
	return *ret0, err
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Relay *RelaySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Relay.Contract.GetRoleAdmin(&_Relay.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Relay *RelayCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Relay.Contract.GetRoleAdmin(&_Relay.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Relay *RelayCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getRoleMember", role, index)
	return *ret0, err
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Relay *RelaySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Relay.Contract.GetRoleMember(&_Relay.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Relay *RelayCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Relay.Contract.GetRoleMember(&_Relay.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Relay *RelayCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "getRoleMemberCount", role)
	return *ret0, err
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Relay *RelaySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Relay.Contract.GetRoleMemberCount(&_Relay.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Relay *RelayCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Relay.Contract.GetRoleMemberCount(&_Relay.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Relay *RelayCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "hasRole", role, account)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Relay *RelaySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Relay.Contract.HasRole(&_Relay.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Relay *RelayCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Relay.Contract.HasRole(&_Relay.CallOpts, role, account)
}

// AddNode is a paid mutator transaction binding the contract method 0x9d95f1cc.
//
// Solidity: function addNode(address newNode) returns(bool)
func (_Relay *RelayTransactor) AddNode(opts *bind.TransactOpts, newNode common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "addNode", newNode)
}

// AddNode is a paid mutator transaction binding the contract method 0x9d95f1cc.
//
// Solidity: function addNode(address newNode) returns(bool)
func (_Relay *RelaySession) AddNode(newNode common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddNode(&_Relay.TransactOpts, newNode)
}

// AddNode is a paid mutator transaction binding the contract method 0x9d95f1cc.
//
// Solidity: function addNode(address newNode) returns(bool)
func (_Relay *RelayTransactorSession) AddNode(newNode common.Address) (*types.Transaction, error) {
	return _Relay.Contract.AddNode(&_Relay.TransactOpts, newNode)
}

// DeleteNode is a paid mutator transaction binding the contract method 0x2d4ede93.
//
// Solidity: function deleteNode(address node) returns()
func (_Relay *RelayTransactor) DeleteNode(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "deleteNode", node)
}

// DeleteNode is a paid mutator transaction binding the contract method 0x2d4ede93.
//
// Solidity: function deleteNode(address node) returns()
func (_Relay *RelaySession) DeleteNode(node common.Address) (*types.Transaction, error) {
	return _Relay.Contract.DeleteNode(&_Relay.TransactOpts, node)
}

// DeleteNode is a paid mutator transaction binding the contract method 0x2d4ede93.
//
// Solidity: function deleteNode(address node) returns()
func (_Relay *RelayTransactorSession) DeleteNode(node common.Address) (*types.Transaction, error) {
	return _Relay.Contract.DeleteNode(&_Relay.TransactOpts, node)
}

// DeployMetaTx is a paid mutator transaction binding the contract method 0x4b802a36.
//
// Solidity: function deployMetaTx(bytes signingData, uint8 v, bytes32 r, bytes32 s) returns(bool success, address deployedAddress)
func (_Relay *RelayTransactor) DeployMetaTx(opts *bind.TransactOpts, signingData []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "deployMetaTx", signingData, v, r, s)
}

// DeployMetaTx is a paid mutator transaction binding the contract method 0x4b802a36.
//
// Solidity: function deployMetaTx(bytes signingData, uint8 v, bytes32 r, bytes32 s) returns(bool success, address deployedAddress)
func (_Relay *RelaySession) DeployMetaTx(signingData []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Relay.Contract.DeployMetaTx(&_Relay.TransactOpts, signingData, v, r, s)
}

// DeployMetaTx is a paid mutator transaction binding the contract method 0x4b802a36.
//
// Solidity: function deployMetaTx(bytes signingData, uint8 v, bytes32 r, bytes32 s) returns(bool success, address deployedAddress)
func (_Relay *RelayTransactorSession) DeployMetaTx(signingData []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Relay.Contract.DeployMetaTx(&_Relay.TransactOpts, signingData, v, r, s)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Relay *RelayTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Relay *RelaySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.Contract.GrantRole(&_Relay.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Relay *RelayTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.Contract.GrantRole(&_Relay.TransactOpts, role, account)
}

// IncreaseGasUsed is a paid mutator transaction binding the contract method 0x2e74335a.
//
// Solidity: function increaseGasUsed(uint256 gasUsed) returns(bool)
func (_Relay *RelayTransactor) IncreaseGasUsed(opts *bind.TransactOpts, gasUsed *big.Int) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "increaseGasUsed", gasUsed)
}

// IncreaseGasUsed is a paid mutator transaction binding the contract method 0x2e74335a.
//
// Solidity: function increaseGasUsed(uint256 gasUsed) returns(bool)
func (_Relay *RelaySession) IncreaseGasUsed(gasUsed *big.Int) (*types.Transaction, error) {
	return _Relay.Contract.IncreaseGasUsed(&_Relay.TransactOpts, gasUsed)
}

// IncreaseGasUsed is a paid mutator transaction binding the contract method 0x2e74335a.
//
// Solidity: function increaseGasUsed(uint256 gasUsed) returns(bool)
func (_Relay *RelayTransactorSession) IncreaseGasUsed(gasUsed *big.Int) (*types.Transaction, error) {
	return _Relay.Contract.IncreaseGasUsed(&_Relay.TransactOpts, gasUsed)
}

// RelayMetaTx is a paid mutator transaction binding the contract method 0xa04fb2ad.
//
// Solidity: function relayMetaTx(bytes signingData, uint8 v, bytes32 r, bytes32 s) returns(bool success)
func (_Relay *RelayTransactor) RelayMetaTx(opts *bind.TransactOpts, signingData []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "relayMetaTx", signingData, v, r, s)
}

// RelayMetaTx is a paid mutator transaction binding the contract method 0xa04fb2ad.
//
// Solidity: function relayMetaTx(bytes signingData, uint8 v, bytes32 r, bytes32 s) returns(bool success)
func (_Relay *RelaySession) RelayMetaTx(signingData []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Relay.Contract.RelayMetaTx(&_Relay.TransactOpts, signingData, v, r, s)
}

// RelayMetaTx is a paid mutator transaction binding the contract method 0xa04fb2ad.
//
// Solidity: function relayMetaTx(bytes signingData, uint8 v, bytes32 r, bytes32 s) returns(bool success)
func (_Relay *RelayTransactorSession) RelayMetaTx(signingData []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Relay.Contract.RelayMetaTx(&_Relay.TransactOpts, signingData, v, r, s)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Relay *RelayTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Relay *RelaySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RenounceRole(&_Relay.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Relay *RelayTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RenounceRole(&_Relay.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Relay *RelayTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Relay *RelaySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RevokeRole(&_Relay.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Relay *RelayTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Relay.Contract.RevokeRole(&_Relay.TransactOpts, role, account)
}

// SetAccounIngress is a paid mutator transaction binding the contract method 0x4473d59d.
//
// Solidity: function setAccounIngress(address _accountIngress) returns()
func (_Relay *RelayTransactor) SetAccounIngress(opts *bind.TransactOpts, _accountIngress common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "setAccounIngress", _accountIngress)
}

// SetAccounIngress is a paid mutator transaction binding the contract method 0x4473d59d.
//
// Solidity: function setAccounIngress(address _accountIngress) returns()
func (_Relay *RelaySession) SetAccounIngress(_accountIngress common.Address) (*types.Transaction, error) {
	return _Relay.Contract.SetAccounIngress(&_Relay.TransactOpts, _accountIngress)
}

// SetAccounIngress is a paid mutator transaction binding the contract method 0x4473d59d.
//
// Solidity: function setAccounIngress(address _accountIngress) returns()
func (_Relay *RelayTransactorSession) SetAccounIngress(_accountIngress common.Address) (*types.Transaction, error) {
	return _Relay.Contract.SetAccounIngress(&_Relay.TransactOpts, _accountIngress)
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

// SetGasUsedRelayHub is a paid mutator transaction binding the contract method 0x78beb3e7.
//
// Solidity: function setGasUsedRelayHub(uint256 _gasUsedRelayHub) returns()
func (_Relay *RelayTransactor) SetGasUsedRelayHub(opts *bind.TransactOpts, _gasUsedRelayHub *big.Int) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "setGasUsedRelayHub", _gasUsedRelayHub)
}

// SetGasUsedRelayHub is a paid mutator transaction binding the contract method 0x78beb3e7.
//
// Solidity: function setGasUsedRelayHub(uint256 _gasUsedRelayHub) returns()
func (_Relay *RelaySession) SetGasUsedRelayHub(_gasUsedRelayHub *big.Int) (*types.Transaction, error) {
	return _Relay.Contract.SetGasUsedRelayHub(&_Relay.TransactOpts, _gasUsedRelayHub)
}

// SetGasUsedRelayHub is a paid mutator transaction binding the contract method 0x78beb3e7.
//
// Solidity: function setGasUsedRelayHub(uint256 _gasUsedRelayHub) returns()
func (_Relay *RelayTransactorSession) SetGasUsedRelayHub(_gasUsedRelayHub *big.Int) (*types.Transaction, error) {
	return _Relay.Contract.SetGasUsedRelayHub(&_Relay.TransactOpts, _gasUsedRelayHub)
}

// SetMaxGasBlockLimit is a paid mutator transaction binding the contract method 0x2a45d599.
//
// Solidity: function setMaxGasBlockLimit(uint256 _maxGasBlockLimit) returns()
func (_Relay *RelayTransactor) SetMaxGasBlockLimit(opts *bind.TransactOpts, _maxGasBlockLimit *big.Int) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "setMaxGasBlockLimit", _maxGasBlockLimit)
}

// SetMaxGasBlockLimit is a paid mutator transaction binding the contract method 0x2a45d599.
//
// Solidity: function setMaxGasBlockLimit(uint256 _maxGasBlockLimit) returns()
func (_Relay *RelaySession) SetMaxGasBlockLimit(_maxGasBlockLimit *big.Int) (*types.Transaction, error) {
	return _Relay.Contract.SetMaxGasBlockLimit(&_Relay.TransactOpts, _maxGasBlockLimit)
}

// SetMaxGasBlockLimit is a paid mutator transaction binding the contract method 0x2a45d599.
//
// Solidity: function setMaxGasBlockLimit(uint256 _maxGasBlockLimit) returns()
func (_Relay *RelayTransactorSession) SetMaxGasBlockLimit(_maxGasBlockLimit *big.Int) (*types.Transaction, error) {
	return _Relay.Contract.SetMaxGasBlockLimit(&_Relay.TransactOpts, _maxGasBlockLimit)
}

// RelayAccountIngressChangedIterator is returned from FilterAccountIngressChanged and is used to iterate over the raw logs and unpacked data for AccountIngressChanged events raised by the Relay contract.
type RelayAccountIngressChangedIterator struct {
	Event *RelayAccountIngressChanged // Event containing the contract specifics and raw log

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
func (it *RelayAccountIngressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayAccountIngressChanged)
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
		it.Event = new(RelayAccountIngressChanged)
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
func (it *RelayAccountIngressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayAccountIngressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayAccountIngressChanged represents a AccountIngressChanged event raised by the Relay contract.
type RelayAccountIngressChanged struct {
	Admin      common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountIngressChanged is a free log retrieval operation binding the contract event 0xf552f6d1d0f097137db64c11c170afb61be6d9a123c50c5fc38c5b1f56a205f3.
//
// Solidity: event AccountIngressChanged(address admin, address newAddress)
func (_Relay *RelayFilterer) FilterAccountIngressChanged(opts *bind.FilterOpts) (*RelayAccountIngressChangedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "AccountIngressChanged")
	if err != nil {
		return nil, err
	}
	return &RelayAccountIngressChangedIterator{contract: _Relay.contract, event: "AccountIngressChanged", logs: logs, sub: sub}, nil
}

// WatchAccountIngressChanged is a free log subscription operation binding the contract event 0xf552f6d1d0f097137db64c11c170afb61be6d9a123c50c5fc38c5b1f56a205f3.
//
// Solidity: event AccountIngressChanged(address admin, address newAddress)
func (_Relay *RelayFilterer) WatchAccountIngressChanged(opts *bind.WatchOpts, sink chan<- *RelayAccountIngressChanged) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "AccountIngressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayAccountIngressChanged)
				if err := _Relay.contract.UnpackLog(event, "AccountIngressChanged", log); err != nil {
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

// ParseAccountIngressChanged is a log parse operation binding the contract event 0xf552f6d1d0f097137db64c11c170afb61be6d9a123c50c5fc38c5b1f56a205f3.
//
// Solidity: event AccountIngressChanged(address admin, address newAddress)
func (_Relay *RelayFilterer) ParseAccountIngressChanged(log types.Log) (*RelayAccountIngressChanged, error) {
	event := new(RelayAccountIngressChanged)
	if err := _Relay.contract.UnpackLog(event, "AccountIngressChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayBadTransactionSentIterator is returned from FilterBadTransactionSent and is used to iterate over the raw logs and unpacked data for BadTransactionSent events raised by the Relay contract.
type RelayBadTransactionSentIterator struct {
	Event *RelayBadTransactionSent // Event containing the contract specifics and raw log

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
func (it *RelayBadTransactionSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayBadTransactionSent)
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
		it.Event = new(RelayBadTransactionSent)
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
func (it *RelayBadTransactionSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayBadTransactionSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayBadTransactionSent represents a BadTransactionSent event raised by the Relay contract.
type RelayBadTransactionSent struct {
	Node           common.Address
	OriginalSender common.Address
	ErrorCode      uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBadTransactionSent is a free log retrieval operation binding the contract event 0xc62bb53370aadcfe652881fc57ef9ca04a7c473e83b963413f2cf2b5d66c3ef3.
//
// Solidity: event BadTransactionSent(address node, address originalSender, uint8 errorCode)
func (_Relay *RelayFilterer) FilterBadTransactionSent(opts *bind.FilterOpts) (*RelayBadTransactionSentIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "BadTransactionSent")
	if err != nil {
		return nil, err
	}
	return &RelayBadTransactionSentIterator{contract: _Relay.contract, event: "BadTransactionSent", logs: logs, sub: sub}, nil
}

// WatchBadTransactionSent is a free log subscription operation binding the contract event 0xc62bb53370aadcfe652881fc57ef9ca04a7c473e83b963413f2cf2b5d66c3ef3.
//
// Solidity: event BadTransactionSent(address node, address originalSender, uint8 errorCode)
func (_Relay *RelayFilterer) WatchBadTransactionSent(opts *bind.WatchOpts, sink chan<- *RelayBadTransactionSent) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "BadTransactionSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayBadTransactionSent)
				if err := _Relay.contract.UnpackLog(event, "BadTransactionSent", log); err != nil {
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

// ParseBadTransactionSent is a log parse operation binding the contract event 0xc62bb53370aadcfe652881fc57ef9ca04a7c473e83b963413f2cf2b5d66c3ef3.
//
// Solidity: event BadTransactionSent(address node, address originalSender, uint8 errorCode)
func (_Relay *RelayFilterer) ParseBadTransactionSent(log types.Log) (*RelayBadTransactionSent, error) {
	event := new(RelayBadTransactionSent)
	if err := _Relay.contract.UnpackLog(event, "BadTransactionSent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayBlockFrequencyChangedIterator is returned from FilterBlockFrequencyChanged and is used to iterate over the raw logs and unpacked data for BlockFrequencyChanged events raised by the Relay contract.
type RelayBlockFrequencyChangedIterator struct {
	Event *RelayBlockFrequencyChanged // Event containing the contract specifics and raw log

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
func (it *RelayBlockFrequencyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayBlockFrequencyChanged)
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
		it.Event = new(RelayBlockFrequencyChanged)
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
func (it *RelayBlockFrequencyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayBlockFrequencyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayBlockFrequencyChanged represents a BlockFrequencyChanged event raised by the Relay contract.
type RelayBlockFrequencyChanged struct {
	Admin           common.Address
	BlocksFrequency uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlockFrequencyChanged is a free log retrieval operation binding the contract event 0x761dd0dd5bb1bfaf8267b9fdad2c2e273a0e661252207ecafc0f97a374c07c21.
//
// Solidity: event BlockFrequencyChanged(address admin, uint8 blocksFrequency)
func (_Relay *RelayFilterer) FilterBlockFrequencyChanged(opts *bind.FilterOpts) (*RelayBlockFrequencyChangedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "BlockFrequencyChanged")
	if err != nil {
		return nil, err
	}
	return &RelayBlockFrequencyChangedIterator{contract: _Relay.contract, event: "BlockFrequencyChanged", logs: logs, sub: sub}, nil
}

// WatchBlockFrequencyChanged is a free log subscription operation binding the contract event 0x761dd0dd5bb1bfaf8267b9fdad2c2e273a0e661252207ecafc0f97a374c07c21.
//
// Solidity: event BlockFrequencyChanged(address admin, uint8 blocksFrequency)
func (_Relay *RelayFilterer) WatchBlockFrequencyChanged(opts *bind.WatchOpts, sink chan<- *RelayBlockFrequencyChanged) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "BlockFrequencyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayBlockFrequencyChanged)
				if err := _Relay.contract.UnpackLog(event, "BlockFrequencyChanged", log); err != nil {
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

// ParseBlockFrequencyChanged is a log parse operation binding the contract event 0x761dd0dd5bb1bfaf8267b9fdad2c2e273a0e661252207ecafc0f97a374c07c21.
//
// Solidity: event BlockFrequencyChanged(address admin, uint8 blocksFrequency)
func (_Relay *RelayFilterer) ParseBlockFrequencyChanged(log types.Log) (*RelayBlockFrequencyChanged, error) {
	event := new(RelayBlockFrequencyChanged)
	if err := _Relay.contract.UnpackLog(event, "BlockFrequencyChanged", log); err != nil {
		return nil, err
	}
	return event, nil
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
	Relay            common.Address
	From             common.Address
	ContractDeployed common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterContractDeployed is a free log retrieval operation binding the contract event 0x8a14d1d7200360982eafa429b53edf408f7f589e6da6558f3c116c7f708327b3.
//
// Solidity: event ContractDeployed(address indexed relay, address indexed from, address contractDeployed)
func (_Relay *RelayFilterer) FilterContractDeployed(opts *bind.FilterOpts, relay []common.Address, from []common.Address) (*RelayContractDeployedIterator, error) {

	var relayRule []interface{}
	for _, relayItem := range relay {
		relayRule = append(relayRule, relayItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Relay.contract.FilterLogs(opts, "ContractDeployed", relayRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &RelayContractDeployedIterator{contract: _Relay.contract, event: "ContractDeployed", logs: logs, sub: sub}, nil
}

// WatchContractDeployed is a free log subscription operation binding the contract event 0x8a14d1d7200360982eafa429b53edf408f7f589e6da6558f3c116c7f708327b3.
//
// Solidity: event ContractDeployed(address indexed relay, address indexed from, address contractDeployed)
func (_Relay *RelayFilterer) WatchContractDeployed(opts *bind.WatchOpts, sink chan<- *RelayContractDeployed, relay []common.Address, from []common.Address) (event.Subscription, error) {

	var relayRule []interface{}
	for _, relayItem := range relay {
		relayRule = append(relayRule, relayItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Relay.contract.WatchLogs(opts, "ContractDeployed", relayRule, fromRule)
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

// ParseContractDeployed is a log parse operation binding the contract event 0x8a14d1d7200360982eafa429b53edf408f7f589e6da6558f3c116c7f708327b3.
//
// Solidity: event ContractDeployed(address indexed relay, address indexed from, address contractDeployed)
func (_Relay *RelayFilterer) ParseContractDeployed(log types.Log) (*RelayContractDeployed, error) {
	event := new(RelayContractDeployed)
	if err := _Relay.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayGasLimitExceededIterator is returned from FilterGasLimitExceeded and is used to iterate over the raw logs and unpacked data for GasLimitExceeded events raised by the Relay contract.
type RelayGasLimitExceededIterator struct {
	Event *RelayGasLimitExceeded // Event containing the contract specifics and raw log

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
func (it *RelayGasLimitExceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayGasLimitExceeded)
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
		it.Event = new(RelayGasLimitExceeded)
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
func (it *RelayGasLimitExceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayGasLimitExceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayGasLimitExceeded represents a GasLimitExceeded event raised by the Relay contract.
type RelayGasLimitExceeded struct {
	Node          common.Address
	BlockNumber   *big.Int
	CountExceeded uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGasLimitExceeded is a free log retrieval operation binding the contract event 0xcbd76994be9202c276713d27fdea7e4e64ef81f143f972ae0ab849ff417bf036.
//
// Solidity: event GasLimitExceeded(address node, uint256 blockNumber, uint8 countExceeded)
func (_Relay *RelayFilterer) FilterGasLimitExceeded(opts *bind.FilterOpts) (*RelayGasLimitExceededIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "GasLimitExceeded")
	if err != nil {
		return nil, err
	}
	return &RelayGasLimitExceededIterator{contract: _Relay.contract, event: "GasLimitExceeded", logs: logs, sub: sub}, nil
}

// WatchGasLimitExceeded is a free log subscription operation binding the contract event 0xcbd76994be9202c276713d27fdea7e4e64ef81f143f972ae0ab849ff417bf036.
//
// Solidity: event GasLimitExceeded(address node, uint256 blockNumber, uint8 countExceeded)
func (_Relay *RelayFilterer) WatchGasLimitExceeded(opts *bind.WatchOpts, sink chan<- *RelayGasLimitExceeded) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "GasLimitExceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayGasLimitExceeded)
				if err := _Relay.contract.UnpackLog(event, "GasLimitExceeded", log); err != nil {
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

// ParseGasLimitExceeded is a log parse operation binding the contract event 0xcbd76994be9202c276713d27fdea7e4e64ef81f143f972ae0ab849ff417bf036.
//
// Solidity: event GasLimitExceeded(address node, uint256 blockNumber, uint8 countExceeded)
func (_Relay *RelayFilterer) ParseGasLimitExceeded(log types.Log) (*RelayGasLimitExceeded, error) {
	event := new(RelayGasLimitExceeded)
	if err := _Relay.contract.UnpackLog(event, "GasLimitExceeded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayGasLimitSetIterator is returned from FilterGasLimitSet and is used to iterate over the raw logs and unpacked data for GasLimitSet events raised by the Relay contract.
type RelayGasLimitSetIterator struct {
	Event *RelayGasLimitSet // Event containing the contract specifics and raw log

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
func (it *RelayGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayGasLimitSet)
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
		it.Event = new(RelayGasLimitSet)
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
func (it *RelayGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayGasLimitSet represents a GasLimitSet event raised by the Relay contract.
type RelayGasLimitSet struct {
	BlockNumber       *big.Int
	GasUsedLastBlocks *big.Int
	AverageLastBlocks *big.Int
	NewGasLimit       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGasLimitSet is a free log retrieval operation binding the contract event 0x1ecdaca0ae98a95eed765c0622982b0f7691f9a345988f8fca91c1c016ce5ee7.
//
// Solidity: event GasLimitSet(uint256 blockNumber, uint256 gasUsedLastBlocks, uint256 averageLastBlocks, uint256 newGasLimit)
func (_Relay *RelayFilterer) FilterGasLimitSet(opts *bind.FilterOpts) (*RelayGasLimitSetIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "GasLimitSet")
	if err != nil {
		return nil, err
	}
	return &RelayGasLimitSetIterator{contract: _Relay.contract, event: "GasLimitSet", logs: logs, sub: sub}, nil
}

// WatchGasLimitSet is a free log subscription operation binding the contract event 0x1ecdaca0ae98a95eed765c0622982b0f7691f9a345988f8fca91c1c016ce5ee7.
//
// Solidity: event GasLimitSet(uint256 blockNumber, uint256 gasUsedLastBlocks, uint256 averageLastBlocks, uint256 newGasLimit)
func (_Relay *RelayFilterer) WatchGasLimitSet(opts *bind.WatchOpts, sink chan<- *RelayGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "GasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayGasLimitSet)
				if err := _Relay.contract.UnpackLog(event, "GasLimitSet", log); err != nil {
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

// ParseGasLimitSet is a log parse operation binding the contract event 0x1ecdaca0ae98a95eed765c0622982b0f7691f9a345988f8fca91c1c016ce5ee7.
//
// Solidity: event GasLimitSet(uint256 blockNumber, uint256 gasUsedLastBlocks, uint256 averageLastBlocks, uint256 newGasLimit)
func (_Relay *RelayFilterer) ParseGasLimitSet(log types.Log) (*RelayGasLimitSet, error) {
	event := new(RelayGasLimitSet)
	if err := _Relay.contract.UnpackLog(event, "GasLimitSet", log); err != nil {
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
	Node              common.Address
	BlockNumber       *big.Int
	GasUsed           *big.Int
	GasLimit          *big.Int
	GasUsedLastBlocks *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGasUsedByTransaction is a free log retrieval operation binding the contract event 0x260359eeed8459102359245337088f93b15364b134b4be9092d508e741bbdee1.
//
// Solidity: event GasUsedByTransaction(address node, uint256 blockNumber, uint256 gasUsed, uint256 gasLimit, uint256 gasUsedLastBlocks)
func (_Relay *RelayFilterer) FilterGasUsedByTransaction(opts *bind.FilterOpts) (*RelayGasUsedByTransactionIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "GasUsedByTransaction")
	if err != nil {
		return nil, err
	}
	return &RelayGasUsedByTransactionIterator{contract: _Relay.contract, event: "GasUsedByTransaction", logs: logs, sub: sub}, nil
}

// WatchGasUsedByTransaction is a free log subscription operation binding the contract event 0x260359eeed8459102359245337088f93b15364b134b4be9092d508e741bbdee1.
//
// Solidity: event GasUsedByTransaction(address node, uint256 blockNumber, uint256 gasUsed, uint256 gasLimit, uint256 gasUsedLastBlocks)
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

// ParseGasUsedByTransaction is a log parse operation binding the contract event 0x260359eeed8459102359245337088f93b15364b134b4be9092d508e741bbdee1.
//
// Solidity: event GasUsedByTransaction(address node, uint256 blockNumber, uint256 gasUsed, uint256 gasLimit, uint256 gasUsedLastBlocks)
func (_Relay *RelayFilterer) ParseGasUsedByTransaction(log types.Log) (*RelayGasUsedByTransaction, error) {
	event := new(RelayGasUsedByTransaction)
	if err := _Relay.contract.UnpackLog(event, "GasUsedByTransaction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayGasUsedRelayHubChangedIterator is returned from FilterGasUsedRelayHubChanged and is used to iterate over the raw logs and unpacked data for GasUsedRelayHubChanged events raised by the Relay contract.
type RelayGasUsedRelayHubChangedIterator struct {
	Event *RelayGasUsedRelayHubChanged // Event containing the contract specifics and raw log

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
func (it *RelayGasUsedRelayHubChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayGasUsedRelayHubChanged)
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
		it.Event = new(RelayGasUsedRelayHubChanged)
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
func (it *RelayGasUsedRelayHubChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayGasUsedRelayHubChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayGasUsedRelayHubChanged represents a GasUsedRelayHubChanged event raised by the Relay contract.
type RelayGasUsedRelayHubChanged struct {
	Admin           common.Address
	GasUsedRelayHub *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterGasUsedRelayHubChanged is a free log retrieval operation binding the contract event 0x3813cab05b71ba7f1b896b5c81bc102fb1329cb18c002d93301454621f6e2dd6.
//
// Solidity: event GasUsedRelayHubChanged(address admin, uint256 gasUsedRelayHub)
func (_Relay *RelayFilterer) FilterGasUsedRelayHubChanged(opts *bind.FilterOpts) (*RelayGasUsedRelayHubChangedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "GasUsedRelayHubChanged")
	if err != nil {
		return nil, err
	}
	return &RelayGasUsedRelayHubChangedIterator{contract: _Relay.contract, event: "GasUsedRelayHubChanged", logs: logs, sub: sub}, nil
}

// WatchGasUsedRelayHubChanged is a free log subscription operation binding the contract event 0x3813cab05b71ba7f1b896b5c81bc102fb1329cb18c002d93301454621f6e2dd6.
//
// Solidity: event GasUsedRelayHubChanged(address admin, uint256 gasUsedRelayHub)
func (_Relay *RelayFilterer) WatchGasUsedRelayHubChanged(opts *bind.WatchOpts, sink chan<- *RelayGasUsedRelayHubChanged) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "GasUsedRelayHubChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayGasUsedRelayHubChanged)
				if err := _Relay.contract.UnpackLog(event, "GasUsedRelayHubChanged", log); err != nil {
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

// ParseGasUsedRelayHubChanged is a log parse operation binding the contract event 0x3813cab05b71ba7f1b896b5c81bc102fb1329cb18c002d93301454621f6e2dd6.
//
// Solidity: event GasUsedRelayHubChanged(address admin, uint256 gasUsedRelayHub)
func (_Relay *RelayFilterer) ParseGasUsedRelayHubChanged(log types.Log) (*RelayGasUsedRelayHubChanged, error) {
	event := new(RelayGasUsedRelayHubChanged)
	if err := _Relay.contract.UnpackLog(event, "GasUsedRelayHubChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayMaxGasBlockLimitChangedIterator is returned from FilterMaxGasBlockLimitChanged and is used to iterate over the raw logs and unpacked data for MaxGasBlockLimitChanged events raised by the Relay contract.
type RelayMaxGasBlockLimitChangedIterator struct {
	Event *RelayMaxGasBlockLimitChanged // Event containing the contract specifics and raw log

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
func (it *RelayMaxGasBlockLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayMaxGasBlockLimitChanged)
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
		it.Event = new(RelayMaxGasBlockLimitChanged)
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
func (it *RelayMaxGasBlockLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayMaxGasBlockLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayMaxGasBlockLimitChanged represents a MaxGasBlockLimitChanged event raised by the Relay contract.
type RelayMaxGasBlockLimitChanged struct {
	BlockNumber      *big.Int
	Admin            common.Address
	MaxGasBlockLimit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMaxGasBlockLimitChanged is a free log retrieval operation binding the contract event 0x2eda5665530e0f918783d2a5e33519c436ef2275f0960978c0a3b9258483339b.
//
// Solidity: event MaxGasBlockLimitChanged(uint256 blockNumber, address admin, uint256 maxGasBlockLimit)
func (_Relay *RelayFilterer) FilterMaxGasBlockLimitChanged(opts *bind.FilterOpts) (*RelayMaxGasBlockLimitChangedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "MaxGasBlockLimitChanged")
	if err != nil {
		return nil, err
	}
	return &RelayMaxGasBlockLimitChangedIterator{contract: _Relay.contract, event: "MaxGasBlockLimitChanged", logs: logs, sub: sub}, nil
}

// WatchMaxGasBlockLimitChanged is a free log subscription operation binding the contract event 0x2eda5665530e0f918783d2a5e33519c436ef2275f0960978c0a3b9258483339b.
//
// Solidity: event MaxGasBlockLimitChanged(uint256 blockNumber, address admin, uint256 maxGasBlockLimit)
func (_Relay *RelayFilterer) WatchMaxGasBlockLimitChanged(opts *bind.WatchOpts, sink chan<- *RelayMaxGasBlockLimitChanged) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "MaxGasBlockLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayMaxGasBlockLimitChanged)
				if err := _Relay.contract.UnpackLog(event, "MaxGasBlockLimitChanged", log); err != nil {
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

// ParseMaxGasBlockLimitChanged is a log parse operation binding the contract event 0x2eda5665530e0f918783d2a5e33519c436ef2275f0960978c0a3b9258483339b.
//
// Solidity: event MaxGasBlockLimitChanged(uint256 blockNumber, address admin, uint256 maxGasBlockLimit)
func (_Relay *RelayFilterer) ParseMaxGasBlockLimitChanged(log types.Log) (*RelayMaxGasBlockLimitChanged, error) {
	event := new(RelayMaxGasBlockLimitChanged)
	if err := _Relay.contract.UnpackLog(event, "MaxGasBlockLimitChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayNodeAddedIterator is returned from FilterNodeAdded and is used to iterate over the raw logs and unpacked data for NodeAdded events raised by the Relay contract.
type RelayNodeAddedIterator struct {
	Event *RelayNodeAdded // Event containing the contract specifics and raw log

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
func (it *RelayNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayNodeAdded)
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
		it.Event = new(RelayNodeAdded)
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
func (it *RelayNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayNodeAdded represents a NodeAdded event raised by the Relay contract.
type RelayNodeAdded struct {
	NewNode common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeAdded is a free log retrieval operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address newNode)
func (_Relay *RelayFilterer) FilterNodeAdded(opts *bind.FilterOpts) (*RelayNodeAddedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "NodeAdded")
	if err != nil {
		return nil, err
	}
	return &RelayNodeAddedIterator{contract: _Relay.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

// WatchNodeAdded is a free log subscription operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address newNode)
func (_Relay *RelayFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *RelayNodeAdded) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "NodeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayNodeAdded)
				if err := _Relay.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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

// ParseNodeAdded is a log parse operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address newNode)
func (_Relay *RelayFilterer) ParseNodeAdded(log types.Log) (*RelayNodeAdded, error) {
	event := new(RelayNodeAdded)
	if err := _Relay.contract.UnpackLog(event, "NodeAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayNodeBlockedIterator is returned from FilterNodeBlocked and is used to iterate over the raw logs and unpacked data for NodeBlocked events raised by the Relay contract.
type RelayNodeBlockedIterator struct {
	Event *RelayNodeBlocked // Event containing the contract specifics and raw log

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
func (it *RelayNodeBlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayNodeBlocked)
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
		it.Event = new(RelayNodeBlocked)
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
func (it *RelayNodeBlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayNodeBlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayNodeBlocked represents a NodeBlocked event raised by the Relay contract.
type RelayNodeBlocked struct {
	Node        common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeBlocked is a free log retrieval operation binding the contract event 0x29894930b6f680a84bc3015c2ee88544ea90c73f564a4dba638e3c55ebe63600.
//
// Solidity: event NodeBlocked(address node, uint256 blockNumber)
func (_Relay *RelayFilterer) FilterNodeBlocked(opts *bind.FilterOpts) (*RelayNodeBlockedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "NodeBlocked")
	if err != nil {
		return nil, err
	}
	return &RelayNodeBlockedIterator{contract: _Relay.contract, event: "NodeBlocked", logs: logs, sub: sub}, nil
}

// WatchNodeBlocked is a free log subscription operation binding the contract event 0x29894930b6f680a84bc3015c2ee88544ea90c73f564a4dba638e3c55ebe63600.
//
// Solidity: event NodeBlocked(address node, uint256 blockNumber)
func (_Relay *RelayFilterer) WatchNodeBlocked(opts *bind.WatchOpts, sink chan<- *RelayNodeBlocked) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "NodeBlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayNodeBlocked)
				if err := _Relay.contract.UnpackLog(event, "NodeBlocked", log); err != nil {
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

// ParseNodeBlocked is a log parse operation binding the contract event 0x29894930b6f680a84bc3015c2ee88544ea90c73f564a4dba638e3c55ebe63600.
//
// Solidity: event NodeBlocked(address node, uint256 blockNumber)
func (_Relay *RelayFilterer) ParseNodeBlocked(log types.Log) (*RelayNodeBlocked, error) {
	event := new(RelayNodeBlocked)
	if err := _Relay.contract.UnpackLog(event, "NodeBlocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayNodeDeletedIterator is returned from FilterNodeDeleted and is used to iterate over the raw logs and unpacked data for NodeDeleted events raised by the Relay contract.
type RelayNodeDeletedIterator struct {
	Event *RelayNodeDeleted // Event containing the contract specifics and raw log

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
func (it *RelayNodeDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayNodeDeleted)
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
		it.Event = new(RelayNodeDeleted)
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
func (it *RelayNodeDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayNodeDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayNodeDeleted represents a NodeDeleted event raised by the Relay contract.
type RelayNodeDeleted struct {
	OldNode common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeDeleted is a free log retrieval operation binding the contract event 0x1629bfc36423a1b4749d3fe1d6970b9d32d42bbee47dd5540670696ab6b9a4ad.
//
// Solidity: event NodeDeleted(address oldNode)
func (_Relay *RelayFilterer) FilterNodeDeleted(opts *bind.FilterOpts) (*RelayNodeDeletedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "NodeDeleted")
	if err != nil {
		return nil, err
	}
	return &RelayNodeDeletedIterator{contract: _Relay.contract, event: "NodeDeleted", logs: logs, sub: sub}, nil
}

// WatchNodeDeleted is a free log subscription operation binding the contract event 0x1629bfc36423a1b4749d3fe1d6970b9d32d42bbee47dd5540670696ab6b9a4ad.
//
// Solidity: event NodeDeleted(address oldNode)
func (_Relay *RelayFilterer) WatchNodeDeleted(opts *bind.WatchOpts, sink chan<- *RelayNodeDeleted) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "NodeDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayNodeDeleted)
				if err := _Relay.contract.UnpackLog(event, "NodeDeleted", log); err != nil {
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

// ParseNodeDeleted is a log parse operation binding the contract event 0x1629bfc36423a1b4749d3fe1d6970b9d32d42bbee47dd5540670696ab6b9a4ad.
//
// Solidity: event NodeDeleted(address oldNode)
func (_Relay *RelayFilterer) ParseNodeDeleted(log types.Log) (*RelayNodeDeleted, error) {
	event := new(RelayNodeDeleted)
	if err := _Relay.contract.UnpackLog(event, "NodeDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayParametersIterator is returned from FilterParameters and is used to iterate over the raw logs and unpacked data for Parameters events raised by the Relay contract.
type RelayParametersIterator struct {
	Event *RelayParameters // Event containing the contract specifics and raw log

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
func (it *RelayParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayParameters)
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
		it.Event = new(RelayParameters)
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
func (it *RelayParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayParameters represents a Parameters event raised by the Relay contract.
type RelayParameters struct {
	Nonce           *big.Int
	GasLimit        *big.Int
	To              common.Address
	DecodedFunction []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterParameters is a free log retrieval operation binding the contract event 0x7cad50c38b02aa7e6f3cbcd18e919ffab3c54a7d49c09ae59bb59bf3c2be6dfe.
//
// Solidity: event Parameters(uint256 nonce, uint256 gasLimit, address to, bytes decodedFunction)
func (_Relay *RelayFilterer) FilterParameters(opts *bind.FilterOpts) (*RelayParametersIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "Parameters")
	if err != nil {
		return nil, err
	}
	return &RelayParametersIterator{contract: _Relay.contract, event: "Parameters", logs: logs, sub: sub}, nil
}

// WatchParameters is a free log subscription operation binding the contract event 0x7cad50c38b02aa7e6f3cbcd18e919ffab3c54a7d49c09ae59bb59bf3c2be6dfe.
//
// Solidity: event Parameters(uint256 nonce, uint256 gasLimit, address to, bytes decodedFunction)
func (_Relay *RelayFilterer) WatchParameters(opts *bind.WatchOpts, sink chan<- *RelayParameters) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "Parameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayParameters)
				if err := _Relay.contract.UnpackLog(event, "Parameters", log); err != nil {
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

// ParseParameters is a log parse operation binding the contract event 0x7cad50c38b02aa7e6f3cbcd18e919ffab3c54a7d49c09ae59bb59bf3c2be6dfe.
//
// Solidity: event Parameters(uint256 nonce, uint256 gasLimit, address to, bytes decodedFunction)
func (_Relay *RelayFilterer) ParseParameters(log types.Log) (*RelayParameters, error) {
	event := new(RelayParameters)
	if err := _Relay.contract.UnpackLog(event, "Parameters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayRecalculatedIterator is returned from FilterRecalculated and is used to iterate over the raw logs and unpacked data for Recalculated events raised by the Relay contract.
type RelayRecalculatedIterator struct {
	Event *RelayRecalculated // Event containing the contract specifics and raw log

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
func (it *RelayRecalculatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayRecalculated)
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
		it.Event = new(RelayRecalculated)
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
func (it *RelayRecalculatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayRecalculatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayRecalculated represents a Recalculated event raised by the Relay contract.
type RelayRecalculated struct {
	Result bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecalculated is a free log retrieval operation binding the contract event 0xa37b1b27143f61d990cfcf145e7f5d21c4419700613094ab29654b7ac6c08724.
//
// Solidity: event Recalculated(bool result)
func (_Relay *RelayFilterer) FilterRecalculated(opts *bind.FilterOpts) (*RelayRecalculatedIterator, error) {

	logs, sub, err := _Relay.contract.FilterLogs(opts, "Recalculated")
	if err != nil {
		return nil, err
	}
	return &RelayRecalculatedIterator{contract: _Relay.contract, event: "Recalculated", logs: logs, sub: sub}, nil
}

// WatchRecalculated is a free log subscription operation binding the contract event 0xa37b1b27143f61d990cfcf145e7f5d21c4419700613094ab29654b7ac6c08724.
//
// Solidity: event Recalculated(bool result)
func (_Relay *RelayFilterer) WatchRecalculated(opts *bind.WatchOpts, sink chan<- *RelayRecalculated) (event.Subscription, error) {

	logs, sub, err := _Relay.contract.WatchLogs(opts, "Recalculated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayRecalculated)
				if err := _Relay.contract.UnpackLog(event, "Recalculated", log); err != nil {
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

// ParseRecalculated is a log parse operation binding the contract event 0xa37b1b27143f61d990cfcf145e7f5d21c4419700613094ab29654b7ac6c08724.
//
// Solidity: event Recalculated(bool result)
func (_Relay *RelayFilterer) ParseRecalculated(log types.Log) (*RelayRecalculated, error) {
	event := new(RelayRecalculated)
	if err := _Relay.contract.UnpackLog(event, "Recalculated", log); err != nil {
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

// RelayRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Relay contract.
type RelayRoleAdminChangedIterator struct {
	Event *RelayRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RelayRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayRoleAdminChanged)
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
		it.Event = new(RelayRoleAdminChanged)
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
func (it *RelayRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayRoleAdminChanged represents a RoleAdminChanged event raised by the Relay contract.
type RelayRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Relay *RelayFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RelayRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Relay.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RelayRoleAdminChangedIterator{contract: _Relay.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Relay *RelayFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RelayRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Relay.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayRoleAdminChanged)
				if err := _Relay.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Relay *RelayFilterer) ParseRoleAdminChanged(log types.Log) (*RelayRoleAdminChanged, error) {
	event := new(RelayRoleAdminChanged)
	if err := _Relay.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Relay contract.
type RelayRoleGrantedIterator struct {
	Event *RelayRoleGranted // Event containing the contract specifics and raw log

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
func (it *RelayRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayRoleGranted)
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
		it.Event = new(RelayRoleGranted)
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
func (it *RelayRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayRoleGranted represents a RoleGranted event raised by the Relay contract.
type RelayRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Relay *RelayFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RelayRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Relay.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RelayRoleGrantedIterator{contract: _Relay.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Relay *RelayFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RelayRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Relay.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayRoleGranted)
				if err := _Relay.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Relay *RelayFilterer) ParseRoleGranted(log types.Log) (*RelayRoleGranted, error) {
	event := new(RelayRoleGranted)
	if err := _Relay.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Relay contract.
type RelayRoleRevokedIterator struct {
	Event *RelayRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RelayRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayRoleRevoked)
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
		it.Event = new(RelayRoleRevoked)
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
func (it *RelayRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayRoleRevoked represents a RoleRevoked event raised by the Relay contract.
type RelayRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Relay *RelayFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RelayRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Relay.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RelayRoleRevokedIterator{contract: _Relay.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Relay *RelayFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RelayRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Relay.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayRoleRevoked)
				if err := _Relay.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Relay *RelayFilterer) ParseRoleRevoked(log types.Log) (*RelayRoleRevoked, error) {
	event := new(RelayRoleRevoked)
	if err := _Relay.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
	Executed bool
	Output   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransactionRelayed is a free log retrieval operation binding the contract event 0x548af85d7bc344f47cbfacdfce1ffea1ecd862e5e235ca9ec919e767c14049a8.
//
// Solidity: event TransactionRelayed(address indexed relay, address indexed from, address indexed to, bool executed, bytes output)
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

// WatchTransactionRelayed is a free log subscription operation binding the contract event 0x548af85d7bc344f47cbfacdfce1ffea1ecd862e5e235ca9ec919e767c14049a8.
//
// Solidity: event TransactionRelayed(address indexed relay, address indexed from, address indexed to, bool executed, bytes output)
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

// ParseTransactionRelayed is a log parse operation binding the contract event 0x548af85d7bc344f47cbfacdfce1ffea1ecd862e5e235ca9ec919e767c14049a8.
//
// Solidity: event TransactionRelayed(address indexed relay, address indexed from, address indexed to, bool executed, bytes output)
func (_Relay *RelayFilterer) ParseTransactionRelayed(log types.Log) (*RelayTransactionRelayed, error) {
	event := new(RelayTransactionRelayed)
	if err := _Relay.contract.UnpackLog(event, "TransactionRelayed", log); err != nil {
		return nil, err
	}
	return event, nil
}
