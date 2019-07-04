package assetservice

import (
	"database/sql"
	"errors"

	"github.com/FuturICT2/fin4-core/server/apperrors"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

// SensorVerifyBlock sensor claim verification
func (db *Service) SensorVerifyBlock(
	sc *datatype.ServiceContainer,
	status int,
	accessToken string,
) error {
	// make sure block exists and the logged in user is the oracle of the related
	// asset
	var assetID datatype.ID
	var doerID datatype.ID
	var blockID datatype.ID
	var creatorID datatype.ID
	err := db.QueryRow(`
    SELECT
      b.id,
      b.assetId,
      b.userId,
      ta.creatorId
    FROM asset_block b
    LEFT JOIN asset ta ON ta.id=b.assetId
    WHERE b.status = ? AND ta.accessToken = ?`,
		datatype.BlockUnverified,
		accessToken,
	).Scan(&blockID, &assetID, &doerID, &creatorID)
	if err == sql.ErrNoRows {
		return errors.New("Not authorized")
	}
	if err != nil {
		apperrors.Critical("assetservice:SensorVerifyBlock:1", err)
		return datatype.ErrServerError
	}

	if status == datatype.BlockAccepted {
		return db.acceptAssetBlock(sc, blockID, assetID, doerID)
	}
	u, _ := sc.UserService.FindByID(creatorID)
	return db.rejectAssetBlock(u, blockID)
}
