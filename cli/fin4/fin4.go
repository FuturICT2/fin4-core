package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "tokens"
	app.Usage = "to help create new fin4 tokens"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "deploy new token to the ethereum blockchain",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "name, n"},
				cli.StringFlag{Name: "totalSupply, t"},
				cli.StringFlag{Name: "symbol, s"},
				cli.IntFlag{Name: "decimals, d"},
			},
			Action: newTokenHandler,
		},
		{
			Name:    "info",
			Aliases: []string{"i"},
			Usage:   "retreive informaion about tokens",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "token"},
				cli.StringFlag{Name: "tx"},
			},
			Action: tokenInfoHandler,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func newTokenHandler(c *cli.Context) error {
	name := c.String("name")
	symbol := c.String("symbol")
	totalSupply := new(big.Int)
	totalSupply, ok := totalSupply.SetString(c.String("totalSupply"), 10)
	if !ok {
		fmt.Println("Please pass valid total supply")
		return nil
	}
	decimals := c.Int("decimals")
	deployNewToken(totalSupply, name, uint8(decimals), symbol)
	return nil
}

func tokenInfoHandler(c *cli.Context) error {
	address := c.String("token")
	printTokenName(address)
	return nil
}

func deployNewToken(
	initialSupply *big.Int,
	tokenName string,
	decimals uint8,
	tokenSymbol string,
) (common.Address, *types.Transaction, *Token) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(key)
	gAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: big.NewInt(10000000000)},
	}
	sim := backends.NewSimulatedBackend(gAlloc, 10393939)

	// Create an IPC based RPC connection to a remote node and an authorized transactor
	// priv := "2cc40d7ad120a03b1806a83f82422a0b9753e76dbe0da0099a3202aa61bb6217"
	// key, err := crypto.HexToECDSA(priv)
	// if err != nil {
	// 	log.Fatalf("Failed to import key: %v", err)
	// }
	// conn, err := ethclient.Dial("http://localhost:8545")
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	// }
	// auth := bind.NewKeyedTransactor(key)

	// Deploy new token
	address, tx, token, err := DeployToken(auth, sim, initialSupply, tokenName, decimals, tokenSymbol)
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Token contract address: 0x%x\n", address)
	fmt.Printf("Tx address: 0x%x\n\n", tx.Hash())
	return address, tx, token
}

func printTokenName(address string) {
	conn, err := ethclient.Dial("https://rinkeby.infura.io/")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	log.Println("Connected to Ethereum")
	log.Printf("Fetching contract info from address : %s\n", address)
	token, err := NewToken(common.HexToAddress(address), conn)
	if err != nil {
		log.Fatalf("cli.tokens.tokens.NewToken.e1: %v", err)
	}
	name, err := token.Name(nil)
	if err != nil {
		log.Fatalf("cli.tokens.tokens.NewToken.e2: %v", err)
	}
	log.Println("Token name is >>> ", name)

}
