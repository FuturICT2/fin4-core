package routes

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/models"
	"github.com/gin-gonic/gin"
)

// Portfolio handler to return json of existing tokens
func (env *Env) Portfolio(c *gin.Context) {
	user := mustGetUser(c)
	userModel := env.DB.NewUserModel()
	balances, err := userModel.GetUserBalances(user.ID)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, struct {
		Count     int
		Limit     int
		Page      int
		Positions []models.Balance
	}{
		Count:     2,
		Limit:     10,
		Page:      0,
		Positions: balances,
	})
}

// PeopleList handler to return json of existing tokens
func (env *Env) PeopleList(c *gin.Context) {
	userModel := env.DB.NewUserModel()
	users, err := userModel.FindUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, struct {
		Count   int
		Limit   int
		Page    int
		Entries []models.User
	}{
		Count:   len(users),
		Limit:   10,
		Page:    0,
		Entries: users,
	})
}

// UserLogin login route
func (env *Env) UserLogin(c *gin.Context) {
	body := struct {
		Name string `json:"name"`
	}{}
	c.BindJSON(&body)
	if len(body.Name) < 2 || len(body.Name) > 35 {
		c.String(http.StatusBadRequest, "Name length should be between than 2 and 35 characters")
		return
	}
	userModel := env.DB.NewUserModel()
	// a unique ethereum address for each user for demonstration puproses
	ethereumAddress, err := env.Ethereum.CreateNewAddress()
	if err != nil {
		c.String(
			http.StatusInternalServerError,
			"We are not able to create your ethereum address, please try again.",
		)
	}
	user, err := userModel.Register(body.Name, ethereumAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	auth.Login(c, user)
	c.JSON(http.StatusOK, user)
}

// UserLogout logout route
func (env *Env) UserLogout(c *gin.Context) {
	auth.Logout(c)
	c.JSON(http.StatusOK, gin.H{})
}

func mustGetUser(c *gin.Context) *models.User {
	return c.MustGet("user").(*models.User)
}
