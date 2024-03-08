package model

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type Register struct {
	TxnIdx      int
	Date        time.Time
	Description string
	Account     string
	Amount      string
	Total       string
}

func LoadRegisters(input string) ([]Register, error) {
	reader := csv.NewReader(strings.NewReader(input))
	reader.Read() // Skip header
	var registers []Register
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		registers = append(registers, parseRegister(record))
	}
	return registers, nil
}

func parseRegister(record []string) Register {
	date, _ := time.Parse("2006-01-02", record[1])
	txnIdx, _ := strconv.Atoi(record[0])
	register := Register{
		TxnIdx:      txnIdx,
		Date:        date,
		Description: record[3],
		Account:     record[4],
		Amount:      record[5],
		Total:       record[6],
	}
	return register
}

func ConvertAmount(money string) (float64, string, error) {
	var amount float64
	var commodity string
	var err error

	// Handle cases with a negative sign before the dollar sign
	if strings.HasPrefix(money, "-$") {
		commodity = "$"
		money = "-" + strings.TrimPrefix(money, "-$")
	} else if strings.HasPrefix(money, "$") {
		// Remove the dollar sign if present and set commodity to USD
		commodity = "$"
		money = strings.TrimPrefix(money, "$")
	} else {
		// Split the string by space to separate amount and commodity
		parts := strings.Split(money, " ")
		if len(parts) == 1 {
			money = parts[0]
		} else if len(parts) == 2 {
			money = parts[0]
			commodity = parts[1]
		} else {
			return 0, "", fmt.Errorf("invalid money format")
		}
	}

	// Convert the string amount to float64
	amount, err = strconv.ParseFloat(money, 64)
	if err != nil {
		return 0, "", fmt.Errorf("invalid amount format")
	}

	return amount, commodity, nil
}

func (r Register) ToAccount() (Account, error) {
	amount, commodity, err := ConvertAmount(r.Amount)
	if err != nil {
		return Account{}, err
	}
	return Account{
		Name:      r.Account,
		Amount:    amount,
		Commodity: commodity,
	}, nil
}
