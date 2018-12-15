package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/FuturICT2/fin4-core/server/models"
	"github.com/gin-gonic/gin"
)

// ActionsList handler to return json of existing tokens
func (env *Env) ActionsList(c *gin.Context) {
	// user := mustGetUser(c)
	userModel := env.DB.NewUserModel()
	actions, err := userModel.FindActions(models.ID(1))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, struct {
		Count   int
		Limit   int
		Page    int
		Entries []interface{}
	}{
		Count:   len(actions),
		Limit:   10,
		Page:    0,
		Entries: toActionsResponse(actions),
	})
}

func toActionsResponse(entries []models.Action) []interface{} {
	res := []interface{}{}
	for _, entry := range entries {
		duration := time.Until(entry.EndsAt)
		endsInHours := strconv.FormatFloat(duration.Hours(), 'f', 0, 64)
		isTimeout := false
		if duration.Hours() < 0 {
			endsInHours = "0"
		}

		endsInMinutes := strconv.FormatFloat(duration.Minutes(), 'f', 0, 64)
		if duration.Minutes() < 0 {
			endsInMinutes = "0"
			isTimeout = true
		}
		res = append(res, struct {
			ID            models.ID                `json:"id"`
			Description   string                   `json:"description"`
			CreatorID     models.ID                `json:"creatorID"`
			CreatorName   string                   `json:"creatorName"`
			Status        int                      `json:"status"`
			LogoFile      string                   `json:"logoFile"`
			StartsAt      time.Time                `json:"startsAt"`
			EndsAt        time.Time                `json:"endsAt"`
			CreatedAt     time.Time                `json:"createdAt"`
			Proposals     []models.ActionProposal  `json:"proposals"`
			Supporters    []models.ActionSupporter `json:"supporters"`
			EndsInHours   string                   `json:"endsInHours"`
			EndsInMinutes string                   `json:"endsInMinutes"`
			IsTimeout     bool                     `json:'isTimeout'`
			IsTimeLimit   bool                     `json:"isTimeLimit"`
		}{
			ID:            entry.ID,
			Description:   entry.Description,
			CreatorID:     entry.CreatorID,
			CreatorName:   entry.CreatorName,
			Status:        entry.Status,
			LogoFile:      entry.LogoFile.String,
			StartsAt:      entry.StartsAt,
			EndsAt:        entry.EndsAt,
			CreatedAt:     entry.CreatedAt,
			Proposals:     entry.Proposals,
			Supporters:    entry.Supporters,
			EndsInHours:   endsInHours,
			EndsInMinutes: endsInMinutes,
			IsTimeout:     isTimeout,
			IsTimeLimit:   !entry.StartsAt.Equal(entry.EndsAt),
		})
	}
	return res
}
