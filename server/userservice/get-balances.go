package userservice

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

// GetBalaces fetches user balaces
func (db *Service) GetBalances(userID datatype.ID) ([]datatype.Balance, error) {
	result := []datatype.Balance{}
	rows, err := db.Query(`SELECT
			b.tokenId,
			b.balance,
			b.reserved,
			t.name,
			t.symbol,
			t.blockchainAddress
		FROM user_balance b
		LEFT JOIN
			token t ON t.id = b.tokenId
		WHERE b.userId=?`,
		userID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("usersercie:GetBalances:1")
		return result, datatype.ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := datatype.Balance{UserID: userID}
		err = rows.Scan(
			&entry.TokenID,
			&entry.Balance,
			&entry.Reserved,
			&entry.TokenName,
			&entry.TokenSymbol,
			&entry.BlockchainAddress,
		)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{"error": err.Error()},
			).Error("usersercie:GetBalances:2")
			return result, datatype.ErrServerError
		}
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("usersercie:GetBalances:3")
		return result, datatype.ErrServerError
	}
	return result, nil
}
