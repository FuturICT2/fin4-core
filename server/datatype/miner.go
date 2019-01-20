package datatype

import "github.com/FuturICT2/fin4-core/server/decimaldt"

// Miner Miner data type
type Miner struct {
	UserID   ID
	UserName string
	Mined    decimaldt.Decimal
}
