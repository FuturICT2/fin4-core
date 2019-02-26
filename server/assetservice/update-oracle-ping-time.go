package assetservice

import (
	"errors"
	"time"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/kjda/exchange/server/apperrors"
)

// UpdateOraclePingTime updates last ping from oracle
func (db *Service) UpdateOraclePingTime(user *datatype.User) error {
	if user == nil {
		return errors.New("User is missing")
	}
	_, err := db.Exec(
		`UPDATE asset SET
	   	lastOraclePing = ?
	   WHERE creatorId=?`,
		time.Now(),
		user.ID,
	)
	if err != nil {
		apperrors.Critical("assetservice:update-oracle-ping-time:1", err)
		return datatype.ErrServerError
	}
	return nil
}
