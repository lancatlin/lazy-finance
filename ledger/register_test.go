package ledger

import (
	"testing"
	"time"

	"github.com/lancatlin/lazy-finance/model"
	"github.com/stretchr/testify/assert"
)

func TestLoadRegisters(t *testing.T) {
	input := `"txnidx","date","code","description","account","amount","total"
"1","2023-06-29","","午餐","expenses:lunch","$100","$100"`
	registers, err := LoadRegisters(input)
	assert.NoError(t, err)
	expectedRegisters := []Register{
		{
			TxnIdx:      1,
			Date:        MustParseTime("2023-06-29"),
			Description: "午餐",
			Account:     "expenses:lunch",
			Amount:      "$100",
			Total:       "$100",
		},
	}
	assert.Equal(t, expectedRegisters, registers)
}

func MustParseTime(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		panic(err)
	}
	return t
}

func TestFromRegisters(t *testing.T) {
	registers := []Register{
		{
			TxnIdx:      1,
			Date:        MustParseTime("2024-01-01"),
			Description: "restaurant",
			Account:     "expenses:food",
			Amount:      "$100",
			Total:       "$100",
		},
		{
			TxnIdx:      1,
			Date:        MustParseTime("2024-01-01"),
			Description: "restaurant",
			Account:     "asset:cash",
			Amount:      "-$100",
			Total:       "$0",
		},
		{
			TxnIdx:      2,
			Date:        MustParseTime("2024-01-02"),
			Description: "groceries",
			Account:     "expenses:groceries",
			Amount:      "$50",
			Total:       "$50",
		},
		{
			TxnIdx:      2,
			Date:        MustParseTime("2024-01-02"),
			Description: "groceries",
			Account:     "asset:cash",
			Amount:      "-$50",
			Total:       "$0",
		},
	}

	transactions, err := toTransactions(registers)
	assert.NoError(t, err)

	expectedTransactions := []model.Transaction{
		{
			Name: "restaurant",
			Date: MustParseTime("2024-01-01"),
			Accounts: []model.Account{
				{
					Name:      "expenses:food",
					Amount:    100,
					Commodity: "$",
				},
				{
					Name:      "asset:cash",
					Amount:    -100,
					Commodity: "$",
				},
			},
		},
		{
			Name: "groceries",
			Date: MustParseTime("2024-01-02"),
			Accounts: []model.Account{
				{
					Name:      "expenses:groceries",
					Amount:    50,
					Commodity: "$",
				},
				{
					Name:      "asset:cash",
					Amount:    -50,
					Commodity: "$",
				},
			},
		},
	}

	assert.Equal(t, expectedTransactions, transactions)
}
