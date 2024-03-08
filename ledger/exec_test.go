package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenCommandArgs(t *testing.T) {
	command := Command{
		Command: "bal",
		Dir:     "test",
		Files:   []string{"test.journal"},
	}

	args := command.genArgs()
	expectedArgs := []string{"-f", "test.journal", "bal"}
	assert.Equal(t, expectedArgs, args)
}

// func TestExecuteLedgerCommand(t *testing.T) {

// }
