package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadTemplates(t *testing.T) {
	templateString := `[{
	"name": "restaurant",
	"accounts": [
		{
			"name": "expenses:food",
			"amount": 100,
			"commodity": "$"
		},
		{
			"name": "asset:cash",
			"amount": -100,
			"commodity": "$"
		}
	]	
}]`
	templates, err := LoadTemplates(templateString)
	assert.NoError(t, err)
	expectedTemplates := []Template{
		{
			Name: "restaurant",
			Accounts: []Account{
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
	}
	assert.Equal(t, expectedTemplates, templates)
}

func TestFromTransaction(t *testing.T) {
	tx := Transaction{
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

	template := FromTransaction(tx)
	expected := Template{
		Name: "restaurant",
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
	assert.Equal(t, expected, template)
}
