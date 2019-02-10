package assetservice

import (
	"errors"
	"time"

	"github.com/FuturICT2/fin4-core/server/apperrors"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

// ToggleFavoriteBlock toggles asset block as fav/not fav
func (db *Service) ToggleFavoriteBlock(user *datatype.User, blockID datatype.ID) error {
	if user == nil {
		return errors.New("User is missing")
	}
	count := 0
	err := db.QueryRow(
		"SELECT count(*) FROM asset_block WHERE id = ?",
		blockID,
	).Scan(&count)
	if err != nil {
		return datatype.ErrServerError
	}
	if count != 1 {
		return errors.New("Invalid asset")
	}
	err = db.QueryRow(
		"SELECT count(*) FROM asset_block_favorites WHERE userId = ? AND blockId = ?",
		user.ID,
		blockID,
	).Scan(&count)
	if err != nil {
		apperrors.Critical("assetservice:toggle-favorite-block:1", err)
		return datatype.ErrServerError
	}
	if count > 0 {
		_, err := db.Exec(
			"DELETE FROM asset_block_favorites WHERE userId = ? AND blockId = ?",
			user.ID,
			blockID,
		)
		if err != nil {
			apperrors.Critical("assetservice:toggle-favorite-block:2", err)
			return datatype.ErrServerError
		}
		_, err = db.Exec(`UPDATE asset_block
			SET favoritesCount = favoritesCount - 1
			WHERE id = ? `,
			blockID,
		)
		if err != nil {
			apperrors.Critical("assetservice:toggle-favorite-block:3", err)
			return datatype.ErrServerError
		}
	} else {
		_, err := db.Exec(
			"INSERT INTO asset_block_favorites SET userId = ?, blockId = ?, createdAt = ?",
			user.ID,
			blockID,
			time.Now(),
		)
		if err != nil {
			apperrors.Critical("assetservice:toggle-favorite-block:4", err)
			return datatype.ErrServerError
		}
		_, err = db.Exec(`UPDATE asset_block
			SET favoritesCount = favoritesCount + 1
			WHERE id = ? `,
			blockID,
		)
		if err != nil {
			apperrors.Critical("assetservice:toggle-favorite-block:5", err)
			return datatype.ErrServerError
		}
	}
	return nil
}
