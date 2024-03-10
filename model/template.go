package model

import "encoding/json"

type Template struct {
	Name     string    `json:"name"`
	Accounts []Account `json:"accounts"`
}

func LoadTemplates(input string) (templates []Template, err error) {
	err = json.Unmarshal([]byte(input), &templates)
	return
}

func FromTransaction(tx Transaction) (t Template) {
	t = Template{
		Name:     tx.Name,
		Accounts: tx.Accounts,
	}
	return
}
