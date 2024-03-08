package model

import (
	"testing"
	"time"
)

func TestGenTransactionText(t *testing.T) {
	// Given
	transaction := Transaction{
		Name: "restaurant",
		Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Accounts: []Account{
			{
				Name:      "expenses:food",
				Amount:    100,
				Commodity: "$",
			},
			{
				Name:      "asset:cash",
				Amount:    0,
				Commodity: "$",
			},
		},
	}

	expectedTx := `2024-01-01 restaurant
  expenses:food  100 $
  asset:cash
`
	actualTx, err := transaction.Generate()
	assertEqual(t, expectedTx, actualTx)
	assertNil(t, err)
}
