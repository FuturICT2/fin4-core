package assetservice

import (
	"database/sql"
	"time"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/helpers"
)

// FindByID finds Asset by ID
func (db *Service) FindByID(id datatype.ID) (*datatype.Asset, error) {
	var c datatype.Asset
	err := db.QueryRow(`
		SELECT
			asset.id,
			asset.name,
			asset.symbol,
			asset.description,
			asset.supply,
			asset.creatorId,
			asset.oracleType,
			user.name,
			asset.minersCounter,
			asset.favoritesCounter,
			asset.ethereumAddress,
			asset.ethereumTransactionAddress,
			asset.createdAt,
			asset.lastOraclePing,
			asset.accessToken
		FROM
			asset asset
		LEFT JOIN user user ON asset.creatorId=user.id
		WHERE asset.id = ?`,
		id,
	).Scan(
		&c.ID,
		&c.Name,
		&c.Symbol,
		&c.Description,
		&c.Supply,
		&c.CreatorID,
		&c.OracleType,
		&c.CreatorName,
		&c.MinersCounter,
		&c.FavoritesCounter,
		&c.EthereumAddress,
		&c.EthereumTransactionAddress,
		&c.CreatedAt,
		&c.LastOraclePing,
		&c.AccessToken,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	c.CreatedAtHuman = helpers.DateToHuman(c.CreatedAt)
	c.LastOraclePingHuman = helpers.DateToHuman(c.LastOraclePing)
	_, _, _, _, min, sec := helpers.Diff(c.LastOraclePing, time.Now())
	if min < 1 && sec < 10 {
		c.IsConnected = true
	} else {
		c.IsConnected = false
	}
	return &c, err
}
