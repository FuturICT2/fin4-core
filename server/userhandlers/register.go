package userhandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/appstrings"
	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

//Register register route
func Register(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := struct {
			Email        string `json:"email"`
			Name         string `json:"name"`
			Password     string `json:"password"`
			AgreeToTerms bool   `json:"agreeToTerms"`
			IsFastSignup bool   `json:"isFastSignup"`
		}{}
		c.BindJSON(&body)
		if body.IsFastSignup {
			body.Password = appstrings.GetRandomString(8)
			body.AgreeToTerms = true
		}
		ethereumAddress, err := sc.Ethereum.CreateNewAddress()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		user, err := sc.UserService.Register(
			body.Email,
			body.Name,
			ethereumAddress,
			body.Password,
			body.AgreeToTerms,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		auth.Login(c, user)
		c.JSON(http.StatusOK, user)
	}
}
