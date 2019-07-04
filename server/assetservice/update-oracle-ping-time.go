package assetservice

import (
	"time"

	"github.com/FuturICT2/fin4-core/server/apperrors"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

// UpdateOraclePingTime updates last ping from oracle
func (db *Service) UpdateOraclePingTime(accessToken string) error {
	_, err := db.Exec(
		`UPDATE asset SET
	   	lastOraclePing = ?
	   WHERE accessToken = ?`,
		time.Now(),
		accessToken,
	)
	if err != nil {
		apperrors.Critical("assetservice:update-oracle-ping-time:1", err)
		return datatype.ErrServerError
	}
	return nil
}
