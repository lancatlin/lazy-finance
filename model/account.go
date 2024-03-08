package model

type Account struct {
	Name      string  `json:"name"`
	Amount    float64 `json:"amount"`
	Commodity string  `json:"commodity"`
}
