package assethandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

// OraclePing updates last ping for the oracle
func OraclePing(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := struct {
			AccessToken string `json:"accessToken"`
		}{}
		c.BindJSON(&body)
		err := sc.AssetService.UpdateOraclePingTime(body.AccessToken)
		if err != nil {
			c.String(http.StatusServiceUnavailable, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
