package userhandlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/FuturICT2/fin4-core/server/appurl"
	"github.com/FuturICT2/fin4-core/server/auth"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/emailer"
)

//ChangeEmail initiates logged in user email change process
func ChangeEmail(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			Email           string `json:"email"`
			CurrentPassword string `json:"password"`
		}{}
		c.BindJSON(&body)
		conf, err := sc.UserService.CreateChangeEmailConfirmation(
			user,
			body.Email,
			body.CurrentPassword,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		url := appurl.URL(fmt.Sprintf(
			"/wapi/user/email/confirm?userid=%d&token=%s",
			user.ID,
			conf.Token,
		))
		_, err = emailer.DefaultEmailer.SendEmailHTML(
			body.Email,
			"Verify your email address",
			fmt.Sprintf(`To change your email address please click the following link
			to verify your email change request:<br> <a href="%s">%s</a>`, url, url),
			fmt.Sprintf("Open the foloowing link \n\n%s\n\n to verify your email change", url),
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

//ConfirmChangeEmail confirms email change & updates user's email
func ConfirmChangeEmail(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDstr, exists := c.GetQuery("userid")
		redirect := func() {
			c.Redirect(http.StatusMovedPermanently, "/#account-noti/email-change-fail")
		}
		if !exists || userIDstr == "" {
			redirect()
			return
		}
		token, exists := c.GetQuery("token")
		if !exists || token == "" {
			redirect()
			return
		}
		userID, err := strconv.Atoi(userIDstr)
		if err != nil {
			redirect()
			return
		}
		success := sc.UserService.ConfirmUserEmailChange(
			datatype.ID(userID),
			token,
		)
		if !success {
			redirect()
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/#account-noti/email-change-success")
	}
}
