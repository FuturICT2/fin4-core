package userhandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

// Balances return user balances
func Balances(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := auth.MustGetUser(c)
		balances, err := sc.UserService.GetBalances(u.ID)
		if err != nil {
			c.String(http.StatusServiceUnavailable, err.Error())
			return
		}
		c.JSON(http.StatusOK, struct {
			Count   int
			Limit   int
			Page    int
			Entries []datatype.Balance
		}{
			Count:   len(balances),
			Limit:   10,
			Page:    0,
			Entries: balances,
		})
	}
}
