package assethandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

//ToggleFavoriteAsset toggles favorite asset
func ToggleFavoriteAsset(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := auth.MustGetUser(c)
		body := struct {
			AssetID int
		}{}
		c.BindJSON(&body)
		err := sc.AssetService.ToggleFavorite(u, datatype.ID(body.AssetID))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
