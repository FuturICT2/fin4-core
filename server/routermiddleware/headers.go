package routermiddleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HeadersNoCache adds no cache headers to the request
func HeadersNoCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		c.Next()
	}
}

// HeadersCors enables CORS
func HeadersCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.Header("Content-Type", "application/json")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
			c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.AbortWithStatus(200)
		}
	}
}
