// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/bind"
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

// MyERC20TokenStudent is an auto generated low-level Go binding around an user-defined struct.
type MyERC20TokenStudent struct {
	Name string
	Age  *big.Int
}

// MyERC20TokenABI is the input ABI used to generate the binding from.
const MyERC20TokenABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimalUnits\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"returnTuple\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"returnTupleWithStruct\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"Name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"Age\",\"type\":\"uint256\"}],\"internalType\":\"structMyERC20Token.Student\",\"name\":\"s\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MyERC20TokenFuncSigs maps the 4-byte function signature to its string representation.
var MyERC20TokenFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"39009482": "returnTuple()",
	"e354439b": "returnTupleWithStruct()",
	"95d89b41": "symbol()",
	"a9059cbb": "transfer(address,uint256)",
}

// MyERC20TokenBin is the compiled bytecode used for deploying new contracts.
var MyERC20TokenBin = "0x608060405234801561001057600080fd5b5060405161081238038061081283398101604081905261002f91610131565b33600090815260036020526040812085905561004b8482610241565b5060016100588282610241565b50506002805460ff191660ff92909216919091179055506103009050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261009d57600080fd5b81516001600160401b03808211156100b7576100b7610076565b604051601f8301601f19908116603f011681019082821181831017156100df576100df610076565b816040528381526020925086838588010111156100fb57600080fd5b600091505b8382101561011d5785820183015181830184015290820190610100565b600093810190920192909252949350505050565b6000806000806080858703121561014757600080fd5b845160208601519094506001600160401b038082111561016657600080fd5b6101728883890161008c565b94506040870151915060ff8216821461018a57600080fd5b60608701519193508082111561019f57600080fd5b506101ac8782880161008c565b91505092959194509250565b600181811c908216806101cc57607f821691505b6020821081036101ec57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561023c57600081815260208120601f850160051c810160208610156102195750805b601f850160051c820191505b8181101561023857828155600101610225565b5050505b505050565b81516001600160401b0381111561025a5761025a610076565b61026e8161026884546101b8565b846101f2565b602080601f8311600181146102a3576000841561028b5750858301515b600019600386901b1c1916600185901b178555610238565b600085815260208120601f198616915b828110156102d2578886015182559484019460019091019084016102b3565b50858210156102f05787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6105038061030f6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806370a082311161005b57806370a08231146100d357806395d89b4114610101578063a9059cbb14610109578063e354439b1461011e57600080fd5b806306fdde0314610082578063313ce567146100a057806339009482146100bf575b600080fd5b61008a610176565b604051610097919061039f565b60405180910390f35b6002546100ad9060ff1681565b60405160ff9091168152602001610097565b604080516001808252602082015201610097565b6100f36100e13660046103d5565b60036020526000908152604090205481565b604051908152602001610097565b61008a610204565b61011c6101173660046103f0565b610211565b005b6101686040805180820190915260608152600060208201525060408051608081018252600691810191825265736f7068696160d01b60608201529081526014602082015290600190565b60405161009792919061041a565b6000805461018390610451565b80601f01602080910402602001604051908101604052809291908181526020018280546101af90610451565b80156101fc5780601f106101d1576101008083540402835291602001916101fc565b820191906000526020600020905b8154815290600101906020018083116101df57829003601f168201915b505050505081565b6001805461018390610451565b3360009081526003602052604090205481106102695760405162461bcd60e51b81526020600482015260126024820152710c4c2d8c2dcc6ca40dcdee840cadcdeeaced60731b60448201526064015b60405180910390fd5b6001600160a01b03821660009081526003602052604090205461028c82826104a1565b116102c45760405162461bcd60e51b81526020600482015260086024820152676f766572666c6f7760c01b6044820152606401610260565b33600090815260036020526040812080548392906102e39084906104ba565b90915550506001600160a01b038216600090815260036020526040812080548392906103109084906104a1565b90915550506040518181526001600160a01b0383169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561037f57602081850181015186830182015201610363565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006103b26020830184610359565b9392505050565b80356001600160a01b03811681146103d057600080fd5b919050565b6000602082840312156103e757600080fd5b6103b2826103b9565b6000806040838503121561040357600080fd5b61040c836103b9565b946020939093013593505050565b60408152600083516040808401526104356080840182610359565b6020958601516060850152931515929094019190915250919050565b600181811c9082168061046557607f821691505b60208210810361048557634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b808201808211156104b4576104b461048b565b92915050565b818103818111156104b4576104b461048b56fea2646970667358221220c5b164406fccf70a7f9af8490cfa5336fa3bcce48a3b42339339e6b558cbbcf564736f6c63430008130033"

// DeployMyERC20Token deploys a new Conflux contract, binding an instance of MyERC20Token to it.
func DeployMyERC20Token(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, decimalUnits uint8, tokenSymbol string) (*types.UnsignedTransaction, *types.Hash, *MyERC20Token, error) {
	parsed, err := abi.JSON(strings.NewReader(MyERC20TokenABI))
	if err != nil {
		return nil, nil, nil, err
	}

	tx, hash, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MyERC20TokenBin), backend, initialSupply, tokenName, decimalUnits, tokenSymbol)
	if err != nil {
		return nil, nil, nil, err
	}
	return tx, hash, &MyERC20Token{MyERC20TokenCaller: MyERC20TokenCaller{contract: contract}, MyERC20TokenTransactor: MyERC20TokenTransactor{contract: contract}, MyERC20TokenFilterer: MyERC20TokenFilterer{contract: contract}}, nil
}

// MyERC20Token is an auto generated Go binding around an Conflux contract.
type MyERC20Token struct {
	MyERC20TokenCaller     // Read-only binding to the contract
	MyERC20TokenTransactor // Write-only binding to the contract
	MyERC20TokenFilterer   // Log filterer for contract events
}

// MyERC20TokenCaller is an auto generated read-only Go binding around an Conflux contract.
type MyERC20TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyERC20TokenBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type MyERC20TokenBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyERC20TokenTransactor is an auto generated write-only Go binding around an Conflux contract.
type MyERC20TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyERC20TokenBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type MyERC20TokenBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyERC20TokenFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type MyERC20TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyERC20TokenSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type MyERC20TokenSession struct {
	Contract     *MyERC20Token     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyERC20TokenCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type MyERC20TokenCallerSession struct {
	Contract *MyERC20TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MyERC20TokenTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type MyERC20TokenTransactorSession struct {
	Contract     *MyERC20TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MyERC20TokenRaw is an auto generated low-level Go binding around an Conflux contract.
type MyERC20TokenRaw struct {
	Contract *MyERC20Token // Generic contract binding to access the raw methods on
}

// MyERC20TokenCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type MyERC20TokenCallerRaw struct {
	Contract *MyERC20TokenCaller // Generic read-only contract binding to access the raw methods on
}

// MyERC20TokenTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type MyERC20TokenTransactorRaw struct {
	Contract *MyERC20TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyERC20Token creates a new instance of MyERC20Token, bound to a specific deployed contract.
func NewMyERC20Token(address types.Address, backend bind.ContractBackend) (*MyERC20Token, error) {
	contract, err := bindMyERC20Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyERC20Token{MyERC20TokenCaller: MyERC20TokenCaller{contract: contract}, MyERC20TokenTransactor: MyERC20TokenTransactor{contract: contract}, MyERC20TokenFilterer: MyERC20TokenFilterer{contract: contract}}, nil
}

// NewMyERC20TokenCaller creates a new read-only instance of MyERC20Token, bound to a specific deployed contract.
func NewMyERC20TokenCaller(address types.Address, caller bind.ContractCaller) (*MyERC20TokenCaller, error) {
	contract, err := bindMyERC20Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyERC20TokenCaller{contract: contract}, nil
}

// NewMyERC20TokenTransactor creates a new write-only instance of MyERC20Token, bound to a specific deployed contract.
func NewMyERC20TokenTransactor(address types.Address, transactor bind.ContractTransactor) (*MyERC20TokenTransactor, error) {
	contract, err := bindMyERC20Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyERC20TokenTransactor{contract: contract}, nil
}

// NewMyERC20TokenFilterer creates a new log filterer instance of MyERC20Token, bound to a specific deployed contract.
func NewMyERC20TokenFilterer(address types.Address, filterer bind.ContractFilterer) (*MyERC20TokenFilterer, error) {
	contract, err := bindMyERC20Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyERC20TokenFilterer{contract: contract}, nil
}

// NewMyERC20TokenCaller creates a new read-only instance of MyERC20Token, bound to a specific deployed contract.
func NewMyERC20TokenBulkCaller(address types.Address, caller bind.ContractCaller) (*MyERC20TokenBulkCaller, error) {
	contract, err := bindMyERC20Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyERC20TokenBulkCaller{contract: contract}, nil
}

// NewMyERC20TokenBulkTransactor creates a new write-only instance of MyERC20Token, bound to a specific deployed contract.
func NewMyERC20TokenBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*MyERC20TokenBulkTransactor, error) {
	contract, err := bindMyERC20Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyERC20TokenBulkTransactor{contract: contract}, nil
}

// bindMyERC20Token binds a generic wrapper to an already deployed contract.
func bindMyERC20Token(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyERC20TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyERC20Token *MyERC20TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyERC20Token.Contract.MyERC20TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyERC20Token *MyERC20TokenRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyERC20Token.Contract.MyERC20TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyERC20Token *MyERC20TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyERC20Token.Contract.MyERC20TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyERC20Token *MyERC20TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyERC20Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyERC20Token *MyERC20TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyERC20Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyERC20Token *MyERC20TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyERC20Token.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MyERC20Token *MyERC20TokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	__err := _MyERC20Token.contract.Call(opts, &out, "balanceOf", arg0)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MyERC20Token *MyERC20TokenBulkCaller) BalanceOf(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, arg0 common.Address) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyERC20Token.contract.GenRequest(opts, "balanceOf", arg0)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyERC20Token.contract.DecodeOutput(&out, rawOut, "balanceOf")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MyERC20Token *MyERC20TokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MyERC20Token.Contract.BalanceOf(&_MyERC20Token.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MyERC20Token *MyERC20TokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MyERC20Token.Contract.BalanceOf(&_MyERC20Token.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyERC20Token *MyERC20TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	__err := _MyERC20Token.contract.Call(opts, &out, "decimals")

	if __err != nil {
		return *new(uint8), __err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, __err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyERC20Token *MyERC20TokenBulkCaller) Decimals(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*uint8, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyERC20Token.contract.GenRequest(opts, "decimals")

	out0 := new(uint8)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyERC20Token.contract.DecodeOutput(&out, rawOut, "decimals")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(uint8)).(*uint8)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyERC20Token *MyERC20TokenSession) Decimals() (uint8, error) {
	return _MyERC20Token.Contract.Decimals(&_MyERC20Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyERC20Token *MyERC20TokenCallerSession) Decimals() (uint8, error) {
	return _MyERC20Token.Contract.Decimals(&_MyERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyERC20Token *MyERC20TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _MyERC20Token.contract.Call(opts, &out, "name")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyERC20Token *MyERC20TokenBulkCaller) Name(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyERC20Token.contract.GenRequest(opts, "name")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyERC20Token.contract.DecodeOutput(&out, rawOut, "name")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyERC20Token *MyERC20TokenSession) Name() (string, error) {
	return _MyERC20Token.Contract.Name(&_MyERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyERC20Token *MyERC20TokenCallerSession) Name() (string, error) {
	return _MyERC20Token.Contract.Name(&_MyERC20Token.CallOpts)
}

// ReturnTuple is a free data retrieval call binding the contract method 0x39009482.
//
// Solidity: function returnTuple() pure returns(uint256 a, bool b)
func (_MyERC20Token *MyERC20TokenCaller) ReturnTuple(opts *bind.CallOpts) (struct {
	A *big.Int
	B bool
}, error) {
	var out []interface{}
	__err := _MyERC20Token.contract.Call(opts, &out, "returnTuple")

	outstruct := new(struct {
		A *big.Int
		B bool
	})
	if __err != nil {
		return *outstruct, __err
	}

	outstruct.A = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.B = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, __err

}

// ReturnTuple is a free data retrieval call binding the contract method 0x39009482.
//
// Solidity: function returnTuple() pure returns(uint256 a, bool b)
func (_MyERC20Token *MyERC20TokenBulkCaller) ReturnTuple(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*struct {
	A *big.Int
	B bool
}, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyERC20Token.contract.GenRequest(opts, "returnTuple")

	outstruct := new(struct {
		A *big.Int
		B bool
	})

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyERC20Token.contract.DecodeOutput(&out, rawOut, "returnTuple")
		if err != nil {
			return err
		}

		outstruct.A = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		outstruct.B = *abi.ConvertType(out[1], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return outstruct, __err

}

// ReturnTuple is a free data retrieval call binding the contract method 0x39009482.
//
// Solidity: function returnTuple() pure returns(uint256 a, bool b)
func (_MyERC20Token *MyERC20TokenSession) ReturnTuple() (struct {
	A *big.Int
	B bool
}, error) {
	return _MyERC20Token.Contract.ReturnTuple(&_MyERC20Token.CallOpts)
}

// ReturnTuple is a free data retrieval call binding the contract method 0x39009482.
//
// Solidity: function returnTuple() pure returns(uint256 a, bool b)
func (_MyERC20Token *MyERC20TokenCallerSession) ReturnTuple() (struct {
	A *big.Int
	B bool
}, error) {
	return _MyERC20Token.Contract.ReturnTuple(&_MyERC20Token.CallOpts)
}

// ReturnTupleWithStruct is a free data retrieval call binding the contract method 0xe354439b.
//
// Solidity: function returnTupleWithStruct() pure returns((string,uint256) s, bool ok)
func (_MyERC20Token *MyERC20TokenCaller) ReturnTupleWithStruct(opts *bind.CallOpts) (struct {
	S  MyERC20TokenStudent
	Ok bool
}, error) {
	var out []interface{}
	__err := _MyERC20Token.contract.Call(opts, &out, "returnTupleWithStruct")

	outstruct := new(struct {
		S  MyERC20TokenStudent
		Ok bool
	})
	if __err != nil {
		return *outstruct, __err
	}

	outstruct.S = *abi.ConvertType(out[0], new(MyERC20TokenStudent)).(*MyERC20TokenStudent)
	outstruct.Ok = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, __err

}

// ReturnTupleWithStruct is a free data retrieval call binding the contract method 0xe354439b.
//
// Solidity: function returnTupleWithStruct() pure returns((string,uint256) s, bool ok)
func (_MyERC20Token *MyERC20TokenBulkCaller) ReturnTupleWithStruct(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*struct {
	S  MyERC20TokenStudent
	Ok bool
}, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyERC20Token.contract.GenRequest(opts, "returnTupleWithStruct")

	outstruct := new(struct {
		S  MyERC20TokenStudent
		Ok bool
	})

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyERC20Token.contract.DecodeOutput(&out, rawOut, "returnTupleWithStruct")
		if err != nil {
			return err
		}

		outstruct.S = *abi.ConvertType(out[0], new(MyERC20TokenStudent)).(*MyERC20TokenStudent)
		outstruct.Ok = *abi.ConvertType(out[1], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return outstruct, __err

}

// ReturnTupleWithStruct is a free data retrieval call binding the contract method 0xe354439b.
//
// Solidity: function returnTupleWithStruct() pure returns((string,uint256) s, bool ok)
func (_MyERC20Token *MyERC20TokenSession) ReturnTupleWithStruct() (struct {
	S  MyERC20TokenStudent
	Ok bool
}, error) {
	return _MyERC20Token.Contract.ReturnTupleWithStruct(&_MyERC20Token.CallOpts)
}

// ReturnTupleWithStruct is a free data retrieval call binding the contract method 0xe354439b.
//
// Solidity: function returnTupleWithStruct() pure returns((string,uint256) s, bool ok)
func (_MyERC20Token *MyERC20TokenCallerSession) ReturnTupleWithStruct() (struct {
	S  MyERC20TokenStudent
	Ok bool
}, error) {
	return _MyERC20Token.Contract.ReturnTupleWithStruct(&_MyERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyERC20Token *MyERC20TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	__err := _MyERC20Token.contract.Call(opts, &out, "symbol")

	if __err != nil {
		return *new(string), __err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, __err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyERC20Token *MyERC20TokenBulkCaller) Symbol(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*string, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _MyERC20Token.contract.GenRequest(opts, "symbol")

	out0 := new(string)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _MyERC20Token.contract.DecodeOutput(&out, rawOut, "symbol")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(string)).(*string)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyERC20Token *MyERC20TokenSession) Symbol() (string, error) {
	return _MyERC20Token.Contract.Symbol(&_MyERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyERC20Token *MyERC20TokenCallerSession) Symbol() (string, error) {
	return _MyERC20Token.Contract.Symbol(&_MyERC20Token.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_MyERC20Token *MyERC20TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyERC20Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_MyERC20Token *MyERC20TokenBulkTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) types.UnsignedTransaction {
	return _MyERC20Token.contract.GenUnsignedTransaction(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_MyERC20Token *MyERC20TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyERC20Token.Contract.Transfer(&_MyERC20Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_MyERC20Token *MyERC20TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _MyERC20Token.Contract.Transfer(&_MyERC20Token.TransactOpts, _to, _value)
}

// MyERC20TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyERC20Token contract.
type MyERC20TokenTransferIterator struct {
	Event *MyERC20TokenTransfer // Event containing the contract specifics and raw log

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
func (it *MyERC20TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyERC20TokenTransfer)
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
		it.Event = new(MyERC20TokenTransfer)
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
func (it *MyERC20TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyERC20TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyERC20TokenTransfer represents a Transfer event raised by the MyERC20Token contract.
type MyERC20TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// MyERC20TokenTransferOrChainReorg represents a Transfer subscription event raised by the MyERC20Token contract.
type MyERC20TokenTransferOrChainReorg struct {
	Event      *MyERC20TokenTransfer
	ChainReorg *types.ChainReorg
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyERC20Token *MyERC20TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MyERC20TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, err := _MyERC20Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyERC20TokenTransferIterator{contract: _MyERC20Token.contract, event: "Transfer", logs: logs}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyERC20Token *MyERC20TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyERC20TokenTransferOrChainReorg, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyERC20Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyERC20TokenTransferOrChainReorg)
				event.Event = new(MyERC20TokenTransfer)

				if log.ChainReorg == nil {
					if err := _MyERC20Token.contract.UnpackLog(event.Event, "Transfer", *log.Log); err != nil {
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
func (_MyERC20Token *MyERC20TokenFilterer) ParseTransfer(log types.Log) (*MyERC20TokenTransfer, error) {
	event := new(MyERC20TokenTransfer)
	if err := _MyERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
