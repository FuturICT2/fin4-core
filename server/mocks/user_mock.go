package models

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/decimaldt"
)

// UserModelMock TradingModel mock
type UserModelMock struct {
}

// Register Register mock
func (db *UserModelMock) Register(string, string, bool) (*datatype.User, error) {
	return &datatype.User{}, nil
}

// Validate Validate mock
func (db *UserModelMock) Validate(string, string) error {
	return nil
}

// FindByID mock
func (db *UserModelMock) FindByID(datatype.ID) (*datatype.User, error) {
	return &datatype.User{}, nil
}

// FindByEmail mock
func (db *UserModelMock) FindByEmail(string) (*datatype.User, error) {
	return &datatype.User{}, nil
}

// IsEmailRegistered mock
func (db *UserModelMock) IsEmailRegistered(string) bool {
	return false
}

// FindUserBalances mock
func (db *UserModelMock) FindUserBalances(datatype.ID) ([]datatype.Balance, error) {
	return []datatype.Balance{}, nil
}

// FindUserBalance mock
func (db *UserModelMock) FindUserBalance(datatype.ID, datatype.ID) (balance decimaldt.Decimal, reservedBalance decimaldt.Decimal, err error) {
	return decimaldt.NewFromFloat(0.0), decimaldt.NewFromFloat(0.0), nil
}

// NewPassResetTokenByEmail mock
func (db *UserModelMock) NewPassResetTokenByEmail(email string) (*datatype.UserResetToken, error) {
	return &datatype.UserResetToken{}, nil
}

// ConfirmPassResetToken mock
func (db *UserModelMock) ConfirmPassResetToken(userID datatype.ID, token, newPassword string) error {
	return nil
}

// ChangePassword mock
func (db *UserModelMock) ChangePassword(*datatype.User, string, string) error {
	return nil
}

// CreateChangeEmailConfirmation mock
func (db *UserModelMock) CreateChangeEmailConfirmation(user *datatype.User, email, password string) (*datatype.EmailChangeConfirmation, error) {
	return nil, nil
}

// ConfirmUserEmailChange mock
func (db *UserModelMock) ConfirmUserEmailChange(userID datatype.ID, token string) bool {
	return false
}

// DepositBalance mock
func (db *UserModelMock) DepositBalance(userID datatype.ID, assetID datatype.ID, amount decimaldt.Decimal) error {
	return nil
}

// ProcessNewDeposits mock
func (db *UserModelMock) ProcessNewDeposits() {

}

// FindDeposits mock
func (db *UserModelMock) FindDeposits(userID datatype.ID, pendingOnly bool) ([]datatype.CryptoDeposit, error) {
	return []datatype.CryptoDeposit{}, nil
}

// GetWithrawals mock
func (db *UserModelMock) GetWithrawals(
	isConfirmedByUser bool,
	isConfirmedByAdmin bool,
	pendingOnly bool,
) ([]datatype.UserWithdrawal, error) {
	return []datatype.UserWithdrawal{}, nil
}

// UpdateWithdrawalStatus mock
func (db *UserModelMock) UpdateWithdrawalStatus(
	withdrawalID datatype.ID,
	status datatype.WithdrawalStatus,
) error {
	return nil
}

// InsertWithdrawal mock
func (db *UserModelMock) InsertWithdrawal(
	user datatype.User,
	assetID datatype.ID,
	amount decimaldt.Decimal,
	toAddress string,
) (*datatype.UserWithdrawal, error) {
	return nil, nil
}

// ConfirmWithdrawal mock
func (db *UserModelMock) ConfirmWithdrawal(withdrawal datatype.ID, token string) error {
	return nil
}

// CancelUnconfirmedWithdrawal mock
func (db *UserModelMock) CancelUnconfirmedWithdrawal(withdrawalID datatype.ID, user *datatype.User) error {
	return nil
}

// SetManyWithdrawalStatus mock
func (db *UserModelMock) SetManyWithdrawalStatus(
	withdrawals []datatype.UserWithdrawal,
	status datatype.WithdrawalStatus,
) error {
	return nil
}

// CompleteManyWithdrawal mock
func (db *UserModelMock) CompleteManyWithdrawal(
	txID string,
	withdrawals []datatype.UserWithdrawal,
) error {
	return nil
}

// GetUserWithdrawals mock
func (db *UserModelMock) GetUserWithdrawals(
	userID datatype.ID,
	pendingOnly bool,
) ([]datatype.UserWithdrawal, error) {
	return nil, nil
}
func (db *UserModelMock) SaveUserTwofactor(userID datatype.ID, tfa datatype.UserTwofactor) error {
	return nil
}

func (db *UserModelMock) GetUserTwofactor(userID datatype.ID) *datatype.UserTwofactor {
	return nil
}

func (db *UserModelMock) ToggleTwofactor(userID datatype.ID, isActive bool) error {
	return nil
}
