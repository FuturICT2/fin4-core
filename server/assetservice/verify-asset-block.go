package assetservice

import (
	"database/sql"
	"errors"

	"github.com/FuturICT2/fin4-core/server/apperrors"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/decimaldt"
)

// VerifyAssetBlock insert asset block
func (db *Service) VerifyAssetBlock(
	sc *datatype.ServiceContainer,
	user *datatype.User,
	blockID datatype.ID,
	status int,
) error {
	// make sure block exists and the logged in user is the oracle of the related
	// asset
	var assetID datatype.ID
	var doerID datatype.ID
	err := db.QueryRow(`
    SELECT
      b.assetId,
      b.userId
    FROM asset_block b
    LEFT JOIN asset ta ON ta.id=b.assetId
    WHERE b.id=? AND ta.creatorId=? AND b.status = ?`,
		blockID,
		user.ID,
		datatype.BlockUnverified,
	).Scan(&assetID, &doerID)
	if err == sql.ErrNoRows {
		return errors.New("Not authorized")
	}
	if err != nil {
		apperrors.Critical("assetservice:VerifyAssetBlock:1", err)
		return datatype.ErrServerError
	}

	if status == datatype.BlockAccepted {
		return db.acceptAssetBlock(sc, blockID, assetID, doerID)
	}
	return db.rejectAssetBlock(user, blockID)
}

// VerifyAssetBlock insert asset block
func (db *Service) acceptAssetBlock(
	sc *datatype.ServiceContainer,
	blockID datatype.ID,
	assetID datatype.ID,
	doerID datatype.ID,
) error {
	isMiner := db.IsMiner(doerID, assetID)
	// increase total supply, add balance to block doer
	// Start db transaction
	tx, err := db.Begin()
	if err != nil {
		apperrors.Critical("assetservice:acceptAssetBlock:1", err)
		return err
	}
	defer tx.Rollback()
	minerIncrement := 1
	if isMiner {
		minerIncrement = 0
	}
	_, err = tx.Exec(
		`UPDATE asset SET
	   	totalSupply = totalSupply + 1,
			minersCounter = minersCounter + ?
	   WHERE id=?`,
		minerIncrement,
		assetID,
	)
	if err != nil {
		apperrors.Critical("assetservice:acceptAssetBlock:2", err)
		return datatype.ErrServerError
	}
	err = db.DepositBalance(doerID, assetID, decimaldt.NewFromFloat(1))
	if err != nil {
		apperrors.Critical("assetservice:acceptAssetBlock:3", err)
		return datatype.ErrServerError
	}
	_, err = tx.Exec(`UPDATE asset_block SET status = ? WHERE id = ?`,
		datatype.BlockAccepted,
		blockID)
	if err != nil {
		apperrors.Critical("assetservice:acceptAssetBlock:4", err)
		return datatype.ErrServerError
	}
	/*** Commit db transaction ***/
	err = tx.Commit()
	if err != nil {
		apperrors.Critical("assetservice:acceptAssetBlock:5", err)
		return datatype.ErrServerError
	}
	return nil
}

// rejectAssetBlock rejects asset block
func (db *Service) rejectAssetBlock(
	user *datatype.User,
	blockID datatype.ID,
) error {
	_, err := db.Exec(
		`UPDATE asset_block SET status = ? WHERE id = ?`,
		datatype.BlockRejected,
		blockID,
	)
	if err != nil {
		apperrors.Critical("assetservice:rejectAssetBlock:1", err)
		return datatype.ErrServerError
	}
	return nil
}
