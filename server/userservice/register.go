package userservice

import (
	"strings"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

//Register Registers a new user
func (db *Service) Register(
	email,
	name string,
	ethereumAddress string,
	password string,
	agreeToTerms bool,
) (*datatype.User, error) {
	var err error
	email = sanitizeEmail(email)
	password = strings.TrimSpace(password)
	if err = db.Validate(email, password); err != nil {
		return nil, err
	}
	if !agreeToTerms {
		return nil, datatype.ErrAgreeToTerms
	}
	user := datatype.NewUser()
	user.Email = email
	user.Name = name
	user.Password, err = hashPassword(password + user.Salt)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("userservice:register:1")
		return nil, datatype.ErrServerError
	}

	res, err := db.Exec(
		`INSERT INTO user SET
			email = ?,
			name = ?,
			password = ?,
			salt = ?,
			ethereumAddress = ?,
			createdAt = ?,
			updatedAt = ?,
			agreeToTerms = ?,
			isDeleted = ?`,
		user.Email,
		user.Name,
		user.Password,
		user.Salt,
		ethereumAddress,
		user.CreatedAt,
		user.UpdatedAt,
		user.AgreeToTerms,
		user.IsDeleted,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("userservice:register:2")
		return nil, datatype.ErrServerError
	}
	id, err := res.LastInsertId()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("userservice:register:3")
		return nil, datatype.ErrServerError
	}
	return db.FindByID(datatype.ID(id))
}
