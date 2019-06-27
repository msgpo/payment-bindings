/*
 * Copyright (C) 2019 The "MysteriumNetwork/payment-bindings" Authors.
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

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"registrationFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDestination\",\"type\":\"address\"}],\"name\":\"setFundsDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dex\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimEthers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimalAccountantStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountantImplementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFundsDestination\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountants\",\"outputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_dexAddress\",\"type\":\"address\"},{\"name\":\"_channelImplementation\",\"type\":\"address\"},{\"name\":\"_accountantImplementation\",\"type\":\"address\"},{\"name\":\"_regFee\",\"type\":\"uint256\"},{\"name\":\"_minimalAccountantStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identityHash\",\"type\":\"address\"}],\"name\":\"RegisteredIdentity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountantId\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"accountantOperator\",\"type\":\"address\"}],\"name\":\"RegisteredAccountant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousDestination\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newDestination\",\"type\":\"address\"}],\"name\":\"DestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_identityHash\",\"type\":\"address\"},{\"name\":\"_accountantId\",\"type\":\"address\"},{\"name\":\"_loanAmount\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"registerIdentity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_accountantOperator\",\"type\":\"address\"},{\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"registerAccountant\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_identityHash\",\"type\":\"address\"}],\"name\":\"getChannelAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_accountantOperator\",\"type\":\"address\"}],\"name\":\"getAccountantAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"getProxyCode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_identityHash\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_accountantId\",\"type\":\"address\"}],\"name\":\"isAccountant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_accountantId\",\"type\":\"address\"}],\"name\":\"isActiveAccountant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"changeRegistrationFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"transferCollectedFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
const RegistryBin = `0x60806040523480156200001157600080fd5b5060405162002f3638038062002f36833981810160405260c08110156200003757600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a38160048190555080600581905550600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156200017f57600080fd5b85600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415620001fb57600080fd5b84600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156200027757600080fd5b83600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415620002f357600080fd5b82600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050505050612bec806200034a6000396000f3fe6080604052600436106101665760003560e01c80638f32d59b116100d1578063e5e894121161008a578063f58c5b6e11610064578063f58c5b6e14610963578063f595cfd2146109ba578063f707fb4a14610a4b578063fc0c546a14610ae357610166565b8063e5e8941214610826578063f2fde38b146108b7578063f4c1a1f61461090857610166565b80638f32d59b146105cb5780639a3ce274146105fa578063ab86721314610651578063c3c5a5471461071b578063df8de3e714610784578063e3252537146107d557610166565b8063692058c211610123578063692058c2146104995780636931b550146104f0578063715018a614610507578063817b1cd21461051e578063824b09d6146105495780638da5cb5b1461057457610166565b806314c44e09146101d457806319233600146101ff5780631a3d9a591461033b578063238e130a146103a45780632a33ddbd146103f5578063500507691461045e575b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f52656a656374696e672074782077697468206574686572732073656e7400000081525060200191505060405180910390fd5b3480156101e057600080fd5b506101e9610b3a565b6040518082815260200191505060405180910390f35b34801561020b57600080fd5b50610339600480360360c081101561022257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156102b357600080fd5b8201836020820111156102c557600080fd5b803590602001918460018302840111640100000000831117156102e757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610b40565b005b34801561034757600080fd5b5061038a6004803603602081101561035e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506114a7565b604051808215151515815260200191505060405180910390f35b3480156103b057600080fd5b506103f3600480360360208110156103c757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506114f6565b005b34801561040157600080fd5b506104446004803603602081101561041857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061166a565b604051808215151515815260200191505060405180910390f35b34801561046a57600080fd5b506104976004803603602081101561048157600080fd5b81019080803590602001909291905050506116f4565b005b3480156104a557600080fd5b506104ae611778565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104fc57600080fd5b5061050561179e565b005b34801561051357600080fd5b5061051c61187c565b005b34801561052a57600080fd5b506105336119b5565b6040518082815260200191505060405180910390f35b34801561055557600080fd5b5061055e6119bb565b6040518082815260200191505060405180910390f35b34801561058057600080fd5b506105896119c1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105d757600080fd5b506105e06119ea565b604051808215151515815260200191505060405180910390f35b34801561060657600080fd5b5061060f611a41565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561065d57600080fd5b506106a06004803603602081101561067457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611a67565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106e05780820151818401526020810190506106c5565b50505050905090810190601f16801561070d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561072757600080fd5b5061076a6004803603602081101561073e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b0b565b604051808215151515815260200191505060405180910390f35b34801561079057600080fd5b506107d3600480360360208110156107a757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b2d565b005b3480156107e157600080fd5b50610824600480360360208110156107f857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611dd5565b005b34801561083257600080fd5b506108756004803603602081101561084957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612023565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156108c357600080fd5b50610906600480360360208110156108da57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612082565b005b34801561091457600080fd5b506109616004803603604081101561092b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612108565b005b34801561096f57600080fd5b5061097861257f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156109c657600080fd5b50610a09600480360360208110156109dd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506125a9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610a5757600080fd5b50610a9a60048036036020811015610a6e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612608565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b348015610aef57600080fd5b50610af861264c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60045481565b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415610b7a57600080fd5b610b8386611b0b565b15610bd9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612b12602a913960400191505060405180910390fd5b610be2856114a7565b610c37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180612b5d6025913960400191505060405180910390fd5b6000610d5682308989898989604051602001808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140196505050505050506040516020818303038152906040528051906020012061267290919063ffffffff16565b90508673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610ddc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180612a3b6024913960400191505060405180910390fd5b6000610e0585610df78860045461277690919063ffffffff16565b61277690919063ffffffff16565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231610e4e8a612023565b6040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610eae57600080fd5b505afa158015610ec2573d6000803e3d6000fd5b505050506040513d6020811015610ed857600080fd5b8101908080519060200190929190505050811115610ef557600080fd5b6000610f398973ffffffffffffffffffffffffffffffffffffffff16600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166127fe565b90508073ffffffffffffffffffffffffffffffffffffffff1663f7013ef6600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168c8c876040518663ffffffff1660e01b8152600401808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200195505050505050600060405180830381600087803b1580156110a257600080fd5b505af11580156110b6573d6000803e3d6000fd5b50505050600087111561136a57600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415611149576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612b3c6021913960400191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095ea7b389896040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156111f257600080fd5b505af1158015611206573d6000803e3d6000fd5b505050506040513d602081101561121c57600080fd5b8101908080519060200190929190505050611282576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526031815260200180612a856031913960400191505060405180910390fd5b8773ffffffffffffffffffffffffffffffffffffffff16630fb595c18a878a6040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200180602001828103825260008152602001602001945050505050600060405180830381600087803b15801561135157600080fd5b505af1158015611365573d6000803e3d6000fd5b505050505b600086111561145957600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33886040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561141c57600080fd5b505af1158015611430573d6000803e3d6000fd5b505050506040513d602081101561144657600080fd5b8101908080519060200190929190505050505b8873ffffffffffffffffffffffffffffffffffffffff167f16826e74d06e02bdda286d1820cf7f113495bfa8c8576c331511a3708902dfcc60405160405180910390a2505050505050505050565b600080600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015414159050919050565b6114fe6119ea565b611570576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156115aa57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad60405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060006116dd826125a9565b90506000813b905060008114159350505050919050565b6116fc6119ea565b61176e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b8060048190555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156117fa57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f19350505050158015611879573d6000803e3d6000fd5b50565b6118846119ea565b6118f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60065481565b60055481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606080604051806060016040528060378152602001612adb60379139905060008360601b905060008090505b60148160ff161015611b0057818160ff1660148110611aae57fe5b1a60f81b838260140160ff1681518110611ac457fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050611a93565b508192505050919050565b600080611b1783612023565b90506000813b9050600081141592505050919050565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611b8957600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611c30576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180612ab66025913960400191505060405180910390fd5b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015611caf57600080fd5b505afa158015611cc3573d6000803e3d6000fd5b505050506040513d6020811015611cd957600080fd5b810190808051906020019092919050505090508173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611d9557600080fd5b505af1158015611da9573d6000803e3d6000fd5b505050506040513d6020811015611dbf57600080fd5b8101908080519060200190929190505050505050565b611ddd6119ea565b611e4f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015611ef057600080fd5b505afa158015611f04573d6000803e3d6000fd5b505050506040513d6020811015611f1a57600080fd5b8101908080519060200190929190505050905060008111611f3a57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611fe357600080fd5b505af1158015611ff7573d6000803e3d6000fd5b505050506040513d602081101561200d57600080fd5b8101908080519060200190929190505050505050565b600080612051600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611a67565b80519060200120905061207a8373ffffffffffffffffffffffffffffffffffffffff1682612830565b915050919050565b61208a6119ea565b6120fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b612105816128f6565b50565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561214257600080fd5b60055481101561219d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526036815260200180612b826036913960400191505060405180910390fd5b60006121a8836125a9565b90506121b38161166a565b156121bd57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561229a57600080fd5b505af11580156122ae573d6000803e3d6000fd5b505050506040513d60208110156122c457600080fd5b8101908080519060200190929190505050506122eb8260065461277690919063ffffffff16565b60068190555060006123358473ffffffffffffffffffffffffffffffffffffffff16600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166127fe565b90508073ffffffffffffffffffffffffffffffffffffffff1663485cc955600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16866040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561240c57600080fd5b505af1158015612420573d6000803e3d6000fd5b5050505060405180604001604052808573ffffffffffffffffffffffffffffffffffffffff16815260200184815250600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101559050507fc9f77826eb4f5ea1c94b1b3ff214796f59d1c5f610af8a3ff52cd83e3eebf64e8185604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000806125d7600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611a67565b8051906020012090506126008373ffffffffffffffffffffffffffffffffffffffff1682612830565b915050919050565b60096020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060418251146126865760009050612770565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c11156126da5760009350505050612770565b601b8160ff16141580156126f25750601c8160ff1614155b156127035760009350505050612770565b60018682858560405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015612760573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b6000808284019050838110156127f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600080606061280c84611a67565b9050848151602083016000f59150813b61282557600080fd5b819250505092915050565b600060ff60f81b308460001b8460405160200180857effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018381526020018281526020019450505050506040516020818303038152906040528051906020012060001c905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561297c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180612a5f6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fe6861766520746f206265207369676e65642062792070726f706572206964656e746974794f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573736163636f756e74616e742073686f756c642067657420617070726f76616c20746f207472616e7366657220746f6b656e736e617469766520746f6b656e2066756e64732063616e2774206265207265636f76657265643d602d80600a3d3981f3363d3d373d3d3d363d73bebebebebebebebebebebebebebebebebebebebe5af43d82803e903d91602b57fd5bf36964656e7469747948617368206861766520746f206265206e6f7420726567697374657265642079657462656e65666963696172792063616e2774206265207a65726f206164647265737370726f7669646564206163636f756e74616e74206861766520746f206265206163746976656163636f756e74616e74206861766520746f207374616b65206174206c65617374206d696e696d616c207374616b6520616d6f756e74a265627a7a7230582098509c12a0a711f5759817faebfa2fb628e5a186ce8d6863e4ffc0eb15b4d52364736f6c63430005090032`

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAddress common.Address, _dexAddress common.Address, _channelImplementation common.Address, _accountantImplementation common.Address, _regFee *big.Int, _minimalAccountantStake *big.Int) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend, _tokenAddress, _dexAddress, _channelImplementation, _accountantImplementation, _regFee, _minimalAccountantStake)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// AccountantImplementation is a free data retrieval call binding the contract method 0x9a3ce274.
//
// Solidity: function accountantImplementation() constant returns(address)
func (_Registry *RegistryCaller) AccountantImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "accountantImplementation")
	return *ret0, err
}

// AccountantImplementation is a free data retrieval call binding the contract method 0x9a3ce274.
//
// Solidity: function accountantImplementation() constant returns(address)
func (_Registry *RegistrySession) AccountantImplementation() (common.Address, error) {
	return _Registry.Contract.AccountantImplementation(&_Registry.CallOpts)
}

// AccountantImplementation is a free data retrieval call binding the contract method 0x9a3ce274.
//
// Solidity: function accountantImplementation() constant returns(address)
func (_Registry *RegistryCallerSession) AccountantImplementation() (common.Address, error) {
	return _Registry.Contract.AccountantImplementation(&_Registry.CallOpts)
}

// Accountants is a free data retrieval call binding the contract method 0xf707fb4a.
//
// Solidity: function accountants(address ) constant returns(address operator, uint256 stake)
func (_Registry *RegistryCaller) Accountants(opts *bind.CallOpts, arg0 common.Address) (struct {
	Operator common.Address
	Stake    *big.Int
}, error) {
	ret := new(struct {
		Operator common.Address
		Stake    *big.Int
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "accountants", arg0)
	return *ret, err
}

// Accountants is a free data retrieval call binding the contract method 0xf707fb4a.
//
// Solidity: function accountants(address ) constant returns(address operator, uint256 stake)
func (_Registry *RegistrySession) Accountants(arg0 common.Address) (struct {
	Operator common.Address
	Stake    *big.Int
}, error) {
	return _Registry.Contract.Accountants(&_Registry.CallOpts, arg0)
}

// Accountants is a free data retrieval call binding the contract method 0xf707fb4a.
//
// Solidity: function accountants(address ) constant returns(address operator, uint256 stake)
func (_Registry *RegistryCallerSession) Accountants(arg0 common.Address) (struct {
	Operator common.Address
	Stake    *big.Int
}, error) {
	return _Registry.Contract.Accountants(&_Registry.CallOpts, arg0)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() constant returns(address)
func (_Registry *RegistryCaller) Dex(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "dex")
	return *ret0, err
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() constant returns(address)
func (_Registry *RegistrySession) Dex() (common.Address, error) {
	return _Registry.Contract.Dex(&_Registry.CallOpts)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() constant returns(address)
func (_Registry *RegistryCallerSession) Dex() (common.Address, error) {
	return _Registry.Contract.Dex(&_Registry.CallOpts)
}

// GetAccountantAddress is a free data retrieval call binding the contract method 0xf595cfd2.
//
// Solidity: function getAccountantAddress(address _accountantOperator) constant returns(address)
func (_Registry *RegistryCaller) GetAccountantAddress(opts *bind.CallOpts, _accountantOperator common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getAccountantAddress", _accountantOperator)
	return *ret0, err
}

// GetAccountantAddress is a free data retrieval call binding the contract method 0xf595cfd2.
//
// Solidity: function getAccountantAddress(address _accountantOperator) constant returns(address)
func (_Registry *RegistrySession) GetAccountantAddress(_accountantOperator common.Address) (common.Address, error) {
	return _Registry.Contract.GetAccountantAddress(&_Registry.CallOpts, _accountantOperator)
}

// GetAccountantAddress is a free data retrieval call binding the contract method 0xf595cfd2.
//
// Solidity: function getAccountantAddress(address _accountantOperator) constant returns(address)
func (_Registry *RegistryCallerSession) GetAccountantAddress(_accountantOperator common.Address) (common.Address, error) {
	return _Registry.Contract.GetAccountantAddress(&_Registry.CallOpts, _accountantOperator)
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe5e89412.
//
// Solidity: function getChannelAddress(address _identityHash) constant returns(address)
func (_Registry *RegistryCaller) GetChannelAddress(opts *bind.CallOpts, _identityHash common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getChannelAddress", _identityHash)
	return *ret0, err
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe5e89412.
//
// Solidity: function getChannelAddress(address _identityHash) constant returns(address)
func (_Registry *RegistrySession) GetChannelAddress(_identityHash common.Address) (common.Address, error) {
	return _Registry.Contract.GetChannelAddress(&_Registry.CallOpts, _identityHash)
}

// GetChannelAddress is a free data retrieval call binding the contract method 0xe5e89412.
//
// Solidity: function getChannelAddress(address _identityHash) constant returns(address)
func (_Registry *RegistryCallerSession) GetChannelAddress(_identityHash common.Address) (common.Address, error) {
	return _Registry.Contract.GetChannelAddress(&_Registry.CallOpts, _identityHash)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() constant returns(address)
func (_Registry *RegistryCaller) GetFundsDestination(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getFundsDestination")
	return *ret0, err
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() constant returns(address)
func (_Registry *RegistrySession) GetFundsDestination() (common.Address, error) {
	return _Registry.Contract.GetFundsDestination(&_Registry.CallOpts)
}

// GetFundsDestination is a free data retrieval call binding the contract method 0xf58c5b6e.
//
// Solidity: function getFundsDestination() constant returns(address)
func (_Registry *RegistryCallerSession) GetFundsDestination() (common.Address, error) {
	return _Registry.Contract.GetFundsDestination(&_Registry.CallOpts)
}

// GetProxyCode is a free data retrieval call binding the contract method 0xab867213.
//
// Solidity: function getProxyCode(address _implementation) constant returns(bytes)
func (_Registry *RegistryCaller) GetProxyCode(opts *bind.CallOpts, _implementation common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getProxyCode", _implementation)
	return *ret0, err
}

// GetProxyCode is a free data retrieval call binding the contract method 0xab867213.
//
// Solidity: function getProxyCode(address _implementation) constant returns(bytes)
func (_Registry *RegistrySession) GetProxyCode(_implementation common.Address) ([]byte, error) {
	return _Registry.Contract.GetProxyCode(&_Registry.CallOpts, _implementation)
}

// GetProxyCode is a free data retrieval call binding the contract method 0xab867213.
//
// Solidity: function getProxyCode(address _implementation) constant returns(bytes)
func (_Registry *RegistryCallerSession) GetProxyCode(_implementation common.Address) ([]byte, error) {
	return _Registry.Contract.GetProxyCode(&_Registry.CallOpts, _implementation)
}

// IsAccountant is a free data retrieval call binding the contract method 0x2a33ddbd.
//
// Solidity: function isAccountant(address _accountantId) constant returns(bool)
func (_Registry *RegistryCaller) IsAccountant(opts *bind.CallOpts, _accountantId common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isAccountant", _accountantId)
	return *ret0, err
}

// IsAccountant is a free data retrieval call binding the contract method 0x2a33ddbd.
//
// Solidity: function isAccountant(address _accountantId) constant returns(bool)
func (_Registry *RegistrySession) IsAccountant(_accountantId common.Address) (bool, error) {
	return _Registry.Contract.IsAccountant(&_Registry.CallOpts, _accountantId)
}

// IsAccountant is a free data retrieval call binding the contract method 0x2a33ddbd.
//
// Solidity: function isAccountant(address _accountantId) constant returns(bool)
func (_Registry *RegistryCallerSession) IsAccountant(_accountantId common.Address) (bool, error) {
	return _Registry.Contract.IsAccountant(&_Registry.CallOpts, _accountantId)
}

// IsActiveAccountant is a free data retrieval call binding the contract method 0x1a3d9a59.
//
// Solidity: function isActiveAccountant(address _accountantId) constant returns(bool)
func (_Registry *RegistryCaller) IsActiveAccountant(opts *bind.CallOpts, _accountantId common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isActiveAccountant", _accountantId)
	return *ret0, err
}

// IsActiveAccountant is a free data retrieval call binding the contract method 0x1a3d9a59.
//
// Solidity: function isActiveAccountant(address _accountantId) constant returns(bool)
func (_Registry *RegistrySession) IsActiveAccountant(_accountantId common.Address) (bool, error) {
	return _Registry.Contract.IsActiveAccountant(&_Registry.CallOpts, _accountantId)
}

// IsActiveAccountant is a free data retrieval call binding the contract method 0x1a3d9a59.
//
// Solidity: function isActiveAccountant(address _accountantId) constant returns(bool)
func (_Registry *RegistryCallerSession) IsActiveAccountant(_accountantId common.Address) (bool, error) {
	return _Registry.Contract.IsActiveAccountant(&_Registry.CallOpts, _accountantId)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Registry *RegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Registry *RegistrySession) IsOwner() (bool, error) {
	return _Registry.Contract.IsOwner(&_Registry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Registry *RegistryCallerSession) IsOwner() (bool, error) {
	return _Registry.Contract.IsOwner(&_Registry.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identityHash) constant returns(bool)
func (_Registry *RegistryCaller) IsRegistered(opts *bind.CallOpts, _identityHash common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isRegistered", _identityHash)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identityHash) constant returns(bool)
func (_Registry *RegistrySession) IsRegistered(_identityHash common.Address) (bool, error) {
	return _Registry.Contract.IsRegistered(&_Registry.CallOpts, _identityHash)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address _identityHash) constant returns(bool)
func (_Registry *RegistryCallerSession) IsRegistered(_identityHash common.Address) (bool, error) {
	return _Registry.Contract.IsRegistered(&_Registry.CallOpts, _identityHash)
}

// MinimalAccountantStake is a free data retrieval call binding the contract method 0x824b09d6.
//
// Solidity: function minimalAccountantStake() constant returns(uint256)
func (_Registry *RegistryCaller) MinimalAccountantStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "minimalAccountantStake")
	return *ret0, err
}

// MinimalAccountantStake is a free data retrieval call binding the contract method 0x824b09d6.
//
// Solidity: function minimalAccountantStake() constant returns(uint256)
func (_Registry *RegistrySession) MinimalAccountantStake() (*big.Int, error) {
	return _Registry.Contract.MinimalAccountantStake(&_Registry.CallOpts)
}

// MinimalAccountantStake is a free data retrieval call binding the contract method 0x824b09d6.
//
// Solidity: function minimalAccountantStake() constant returns(uint256)
func (_Registry *RegistryCallerSession) MinimalAccountantStake() (*big.Int, error) {
	return _Registry.Contract.MinimalAccountantStake(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_Registry *RegistryCaller) RegistrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "registrationFee")
	return *ret0, err
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_Registry *RegistrySession) RegistrationFee() (*big.Int, error) {
	return _Registry.Contract.RegistrationFee(&_Registry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() constant returns(uint256)
func (_Registry *RegistryCallerSession) RegistrationFee() (*big.Int, error) {
	return _Registry.Contract.RegistrationFee(&_Registry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Registry *RegistryCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Registry *RegistrySession) Token() (common.Address, error) {
	return _Registry.Contract.Token(&_Registry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Registry *RegistryCallerSession) Token() (common.Address, error) {
	return _Registry.Contract.Token(&_Registry.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Registry *RegistryCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "totalStaked")
	return *ret0, err
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Registry *RegistrySession) TotalStaked() (*big.Int, error) {
	return _Registry.Contract.TotalStaked(&_Registry.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Registry *RegistryCallerSession) TotalStaked() (*big.Int, error) {
	return _Registry.Contract.TotalStaked(&_Registry.CallOpts)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(uint256 _newFee) returns()
func (_Registry *RegistryTransactor) ChangeRegistrationFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "changeRegistrationFee", _newFee)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(uint256 _newFee) returns()
func (_Registry *RegistrySession) ChangeRegistrationFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.ChangeRegistrationFee(&_Registry.TransactOpts, _newFee)
}

// ChangeRegistrationFee is a paid mutator transaction binding the contract method 0x50050769.
//
// Solidity: function changeRegistrationFee(uint256 _newFee) returns()
func (_Registry *RegistryTransactorSession) ChangeRegistrationFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.ChangeRegistrationFee(&_Registry.TransactOpts, _newFee)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_Registry *RegistryTransactor) ClaimEthers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "claimEthers")
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_Registry *RegistrySession) ClaimEthers() (*types.Transaction, error) {
	return _Registry.Contract.ClaimEthers(&_Registry.TransactOpts)
}

// ClaimEthers is a paid mutator transaction binding the contract method 0x6931b550.
//
// Solidity: function claimEthers() returns()
func (_Registry *RegistryTransactorSession) ClaimEthers() (*types.Transaction, error) {
	return _Registry.Contract.ClaimEthers(&_Registry.TransactOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_Registry *RegistryTransactor) ClaimTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "claimTokens", _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_Registry *RegistrySession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ClaimTokens(&_Registry.TransactOpts, _token)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xdf8de3e7.
//
// Solidity: function claimTokens(address _token) returns()
func (_Registry *RegistryTransactorSession) ClaimTokens(_token common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ClaimTokens(&_Registry.TransactOpts, _token)
}

// RegisterAccountant is a paid mutator transaction binding the contract method 0xf4c1a1f6.
//
// Solidity: function registerAccountant(address _accountantOperator, uint256 _stakeAmount) returns()
func (_Registry *RegistryTransactor) RegisterAccountant(opts *bind.TransactOpts, _accountantOperator common.Address, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerAccountant", _accountantOperator, _stakeAmount)
}

// RegisterAccountant is a paid mutator transaction binding the contract method 0xf4c1a1f6.
//
// Solidity: function registerAccountant(address _accountantOperator, uint256 _stakeAmount) returns()
func (_Registry *RegistrySession) RegisterAccountant(_accountantOperator common.Address, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RegisterAccountant(&_Registry.TransactOpts, _accountantOperator, _stakeAmount)
}

// RegisterAccountant is a paid mutator transaction binding the contract method 0xf4c1a1f6.
//
// Solidity: function registerAccountant(address _accountantOperator, uint256 _stakeAmount) returns()
func (_Registry *RegistryTransactorSession) RegisterAccountant(_accountantOperator common.Address, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RegisterAccountant(&_Registry.TransactOpts, _accountantOperator, _stakeAmount)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x19233600.
//
// Solidity: function registerIdentity(address _identityHash, address _accountantId, uint256 _loanAmount, uint256 _fee, address _beneficiary, bytes _signature) returns()
func (_Registry *RegistryTransactor) RegisterIdentity(opts *bind.TransactOpts, _identityHash common.Address, _accountantId common.Address, _loanAmount *big.Int, _fee *big.Int, _beneficiary common.Address, _signature []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerIdentity", _identityHash, _accountantId, _loanAmount, _fee, _beneficiary, _signature)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x19233600.
//
// Solidity: function registerIdentity(address _identityHash, address _accountantId, uint256 _loanAmount, uint256 _fee, address _beneficiary, bytes _signature) returns()
func (_Registry *RegistrySession) RegisterIdentity(_identityHash common.Address, _accountantId common.Address, _loanAmount *big.Int, _fee *big.Int, _beneficiary common.Address, _signature []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterIdentity(&_Registry.TransactOpts, _identityHash, _accountantId, _loanAmount, _fee, _beneficiary, _signature)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x19233600.
//
// Solidity: function registerIdentity(address _identityHash, address _accountantId, uint256 _loanAmount, uint256 _fee, address _beneficiary, bytes _signature) returns()
func (_Registry *RegistryTransactorSession) RegisterIdentity(_identityHash common.Address, _accountantId common.Address, _loanAmount *big.Int, _fee *big.Int, _beneficiary common.Address, _signature []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterIdentity(&_Registry.TransactOpts, _identityHash, _accountantId, _loanAmount, _fee, _beneficiary, _signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_Registry *RegistryTransactor) SetFundsDestination(opts *bind.TransactOpts, _newDestination common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setFundsDestination", _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_Registry *RegistrySession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetFundsDestination(&_Registry.TransactOpts, _newDestination)
}

// SetFundsDestination is a paid mutator transaction binding the contract method 0x238e130a.
//
// Solidity: function setFundsDestination(address _newDestination) returns()
func (_Registry *RegistryTransactorSession) SetFundsDestination(_newDestination common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetFundsDestination(&_Registry.TransactOpts, _newDestination)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address _beneficiary) returns()
func (_Registry *RegistryTransactor) TransferCollectedFeeTo(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferCollectedFeeTo", _beneficiary)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address _beneficiary) returns()
func (_Registry *RegistrySession) TransferCollectedFeeTo(_beneficiary common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferCollectedFeeTo(&_Registry.TransactOpts, _beneficiary)
}

// TransferCollectedFeeTo is a paid mutator transaction binding the contract method 0xe3252537.
//
// Solidity: function transferCollectedFeeTo(address _beneficiary) returns()
func (_Registry *RegistryTransactorSession) TransferCollectedFeeTo(_beneficiary common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferCollectedFeeTo(&_Registry.TransactOpts, _beneficiary)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// RegistryDestinationChangedIterator is returned from FilterDestinationChanged and is used to iterate over the raw logs and unpacked data for DestinationChanged events raised by the Registry contract.
type RegistryDestinationChangedIterator struct {
	Event *RegistryDestinationChanged // Event containing the contract specifics and raw log

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
func (it *RegistryDestinationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryDestinationChanged)
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
		it.Event = new(RegistryDestinationChanged)
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
func (it *RegistryDestinationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryDestinationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryDestinationChanged represents a DestinationChanged event raised by the Registry contract.
type RegistryDestinationChanged struct {
	PreviousDestination common.Address
	NewDestination      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDestinationChanged is a free log retrieval operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_Registry *RegistryFilterer) FilterDestinationChanged(opts *bind.FilterOpts, previousDestination []common.Address, newDestination []common.Address) (*RegistryDestinationChangedIterator, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return &RegistryDestinationChangedIterator{contract: _Registry.contract, event: "DestinationChanged", logs: logs, sub: sub}, nil
}

// WatchDestinationChanged is a free log subscription operation binding the contract event 0xe1a66d77649cf0a57b9937073549f30f1c82bb865aaf066d2f299e37a62c6aad.
//
// Solidity: event DestinationChanged(address indexed previousDestination, address indexed newDestination)
func (_Registry *RegistryFilterer) WatchDestinationChanged(opts *bind.WatchOpts, sink chan<- *RegistryDestinationChanged, previousDestination []common.Address, newDestination []common.Address) (event.Subscription, error) {

	var previousDestinationRule []interface{}
	for _, previousDestinationItem := range previousDestination {
		previousDestinationRule = append(previousDestinationRule, previousDestinationItem)
	}
	var newDestinationRule []interface{}
	for _, newDestinationItem := range newDestination {
		newDestinationRule = append(newDestinationRule, newDestinationItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "DestinationChanged", previousDestinationRule, newDestinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryDestinationChanged)
				if err := _Registry.contract.UnpackLog(event, "DestinationChanged", log); err != nil {
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

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RegistryRegisteredAccountantIterator is returned from FilterRegisteredAccountant and is used to iterate over the raw logs and unpacked data for RegisteredAccountant events raised by the Registry contract.
type RegistryRegisteredAccountantIterator struct {
	Event *RegistryRegisteredAccountant // Event containing the contract specifics and raw log

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
func (it *RegistryRegisteredAccountantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRegisteredAccountant)
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
		it.Event = new(RegistryRegisteredAccountant)
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
func (it *RegistryRegisteredAccountantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRegisteredAccountantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRegisteredAccountant represents a RegisteredAccountant event raised by the Registry contract.
type RegistryRegisteredAccountant struct {
	AccountantId       common.Address
	AccountantOperator common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRegisteredAccountant is a free log retrieval operation binding the contract event 0xc9f77826eb4f5ea1c94b1b3ff214796f59d1c5f610af8a3ff52cd83e3eebf64e.
//
// Solidity: event RegisteredAccountant(address accountantId, address accountantOperator)
func (_Registry *RegistryFilterer) FilterRegisteredAccountant(opts *bind.FilterOpts) (*RegistryRegisteredAccountantIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RegisteredAccountant")
	if err != nil {
		return nil, err
	}
	return &RegistryRegisteredAccountantIterator{contract: _Registry.contract, event: "RegisteredAccountant", logs: logs, sub: sub}, nil
}

// WatchRegisteredAccountant is a free log subscription operation binding the contract event 0xc9f77826eb4f5ea1c94b1b3ff214796f59d1c5f610af8a3ff52cd83e3eebf64e.
//
// Solidity: event RegisteredAccountant(address accountantId, address accountantOperator)
func (_Registry *RegistryFilterer) WatchRegisteredAccountant(opts *bind.WatchOpts, sink chan<- *RegistryRegisteredAccountant) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RegisteredAccountant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRegisteredAccountant)
				if err := _Registry.contract.UnpackLog(event, "RegisteredAccountant", log); err != nil {
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

// RegistryRegisteredIdentityIterator is returned from FilterRegisteredIdentity and is used to iterate over the raw logs and unpacked data for RegisteredIdentity events raised by the Registry contract.
type RegistryRegisteredIdentityIterator struct {
	Event *RegistryRegisteredIdentity // Event containing the contract specifics and raw log

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
func (it *RegistryRegisteredIdentityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRegisteredIdentity)
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
		it.Event = new(RegistryRegisteredIdentity)
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
func (it *RegistryRegisteredIdentityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRegisteredIdentityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRegisteredIdentity represents a RegisteredIdentity event raised by the Registry contract.
type RegistryRegisteredIdentity struct {
	IdentityHash common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegisteredIdentity is a free log retrieval operation binding the contract event 0x16826e74d06e02bdda286d1820cf7f113495bfa8c8576c331511a3708902dfcc.
//
// Solidity: event RegisteredIdentity(address indexed identityHash)
func (_Registry *RegistryFilterer) FilterRegisteredIdentity(opts *bind.FilterOpts, identityHash []common.Address) (*RegistryRegisteredIdentityIterator, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RegisteredIdentity", identityHashRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRegisteredIdentityIterator{contract: _Registry.contract, event: "RegisteredIdentity", logs: logs, sub: sub}, nil
}

// WatchRegisteredIdentity is a free log subscription operation binding the contract event 0x16826e74d06e02bdda286d1820cf7f113495bfa8c8576c331511a3708902dfcc.
//
// Solidity: event RegisteredIdentity(address indexed identityHash)
func (_Registry *RegistryFilterer) WatchRegisteredIdentity(opts *bind.WatchOpts, sink chan<- *RegistryRegisteredIdentity, identityHash []common.Address) (event.Subscription, error) {

	var identityHashRule []interface{}
	for _, identityHashItem := range identityHash {
		identityHashRule = append(identityHashRule, identityHashItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RegisteredIdentity", identityHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRegisteredIdentity)
				if err := _Registry.contract.UnpackLog(event, "RegisteredIdentity", log); err != nil {
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
