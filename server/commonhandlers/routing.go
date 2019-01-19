package commonhandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/FuturICT2/fin4-core/server/datatype"
)

//InjectHandlers injects asset handlers into the application
func InjectHandlers(sc datatype.ServiceContainer, rg *gin.RouterGroup) {
	//authenticator := middleware.SessionMustAuth()
	rg.GET("/", Index(sc))
}
