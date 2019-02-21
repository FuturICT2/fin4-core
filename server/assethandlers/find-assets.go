package assethandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

// FindAssets fetch assets from db
func FindAssets(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		assets, err := sc.AssetService.FindAll(user)
		if err != nil {
			c.String(http.StatusServiceUnavailable, err.Error())
			return
		}
		c.JSON(http.StatusOK, toAssetsResponse(assets))
	}
}
