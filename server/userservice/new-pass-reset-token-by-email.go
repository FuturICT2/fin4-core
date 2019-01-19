package userservice

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

// NewPassResetTokenByEmail creates a new reset password token for given email
func (db *Service) NewPassResetTokenByEmail(email string) (*datatype.UserResetToken, error) {
	email = sanitizeEmail(email)
	if !govalidator.IsEmail(email) {
		return nil, datatype.ErrEmailInvalid
	}
	user, err := db.FindByEmail(email)
	if err != nil || user == nil {
		return nil, datatype.ErrEmailDoesntExist
	}

	tokenEntry := datatype.NewUserResetToken(user.ID)
	res, err := db.Exec(
		"INSERT INTO user_password_reset SET userId=?, token=?, createdAt=?",
		tokenEntry.UserID,
		tokenEntry.Token,
		time.Now(),
	)
	if err != nil {
		return nil, datatype.ErrServerError
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, datatype.ErrServerError
	}
	tokenEntry.ID = datatype.ID(lastID)
	return tokenEntry, nil
}
