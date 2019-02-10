package assethandlers

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
)

func toAssetsResponse(entries []datatype.Asset) []interface{} {
	res := []interface{}{}
	for _, entry := range entries {
		res = append(res, struct {
			ID          datatype.ID
			Symbol      string
			Name        string
			CreatorID   datatype.ID
			CreatorName string
			Description string
			TotalSupply int64
			MinersCount int
		}{
			ID:          entry.ID,
			Symbol:      entry.Symbol,
			Name:        entry.Name,
			CreatorID:   entry.CreatorID,
			CreatorName: entry.CreatorName,
			Description: entry.Description,
			TotalSupply: entry.Supply,
			MinersCount: entry.MinersCounter,
		})
	}
	return res
}
