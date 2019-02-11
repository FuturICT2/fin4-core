package assethandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// CreateAsset creat a new asset
func CreateAsset(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			Name    string `json:"name"`
			Purpose string `json:"purpose"`
			Symbol  string `json:"symbol"`
		}{}
		c.BindJSON(&body)
		add, tx, err := sc.Ethereum.DeployMintable(
			body.Name,
			body.Symbol,
			8,
			common.HexToAddress(user.EthereumAddress),
		)
		if err != nil {
			c.String(http.StatusServiceUnavailable, err.Error())
			return
		}
		asset, err := sc.AssetService.Create(
			sc,
			user.ID,
			body.Name,
			body.Symbol,
			body.Purpose,
			add.Hex(),
			tx.Hash().Hex(),
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, asset)
	}
}
