package userhandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

// SessionGet get session route
func SessionGet(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*datatype.User)
		c.JSON(http.StatusOK, user)
	}
}

// SessionDestroy logout route
func SessionDestroy(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.Logout(c)
		c.JSON(http.StatusOK, nil)
	}
}
