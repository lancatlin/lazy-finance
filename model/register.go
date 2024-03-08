package model

import (
	"encoding/csv"
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
