package datatype

import "github.com/FuturICT2/fin4-core/server/decimaldt"

// Miner user type
type Miner struct {
	ID               ID
	UserName         string
	Mined            decimaldt.Decimal
	MiningPercentage string
}
