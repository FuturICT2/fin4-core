package routes

import (
	"math/big"
	"net/http"
	"time"

	"github.com/FuturICT2/fin4-core/server/ethereum"
	"github.com/gin-gonic/gin"
)

// Env router env type
type Env struct {
	Ethereum *ethereum.Ethereum
}

// CreateToken creat a token
func (env *Env) CreateToken(c *gin.Context) {
	time.Sleep(2000 * time.Millisecond)
	body := struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Category    string   `json:"category"`
		Shares      string   `json:"shares"`
		Hashtags    []string `json:"hashtags"`
	}{}
	c.BindJSON(&body)
	totalSupply := new(big.Int)
	totalSupply, _ = totalSupply.SetString(body.Shares, 10)
	address, _, _ := env.Ethereum.DeployNewToken(totalSupply, body.Name, 0, body.Name[0:3])
	// market, err := tradeModel.InsertAsset(
	// 	body.Name,
	// 	body.Description,
	// 	body.Shares,
	// 	body.Category,
	// 	body.Hashtags,
	// )
	// if err != nil {
	// 	c.String(http.StatusBadRequest, err.Error())
	// 	return
	// }
	c.JSON(http.StatusOK, address)
}
