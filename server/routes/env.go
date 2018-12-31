package routes

import (
	"math/big"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/decimaldt"
	"github.com/FuturICT2/fin4-core/server/ethereum"
	"github.com/FuturICT2/fin4-core/server/models"
	"github.com/FuturICT2/fin4-core/server/pkg/filestorage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// Env router env type
type Env struct {
	Ethereum    *ethereum.Ethereum
	DB          *models.DB
	FileStorage *filestorage.Storage
}

// CreateToken creat a token
func (env *Env) CreateToken(c *gin.Context) {
	user := mustGetUser(c)
	logos := []string{
		"https://trello-attachments.s3.amazonaws.com/5b39f1d06f761ae7c1c7d22c/5b9f76d5afa89e794357c6d9/6b63377efe7ee3d2e6d1b2af22ca51b7/coin1.png",
		"https://trello-attachments.s3.amazonaws.com/5b39f1d06f761ae7c1c7d22c/5b9f76d5afa89e794357c6d9/8b0075fcddb11360a8bbd9bc2673b73f/coin3.png",
		"https://trello-attachments.s3.amazonaws.com/5b39f1d06f761ae7c1c7d22c/5b9f76d5afa89e794357c6d9/f7ff863ac3618c3c9d7b1d0ad7ad41e0/coin4.png",
		"https://trello-attachments.s3.amazonaws.com/5b39f1d06f761ae7c1c7d22c/5b9f76d5afa89e794357c6d9/df4703eeb10e3fbcf1bd63288281cf9d/coin2.png",
	}
	body := struct {
		Name        string `json:"name"`
		Symbol      string `json:"symbol"`
		Purpose     string `json:"purpose"`
		TotalSupply string `json:"totalSupply"`
	}{}
	c.BindJSON(&body)
	totalSupply := new(big.Int)
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
		_, ok := totalSupply.SetString(body.TotalSupply, 10)
		if !ok || totalSupply.Int64() < 0 {
			c.JSON(http.StatusBadRequest, "TotalSupply must be a positive valid int")
			return
		}
	}
	address, tx, err := env.Ethereum.DeployMintable(
		body.Name,
		body.Symbol,
		8,
		common.HexToAddress(user.EthereumAddress),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	rand := rand.Intn(len(logos))
	a, err := env.DB.NewUserModel().InsertToken(
		user.ID, body.Name, body.Symbol, body.Purpose,
		body.TotalSupply, address.Hash().Hex(), tx.Hash().Hex(), logos[rand])
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, a)
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
	ethereumAddress, err := env.Ethereum.CreateNewAddress()
	if err != nil {
		c.String(http.StatusInternalServerError, "We are not able to create your ethereum address, please try again.")
	}
	user, err := userModel.Register(body.Name, ethereumAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	auth.Login(c, user)
	c.JSON(http.StatusOK, user)
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

// CreateAction creates an action
func (env *Env) CreateAction(c *gin.Context) {
	user := mustGetUser(c)
	body := struct {
		Description string `json:"description"`
		TimeLimit   string `json:"timeLimit"`
	}{}
	c.BindJSON(&body)
	if len(body.Description) < 1 || len(body.Description) > 10000 {
		c.String(http.StatusBadRequest, "Description length should be between than 1 and 10000 characters")
		return
	}
	userModel := env.DB.NewUserModel()
	now := time.Now()
	timeLimit, err := strconv.ParseFloat(body.TimeLimit, 64)
	if err != nil || timeLimit < 0 || timeLimit > 48 {
		c.String(http.StatusBadRequest, "Time limit should be a positive valid number and less than 48hrs")
		return
	}
	err = userModel.InsertAction(
		user.ID,
		body.Description,
		now,
		now.Add(time.Duration(timeLimit*60*60*1000000000)),
	)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
}

// AddSupportToAction users can add tokens to actions to increase incentives
func (env *Env) AddSupportToAction(c *gin.Context) {
	user := mustGetUser(c)
	body := struct {
		ActionID models.ID         `json:"actionId"`
		TokenID  models.ID         `json:"tokenId"`
		Amount   decimaldt.Decimal `json:"amount"`
	}{}
	c.BindJSON(&body)
	userModel := env.DB.NewUserModel()
	err := userModel.ReserveRewardsForAction(
		user.ID,
		body.TokenID,
		body.ActionID,
		body.Amount,
	)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
}

// UserLogout logout route
func (env *Env) UserLogout(c *gin.Context) {
	auth.Logout(c)
	c.JSON(http.StatusOK, gin.H{})
}
