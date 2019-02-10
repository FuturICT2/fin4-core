package assetservice

import (
	"github.com/FuturICT2/fin4-core/server/apperrors"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/helpers"
)

// GetAssetBlocks finds all user's favored assits
func (db *Service) GetAssetBlocks(
	assetID datatype.ID,
) ([]datatype.Block, error) {
	result := []datatype.Block{}
	rows, err := db.Query(
		`SELECT
      block.id,
      block.assetId,
      block.userId,
			user.username,
      block.text,
			block.videoID,
			block.status,
			block.favoritesCounter,
			block.ethereumTransactionAddress,
      block.createdAt
    FROM asset_block block
    LEFT JOIN user user ON block.userId = user.id
    WHERE block.assetId = ? AND block.status <> ?
    ORDER BY id DESC`,
		assetID,
		datatype.BlockRejected,
	)
	if err != nil {
		apperrors.Critical("assetservice:GetAssetBlock:1", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c datatype.Block
		err := rows.Scan(
			&c.ID,
			&c.AssetID,
			&c.UserID,
			&c.UserName,
			&c.Text,
			&c.YtVideoID,
			&c.Status,
			&c.FavoritesCounter,
			&c.EthereumTransactionAddress,
			&c.CreatedAt,
		)
		c.CreatedAtHuman = helpers.DateToHuman(c.CreatedAt)
		if err != nil {
			apperrors.Critical("assetservice:GetAssetBlock:2", err)
			return nil, err
		}
		c.Images, err = db.GetAssetBlockImages(c.ID)
		if err != nil {
			apperrors.Critical("assetservice:GetAssetBlock:3", err)
			return nil, err
		}
		result = append(result, c)
	}
	if err := rows.Err(); err != nil {
		apperrors.Critical("assetservice:GetAssetBlock:4", err)
		return nil, err
	}
	return result, nil
}
