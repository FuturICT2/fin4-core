package assethandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

// SensorVerifyBlock for f2forum demo sensor hack API
func SensorVerifyBlock(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := struct {
			IsAccepted  bool   `json:"isAccepted"`
			AccessToken string `json:"accessToken"`
		}{}
		c.BindJSON(&body)
		status := 2
		if body.IsAccepted {
			status = 1
		}
		err := sc.AssetService.SensorVerifyBlock(
			&sc,
			status,
			body.AccessToken,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}
