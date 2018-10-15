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
	ID        ID        `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
	Register(name string) (*User, error)
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
	FindUserBalances(userID ID) ([]Balance, error)
	FindUserBalance(userID ID, assetID ID) (
		availableBalance decimaldt.Decimal,
		reserved decimaldt.Decimal,
		err error,
	)
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
		fmt.Sprintf(`SELECT id, name FROM user`),
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
func (db *UserModel) Register(name string) (*User, error) {
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
			createdAt = ?,
			updatedAt = ?`,
		user.Name,
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

// FindUserBalances finds all balances belonging to given user
func (db *UserModel) FindUserBalances(id ID) ([]Balance, error) {
	result := []Balance{}
	rows, err := db.Query(`SELECT
			b.assetId,
			b.balance,
			b.reserved,
			c.name,
			c.symbol,
			c.isDepositEnabled,
			c.isWithdrawalEnabled,
			c.networkFee,
			a.address
		FROM user_balance b
		LEFT JOIN
			trade_asset c ON c.id = b.assetId
		LEFT JOIN
			crypto_user_deposit_address a ON a.userId = b.userId AND a.assetId = b.assetId
		WHERE b.userId=?`,
		id,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": id},
		).Error("user:FindUserBalances:1")
		return result, ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := Balance{UserID: id}
		err = rows.Scan(
			&entry.AssetID,
			&entry.Balance,
			&entry.Reserved,
			&entry.AssetName,
			&entry.AssetSymbol,
			&entry.IsDepositEnabled,
			&entry.IsWithdrawalEnabled,
			&entry.NetworkFee,
			&entry.DepositAddress,
		)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{"e": err.Error(), "userID": id},
			).Error("user:FindUserBalances:2")
			return result, ErrServerError
		}
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error(), "userID": id},
		).Error("user:FindUserBalances:3")
		return result, ErrServerError
	}
	return result, nil
}

// FindUserBalance finds user's balance of given currecncy
func (db *UserModel) FindUserBalance(
	userID ID,
	currecncyID ID,
) (
	decimaldt.Decimal,
	decimaldt.Decimal,
	error,
) {
	var balance decimaldt.Decimal
	var reservedBalace decimaldt.Decimal
	err := db.QueryRow(
		`SELECT balance
			FROM user_balance
			WHERE userId = ? AND assetId = ?`,
		userID,
		currecncyID,
	).Scan(&balance)
	if err == sql.ErrNoRows {
		if db.createBalance(userID, currecncyID) {
			return decimaldt.NewFromFloat(0.0), decimaldt.NewFromFloat(0.0), nil
		}
	} else if err != nil {
		logrus.WithFields(
			logrus.Fields{
				"e":           err.Error(),
				"userID":      userID,
				"currecncyID": currecncyID,
			},
		).Error("user:FindUserBalance:1")
		return decimaldt.NewFromFloat(0.0), decimaldt.NewFromFloat(0.0), ErrServerError
	}
	return balance, reservedBalace, err
}

func (db *UserModel) createBalance(userID ID, assetID ID) bool {
	res, err := db.Exec(`INSERT INTO user_balance SET
			userId = ?, assetId = ?, balance = 0, reserved = 0 `,
		userID, assetID,
	)
	if err != nil {
		return false
	}
	_, err = res.LastInsertId()
	return err == nil
}

func getUserCols() string {
	return "id, name, createdAt, updatedAt"
}

func scanUser(row *sql.Row) (*User, error) {
	var user User
	err := row.Scan(
		&user.ID,
		&user.Name,
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
