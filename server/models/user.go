package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/FuturICT2/fin4-core/server/decimaldt"
	"github.com/sirupsen/logrus"
)

// User user type
type User struct {
	ID              ID        `json:"id"`
	Name            string    `json:"name"`
	EthereumAddress string    `json:"ethereumAddress"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// Balance balance type
type Balance struct {
	UserID              ID
	AssetID             ID
	AssetName           string
	AssetSymbol         string
	Balance             decimaldt.Decimal
	Reserved            decimaldt.Decimal
	DepositAddress      sql.NullString
	IsDepositEnabled    bool
	IsWithdrawalEnabled bool
	NetworkFee          decimaldt.Decimal
}

// NewUser creates a new user
func NewUser() *User {
	return &User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

const (
	tableUser = "user"
)

const tokenMaxTTLInHours = 48

// UserStore interface for user store
type UserStore interface {
	Register(name string, address string) (*User, error)
	FindByID(ID) (*User, error)
	InsertToken(
		userID ID,
		name string,
		symbol string,
		purpose string,
		totalSupply string,
		blockchainAddress string,
		txAddress string,
		logo string,
	) (*Token, error)
	FindByName(string) (*User, error)
	IsNameRegistered(string) bool
	FindTokens() ([]Token, error)
	FindUsers() ([]User, error)
	DoLike(userID ID, tokenID ID) error
}

//DoLike Registers a new user
func (db *UserModel) DoLike(userID ID, tokenID ID) error {
	_, err := db.Exec(
		`INSERT INTO token_like SET
			userId = ?,
			tokenId = ?`,
		userID,
		tokenID,
	)
	if err != nil {
		return ErrServerError
	}
	return nil
}

// FindUsers finds all users
func (db *UserModel) FindUsers() ([]User, error) {
	result := []User{}
	rows, err := db.Query(
		fmt.Sprintf(`SELECT id, name, ethereumAddress FROM user`),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c User
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.EthereumAddress,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

//Register Registers a new user
func (db *UserModel) Register(name string, address string) (*User, error) {
	var err error
	u, err := db.FindByName(name)
	if err == nil {
		return u, nil
	}
	user := NewUser()
	user.Name = name
	res, err := db.Exec(
		`INSERT INTO user SET
			name = ?,
			ethereumAddress = ?,
			createdAt = ?,
			updatedAt = ?`,
		user.Name,
		address,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return nil, ErrServerError
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, ErrServerError
	}
	return db.FindByID(ID(id))
}

// FindByID finds a user by id
func (db *UserModel) FindByID(id ID) (*User, error) {
	row := db.QueryRow(
		fmt.Sprintf("SELECT %s FROM user WHERE id=?", getUserCols()),
		id,
	)
	return scanUser(row)
}

//FindByName finds a user by email
func (db *UserModel) FindByName(name string) (*User, error) {
	row := db.QueryRow(
		fmt.Sprintf("SELECT %s FROM user WHERE name=?", getUserCols()),
		name,
	)

	return scanUser(row)
}

//IsNameRegistered checks whether email is already registered or not
func (db *UserModel) IsNameRegistered(name string) bool {
	var count int32
	err := db.QueryRow(
		"SELECT count(*) as counter FROM user WHERE name=?",
		name,
	).Scan(&count)
	return err == nil && count > 0
}

func getUserCols() string {
	return "id, name, ethereumAddress, createdAt, updatedAt"
}

func scanUser(row *sql.Row) (*User, error) {
	var user User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.EthereumAddress,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("Not found")
	} else if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:scanUser:1")
		return nil, err
	}
	return &user, nil
}
