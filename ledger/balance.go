package ledger

import (
	"encoding/csv"
	"fmt"
	"strings"
)

type Balance struct {
	Account string `json:"account"`
	Balance string `json:"balance"`
}

func LoadBalances(input string) ([]Balance, error) {
	reader := csv.NewReader(strings.NewReader(input))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	balances := make([]Balance, len(records)-1)
	for i, record := range records[1:] { // Skipping header row
		if len(record) != 2 {
			return nil, fmt.Errorf("expected 2 fields, got %d in %+v", len(record), record)
		}
		balances[i] = Balance{
			Account: record[0],
			Balance: record[1],
		}
	}
	return balances, nil
}
