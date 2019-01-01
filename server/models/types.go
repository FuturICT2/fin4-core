package models

import (
	"database/sql"
	"time"

	"github.com/FuturICT2/fin4-core/server/decimaldt"
)

// User user type
type User struct {
	ID              ID        `json:"id"`
	Name            string    `json:"name"`
	EthereumAddress string    `json:"ethereumAddress"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// UserStore interface for user store
type UserStore interface {
	Register(name string, address string) (*User, error)
	FindByID(ID) (*User, error)
	GetUserBalances(userId ID) ([]Balance, error)
	InsertBalance(
		userID ID,
		tokenID ID,
		balance string,
	) error
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
	FindTokens(userID ID) ([]Token, error)
	FindToken(id ID) *Token
	FindUsers() ([]User, error)
	DoLike(userID ID, tokenID ID, state bool) error
	InsertAction(
		userID ID,
		description string,
		startsAt time.Time,
		endsAt time.Time,
	) error
	ReserveRewardsForAction(
		userID ID,
		tokenID ID,
		actionID ID,
		amount decimaldt.Decimal,
	) error
	FindActions(userid ID) ([]Action, error)
	NewActionClaim(userID ID, proposal string, actionID ID, logoPath string) error
	ApproveActionClaim(claimID ID, approverID ID) error
	FindClaim(claimID ID) (*Claim, error)
}

type ActionProposal struct {
	ID          ID             `json:"id"`
	Description string         `json:"description"`
	DoerID      ID             `json:"doerId"`
	DoerName    string         `json:"doerName"`
	LogoFile    sql.NullString `json:"logoFile"`
	Approved    bool           `json:"approved"`
	CreatedAt   time.Time      `json:"createdAt"`
	IsOwner     bool           `json:"isOwner"`
}

type ActionSupporter struct {
	Amount    decimaldt.Decimal `json:"amount"`
	TokenId   ID                `json:"tokenId"`
	TokenName string            `json:"tokenName"`
	UserId    ID                `json:"userId"`
	UserName  string            `json:"userName"`
	Status    int               `json:"status"`
}

type Action struct {
	ID          ID                `json:"id"`
	Description string            `json:"description"`
	CreatorID   ID                `json:"creatorID"`
	CreatorName string            `json:"creatorName"`
	Status      int               `json:"status"`
	LogoFile    sql.NullString    `json:"logoFile"`
	StartsAt    time.Time         `json:"startsAt"`
	EndsAt      time.Time         `json:"endsAt"`
	CreatedAt   time.Time         `json:"createdAt"`
	Proposals   []ActionProposal  `json:"proposals"`
	Supporters  []ActionSupporter `json: "supporters"`
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
}

// Balance balance type
type Balance struct {
	UserID            ID
	TokenID           ID
	Balance           decimaldt.Decimal
	Reserved          decimaldt.Decimal
	TokenName         string
	TokenSymbol       string
	LogoFile          string
	BlockchainAddress string
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
