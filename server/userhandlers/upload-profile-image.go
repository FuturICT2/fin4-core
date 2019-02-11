package userhandlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FuturICT2/fin4-core/server/apperrors"
	"github.com/FuturICT2/fin4-core/server/appstrings"
	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/filestorage"
	"github.com/FuturICT2/fin4-core/server/img"
	"github.com/gin-gonic/gin"
)

// UploadProfileImage creat a new asset
func UploadProfileImage(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			Image string `json:"image"`
		}{}
		c.BindJSON(&body)
		if len(body.Image) == 0 {
			c.String(http.StatusBadRequest, "Missing image")
			return
		}
		var logoPath string
		if body.Image != "" {
			imgData, contentType, ext, err := img.FromBase64(body.Image)
			logoPath = fmt.Sprintf(
				"profile-pictures/%d-%s.%s",
				user.ID,
				appstrings.GetRandomString(20),
				ext,
			)
			err = sc.FileStorage.Put(
				logoPath,
				contentType,
				imgData,
				filestorage.AclPublicRead,
			)
			if err != nil {
				apperrors.Critical("assethandlers:upload-profile-image:1", err)
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			logoPath = "https://s3.amazonaws.com/anychange/" + logoPath
		}
		log.Println("---------------------------", logoPath)
		err := sc.UserService.UpdateProfileImage(user, logoPath)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
