package userservice

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (db *Service) createBalance(userID datatype.ID, assetID datatype.ID) bool {
	res, err := db.Exec(`INSERT INTO user_balance SET
			userId = ?, assetId = ?, balance = 0, reserved = 0 `,
		userID, assetID,
	)
	if err != nil {
		return false
	}
	_, err = res.LastInsertId()
	return err == nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Couldn't save password: " + err.Error())
	}
	return string(hash[:]), nil
}

func sanitizeEmail(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)
	return email
}

// ValidatePassword validates password for a given email
func validatePassword(password string, forEmail string) error {
	if len(password) < 6 {
		return datatype.ErrPasswordTooShort
	}
	if password == forEmail {
		//return datatype.ErrPasswordSameAsEmail
	}
	return nil
}

func getUserCols() string {
	return "id, email, name, ethereumAddress, password, salt, createdAt, updatedAt"
}

func scanUser(row *sql.Row) (*datatype.User, error) {
	var user datatype.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.EthereumAddress,
		&user.Password,
		&user.Salt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("Not found")
	} else if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:scanUser:1")
		return nil, err
	}
	return &user, nil
}
