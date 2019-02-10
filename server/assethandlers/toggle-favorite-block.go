package assethandlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

//ToggleFavoriteBlock toggles favorite block
func ToggleFavoriteBlock(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := auth.MustGetUser(c)
		body := struct {
			BlockID int
		}{}
		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, errors.New("Bad payload"))
			return
		}
		err = sc.AssetService.ToggleFavoriteBlock(u, datatype.ID(body.BlockID))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
