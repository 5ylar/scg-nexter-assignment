package assignment

import (
	"net/http"
	"scg-assignment/internal/cashier"
	"scg-assignment/internal/search"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type searcher interface {
	SearchMultipleKeysFromStringDataset(keys []string, dataset []string, caseSensitive bool) []search.ResultFromStringDataset
}

type cashierer interface {
	Change(price decimal.Decimal, moneyForBuy cashier.MoneyMap) (cashier.MoneyMap, error)
	AddMoney(value float64, amount int) error
	GetMoneyStock() cashier.MoneyMap
}

type server struct {
	search  searcher
	cashier cashierer
}

func New(search searcher, cashier cashierer) *server {
	return &server{
		search,
		cashier,
	}
}

func (s *server) Start(address string) error {

	e := echo.New()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	gs := e.Group("/search")
	gc := e.Group("/cashier")

	gs.POST("/multiple", s.multipleSearchFromDataset)

	gc.POST("/change", s.change)
	gc.GET("/money", s.getMoneyStock)
	gc.POST("/money", s.addMoney)

	return e.Start(address)
}
