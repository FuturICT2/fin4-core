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
	rg.GET("/assets/:assetId", GetAsset(sc))
	rg.GET("/assets", GetAssets(sc))
	rg.POST("/assets", authenticator, CreateAsset(sc))
	rg.POST("/assets/:assetId/toggle-favorite", authenticator, ToggleFavoriteAsset(sc))

	// Claims APIs
	rg.POST("/asset-claims", authenticator, CreateAssetBlock(sc))
	rg.POST("/claim-verify", authenticator, VerifyAssetBlock(sc))
	rg.POST("/claim/:claimId/toggle-favorite", authenticator, ToggleFavoriteBlock(sc))
}
