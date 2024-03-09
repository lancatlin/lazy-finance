package ledger

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteLedgerCommandBalance(t *testing.T) {
	command := Command{
		Query: Query{
			Command: "bal",
		},
		Input: strings.NewReader(`2024-01-01 restaurant
  Expenses:Food:Restaurant  $10
  Assets:Checking
`),
	}
	out, err := command.Execute()
	t.Log(out)
	assert.NoError(t, err)
	assert.Equal(t, `"account","balance"
"Assets:Checking","$-10"
"Expenses:Food:Restaurant","$10"
"total","0"
`, out)
}

func TestExecuteLedgerCommandRegister(t *testing.T) {
	command := Command{
		Query: Query{
			Command: "reg",
		},
		Input: strings.NewReader(`2024-01-01 restaurant
  Expenses:Food:Restaurant  $10
  Assets:Checking
`),
	}
	out, err := command.Execute()
	t.Log(out)
	assert.NoError(t, err)
	assert.Equal(t, `"txnidx","date","code","description","account","amount","total"
"1","2024-01-01","","restaurant","Expenses:Food:Restaurant","$10","$10"
"1","2024-01-01","","restaurant","Assets:Checking","$-10","0"
`, out)
}
