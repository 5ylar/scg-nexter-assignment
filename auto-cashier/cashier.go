package autocashier

import (
	"errors"
	"sync"

	"github.com/shopspring/decimal"
)

type autoCashier struct {
	monies []MoneyStock
	mu     sync.Mutex
}

func New(monies []MoneyStock) *autoCashier {
	return &autoCashier{
		monies: monies,
	}
}

func (c *autoCashier) Change(currency Currency, price decimal.Decimal, moneyKeys []MoneyKeyWithAmont) ([]Money, error) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()

	var (
		changes     []Money
		totalMonies decimal.Decimal
	)

	// price > totalMonies
	if price.Cmp(totalMonies) == 1 {
		return []Money{}, ErrNotEnoughMoney
	}

	return changes, nil
}

func (c *autoCashier) addMoney(key string, amount int) error {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()

	for i, m := range c.monies {
		if m.Key == key {

			// absolute value
			if amount < 0 {
				amount = -amount
			}

			// add
			c.monies[i].Amount += amount

			return nil
		}
	}

	return ErrNotSupportMoney
}

func (c *autoCashier) subMoney(key string, amount int) error {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()

	for i, m := range c.monies {
		if m.Key == key {

			// absolute value
			if amount < 0 {
				amount = -amount
			}

			// subtract
			c.monies[i].Amount -= amount

			return nil
		}
	}

	return ErrNotSupportMoney
}

func (c *autoCashier) ListMoniesByCurrency(currency Currency) ([]Money, error) {
	var monies []Money

	return monies, nil
}

type MoneyKeyWithAmont struct {
	Key    string
	Amount int
}

type MoneyStock struct {
	Money
	Amount int
}

type Money struct {
	Key      string
	Value    decimal.Decimal
	Currency Currency
	Type     MoneyType
}

type MoneyType string

const (
	MoneyTypeBankNote MoneyType = "bank_note"
	MoneyTypeCoin     MoneyType = "coin"
)

type Currency string

const (
	CurrencyTHB Currency = "THB"
)

var (
	ErrNotSupportMoney error = errors.New("money isn't support")
	ErrNotEnoughMoney  error = errors.New("not enough money")
)
