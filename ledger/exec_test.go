package ledger

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenCommandArgs(t *testing.T) {
	command := Command{
		Command: "bal",
	}

	args := command.genArgs()
	expectedArgs := []string{"-f-", "bal", "-O", "csv"}
	assert.Equal(t, expectedArgs, args)
}

func TestExecuteLedgerCommand(t *testing.T) {
	command := Command{
		Command: "bal",
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
