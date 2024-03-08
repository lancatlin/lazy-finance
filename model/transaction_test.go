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

func TestValidateTransaction(t *testing.T) {
	txs := []struct {
		tx            Transaction
		expectedError string
	}{
		{tx: Transaction{}, expectedError: "name is required. at least two accounts are required. "},
		{tx: Transaction{Name: "test"}, expectedError: "at least two accounts are required. "},
		{tx: Transaction{Name: "test", Accounts: []Account{{Name: "test"}}}, expectedError: "at least two accounts are required. "},
		{tx: Transaction{Name: "test", Accounts: []Account{{Amount: 100}, {Name: "expenses", Amount: -10}}}, expectedError: "account names are required. amounts must sum to 0. "},
		{tx: Transaction{Name: "test", Accounts: []Account{{Name: "asset", Amount: 100}, {Name: "expenses", Amount: -10}}}, expectedError: "amounts must sum to 0. "},
		{tx: Transaction{Name: "test", Accounts: []Account{{Name: "asset", Amount: 100}, {Name: "expenses", Amount: 0}}}, expectedError: ""},
		{tx: Transaction{
			Name: "test",
			Accounts: []Account{
				{Name: "asset", Amount: 100},
				{Name: "expenses", Amount: 0},
				{Name: "liabilities", Amount: 0},
			}}, expectedError: "only one zero amount is allowed. "},
	}
	for _, tx := range txs {
		err := tx.tx.Validate()
		if tx.expectedError == "" {
			assertNil(t, err)
		} else {
			if err == nil {
				t.Errorf("Expected [%s], got nil", tx.expectedError)
			} else {
				assertEqual(t, tx.expectedError, err.Error())
			}
		}
	}
}