package tokenhandlers

import (
	"fmt"
	"net/http"

	"github.com/FuturICT2/fin4-core/server/appstrings"
	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/filestorage"
	"github.com/FuturICT2/fin4-core/server/img"
	"github.com/gin-gonic/gin"
	"github.com/lytics/logrus"
)

// CreateClaim creat a mining claim
func CreateClaim(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			Proposal string `json:"proposal"`
			TokenID  int    `json:"tokenId"`
			Image64  string `json:"image64"`
		}{}
		c.BindJSON(&body)
		var logoPath string
		if body.Image64 != "" {
			imgData, contentType, ext, err := img.FromBase64(body.Image64)
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
				logrus.WithFields(logrus.Fields{
					"error": err.Error(),
				}).Error("actions/create-new-claim:3")
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			logoPath = "https://s3.amazonaws.com/anychange/" + logoPath
		}
		if (len(body.Proposal) < 1 || len(body.Proposal) > 10000) && body.Image64 == "" {
			c.String(http.StatusBadRequest, "Claim should have an image or a text with length between than 1 and 10000 characters")
			return
		}
		err := sc.TokenService.InsertClaim(
			user.ID,
			body.Proposal,
			datatype.ID(body.TokenID),
			logoPath,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}
