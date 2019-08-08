package assethandlers

import (
	"time"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/helpers"
)

func toAssetsResponse(entries []datatype.Asset) []interface{} {
	res := []interface{}{}
	for _, entry := range entries {
		_, _, _, _, min, sec := helpers.Diff(entry.LastOraclePing, time.Now())
		if min < 1 && sec > 10 {
			entry.IsConnected = false
		} else {
			entry.IsConnected = true
		}
		res = append(res, struct {
			ID                         datatype.ID
			Symbol                     string
			Name                       string
			CreatorID                  datatype.ID
			CreatorName                string
			Description                string
			Supply                     int64
			MinersCounter              int
			FavoritesCounter           int
			DidUserLike                bool
			EthereumAddress            string
			EthereumTransactionAddress string
			LastOraclePingHuman        string
			IsConnected                bool
			OracleType                 int
			IsUserOracle               bool
			AccessToken                string
		}{
			ID:                         entry.ID,
			Symbol:                     entry.Symbol,
			Name:                       entry.Name,
			CreatorID:                  entry.CreatorID,
			CreatorName:                entry.CreatorName,
			Description:                entry.Description,
			Supply:                     entry.Supply,
			MinersCounter:              entry.MinersCounter,
			FavoritesCounter:           entry.FavoritesCounter,
			DidUserLike:                entry.DidUserLike,
			EthereumAddress:            entry.EthereumAddress,
			EthereumTransactionAddress: entry.EthereumTransactionAddress,
			LastOraclePingHuman:        entry.LastOraclePingHuman,
			IsConnected:                entry.IsConnected,
			OracleType:                 entry.OracleType,
			AccessToken:                entry.AccessToken,
		})
	}
	return res
}
