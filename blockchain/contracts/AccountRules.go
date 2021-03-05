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

// AccountABI is the input ABI used to generate the binding from.
const AccountABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"exitReadOnly\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"accountPermitted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getByIndex\",\"outputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAllowList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ON_CHAIN_PRIVACY_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargets\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"addTarget\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAccounts\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"transactionAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"destinationPermitted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"addAccounts\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"targetAllowList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_relayHub\",\"type\":\"address\"}],\"name\":\"setRelay\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"removeTarget\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enterReadOnly\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isReadOnly\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\"}],\"name\":\"addTargets\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ingressContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountAdded\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"AccountVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountAdded\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"AccountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountRemoved\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"AccountRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"targetAdded\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"TargetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"targetRemoved\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"TargetRemoved\",\"type\":\"event\"}]"

// Account is an auto generated Go binding around an Ethereum contract.
type Account struct {
	AccountCaller     // Read-only binding to the contract
	AccountTransactor // Write-only binding to the contract
	AccountFilterer   // Log filterer for contract events
}

// AccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountSession struct {
	Contract     *Account          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountCallerSession struct {
	Contract *AccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountTransactorSession struct {
	Contract     *AccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountRaw struct {
	Contract *Account // Generic contract binding to access the raw methods on
}

// AccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountCallerRaw struct {
	Contract *AccountCaller // Generic read-only contract binding to access the raw methods on
}

// AccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountTransactorRaw struct {
	Contract *AccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccount creates a new instance of Account, bound to a specific deployed contract.
func NewAccount(address common.Address, backend bind.ContractBackend) (*Account, error) {
	contract, err := bindAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Account{AccountCaller: AccountCaller{contract: contract}, AccountTransactor: AccountTransactor{contract: contract}, AccountFilterer: AccountFilterer{contract: contract}}, nil
}

// NewAccountCaller creates a new read-only instance of Account, bound to a specific deployed contract.
func NewAccountCaller(address common.Address, caller bind.ContractCaller) (*AccountCaller, error) {
	contract, err := bindAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountCaller{contract: contract}, nil
}

// NewAccountTransactor creates a new write-only instance of Account, bound to a specific deployed contract.
func NewAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountTransactor, error) {
	contract, err := bindAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountTransactor{contract: contract}, nil
}

// NewAccountFilterer creates a new log filterer instance of Account, bound to a specific deployed contract.
func NewAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountFilterer, error) {
	contract, err := bindAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountFilterer{contract: contract}, nil
}

// bindAccount binds a generic wrapper to an already deployed contract.
func bindAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Account *AccountRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Account.Contract.AccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Account *AccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.Contract.AccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Account *AccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Account.Contract.AccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Account *AccountCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Account.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Account *AccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Account *AccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Account.Contract.contract.Transact(opts, method, params...)
}

// ONCHAINPRIVACYADDRESS is a free data retrieval call binding the contract method 0x556dcfee.
//
// Solidity: function ON_CHAIN_PRIVACY_ADDRESS() view returns(address)
func (_Account *AccountCaller) ONCHAINPRIVACYADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Account.contract.Call(opts, out, "ON_CHAIN_PRIVACY_ADDRESS")
	return *ret0, err
}

// ONCHAINPRIVACYADDRESS is a free data retrieval call binding the contract method 0x556dcfee.
//
// Solidity: function ON_CHAIN_PRIVACY_ADDRESS() view returns(address)
func (_Account *AccountSession) ONCHAINPRIVACYADDRESS() (common.Address, error) {
	return _Account.Contract.ONCHAINPRIVACYADDRESS(&_Account.CallOpts)
}

// ONCHAINPRIVACYADDRESS is a free data retrieval call binding the contract method 0x556dcfee.
//
// Solidity: function ON_CHAIN_PRIVACY_ADDRESS() view returns(address)
func (_Account *AccountCallerSession) ONCHAINPRIVACYADDRESS() (common.Address, error) {
	return _Account.Contract.ONCHAINPRIVACYADDRESS(&_Account.CallOpts)
}

// AccountAllowList is a free data retrieval call binding the contract method 0x368c5edf.
//
// Solidity: function accountAllowList(uint256 ) view returns(address)
func (_Account *AccountCaller) AccountAllowList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Account.contract.Call(opts, out, "accountAllowList", arg0)
	return *ret0, err
}

// AccountAllowList is a free data retrieval call binding the contract method 0x368c5edf.
//
// Solidity: function accountAllowList(uint256 ) view returns(address)
func (_Account *AccountSession) AccountAllowList(arg0 *big.Int) (common.Address, error) {
	return _Account.Contract.AccountAllowList(&_Account.CallOpts, arg0)
}

// AccountAllowList is a free data retrieval call binding the contract method 0x368c5edf.
//
// Solidity: function accountAllowList(uint256 ) view returns(address)
func (_Account *AccountCallerSession) AccountAllowList(arg0 *big.Int) (common.Address, error) {
	return _Account.Contract.AccountAllowList(&_Account.CallOpts, arg0)
}

// AccountPermitted is a free data retrieval call binding the contract method 0x0f68f0b3.
//
// Solidity: function accountPermitted(address _account) view returns(bool)
func (_Account *AccountCaller) AccountPermitted(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Account.contract.Call(opts, out, "accountPermitted", _account)
	return *ret0, err
}

// AccountPermitted is a free data retrieval call binding the contract method 0x0f68f0b3.
//
// Solidity: function accountPermitted(address _account) view returns(bool)
func (_Account *AccountSession) AccountPermitted(_account common.Address) (bool, error) {
	return _Account.Contract.AccountPermitted(&_Account.CallOpts, _account)
}

// AccountPermitted is a free data retrieval call binding the contract method 0x0f68f0b3.
//
// Solidity: function accountPermitted(address _account) view returns(bool)
func (_Account *AccountCallerSession) AccountPermitted(_account common.Address) (bool, error) {
	return _Account.Contract.AccountPermitted(&_Account.CallOpts, _account)
}

// DestinationPermitted is a free data retrieval call binding the contract method 0x9b1877ad.
//
// Solidity: function destinationPermitted(address _target) view returns(bool)
func (_Account *AccountCaller) DestinationPermitted(opts *bind.CallOpts, _target common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Account.contract.Call(opts, out, "destinationPermitted", _target)
	return *ret0, err
}

// DestinationPermitted is a free data retrieval call binding the contract method 0x9b1877ad.
//
// Solidity: function destinationPermitted(address _target) view returns(bool)
func (_Account *AccountSession) DestinationPermitted(_target common.Address) (bool, error) {
	return _Account.Contract.DestinationPermitted(&_Account.CallOpts, _target)
}

// DestinationPermitted is a free data retrieval call binding the contract method 0x9b1877ad.
//
// Solidity: function destinationPermitted(address _target) view returns(bool)
func (_Account *AccountCallerSession) DestinationPermitted(_target common.Address) (bool, error) {
	return _Account.Contract.DestinationPermitted(&_Account.CallOpts, _target)
}

// GetAccounts is a free data retrieval call binding the contract method 0x8a48ac03.
//
// Solidity: function getAccounts() view returns(address[])
func (_Account *AccountCaller) GetAccounts(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Account.contract.Call(opts, out, "getAccounts")
	return *ret0, err
}

// GetAccounts is a free data retrieval call binding the contract method 0x8a48ac03.
//
// Solidity: function getAccounts() view returns(address[])
func (_Account *AccountSession) GetAccounts() ([]common.Address, error) {
	return _Account.Contract.GetAccounts(&_Account.CallOpts)
}

// GetAccounts is a free data retrieval call binding the contract method 0x8a48ac03.
//
// Solidity: function getAccounts() view returns(address[])
func (_Account *AccountCallerSession) GetAccounts() ([]common.Address, error) {
	return _Account.Contract.GetAccounts(&_Account.CallOpts)
}

// GetByIndex is a free data retrieval call binding the contract method 0x2d883a73.
//
// Solidity: function getByIndex(uint256 index) view returns(address account)
func (_Account *AccountCaller) GetByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Account.contract.Call(opts, out, "getByIndex", index)
	return *ret0, err
}

// GetByIndex is a free data retrieval call binding the contract method 0x2d883a73.
//
// Solidity: function getByIndex(uint256 index) view returns(address account)
func (_Account *AccountSession) GetByIndex(index *big.Int) (common.Address, error) {
	return _Account.Contract.GetByIndex(&_Account.CallOpts, index)
}

// GetByIndex is a free data retrieval call binding the contract method 0x2d883a73.
//
// Solidity: function getByIndex(uint256 index) view returns(address account)
func (_Account *AccountCallerSession) GetByIndex(index *big.Int) (common.Address, error) {
	return _Account.Contract.GetByIndex(&_Account.CallOpts, index)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Account *AccountCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Account.contract.Call(opts, out, "getContractVersion")
	return *ret0, err
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Account *AccountSession) GetContractVersion() (*big.Int, error) {
	return _Account.Contract.GetContractVersion(&_Account.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Account *AccountCallerSession) GetContractVersion() (*big.Int, error) {
	return _Account.Contract.GetContractVersion(&_Account.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_Account *AccountCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Account.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_Account *AccountSession) GetSize() (*big.Int, error) {
	return _Account.Contract.GetSize(&_Account.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_Account *AccountCallerSession) GetSize() (*big.Int, error) {
	return _Account.Contract.GetSize(&_Account.CallOpts)
}

// GetTargets is a free data retrieval call binding the contract method 0x63fe3b56.
//
// Solidity: function getTargets() view returns(address[])
func (_Account *AccountCaller) GetTargets(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Account.contract.Call(opts, out, "getTargets")
	return *ret0, err
}

// GetTargets is a free data retrieval call binding the contract method 0x63fe3b56.
//
// Solidity: function getTargets() view returns(address[])
func (_Account *AccountSession) GetTargets() ([]common.Address, error) {
	return _Account.Contract.GetTargets(&_Account.CallOpts)
}

// GetTargets is a free data retrieval call binding the contract method 0x63fe3b56.
//
// Solidity: function getTargets() view returns(address[])
func (_Account *AccountCallerSession) GetTargets() ([]common.Address, error) {
	return _Account.Contract.GetTargets(&_Account.CallOpts)
}

// IsReadOnly is a free data retrieval call binding the contract method 0xdc2a60f6.
//
// Solidity: function isReadOnly() view returns(bool)
func (_Account *AccountCaller) IsReadOnly(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Account.contract.Call(opts, out, "isReadOnly")
	return *ret0, err
}

// IsReadOnly is a free data retrieval call binding the contract method 0xdc2a60f6.
//
// Solidity: function isReadOnly() view returns(bool)
func (_Account *AccountSession) IsReadOnly() (bool, error) {
	return _Account.Contract.IsReadOnly(&_Account.CallOpts)
}

// IsReadOnly is a free data retrieval call binding the contract method 0xdc2a60f6.
//
// Solidity: function isReadOnly() view returns(bool)
func (_Account *AccountCallerSession) IsReadOnly() (bool, error) {
	return _Account.Contract.IsReadOnly(&_Account.CallOpts)
}

// TargetAllowList is a free data retrieval call binding the contract method 0xb8ab17b9.
//
// Solidity: function targetAllowList(uint256 ) view returns(address)
func (_Account *AccountCaller) TargetAllowList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Account.contract.Call(opts, out, "targetAllowList", arg0)
	return *ret0, err
}

// TargetAllowList is a free data retrieval call binding the contract method 0xb8ab17b9.
//
// Solidity: function targetAllowList(uint256 ) view returns(address)
func (_Account *AccountSession) TargetAllowList(arg0 *big.Int) (common.Address, error) {
	return _Account.Contract.TargetAllowList(&_Account.CallOpts, arg0)
}

// TargetAllowList is a free data retrieval call binding the contract method 0xb8ab17b9.
//
// Solidity: function targetAllowList(uint256 ) view returns(address)
func (_Account *AccountCallerSession) TargetAllowList(arg0 *big.Int) (common.Address, error) {
	return _Account.Contract.TargetAllowList(&_Account.CallOpts, arg0)
}

// AddAccount is a paid mutator transaction binding the contract method 0xe89b0e1e.
//
// Solidity: function addAccount(address account) returns(bool)
func (_Account *AccountTransactor) AddAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "addAccount", account)
}

// AddAccount is a paid mutator transaction binding the contract method 0xe89b0e1e.
//
// Solidity: function addAccount(address account) returns(bool)
func (_Account *AccountSession) AddAccount(account common.Address) (*types.Transaction, error) {
	return _Account.Contract.AddAccount(&_Account.TransactOpts, account)
}

// AddAccount is a paid mutator transaction binding the contract method 0xe89b0e1e.
//
// Solidity: function addAccount(address account) returns(bool)
func (_Account *AccountTransactorSession) AddAccount(account common.Address) (*types.Transaction, error) {
	return _Account.Contract.AddAccount(&_Account.TransactOpts, account)
}

// AddAccounts is a paid mutator transaction binding the contract method 0xac71abde.
//
// Solidity: function addAccounts(address[] accounts) returns(bool)
func (_Account *AccountTransactor) AddAccounts(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "addAccounts", accounts)
}

// AddAccounts is a paid mutator transaction binding the contract method 0xac71abde.
//
// Solidity: function addAccounts(address[] accounts) returns(bool)
func (_Account *AccountSession) AddAccounts(accounts []common.Address) (*types.Transaction, error) {
	return _Account.Contract.AddAccounts(&_Account.TransactOpts, accounts)
}

// AddAccounts is a paid mutator transaction binding the contract method 0xac71abde.
//
// Solidity: function addAccounts(address[] accounts) returns(bool)
func (_Account *AccountTransactorSession) AddAccounts(accounts []common.Address) (*types.Transaction, error) {
	return _Account.Contract.AddAccounts(&_Account.TransactOpts, accounts)
}

// AddTarget is a paid mutator transaction binding the contract method 0x6de45dee.
//
// Solidity: function addTarget(address target) returns(bool)
func (_Account *AccountTransactor) AddTarget(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "addTarget", target)
}

// AddTarget is a paid mutator transaction binding the contract method 0x6de45dee.
//
// Solidity: function addTarget(address target) returns(bool)
func (_Account *AccountSession) AddTarget(target common.Address) (*types.Transaction, error) {
	return _Account.Contract.AddTarget(&_Account.TransactOpts, target)
}

// AddTarget is a paid mutator transaction binding the contract method 0x6de45dee.
//
// Solidity: function addTarget(address target) returns(bool)
func (_Account *AccountTransactorSession) AddTarget(target common.Address) (*types.Transaction, error) {
	return _Account.Contract.AddTarget(&_Account.TransactOpts, target)
}

// AddTargets is a paid mutator transaction binding the contract method 0xde866db3.
//
// Solidity: function addTargets(address[] targets) returns(bool)
func (_Account *AccountTransactor) AddTargets(opts *bind.TransactOpts, targets []common.Address) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "addTargets", targets)
}

// AddTargets is a paid mutator transaction binding the contract method 0xde866db3.
//
// Solidity: function addTargets(address[] targets) returns(bool)
func (_Account *AccountSession) AddTargets(targets []common.Address) (*types.Transaction, error) {
	return _Account.Contract.AddTargets(&_Account.TransactOpts, targets)
}

// AddTargets is a paid mutator transaction binding the contract method 0xde866db3.
//
// Solidity: function addTargets(address[] targets) returns(bool)
func (_Account *AccountTransactorSession) AddTargets(targets []common.Address) (*types.Transaction, error) {
	return _Account.Contract.AddTargets(&_Account.TransactOpts, targets)
}

// EnterReadOnly is a paid mutator transaction binding the contract method 0xd8cec925.
//
// Solidity: function enterReadOnly() returns(bool)
func (_Account *AccountTransactor) EnterReadOnly(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "enterReadOnly")
}

// EnterReadOnly is a paid mutator transaction binding the contract method 0xd8cec925.
//
// Solidity: function enterReadOnly() returns(bool)
func (_Account *AccountSession) EnterReadOnly() (*types.Transaction, error) {
	return _Account.Contract.EnterReadOnly(&_Account.TransactOpts)
}

// EnterReadOnly is a paid mutator transaction binding the contract method 0xd8cec925.
//
// Solidity: function enterReadOnly() returns(bool)
func (_Account *AccountTransactorSession) EnterReadOnly() (*types.Transaction, error) {
	return _Account.Contract.EnterReadOnly(&_Account.TransactOpts)
}

// ExitReadOnly is a paid mutator transaction binding the contract method 0x0c6e35d5.
//
// Solidity: function exitReadOnly() returns(bool)
func (_Account *AccountTransactor) ExitReadOnly(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "exitReadOnly")
}

// ExitReadOnly is a paid mutator transaction binding the contract method 0x0c6e35d5.
//
// Solidity: function exitReadOnly() returns(bool)
func (_Account *AccountSession) ExitReadOnly() (*types.Transaction, error) {
	return _Account.Contract.ExitReadOnly(&_Account.TransactOpts)
}

// ExitReadOnly is a paid mutator transaction binding the contract method 0x0c6e35d5.
//
// Solidity: function exitReadOnly() returns(bool)
func (_Account *AccountTransactorSession) ExitReadOnly() (*types.Transaction, error) {
	return _Account.Contract.ExitReadOnly(&_Account.TransactOpts)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(address account) returns(bool)
func (_Account *AccountTransactor) RemoveAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "removeAccount", account)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(address account) returns(bool)
func (_Account *AccountSession) RemoveAccount(account common.Address) (*types.Transaction, error) {
	return _Account.Contract.RemoveAccount(&_Account.TransactOpts, account)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(address account) returns(bool)
func (_Account *AccountTransactorSession) RemoveAccount(account common.Address) (*types.Transaction, error) {
	return _Account.Contract.RemoveAccount(&_Account.TransactOpts, account)
}

// RemoveTarget is a paid mutator transaction binding the contract method 0xd5d7ff3c.
//
// Solidity: function removeTarget(address target) returns(bool)
func (_Account *AccountTransactor) RemoveTarget(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "removeTarget", target)
}

// RemoveTarget is a paid mutator transaction binding the contract method 0xd5d7ff3c.
//
// Solidity: function removeTarget(address target) returns(bool)
func (_Account *AccountSession) RemoveTarget(target common.Address) (*types.Transaction, error) {
	return _Account.Contract.RemoveTarget(&_Account.TransactOpts, target)
}

// RemoveTarget is a paid mutator transaction binding the contract method 0xd5d7ff3c.
//
// Solidity: function removeTarget(address target) returns(bool)
func (_Account *AccountTransactorSession) RemoveTarget(target common.Address) (*types.Transaction, error) {
	return _Account.Contract.RemoveTarget(&_Account.TransactOpts, target)
}

// SetRelay is a paid mutator transaction binding the contract method 0xc805f68b.
//
// Solidity: function setRelay(address _relayHub) returns(bool)
func (_Account *AccountTransactor) SetRelay(opts *bind.TransactOpts, _relayHub common.Address) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "setRelay", _relayHub)
}

// SetRelay is a paid mutator transaction binding the contract method 0xc805f68b.
//
// Solidity: function setRelay(address _relayHub) returns(bool)
func (_Account *AccountSession) SetRelay(_relayHub common.Address) (*types.Transaction, error) {
	return _Account.Contract.SetRelay(&_Account.TransactOpts, _relayHub)
}

// SetRelay is a paid mutator transaction binding the contract method 0xc805f68b.
//
// Solidity: function setRelay(address _relayHub) returns(bool)
func (_Account *AccountTransactorSession) SetRelay(_relayHub common.Address) (*types.Transaction, error) {
	return _Account.Contract.SetRelay(&_Account.TransactOpts, _relayHub)
}

// TransactionAllowed is a paid mutator transaction binding the contract method 0x936421d5.
//
// Solidity: function transactionAllowed(address sender, address target, uint256 , uint256 , uint256 , bytes ) returns(bool)
func (_Account *AccountTransactor) TransactionAllowed(opts *bind.TransactOpts, sender common.Address, target common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "transactionAllowed", sender, target, arg2, arg3, arg4, arg5)
}

// TransactionAllowed is a paid mutator transaction binding the contract method 0x936421d5.
//
// Solidity: function transactionAllowed(address sender, address target, uint256 , uint256 , uint256 , bytes ) returns(bool)
func (_Account *AccountSession) TransactionAllowed(sender common.Address, target common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _Account.Contract.TransactionAllowed(&_Account.TransactOpts, sender, target, arg2, arg3, arg4, arg5)
}

// TransactionAllowed is a paid mutator transaction binding the contract method 0x936421d5.
//
// Solidity: function transactionAllowed(address sender, address target, uint256 , uint256 , uint256 , bytes ) returns(bool)
func (_Account *AccountTransactorSession) TransactionAllowed(sender common.Address, target common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _Account.Contract.TransactionAllowed(&_Account.TransactOpts, sender, target, arg2, arg3, arg4, arg5)
}

// AccountAccountAddedIterator is returned from FilterAccountAdded and is used to iterate over the raw logs and unpacked data for AccountAdded events raised by the Account contract.
type AccountAccountAddedIterator struct {
	Event *AccountAccountAdded // Event containing the contract specifics and raw log

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
func (it *AccountAccountAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountAccountAdded)
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
		it.Event = new(AccountAccountAdded)
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
func (it *AccountAccountAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountAccountAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountAccountAdded represents a AccountAdded event raised by the Account contract.
type AccountAccountAdded struct {
	AccountAdded   bool
	AccountAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountAdded is a free log retrieval operation binding the contract event 0xe39119db1877d19873ffb44540ac1dbd9ca72da5413d351392ce967885031aa4.
//
// Solidity: event AccountAdded(bool accountAdded, address accountAddress)
func (_Account *AccountFilterer) FilterAccountAdded(opts *bind.FilterOpts) (*AccountAccountAddedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "AccountAdded")
	if err != nil {
		return nil, err
	}
	return &AccountAccountAddedIterator{contract: _Account.contract, event: "AccountAdded", logs: logs, sub: sub}, nil
}

// WatchAccountAdded is a free log subscription operation binding the contract event 0xe39119db1877d19873ffb44540ac1dbd9ca72da5413d351392ce967885031aa4.
//
// Solidity: event AccountAdded(bool accountAdded, address accountAddress)
func (_Account *AccountFilterer) WatchAccountAdded(opts *bind.WatchOpts, sink chan<- *AccountAccountAdded) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "AccountAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountAccountAdded)
				if err := _Account.contract.UnpackLog(event, "AccountAdded", log); err != nil {
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

// ParseAccountAdded is a log parse operation binding the contract event 0xe39119db1877d19873ffb44540ac1dbd9ca72da5413d351392ce967885031aa4.
//
// Solidity: event AccountAdded(bool accountAdded, address accountAddress)
func (_Account *AccountFilterer) ParseAccountAdded(log types.Log) (*AccountAccountAdded, error) {
	event := new(AccountAccountAdded)
	if err := _Account.contract.UnpackLog(event, "AccountAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountAccountRemovedIterator is returned from FilterAccountRemoved and is used to iterate over the raw logs and unpacked data for AccountRemoved events raised by the Account contract.
type AccountAccountRemovedIterator struct {
	Event *AccountAccountRemoved // Event containing the contract specifics and raw log

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
func (it *AccountAccountRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountAccountRemoved)
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
		it.Event = new(AccountAccountRemoved)
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
func (it *AccountAccountRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountAccountRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountAccountRemoved represents a AccountRemoved event raised by the Account contract.
type AccountAccountRemoved struct {
	AccountRemoved bool
	AccountAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountRemoved is a free log retrieval operation binding the contract event 0xf9cfee605255fa33725274ecbb6100757021c2c1679bb4538c8fad791751a4d9.
//
// Solidity: event AccountRemoved(bool accountRemoved, address accountAddress)
func (_Account *AccountFilterer) FilterAccountRemoved(opts *bind.FilterOpts) (*AccountAccountRemovedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "AccountRemoved")
	if err != nil {
		return nil, err
	}
	return &AccountAccountRemovedIterator{contract: _Account.contract, event: "AccountRemoved", logs: logs, sub: sub}, nil
}

// WatchAccountRemoved is a free log subscription operation binding the contract event 0xf9cfee605255fa33725274ecbb6100757021c2c1679bb4538c8fad791751a4d9.
//
// Solidity: event AccountRemoved(bool accountRemoved, address accountAddress)
func (_Account *AccountFilterer) WatchAccountRemoved(opts *bind.WatchOpts, sink chan<- *AccountAccountRemoved) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "AccountRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountAccountRemoved)
				if err := _Account.contract.UnpackLog(event, "AccountRemoved", log); err != nil {
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

// ParseAccountRemoved is a log parse operation binding the contract event 0xf9cfee605255fa33725274ecbb6100757021c2c1679bb4538c8fad791751a4d9.
//
// Solidity: event AccountRemoved(bool accountRemoved, address accountAddress)
func (_Account *AccountFilterer) ParseAccountRemoved(log types.Log) (*AccountAccountRemoved, error) {
	event := new(AccountAccountRemoved)
	if err := _Account.contract.UnpackLog(event, "AccountRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountAccountVerifiedIterator is returned from FilterAccountVerified and is used to iterate over the raw logs and unpacked data for AccountVerified events raised by the Account contract.
type AccountAccountVerifiedIterator struct {
	Event *AccountAccountVerified // Event containing the contract specifics and raw log

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
func (it *AccountAccountVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountAccountVerified)
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
		it.Event = new(AccountAccountVerified)
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
func (it *AccountAccountVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountAccountVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountAccountVerified represents a AccountVerified event raised by the Account contract.
type AccountAccountVerified struct {
	AccountAdded   bool
	AccountAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountVerified is a free log retrieval operation binding the contract event 0x2ef3463cbe37b822a67582095e37c3be83bb71bfb7659f8e57f1cca0fa41be88.
//
// Solidity: event AccountVerified(bool accountAdded, address accountAddress)
func (_Account *AccountFilterer) FilterAccountVerified(opts *bind.FilterOpts) (*AccountAccountVerifiedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "AccountVerified")
	if err != nil {
		return nil, err
	}
	return &AccountAccountVerifiedIterator{contract: _Account.contract, event: "AccountVerified", logs: logs, sub: sub}, nil
}

// WatchAccountVerified is a free log subscription operation binding the contract event 0x2ef3463cbe37b822a67582095e37c3be83bb71bfb7659f8e57f1cca0fa41be88.
//
// Solidity: event AccountVerified(bool accountAdded, address accountAddress)
func (_Account *AccountFilterer) WatchAccountVerified(opts *bind.WatchOpts, sink chan<- *AccountAccountVerified) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "AccountVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountAccountVerified)
				if err := _Account.contract.UnpackLog(event, "AccountVerified", log); err != nil {
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

// ParseAccountVerified is a log parse operation binding the contract event 0x2ef3463cbe37b822a67582095e37c3be83bb71bfb7659f8e57f1cca0fa41be88.
//
// Solidity: event AccountVerified(bool accountAdded, address accountAddress)
func (_Account *AccountFilterer) ParseAccountVerified(log types.Log) (*AccountAccountVerified, error) {
	event := new(AccountAccountVerified)
	if err := _Account.contract.UnpackLog(event, "AccountVerified", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountTargetAddedIterator is returned from FilterTargetAdded and is used to iterate over the raw logs and unpacked data for TargetAdded events raised by the Account contract.
type AccountTargetAddedIterator struct {
	Event *AccountTargetAdded // Event containing the contract specifics and raw log

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
func (it *AccountTargetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountTargetAdded)
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
		it.Event = new(AccountTargetAdded)
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
func (it *AccountTargetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountTargetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountTargetAdded represents a TargetAdded event raised by the Account contract.
type AccountTargetAdded struct {
	TargetAdded    bool
	AccountAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTargetAdded is a free log retrieval operation binding the contract event 0x32a4165e84e7995883925a7df45e898fcc735f3b8d1c4f481f5807d722b4e161.
//
// Solidity: event TargetAdded(bool targetAdded, address accountAddress)
func (_Account *AccountFilterer) FilterTargetAdded(opts *bind.FilterOpts) (*AccountTargetAddedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "TargetAdded")
	if err != nil {
		return nil, err
	}
	return &AccountTargetAddedIterator{contract: _Account.contract, event: "TargetAdded", logs: logs, sub: sub}, nil
}

// WatchTargetAdded is a free log subscription operation binding the contract event 0x32a4165e84e7995883925a7df45e898fcc735f3b8d1c4f481f5807d722b4e161.
//
// Solidity: event TargetAdded(bool targetAdded, address accountAddress)
func (_Account *AccountFilterer) WatchTargetAdded(opts *bind.WatchOpts, sink chan<- *AccountTargetAdded) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "TargetAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountTargetAdded)
				if err := _Account.contract.UnpackLog(event, "TargetAdded", log); err != nil {
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

// ParseTargetAdded is a log parse operation binding the contract event 0x32a4165e84e7995883925a7df45e898fcc735f3b8d1c4f481f5807d722b4e161.
//
// Solidity: event TargetAdded(bool targetAdded, address accountAddress)
func (_Account *AccountFilterer) ParseTargetAdded(log types.Log) (*AccountTargetAdded, error) {
	event := new(AccountTargetAdded)
	if err := _Account.contract.UnpackLog(event, "TargetAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountTargetRemovedIterator is returned from FilterTargetRemoved and is used to iterate over the raw logs and unpacked data for TargetRemoved events raised by the Account contract.
type AccountTargetRemovedIterator struct {
	Event *AccountTargetRemoved // Event containing the contract specifics and raw log

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
func (it *AccountTargetRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountTargetRemoved)
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
		it.Event = new(AccountTargetRemoved)
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
func (it *AccountTargetRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountTargetRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountTargetRemoved represents a TargetRemoved event raised by the Account contract.
type AccountTargetRemoved struct {
	TargetRemoved  bool
	AccountAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTargetRemoved is a free log retrieval operation binding the contract event 0x4582e636d58982a8a89fa653199811b9a9e555cc5fa954d3d29ab6cc3c2fce1c.
//
// Solidity: event TargetRemoved(bool targetRemoved, address accountAddress)
func (_Account *AccountFilterer) FilterTargetRemoved(opts *bind.FilterOpts) (*AccountTargetRemovedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "TargetRemoved")
	if err != nil {
		return nil, err
	}
	return &AccountTargetRemovedIterator{contract: _Account.contract, event: "TargetRemoved", logs: logs, sub: sub}, nil
}

// WatchTargetRemoved is a free log subscription operation binding the contract event 0x4582e636d58982a8a89fa653199811b9a9e555cc5fa954d3d29ab6cc3c2fce1c.
//
// Solidity: event TargetRemoved(bool targetRemoved, address accountAddress)
func (_Account *AccountFilterer) WatchTargetRemoved(opts *bind.WatchOpts, sink chan<- *AccountTargetRemoved) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "TargetRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountTargetRemoved)
				if err := _Account.contract.UnpackLog(event, "TargetRemoved", log); err != nil {
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

// ParseTargetRemoved is a log parse operation binding the contract event 0x4582e636d58982a8a89fa653199811b9a9e555cc5fa954d3d29ab6cc3c2fce1c.
//
// Solidity: event TargetRemoved(bool targetRemoved, address accountAddress)
func (_Account *AccountFilterer) ParseTargetRemoved(log types.Log) (*AccountTargetRemoved, error) {
	event := new(AccountTargetRemoved)
	if err := _Account.contract.UnpackLog(event, "TargetRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}
