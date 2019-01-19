package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

func mustGetUser(c *gin.Context) *datatype.User {
	return c.MustGet("user").(*datatype.User)
}
