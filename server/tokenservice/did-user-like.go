package tokenservice

import "github.com/FuturICT2/fin4-core/server/datatype"

// DidUserLike returns true or false based on if user liked a token
func (db *Service) DidUserLike(userID datatype.ID, tokenID datatype.ID) bool {
	var count int
	db.QueryRow(
		`SELECT
			count(*)
		FROM token_like
		WHERE tokenId = ? and userId =? `,
		tokenID,
		userID,
	).Scan(&count)
	return count > 0
}
