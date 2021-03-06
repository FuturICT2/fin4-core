package userhandlers

import (
	"math/rand"
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

//Login login route
func Login(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := struct {
			Name string `json:"name"`
		}{}
		c.BindJSON(&body)
		user, err := sc.UserService.FindByEmail(body.Name)
		if err != nil || user == nil {
			ethereumAddress, err := sc.Ethereum.CreateNewAddress()
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			user, err := sc.UserService.Register(
				generateToken(20)+"@finfour.net",
				body.Name,
				ethereumAddress,
				"dummypass",
				true,
			)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			auth.Login(c, user)
			c.JSON(http.StatusOK, user)
		}
		auth.Login(c, user)
		c.JSON(http.StatusOK, user)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateToken(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
