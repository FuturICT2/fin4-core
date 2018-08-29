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

// Token token struct
type Token struct {
	ID             int     `json:"Id"`
	Name           string  `json:"Name"`
	Symbol         string  `json:"Symbol"`
	TotalSupply    string  `json:"TotalSupply"`
	Logo           string  `json:"Logo"`
	Decimals       int     `json:"Decimals"`
	FavoritesCount int     `json:"FavoritesCount"`
	Volume24       float32 `json:"Volume24"`
	Change24       float32 `json:"Change24"`
}

// TokensList handler to return json of existing tokens
func (env *Env) TokensList(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		Count      int
		Limit      int
		Page       int
		ValueInUSD string
		Entries    []Token
	}{
		Count:      3,
		Limit:      10,
		Page:       0,
		ValueInUSD: "22454",
		Entries: []Token{
			Token{
				ID:             1,
				Name:           "Fin4",
				Symbol:         "FIN",
				TotalSupply:    "100000000",
				Logo:           "https://en.bitcoin.it/w/images/en/2/29/BC_Logo_.png",
				FavoritesCount: 34,
				Volume24:       10000,
				Change24:       2.3,
				Decimals:       8,
			},
			Token{
				ID:             4,
				Name:           "FinDevCoin",
				Symbol:         "FDC",
				TotalSupply:    "21000000",
				Logo:           "https://en.bitcoin.it/w/images/en/2/29/BC_Logo_.png",
				FavoritesCount: 34,
				Volume24:       3434,
				Change24:       1.3,
				Decimals:       8,
			},
			Token{
				ID:             2,
				Name:           "TreeCoin",
				Symbol:         "TCN",
				TotalSupply:    "12100000",
				Logo:           "https://en.bitcoin.it/w/images/en/2/29/BC_Logo_.png",
				FavoritesCount: 3,
				Volume24:       1234,
				Change24:       -1.3,
				Decimals:       18,
			},
			Token{
				ID:             3,
				Name:           "CyclingCoin",
				Symbol:         "CCN",
				TotalSupply:    "342100000",
				Logo:           "https://en.bitcoin.it/w/images/en/2/29/BC_Logo_.png",
				FavoritesCount: 534,
				Volume24:       9234,
				Change24:       10.5,
				Decimals:       18,
			},
		},
	})
}

// Position token struct
type Position struct {
	TokenID    string `json:"TokenID"`
	Symbol     string `json:"Symbol"`
	Name       string `json:"Name"`
	Balance    string `json:"Balance"`
	ValueInUSD string `json:"ValueInUSD"`
}

// Portfolio handler to return json of existing tokens
func (env *Env) Portfolio(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		Count      int
		Limit      int
		Page       int
		ValueInUSD string
		Positions  []Position
	}{
		Count:      2,
		Limit:      10,
		Page:       0,
		ValueInUSD: "22454",
		Positions: []Position{
			Position{
				TokenID:    "1",
				Name:       "Fin4",
				Symbol:     "FIN",
				Balance:    "23.4",
				ValueInUSD: "2344",
			},
			Position{
				TokenID:    "2",
				Name:       "FinDevCoin",
				Symbol:     "FDC",
				Balance:    "3443.4",
				ValueInUSD: "23344",
			},
			Position{
				TokenID:    "3",
				Name:       "TreeCoint",
				Symbol:     "TCN",
				Balance:    "13.4",
				ValueInUSD: "3",
			},
		},
	})
}
