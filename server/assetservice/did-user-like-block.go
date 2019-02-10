package assetservice

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
)

// DidUserLikeBlock checks whether user liked a block or not
func (db *Service) DidUserLikeBlock(userID datatype.ID, blockID datatype.ID) bool {
	var count int32
	err := db.QueryRow(
		"SELECT count(*) as counter FROM asset_block_favorites WHERE blockId=? AND userId=?",
		blockID,
		userID,
	).Scan(&count)
	return err == nil && count > 0
}
