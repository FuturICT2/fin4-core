package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/FuturICT2/fin4-core/server/appurl"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/emailer"
)

// PasswordRequestNew starts a new password reset procedure
func PasswordRequestNew(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Email string `json:"email"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.String(http.StatusBadRequest, "Please enter a valid email!")
			return
		}
		token, err := sc.UserService.NewPassResetTokenByEmail(body.Email)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		urlFragment := fmt.Sprintf(
			"/#forgot-reset-password/%d/%s",
			token.UserID,
			token.Token,
		)
		url := appurl.URL(urlFragment)
		_, err = emailer.DefaultEmailer.SendEmailHTML(
			body.Email,
			"Reset your password",
			fmt.Sprintf(`We received a request to reset the password for your account.
			<br><br>
			To reset your password, click the following link:
			<br><br> <a href="%s">%s</a>
			<br><br>
			If didn't initiate this request, then please ignore this email.`, url, url),
			fmt.Sprintf("Click the following link to reset your password %s", url),
		)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
				"email": body.Email,
			}).Error("could not send reset password email")
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}

// PasswordRequestReset starts a new password reset procedure
func PasswordRequestReset(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			UserID   datatype.ID `json:"userId"`
			Token    string      `json:"token"`
			Password string      `json:"password"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.String(http.StatusBadRequest, "Please enter a valid email!")
			return
		}
		err := sc.UserService.ConfirmPassResetToken(
			body.UserID,
			body.Token,
			body.Password,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

//ChangePassword changes logged in user password
func ChangePassword(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			NewPassword     string `json:"newPassword"`
			CurrentPassword string `json:"password"`
		}{}
		c.BindJSON(&body)
		err := sc.UserService.ChangePassword(
			user, body.NewPassword, body.CurrentPassword,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
