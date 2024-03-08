package model

import (
	"reflect"
	"testing"
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
	assertNil(t, err)
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
	assertDeepEqual(t, expectedTemplates, templates)
}

func assertNil(t *testing.T, value interface{}) {
	if value != nil {
		t.Errorf("Expected nil, got `%+v`", value)
	}
}

func assertNotNil(t *testing.T, value interface{}) {
	if value == nil {
		t.Fatalf("Expected not nil, got `%+v`", value)
	}
}

func assertEqual[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Errorf("Expected `%+v`, got `%+v`", expected, actual)
	}
}

func assertDeepEqual(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected and actual do not match. Expected: `%+v`, got: `%+v`", expected, actual)
	}
}
