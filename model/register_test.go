package model

import (
	"testing"
	"time"
)

func TestLoadRegisters(t *testing.T) {
	input := `"txnidx","date","code","description","account","amount","total"
"1","2023-06-29","","午餐","expenses:lunch","$100","$100"`
	registers, err := LoadRegisters(input)
	assertNil(t, err)
	expectedRegisters := []Register{
		{
			TxnIdx:      1,
			Date:        MustParseTime("2023-06-29"),
			Description: "午餐",
			Account:     "expenses:lunch",
			Amount:      "$100",
			Total:       "$100",
		},
	}
	assertDeepEqual(t, expectedRegisters, registers)
}

func MustParseTime(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		panic(err)
	}
	return t
}
