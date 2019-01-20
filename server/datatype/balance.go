package datatype

import "github.com/FuturICT2/fin4-core/server/decimaldt"

// Balance balance type
type Balance struct {
	UserID            ID
	TokenID           ID
	Balance           decimaldt.Decimal
	Reserved          decimaldt.Decimal
	Mined             decimaldt.Decimal
	TokenName         string
	TokenSymbol       string
	LogoFile          string
	BlockchainAddress string
}
