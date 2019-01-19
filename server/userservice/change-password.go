package userservice

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

// ChangePassword changes user password
func (db *Service) ChangePassword(user *datatype.User, newPassword string, password string) error {
	err := validatePassword(newPassword, user.Email)
	if err != nil {
		return err
	}
	if !user.IsPassword(password) {
		return datatype.ErrInvalidCurrentPassword
	}
	newPassword, err = hashPassword(newPassword + user.Salt)
	user.Password = newPassword
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": user.ID},
		).Error("user:ChangePassword:1")
		return datatype.ErrServerError
	}
	res, err := db.Exec("UPDATE user SET password = ? WHERE id = ?", user.Password, user.ID)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": user.ID},
		).Error("user:ChangePassword:2")
		return datatype.ErrServerError
	}
	if affected, err := res.RowsAffected(); err != nil || affected == 0 {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": user.ID},
		).Error("user:ChangePassword:3")
		return datatype.ErrServerError
	}
	return nil
}
