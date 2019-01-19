package userservice

import (
	"fmt"

	"github.com/FuturICT2/fin4-core/server/datatype"
)

// FindByID finds a user by id
func (db *Service) FindByID(id datatype.ID) (*datatype.User, error) {
	row := db.QueryRow(
		fmt.Sprintf("SELECT %s FROM user WHERE id=?", getUserCols()),
		id,
	)
	return scanUser(row)
}
