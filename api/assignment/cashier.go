package assignment

import (
	"net/http"
	"scg-assignment/internal/cashier"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

func (s *server) change(ctx echo.Context) error {
	var req changeMoneyReq

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	// convert map[decimal.Decimal]int to map[float64]int
	moneyForBuy := make(cashier.MoneyMap)
	for v, a := range req.MoneyForBuy {
		f, _ := v.Float64()
		moneyForBuy[f] = a
	}

	change, err := s.cashier.Change(req.Price, moneyForBuy)
	if err != nil {
		return err
	}

	changeDec := moneyMapToMoneyMapDec(change)

	return ctx.JSON(http.StatusOK, changeDec)
}

func (s *server) addMoney(ctx echo.Context) error {
	var req struct {
		Value  float64 `json:"value"`
		Amount int     `json:"amount"`
	}

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := s.cashier.AddMoney(req.Value, req.Amount); err != nil {
		return err
	}

	return nil
}

func (s *server) getMoneyStock(ctx echo.Context) error {
	money := s.cashier.GetMoneyStock()

	moneyDec := moneyMapToMoneyMapDec(money)

	return ctx.JSON(http.StatusOK, moneyDec)
}

type moneyMapDec map[decimal.Decimal]int

type changeMoneyReq struct {
	Price       decimal.Decimal `json:"price"`
	MoneyForBuy moneyMapDec     `json:"moneyForBuy"`
}

// change map[float64]int to map[decimal.Decimal]int because we can't marshal json from map[float64]int
func moneyMapToMoneyMapDec(m cashier.MoneyMap) moneyMapDec {
	m2 := make(moneyMapDec)
	for v, a := range m {
		d := decimal.NewFromFloat(v)
		m2[d] = a
	}

	return m2
}
