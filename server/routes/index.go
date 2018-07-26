package routes

import "github.com/gin-gonic/gin"

// Index index route
func (env *Env) Index(c *gin.Context) {
	c.File("public/index.html")
}
