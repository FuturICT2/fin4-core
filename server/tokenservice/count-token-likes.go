package tokenservice

import "github.com/FuturICT2/fin4-core/server/datatype"

// CountTokenLikes finds all assets
func (db *Service) CountTokenLikes(tokenID datatype.ID) int {
	var count int
	db.QueryRow(
		`SELECT
			count(*)
		FROM token_like
		WHERE tokenId = ? `,
		tokenID,
	).Scan(&count)
	return count
}
