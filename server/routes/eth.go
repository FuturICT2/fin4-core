package routes

// UserGetBalances gets user balances
import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	// ErrServerError for any internal server error
	ErrServerError = errors.New("Server error")
)

// BestBlock route handler that returns best block number of fin4 ethereum node
func (env *Env) BestBlock(c *gin.Context) {
	blockNumber, err := env.Ethereum.GetBlockNumber()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("server:eth:e1")
		c.String(http.StatusFailedDependency, ErrServerError.Error())
		return
	}
	c.JSON(http.StatusOK, blockNumber)
}
