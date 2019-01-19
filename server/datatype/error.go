package datatype

import "errors"

var (
	//ErrInvalidRequest invalid request
	ErrInvalidRequest = errors.New("Invalid request")

	// ErrNotEnoughBalance not enough balance
	ErrNotEnoughBalance = errors.New("Not enough balance")

	// ErrServerError for any internal server error
	ErrServerError = errors.New("Server error")

	// ErrBadRequest send it when user tries to do something illegal
	ErrBadRequest = errors.New("Bad request")

	// ErrTryAgain for any try again error
	ErrTryAgain = errors.New("Can't complete this operation at the moment, please try again")

	// ErrMarketCantBeNil when required market is missing
	ErrMarketCantBeNil = errors.New("Market can't be nil")

	// ErrOrderCantBeNil when required order is missing
	ErrOrderCantBeNil = errors.New("Order can't be nil")

	//ErrOrderQuantityMustBeBiggerThanZero for when quantity is too little
	ErrOrderQuantityMustBeBiggerThanZero = errors.New("Quantity must be bigger than 0")

	//ErrOrderPriceMustBeBiggerThanZero for when quantity is too little
	ErrOrderPriceMustBeBiggerThanZero = errors.New("Price must be bigger than 0")

	ErrEmailInvalid           = errors.New("Email address is invalid")
	ErrEmailDoesntExist       = errors.New("Email doesn't exist")
	ErrEmailExists            = errors.New("Email is already registered")
	ErrPasswordTooShort       = errors.New("Password must be at least 6 characters long")
	ErrPasswordSameAsEmail    = errors.New("Password can not be the same as email")
	ErrAgreeToTerms           = errors.New("You must agree to site terms & privacy")
	ErrInvalidCurrentPassword = errors.New("Invalid current password")

	ErrOrderDeosntBelongToMarket = errors.New("Order doesn't belong to market")
)
