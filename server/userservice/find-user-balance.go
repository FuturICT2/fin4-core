package userservice

import (
	"database/sql"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/decimaldt"
	"github.com/sirupsen/logrus"
)

// FindUserBalance finds user's balance of given currecncy
func (db *Service) FindUserBalance(
	userID datatype.ID,
	currecncyID datatype.ID,
) (
	decimaldt.Decimal,
	decimaldt.Decimal,
	error,
) {
	var balance decimaldt.Decimal
	var reservedBalace decimaldt.Decimal
	err := db.QueryRow(
		`SELECT balance, reserved
			FROM user_balance
			WHERE userId = ? AND assetId = ?`,
		userID,
		currecncyID,
	).Scan(&balance, &reservedBalace)
	if err == sql.ErrNoRows {
		if db.createBalance(userID, currecncyID) {
			return decimaldt.NewFromFloat(0.0), decimaldt.NewFromFloat(0.0), nil
		}
	} else if err != nil {
		logrus.WithFields(
			logrus.Fields{
				"e":           err.Error(),
				"userID":      userID,
				"currecncyID": currecncyID,
			},
		).Error("user:FindUserBalance:1")
		return decimaldt.NewFromFloat(0.0), decimaldt.NewFromFloat(0.0), datatype.ErrServerError
	}
	return balance, reservedBalace, err
}
