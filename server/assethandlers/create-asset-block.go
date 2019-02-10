package assethandlers

import (
	"fmt"
	"net/http"

	"github.com/FuturICT2/fin4-core/server/apperrors"
	"github.com/FuturICT2/fin4-core/server/appstrings"
	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/filestorage"
	"github.com/FuturICT2/fin4-core/server/img"
	"github.com/gin-gonic/gin"
)

// CreateAssetBlock creat a new asset
func CreateAssetBlock(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			AssetID   datatype.ID `json:"assetId"`
			BlockText string      `json:"blockText"`
			Images    []struct {
				Contents string `json:"contents"`
				FileName string `json:"filename"`
			}
		}{}
		c.BindJSON(&body)
		if len(body.Images) > 3 {
			c.String(http.StatusBadRequest, "Post images are limited to 3 images")
			return
		}
		var logoPaths []string
		{ // upload photos
			for _, image := range body.Images {
				var logoPath string
				if image.Contents != "" {
					imgData, contentType, ext, err := img.FromBase64(image.Contents)
					logoPath = fmt.Sprintf(
						"tokenlogos/%d-%s.%s",
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
						apperrors.Critical("assethandlers:CreateAssetBlock:1", err)
						c.String(http.StatusBadRequest, err.Error())
						return
					}
					logoPath = "https://s3.amazonaws.com/anychange/" + logoPath
				}
				logoPaths = append(logoPaths, logoPath)
			}
		} // end of uploading images
		block, err := sc.AssetService.CreateAssetBlock(
			user.ID,
			body.AssetID,
			body.BlockText,
			logoPaths,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, block)
	}
}
