package datatype

// TimelineType a data type representing timeline type (SELL, BUY)
type TimelineType uint8

const (
	// HOMETIMELINE home asset timeline
	HOMETIMELINE = 0
	// USERTIMELINE user asset timeline
	USERTIMELINE = 1
	// ASSETTIMELINE asset timeline
	ASSETTIMELINE = 2
)

// TimelineFilter timeline entry data type
type TimelineFilter struct {
	Type    TimelineType
	UserID  ID
	AssetID ID
}
