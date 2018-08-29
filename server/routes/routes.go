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
	wapi.Use(middleware.CheckCsrfToken())
	{
		wapi.GET("/csrf", middleware.SetCsrfToken())
		wapi.GET("/tokens", env.TokensList)
		wapi.GET("/portfolio/positions", env.Portfolio)
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
