// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// AllPurposeCappedABI is the input ABI used to generate the binding from.
const AllPurposeCappedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isMintable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isBurnable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\"},{\"name\":\"minter\",\"type\":\"address\"},{\"name\":\"isBurnable_\",\"type\":\"bool\"},{\"name\":\"cap_\",\"type\":\"uint256\"},{\"name\":\"isTransferable_\",\"type\":\"bool\"},{\"name\":\"isMintable_\",\"type\":\"bool\"},{\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// AllPurposeCappedBin is the compiled bytecode used for deploying new contracts.
const AllPurposeCappedBin = `60806040526001600960046101000a81548160ff0219169083151502179055503480156200002c57600080fd5b5060405162002c0038038062002c0083398101806040528101908080518201929190602001805182019291906020018051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919050505088888888888787878b620000c73362000238640100000000026401000000009004565b600081111515620000d757600080fd5b8060048190555050620000f933620002a2640100000000026401000000009004565b6000600660006101000a81548160ff02191690831515021790555087600790805190602001906200012c9291906200081a565b508660089080519060200190620001459291906200081a565b5085600960006101000a81548160ff021916908360ff16021790555083600960016101000a81548160ff02191690831515021790555082600960026101000a81548160ff02191690831515021790555081600960036101000a81548160ff021916908315150217905550821515620001d157620001d06200030c640100000000026401000000009004565b5b620001eb8562000238640100000000026401000000009004565b620002063382620003d7640100000000026401000000009004565b6000600960046101000a81548160ff0219169083151502179055505050505050505050505050505050505050620008c9565b6200025c81600362000447640100000000026200210d179091906401000000009004565b8073ffffffffffffffffffffffffffffffffffffffff167f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f660405160405180910390a250565b620002c681600562000447640100000000026200210d179091906401000000009004565b8073ffffffffffffffffffffffffffffffffffffffff167f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f860405160405180910390a250565b600960049054906101000a900460ff161515620003b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001807f746869732066756e6374696f6e2063616e206f6e6c792062652072756e206f6e81526020017f206372656174696f6e000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b620003d56200050a64010000000002620016e6176401000000009004565b565b6004546200041582620003f8620005cd640100000000026401000000009004565b620005d76401000000000262001cd2179091906401000000009004565b111515156200042357600080fd5b620004438282620005f964010000000002620021bd176401000000009004565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200048457600080fd5b6200049f828262000758640100000000026401000000009004565b151515620004ac57600080fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6200052433620007ed640100000000026401000000009004565b15156200053057600080fd5b600660009054906101000a900460ff161515156200054d57600080fd5b6001600660006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b6000600254905090565b6000808284019050838110151515620005ef57600080fd5b8091505092915050565b60008273ffffffffffffffffffffffffffffffffffffffff16141515156200062057600080fd5b6200064581600254620005d76401000000000262001cd2179091906401000000009004565b600281905550620006ac816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054620005d76401000000000262001cd2179091906401000000009004565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156200079657600080fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600062000813826005620007586401000000000262001590179091906401000000009004565b9050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200085d57805160ff19168380011785556200088e565b828001600101855582156200088e579182015b828111156200088d57825182559160200191906001019062000870565b5b5090506200089d9190620008a1565b5090565b620008c691905b80821115620008c2576000816000905550600101620008a8565b5090565b90565b61232780620008d96000396000f30060806040526004361061015f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde0314610164578063095ea7b3146101f457806318160ddd146102595780632121dc751461028457806323b872dd146102b3578063313ce56714610338578063355274ea1461036957806339509351146103945780633f4ba83a146103f957806340c10f191461041057806342966c681461047557806346b45af7146104a257806346fbf68e146104d15780635c975abb1461052c5780636ef8d66d1461055b57806370a082311461057257806379cc6790146105c957806382dc1ec4146106165780638456cb5914610659578063883356d91461067057806395d89b411461069f578063983b2d561461072f5780639865027514610772578063a457c2d714610789578063a9059cbb146107ee578063aa271e1a14610853578063dd62ed3e146108ae575b600080fd5b34801561017057600080fd5b50610179610925565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101b957808201518184015260208101905061019e565b50505050905090810190601f1680156101e65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561020057600080fd5b5061023f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506109c3565b604051808215151515815260200191505060405180910390f35b34801561026557600080fd5b5061026e6109f3565b6040518082815260200191505060405180910390f35b34801561029057600080fd5b506102996109fd565b604051808215151515815260200191505060405180910390f35b3480156102bf57600080fd5b5061031e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a10565b604051808215151515815260200191505060405180910390f35b34801561034457600080fd5b5061034d610a42565b604051808260ff1660ff16815260200191505060405180910390f35b34801561037557600080fd5b5061037e610a55565b6040518082815260200191505060405180910390f35b3480156103a057600080fd5b506103df600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a5f565b604051808215151515815260200191505060405180910390f35b34801561040557600080fd5b5061040e610a8f565b005b34801561041c57600080fd5b5061045b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b3b565b604051808215151515815260200191505060405180910390f35b34801561048157600080fd5b506104a060048036038101908080359060200190929190505050610bd3565b005b3480156104ae57600080fd5b506104b7610c63565b604051808215151515815260200191505060405180910390f35b3480156104dd57600080fd5b50610512600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c76565b604051808215151515815260200191505060405180910390f35b34801561053857600080fd5b50610541610c93565b604051808215151515815260200191505060405180910390f35b34801561056757600080fd5b50610570610caa565b005b34801561057e57600080fd5b506105b3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610cb5565b6040518082815260200191505060405180910390f35b3480156105d557600080fd5b50610614600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610cfd565b005b34801561062257600080fd5b50610657600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d8f565b005b34801561066557600080fd5b5061066e610daf565b005b34801561067c57600080fd5b50610685610e63565b604051808215151515815260200191505060405180910390f35b3480156106ab57600080fd5b506106b4610e76565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106f45780820151818401526020810190506106d9565b50505050905090810190601f1680156107215780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561073b57600080fd5b50610770600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f14565b005b34801561077e57600080fd5b50610787610f34565b005b34801561079557600080fd5b506107d4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f3f565b604051808215151515815260200191505060405180910390f35b3480156107fa57600080fd5b50610839600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f6f565b604051808215151515815260200191505060405180910390f35b34801561085f57600080fd5b50610894600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f9f565b604051808215151515815260200191505060405180910390f35b3480156108ba57600080fd5b5061090f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fbc565b6040518082815260200191505060405180910390f35b60078054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109bb5780601f10610990576101008083540402835291602001916109bb565b820191906000526020600020905b81548152906001019060200180831161099e57829003601f168201915b505050505081565b6000600660009054906101000a900460ff161515156109e157600080fd5b6109eb8383611043565b905092915050565b6000600254905090565b600960029054906101000a900460ff1681565b6000600660009054906101000a900460ff16151515610a2e57600080fd5b610a39848484611170565b90509392505050565b600960009054906101000a900460ff1681565b6000600454905090565b6000600660009054906101000a900460ff16151515610a7d57600080fd5b610a878383611322565b905092915050565b600960049054906101000a900460ff161515610b39576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001807f746869732066756e6374696f6e2063616e206f6e6c792062652072756e206f6e81526020017f206372656174696f6e000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b565b6000600960039054906101000a900460ff161515610bc1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f436f696e206e6f74206d696e7461626c6500000000000000000000000000000081525060200191505060405180910390fd5b610bcb8383611559565b905092915050565b600960019054906101000a900460ff161515610c57576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f436f696e206e6f74206275726e61626c6500000000000000000000000000000081525060200191505060405180910390fd5b610c6081611583565b50565b600960039054906101000a900460ff1681565b6000610c8c82600561159090919063ffffffff16565b9050919050565b6000600660009054906101000a900460ff16905090565b610cb333611624565b565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600960019054906101000a900460ff161515610d81576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f436f696e206e6f74206275726e61626c6500000000000000000000000000000081525060200191505060405180910390fd5b610d8b828261167e565b5050565b610d9833610c76565b1515610da357600080fd5b610dac8161168c565b50565b600960049054906101000a900460ff161515610e59576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001807f746869732066756e6374696f6e2063616e206f6e6c792062652072756e206f6e81526020017f206372656174696f6e000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b610e616116e6565b565b600960019054906101000a900460ff1681565b60088054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f0c5780601f10610ee157610100808354040283529160200191610f0c565b820191906000526020600020905b815481529060010190602001808311610eef57829003601f168201915b505050505081565b610f1d33610f9f565b1515610f2857600080fd5b610f3181611796565b50565b610f3d336117f0565b565b6000600660009054906101000a900460ff16151515610f5d57600080fd5b610f67838361184a565b905092915050565b6000600660009054906101000a900460ff16151515610f8d57600080fd5b610f978383611a81565b905092915050565b6000610fb582600361159090919063ffffffff16565b9050919050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561108057600080fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111515156111fd57600080fd5b61128c82600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a9890919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611317848484611ab9565b600190509392505050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561135f57600080fd5b6113ee82600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cd290919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a36001905092915050565b600061156433610f9f565b151561156f57600080fd5b6115798383611cf3565b6001905092915050565b61158d3382611d2b565b50565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156115cd57600080fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b611638816005611eb690919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e60405160405180910390a250565b6116888282611f65565b5050565b6116a081600561210d90919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f860405160405180910390a250565b6116ef33610c76565b15156116fa57600080fd5b600660009054906101000a900460ff1615151561171657600080fd5b6001600660006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b6117aa81600361210d90919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f660405160405180910390a250565b611804816003611eb690919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669260405160405180910390a250565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561188757600080fd5b61191682600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a9890919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a36001905092915050565b6000611a8e338484611ab9565b6001905092915050565b600080838311151515611aaa57600080fd5b82840390508091505092915050565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111151515611b0657600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515611b4257600080fd5b611b93816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a9890919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611c26816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cd290919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000808284019050838110151515611ce957600080fd5b8091505092915050565b600454611d1082611d026109f3565b611cd290919063ffffffff16565b11151515611d1d57600080fd5b611d2782826121bd565b5050565b60008273ffffffffffffffffffffffffffffffffffffffff1614151515611d5157600080fd5b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111151515611d9e57600080fd5b611db381600254611a9890919063ffffffff16565b600281905550611e0a816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a9890919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611ef257600080fd5b611efc8282611590565b1515611f0757600080fd5b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111151515611ff057600080fd5b61207f81600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a9890919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506121098282611d2b565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561214957600080fd5b6121538282611590565b15151561215f57600080fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008273ffffffffffffffffffffffffffffffffffffffff16141515156121e357600080fd5b6121f881600254611cd290919063ffffffff16565b60028190555061224f816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cd290919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a350505600a165627a7a723058207d8b704ebae4eec15ef127de9fee43b6b0a2acfac252cb4cb7f614ba34ae5d160029`

// DeployAllPurposeCapped deploys a new Ethereum contract, binding an instance of AllPurposeCapped to it.
func DeployAllPurposeCapped(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8, minter common.Address, isBurnable_ bool, cap_ *big.Int, isTransferable_ bool, isMintable_ bool, initialSupply *big.Int) (common.Address, *types.Transaction, *AllPurposeCapped, error) {
	parsed, err := abi.JSON(strings.NewReader(AllPurposeCappedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AllPurposeCappedBin), backend, name_, symbol_, decimals_, minter, isBurnable_, cap_, isTransferable_, isMintable_, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AllPurposeCapped{AllPurposeCappedCaller: AllPurposeCappedCaller{contract: contract}, AllPurposeCappedTransactor: AllPurposeCappedTransactor{contract: contract}, AllPurposeCappedFilterer: AllPurposeCappedFilterer{contract: contract}}, nil
}

// AllPurposeCapped is an auto generated Go binding around an Ethereum contract.
type AllPurposeCapped struct {
	AllPurposeCappedCaller     // Read-only binding to the contract
	AllPurposeCappedTransactor // Write-only binding to the contract
	AllPurposeCappedFilterer   // Log filterer for contract events
}

// AllPurposeCappedCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllPurposeCappedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllPurposeCappedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllPurposeCappedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllPurposeCappedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllPurposeCappedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllPurposeCappedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllPurposeCappedSession struct {
	Contract     *AllPurposeCapped // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AllPurposeCappedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllPurposeCappedCallerSession struct {
	Contract *AllPurposeCappedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AllPurposeCappedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllPurposeCappedTransactorSession struct {
	Contract     *AllPurposeCappedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AllPurposeCappedRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllPurposeCappedRaw struct {
	Contract *AllPurposeCapped // Generic contract binding to access the raw methods on
}

// AllPurposeCappedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllPurposeCappedCallerRaw struct {
	Contract *AllPurposeCappedCaller // Generic read-only contract binding to access the raw methods on
}

// AllPurposeCappedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllPurposeCappedTransactorRaw struct {
	Contract *AllPurposeCappedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllPurposeCapped creates a new instance of AllPurposeCapped, bound to a specific deployed contract.
func NewAllPurposeCapped(address common.Address, backend bind.ContractBackend) (*AllPurposeCapped, error) {
	contract, err := bindAllPurposeCapped(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AllPurposeCapped{AllPurposeCappedCaller: AllPurposeCappedCaller{contract: contract}, AllPurposeCappedTransactor: AllPurposeCappedTransactor{contract: contract}, AllPurposeCappedFilterer: AllPurposeCappedFilterer{contract: contract}}, nil
}

// NewAllPurposeCappedCaller creates a new read-only instance of AllPurposeCapped, bound to a specific deployed contract.
func NewAllPurposeCappedCaller(address common.Address, caller bind.ContractCaller) (*AllPurposeCappedCaller, error) {
	contract, err := bindAllPurposeCapped(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllPurposeCappedCaller{contract: contract}, nil
}

// NewAllPurposeCappedTransactor creates a new write-only instance of AllPurposeCapped, bound to a specific deployed contract.
func NewAllPurposeCappedTransactor(address common.Address, transactor bind.ContractTransactor) (*AllPurposeCappedTransactor, error) {
	contract, err := bindAllPurposeCapped(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllPurposeCappedTransactor{contract: contract}, nil
}

// NewAllPurposeCappedFilterer creates a new log filterer instance of AllPurposeCapped, bound to a specific deployed contract.
func NewAllPurposeCappedFilterer(address common.Address, filterer bind.ContractFilterer) (*AllPurposeCappedFilterer, error) {
	contract, err := bindAllPurposeCapped(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllPurposeCappedFilterer{contract: contract}, nil
}

// bindAllPurposeCapped binds a generic wrapper to an already deployed contract.
func bindAllPurposeCapped(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AllPurposeCappedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllPurposeCapped *AllPurposeCappedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AllPurposeCapped.Contract.AllPurposeCappedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllPurposeCapped *AllPurposeCappedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.AllPurposeCappedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllPurposeCapped *AllPurposeCappedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.AllPurposeCappedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllPurposeCapped *AllPurposeCappedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AllPurposeCapped.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllPurposeCapped *AllPurposeCappedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllPurposeCapped *AllPurposeCappedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_AllPurposeCapped *AllPurposeCappedCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_AllPurposeCapped *AllPurposeCappedSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AllPurposeCapped.Contract.Allowance(&_AllPurposeCapped.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AllPurposeCapped.Contract.Allowance(&_AllPurposeCapped.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_AllPurposeCapped *AllPurposeCappedCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_AllPurposeCapped *AllPurposeCappedSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _AllPurposeCapped.Contract.BalanceOf(&_AllPurposeCapped.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _AllPurposeCapped.Contract.BalanceOf(&_AllPurposeCapped.CallOpts, owner)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_AllPurposeCapped *AllPurposeCappedCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_AllPurposeCapped *AllPurposeCappedSession) Cap() (*big.Int, error) {
	return _AllPurposeCapped.Contract.Cap(&_AllPurposeCapped.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) Cap() (*big.Int, error) {
	return _AllPurposeCapped.Contract.Cap(&_AllPurposeCapped.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_AllPurposeCapped *AllPurposeCappedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_AllPurposeCapped *AllPurposeCappedSession) Decimals() (uint8, error) {
	return _AllPurposeCapped.Contract.Decimals(&_AllPurposeCapped.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) Decimals() (uint8, error) {
	return _AllPurposeCapped.Contract.Decimals(&_AllPurposeCapped.CallOpts)
}

// IsBurnable is a free data retrieval call binding the contract method 0x883356d9.
//
// Solidity: function isBurnable() constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedCaller) IsBurnable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "isBurnable")
	return *ret0, err
}

// IsBurnable is a free data retrieval call binding the contract method 0x883356d9.
//
// Solidity: function isBurnable() constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedSession) IsBurnable() (bool, error) {
	return _AllPurposeCapped.Contract.IsBurnable(&_AllPurposeCapped.CallOpts)
}

// IsBurnable is a free data retrieval call binding the contract method 0x883356d9.
//
// Solidity: function isBurnable() constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) IsBurnable() (bool, error) {
	return _AllPurposeCapped.Contract.IsBurnable(&_AllPurposeCapped.CallOpts)
}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedCaller) IsMintable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "isMintable")
	return *ret0, err
}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedSession) IsMintable() (bool, error) {
	return _AllPurposeCapped.Contract.IsMintable(&_AllPurposeCapped.CallOpts)
}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) IsMintable() (bool, error) {
	return _AllPurposeCapped.Contract.IsMintable(&_AllPurposeCapped.CallOpts)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedSession) IsMinter(account common.Address) (bool, error) {
	return _AllPurposeCapped.Contract.IsMinter(&_AllPurposeCapped.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) IsMinter(account common.Address) (bool, error) {
	return _AllPurposeCapped.Contract.IsMinter(&_AllPurposeCapped.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedSession) IsPauser(account common.Address) (bool, error) {
	return _AllPurposeCapped.Contract.IsPauser(&_AllPurposeCapped.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) IsPauser(account common.Address) (bool, error) {
	return _AllPurposeCapped.Contract.IsPauser(&_AllPurposeCapped.CallOpts, account)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedSession) IsTransferable() (bool, error) {
	return _AllPurposeCapped.Contract.IsTransferable(&_AllPurposeCapped.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) IsTransferable() (bool, error) {
	return _AllPurposeCapped.Contract.IsTransferable(&_AllPurposeCapped.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_AllPurposeCapped *AllPurposeCappedCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_AllPurposeCapped *AllPurposeCappedSession) Name() (string, error) {
	return _AllPurposeCapped.Contract.Name(&_AllPurposeCapped.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) Name() (string, error) {
	return _AllPurposeCapped.Contract.Name(&_AllPurposeCapped.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedSession) Paused() (bool, error) {
	return _AllPurposeCapped.Contract.Paused(&_AllPurposeCapped.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) Paused() (bool, error) {
	return _AllPurposeCapped.Contract.Paused(&_AllPurposeCapped.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_AllPurposeCapped *AllPurposeCappedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_AllPurposeCapped *AllPurposeCappedSession) Symbol() (string, error) {
	return _AllPurposeCapped.Contract.Symbol(&_AllPurposeCapped.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) Symbol() (string, error) {
	return _AllPurposeCapped.Contract.Symbol(&_AllPurposeCapped.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_AllPurposeCapped *AllPurposeCappedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllPurposeCapped.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_AllPurposeCapped *AllPurposeCappedSession) TotalSupply() (*big.Int, error) {
	return _AllPurposeCapped.Contract.TotalSupply(&_AllPurposeCapped.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_AllPurposeCapped *AllPurposeCappedCallerSession) TotalSupply() (*big.Int, error) {
	return _AllPurposeCapped.Contract.TotalSupply(&_AllPurposeCapped.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_AllPurposeCapped *AllPurposeCappedTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_AllPurposeCapped *AllPurposeCappedSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.AddMinter(&_AllPurposeCapped.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.AddMinter(&_AllPurposeCapped.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_AllPurposeCapped *AllPurposeCappedTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_AllPurposeCapped *AllPurposeCappedSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.AddPauser(&_AllPurposeCapped.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.AddPauser(&_AllPurposeCapped.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AllPurposeCapped *AllPurposeCappedTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AllPurposeCapped *AllPurposeCappedSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.Approve(&_AllPurposeCapped.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.Approve(&_AllPurposeCapped.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_AllPurposeCapped *AllPurposeCappedTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_AllPurposeCapped *AllPurposeCappedSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.Burn(&_AllPurposeCapped.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.Burn(&_AllPurposeCapped.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 value) returns()
func (_AllPurposeCapped *AllPurposeCappedTransactor) BurnFrom(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "burnFrom", from, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 value) returns()
func (_AllPurposeCapped *AllPurposeCappedSession) BurnFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.BurnFrom(&_AllPurposeCapped.TransactOpts, from, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 value) returns()
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) BurnFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.BurnFrom(&_AllPurposeCapped.TransactOpts, from, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool success)
func (_AllPurposeCapped *AllPurposeCappedTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool success)
func (_AllPurposeCapped *AllPurposeCappedSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.DecreaseAllowance(&_AllPurposeCapped.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool success)
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.DecreaseAllowance(&_AllPurposeCapped.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool success)
func (_AllPurposeCapped *AllPurposeCappedTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool success)
func (_AllPurposeCapped *AllPurposeCappedSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.IncreaseAllowance(&_AllPurposeCapped.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool success)
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.IncreaseAllowance(&_AllPurposeCapped.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_AllPurposeCapped *AllPurposeCappedTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_AllPurposeCapped *AllPurposeCappedSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.Mint(&_AllPurposeCapped.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.Mint(&_AllPurposeCapped.TransactOpts, to, value)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AllPurposeCapped *AllPurposeCappedTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AllPurposeCapped *AllPurposeCappedSession) Pause() (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.Pause(&_AllPurposeCapped.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) Pause() (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.Pause(&_AllPurposeCapped.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_AllPurposeCapped *AllPurposeCappedTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_AllPurposeCapped *AllPurposeCappedSession) RenounceMinter() (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.RenounceMinter(&_AllPurposeCapped.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.RenounceMinter(&_AllPurposeCapped.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_AllPurposeCapped *AllPurposeCappedTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_AllPurposeCapped *AllPurposeCappedSession) RenouncePauser() (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.RenouncePauser(&_AllPurposeCapped.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.RenouncePauser(&_AllPurposeCapped.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AllPurposeCapped *AllPurposeCappedTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AllPurposeCapped *AllPurposeCappedSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.Transfer(&_AllPurposeCapped.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.Transfer(&_AllPurposeCapped.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AllPurposeCapped *AllPurposeCappedTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AllPurposeCapped *AllPurposeCappedSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.TransferFrom(&_AllPurposeCapped.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.TransferFrom(&_AllPurposeCapped.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AllPurposeCapped *AllPurposeCappedTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllPurposeCapped.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AllPurposeCapped *AllPurposeCappedSession) Unpause() (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.Unpause(&_AllPurposeCapped.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AllPurposeCapped *AllPurposeCappedTransactorSession) Unpause() (*types.Transaction, error) {
	return _AllPurposeCapped.Contract.Unpause(&_AllPurposeCapped.TransactOpts)
}

// AllPurposeCappedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AllPurposeCapped contract.
type AllPurposeCappedApprovalIterator struct {
	Event *AllPurposeCappedApproval // Event containing the contract specifics and raw log

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
func (it *AllPurposeCappedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeCappedApproval)
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
		it.Event = new(AllPurposeCappedApproval)
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
func (it *AllPurposeCappedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeCappedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeCappedApproval represents a Approval event raised by the AllPurposeCapped contract.
type AllPurposeCappedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AllPurposeCapped *AllPurposeCappedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AllPurposeCappedApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AllPurposeCapped.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AllPurposeCappedApprovalIterator{contract: _AllPurposeCapped.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AllPurposeCapped *AllPurposeCappedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AllPurposeCappedApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AllPurposeCapped.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeCappedApproval)
				if err := _AllPurposeCapped.contract.UnpackLog(event, "Approval", log); err != nil {
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

// AllPurposeCappedMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the AllPurposeCapped contract.
type AllPurposeCappedMinterAddedIterator struct {
	Event *AllPurposeCappedMinterAdded // Event containing the contract specifics and raw log

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
func (it *AllPurposeCappedMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeCappedMinterAdded)
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
		it.Event = new(AllPurposeCappedMinterAdded)
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
func (it *AllPurposeCappedMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeCappedMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeCappedMinterAdded represents a MinterAdded event raised by the AllPurposeCapped contract.
type AllPurposeCappedMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_AllPurposeCapped *AllPurposeCappedFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*AllPurposeCappedMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurposeCapped.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &AllPurposeCappedMinterAddedIterator{contract: _AllPurposeCapped.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_AllPurposeCapped *AllPurposeCappedFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *AllPurposeCappedMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurposeCapped.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeCappedMinterAdded)
				if err := _AllPurposeCapped.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// AllPurposeCappedMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the AllPurposeCapped contract.
type AllPurposeCappedMinterRemovedIterator struct {
	Event *AllPurposeCappedMinterRemoved // Event containing the contract specifics and raw log

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
func (it *AllPurposeCappedMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeCappedMinterRemoved)
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
		it.Event = new(AllPurposeCappedMinterRemoved)
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
func (it *AllPurposeCappedMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeCappedMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeCappedMinterRemoved represents a MinterRemoved event raised by the AllPurposeCapped contract.
type AllPurposeCappedMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_AllPurposeCapped *AllPurposeCappedFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*AllPurposeCappedMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurposeCapped.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AllPurposeCappedMinterRemovedIterator{contract: _AllPurposeCapped.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_AllPurposeCapped *AllPurposeCappedFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *AllPurposeCappedMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurposeCapped.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeCappedMinterRemoved)
				if err := _AllPurposeCapped.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// AllPurposeCappedPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AllPurposeCapped contract.
type AllPurposeCappedPausedIterator struct {
	Event *AllPurposeCappedPaused // Event containing the contract specifics and raw log

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
func (it *AllPurposeCappedPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeCappedPaused)
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
		it.Event = new(AllPurposeCappedPaused)
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
func (it *AllPurposeCappedPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeCappedPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeCappedPaused represents a Paused event raised by the AllPurposeCapped contract.
type AllPurposeCappedPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AllPurposeCapped *AllPurposeCappedFilterer) FilterPaused(opts *bind.FilterOpts) (*AllPurposeCappedPausedIterator, error) {

	logs, sub, err := _AllPurposeCapped.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AllPurposeCappedPausedIterator{contract: _AllPurposeCapped.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AllPurposeCapped *AllPurposeCappedFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AllPurposeCappedPaused) (event.Subscription, error) {

	logs, sub, err := _AllPurposeCapped.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeCappedPaused)
				if err := _AllPurposeCapped.contract.UnpackLog(event, "Paused", log); err != nil {
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

// AllPurposeCappedPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the AllPurposeCapped contract.
type AllPurposeCappedPauserAddedIterator struct {
	Event *AllPurposeCappedPauserAdded // Event containing the contract specifics and raw log

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
func (it *AllPurposeCappedPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeCappedPauserAdded)
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
		it.Event = new(AllPurposeCappedPauserAdded)
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
func (it *AllPurposeCappedPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeCappedPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeCappedPauserAdded represents a PauserAdded event raised by the AllPurposeCapped contract.
type AllPurposeCappedPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_AllPurposeCapped *AllPurposeCappedFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*AllPurposeCappedPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurposeCapped.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &AllPurposeCappedPauserAddedIterator{contract: _AllPurposeCapped.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_AllPurposeCapped *AllPurposeCappedFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *AllPurposeCappedPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurposeCapped.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeCappedPauserAdded)
				if err := _AllPurposeCapped.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// AllPurposeCappedPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the AllPurposeCapped contract.
type AllPurposeCappedPauserRemovedIterator struct {
	Event *AllPurposeCappedPauserRemoved // Event containing the contract specifics and raw log

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
func (it *AllPurposeCappedPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeCappedPauserRemoved)
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
		it.Event = new(AllPurposeCappedPauserRemoved)
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
func (it *AllPurposeCappedPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeCappedPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeCappedPauserRemoved represents a PauserRemoved event raised by the AllPurposeCapped contract.
type AllPurposeCappedPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_AllPurposeCapped *AllPurposeCappedFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*AllPurposeCappedPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurposeCapped.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AllPurposeCappedPauserRemovedIterator{contract: _AllPurposeCapped.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_AllPurposeCapped *AllPurposeCappedFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *AllPurposeCappedPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurposeCapped.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeCappedPauserRemoved)
				if err := _AllPurposeCapped.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// AllPurposeCappedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AllPurposeCapped contract.
type AllPurposeCappedTransferIterator struct {
	Event *AllPurposeCappedTransfer // Event containing the contract specifics and raw log

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
func (it *AllPurposeCappedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeCappedTransfer)
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
		it.Event = new(AllPurposeCappedTransfer)
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
func (it *AllPurposeCappedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeCappedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeCappedTransfer represents a Transfer event raised by the AllPurposeCapped contract.
type AllPurposeCappedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AllPurposeCapped *AllPurposeCappedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AllPurposeCappedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AllPurposeCapped.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AllPurposeCappedTransferIterator{contract: _AllPurposeCapped.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AllPurposeCapped *AllPurposeCappedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AllPurposeCappedTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AllPurposeCapped.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeCappedTransfer)
				if err := _AllPurposeCapped.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// AllPurposeCappedUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AllPurposeCapped contract.
type AllPurposeCappedUnpausedIterator struct {
	Event *AllPurposeCappedUnpaused // Event containing the contract specifics and raw log

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
func (it *AllPurposeCappedUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeCappedUnpaused)
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
		it.Event = new(AllPurposeCappedUnpaused)
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
func (it *AllPurposeCappedUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeCappedUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeCappedUnpaused represents a Unpaused event raised by the AllPurposeCapped contract.
type AllPurposeCappedUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AllPurposeCapped *AllPurposeCappedFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AllPurposeCappedUnpausedIterator, error) {

	logs, sub, err := _AllPurposeCapped.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AllPurposeCappedUnpausedIterator{contract: _AllPurposeCapped.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AllPurposeCapped *AllPurposeCappedFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AllPurposeCappedUnpaused) (event.Subscription, error) {

	logs, sub, err := _AllPurposeCapped.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeCappedUnpaused)
				if err := _AllPurposeCapped.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
