package datatype

import "time"

// TimelineEntry timeline entry data type
type TimelineEntry struct {
	BlockID             ID
	UserID              ID
	UserName            string
	UserProfileImageURL string
	AssetID             ID
	AssetName           string
	AssetSymbol         string
	OracleID            ID
	OracleName          string
	Text                string
	Status              int
	YtVideoID           string
	FavoritesCount      int
	DidUserLike         bool
	CreatedAt           time.Time
	CreatedAtHuman      string
	Images              []string
}
