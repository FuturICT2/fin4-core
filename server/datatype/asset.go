package datatype

import "time"

// Asset asset data type
type Asset struct {
	ID                         ID
	Name                       string
	Symbol                     string
	Description                string
	Supply                     int64
	CreatorID                  ID
	CreatorName                string
	MinersCounter              int
	FavoritesCounter           int
	EthereumAddress            string
	EthereumTransactionAddress string
	DidUserLike                bool
	CreatedAt                  time.Time
	CreatedAtHuman             string
}
