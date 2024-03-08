package model

import (
	"bytes"
	"fmt"
	"text/template"
	"time"
)

type Transaction struct {
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Accounts []Account `json:"accounts"`
}

const txTemplate = `{{.Date.Format "2006/01/02" }} {{.Name}}
{{range .Accounts}}  {{.Name}}{{ if ne .Amount 0.0 }}  {{.Amount}} {{.Commodity}}{{end}}
{{end}}
`

func (tx Transaction) Validate() error {
	errorList := make([]error, 0)
	if tx.Name == "" {
		errorList = append(errorList, fmt.Errorf("name is required"))
	}
	if len(tx.Accounts) < 2 {
		errorList = append(errorList, fmt.Errorf("at least two accounts are required"))
	}
	sum := 0.0
	zeroCount := 0
	for _, account := range tx.Accounts {
		if account.Name == "" {
			errorList = append(errorList, fmt.Errorf("account names are required"))
		}
		if account.Amount == 0.0 {
			zeroCount++
		}
		sum += account.Amount
	}
	if zeroCount > 1 {
		errorList = append(errorList, fmt.Errorf("only one zero amount is allowed"))
	}
	if zeroCount == 0 && sum != 0.0 {
		errorList = append(errorList, fmt.Errorf("amounts must sum to 0"))
	}
	if len(errorList) > 0 {
		return genErrorMsg(errorList)
	}
	return nil
}

func genErrorMsg(errorList []error) error {
	errorMsg := ""
	for _, err := range errorList {
		errorMsg += err.Error() + ". "
	}
	return fmt.Errorf(errorMsg)
}

func (tx Transaction) Generate() (string, error) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("transaction").Parse(txTemplate)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(buf, tx)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func LoadTransactions(input string) ([]Transaction, error) {
	registers, err := LoadRegisters(input)
	if err != nil {
		return nil, err
	}
	return fromRegisters(registers)
}

func fromRegisters(registers []Register) ([]Transaction, error) {
	var transactions = make([]Transaction, 0, len(registers)/2)
	for _, reg := range registers {
		acc, err := reg.ToAccount()
		if err != nil {
			return nil, err
		}

		idx := reg.TxnIdx - 1

		if idx >= len(transactions) {
			transactions = append(transactions, Transaction{
				Name:     reg.Description,
				Date:     reg.Date,
				Accounts: []Account{acc},
			})
		} else {
			transactions[idx].Accounts = append(transactions[idx].Accounts, acc)
		}
	}

	return transactions, nil
}
