package assethandlers

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/routermiddleware"
	"github.com/gin-gonic/gin"
)

//InjectHandlers injects asset handlers into the application
func InjectHandlers(sc datatype.ServiceContainer, rg *gin.RouterGroup) {
	authenticator := routermiddleware.SessionMustAuth()
	// Assets APIs
	rg.GET("/v2/assets/:assetId", GetAsset(sc))
	rg.GET("/v2/assets", GetAssets(sc))
	rg.POST("/v2/assets", authenticator, CreateAsset(sc))
	rg.POST("/v2/assets/:assetId/toggle-favorite", authenticator, ToggleFavoriteAsset(sc))

	// Claims APIs
	rg.POST("/v2/asset-blocks", authenticator, CreateAssetBlock(sc))
	rg.POST("/v2/asset-block/:blockId/verify", authenticator, VerifyAssetBlock(sc))
	rg.POST("/v2/asset-block/:blockId/toggle-favorite", authenticator, ToggleFavoriteBlock(sc))
}
