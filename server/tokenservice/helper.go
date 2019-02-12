package tokenservice

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/decimaldt"
	"github.com/lytics/logrus"
)

func (db *Service) getTokenMiners(tokenID datatype.ID) ([]datatype.Miner, error) {
	result := []datatype.Miner{}
	rows, err := db.Query(`SELECT
			c.userId,
			u.name
		FROM claim c
		LEFT JOIN
			user u ON u.id = c.userId
		WHERE c.tokenId=? AND c.isApproved=1`,
		tokenID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("getTokenMiners:1")
		return result, datatype.ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := datatype.Miner{}
		err = rows.Scan(
			&entry.ID,
			&entry.UserName,
		)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{"error": err.Error()},
			).Error("getTokenMiners:2")
			return result, datatype.ErrServerError
		}
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("getTokenMiners:3")
		return result, datatype.ErrServerError
	}

	minersMap := map[int]datatype.Miner{}
	one, _ := decimaldt.NewFromString("1.0")
	// TODO make more clean using SQL
	for _, miner := range result {
		key := int(miner.ID)
		m := minersMap[key]
		if m.UserName == "" {
			miner.Mined = miner.Mined.Add(one)
			minersMap[key] = miner
		} else {
			newMiner := minersMap[key]
			newMiner.Mined = newMiner.Mined.Add(one)
			minersMap[key] = newMiner
		}
	}
	clean := []datatype.Miner{}
	for _, v := range minersMap {
		clean = append(clean, v)
	}
	return clean, nil
}

func (db *Service) getTokenLikers(tokenID datatype.ID) ([]datatype.Liker, error) {
	result := []datatype.Liker{}
	rows, err := db.Query(`SELECT
			tl.userId,
			u.name
		FROM token_like tl
		LEFT JOIN
			user u ON u.id = tl.userId
		WHERE tl.tokenId=?`,
		tokenID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("getTokenLikers:1")
		return result, datatype.ErrServerError
	}
	defer rows.Close()
	for rows.Next() {
		entry := datatype.Liker{}
		err = rows.Scan(
			&entry.UserID,
			&entry.UserName,
		)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{"error": err.Error()},
			).Error("getTokenLikers:2")
			return result, datatype.ErrServerError
		}
		result = append(result, entry)
	}
	if err = rows.Err(); err != nil {
		logrus.WithFields(
			logrus.Fields{"error": err.Error()},
		).Error("getTokenLikers:3")
		return result, datatype.ErrServerError
	}
	return result, nil
}

func (db *Service) findTokenBy(by string, value string) (*datatype.Token, error) {
	var c datatype.Token
	sqlStmt := fmt.Sprintf(`SELECT %s FROM token t WHERE %s = ?`, getTokenCols(), by)
	err := db.QueryRow(
		sqlStmt,
		value,
	).Scan(
		&c.ID,
		&c.CreatorID,
		&c.Name,
		&c.Symbol,
		&c.TotalSupply,
		&c.Purpose,
		&c.BlockchainAddress,
		&c.TxAddress,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New(fmt.Sprintf("No token with %s equals to %s", by, value))
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("tokenservice:helper:findTokenBy:e1")
		return nil, errors.New("Error scanning token")
	}
	return &c, nil
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
		t.txAddress`
}
