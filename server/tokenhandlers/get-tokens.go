package tokenhandlers

import (
	"net/http"
	"log"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

// GetTokens fetch tokens from db
func GetTokens(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Print("hallo")

		u := auth.MustGetUser(c)
		tokens, err := sc.TokenService.FindAll(u.ID)
		if err != nil {
			c.String(http.StatusServiceUnavailable, err.Error())
			return
		}
		c.JSON(http.StatusOK, struct {
			Count   int
			Limit   int
			Page    int
			Entries []datatype.Token
		}{
			Count:   len(tokens),
			Limit:   10,
			Page:    0,
			Entries: tokens,
		})
	}
}
