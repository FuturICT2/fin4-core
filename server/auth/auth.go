package auth

import (
	"fmt"

	"github.com/FuturICT2/fin4-core/server/models"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Login logs user in
func Login(c *gin.Context, user *models.User) {
	session := sessions.Default(c)
	session.Set("userId", fmt.Sprintf("%d", user.ID))
	session.Save()
}

// Logout logs user out
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userId")
	session.Save()
}
