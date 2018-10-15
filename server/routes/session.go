package routes

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/models"
	"github.com/gin-gonic/gin"
)

// SessionGet get session route
func (env *Env) SessionGet(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	c.JSON(http.StatusOK, user)
}

// SessionDestroy logout route
func (env *Env) SessionDestroy(c *gin.Context) {
	auth.Logout(c)
	c.JSON(http.StatusOK, nil)
}
