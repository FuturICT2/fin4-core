package ethereum

import (
	"github.com/onrik/ethrpc"
	"github.com/sirupsen/logrus"
)

// MustNewEthRPC creates a new eth rpc client
func MustNewEthRPC(url string) *ethrpc.EthRPC {
	client := ethrpc.NewEthRPC(url)
	verion, err := client.Web3ClientVersion()
	if err != nil {
		logrus.Fatal("Error connecting to Ethereum: ", err)
		return nil
	}
	logrus.Println("Connected to ethereum --> version: ", verion)
	return client
}
