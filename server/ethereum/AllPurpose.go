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

// AllPurposeABI is the input ABI used to generate the binding from.
const AllPurposeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isMintable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isBurnable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\"},{\"name\":\"minter\",\"type\":\"address\"},{\"name\":\"isBurnable_\",\"type\":\"bool\"},{\"name\":\"isTransferable_\",\"type\":\"bool\"},{\"name\":\"isMintable_\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// AllPurposeBin is the compiled bytecode used for deploying new contracts.
const AllPurposeBin = `60806040526001600860046101000a81548160ff021916908315150217905550600a6009553480156200003157600080fd5b5060405162002b1238038062002b1283398101806040528101908080518201929190602001805182019291906020018051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050620000af3362000200640100000000026401000000009004565b620000c9336200026a640100000000026401000000009004565b6000600560006101000a81548160ff0219169083151502179055508660069080519060200190620000fc92919062000768565b5085600790805190602001906200011592919062000768565b5084600860006101000a81548160ff021916908360ff16021790555082600860016101000a81548160ff02191690831515021790555081600860026101000a81548160ff02191690831515021790555080600860036101000a81548160ff021916908315150217905550811515620001a157620001a0620002d4640100000000026401000000009004565b5b620001bb8462000200640100000000026401000000009004565b620001d8336009546200039f640100000000026401000000009004565b6000600860046101000a81548160ff0219169083151502179055505050505050505062000817565b62000224816003620004fe640100000000026200220f179091906401000000009004565b8073ffffffffffffffffffffffffffffffffffffffff167f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f660405160405180910390a250565b6200028e816004620004fe640100000000026200220f179091906401000000009004565b8073ffffffffffffffffffffffffffffffffffffffff167f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f860405160405180910390a250565b600860049054906101000a900460ff1615156200037f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001807f746869732066756e6374696f6e2063616e206f6e6c792062652072756e206f6e81526020017f206372656174696f6e000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b6200039d620005c164010000000002620016e2176401000000009004565b565b60008273ffffffffffffffffffffffffffffffffffffffff1614151515620003c657600080fd5b620003eb81600254620006846401000000000262001cce179091906401000000009004565b60028190555062000452816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054620006846401000000000262001cce179091906401000000009004565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200053b57600080fd5b620005568282620006a6640100000000026401000000009004565b1515156200056357600080fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b620005db336200073b640100000000026401000000009004565b1515620005e757600080fd5b600560009054906101000a900460ff161515156200060457600080fd5b6001600560006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b60008082840190508381101515156200069c57600080fd5b8091505092915050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515620006e457600080fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600062000761826004620006a6640100000000026200158c179091906401000000009004565b9050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620007ab57805160ff1916838001178555620007dc565b82800160010185558215620007dc579182015b82811115620007db578251825591602001919060010190620007be565b5b509050620007eb9190620007ef565b5090565b6200081491905b8082111562000810576000816000905550600101620007f6565b5090565b90565b6122eb80620008276000396000f30060806040526004361061015f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde0314610164578063095ea7b3146101f457806318160ddd146102595780632121dc751461028457806323b872dd146102b35780632ff2e9dc14610338578063313ce5671461036357806339509351146103945780633f4ba83a146103f957806340c10f191461041057806342966c681461047557806346b45af7146104a257806346fbf68e146104d15780635c975abb1461052c5780636ef8d66d1461055b57806370a082311461057257806379cc6790146105c957806382dc1ec4146106165780638456cb5914610659578063883356d91461067057806395d89b411461069f578063983b2d561461072f5780639865027514610772578063a457c2d714610789578063a9059cbb146107ee578063aa271e1a14610853578063dd62ed3e146108ae575b600080fd5b34801561017057600080fd5b50610179610925565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101b957808201518184015260208101905061019e565b50505050905090810190601f1680156101e65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561020057600080fd5b5061023f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506109c3565b604051808215151515815260200191505060405180910390f35b34801561026557600080fd5b5061026e6109f3565b6040518082815260200191505060405180910390f35b34801561029057600080fd5b506102996109fd565b604051808215151515815260200191505060405180910390f35b3480156102bf57600080fd5b5061031e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a10565b604051808215151515815260200191505060405180910390f35b34801561034457600080fd5b5061034d610a42565b6040518082815260200191505060405180910390f35b34801561036f57600080fd5b50610378610a48565b604051808260ff1660ff16815260200191505060405180910390f35b3480156103a057600080fd5b506103df600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a5b565b604051808215151515815260200191505060405180910390f35b34801561040557600080fd5b5061040e610a8b565b005b34801561041c57600080fd5b5061045b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b37565b604051808215151515815260200191505060405180910390f35b34801561048157600080fd5b506104a060048036038101908080359060200190929190505050610bcf565b005b3480156104ae57600080fd5b506104b7610c5f565b604051808215151515815260200191505060405180910390f35b3480156104dd57600080fd5b50610512600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c72565b604051808215151515815260200191505060405180910390f35b34801561053857600080fd5b50610541610c8f565b604051808215151515815260200191505060405180910390f35b34801561056757600080fd5b50610570610ca6565b005b34801561057e57600080fd5b506105b3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610cb1565b6040518082815260200191505060405180910390f35b3480156105d557600080fd5b50610614600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610cf9565b005b34801561062257600080fd5b50610657600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d8b565b005b34801561066557600080fd5b5061066e610dab565b005b34801561067c57600080fd5b50610685610e5f565b604051808215151515815260200191505060405180910390f35b3480156106ab57600080fd5b506106b4610e72565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106f45780820151818401526020810190506106d9565b50505050905090810190601f1680156107215780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561073b57600080fd5b50610770600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f10565b005b34801561077e57600080fd5b50610787610f30565b005b34801561079557600080fd5b506107d4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f3b565b604051808215151515815260200191505060405180910390f35b3480156107fa57600080fd5b50610839600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f6b565b604051808215151515815260200191505060405180910390f35b34801561085f57600080fd5b50610894600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f9b565b604051808215151515815260200191505060405180910390f35b3480156108ba57600080fd5b5061090f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fb8565b6040518082815260200191505060405180910390f35b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109bb5780601f10610990576101008083540402835291602001916109bb565b820191906000526020600020905b81548152906001019060200180831161099e57829003601f168201915b505050505081565b6000600560009054906101000a900460ff161515156109e157600080fd5b6109eb838361103f565b905092915050565b6000600254905090565b600860029054906101000a900460ff1681565b6000600560009054906101000a900460ff16151515610a2e57600080fd5b610a3984848461116c565b90509392505050565b60095481565b600860009054906101000a900460ff1681565b6000600560009054906101000a900460ff16151515610a7957600080fd5b610a83838361131e565b905092915050565b600860049054906101000a900460ff161515610b35576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001807f746869732066756e6374696f6e2063616e206f6e6c792062652072756e206f6e81526020017f206372656174696f6e000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b565b6000600860039054906101000a900460ff161515610bbd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f436f696e206e6f74206d696e7461626c6500000000000000000000000000000081525060200191505060405180910390fd5b610bc78383611555565b905092915050565b600860019054906101000a900460ff161515610c53576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f436f696e206e6f74206275726e61626c6500000000000000000000000000000081525060200191505060405180910390fd5b610c5c8161157f565b50565b600860039054906101000a900460ff1681565b6000610c8882600461158c90919063ffffffff16565b9050919050565b6000600560009054906101000a900460ff16905090565b610caf33611620565b565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600860019054906101000a900460ff161515610d7d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f436f696e206e6f74206275726e61626c6500000000000000000000000000000081525060200191505060405180910390fd5b610d87828261167a565b5050565b610d9433610c72565b1515610d9f57600080fd5b610da881611688565b50565b600860049054906101000a900460ff161515610e55576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001807f746869732066756e6374696f6e2063616e206f6e6c792062652072756e206f6e81526020017f206372656174696f6e000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b610e5d6116e2565b565b600860019054906101000a900460ff1681565b60078054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f085780601f10610edd57610100808354040283529160200191610f08565b820191906000526020600020905b815481529060010190602001808311610eeb57829003601f168201915b505050505081565b610f1933610f9b565b1515610f2457600080fd5b610f2d81611792565b50565b610f39336117ec565b565b6000600560009054906101000a900460ff16151515610f5957600080fd5b610f638383611846565b905092915050565b6000600560009054906101000a900460ff16151515610f8957600080fd5b610f938383611a7d565b905092915050565b6000610fb182600361158c90919063ffffffff16565b9050919050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561107c57600080fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111515156111f957600080fd5b61128882600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a9490919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611313848484611ab5565b600190509392505050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561135b57600080fd5b6113ea82600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cce90919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a36001905092915050565b600061156033610f9b565b151561156b57600080fd5b6115758383611cef565b6001905092915050565b6115893382611e2d565b50565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156115c957600080fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b611634816004611fb890919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e60405160405180910390a250565b6116848282612067565b5050565b61169c81600461220f90919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f860405160405180910390a250565b6116eb33610c72565b15156116f657600080fd5b600560009054906101000a900460ff1615151561171257600080fd5b6001600560006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b6117a681600361220f90919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f660405160405180910390a250565b611800816003611fb890919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669260405160405180910390a250565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561188357600080fd5b61191282600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a9490919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a36001905092915050565b6000611a8a338484611ab5565b6001905092915050565b600080838311151515611aa657600080fd5b82840390508091505092915050565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111151515611b0257600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515611b3e57600080fd5b611b8f816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a9490919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611c22816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cce90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000808284019050838110151515611ce557600080fd5b8091505092915050565b60008273ffffffffffffffffffffffffffffffffffffffff1614151515611d1557600080fd5b611d2a81600254611cce90919063ffffffff16565b600281905550611d81816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cce90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b60008273ffffffffffffffffffffffffffffffffffffffff1614151515611e5357600080fd5b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111151515611ea057600080fd5b611eb581600254611a9490919063ffffffff16565b600281905550611f0c816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a9490919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611ff457600080fd5b611ffe828261158c565b151561200957600080fd5b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481111515156120f257600080fd5b61218181600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611a9490919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061220b8282611e2d565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561224b57600080fd5b612255828261158c565b15151561226157600080fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505600a165627a7a72305820522d4c107c9a85ebc563757a8562cbab8f5162a2a55e472416e62fed8bfd3a080029`

// DeployAllPurpose deploys a new Ethereum contract, binding an instance of AllPurpose to it.
func DeployAllPurpose(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8, minter common.Address, isBurnable_ bool, isTransferable_ bool, isMintable_ bool) (common.Address, *types.Transaction, *AllPurpose, error) {
	parsed, err := abi.JSON(strings.NewReader(AllPurposeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AllPurposeBin), backend, name_, symbol_, decimals_, minter, isBurnable_, isTransferable_, isMintable_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AllPurpose{AllPurposeCaller: AllPurposeCaller{contract: contract}, AllPurposeTransactor: AllPurposeTransactor{contract: contract}, AllPurposeFilterer: AllPurposeFilterer{contract: contract}}, nil
}

// AllPurpose is an auto generated Go binding around an Ethereum contract.
type AllPurpose struct {
	AllPurposeCaller     // Read-only binding to the contract
	AllPurposeTransactor // Write-only binding to the contract
	AllPurposeFilterer   // Log filterer for contract events
}

// AllPurposeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllPurposeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllPurposeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllPurposeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllPurposeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllPurposeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllPurposeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllPurposeSession struct {
	Contract     *AllPurpose       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AllPurposeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllPurposeCallerSession struct {
	Contract *AllPurposeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AllPurposeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllPurposeTransactorSession struct {
	Contract     *AllPurposeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AllPurposeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllPurposeRaw struct {
	Contract *AllPurpose // Generic contract binding to access the raw methods on
}

// AllPurposeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllPurposeCallerRaw struct {
	Contract *AllPurposeCaller // Generic read-only contract binding to access the raw methods on
}

// AllPurposeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllPurposeTransactorRaw struct {
	Contract *AllPurposeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllPurpose creates a new instance of AllPurpose, bound to a specific deployed contract.
func NewAllPurpose(address common.Address, backend bind.ContractBackend) (*AllPurpose, error) {
	contract, err := bindAllPurpose(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AllPurpose{AllPurposeCaller: AllPurposeCaller{contract: contract}, AllPurposeTransactor: AllPurposeTransactor{contract: contract}, AllPurposeFilterer: AllPurposeFilterer{contract: contract}}, nil
}

// NewAllPurposeCaller creates a new read-only instance of AllPurpose, bound to a specific deployed contract.
func NewAllPurposeCaller(address common.Address, caller bind.ContractCaller) (*AllPurposeCaller, error) {
	contract, err := bindAllPurpose(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllPurposeCaller{contract: contract}, nil
}

// NewAllPurposeTransactor creates a new write-only instance of AllPurpose, bound to a specific deployed contract.
func NewAllPurposeTransactor(address common.Address, transactor bind.ContractTransactor) (*AllPurposeTransactor, error) {
	contract, err := bindAllPurpose(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllPurposeTransactor{contract: contract}, nil
}

// NewAllPurposeFilterer creates a new log filterer instance of AllPurpose, bound to a specific deployed contract.
func NewAllPurposeFilterer(address common.Address, filterer bind.ContractFilterer) (*AllPurposeFilterer, error) {
	contract, err := bindAllPurpose(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllPurposeFilterer{contract: contract}, nil
}

// bindAllPurpose binds a generic wrapper to an already deployed contract.
func bindAllPurpose(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AllPurposeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllPurpose *AllPurposeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AllPurpose.Contract.AllPurposeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllPurpose *AllPurposeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllPurpose.Contract.AllPurposeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllPurpose *AllPurposeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllPurpose.Contract.AllPurposeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllPurpose *AllPurposeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AllPurpose.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllPurpose *AllPurposeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllPurpose.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllPurpose *AllPurposeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllPurpose.Contract.contract.Transact(opts, method, params...)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_AllPurpose *AllPurposeCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_AllPurpose *AllPurposeSession) INITIALSUPPLY() (*big.Int, error) {
	return _AllPurpose.Contract.INITIALSUPPLY(&_AllPurpose.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_AllPurpose *AllPurposeCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _AllPurpose.Contract.INITIALSUPPLY(&_AllPurpose.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_AllPurpose *AllPurposeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_AllPurpose *AllPurposeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AllPurpose.Contract.Allowance(&_AllPurpose.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_AllPurpose *AllPurposeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AllPurpose.Contract.Allowance(&_AllPurpose.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_AllPurpose *AllPurposeCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_AllPurpose *AllPurposeSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _AllPurpose.Contract.BalanceOf(&_AllPurpose.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_AllPurpose *AllPurposeCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _AllPurpose.Contract.BalanceOf(&_AllPurpose.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_AllPurpose *AllPurposeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_AllPurpose *AllPurposeSession) Decimals() (uint8, error) {
	return _AllPurpose.Contract.Decimals(&_AllPurpose.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_AllPurpose *AllPurposeCallerSession) Decimals() (uint8, error) {
	return _AllPurpose.Contract.Decimals(&_AllPurpose.CallOpts)
}

// IsBurnable is a free data retrieval call binding the contract method 0x883356d9.
//
// Solidity: function isBurnable() constant returns(bool)
func (_AllPurpose *AllPurposeCaller) IsBurnable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "isBurnable")
	return *ret0, err
}

// IsBurnable is a free data retrieval call binding the contract method 0x883356d9.
//
// Solidity: function isBurnable() constant returns(bool)
func (_AllPurpose *AllPurposeSession) IsBurnable() (bool, error) {
	return _AllPurpose.Contract.IsBurnable(&_AllPurpose.CallOpts)
}

// IsBurnable is a free data retrieval call binding the contract method 0x883356d9.
//
// Solidity: function isBurnable() constant returns(bool)
func (_AllPurpose *AllPurposeCallerSession) IsBurnable() (bool, error) {
	return _AllPurpose.Contract.IsBurnable(&_AllPurpose.CallOpts)
}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() constant returns(bool)
func (_AllPurpose *AllPurposeCaller) IsMintable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "isMintable")
	return *ret0, err
}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() constant returns(bool)
func (_AllPurpose *AllPurposeSession) IsMintable() (bool, error) {
	return _AllPurpose.Contract.IsMintable(&_AllPurpose.CallOpts)
}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() constant returns(bool)
func (_AllPurpose *AllPurposeCallerSession) IsMintable() (bool, error) {
	return _AllPurpose.Contract.IsMintable(&_AllPurpose.CallOpts)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_AllPurpose *AllPurposeCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_AllPurpose *AllPurposeSession) IsMinter(account common.Address) (bool, error) {
	return _AllPurpose.Contract.IsMinter(&_AllPurpose.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_AllPurpose *AllPurposeCallerSession) IsMinter(account common.Address) (bool, error) {
	return _AllPurpose.Contract.IsMinter(&_AllPurpose.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_AllPurpose *AllPurposeCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_AllPurpose *AllPurposeSession) IsPauser(account common.Address) (bool, error) {
	return _AllPurpose.Contract.IsPauser(&_AllPurpose.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_AllPurpose *AllPurposeCallerSession) IsPauser(account common.Address) (bool, error) {
	return _AllPurpose.Contract.IsPauser(&_AllPurpose.CallOpts, account)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_AllPurpose *AllPurposeCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_AllPurpose *AllPurposeSession) IsTransferable() (bool, error) {
	return _AllPurpose.Contract.IsTransferable(&_AllPurpose.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_AllPurpose *AllPurposeCallerSession) IsTransferable() (bool, error) {
	return _AllPurpose.Contract.IsTransferable(&_AllPurpose.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_AllPurpose *AllPurposeCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_AllPurpose *AllPurposeSession) Name() (string, error) {
	return _AllPurpose.Contract.Name(&_AllPurpose.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_AllPurpose *AllPurposeCallerSession) Name() (string, error) {
	return _AllPurpose.Contract.Name(&_AllPurpose.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AllPurpose *AllPurposeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AllPurpose *AllPurposeSession) Paused() (bool, error) {
	return _AllPurpose.Contract.Paused(&_AllPurpose.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AllPurpose *AllPurposeCallerSession) Paused() (bool, error) {
	return _AllPurpose.Contract.Paused(&_AllPurpose.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_AllPurpose *AllPurposeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_AllPurpose *AllPurposeSession) Symbol() (string, error) {
	return _AllPurpose.Contract.Symbol(&_AllPurpose.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_AllPurpose *AllPurposeCallerSession) Symbol() (string, error) {
	return _AllPurpose.Contract.Symbol(&_AllPurpose.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_AllPurpose *AllPurposeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllPurpose.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_AllPurpose *AllPurposeSession) TotalSupply() (*big.Int, error) {
	return _AllPurpose.Contract.TotalSupply(&_AllPurpose.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_AllPurpose *AllPurposeCallerSession) TotalSupply() (*big.Int, error) {
	return _AllPurpose.Contract.TotalSupply(&_AllPurpose.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_AllPurpose *AllPurposeTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_AllPurpose *AllPurposeSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _AllPurpose.Contract.AddMinter(&_AllPurpose.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_AllPurpose *AllPurposeTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _AllPurpose.Contract.AddMinter(&_AllPurpose.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_AllPurpose *AllPurposeTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_AllPurpose *AllPurposeSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _AllPurpose.Contract.AddPauser(&_AllPurpose.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_AllPurpose *AllPurposeTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _AllPurpose.Contract.AddPauser(&_AllPurpose.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AllPurpose *AllPurposeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AllPurpose *AllPurposeSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.Approve(&_AllPurpose.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AllPurpose *AllPurposeTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.Approve(&_AllPurpose.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_AllPurpose *AllPurposeTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_AllPurpose *AllPurposeSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.Burn(&_AllPurpose.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_AllPurpose *AllPurposeTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.Burn(&_AllPurpose.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 value) returns()
func (_AllPurpose *AllPurposeTransactor) BurnFrom(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "burnFrom", from, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 value) returns()
func (_AllPurpose *AllPurposeSession) BurnFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.BurnFrom(&_AllPurpose.TransactOpts, from, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 value) returns()
func (_AllPurpose *AllPurposeTransactorSession) BurnFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.BurnFrom(&_AllPurpose.TransactOpts, from, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool success)
func (_AllPurpose *AllPurposeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool success)
func (_AllPurpose *AllPurposeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.DecreaseAllowance(&_AllPurpose.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool success)
func (_AllPurpose *AllPurposeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.DecreaseAllowance(&_AllPurpose.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool success)
func (_AllPurpose *AllPurposeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool success)
func (_AllPurpose *AllPurposeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.IncreaseAllowance(&_AllPurpose.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool success)
func (_AllPurpose *AllPurposeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.IncreaseAllowance(&_AllPurpose.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_AllPurpose *AllPurposeTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_AllPurpose *AllPurposeSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.Mint(&_AllPurpose.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_AllPurpose *AllPurposeTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.Mint(&_AllPurpose.TransactOpts, to, value)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AllPurpose *AllPurposeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AllPurpose *AllPurposeSession) Pause() (*types.Transaction, error) {
	return _AllPurpose.Contract.Pause(&_AllPurpose.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AllPurpose *AllPurposeTransactorSession) Pause() (*types.Transaction, error) {
	return _AllPurpose.Contract.Pause(&_AllPurpose.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_AllPurpose *AllPurposeTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_AllPurpose *AllPurposeSession) RenounceMinter() (*types.Transaction, error) {
	return _AllPurpose.Contract.RenounceMinter(&_AllPurpose.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_AllPurpose *AllPurposeTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _AllPurpose.Contract.RenounceMinter(&_AllPurpose.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_AllPurpose *AllPurposeTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_AllPurpose *AllPurposeSession) RenouncePauser() (*types.Transaction, error) {
	return _AllPurpose.Contract.RenouncePauser(&_AllPurpose.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_AllPurpose *AllPurposeTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _AllPurpose.Contract.RenouncePauser(&_AllPurpose.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AllPurpose *AllPurposeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AllPurpose *AllPurposeSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.Transfer(&_AllPurpose.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AllPurpose *AllPurposeTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.Transfer(&_AllPurpose.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AllPurpose *AllPurposeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AllPurpose *AllPurposeSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.TransferFrom(&_AllPurpose.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AllPurpose *AllPurposeTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllPurpose.Contract.TransferFrom(&_AllPurpose.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AllPurpose *AllPurposeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllPurpose.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AllPurpose *AllPurposeSession) Unpause() (*types.Transaction, error) {
	return _AllPurpose.Contract.Unpause(&_AllPurpose.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AllPurpose *AllPurposeTransactorSession) Unpause() (*types.Transaction, error) {
	return _AllPurpose.Contract.Unpause(&_AllPurpose.TransactOpts)
}

// AllPurposeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AllPurpose contract.
type AllPurposeApprovalIterator struct {
	Event *AllPurposeApproval // Event containing the contract specifics and raw log

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
func (it *AllPurposeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeApproval)
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
		it.Event = new(AllPurposeApproval)
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
func (it *AllPurposeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeApproval represents a Approval event raised by the AllPurpose contract.
type AllPurposeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AllPurpose *AllPurposeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AllPurposeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AllPurpose.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AllPurposeApprovalIterator{contract: _AllPurpose.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AllPurpose *AllPurposeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AllPurposeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AllPurpose.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeApproval)
				if err := _AllPurpose.contract.UnpackLog(event, "Approval", log); err != nil {
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

// AllPurposeMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the AllPurpose contract.
type AllPurposeMinterAddedIterator struct {
	Event *AllPurposeMinterAdded // Event containing the contract specifics and raw log

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
func (it *AllPurposeMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeMinterAdded)
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
		it.Event = new(AllPurposeMinterAdded)
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
func (it *AllPurposeMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeMinterAdded represents a MinterAdded event raised by the AllPurpose contract.
type AllPurposeMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_AllPurpose *AllPurposeFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*AllPurposeMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurpose.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &AllPurposeMinterAddedIterator{contract: _AllPurpose.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_AllPurpose *AllPurposeFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *AllPurposeMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurpose.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeMinterAdded)
				if err := _AllPurpose.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// AllPurposeMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the AllPurpose contract.
type AllPurposeMinterRemovedIterator struct {
	Event *AllPurposeMinterRemoved // Event containing the contract specifics and raw log

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
func (it *AllPurposeMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeMinterRemoved)
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
		it.Event = new(AllPurposeMinterRemoved)
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
func (it *AllPurposeMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeMinterRemoved represents a MinterRemoved event raised by the AllPurpose contract.
type AllPurposeMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_AllPurpose *AllPurposeFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*AllPurposeMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurpose.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AllPurposeMinterRemovedIterator{contract: _AllPurpose.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_AllPurpose *AllPurposeFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *AllPurposeMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurpose.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeMinterRemoved)
				if err := _AllPurpose.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// AllPurposePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AllPurpose contract.
type AllPurposePausedIterator struct {
	Event *AllPurposePaused // Event containing the contract specifics and raw log

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
func (it *AllPurposePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposePaused)
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
		it.Event = new(AllPurposePaused)
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
func (it *AllPurposePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposePaused represents a Paused event raised by the AllPurpose contract.
type AllPurposePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AllPurpose *AllPurposeFilterer) FilterPaused(opts *bind.FilterOpts) (*AllPurposePausedIterator, error) {

	logs, sub, err := _AllPurpose.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AllPurposePausedIterator{contract: _AllPurpose.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AllPurpose *AllPurposeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AllPurposePaused) (event.Subscription, error) {

	logs, sub, err := _AllPurpose.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposePaused)
				if err := _AllPurpose.contract.UnpackLog(event, "Paused", log); err != nil {
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

// AllPurposePauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the AllPurpose contract.
type AllPurposePauserAddedIterator struct {
	Event *AllPurposePauserAdded // Event containing the contract specifics and raw log

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
func (it *AllPurposePauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposePauserAdded)
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
		it.Event = new(AllPurposePauserAdded)
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
func (it *AllPurposePauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposePauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposePauserAdded represents a PauserAdded event raised by the AllPurpose contract.
type AllPurposePauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_AllPurpose *AllPurposeFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*AllPurposePauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurpose.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &AllPurposePauserAddedIterator{contract: _AllPurpose.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_AllPurpose *AllPurposeFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *AllPurposePauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurpose.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposePauserAdded)
				if err := _AllPurpose.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// AllPurposePauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the AllPurpose contract.
type AllPurposePauserRemovedIterator struct {
	Event *AllPurposePauserRemoved // Event containing the contract specifics and raw log

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
func (it *AllPurposePauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposePauserRemoved)
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
		it.Event = new(AllPurposePauserRemoved)
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
func (it *AllPurposePauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposePauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposePauserRemoved represents a PauserRemoved event raised by the AllPurpose contract.
type AllPurposePauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_AllPurpose *AllPurposeFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*AllPurposePauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurpose.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AllPurposePauserRemovedIterator{contract: _AllPurpose.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_AllPurpose *AllPurposeFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *AllPurposePauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AllPurpose.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposePauserRemoved)
				if err := _AllPurpose.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// AllPurposeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AllPurpose contract.
type AllPurposeTransferIterator struct {
	Event *AllPurposeTransfer // Event containing the contract specifics and raw log

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
func (it *AllPurposeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeTransfer)
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
		it.Event = new(AllPurposeTransfer)
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
func (it *AllPurposeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeTransfer represents a Transfer event raised by the AllPurpose contract.
type AllPurposeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AllPurpose *AllPurposeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AllPurposeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AllPurpose.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AllPurposeTransferIterator{contract: _AllPurpose.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AllPurpose *AllPurposeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AllPurposeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AllPurpose.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeTransfer)
				if err := _AllPurpose.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// AllPurposeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AllPurpose contract.
type AllPurposeUnpausedIterator struct {
	Event *AllPurposeUnpaused // Event containing the contract specifics and raw log

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
func (it *AllPurposeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllPurposeUnpaused)
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
		it.Event = new(AllPurposeUnpaused)
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
func (it *AllPurposeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllPurposeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllPurposeUnpaused represents a Unpaused event raised by the AllPurpose contract.
type AllPurposeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AllPurpose *AllPurposeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AllPurposeUnpausedIterator, error) {

	logs, sub, err := _AllPurpose.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AllPurposeUnpausedIterator{contract: _AllPurpose.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AllPurpose *AllPurposeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AllPurposeUnpaused) (event.Subscription, error) {

	logs, sub, err := _AllPurpose.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllPurposeUnpaused)
				if err := _AllPurpose.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
