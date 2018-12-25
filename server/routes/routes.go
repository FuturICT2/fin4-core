package routes

import (
	"github.com/FuturICT2/fin4-core/server/routes/middleware"
	"github.com/gin-gonic/gin"
)

// StartRouter sets up routes
func (env *Env) StartRouter() *gin.Engine {
	r := gin.Default()
	r.Static("./assets", "./public")
	// html
	web := r.Group("/")
	{
		web.GET("/", env.Index)
	}
	// website specific api
	wapi := r.Group("/wapi")
	wapi.Use(middleware.Session())
	wapi.Use(middleware.SessionSetUser(env.DB))
	wapi.Use(middleware.CheckCsrfToken())
	{
		mustAuth := middleware.SessionMustAuth()
		wapi.GET("/csrf", middleware.SetCsrfToken())
		wapi.GET("/session", mustAuth, env.SessionGet)
		wapi.POST("/login", env.UserLogin)
		wapi.POST("/logout", mustAuth, env.UserLogout)
		wapi.GET("/tokens", env.TokensList)
		wapi.GET("/people", env.PeopleList)
		wapi.GET("/portfolio/positions", mustAuth, env.Portfolio)
		wapi.POST("/create-token", mustAuth, env.CreateToken)
		wapi.GET("/like/:tokenID", mustAuth, env.DoLike)
		wapi.GET("/unlike/:tokenID", mustAuth, env.DoUnLike)
		wapi.POST("/actions", mustAuth, env.CreateAction)
		wapi.GET("/actions", env.ActionsList)
		wapi.POST("/submit-proposal", mustAuth, env.AddActionProposal)
		// wapi.POST("/approve-proposal", mustAuth, env.AprroveProposal)
		wapi.POST("/add-rewards-to-action", mustAuth, env.AddSupportToAction)
	}

	return r
}
