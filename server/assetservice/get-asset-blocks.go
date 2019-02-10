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
      b.id,
      b.userId,
      b.text,
      b.status,
			b.videoID,
			b.createdAt,
      u.username
    FROM asset_block b
    LEFT JOIN user u ON b.userId = u.id
    WHERE b.assetId = ? AND b.status <> ?
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
			&c.UserID,
			&c.Text,
			&c.Status,
			&c.YtVideoID,
			&c.CreatedAt,
			&c.UserName,
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
