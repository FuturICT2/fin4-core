package models

import "errors"

var (
	// ErrServerError error type
	ErrServerError = errors.New("Server error")
	// ErrBadRequest error type
	ErrBadRequest = errors.New("Bad request")
	// ErrTryAgain error type
	ErrTryAgain = errors.New("Can't complete this operation at the moment, please try again")
	// ErrEmailInvalid error type
	ErrEmailInvalid = errors.New("Email address is invalid")
	// ErrEmailDoesntExist  error type
	ErrEmailDoesntExist = errors.New("Email doesn't exist")
	// ErrEmailExists  error type
	ErrEmailExists = errors.New("Email is already registered")
	// ErrPasswordTooShort error type
	ErrPasswordTooShort = errors.New("Password must be at least 6 characters long")
)
