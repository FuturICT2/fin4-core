package datatype

import (
	"github.com/FuturICT2/fin4-core/server/decimaldt"
)

//UserService defines user service interface
type UserService interface {
	Register(
		email string,
		name string,
		ethereumAddress string,
		password string,
		agreeToTerms bool,
	) (*User, error)
	Validate(string, string) error
	FindByID(ID) (*User, error)
	FindByEmail(string) (*User, error)
	IsEmailRegistered(string) bool
	FindUserBalances(userID ID) ([]Balance, error)
	FindUserBalance(userID ID, assetID ID) (
		availableBalance decimaldt.Decimal,
		reserved decimaldt.Decimal,
		err error,
	)
	NewPassResetTokenByEmail(email string) (*UserResetToken, error)
	ConfirmPassResetToken(userID ID, token, newPassword string) error
	ChangePassword(*User, string, string) error
	CreateChangeEmailConfirmation(*User, string, string) (*EmailChangeConfirmation, error)
	ConfirmUserEmailChange(ID, string) bool
	GetBalances(userID ID) ([]Balance, error)
	GetPerson(userID ID) (*Person, error)
}
