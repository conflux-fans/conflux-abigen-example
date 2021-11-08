// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/Conflux-Chain/conflux-abigen/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient/bulk"

	types "github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = ethBind.Bind
	_ = common.Big1
	_ = ethtypes.BloomLookup
	_ = event.NewSubscription
)

// MyTokenStudent is an auto generated low-level Go binding around an user-defined struct.
type MyTokenStudent struct {
	Name string
	Age  *big.Int
}

// MyTokenABI is the input ABI used to generate the binding from.
const MyTokenABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimalUnits\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"returnTuple\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"returnTupleWithStruct\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"Name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"Age\",\"type\":\"uint256\"}],\"internalType\":\"structMyToken.Student\",\"name\":\"s\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MyTokenFuncSigs maps the 4-byte function signature to its string representation.
var MyTokenFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"39009482": "returnTuple()",
	"e354439b": "returnTupleWithStruct()",
	"95d89b41": "symbol()",
	"a9059cbb": "transfer(address,uint256)",
}

// MyTokenBin is the compiled bytecode used for deploying new contracts.
var MyTokenBin = "0x608060405234801561001057600080fd5b506040516107be3803806107be83398101604081905261002f916101c9565b336000908152600360209081526040822086905584516100529291860190610084565b508051610066906001906020840190610084565b50506002805460ff191660ff92909216919091179055506102a19050565b82805461009090610250565b90600052602060002090601f0160209004810192826100b257600085556100f8565b82601f106100cb57805160ff19168380011785556100f8565b828001600101855582156100f8579182015b828111156100f85782518255916020019190600101906100dd565b50610104929150610108565b5090565b5b808211156101045760008155600101610109565b600082601f83011261012e57600080fd5b81516001600160401b03808211156101485761014861028b565b604051601f8301601f19908116603f011681019082821181831017156101705761017061028b565b8160405283815260209250868385880101111561018c57600080fd5b600091505b838210156101ae5785820183015181830184015290820190610191565b838211156101bf5760008385830101525b9695505050505050565b600080600080608085870312156101df57600080fd5b845160208601519094506001600160401b03808211156101fe57600080fd5b61020a8883890161011d565b94506040870151915060ff8216821461022257600080fd5b60608701519193508082111561023757600080fd5b506102448782880161011d565b91505092959194509250565b600181811c9082168061026457607f821691505b6020821081141561028557634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b61050e806102b06000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806370a082311161005b57806370a08231146100d357806395d89b4114610101578063a9059cbb14610109578063e354439b1461011e57600080fd5b806306fdde0314610082578063313ce567146100a057806339009482146100bf575b600080fd5b61008a610176565b604051610097919061040e565b60405180910390f35b6002546100ad9060ff1681565b60405160ff9091168152602001610097565b604080516001808252602082015201610097565b6100f36100e1366004610375565b60036020526000908152604090205481565b604051908152602001610097565b61008a610204565b61011c610117366004610397565b610211565b005b6101686040805180820190915260608152600060208201525060408051608081018252600691810191825265736f7068696160d01b60608201529081526014602082015290600190565b604051610097929190610421565b6000805461018390610487565b80601f01602080910402602001604051908101604052809291908181526020018280546101af90610487565b80156101fc5780601f106101d1576101008083540402835291602001916101fc565b820191906000526020600020905b8154815290600101906020018083116101df57829003601f168201915b505050505081565b6001805461018390610487565b3360009081526003602052604090205481106102695760405162461bcd60e51b81526020600482015260126024820152710c4c2d8c2dcc6ca40dcdee840cadcdeeaced60731b60448201526064015b60405180910390fd5b6001600160a01b03821660009081526003602052604090205461028c8282610458565b116102c45760405162461bcd60e51b81526020600482015260086024820152676f766572666c6f7760c01b6044820152606401610260565b33600090815260036020526040812080548392906102e3908490610470565b90915550506001600160a01b03821660009081526003602052604081208054839290610310908490610458565b90915550506040518181526001600160a01b0383169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b80356001600160a01b038116811461037057600080fd5b919050565b60006020828403121561038757600080fd5b61039082610359565b9392505050565b600080604083850312156103aa57600080fd5b6103b383610359565b946020939093013593505050565b6000815180845260005b818110156103e7576020818501810151868301820152016103cb565b818111156103f9576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061039060208301846103c1565b604081526000835160408084015261043c60808401826103c1565b6020958601516060850152931515929094019190915250919050565b6000821982111561046b5761046b6104c2565b500190565b600082821015610482576104826104c2565b500390565b600181811c9082168061049b57607f821691505b602082108114156104bc57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea264697066735822122000b371da5e090505eac09c59785844e40dd1c8985bf2b2ac985b958c996caf4b64736f6c63430008060033"

// DeployMyToken deploys a new Conflux contract, binding an instance of MyToken to it.
func DeployMyToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, decimalUnits uint8, tokenSymbol string) (*types.UnsignedTransaction, *types.Hash, *MyToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MyTokenABI))
	if err != nil {
		return nil, nil, nil, err
	}

	tx, hash, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MyTokenBin), backend, initialSupply, tokenName, decimalUnits, tokenSymbol)
	if err != nil {
		return nil, nil, nil, err
	}
	return tx, hash, &MyToken{MyTokenCaller: MyTokenCaller{contract: contract}, MyTokenTransactor: MyTokenTransactor{contract: contract}, MyTokenFilterer: MyTokenFilterer{contract: contract}}, nil
}

// MyToken is an auto generated Go binding around an Conflux contract.
type MyToken struct {
	MyTokenCaller     // Read-only binding to the contract
	MyTokenTransactor // Write-only binding to the contract
	MyTokenFilterer   // Log filterer for contract events
}

// MyTokenCaller is an auto generated read-only Go binding around an Conflux contract.
type MyTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type MyTokenBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenTransactor is an auto generated write-only Go binding around an Conflux contract.
type MyTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type MyTokenBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type MyTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type MyTokenSession struct {
	Contract     *MyToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyTokenCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type MyTokenCallerSession struct {
	Contract *MyTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MyTokenTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type MyTokenTransactorSession struct {
	Contract     *MyTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MyTokenRaw is an auto generated low-level Go binding around an Conflux contract.
type MyTokenRaw struct {
	Contract *MyToken // Generic contract binding to access the raw methods on
}

// MyTokenCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type MyTokenCallerRaw struct {
	Contract *MyTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MyTokenTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type MyTokenTransactorRaw struct {
	Contract *MyTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyToken creates a new instance of MyToken, bound to a specific deployed contract.
func NewMyToken(address types.Address, backend bind.ContractBackend) (*MyToken, error) {
	contract, err := bindMyToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyToken{MyTokenCaller: MyTokenCaller{contract: contract}, MyTokenTransactor: MyTokenTransactor{contract: contract}, MyTokenFilterer: MyTokenFilterer{contract: contract}}, nil
}

// NewMyTokenCaller creates a new read-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenCaller(address types.Address, caller bind.ContractCaller) (*MyTokenCaller, error) {
	contract, err := bindMyToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenCaller{contract: contract}, nil
}

// NewMyTokenTransactor creates a new write-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenTransactor(address types.Address, transactor bind.ContractTransactor) (*MyTokenTransactor, error) {
	contract, err := bindMyToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenTransactor{contract: contract}, nil
}

// NewMyTokenFilterer creates a new log filterer instance of MyToken, bound to a specific deployed contract.
func NewMyTokenFilterer(address types.Address, filterer bind.ContractFilterer) (*MyTokenFilterer, error) {
	contract, err := bindMyToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyTokenFilterer{contract: contract}, nil
}

// NewMyTokenCaller creates a new read-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenBulkCaller(address types.Address, caller bind.ContractCaller) (*MyTokenBulkCaller, error) {
	contract, err := bindMyToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenBulkCaller{contract: contract}, nil
}

// NewMyTokenBulkTransactor creates a new write-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*MyTokenBulkTransactor, error) {
	contract, err := bindMyToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenBulkTransactor{contract: contract}, nil
}

// bindMyToken binds a generic wrapper to an already deployed contract.
func bindMyToken(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyToken *MyTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyToken.Contract.MyTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyToken *MyTokenRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.MyTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyToken *MyTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.MyTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyToken *MyTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyToken *MyTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyToken *MyTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MyToken *MyTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MyToken *MyTokenBulkCaller) BalanceOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, arg0 common.Address) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	msg := _MyToken.contract.GenRequest(opts, "balanceOf", arg0)

	out0 := new(*big.Int)

	err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "balanceOf")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(msg, opts.EpochNumber, outDecoder, err)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MyToken *MyTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MyToken.Contract.BalanceOf(&_MyToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MyToken *MyTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MyToken.Contract.BalanceOf(&_MyToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyToken *MyTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyToken *MyTokenBulkCaller) Decimals(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*uint8, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	msg := _MyToken.contract.GenRequest(opts, "decimals")

	out0 := new(uint8)

	err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "decimals")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(uint8)).(*uint8)

		return nil
	}

	bulkcaller.Customer().ContractCall(msg, opts.EpochNumber, outDecoder, err)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyToken *MyTokenSession) Decimals() (uint8, error) {
	return _MyToken.Contract.Decimals(&_MyToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyToken *MyTokenCallerSession) Decimals() (uint8, error) {
	return _MyToken.Contract.Decimals(&_MyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyToken *MyTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyToken *MyTokenBulkCaller) Name(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	msg := _MyToken.contract.GenRequest(opts, "name")

	out0 := new(string)

	err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "name")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(msg, opts.EpochNumber, outDecoder, err)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyToken *MyTokenSession) Name() (string, error) {
	return _MyToken.Contract.Name(&_MyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyToken *MyTokenCallerSession) Name() (string, error) {
	return _MyToken.Contract.Name(&_MyToken.CallOpts)
}

// ReturnTuple is a free data retrieval call binding the contract method 0x39009482.
//
// Solidity: function returnTuple() pure returns(uint256 a, bool b)
func (_MyToken *MyTokenCaller) ReturnTuple(opts *bind.CallOpts) (struct {
	A *big.Int
	B bool
}, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "returnTuple")

	outstruct := new(struct {
		A *big.Int
		B bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.A = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.B = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// ReturnTuple is a free data retrieval call binding the contract method 0x39009482.
//
// Solidity: function returnTuple() pure returns(uint256 a, bool b)
func (_MyToken *MyTokenBulkCaller) ReturnTuple(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*struct {
	A *big.Int
	B bool
}, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	msg := _MyToken.contract.GenRequest(opts, "returnTuple")

	outstruct := new(struct {
		A *big.Int
		B bool
	})

	err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "returnTuple")
		if err != nil {
			return err
		}

		outstruct.A = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		outstruct.B = *abi.ConvertType(out[1], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(msg, opts.EpochNumber, outDecoder, err)

	return outstruct, err

}

// ReturnTuple is a free data retrieval call binding the contract method 0x39009482.
//
// Solidity: function returnTuple() pure returns(uint256 a, bool b)
func (_MyToken *MyTokenSession) ReturnTuple() (struct {
	A *big.Int
	B bool
}, error) {
	return _MyToken.Contract.ReturnTuple(&_MyToken.CallOpts)
}

// ReturnTuple is a free data retrieval call binding the contract method 0x39009482.
//
// Solidity: function returnTuple() pure returns(uint256 a, bool b)
func (_MyToken *MyTokenCallerSession) ReturnTuple() (struct {
	A *big.Int
	B bool
}, error) {
	return _MyToken.Contract.ReturnTuple(&_MyToken.CallOpts)
}

// ReturnTupleWithStruct is a free data retrieval call binding the contract method 0xe354439b.
//
// Solidity: function returnTupleWithStruct() pure returns((string,uint256) s, bool ok)
func (_MyToken *MyTokenCaller) ReturnTupleWithStruct(opts *bind.CallOpts) (struct {
	S  MyTokenStudent
	Ok bool
}, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "returnTupleWithStruct")

	outstruct := new(struct {
		S  MyTokenStudent
		Ok bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.S = *abi.ConvertType(out[0], new(MyTokenStudent)).(*MyTokenStudent)
	outstruct.Ok = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// ReturnTupleWithStruct is a free data retrieval call binding the contract method 0xe354439b.
//
// Solidity: function returnTupleWithStruct() pure returns((string,uint256) s, bool ok)
func (_MyToken *MyTokenBulkCaller) ReturnTupleWithStruct(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*struct {
	S  MyTokenStudent
	Ok bool
}, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	msg := _MyToken.contract.GenRequest(opts, "returnTupleWithStruct")

	outstruct := new(struct {
		S  MyTokenStudent
		Ok bool
	})

	err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "returnTupleWithStruct")
		if err != nil {
			return err
		}

		outstruct.S = *abi.ConvertType(out[0], new(MyTokenStudent)).(*MyTokenStudent)
		outstruct.Ok = *abi.ConvertType(out[1], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(msg, opts.EpochNumber, outDecoder, err)

	return outstruct, err

}

// ReturnTupleWithStruct is a free data retrieval call binding the contract method 0xe354439b.
//
// Solidity: function returnTupleWithStruct() pure returns((string,uint256) s, bool ok)
func (_MyToken *MyTokenSession) ReturnTupleWithStruct() (struct {
	S  MyTokenStudent
	Ok bool
}, error) {
	return _MyToken.Contract.ReturnTupleWithStruct(&_MyToken.CallOpts)
}

// ReturnTupleWithStruct is a free data retrieval call binding the contract method 0xe354439b.
//
// Solidity: function returnTupleWithStruct() pure returns((string,uint256) s, bool ok)
func (_MyToken *MyTokenCallerSession) ReturnTupleWithStruct() (struct {
	S  MyTokenStudent
	Ok bool
}, error) {
	return _MyToken.Contract.ReturnTupleWithStruct(&_MyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyToken *MyTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyToken *MyTokenBulkCaller) Symbol(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	msg := _MyToken.contract.GenRequest(opts, "symbol")

	out0 := new(string)

	err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyToken.contract.DecodeOutput(&out, rawOut, "symbol")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(msg, opts.EpochNumber, outDecoder, err)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyToken *MyTokenSession) Symbol() (string, error) {
	return _MyToken.Contract.Symbol(&_MyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyToken *MyTokenCallerSession) Symbol() (string, error) {
	return _MyToken.Contract.Symbol(&_MyToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_MyToken *MyTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_MyToken *MyTokenBulkTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) types.UnsignedTransaction {
	return _MyToken.contract.GenUnsignedTransaction(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_MyToken *MyTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.Transfer(&_MyToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_MyToken *MyTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyToken.Contract.Transfer(&_MyToken.TransactOpts, _to, _value)
}

// MyTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyToken contract.
type MyTokenTransferIterator struct {
	Event *MyTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MyTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyTokenTransfer)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyTokenTransfer)
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
func (it *MyTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyTokenTransfer represents a Transfer event raised by the MyToken contract.
type MyTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// MyTokenTransferOrChainReorg represents a Transfer subscription event raised by the MyToken contract.
type MyTokenTransferOrChainReorg struct {
	Event      *MyTokenTransfer
	ChainReorg *types.ChainReorg
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyToken *MyTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MyTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, err := _MyToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyTokenTransferIterator{contract: _MyToken.contract, event: "Transfer", logs: logs}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyToken *MyTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyTokenTransferOrChainReorg, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}

	// go func() {
	// 	for {
	// 		r, ok := <-_reorgs
	// 		if !ok {
	// 			return
	// 		}
	// 		reorgs <- r
	// 	}
	// }()

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyTokenTransferOrChainReorg)
				event.Event = new(MyTokenTransfer)

				if log.ChainReorg == nil {
					if err := _MyToken.contract.UnpackLog(event.Event, "Transfer", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyToken *MyTokenFilterer) ParseTransfer(log types.Log) (*MyTokenTransfer, error) {
	event := new(MyTokenTransfer)
	if err := _MyToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
