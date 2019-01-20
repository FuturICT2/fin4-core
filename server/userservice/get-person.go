package userservice

import (
	"fmt"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/decimaldt"
	"github.com/lytics/logrus"
)

//GetPerson finds a person with stats
func (db *Service) GetPerson(userID datatype.ID) (*datatype.Person, error) {
	res := db.QueryRow(
		fmt.Sprintf("SELECT %s FROM user WHERE id=?", getUserCols()),
		userID,
	)
	user, err := scanUser(res)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("userservice:GetPerson:1")
		return nil, datatype.ErrServerError
	}
	result := []datatype.Balance{}
	rows, err := db.Query(`SELECT
			t.id,
			IFNULL(b.balance, "0.0"),
			t.name,
			t.symbol,
			t.blockchainAddress
		FROM token t
		INNER JOIN
			user_balance b ON b.tokenId = t.id AND b.userId = ?
		WHERE t.creatorId=?`,
		userID,
		userID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("userservice:GetPerson:2")
		return nil, datatype.ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := datatype.Balance{UserID: userID}
		err = rows.Scan(
			&entry.TokenID,
			&entry.Balance,
			&entry.TokenName,
			&entry.TokenSymbol,
			&entry.BlockchainAddress,
		)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{"error": err.Error()},
			).Error("userservice:GetPerson:3")
			return nil, datatype.ErrServerError
		}
		var count int
		db.QueryRow(`
      SELECT COUNT(tokenId)
      FROM claim
      WHERE tokenId=? AND userId=? AND isApproved=1`,
			entry.TokenID,
			userID,
		).Scan(&count)
		d := decimaldt.New(int64(count), 0)
		entry.Mined = d
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("userservice:GetPerson:4")
		return nil, datatype.ErrServerError
	}
	person := &datatype.Person{}
	person.ID = userID
	person.Email = user.Email
	person.Name = user.Name
	person.MinedTokens = result
	return person, nil
}
