package datatype

import "time"

const (
	// BlockUnverified unverified status for block
	BlockUnverified = 0
	// BlockAccepted accepted status for block
	BlockAccepted = 1
	// BlockRejected rejected status for block
	BlockRejected = 2
)

// Block asset block data type
type Block struct {
	ID             ID
	UserID         ID
	UserName       string
	AssetID        ID
	Text           string
	Images         []string
	Status         int
	CreatedAt      time.Time
	CreatedAtHuman string
}
