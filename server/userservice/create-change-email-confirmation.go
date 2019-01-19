package userservice

import (
	"time"

	"github.com/FuturICT2/fin4-core/server/appstrings"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

// CreateChangeEmailConfirmation creates a new email change confirmation db entry
func (db *Service) CreateChangeEmailConfirmation(
	user *datatype.User,
	email string,
	password string,
) (*datatype.EmailChangeConfirmation, error) {
	if !user.IsPassword(password) {
		return nil, datatype.ErrInvalidCurrentPassword
	}
	if err := db.Validate(email, password); err != nil {
		return nil, err
	}
	token := appstrings.GetRandomString(64)
	res, err := db.Exec(`INSERT INTO user_change_email_confirm
		SET userId = ?, email = ?, token = ?, createdAt = ?
		ON DUPLICATE KEY UPDATE email = ?, token = ?, createdAt = ?`,
		user.ID, email, token, time.Now(), email, token, time.Now(),
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": user.ID, "email": email},
		).Error("user:CreateChangeEmailConfirmation:1")
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": user.ID, "email": email},
		).Error("user:CreateChangeEmailConfirmation:2")
		return nil, err
	}
	return &datatype.EmailChangeConfirmation{
		UserID: datatype.ID(id),
		Email:  email,
		Token:  token,
	}, nil
}
