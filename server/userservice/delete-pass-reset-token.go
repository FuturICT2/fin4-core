package userservice

import "github.com/FuturICT2/fin4-core/server/datatype"

// DeletePassResetToken deletes a pass-reset token
func (db *Service) DeletePassResetToken(userID datatype.ID, token string) bool {
	_, err := db.Exec("DELETE FROM user_password_reset WHERE userId = ? AND token = ?", userID, token)
	return err == nil
}
