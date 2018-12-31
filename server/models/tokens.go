package models

import (
	"errors"
	"fmt"

	"github.com/FuturICT2/fin4-core/server/decimaldt"
	"github.com/lytics/logrus"
)

// FindTokens finds all tokens
func (db *UserModel) FindTokens(userID ID) ([]Token, error) {
	result := []Token{}
	rows, err := db.Query(
		fmt.Sprintf(`
			SELECT %s, u.name
			FROM token t
			LEFT JOIN
				user u ON u.id = t.creatorId
			ORDER BY id DESC`, getTokenCols()),
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("models:FindTokens:e1")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c Token
		rows.Scan(
			&c.ID,
			&c.CreatorID,
			&c.Name,
			&c.Symbol,
			&c.TotalSupply,
			&c.Purpose,
			&c.BlockchainAddress,
			&c.TxAddress,
			&c.Logo,
			&c.CreatorName,
		)
		c.FavouriteCount = db.CountLikes(c.ID)
		c.DidUserLike = db.DidUserLike(userID, c.ID)
		claims, err := db.GetTokenClaims(userID, c.ID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("models:FindTokens:e1.2")
			return nil, err
		}
		c.Claims = claims
		result = append(result, c)
	}
	if err := rows.Err(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("models:FindTokens:e2")
		return nil, err
	}
	return result, nil
}

// GetUserBalance returns token balance of a user
func (db *UserModel) GetUserBalances(userId ID) ([]Balance, error) {
	result := []Balance{}
	rows, err := db.Query(`SELECT
			b.tokenId,
			b.balance,
			b.reserved,
			t.name,
			t.symbol,
			t.logo,
			t.blockchainAddress
		FROM user_holding b
		LEFT JOIN
			token t ON t.id = b.tokenId
		WHERE b.userId=?`,
		userId,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("GetUserBalances:1")
		return result, ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := Balance{UserID: userId}
		err = rows.Scan(
			&entry.TokenID,
			&entry.Balance,
			&entry.Reserved,
			&entry.TokenName,
			&entry.TokenSymbol,
			&entry.LogoFile,
			&entry.BlockchainAddress,
		)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{"error": err.Error()},
			).Error("GetUserBalances:2")
			return result, ErrServerError
		}
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("GetUserBalances:3")
		return result, ErrServerError
	}
	return result, nil
}

// CountLikes returnes number of likes for the passed token
func (db *UserModel) CountLikes(tokenID ID) int {
	var count int
	db.QueryRow(
		`SELECT
			count(*)
		FROM token_like
		WHERE tokenId = ? `,
		tokenID,
	).Scan(&count)
	return count
}

func (db *UserModel) GetTokenClaims(userID ID, tokenID ID) ([]Claim, error) {
	result := []Claim{}
	rows, err := db.Query(`SELECT
			c.id,
			c.tokenId,
			c.userId,
			u.name,
			c.text,
			c.logoFile,
			c.isApproved
		FROM claim c
		LEFT JOIN
			user u ON u.id = c.userId
		WHERE c.tokenId=?`,
		tokenID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("GetTokenClaims:1")
		return result, ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := Claim{}
		err = rows.Scan(
			&entry.ID,
			&entry.TokenID,
			&entry.UserID,
			&entry.UserName,
			&entry.Text,
			&entry.LogoFile,
			&entry.IsApproved,
		)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{"error": err.Error()},
			).Error("GetTokenClaims:2")
			return result, ErrServerError
		}
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("GetTokenClaims:3")
		return result, ErrServerError
	}
	return result, nil
}

// DidUserLike returns true or false based on if user liked a token
func (db *UserModel) DidUserLike(userID ID, tokenID ID) bool {
	var count int
	db.QueryRow(
		`SELECT
			count(*)
		FROM token_like
		WHERE tokenId = ? and userId =? `,
		tokenID,
		userID,
	).Scan(&count)
	return count > 0
}

// FindToken finds Token by ID
func (db *UserModel) FindToken(id ID) *Token {
	var c Token
	db.QueryRow(
		fmt.Sprintf(`SELECT %s FROM token t WHERE id = ?`, getTokenCols()),
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
	return &c
}

// FindTokenBySymbol find token by symbol
func (db *UserModel) FindTokenBySymbol(symbol string) *Token {
	var c Token
	err := db.QueryRow(
		fmt.Sprintf(`SELECT %s FROM token t WHERE symbol = ?`, getTokenCols()),
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
	if err != nil {
		return nil
	}
	return &c
}

// FindTokenByName find token by name
func (db *UserModel) FindTokenByName(name string) *Token {
	var c Token
	err := db.QueryRow(
		fmt.Sprintf(`SELECT %s FROM token t WHERE name = ?`, getTokenCols()),
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
	if err != nil {
		return nil
	}
	return &c
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
		token := db.FindTokenByName(name)
		if token != nil {
			return nil, errors.New("Token with this name already exists")
		}
		token = db.FindTokenBySymbol(symbol)
		if token != nil {
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

// InsertBalance insert balance
func (db *UserModel) InsertBalance(
	userID ID,
	tokenID ID,
	balance string,
) error {
	_, err := db.Exec(
		`INSERT INTO user_holding SET
          userId = ?,
	        tokenId = ?,
					balance = ?
	      `,
		userID,
		tokenID,
		balance,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("InsertBalance:1")
		return ErrServerError
	}
	return nil
}

// AddToBalance insert balance
func (db *UserModel) AddToBalance(
	userID ID,
	tokenID ID,
	value string,
) error {
	var balance string
	err := db.QueryRow(
		fmt.Sprintf(`SELECT balance FROM user_holding WHERE userID = ? and tokenID = ?`),
		userID,
		tokenID,
	).Scan(
		&balance,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("AddToBalance:0")
		return ErrServerError
	}

	b, _ := decimaldt.NewFromString(balance)
	v, _ := decimaldt.NewFromString(value)
	all := b.Add(v)
	_, err = db.Exec(
		`update user_holding SET
					balance = ? where userId =? and tokenId = ?
	      `,
		all,
		userID,
		tokenID,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("AddToBalance:1")
		return ErrServerError
	}
	return nil
}

func getTokenCols() string {
	return `
		t.id,
		t.creatorId,
		t.name,
		t.symbol,
		t.totalSupply,
		t.purpose,
		t.blockchainAddress,
		t.txAddress,
		t.logo`
}
