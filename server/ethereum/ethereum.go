package ethereum

import (
	serverUtil "github.com/FuturICT2/fin4-core/server/util"
	"github.com/onrik/ethrpc"
)

// Ethereum ethereum struct to implement crypto interface
type Ethereum struct {
	rpc *ethrpc.EthRPC
}

// MustNewEthereum create new Ethereum interface, panic if no connection
func MustNewEthereum() *Ethereum {
	net := serverUtil.MustGetenv("ETHEREUM_NET")
	url := serverUtil.MustGetenv(net + "_ETH_HOST")
	return &Ethereum{
		rpc: MustNewEthRPC(url),
	}
}

// GetBlockNumber returns best blocknumber in the blockchain
func (b *Ethereum) GetBlockNumber() (int, error) {
	return b.rpc.EthBlockNumber()
}
