package userservice

import (
	"github.com/FuturICT2/fin4-core/server/apperrors"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

// UpdateProfileImage changes user profile image url
func (db *Service) UpdateProfileImage(
	user *datatype.User, imageURL string,
) error {
	res, err := db.Exec("UPDATE user SET profileImageUrl = ? WHERE id = ?", imageURL, user.ID)
	if err != nil {
		apperrors.Critical("userservice:update-prfoile-image:1", err)
		return datatype.ErrServerError
	}
	if affected, err := res.RowsAffected(); err != nil || affected == 0 {
		apperrors.Critical("userservice:update-prfoile-image:2", err)
		return datatype.ErrServerError
	}
	return nil
}
