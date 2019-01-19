package tokenhandlers

import (
	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

// ToggleTokenLike toggle user a liking token
func ToggleTokenLike(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			TokenID int `json:"tokenId"`
		}{}
		c.BindJSON(&body)
		sc.TokenService.ToggleTokenLike(user.ID, datatype.ID(body.TokenID))
	}
}
