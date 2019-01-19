package tokenservice

import (
	"fmt"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

// FindAll finds all assets
func (db *Service) FindAll(userID datatype.ID) ([]datatype.Token, error) {
	result := []datatype.Token{}
	rows, err := db.Query(
		fmt.Sprintf(`
			SELECT
				%s,
				u.name
			FROM token t
			LEFT JOIN
				user u ON u.id = t.creatorId
			ORDER BY id DESC`, getTokenCols()),
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("models:FindTokens:e1")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c datatype.Token
		rows.Scan(
			&c.ID,
			&c.CreatorID,
			&c.Name,
			&c.Symbol,
			&c.TotalSupply,
			&c.Purpose,
			&c.BlockchainAddress,
			&c.TxAddress,
			&c.CreatorName,
		)
		c.FavouriteCount = db.CountTokenLikes(c.ID)
		c.DidUserLike = db.DidUserLike(userID, c.ID)
		claims, err := db.GetTokenClaims(userID, c.ID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("tokenservice:findall:e1.2")
			return nil, err
		}
		c.Claims = claims
		result = append(result, c)
	}
	if err := rows.Err(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("models:FindTokens:e2")
		return nil, err
	}
	return result, nil
}
