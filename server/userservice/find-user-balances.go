package userservice

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/sirupsen/logrus"
)

// FindUserBalances finds all balances belonging to given user
func (db *Service) FindUserBalances(id datatype.ID) ([]datatype.Balance, error) {
	result := []datatype.Balance{}
	rows, err := db.Query(`SELECT
			b.tokenId,
			b.balance,
			b.reserved,
			c.name,
			c.symbol,
			c.logoFile,
			m.blockchainAddress
		FROM user_balance b
		LEFT JOIN
			token c ON c.id = b.tokenId
		WHERE b.userId=?`,
		id,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": id},
		).Error("user:FindUserBalances:1")
		return result, datatype.ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := datatype.Balance{UserID: id}
		err = rows.Scan(
			&entry.TokenID,
			&entry.Balance,
			&entry.Reserved,
			&entry.TokenName,
			&entry.TokenSymbol,
			&entry.LogoFile,
			&entry.BlockchainAddress,
		)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{"e": err.Error(), "userID": id},
			).Error("user:FindUserBalances:2")
			return result, datatype.ErrServerError
		}
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": id},
		).Error("user:FindUserBalances:3")
		return result, datatype.ErrServerError
	}
	return result, nil
}
