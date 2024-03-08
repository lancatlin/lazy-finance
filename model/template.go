package model

import "encoding/json"

type Template struct {
	Name     string    `json:"name"`
	Accounts []Account `json:"accounts"`
}

func loadTemplates(input string) (templates []Template, err error) {
	err = json.Unmarshal([]byte(input), &templates)
	return
}
