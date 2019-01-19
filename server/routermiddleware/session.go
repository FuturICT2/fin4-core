package routermiddleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

// Session starts session
func Session() gin.HandlerFunc {
	store := sessions.NewCookieStore([]byte("|FQyu8eMGdCx2D7h!c6qBKnhgAv8jQmlWVgaf&MCamrylarK!scXYwkeeL'MuQlq"))
	return sessions.Sessions("session", store)
}

// SessionSetUser sets user on context
func SessionSetUser(userService datatype.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userIDRaw := session.Get("userId")
		var user *datatype.User
		switch userIDStr := userIDRaw.(type) {
		case string:
			if userID, err := strconv.Atoi(userIDStr); err == nil {
				user, _ = userService.FindByID(datatype.ID(userID))
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
		user := c.MustGet("user").(*datatype.User)
		if user == nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, nil)
		}
	}
}
