package userhandlers

import (
	"net/http"
	"strconv"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/gin-gonic/gin"
)

// GetPerson return person information
func GetPerson(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		personID, err := strconv.Atoi(c.Params.ByName("personID"))
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}
		person, err := sc.UserService.GetPerson(datatype.ID(personID))
		if err != nil {
			c.String(http.StatusServiceUnavailable, err.Error())
			return
		}
		c.JSON(http.StatusOK, person)
	}
}
