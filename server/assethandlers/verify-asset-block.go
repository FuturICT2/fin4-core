package assethandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/dbservice"
	"github.com/gin-gonic/gin"
)

// VerifyAssetBlock creat a new asset
func VerifyAssetBlock(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			IsAccepted bool `json:"isAccepted"`
		}{}
		c.BindJSON(&body)
		blockID, isOk := dbservice.StringToID(c.Params.ByName("blockId"))
		if !isOk {
			c.String(http.StatusBadRequest, "Block id is not valid")
			return
		}
		status := 2
		if body.IsAccepted {
			status = 1
		}
		err := sc.AssetService.VerifyAssetBlock(
			&sc,
			user,
			blockID,
			status,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}
