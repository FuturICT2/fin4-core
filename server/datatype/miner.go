package datatype

import "github.com/kjda/exchange/server/decimaldt"

// Miner Miner data type
type Miner struct {
	UserID   ID
	UserName string
	Mined    decimaldt.Decimal
}
