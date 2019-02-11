package assethandlers

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
)

func toAssetsResponse(entries []datatype.Asset) []interface{} {
	res := []interface{}{}
	for _, entry := range entries {
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
		})
	}
	return res
}
