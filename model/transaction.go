package model

import "time"

type Transaction struct {
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Accounts []Account `json:"accounts"`
}
