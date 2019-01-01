package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/FuturICT2/fin4-core/server/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// TokensList handler to return json of existing tokens
func (env *Env) TokensList(c *gin.Context) {
	user := mustGetUser(c)
	userModel := env.DB.NewUserModel()
	log.Println("---------------------------", user.ID)
	tokens, err := userModel.FindTokens(user.ID)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	users, err := userModel.FindUsers()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, struct {
		Count   int
		Limit   int
		Page    int
		Entries []models.Token
		People  []models.User
	}{
		Count:   len(tokens),
		Limit:   10,
		Page:    0,
		Entries: tokens,
		People:  users,
	})
}

// CreateToken creat a token
func (env *Env) CreateToken(c *gin.Context) {
	user := mustGetUser(c)
	body := struct {
		Name    string `json:"name"`
		Symbol  string `json:"symbol"`
		Purpose string `json:"purpose"`
	}{}
	c.BindJSON(&body)
	{ // verify data formats
		if len(body.Name) < 3 || len(body.Name) > 35 {
			c.String(http.StatusBadRequest, "Name length should be between than 3 and 35 characters")
			return
		}
		if len(body.Purpose) < 3 || len(body.Purpose) > 255 {
			c.String(http.StatusBadRequest, "Purpose length should be between than 3 and 255 characters")
			return
		}
		if body.Symbol != "" && (len(body.Symbol) > 4 || len(body.Symbol) < 3) {
			c.String(http.StatusBadRequest, "Symbol length should be 3 or 4 characters")
			return
		}
	}
	address, tx, err := env.Ethereum.DeployMintable(
		body.Name,
		body.Symbol,
		8,
		common.HexToAddress(user.EthereumAddress),
	)
	// @TODO MySQL caching of the tokens should be done in a synching process
	// so we can get token contract number of confirmations in our db
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userModel := env.DB.NewUserModel()
	a, err := userModel.InsertToken(
		user.ID,
		body.Name,
		body.Symbol,
		body.Purpose,
		"0", // total supply
		address.Hex(),
		tx.Hash().Hex(),
		"", // logo file
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, a)
}

// DoLike like from logged in user to the passed token
func (env *Env) DoLike(c *gin.Context) {
	user := mustGetUser(c)
	tokenID, err := strconv.Atoi(c.Params.ByName("tokenID"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	tm := env.DB.NewUserModel()
	err = tm.DoLike(user.ID, models.ID(tokenID), true)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, struct{}{})
}

// DoUnLike unlike from logged in user to the passed token
func (env *Env) DoUnLike(c *gin.Context) {
	user := mustGetUser(c)
	tokenID, err := strconv.Atoi(c.Params.ByName("tokenID"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	tm := env.DB.NewUserModel()
	err = tm.DoLike(user.ID, models.ID(tokenID), false)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, struct{}{})
}
