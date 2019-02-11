package assetservice

import (
	"errors"
	"fmt"

	"github.com/FuturICT2/fin4-core/server/apperrors"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/decimaldt"
)

// GetAssetMiners finds all user's favored assits
func (db *Service) GetAssetMiners(
	assetID datatype.ID,
) ([]datatype.Miner, error) {
	asset, err := db.FindByID(assetID)
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, errors.New("Asset doesn't exist")
	}
	result := []datatype.Miner{}
	rows, err := db.Query(`SELECT
			c.userId,
			u.name
		FROM asset_block c
		LEFT JOIN
			user u ON u.id = c.userId
		WHERE c.assetId=? AND c.status=?`,
		assetID,
		datatype.BlockAccepted,
	)
	if err != nil {
		apperrors.Critical("assetservice:GetAssetMiners:1", err)
		return nil, datatype.ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := datatype.Miner{}
		err = rows.Scan(
			&entry.ID,
			&entry.UserName,
		)
		if err != nil {
			apperrors.Critical("assetservice:GetAssetMiners:2", err)
			return result, datatype.ErrServerError
		}
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		apperrors.Critical("assetservice:GetAssetMiners:3", err)
		return nil, datatype.ErrServerError
	}
	minersMap := map[int]datatype.Miner{}
	one, _ := decimaldt.NewFromString("1.0")
	// TODO make more clean using SQL
	for _, miner := range result {
		key := int(miner.ID)
		m := minersMap[key]
		if m.UserName == "" {
			miner.Mined = miner.Mined.Add(one)
			minersMap[key] = miner
		} else {
			newMiner := minersMap[key]
			newMiner.Mined = newMiner.Mined.Add(one)
			minersMap[key] = newMiner
		}
	}
	clean := []datatype.Miner{}
	for _, v := range minersMap {
		minedFloat, _ := v.Mined.Float64()
		percentage := minedFloat / float64(asset.Supply)
		v.MiningPercentage = fmt.Sprintf("%f", percentage*100.0)
		clean = append(clean, v)
	}
	return clean, nil
}
