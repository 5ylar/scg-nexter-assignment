package cashier

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestChange(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		price                decimal.Decimal
		moneyStock           MoneyMap
		money                MoneyMap
		expectChange         MoneyMap
		expectError          error
		expectMoneyStockLeft MoneyMap // (moneyStock[n] + money[n]) - expectChange[n]
	}{
		{
			name:  "case 1",
			price: decimal.NewFromInt(300),
			moneyStock: MoneyMap{
				100: 1,
			},
			money: MoneyMap{
				100: 4,
			},
			expectChange: MoneyMap{
				100: 1,
			},
			expectMoneyStockLeft: MoneyMap{
				100: 4,
			},
		},
		{
			name:  "case 2",
			price: decimal.NewFromInt(300),
			moneyStock: MoneyMap{
				100: 1,
			},
			money: MoneyMap{
				100: 3,
			},
			expectChange: EmptyMoneyMap,
			expectMoneyStockLeft: MoneyMap{
				100: 4,
			},
		},
		{
			name:  "case 3",
			price: decimal.NewFromInt(300),
			moneyStock: MoneyMap{
				100: 1,
			},
			money: MoneyMap{
				100: 1,
			},
			expectChange: EmptyMoneyMap,
			expectMoneyStockLeft: MoneyMap{
				100: 1,
			},
			expectError: ErrNotEnoughMoney,
		},
		{
			name:       "case 4",
			price:      decimal.NewFromInt(300),
			moneyStock: EmptyMoneyMap,
			money: MoneyMap{
				1000: 1,
			},
			expectChange:         EmptyMoneyMap,
			expectMoneyStockLeft: EmptyMoneyMap,
			expectError:          ErrCantChangeMoney,
		},
		{
			name:  "case 5",
			price: decimal.NewFromInt(300),
			moneyStock: MoneyMap{
				100: 1,
				50:  1,
				20:  2,
			},
			money: MoneyMap{
				500: 1,
			},
			expectChange: EmptyMoneyMap,
			expectMoneyStockLeft: MoneyMap{
				100: 1,
				50:  1,
				20:  2,
			},
			expectError: ErrCantChangeMoney,
		},
		{
			name:  "case 6",
			price: decimal.NewFromInt(300),
			moneyStock: MoneyMap{
				100: 1,
				50:  1,
				20:  2,
				10:  2,
			},
			money: MoneyMap{
				500: 1,
			},
			expectChange: MoneyMap{
				100: 1,
				50:  1,
				20:  2,
				10:  1,
			},
			expectMoneyStockLeft: MoneyMap{
				500: 1,
				10:  1,
			},
		},
		{
			name:  "case 7",
			price: decimal.NewFromInt(300),
			moneyStock: MoneyMap{
				100: 1,
				50:  1,
				20:  2,
				10:  2,
			},
			money: MoneyMap{
				600: 1,
			},
			expectChange: EmptyMoneyMap,
			expectMoneyStockLeft: MoneyMap{
				100: 1,
				50:  1,
				20:  2,
				10:  2,
			},
			expectError: ErrNotAllowedMoney,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ac := New(tc.moneyStock)

			change, err := ac.Change(tc.price, tc.money)

			assert.Equal(t, tc.expectError, err)
			assert.Equal(t, tc.expectChange, change)
			assert.Equal(t, tc.expectMoneyStockLeft, ac.GetMoneyStock())
		})
	}
}
