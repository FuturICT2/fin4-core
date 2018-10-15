package routes

import (
	"github.com/FuturICT2/fin4-core/server/routes/middleware"
	"github.com/gin-gonic/gin"

	api "gopkg.in/appleboy/gin-status-api.v1"
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
		wapi.GET("/tokens", env.TokensList)
		wapi.GET("/people", env.PeopleList)
		wapi.GET("/portfolio/positions", env.Portfolio)
		wapi.POST("/create-token", env.CreateToken)
	}

	// Ethereum specific APIs
	eth := r.Group("/eth")
	eth.Use(middleware.CheckCsrfToken())
	{
		eth.GET("/best-block", env.BestBlock)
	}

	// Core API
	v1 := r.Group("/api")
	v1.Use(middleware.HeadersNoCache())
	v1.Use(middleware.HeadersCors())
	{
		v1.GET("/status", middleware.APIAuth(), api.StatusHandler)
	}
	return r
}
