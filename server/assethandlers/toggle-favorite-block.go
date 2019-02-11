package assethandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/dbservice"
	"github.com/gin-gonic/gin"
)

//ToggleFavoriteBlock toggles favorite block
func ToggleFavoriteBlock(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := auth.MustGetUser(c)
		blockID, isOk := dbservice.StringToID(c.Params.ByName("blockId"))
		if !isOk {
			c.String(http.StatusBadRequest, "Block id is not valid")
			return
		}
		err := sc.AssetService.ToggleFavoriteBlock(u, blockID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
