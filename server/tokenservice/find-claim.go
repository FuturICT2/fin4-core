package tokenservice

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

// FindClaim fetches a claim
func (db *Service) FindClaim(id datatype.ID) (*datatype.Claim, error) {
	row := db.QueryRow(`SELECT
			c.id,
			c.tokenId,
			c.userId,
			u.name,
			c.text,
			c.logoFile,
			c.isApproved
		FROM claim c
		LEFT JOIN
			user u ON u.id = c.userId
		WHERE c.id=?`,
		id,
	)
	entry := datatype.Claim{}
	err := row.Scan(
		&entry.ID,
		&entry.TokenID,
		&entry.UserID,
		&entry.UserName,
		&entry.Text,
		&entry.LogoFile,
		&entry.IsApproved,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("tokenservie:FindClaim:1")
		return nil, err
	}
	return &entry, nil
}
