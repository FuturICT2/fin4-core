package userhandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

//Logout logout route
func Logout(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.Logout(c)
		c.JSON(http.StatusOK, gin.H{})
		// c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}
