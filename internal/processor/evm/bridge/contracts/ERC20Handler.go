// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Erc20HandlerMetaData contains all meta data concerning the Erc20Handler contract.
var Erc20HandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Erc20HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20HandlerMetaData.ABI instead.
var Erc20HandlerABI = Erc20HandlerMetaData.ABI

// Erc20Handler is an auto generated Go binding around an Ethereum contract.
type Erc20Handler struct {
	Erc20HandlerCaller     // Read-only binding to the contract
	Erc20HandlerTransactor // Write-only binding to the contract
	Erc20HandlerFilterer   // Log filterer for contract events
}

// Erc20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20HandlerSession struct {
	Contract     *Erc20Handler           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20HandlerCallerSession struct {
	Contract *Erc20HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc20HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20HandlerTransactorSession struct {
	Contract     *Erc20HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20HandlerRaw struct {
	Contract *Erc20Handler // Generic contract binding to access the raw methods on
}

// Erc20HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20HandlerCallerRaw struct {
	Contract *Erc20HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20HandlerTransactorRaw struct {
	Contract *Erc20HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20Handler creates a new instance of Erc20Handler, bound to a specific deployed contract.
func NewErc20Handler(address common.Address, backend bind.ContractBackend) (*Erc20Handler, error) {
	contract, err := bindErc20Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20Handler{Erc20HandlerCaller: Erc20HandlerCaller{contract: contract}, Erc20HandlerTransactor: Erc20HandlerTransactor{contract: contract}, Erc20HandlerFilterer: Erc20HandlerFilterer{contract: contract}}, nil
}

// NewErc20HandlerCaller creates a new read-only instance of Erc20Handler, bound to a specific deployed contract.
func NewErc20HandlerCaller(address common.Address, caller bind.ContractCaller) (*Erc20HandlerCaller, error) {
	contract, err := bindErc20Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20HandlerCaller{contract: contract}, nil
}

// NewErc20HandlerTransactor creates a new write-only instance of Erc20Handler, bound to a specific deployed contract.
func NewErc20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20HandlerTransactor, error) {
	contract, err := bindErc20Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20HandlerTransactor{contract: contract}, nil
}

// NewErc20HandlerFilterer creates a new log filterer instance of Erc20Handler, bound to a specific deployed contract.
func NewErc20HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20HandlerFilterer, error) {
	contract, err := bindErc20Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20HandlerFilterer{contract: contract}, nil
}

// bindErc20Handler binds a generic wrapper to an already deployed contract.
func bindErc20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Handler *Erc20HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Handler.Contract.Erc20HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Handler *Erc20HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Handler.Contract.Erc20HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Handler *Erc20HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Handler.Contract.Erc20HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Handler *Erc20HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Handler *Erc20HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Handler *Erc20HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Handler.Contract.contract.Transact(opts, method, params...)
}

// Erc20HandlerAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_Erc20Handler *Erc20HandlerCaller) Erc20HandlerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20Handler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20HandlerAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_Erc20Handler *Erc20HandlerSession) Erc20HandlerAddress() (common.Address, error) {
	return _Erc20Handler.Contract.Erc20HandlerAddress(&_Erc20Handler.CallOpts)
}

// Erc20HandlerAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_Erc20Handler *Erc20HandlerCallerSession) Erc20HandlerAddress() (common.Address, error) {
	return _Erc20Handler.Contract.Erc20HandlerAddress(&_Erc20Handler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_Erc20Handler *Erc20HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Erc20Handler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_Erc20Handler *Erc20HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _Erc20Handler.Contract.BurnList(&_Erc20Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_Erc20Handler *Erc20HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _Erc20Handler.Contract.BurnList(&_Erc20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_Erc20Handler *Erc20HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Erc20Handler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_Erc20Handler *Erc20HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _Erc20Handler.Contract.ContractWhitelist(&_Erc20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_Erc20Handler *Erc20HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _Erc20Handler.Contract.ContractWhitelist(&_Erc20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_Erc20Handler *Erc20HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Erc20Handler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_Erc20Handler *Erc20HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _Erc20Handler.Contract.ResourceIDToTokenContractAddress(&_Erc20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_Erc20Handler *Erc20HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _Erc20Handler.Contract.ResourceIDToTokenContractAddress(&_Erc20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_Erc20Handler *Erc20HandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Erc20Handler.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_Erc20Handler *Erc20HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _Erc20Handler.Contract.TokenContractAddressToResourceID(&_Erc20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_Erc20Handler *Erc20HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _Erc20Handler.Contract.TokenContractAddressToResourceID(&_Erc20Handler.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes)
func (_Erc20Handler *Erc20HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _Erc20Handler.contract.Transact(opts, "deposit", resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes)
func (_Erc20Handler *Erc20HandlerSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _Erc20Handler.Contract.Deposit(&_Erc20Handler.TransactOpts, resourceID, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0xb07e54bb.
//
// Solidity: function deposit(bytes32 resourceID, address depositer, bytes data) returns(bytes)
func (_Erc20Handler *Erc20HandlerTransactorSession) Deposit(resourceID [32]byte, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _Erc20Handler.Contract.Deposit(&_Erc20Handler.TransactOpts, resourceID, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_Erc20Handler *Erc20HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _Erc20Handler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_Erc20Handler *Erc20HandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _Erc20Handler.Contract.ExecuteProposal(&_Erc20Handler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_Erc20Handler *Erc20HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _Erc20Handler.Contract.ExecuteProposal(&_Erc20Handler.TransactOpts, resourceID, data)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_Erc20Handler *Erc20HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Erc20Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_Erc20Handler *Erc20HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _Erc20Handler.Contract.SetBurnable(&_Erc20Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_Erc20Handler *Erc20HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _Erc20Handler.Contract.SetBurnable(&_Erc20Handler.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_Erc20Handler *Erc20HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _Erc20Handler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_Erc20Handler *Erc20HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _Erc20Handler.Contract.SetResource(&_Erc20Handler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_Erc20Handler *Erc20HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _Erc20Handler.Contract.SetResource(&_Erc20Handler.TransactOpts, resourceID, contractAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_Erc20Handler *Erc20HandlerTransactor) Withdraw(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Erc20Handler.contract.Transact(opts, "withdraw", data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_Erc20Handler *Erc20HandlerSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _Erc20Handler.Contract.Withdraw(&_Erc20Handler.TransactOpts, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0968f264.
//
// Solidity: function withdraw(bytes data) returns()
func (_Erc20Handler *Erc20HandlerTransactorSession) Withdraw(data []byte) (*types.Transaction, error) {
	return _Erc20Handler.Contract.Withdraw(&_Erc20Handler.TransactOpts, data)
}
