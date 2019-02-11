package assetservice

import (
	"errors"
	"time"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/apperrors"
)

//ToggleFavorite toggles asset as fav/not fav
func (db *Service) ToggleFavorite(user *datatype.User, assetID datatype.ID) error {
	if user == nil {
		return errors.New("user is missing")
	}
	count := 0
	err := db.QueryRow(
		"SELECT count(*) FROM asset WHERE id = ?",
		assetID,
	).Scan(&count)
	if err != nil {
		return datatype.ErrServerError
	}
	if count != 1 {
		return errors.New("Invalid asset")
	}
	err = db.QueryRow(
		"SELECT count(*) FROM asset_favorites WHERE userId = ? AND assetId = ?",
		user.ID,
		assetID,
	).Scan(&count)
	if err != nil {
		return datatype.ErrServerError
	}
	if count > 0 {
		db.Exec(
			"DELETE FROM asset_favorites WHERE userId = ? AND assetId = ?",
			user.ID,
			assetID,
		)
		if err != nil {
			return datatype.ErrServerError
		}
		_, err = db.Exec(`UPDATE asset
			SET favoritesCounter = favoritesCounter - 1
			WHERE id = ? `,
			assetID,
		)
		if err != nil {
			apperrors.Critical("assetservice:toggle-favorite:3", err)
			return datatype.ErrServerError
		}
	} else {
		_, err := db.Exec(
			"INSERT INTO asset_favorites SET userId = ?, assetId = ?, addedAt = ?",
			user.ID,
			assetID,
			time.Now(),
		)
		if err != nil {
			return datatype.ErrServerError
		}
		_, err = db.Exec(`UPDATE asset
			SET favoritesCounter = favoritesCounter + 1
			WHERE id = ? `,
			assetID,
		)
		if err != nil {
			apperrors.Critical("assetservice:toggle-favorite:3", err)
			return datatype.ErrServerError
		}
	}
	return nil
}
