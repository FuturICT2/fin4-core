package tokenservice

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
)

func (db *Service) FindTokenByName(name string) (*datatype.Token, error) {
	return db.findTokenBy("name", name)
}
