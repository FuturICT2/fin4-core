package tokens

import "github.com/hunterlong/tokenbalance"

// EthToken
type EthToken struct {
	ContractAddress string
}

/*
Fin4 ERC Rinkbey
contract 0x3A26aF8d142528A186BfaeDE37a83bf0D4C9768b
tx 0xcb64c247d6002258f4b5b344459d53deef4dec3b50b6916db28e0d3629abcb96

FinDevCoin
contract 0xF9fcCedcBcd0c00E25A330aD63cc5d52c5eF9Ca0
tx 0x05a6b46db6a8c3799fa95d2db1fcd52874befe31d13f5490061fa527bb776244
*/

const (
	fin4Address   = "0x3A26aF8d142528A186BfaeDE37a83bf0D4C9768b"
	finDevAddress = "0xF9fcCedcBcd0c00E25A330aD63cc5d52c5eF9Ca0"
)

func (E *EthToken) getBalance(address string) int {

}

func main() {
	// initLogger()
	// ethereum := ethereum.MustNewEthereum()
	// routesEnv := routes.Env{
	// 	// We can add general types like database connection to the ENV when needed
	// 	Ethereum: ethereum,
	// }
	// routesEnv.StartRouter().Run(":" + getPort())

	configs := &tokenbalance.Config{
		GethLocation: "https://rinkeby.infura.io/",
		Logs:         true,
	}
	configs.Connect()

	// insert a Token Contract address and Wallet address
	contract := "0x3A26aF8d142528A186BfaeDE37a83bf0D4C9768b"
	wallet := "0x81bec362eb7b7cd838b155eb8845cbf1f2b50e0f"

	// query the blockchain and wallet details
	token, _ := tokenbalance.New(contract, wallet)

	// Token Balance will respond back useful things
	// token.BalanceString() // "600000.0"
	// token.ETHString()     // "1.020095885777777767"
	// token.Name            // "OmiseGO"
	// token.Symbol          // "OMG"
	// token.Decimals        // 18
	// token.Balance         // big.Int() (token balance)
	// token.ETH             // big.Int() (ether balance)
	token.
		log.Println(token.Balance)
}
