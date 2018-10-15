package models

import (
	"errors"
	"fmt"

	"github.com/lytics/logrus"
)

// Token Token data type
type Token struct {
	ID                ID
	CreatorID         ID
	Name              string
	Symbol            string
	BlockchainAddress string
	TotalSupply       string
	Purpose           string
	TxAddress         string
	Logo              string
}

// FindTokens finds all tokens
func (db *UserModel) FindTokens() ([]Token, error) {
	result := []Token{}
	rows, err := db.Query(
		fmt.Sprintf(`SELECT %s FROM token`, getTokenCols()),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c Token
		err := rows.Scan(
			&c.ID,
			&c.CreatorID,
			&c.Name,
			&c.Symbol,
			&c.TotalSupply,
			&c.Purpose,
			&c.BlockchainAddress,
			&c.TxAddress,
			&c.Logo,
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

// FindToken finds Token by ID
func (db *UserModel) FindToken(id ID) (*Token, error) {
	var c Token
	err := db.QueryRow(
		fmt.Sprintf(`SELECT %s FROM token WHERE id = ?`, getTokenCols()),
		id,
	).Scan(
		&c.ID,
		&c.CreatorID,
		&c.Name,
		&c.Symbol,
		&c.TotalSupply,
		&c.Purpose,
		&c.BlockchainAddress,
		&c.TxAddress,
		&c.Logo,
	)
	return &c, err
}

// FindTokenBySymbol find token by symbol
func (db *UserModel) FindTokenBySymbol(symbol string) (*Token, error) {
	var c Token
	err := db.QueryRow(
		fmt.Sprintf(`SELECT %s FROM token WHERE symbol = ?`, getTokenCols()),
		symbol,
	).Scan(
		&c.ID,
		&c.CreatorID,
		&c.Name,
		&c.Symbol,
		&c.TotalSupply,
		&c.Purpose,
		&c.BlockchainAddress,
		&c.TxAddress,
		&c.Logo,
	)
	return &c, err
}

// FindTokenByName find token by name
func (db *UserModel) FindTokenByName(name string) (*Token, error) {
	var c Token
	err := db.QueryRow(
		fmt.Sprintf(`SELECT %s FROM token WHERE name = ?`, getTokenCols()),
		name,
	).Scan(
		&c.ID,
		&c.CreatorID,
		&c.Name,
		&c.Symbol,
		&c.TotalSupply,
		&c.Purpose,
		&c.BlockchainAddress,
		&c.TxAddress,
		&c.Logo,
	)
	return &c, err
}

// InsertToken insert token
func (db *UserModel) InsertToken(
	userID ID,
	name string,
	symbol string,
	purpose string,
	totalSupply string,
	blockchainAddress string,
	txAddress string,
	logo string,
) (*Token, error) {

	{ // check if token name and symbol dont exist already
		_, err := db.FindTokenByName(name)
		if err == nil {
			return nil, errors.New("Token with this name already exists")
		}
		_, err = db.FindTokenBySymbol(symbol)
		if err == nil {
			return nil, errors.New("Token with this symbol already exists")
		}
	}
	res, err := db.Exec(
		`INSERT INTO token SET
          creatorId = ?,
	        name = ?,
					symbol = ?,
					purpose = ?,
					totalSupply = ?,
					blockchainAddress = ?,
          txAddress = ?,
					logo = ?
	      `,
		userID,
		name,
		symbol,
		purpose,
		totalSupply,
		blockchainAddress,
		txAddress,
		logo,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("token:InsertToken:e4")
		return nil, ErrServerError
	}
	tokenID, err := res.LastInsertId()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("token:InsertToken:e5")
		return nil, ErrServerError
	}
	var token Token
	token.ID = ID(tokenID)
	token.CreatorID = userID
	token.Name = name
	token.Symbol = symbol
	token.TotalSupply = totalSupply
	token.Purpose = purpose
	token.BlockchainAddress = blockchainAddress
	token.TxAddress = txAddress
	token.Logo = logo
	return &token, nil
}

func getTokenCols() string {
	return `id, creatorId, name, symbol, totalSupply, purpose, blockchainAddress, txAddress, logo`
}
