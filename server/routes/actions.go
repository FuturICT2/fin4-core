package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/FuturICT2/fin4-core/server/decimaldt"
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

// AprroveProposal approve a proposal, disburse funds, and close action
// func (env *Env) AprroveProposal(c *gin.Context) {
// 	user := mustGetUser(c)
// 	body := struct {
// 		ActionID   int `json:"actionId"`
// 		ProposalID int `json:"proposalId"`
// 		DoerID     int `json:"doerId"`
// 	}{}
// 	c.BindJSON(&body)
// 	userModel := env.DB.NewUserModel()
// 	err := userModel.AprroveProposal(
// 		body.ActionID, body.ProposalID, body.DoerID, user.ID)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	c.String(http.StatusOK, "")
// }

// ProposeActionSolution API to add action propsal soltion
func (env *Env) AddActionProposal(c *gin.Context) {
	user := mustGetUser(c)
	body := struct {
		Proposal string `json:"proposal"`
		ActionID int    `json:"actionId"`
	}{}
	c.BindJSON(&body)
	if len(body.Proposal) < 1 || len(body.Proposal) > 10000 {
		c.String(http.StatusBadRequest, "Proposal length should be between than 1 and 10000 characters")
		return
	}
	userModel := env.DB.NewUserModel()
	err := userModel.AddActionProposal(user.ID, body.Proposal, models.ID(body.ActionID))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "")
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
			TotalRewards  decimaldt.Decimal        `json:"totalRewrads"`
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
			TotalRewards:  getTotalRewards(entry.Supporters),
		})
	}
	return res
}

func getTotalRewards(supporters []models.ActionSupporter) decimaldt.Decimal {
	var total decimaldt.Decimal
	for _, supporter := range supporters {
		total = total.Add(supporter.Amount)
	}
	return total
}
