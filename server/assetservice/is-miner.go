package assetservice

import (
	"database/sql"

	"github.com/FuturICT2/fin4-core/server/apperrors"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

//IsMiner checks whether a user is a miner of a token or not
func (db *Service) IsMiner(userID datatype.ID, assetID datatype.ID) bool {
	var id datatype.ID
	err := db.QueryRow(
		`SELECT
			assetId as counter
		FROM asset_block
		WHERE assetId=? AND userId=? and status=? LIMIT 1`,
		assetID,
		userID,
		datatype.BlockAccepted,
	).Scan(&id)
	if err != nil {
		apperrors.Critical("assetservice:IsMiner:1", err)
		return false
	}
	if err == sql.ErrNoRows {
		return false
	}
	return true
}
