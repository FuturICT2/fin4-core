package assethandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
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
		asset, err := sc.AssetService.Create(
			sc,
			user.ID,
			body.Name,
			body.Symbol,
			body.Purpose,
			"PLACE_HOLDER_ETH_ADDRESS",
			"PLACE_HOLDER_ETH_TX_ADDRESS",
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, asset)
	}
}
