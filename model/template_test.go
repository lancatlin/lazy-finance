package model

import (
	"testing"

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
