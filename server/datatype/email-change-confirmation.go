package datatype

// EmailChangeConfirmation email change confirmation
type EmailChangeConfirmation struct {
	UserID ID
	Email  string
	Token  string
}
