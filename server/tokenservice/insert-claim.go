package tokenservice

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

// NewClaim insert a mining claim
func (db *Service) InsertClaim(
	userID datatype.ID,
	proposal string,
	actionID datatype.ID,
	logoPath string,
) error {
	res, err := db.Exec(
		`INSERT INTO claim SET
			tokenId = ?,
			userId = ?,
			text = ?,
			isApproved = 0,
			logoFile = ?`,
		actionID,
		userID,
		proposal,
		logoPath,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:NewActionClaim:0")
		return datatype.ErrServerError
	}
	_, err = res.LastInsertId()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:NewActionClaim:1")
		return datatype.ErrServerError
	}
	return nil
}
