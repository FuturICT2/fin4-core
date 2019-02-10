package timelinehandlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/dbservice"
)

// FindEntries get time line entries handler
func FindEntries(sc datatype.ServiceContainer, timelineType datatype.TimelineType) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 1
		}
		limit := 25
		offset := page*limit - limit
		timelineFilter := datatype.TimelineFilter{}

		timelineFilter.Type = timelineType
		assetID, _ := dbservice.StringToID(c.Params.ByName("assetID"))
		timelineFilter.AssetID = assetID
		userID, _ := dbservice.StringToID(c.Params.ByName("userID"))
		timelineFilter.UserID = userID
		entries, hasMore, err := sc.TimelineService.FindAll(
			sc,
			user,
			timelineFilter,
			offset,
			limit,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, struct {
			HasMore bool
			Page    int
			Entries []datatype.TimelineEntry
		}{
			HasMore: hasMore,
			Page:    page,
			Entries: entries,
		})
	}
}
