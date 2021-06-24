package cashier

import (
	"errors"
	"sort"
	"sync"

	"github.com/shopspring/decimal"
)

type cashier struct {
	moneyStock MoneyMap
	mu         sync.Mutex
}

func New(ms MoneyMap) *cashier {
	return &cashier{
		moneyStock: ms,
	}
}

func (c *cashier) Change(price decimal.Decimal, moneyForBuy MoneyMap) (MoneyMap, error) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()

	if price.IsZero() || price.IsNegative() {
		return EmptyMoneyMap, ErrInvalidPrice
	}

	var (
		change           MoneyMap = make(MoneyMap)
		totalMoneyStock  decimal.Decimal
		totalMoneyForBuy decimal.Decimal
		totalChange      decimal.Decimal
	)

	// calculate total money
	for v, a := range moneyForBuy {
		totalMoneyForBuy = totalMoneyForBuy.Add(decimal.NewFromFloat(v * float64(a)))

		// check allowd money
		if !checkContainsForFloat(AllowedMoney, v) {
			return EmptyMoneyMap, ErrNotAllowedMoney
		}
	}

	// expected change
	totalChange = totalMoneyForBuy.Sub(price)

	// totalMoneyForBuy < price => error
	if totalChange.IsNegative() {
		return EmptyMoneyMap, ErrNotEnoughMoney
	}

	// find all money values in cashier
	var moneyValues []float64
	for v, a := range c.moneyStock {
		if a > 0 {
			moneyValues = append(moneyValues, v)
			totalMoneyStock = totalMoneyStock.Add(decimal.NewFromFloat(v * float64(a)))
		}
	}

	// sort money values (DESC)
	sort.SliceStable(moneyValues, func(i, j int) bool {
		return moneyValues[i] > (moneyValues[j])
	})

	// copy money stock to tmp variable by value
	tmpMoneyStock := make(MoneyMap)
	for v, a := range c.moneyStock {
		tmpMoneyStock[v] = a
	}

	// add money to money stock (tmp)
	for v := range moneyForBuy {
		tmpMoneyStock[v] += moneyForBuy[v]
	}

	remainingChange := totalChange

	// find money change
REMAIN_CHANGE_LOOP:
	for remainingChange.Cmp(decimal.Zero) == 1 { // remainingChange > 0

		for _, v := range moneyValues {

			// out of money or money value is less than change, skiped
			if tmpMoneyStock[v] == 0 || decimal.NewFromFloat(v).Cmp(remainingChange) == 1 {
				continue
			}

			remainingChange = remainingChange.Sub(decimal.NewFromFloat(v))

			tmpMoneyStock[v]--
			if tmpMoneyStock[v] == 0 {
				delete(tmpMoneyStock, v)
			}

			change[v]++
			continue REMAIN_CHANGE_LOOP
		}

		// can't change money
		return EmptyMoneyMap, ErrCantChangeMoney
	}

	// commit new money stock
	c.moneyStock = tmpMoneyStock

	return change, nil
}

func (c *cashier) AddMoney(value float64, amount int) error {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()

	v := c.moneyStock[value] + amount

	// prevent money amount be negative
	if v < 0 {
		return ErrMoneyAmountCantBeNegative
	}

	c.moneyStock[value] = v

	return nil
}

func (c *cashier) GetMoneyStock() MoneyMap {
	return c.moneyStock
}

type MoneyMap map[float64]int

var EmptyMoneyMap = make(MoneyMap)

var (
	ErrInvalidPrice              error = errors.New("invalid price")
	ErrNotEnoughMoney            error = errors.New("not enough money")
	ErrMoneyAmountCantBeNegative error = errors.New("money amount can't be negative")
	ErrCantChangeMoney           error = errors.New("can't change money")
	ErrNotAllowedMoney           error = errors.New("not allowed money")
)

var AllowedMoney = []float64{
	1000,
	500,
	100,
	50,
	20,
	10,
	5,
	1,
	0.25,
}

func checkContainsForFloat(d []float64, t float64) bool {
	for _, v := range d {
		if v == t {
			return true
		}
	}

	return false
}
