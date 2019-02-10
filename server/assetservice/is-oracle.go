package assetservice

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
)

//IsOracle checks whether email is already registered or not
func (db *Service) IsOracle(userID datatype.ID, assetID datatype.ID) bool {
	var count int32
	err := db.QueryRow(
		"SELECT count(*) as counter FROM asset WHERE id=? AND creatorId=?",
		assetID,
		userID,
	).Scan(&count)
	return err == nil && count > 0
}
