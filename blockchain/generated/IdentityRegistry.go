/*
 * Copyright (C) 2019 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IdentityRegistryABI is the input ABI used to generate the binding from.
const IdentityRegistryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"registrationFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dex\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimEthers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_dexAddress\",\"type\":\"address\"},{\"name\":\"_regFee\",\"type\":\"uint256\"},{\"name\":\"_implementation\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identityHash\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_identityHash\",\"type\":\"address\"},{\"name\":\"_hubId\",\"type\":\"address\"}],\"name\":\"registerIdentity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProxyCode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_identityHash\",\"type\":\"address\"}],\"name\":\"getChannelAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_identityHash\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"changeRegistrationFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"transferCollectedFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IdentityRegistryBin is the compiled bytecode used for deploying new contracts.
const IdentityRegistryBin = `0x608060405234801561001057600080fd5b506040516080806117fe8339810180604052608081101561003057600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a381600481905550600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561015c57600080fd5b83600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156101d757600080fd5b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561025257600080fd5b80600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050611558806102a66000396000f3fe6080604052600436106100fe5760003560e01c80638f32d59b11610095578063e325253711610064578063e3252537146104e9578063e5e894121461053a578063f2fde38b146105cb578063f58c5b6e1461061c578063fc0c546a14610673576100fe565b80638f32d59b14610370578063b92234f11461039f578063c3c5a5471461042f578063df8de3e714610498576100fe565b80636931b550116100d15780636931b5501461027a578063715018a614610291578063781f5a83146102a85780638da5cb5b14610319576100fe565b806314c44e091461016c578063238e130a1461019757806350050769146101e8578063692058c214610223575b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f52656a656374696e672074782077697468206574686572732073656e7400000081525060200191505060405180910390fd5b34801561017857600080fd5b506101816106ca565b6040518082815260200191505060405180910390f35b3480156101a357600080fd5b506101e6600480360360208110156101ba57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106d0565b005b3480156101f457600080fd5b506102216004803603602081101561020b57600080fd5b81019080803590602001909291905050506107db565b005b34801561022f57600080fd5b506102386107f6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561028657600080fd5b5061028f61081c565b005b34801561029d57600080fd5b506102a66108fa565b005b3480156102b457600080fd5b50610317600480360360408110156102cb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109ca565b005b34801561032557600080fd5b5061032e610d2b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561037c57600080fd5b50610385610d54565b604051808215151515815260200191505060405180910390f35b3480156103ab57600080fd5b506103b4610dab565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103f45780820151818401526020810190506103d9565b50505050905090810190601f1680156104215780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561043b57600080fd5b5061047e6004803603602081101561045257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e6f565b604051808215151515815260200191505060405180910390f35b3480156104a457600080fd5b506104e7600480360360208110156104bb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e91565b005b3480156104f557600080fd5b506105386004803603602081101561050c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611092565b005b34801561054657600080fd5b506105896004803603602081101561055d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611277565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105d757600080fd5b5061061a600480360360208110156105ee57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611360565b005b34801561062857600080fd5b5061063161137d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561067f57600080fd5b506106886113a7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60045481565b6106d8610d54565b6106e157600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561071b57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad60405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6107e3610d54565b6107ec57600080fd5b8060048190555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561087857600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f193505050501580156108f7573d6000803e3d6000fd5b50565b610902610d54565b61090b57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a0457600080fd5b610a0d82610e6f565b15610a1757600080fd5b60006004541115610b3e57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd33306004546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610b0157600080fd5b505af1158015610b15573d6000803e3d6000fd5b505050506040513d6020811015610b2b57600080fd5b8101908080519060200190929190505050505b6000610b5f8373ffffffffffffffffffffffffffffffffffffffff166113cd565b90508073ffffffffffffffffffffffffffffffffffffffff1663f7013ef6600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1686866203f4806040518663ffffffff1660e01b8152600401808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200195505050505050600060405180830381600087803b158015610ccb57600080fd5b505af1158015610cdf573d6000803e3d6000fd5b505050508273ffffffffffffffffffffffffffffffffffffffff167f2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e960405160405180910390a2505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6060806040518060600160405280603781526020016114f66037913990506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660601b905060008090505b60148160ff161015610e6657818160ff1660148110610e1457fe5b1a60f81b838260140160ff1681518110610e2a57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050610df9565b50819250505090565b600080610e7b83611277565b90506000813b9050600081141592505050919050565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610eed57600080fd5b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610f6c57600080fd5b505afa158015610f80573d6000803e3d6000fd5b505050506040513d6020811015610f9657600080fd5b810190808051906020019092919050505090508173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561105257600080fd5b505af1158015611066573d6000803e3d6000fd5b505050506040513d602081101561107c57600080fd5b8101908080519060200190929190505050505050565b61109a610d54565b6110a357600080fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561114457600080fd5b505afa158015611158573d6000803e3d6000fd5b505050506040513d602081101561116e57600080fd5b810190808051906020019092919050505090506000811161118e57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561123757600080fd5b505af115801561124b573d6000803e3d6000fd5b505050506040513d602081101561126157600080fd5b8101908080519060200190929190505050505050565b600060ff60f81b308373ffffffffffffffffffffffffffffffffffffffff1660001b6112a1610dab565b8051906020012060405160200180857effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018381526020018281526020019450505050506040516020818303038152906040528051906020012060001c9050919050565b611368610d54565b61137157600080fd5b61137a816113fd565b50565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060606113da610dab565b9050838151602083016000f59150813b6113f357600080fd5b8192505050919050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561143757600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fe3d602d80600a3d3981f3363d3d373d3d3d363d73bebebebebebebebebebebebebebebebebebebebe5af43d82803e903d91602b57fd5bf3a165627a7a72305820419d88c004750fde7301a913128d9b18c397694207c4ef64876225c58f0155730029`

// DeployIdentityRegistry deploys a new Ethereum contract, binding an instance of IdentityRegistry to it.
func DeployIdentityRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAddress common.Address, _dexAddress common.Address, _regFee *big.Int, _implementation common.Address) (common.Address, *types.Transaction, *IdentityRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityRegistryBin), backend, _tokenAddress, _dexAddress, _regFee, _implementation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// IdentityRegistry is an auto generated Go binding around an Ethereum contract.
type IdentityRegistry struct {
	IdentityRegistryCaller     // Read-only binding to the contract
	IdentityRegistryTransactor // Write-only binding to the contract
	IdentityRegistryFilterer   // Log filterer for contract events
}

// IdentityRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityRegistrySession struct {
	Contract     *IdentityRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityRegistryCallerSession struct {
	Contract *IdentityRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IdentityRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityRegistryTransactorSession struct {
	Contract     *IdentityRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IdentityRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRegistryRaw struct {
	Contract *IdentityRegistry // Generic contract binding to access the raw methods on
}

// IdentityRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityRegistryCallerRaw struct {
	Contract *IdentityRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactorRaw struct {
	Contract *IdentityRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityRegistry creates a new instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistry(address common.Address, backend bind.ContractBackend) (*IdentityRegistry, error) {
	contract, err := bindIdentityRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// NewIdentityRegistryCaller creates a new read-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryCaller(address common.Address, caller bind.ContractCaller) (*IdentityRegistryCaller, error) {
	contract, err := bindIdentityRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryCaller{contract: contract}, nil
}

// NewIdentityRegistryTransactor creates a new write-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityRegistryTransactor, error) {
	contract, err := bindIdentityRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryTransactor{contract: contract}, nil
}

// NewIdentityRegistryFilterer creates a new log filterer instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityRegistryFilterer, error) {
	contract, err := bindIdentityRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFilterer{contract: contract}, nil
}

// bindIdentityRegistry binds a generic wrapper to an already deployed contract.
func bindIdentityRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.IdentityRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transact(opts, method, params...)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) Dex(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "dex")
	return *ret0, err
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() constant returns(address)
func (_IdentityRegistry *IdentityRegistrySession) Dex() (common.Address, error) {
	return _IdentityRegistry.Contract.Dex(&_IdentityRegistry.CallOpts)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) Dex() (common.Address, error) {
	return _IdentityRegistry.Contract.Dex(&_IdentityRegistry.CallOpts)
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe5e89412.
//
// Solidity: function getChannelAddress(address _identityHash) constant returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) GetChannelAddress(opts *bind.CallOpts, _identityHash common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "getChannelAddress", _identityHash)
	return *ret0, err
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe5e89412.
//
// Solidity: function getChannelAddress(address _identityHash) constant returns(address)
func (_IdentityRegistry *IdentityRegistrySession) GetChannelAddress(_identityHash common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.GetChannelAddress(&_IdentityRegistry.CallOpts, _identityHash)
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe5e89412.
//
// Solidity: function getChannelAddress(address _identityHash) constant returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetChannelAddress(_identityHash common.Address) (common.Address, error) {
	return _IdentityRegistry.Contract.GetChannelAddress(&_IdentityRegistry.CallOpts, _identityHash)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) GetFundsDestination(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "getFundsDestination")
	return *ret0, err
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() constant returns(address)
func (_IdentityRegistry *IdentityRegistrySession) GetFundsDestination() (common.Address, error) {
	return _IdentityRegistry.Contract.GetFundsDestination(&_IdentityRegistry.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetFundsDestination() (common.Address, error) {
	return _IdentityRegistry.Contract.GetFundsDestination(&_IdentityRegistry.CallOpts)
}

// GetProxyCode is a free data retrieval call binding the contract method 0xb92234f1.
//
// Solidity: function getProxyCode() constant returns(bytes)
func (_IdentityRegistry *IdentityRegistryCaller) GetProxyCode(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "getProxyCode")
	return *ret0, err
}

// GetProxyCode is a free data retrieval call binding the contract method 0xb92234f1.
//
// Solidity: function getProxyCode() constant returns(bytes)
func (_IdentityRegistry *IdentityRegistrySession) GetProxyCode() ([]byte, error) {
	return _IdentityRegistry.Contract.GetProxyCode(&_IdentityRegistry.CallOpts)
}

// GetProxyCode is a free data retrieval call binding the contract method 0xb92234f1.
//
// Solidity: function getProxyCode() constant returns(bytes)
func (_IdentityRegistry *IdentityRegistryCallerSession) GetProxyCode() ([]byte, error) {
	return _IdentityRegistry.Contract.GetProxyCode(&_IdentityRegistry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) IsOwner() (bool, error) {
	return _IdentityRegistry.Contract.IsOwner(&_IdentityRegistry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) IsOwner() (bool, error) {
	return _IdentityRegistry.Contract.IsOwner(&_IdentityRegistry.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identityHash) constant returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) IsRegistered(opts *bind.CallOpts, _identityHash common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "isRegistered", _identityHash)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identityHash) constant returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) IsRegistered(_identityHash common.Address) (bool, error) {
	return _IdentityRegistry.Contract.IsRegistered(&_IdentityRegistry.CallOpts, _identityHash)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identityHash) constant returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) IsRegistered(_identityHash common.Address) (bool, error) {
	return _IdentityRegistry.Contract.IsRegistered(&_IdentityRegistry.CallOpts, _identityHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IdentityRegistry *IdentityRegistrySession) Owner() (common.Address, error) {
	return _IdentityRegistry.Contract.Owner(&_IdentityRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) Owner() (common.Address, error) {
	return _IdentityRegistry.Contract.Owner(&_IdentityRegistry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistryCaller) RegistrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "registrationFee")
	return *ret0, err
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistrySession) RegistrationFee() (*big.Int, error) {
	return _IdentityRegistry.Contract.RegistrationFee(&_IdentityRegistry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_IdentityRegistry *IdentityRegistryCallerSession) RegistrationFee() (*big.Int, error) {
	return _IdentityRegistry.Contract.RegistrationFee(&_IdentityRegistry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IdentityRegistry *IdentityRegistrySession) Token() (common.Address, error) {
	return _IdentityRegistry.Contract.Token(&_IdentityRegistry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) Token() (common.Address, error) {
	return _IdentityRegistry.Contract.Token(&_IdentityRegistry.CallOpts)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(uint256 _newFee) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) ChangeRegistrationFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "changeRegistrationFee", _newFee)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(uint256 _newFee) returns()
func (_IdentityRegistry *IdentityRegistrySession) ChangeRegistrationFee(_newFee *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.ChangeRegistrationFee(&_IdentityRegistry.TransactOpts, _newFee)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(uint256 _newFee) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) ChangeRegistrationFee(_newFee *big.Int) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.ChangeRegistrationFee(&_IdentityRegistry.TransactOpts, _newFee)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_IdentityRegistry *IdentityRegistryTransactor) ClaimEthers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "claimEthers")
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_IdentityRegistry *IdentityRegistrySession) ClaimEthers() (*types.Transaction, error) {
	return _IdentityRegistry.Contract.ClaimEthers(&_IdentityRegistry.TransactOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) ClaimEthers() (*types.Transaction, error) {
	return _IdentityRegistry.Contract.ClaimEthers(&_IdentityRegistry.TransactOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address token) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) ClaimTokens(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "claimTokens", token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address token) returns()
func (_IdentityRegistry *IdentityRegistrySession) ClaimTokens(token common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.ClaimTokens(&_IdentityRegistry.TransactOpts, token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address token) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) ClaimTokens(token common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.ClaimTokens(&_IdentityRegistry.TransactOpts, token)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x781f5a83.
//
// Solidity: function registerIdentity(address _identityHash, address _hubId) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) RegisterIdentity(opts *bind.TransactOpts, _identityHash common.Address, _hubId common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "registerIdentity", _identityHash, _hubId)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x781f5a83.
//
// Solidity: function registerIdentity(address _identityHash, address _hubId) returns()
func (_IdentityRegistry *IdentityRegistrySession) RegisterIdentity(_identityHash common.Address, _hubId common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.RegisterIdentity(&_IdentityRegistry.TransactOpts, _identityHash, _hubId)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x781f5a83.
//
// Solidity: function registerIdentity(address _identityHash, address _hubId) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) RegisterIdentity(_identityHash common.Address, _hubId common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.RegisterIdentity(&_IdentityRegistry.TransactOpts, _identityHash, _hubId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityRegistry *IdentityRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityRegistry *IdentityRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _IdentityRegistry.Contract.RenounceOwnership(&_IdentityRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IdentityRegistry.Contract.RenounceOwnership(&_IdentityRegistry.TransactOpts)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) SetFundsDestination(opts *bind.TransactOpts, _newDestination common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "setFundsDestination", _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_IdentityRegistry *IdentityRegistrySession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetFundsDestination(&_IdentityRegistry.TransactOpts, _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.SetFundsDestination(&_IdentityRegistry.TransactOpts, _newDestination)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address _beneficiary) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) TransferCollectedFeeTo(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "transferCollectedFeeTo", _beneficiary)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address _beneficiary) returns()
func (_IdentityRegistry *IdentityRegistrySession) TransferCollectedFeeTo(_beneficiary common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.TransferCollectedFeeTo(&_IdentityRegistry.TransactOpts, _beneficiary)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address _beneficiary) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) TransferCollectedFeeTo(_beneficiary common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.TransferCollectedFeeTo(&_IdentityRegistry.TransactOpts, _beneficiary)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityRegistry *IdentityRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.TransferOwnership(&_IdentityRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.TransferOwnership(&_IdentityRegistry.TransactOpts, newOwner)
}

// IdentityRegistryDestinationChangedIterator is returned from FilterDestinationChanged and is used to iterate over the raw logs and unpacked data for DestinationChanged events raised by the IdentityRegistry contract.
type IdentityRegistryDestinationChangedIterator struct {
	Event *IdentityRegistryDestinationChanged // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryDestinationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryDestinationChanged)
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
		it.Event = new(IdentityRegistryDestinationChanged)
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
func (it *IdentityRegistryDestinationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryDestinationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryDestinationChanged represents a DestinationChanged event raised by the IdentityRegistry contract.
type IdentityRegistryDestinationChanged struct {
	PreviousDestination common.Address
	NewDestination      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDestinationChanged is a free log retrieval operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterDestinationChanged(opts *bind.FilterOpts, previousDestination []common.Address, newDestination []common.Address) (*IdentityRegistryDestinationChangedIterator, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryDestinationChangedIterator{contract: _IdentityRegistry.contract, event: "DestinationChanged", logs: logs, sub: sub}, nil
}

// WatchDestinationChanged is a free log subscription operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchDestinationChanged(opts *bind.WatchOpts, sink chan<- *IdentityRegistryDestinationChanged, previousDestination []common.Address, newDestination []common.Address) (event.Subscription, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryDestinationChanged)
				if err := _IdentityRegistry.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
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

// IdentityRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IdentityRegistry contract.
type IdentityRegistryOwnershipTransferredIterator struct {
	Event *IdentityRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryOwnershipTransferred)
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
		it.Event = new(IdentityRegistryOwnershipTransferred)
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
func (it *IdentityRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the IdentityRegistry contract.
type IdentityRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentityRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryOwnershipTransferredIterator{contract: _IdentityRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdentityRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryOwnershipTransferred)
				if err := _IdentityRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// IdentityRegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the IdentityRegistry contract.
type IdentityRegistryRegisteredIterator struct {
	Event *IdentityRegistryRegistered // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryRegistered)
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
		it.Event = new(IdentityRegistryRegistered)
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
func (it *IdentityRegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryRegistered represents a Registered event raised by the IdentityRegistry contract.
type IdentityRegistryRegistered struct {
	IdentityHash common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e9.
//
// Solidity: event Registered(address indexed identityHash)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterRegistered(opts *bind.FilterOpts, identityHash []common.Address) (*IdentityRegistryRegisteredIterator, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "Registered", identityHashRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryRegisteredIterator{contract: _IdentityRegistry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e9.
//
// Solidity: event Registered(address indexed identityHash)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *IdentityRegistryRegistered, identityHash []common.Address) (event.Subscription, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "Registered", identityHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryRegistered)
				if err := _IdentityRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
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
