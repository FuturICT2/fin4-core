package tokenservice

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

func (db *Service) GetTokenClaims(userID datatype.ID, tokenID datatype.ID) ([]datatype.Claim, error) {
	result := []datatype.Claim{}
	rows, err := db.Query(`SELECT
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
		WHERE c.tokenId=?
		ORDER BY id DESC`,
		tokenID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("GetTokenClaims:1")
		return result, datatype.ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := datatype.Claim{}
		err = rows.Scan(
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
				logrus.Fields{"error": err.Error()},
			).Error("GetTokenClaims:2")
			return result, datatype.ErrServerError
		}
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("GetTokenClaims:3")
		return result, datatype.ErrServerError
	}
	return result, nil
}
