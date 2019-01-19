package tokenhandlers

import (
	"net/http"
	"strconv"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

// FindTokenForUser handler to return the required token
func FindTokenForUser(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		tokenID, err := strconv.Atoi(c.Params.ByName("tokenID"))
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}
		token, err := sc.TokenService.FindTokenForUser(datatype.ID(tokenID), user.ID)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, token)
	}
}
