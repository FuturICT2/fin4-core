package tokenservice

import (
	"errors"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

// InsertToken insert token
func (db *Service) InsertToken(
	userID datatype.ID,
	name string,
	symbol string,
	purpose string,
	hardCap string,
	blockchainAddress string,
	txAddress string,
) (*datatype.Token, error) {

	{ // check if token name and symbol dont exist already
		token, _ := db.FindTokenByName(name)
		if token != nil {
			return nil, errors.New("Token with this name already exists")
		}
		token, _ = db.FindTokenBySymbol(symbol)
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
					hardCap = ?,
					blockchainAddress = ?,
          txAddress = ?
	      `,
		userID,
		name,
		symbol,
		purpose,
		hardCap,
		blockchainAddress,
		txAddress,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("token:InsertToken:e4")
		return nil, datatype.ErrServerError
	}
	tokenID, err := res.LastInsertId()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("token:InsertToken:e5")
		return nil, datatype.ErrServerError
	}
	var token datatype.Token
	token.ID = datatype.ID(tokenID)
	token.CreatorID = userID
	token.Name = name
	token.Symbol = symbol
	token.Purpose = purpose
	token.BlockchainAddress = blockchainAddress
	token.TxAddress = txAddress
	return &token, nil
}
