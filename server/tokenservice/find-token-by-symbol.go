package tokenservice

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
)

func (db *Service) FindTokenBySymbol(symbol string) (*datatype.Token, error) {
	return db.findTokenBy("symbol", symbol)
}
