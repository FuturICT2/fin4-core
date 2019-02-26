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
	// Find listed assets
	rg.GET("/assets", FindAssets(sc))
	// Create new asset
	rg.POST("/assets", authenticator, CreateAsset(sc))
	// Find asset by its ID
	rg.GET("/assets/:assetId", FindAsset(sc))
	// Toggle user favorite asset
	rg.POST("/assets/:assetId/toggle-favorite", authenticator, ToggleFavoriteAsset(sc))

	// Oracle APIs
	// Create new block in an asset
	rg.POST("/asset-blocks", authenticator, CreateAssetBlock(sc))
	rg.POST("/asset-blocks/:blockId/verify", authenticator, VerifyAssetBlock(sc))
	rg.POST("/asset-blocks/:blockId/toggle-favorite", authenticator, ToggleFavoriteBlock(sc))
	// a hack api to demonstrate sensor verification
	// @TODO extend the APIs to enable sensors or app registrations and application specific APIs
	rg.POST("/asset-block/sensor-verify", SensorVerifyBlock(sc))
	rg.POST("/oracle-ping", OraclePing(sc))
}
