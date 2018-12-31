package routes

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/FuturICT2/fin4-core/server/decimaldt"
	"github.com/FuturICT2/fin4-core/server/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// ActionsList handler to return existing socially actionable tokens
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

// ApproveActionClaim approve a proposal, disburse funds, and close action
func (env *Env) ApproveActionClaim(c *gin.Context) {
	user := mustGetUser(c)
	body := struct {
		ClaimID int `json:"claimId"`
	}{}
	c.BindJSON(&body)
	userModel := env.DB.NewUserModel()
	claim, err := userModel.FindClaim(models.ID(body.ClaimID))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	token := userModel.FindToken(claim.TokenID)
	if token == nil || token.CreatorID != user.ID {
		c.String(http.StatusBadRequest, "Token not found or not authorized")
		return
	}
	doer, err := userModel.FindByID(claim.UserID)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	tx, err := env.Ethereum.Mint(
		common.HexToAddress(token.BlockchainAddress),
		common.HexToAddress(doer.EthereumAddress),
		1,
	)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	log.Println("TX of minting --=-->>>>>", tx.Hash())
	err = userModel.ApproveActionClaim(models.ID(body.ClaimID), user.ID)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "")
}

// ProposeActionSolution API to add action propsal soltion
func (env *Env) NewActionClaim(c *gin.Context) {
	user := mustGetUser(c)
	body := struct {
		Proposal string `json:"proposal"`
		TokenID  int    `json:"tokenId"`
	}{}
	c.BindJSON(&body)
	if len(body.Proposal) < 1 || len(body.Proposal) > 10000 {
		c.String(http.StatusBadRequest, "Proposal length should be between than 1 and 10000 characters")
		return
	}
	userModel := env.DB.NewUserModel()
	err := userModel.NewActionClaim(user.ID, body.Proposal, models.ID(body.TokenID))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "")
}

// CreateAction creates an action
func (env *Env) CreateAction(c *gin.Context) {
	user := mustGetUser(c)
	body := struct {
		Description string `json:"description"`
		TimeLimit   string `json:"timeLimit"`
	}{}
	c.BindJSON(&body)
	if len(body.Description) < 1 || len(body.Description) > 10000 {
		c.String(http.StatusBadRequest, "Description length should be between than 1 and 10000 characters")
		return
	}
	userModel := env.DB.NewUserModel()
	now := time.Now()
	timeLimit, err := strconv.ParseFloat(body.TimeLimit, 64)
	if err != nil || timeLimit < 0 || timeLimit > 48 {
		c.String(http.StatusBadRequest, "Time limit should be a positive valid number and less than 48hrs")
		return
	}
	err = userModel.InsertAction(
		user.ID,
		body.Description,
		now,
		now.Add(time.Duration(timeLimit*60*60*1000000000)),
	)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
}

// AddSupportToAction users can add tokens to actions to increase incentives
func (env *Env) AddSupportToAction(c *gin.Context) {
	user := mustGetUser(c)
	body := struct {
		ActionID models.ID         `json:"actionId"`
		TokenID  models.ID         `json:"tokenId"`
		Amount   decimaldt.Decimal `json:"amount"`
	}{}
	c.BindJSON(&body)
	userModel := env.DB.NewUserModel()
	err := userModel.ReserveRewardsForAction(
		user.ID,
		body.TokenID,
		body.ActionID,
		body.Amount,
	)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
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
