package datatype

// TimelineService defines timeline service interface
type TimelineService interface {
	FindAll(
		sc ServiceContainer,
		user *User,
		timelineFilter TimelineFilter,
		offset int,
		limit int,
	) ([]TimelineEntry, bool, error) // entries, hasMore, error
}
