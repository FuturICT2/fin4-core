package routes

import (
	"github.com/FuturICT2/fin4-core/server/models"
	"github.com/gin-gonic/gin"
)

func mustGetUser(c *gin.Context) *models.User {
	return c.MustGet("user").(*models.User)
}
