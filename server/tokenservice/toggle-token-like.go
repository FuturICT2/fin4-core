package tokenservice

import (
	"log"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

// ToggleTokenLike insert token
func (db *Service) ToggleTokenLike(
	userID datatype.ID,
	tokenID datatype.ID,
) error {
	state := db.DidUserLike(userID, tokenID)
	log.Println("((((()))))userID=", userID, "... tokenID", tokenID, " status = ", state)
	if state == true {
		_, err := db.Exec(
			`delete from token_like where
      userId = ? and  tokenId = ?`,
			userID,
			tokenID,
		)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("tokenservice:ToggleTokenLike:e1")
			return datatype.ErrServerError
		}
		return nil
	} else {
		_, err := db.Exec(
			`INSERT INTO token_like SET
      userId = ?,
      tokenId = ?`,
			userID,
			tokenID,
		)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("tokenservice:ToggleTokenLike:e2")
			return datatype.ErrServerError
		}
		return nil
	}
}
