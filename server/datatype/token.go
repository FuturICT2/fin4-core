package datatype

import (
	"database/sql"
	"time"
)

type Action struct {
	ID          ID             `json:"id"`
	Description string         `json:"description"`
	CreatorID   ID             `json:"creatorID"`
	CreatorName string         `json:"creatorName"`
	Status      int            `json:"status"`
	LogoFile    sql.NullString `json:"logoFile"`
	StartsAt    time.Time      `json:"startsAt"`
	EndsAt      time.Time      `json:"endsAt"`
	CreatedAt   time.Time      `json:"createdAt"`
}

// Token Token data type
type Token struct {
	ID                ID
	CreatorID         ID
	CreatorName       string
	Name              string
	Symbol            string
	BlockchainAddress string
	TotalSupply       string
	Purpose           string
	TxAddress         string
	Logo              string
	FavouriteCount    int
	DidUserLike       bool
	Claims            []Claim
	Miners            []Miner
	Likers            []Liker
}

// Balance balance type
type Claim struct {
	ID         ID
	UserID     ID
	UserName   string
	TokenID    ID
	Text       string
	IsApproved bool
	LogoFile   string
}
