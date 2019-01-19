package userservice

import (
	"fmt"
	"log"

	"github.com/FuturICT2/fin4-core/server/datatype"
)

const tokenMaxTTLInHours = 48

// ConfirmPassResetToken validates a pass-reset token and update user's password
func (db *Service) ConfirmPassResetToken(
	userID datatype.ID,
	token string,
	newPassword string,
) error {
	var count int64
	user, err := db.FindByID(userID)
	if err != nil || user == nil {
		return datatype.ErrBadRequest
	}
	err = db.QueryRow(
		"SELECT count(*) FROM user_password_reset WHERE userId = ? AND token = ? AND createdAt > (now() - INTERVAL ? DAY)",
		userID,
		token,
		tokenMaxTTLInHours,
	).Scan(&count)
	if err != nil {
		return datatype.ErrServerError
	}
	if count == 0 {
		return datatype.ErrBadRequest
	}

	if err = validatePassword(newPassword, user.Email); err != nil {
		return err
	}

	user.Password, err = hashPassword(newPassword + user.Salt)
	if err != nil {
		return datatype.ErrServerError
	}
	_, err = db.Exec("UPDATE user SET password = ? WHERE id = ?", user.Password, user.ID)
	if err != nil {
		return datatype.ErrServerError
	}
	db.Exec("DELETE from user_password_reset WHERE userId = ? AND token = ?", userID, token)
	log.Println(fmt.Sprintf("DELETE from user_password_reset WHERE userId = %d AND token = '%s'", userID, token))
	return nil
}
