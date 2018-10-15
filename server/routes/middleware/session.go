package middleware

import (
	"net/http"
	"strconv"

	"github.com/FuturICT2/fin4-core/server/models"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Session starts session
func Session() gin.HandlerFunc {
	store := sessions.NewCookieStore([]byte("|FQyu8eMGdCx2D7h!c6qBKnhgAv8jQmlWVgaf&MCamrylarK!scXYwkeeL'MuQlq"))
	return sessions.Sessions("session", store)
}

// SessionSetUser sets user on context
func SessionSetUser(db *models.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userIDRaw := session.Get("userId")
		var user *models.User
		switch userIDStr := userIDRaw.(type) {
		case string:
			if userID, err := strconv.Atoi(userIDStr); err == nil {
				user, _ = db.NewUserModel().FindByID(models.ID(userID))
			}
			break
		}
		c.Set("user", user)
		c.Next()
	}
}

// SessionMustAuth ensures user is authenticated
func SessionMustAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*models.User)
		if user == nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, nil)
		}
	}
}
