package userhandlers

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/routermiddleware"
	"github.com/gin-gonic/gin"
)

//InjectHandlers injects asset handlers into the application
func InjectHandlers(sc datatype.ServiceContainer, rg *gin.RouterGroup) {
	authenticator := routermiddleware.SessionMustAuth()
	rg.POST("/register", Register(sc))
	rg.POST("/login", Login(sc))
	rg.POST("/logout", Logout(sc))
	rg.GET("/session", authenticator, SessionGet(sc))
	rg.DELETE("/session", authenticator, SessionDestroy(sc))
	rg.POST("/forgotpass-requests/new", PasswordRequestNew(sc))
	rg.POST("/forgotpass-requests/reset", PasswordRequestReset(sc))
	rg.POST("/user/passwords", authenticator, ChangePassword(sc))
	rg.POST("/user/email", authenticator, ChangeEmail(sc))
	rg.GET("/user/email/confirm", ConfirmChangeEmail(sc))
	rg.GET("/balances", authenticator, Balances(sc))
}
