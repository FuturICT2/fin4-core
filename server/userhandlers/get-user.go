package userhandlers

import (
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/dbservice"
	"github.com/gin-gonic/gin"
)

// GetUser handle get profile route
func GetUser(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		userID, isOk := dbservice.StringToID(c.Params.ByName("userID"))
		if !isOk {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}
		profileUser, err := sc.UserService.FindByID(userID)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		isOwner := user.ID == userID
		c.JSON(http.StatusOK, struct {
			ID              datatype.ID
			UserName        string
			IsOwner         bool
			ProfileImageURL string
		}{
			profileUser.ID,
			profileUser.Name,
			isOwner,
			profileUser.ProfileImageURL,
		})
	}
}
