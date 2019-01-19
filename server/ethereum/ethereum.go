package ethereum

import (
	"errors"
	"math/big"

	"github.com/FuturICT2/fin4-core/server/env"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lytics/logrus"
)

// Ethereum ethereum struct to implement crypto interface
type Ethereum struct {
	rpc      *ethclient.Client
	sim      *backends.SimulatedBackend
	auth     *bind.TransactOpts
	keystore *keystore.KeyStore
}

// MustNewEthereum create new Ethereum interface, panic if no connection
func MustNewEthereum() *Ethereum {
	conn, err := ethclient.Dial("https://rinkeby.infura.io/")
	if err != nil {
		logrus.Fatal("Failed to connect to the Ethereum client: %v", err)
	}
	// server key
	rawKey := env.MustGetenv("ETH_KEY_RAW")
	rawKeyECDSA, err := crypto.HexToECDSA(rawKey)
	if err != nil {
		logrus.Fatal("Something wrong with server private key.", err)
	}
	ks := keystore.NewKeyStore(
		env.MustGetenv("ETH_KEY_STORE_DIR"),
		keystore.LightScryptN,
		keystore.LightScryptP)
	ks.ImportECDSA(rawKeyECDSA, "passphrase")
	// Create an authorized transactor
	auth := bind.NewKeyedTransactor(rawKeyECDSA)
	if err != nil {
		logrus.Fatal("Failed to create transactor: %v", err)
	}
	// Setup blockchain simmulator
	gAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: big.NewInt(10000000000)},
	}
	sim := backends.NewSimulatedBackend(gAlloc, 21000)
	return &Ethereum{
		rpc:      conn,
		sim:      sim,
		auth:     auth,
		keystore: ks,
	}
}

// CreateNewAddress returns best blocknumber in the blockchain
func (b *Ethereum) CreateNewAddress() (string, error) {
	acc, err := b.keystore.NewAccount("demo")
	return acc.Address.String(), err
}

// DeployMintable deployes new Mintable token to Ethereum from server account
func (b *Ethereum) DeployMintable(
	name_ string,
	symbol_ string,
	decimals_ uint8,
	minter common.Address,
) (common.Address, *types.Transaction, error) {
	address, tx, _, err := DeployMintable(
		b.auth,
		b.sim,
		name_,
		symbol_,
		decimals_,
		minter,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("ethereum:DeployMintable:e1")
		return address, nil, errors.New("Error deploying token contract to Ethereum")
	}
	return address, tx, nil
}

// DeployMintable deployes new Mintable token to Ethereum from server account
func (b *Ethereum) Mint(
	tokenAddress common.Address,
	toAddress common.Address,
	amount int64,
) (*types.Transaction, error) {
	mintable, err := NewMintable(tokenAddress, b.sim)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("ethereum:Mint:1")
		return nil, err
	}
	// @TODO change b.auth to the address of the user who is minting the new
	// tokens i.e claim approver
	txAddress, err := mintable.Mint(b.auth, toAddress, big.NewInt(amount))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("ethereum:Mint:2")
		return nil, err
	}
	return txAddress, nil
}
