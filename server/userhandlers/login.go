package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

//Login login route
func Login(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
		c.BindJSON(&body)
		user, err := sc.UserService.FindByEmail(body.Email)
		if err != nil || user == nil || !user.IsPassword(body.Password) {
			c.String(http.StatusBadRequest, "Invalid login")
			return
		}
		auth.Login(c, user)
		c.JSON(http.StatusOK, user)
	}
}
