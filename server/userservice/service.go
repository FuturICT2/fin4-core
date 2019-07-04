package userservice

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/dbservice"
	"github.com/asaskevich/govalidator"
)

//Service defines service type
type Service struct {
	*dbservice.DB
}

//NewService factory for Service
func NewService(db *dbservice.DB) *Service {
	return &Service{db}
}

//TODO do we need this? maybe get rid of it
// Validate validates email and password for signup
func (db *Service) Validate(email string, password string) error {
	if !govalidator.IsEmail(email) {
		return datatype.ErrEmailInvalid
	}
	// if db.IsEmailRegistered(email) {
	// 	return datatype.ErrEmailExists
	// }
	if err := validatePassword(password, email); err != nil {
		return err
	}
	return nil
}
