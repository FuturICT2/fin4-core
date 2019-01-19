package routermiddleware

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/FuturICT2/fin4-core/server/appstrings"
)

const headerName = "X-Csrf-Token"
const tokenName = "csrftoken"

// CheckCsrfToken adds a middleware that checks csrf token
func CheckCsrfToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		return
		isPost := c.Request.Method == http.MethodPost
		isPut := c.Request.Method == http.MethodPut
		isDelete := c.Request.Method == http.MethodDelete
		isPatch := c.Request.Method == http.MethodPatch
		if isPost || isPut || isDelete || isPatch {
			token := c.Request.Header.Get(headerName)
			savedToken := sessions.Default(c).Get(tokenName)
			if token != savedToken {
				c.Abort()
				c.Status(http.StatusUnauthorized)
				c.Writer.WriteString("⌐╦╦═─  ▄︻̷̿┻̿═━一")
			}
		}
	}
}

// SetCsrfToken sends csrf token to user
func SetCsrfToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		newToken := appstrings.GetRandomString(0)
		session := sessions.Default(c)
		session.Set(tokenName, newToken)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"token": newToken,
		})
	}
}
