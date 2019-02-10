package assethandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

// VerifyAssetBlock creat a new asset
func VerifyAssetBlock(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			BlockID    datatype.ID `json:"blockId"`
			IsAccepted bool        `json:"isAccepted"`
		}{}
		c.BindJSON(&body)
		status := 2
		if body.IsAccepted {
			status = 1
		}
		err := sc.AssetService.VerifyAssetBlock(
			&sc,
			user,
			datatype.ID(body.BlockID),
			status,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}
