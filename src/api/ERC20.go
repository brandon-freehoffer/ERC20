// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferAnyERC20Token\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600381526020017f5453540000000000000000000000000000000000000000000000000000000000815250600390805190602001906200005f929190620001ba565b506040518060400160405280600a81526020017f5445535420546f6b656e0000000000000000000000000000000000000000000081525060049080519060200190620000ad929190620001ba565b506012600560006101000a81548160ff021916908360ff1602179055506a52b7d2dcc80cd2e400000060068190555060065460076000735a86f0cafd4ef3ba4f0344c138afcc84bd1ed22273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550735a86f0cafd4ef3ba4f0344c138afcc84bd1ed22273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600654604051620001ac919062000285565b60405180910390a362000306565b828054620001c890620002d1565b90600052602060002090601f016020900481019282620001ec576000855562000238565b82601f106200020757805160ff191683800117855562000238565b8280016001018555821562000238579182015b82811115620002375782518255916020019190600101906200021a565b5b5090506200024791906200024b565b5090565b5b80821115620002665760008160009055506001016200024c565b5090565b6000819050919050565b6200027f816200026a565b82525050565b60006020820190506200029c600083018462000274565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002ea57607f821691505b6020821081036200030057620002ff620002a2565b5b50919050565b611add80620003166000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806379ba5097116100a2578063cae9ca5111610071578063cae9ca51146102aa578063d4ee1d90146102da578063dc39d06d146102f8578063dd62ed3e14610328578063f2fde38b146103585761010b565b806379ba5097146102345780638da5cb5b1461023e57806395d89b411461025c578063a9059cbb1461027a5761010b565b8063313ce567116100de578063313ce567146101ac5780633eaaf86b146101ca57806340c10f19146101e857806370a08231146102045761010b565b806306fdde0314610110578063095ea7b31461012e57806318160ddd1461015e57806323b872dd1461017c575b600080fd5b610118610374565b60405161012591906112e7565b60405180910390f35b610148600480360381019061014391906113b1565b610402565b604051610155919061140c565b60405180910390f35b6101666104f4565b6040516101739190611436565b60405180910390f35b61019660048036038101906101919190611451565b610548565b6040516101a3919061140c565b60405180910390f35b6101b46107f3565b6040516101c191906114c0565b60405180910390f35b6101d2610806565b6040516101df9190611436565b60405180910390f35b61020260048036038101906101fd91906113b1565b61080c565b005b61021e600480360381019061021991906114db565b610872565b60405161022b9190611436565b60405180910390f35b61023c6108bb565b005b610246610a57565b6040516102539190611517565b60405180910390f35b610264610a7b565b60405161027191906112e7565b60405180910390f35b610294600480360381019061028f91906113b1565b610b09565b6040516102a1919061140c565b60405180910390f35b6102c460048036038101906102bf9190611667565b610ca4565b6040516102d1919061140c565b60405180910390f35b6102e2610e08565b6040516102ef9190611517565b60405180910390f35b610312600480360381019061030d91906113b1565b610e2e565b60405161031f919061140c565b60405180910390f35b610342600480360381019061033d91906116d6565b610f2f565b60405161034f9190611436565b60405180910390f35b610372600480360381019061036d91906114db565b610fb6565b005b6004805461038190611745565b80601f01602080910402602001604051908101604052809291908181526020018280546103ad90611745565b80156103fa5780601f106103cf576101008083540402835291602001916103fa565b820191906000526020600020905b8154815290600101906020018083116103dd57829003601f168201915b505050505081565b600081600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104e29190611436565b60405180910390a36001905092915050565b6000600760008073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460065461054391906117a5565b905090565b600061059c82600760008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461105290919063ffffffff16565b600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061066e82600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461105290919063ffffffff16565b600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061074082600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461107990919063ffffffff16565b600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516107e09190611436565b60405180910390a3600190509392505050565b600560009054906101000a900460ff1681565b60065481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461086457600080fd5b61086e82826110c1565b5050565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461091557600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60038054610a8890611745565b80601f0160208091040260200160405190810160405280929190818152602001828054610ab490611745565b8015610b015780601f10610ad657610100808354040283529160200191610b01565b820191906000526020600020905b815481529060010190602001808311610ae457829003601f168201915b505050505081565b6000610b5d82600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461105290919063ffffffff16565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610bf282600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461121890919063ffffffff16565b600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610c929190611436565b60405180910390a36001905092915050565b600082600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051610d849190611436565b60405180910390a38373ffffffffffffffffffffffffffffffffffffffff16638f4ffcb1338587866040518563ffffffff1660e01b8152600401610dcb949392919061182e565b600060405180830381600087803b158015610de557600080fd5b505af1158015610df9573d6000803e3d6000fd5b50505050600190509392505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e8957600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b8152600401610ee492919061187a565b6020604051808303816000875af1158015610f03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2791906118cf565b905092915050565b6000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461100e57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600082821115611065576110646118fc565b5b818361107191906117a5565b905092915050565b600080830361108b57600090506110bb565b60008284611099919061192b565b90508284826110a891906119b4565b146110b6576110b56118fc565b5b809150505b92915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611130576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161112790611a31565b60405180910390fd5b61113c60008383611244565b806006600082825461114e9190611a51565b9250508190555080600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516112009190611436565b60405180910390a361121460008383611249565b5050565b60008082846112279190611a51565b90508381101561123a576112396118fc565b5b8091505092915050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561128857808201518184015260208101905061126d565b83811115611297576000848401525b50505050565b6000601f19601f8301169050919050565b60006112b98261124e565b6112c38185611259565b93506112d381856020860161126a565b6112dc8161129d565b840191505092915050565b6000602082019050818103600083015261130181846112ae565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006113488261131d565b9050919050565b6113588161133d565b811461136357600080fd5b50565b6000813590506113758161134f565b92915050565b6000819050919050565b61138e8161137b565b811461139957600080fd5b50565b6000813590506113ab81611385565b92915050565b600080604083850312156113c8576113c7611313565b5b60006113d685828601611366565b92505060206113e78582860161139c565b9150509250929050565b60008115159050919050565b611406816113f1565b82525050565b600060208201905061142160008301846113fd565b92915050565b6114308161137b565b82525050565b600060208201905061144b6000830184611427565b92915050565b60008060006060848603121561146a57611469611313565b5b600061147886828701611366565b935050602061148986828701611366565b925050604061149a8682870161139c565b9150509250925092565b600060ff82169050919050565b6114ba816114a4565b82525050565b60006020820190506114d560008301846114b1565b92915050565b6000602082840312156114f1576114f0611313565b5b60006114ff84828501611366565b91505092915050565b6115118161133d565b82525050565b600060208201905061152c6000830184611508565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6115748261129d565b810181811067ffffffffffffffff821117156115935761159261153c565b5b80604052505050565b60006115a6611309565b90506115b2828261156b565b919050565b600067ffffffffffffffff8211156115d2576115d161153c565b5b6115db8261129d565b9050602081019050919050565b82818337600083830152505050565b600061160a611605846115b7565b61159c565b90508281526020810184848401111561162657611625611537565b5b6116318482856115e8565b509392505050565b600082601f83011261164e5761164d611532565b5b813561165e8482602086016115f7565b91505092915050565b6000806000606084860312156116805761167f611313565b5b600061168e86828701611366565b935050602061169f8682870161139c565b925050604084013567ffffffffffffffff8111156116c0576116bf611318565b5b6116cc86828701611639565b9150509250925092565b600080604083850312156116ed576116ec611313565b5b60006116fb85828601611366565b925050602061170c85828601611366565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061175d57607f821691505b6020821081036117705761176f611716565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006117b08261137b565b91506117bb8361137b565b9250828210156117ce576117cd611776565b5b828203905092915050565b600081519050919050565b600082825260208201905092915050565b6000611800826117d9565b61180a81856117e4565b935061181a81856020860161126a565b6118238161129d565b840191505092915050565b60006080820190506118436000830187611508565b6118506020830186611427565b61185d6040830185611508565b818103606083015261186f81846117f5565b905095945050505050565b600060408201905061188f6000830185611508565b61189c6020830184611427565b9392505050565b6118ac816113f1565b81146118b757600080fd5b50565b6000815190506118c9816118a3565b92915050565b6000602082840312156118e5576118e4611313565b5b60006118f3848285016118ba565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60006119368261137b565b91506119418361137b565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561197a57611979611776565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006119bf8261137b565b91506119ca8361137b565b9250826119da576119d9611985565b5b828204905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611a1b601f83611259565b9150611a26826119e5565b602082019050919050565b60006020820190508181036000830152611a4a81611a0e565b9050919050565b6000611a5c8261137b565b9150611a678361137b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611a9c57611a9b611776565b5b82820190509291505056fea2646970667358221220cb7d09e7de9b410599ec2e25df4effc6e71632a776980c6154697e6146c48ea564736f6c634300080e0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// TotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() view returns(uint256)
func (_Api *ApiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "_totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() view returns(uint256)
func (_Api *ApiSession) TotalSupply() (*big.Int, error) {
	return _Api.Contract.TotalSupply(&_Api.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() view returns(uint256)
func (_Api *ApiCallerSession) TotalSupply() (*big.Int, error) {
	return _Api.Contract.TotalSupply(&_Api.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_Api *ApiCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "allowance", tokenOwner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_Api *ApiSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Api.Contract.Allowance(&_Api.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) view returns(uint256 remaining)
func (_Api *ApiCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Api.Contract.Allowance(&_Api.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_Api *ApiCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_Api *ApiSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Api.Contract.BalanceOf(&_Api.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256 balance)
func (_Api *ApiCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Api.Contract.BalanceOf(&_Api.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Api *ApiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Api *ApiSession) Decimals() (uint8, error) {
	return _Api.Contract.Decimals(&_Api.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Api *ApiCallerSession) Decimals() (uint8, error) {
	return _Api.Contract.Decimals(&_Api.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiSession) Name() (string, error) {
	return _Api.Contract.Name(&_Api.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Api *ApiCallerSession) Name() (string, error) {
	return _Api.Contract.Name(&_Api.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Api *ApiCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "newOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Api *ApiSession) NewOwner() (common.Address, error) {
	return _Api.Contract.NewOwner(&_Api.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Api *ApiCallerSession) NewOwner() (common.Address, error) {
	return _Api.Contract.NewOwner(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCallerSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiSession) Symbol() (string, error) {
	return _Api.Contract.Symbol(&_Api.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Api *ApiCallerSession) Symbol() (string, error) {
	return _Api.Contract.Symbol(&_Api.CallOpts)
}

// TotalSupply1 is a free data retrieval call binding the contract method 0x976d36d0.
//
// Solidity: function totalSupply1() view returns(uint256)
func (_Api *ApiCaller) TotalSupply1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "totalSupply1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply1 is a free data retrieval call binding the contract method 0x976d36d0.
//
// Solidity: function totalSupply1() view returns(uint256)
func (_Api *ApiSession) TotalSupply1() (*big.Int, error) {
	return _Api.Contract.TotalSupply1(&_Api.CallOpts)
}

// TotalSupply1 is a free data retrieval call binding the contract method 0x976d36d0.
//
// Solidity: function totalSupply1() view returns(uint256)
func (_Api *ApiCallerSession) TotalSupply1() (*big.Int, error) {
	return _Api.Contract.TotalSupply1(&_Api.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Api *ApiTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Api *ApiSession) AcceptOwnership() (*types.Transaction, error) {
	return _Api.Contract.AcceptOwnership(&_Api.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Api *ApiTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Api.Contract.AcceptOwnership(&_Api.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Api *ApiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Api *ApiSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Approve(&_Api.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Api *ApiTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Approve(&_Api.TransactOpts, spender, tokens)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 tokens, bytes data) returns(bool success)
func (_Api *ApiTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "approveAndCall", spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 tokens, bytes data) returns(bool success)
func (_Api *ApiSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.ApproveAndCall(&_Api.TransactOpts, spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 tokens, bytes data) returns(bool success)
func (_Api *ApiTransactorSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Api.Contract.ApproveAndCall(&_Api.TransactOpts, spender, tokens, data)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Api *ApiTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Api *ApiSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Mint(&_Api.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Api *ApiTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Mint(&_Api.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Api *ApiTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Api *ApiSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Transfer(&_Api.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Api *ApiTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Transfer(&_Api.TransactOpts, to, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, uint256 tokens) returns(bool success)
func (_Api *ApiTransactor) TransferAnyERC20Token(opts *bind.TransactOpts, tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferAnyERC20Token", tokenAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, uint256 tokens) returns(bool success)
func (_Api *ApiSession) TransferAnyERC20Token(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferAnyERC20Token(&_Api.TransactOpts, tokenAddress, tokens)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, uint256 tokens) returns(bool success)
func (_Api *ApiTransactorSession) TransferAnyERC20Token(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferAnyERC20Token(&_Api.TransactOpts, tokenAddress, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Api *ApiTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Api *ApiSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferFrom(&_Api.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Api *ApiTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Api.Contract.TransferFrom(&_Api.TransactOpts, from, to, tokens)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Api *ApiTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Api *ApiSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Api.Contract.TransferOwnership(&_Api.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Api *ApiTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Api.Contract.TransferOwnership(&_Api.TransactOpts, _newOwner)
}

// ApiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Api contract.
type ApiApprovalIterator struct {
	Event *ApiApproval // Event containing the contract specifics and raw log

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
func (it *ApiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiApproval)
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
		it.Event = new(ApiApproval)
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
func (it *ApiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiApproval represents a Approval event raised by the Api contract.
type ApiApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Api *ApiFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*ApiApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ApiApprovalIterator{contract: _Api.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Api *ApiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ApiApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiApproval)
				if err := _Api.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Api *ApiFilterer) ParseApproval(log types.Log) (*ApiApproval, error) {
	event := new(ApiApproval)
	if err := _Api.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Api contract.
type ApiOwnershipTransferredIterator struct {
	Event *ApiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ApiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiOwnershipTransferred)
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
		it.Event = new(ApiOwnershipTransferred)
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
func (it *ApiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiOwnershipTransferred represents a OwnershipTransferred event raised by the Api contract.
type ApiOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _from, address indexed _to)
func (_Api *ApiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ApiOwnershipTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ApiOwnershipTransferredIterator{contract: _Api.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _from, address indexed _to)
func (_Api *ApiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ApiOwnershipTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiOwnershipTransferred)
				if err := _Api.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _from, address indexed _to)
func (_Api *ApiFilterer) ParseOwnershipTransferred(log types.Log) (*ApiOwnershipTransferred, error) {
	event := new(ApiOwnershipTransferred)
	if err := _Api.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Api contract.
type ApiTransferIterator struct {
	Event *ApiTransfer // Event containing the contract specifics and raw log

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
func (it *ApiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiTransfer)
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
		it.Event = new(ApiTransfer)
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
func (it *ApiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiTransfer represents a Transfer event raised by the Api contract.
type ApiTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Api *ApiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ApiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ApiTransferIterator{contract: _Api.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Api *ApiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ApiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiTransfer)
				if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Api *ApiFilterer) ParseTransfer(log types.Log) (*ApiTransfer, error) {
	event := new(ApiTransfer)
	if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
