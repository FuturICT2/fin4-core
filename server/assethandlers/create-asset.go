package assethandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// CreateAllPurposeAsset create a new configured asset
func CreateAllPurposeAsset(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			Name           string `json:"name"`
			Purpose        string `json:"purpose"`
			Symbol         string `json:"symbol"`
			Decimals       uint8  `json:"decimals"`
			Cap            uint64 `json:"cap"`
			IsBurnable     bool   `json:"isBurnable"`
			IsTransferable bool   `json:"isTransferable"`
			IsMintable     bool   `json:"isMintable"`
		}{}
		c.BindJSON(&body)
		add, tx, err := sc.Ethereum.DeployAllPurpose(
			body.Name,
			body.Symbol,
			body.Decimals,
			common.HexToAddress(user.EthereumAddress),
			body.IsBurnable,
			body.IsTransferable,
			body.IsMintable,
			body.Cap,
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

// CreateAsset create a new mintable, burnable, transferable, uncapped asset
func CreateAsset(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			Name    string `json:"name"`
			Purpose string `json:"purpose"`
			Symbol  string `json:"symbol"`
		}{}
		c.BindJSON(&body)
		add, tx, err := sc.Ethereum.DeployAllPurpose(
			body.Name,
			body.Symbol,
			8,
			common.HexToAddress(user.EthereumAddress),
			true,
			true,
			true,
			0,
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
