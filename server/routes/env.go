package routes

import (
	"math/big"
	"math/rand"
	"net/http"
	"time"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/ethereum"
	"github.com/FuturICT2/fin4-core/server/models"
	"github.com/gin-gonic/gin"
)

// Env router env type
type Env struct {
	Ethereum *ethereum.Ethereum
	DB       *models.DB
}

// CreateToken creat a token
func (env *Env) CreateToken(c *gin.Context) {
	logos := []string{
		"https://trello-attachments.s3.amazonaws.com/5b39f1d06f761ae7c1c7d22c/5b9f76d5afa89e794357c6d9/6b63377efe7ee3d2e6d1b2af22ca51b7/coin1.png",
		"https://trello-attachments.s3.amazonaws.com/5b39f1d06f761ae7c1c7d22c/5b9f76d5afa89e794357c6d9/8b0075fcddb11360a8bbd9bc2673b73f/coin3.png",
		"https://trello-attachments.s3.amazonaws.com/5b39f1d06f761ae7c1c7d22c/5b9f76d5afa89e794357c6d9/f7ff863ac3618c3c9d7b1d0ad7ad41e0/coin4.png",
		"https://trello-attachments.s3.amazonaws.com/5b39f1d06f761ae7c1c7d22c/5b9f76d5afa89e794357c6d9/df4703eeb10e3fbcf1bd63288281cf9d/coin2.png",
	}
	time.Sleep(1000 * time.Millisecond)
	body := struct {
		Name        string   `json:"name"`
		Symbol      string   `json:"symbol"`
		Purpose     string   `json:"purpose"`
		TotalSupply string   `json:"totalSupply"`
		Hashtags    []string `json:"hashtags"`
	}{}
	c.BindJSON(&body)
	totalSupply := new(big.Int)
	totalSupply, ok := totalSupply.SetString(body.TotalSupply, 10)
	if !ok {
		c.JSON(http.StatusFailedDependency, models.ErrServerError)
	}
	address, tx, err := env.Ethereum.DeployNewToken(totalSupply, body.Name, 0, body.Symbol)
	if err != nil {
		c.JSON(http.StatusFailedDependency, models.ErrServerError)
	}
	rand := rand.Intn(len(logos))
	a, x := env.DB.NewUserModel().InsertToken(
		models.ID(1), body.Name, body.Symbol, body.Purpose,
		body.TotalSupply, address.Hash().Hex(), tx.Hash().Hex(), logos[rand])
	if x != nil {
		c.JSON(http.StatusFailedDependency, models.ErrServerError)
	}
	c.JSON(http.StatusOK, a)
}

// UserLogin login route
func (env *Env) UserLogin(c *gin.Context) {
	body := struct {
		Name string `json:"name"`
	}{}
	c.BindJSON(&body)
	userModel := env.DB.NewUserModel()
	user, err := userModel.Register(body.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	auth.Login(c, user)
	c.JSON(http.StatusOK, user)
}
