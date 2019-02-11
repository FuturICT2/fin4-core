package tokenhandlers

import (
	"net/http"
	"strings"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// CreateToken creat a token
func CreateToken(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			Name    string `json:"name"`
			Symbol  string `json:"symbol"`
			Purpose string `json:"purpose"`
		}{}
		c.BindJSON(&body)
		body.Symbol = strings.ToUpper(body.Symbol)
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
		add, tx, err := sc.Ethereum.DeployMintable(
			body.Name,
			body.Symbol,
			8,
			common.HexToAddress(user.EthereumAddress),
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		token, err := sc.TokenService.InsertToken(
			user.ID, // creator ID
			body.Name,
			body.Symbol,
			body.Purpose,
			"0.0", // hard cap
			add.Hex(),
			tx.Hash().Hex(),
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		token.Claims = []datatype.Claim{}
		token.Miners = []datatype.Miner{}
		token.Likers = []datatype.Liker{}
		c.JSON(http.StatusOK, token)
	}
}
