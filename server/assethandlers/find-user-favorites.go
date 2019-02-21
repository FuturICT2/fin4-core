package assethandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

//GetUserFavoriteAssets get user's favorite assets
func GetUserFavoriteAssets(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := auth.MustGetUser(c)
		assets, err := sc.AssetService.FindUserFavoriteAssets(u)
		if err != nil {
			c.String(http.StatusServiceUnavailable, err.Error())
			return
		}
		c.JSON(http.StatusOK, toAssetsResponse(assets))
	}
}
