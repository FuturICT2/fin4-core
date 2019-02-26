package datatype

import "time"

const (
	// OraclePerson person oracle type
	OraclePerson = 0
	// OracleSensor sensor oracle type
	OracleSensor = 1
)

// Asset asset data type
type Asset struct {
	ID                         ID
	Name                       string
	Symbol                     string
	Description                string
	Supply                     int64
	CreatorID                  ID
	OracleType                 int
	CreatorName                string
	MinersCounter              int
	FavoritesCounter           int
	EthereumAddress            string
	EthereumTransactionAddress string
	DidUserLike                bool
	CreatedAt                  time.Time
	CreatedAtHuman             string
	LastOraclePing             time.Time
	LastOraclePingHuman        string
	IsConnected                bool
	AccessToken                string
}
