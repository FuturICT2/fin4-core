package assetservice

import (
	"database/sql"

	"github.com/FuturICT2/fin4-core/server/apperrors"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/decimaldt"
)

// FindUserBalance finds user's balance of given asset
func (db *Service) FindUserBalance(
	userID datatype.ID,
	assetID datatype.ID,
) (
	decimaldt.Decimal,
	decimaldt.Decimal,
	error,
) {
	var balance decimaldt.Decimal
	var reservedBalace decimaldt.Decimal
	err := db.QueryRow(
		`SELECT balance, reserved
			FROM asset_user_balance
			WHERE userId = ? AND assetId = ?`,
		userID,
		assetID,
	).Scan(&balance, &reservedBalace)
	if err == sql.ErrNoRows {
		return decimaldt.NewFromFloat(0.0), decimaldt.NewFromFloat(0.0), nil
	} else if err != nil {
		apperrors.Critical("user:FindUserBalance:1", err)
		return decimaldt.NewFromFloat(0.0), decimaldt.NewFromFloat(0.0), datatype.ErrServerError
	}
	return balance, reservedBalace, err
}
