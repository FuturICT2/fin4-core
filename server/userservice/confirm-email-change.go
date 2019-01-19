package userservice

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

// ConfirmUserEmailChange confirms email change and changes user's email if successful
func (db *Service) ConfirmUserEmailChange(userID datatype.ID, token string) bool {
	conf := datatype.EmailChangeConfirmation{}
	err := db.QueryRow(`SELECT userId, email, token FROM user_change_email_confirm
		WHERE userId = ? AND createdAt > (now() - INTERVAL ? DAY)`,
		userID, 1,
	).Scan(&conf.UserID, &conf.Email, &conf.Token)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": userID, "token": token},
		).Error("user:ConfirmUserEmailChange:1")
		return false
	}
	if token != conf.Token {
		return false
	}
	defer db.Exec("DELETE FROM user_change_email_confirm WHERE userId = ?", userID)
	// reinsure that the email is still available
	if db.IsEmailRegistered(conf.Email) {
		return false
	}
	_, err = db.Exec("UPDATE user SET email = ? WHERE id = ?",
		conf.Email,
		userID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": userID, "token": token},
		).Error("user:ConfirmUserEmailChange:2")
		return false
	}
	return true
}
