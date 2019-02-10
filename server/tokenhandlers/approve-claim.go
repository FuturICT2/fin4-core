package tokenhandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

// ApproveClaim approves a mining claim
func ApproveClaim(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			ClaimID int `json:"claimId"`
		}{}
		c.BindJSON(&body)
		claim, err := sc.TokenService.FindClaim(datatype.ID(body.ClaimID))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		token, err := sc.TokenService.FindTokenForUser(claim.TokenID, user.ID)
		if token == nil || token.CreatorID != user.ID || err != nil {
			c.String(http.StatusBadRequest, "Token not found or not authorized")
			return
		}
		// doer, err := sc.UserService.FindByID(claim.UserID)
		// if err != nil {
		// 	c.String(http.StatusBadRequest, err.Error())
		// 	return
		// }
		err = sc.TokenService.ApproveClaim(datatype.ID(body.ClaimID), user.ID)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "")
	}
}
