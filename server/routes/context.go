package routes

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

func mustGetUser(c *gin.Context) *datatype.User {
	return c.MustGet("user").(*datatype.User)
}
