package ledger

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

type Command struct {
	Command string
	Dir     string
	Input   io.Reader
}

func (c Command) genArgs() []string {
	args := []string{"-f-", c.Command, "-O", "csv"}
	return args
}

func (c Command) Execute() (string, error) {
	cmd := exec.Command("hledger", c.genArgs()...)
	log.Println(cmd.Args)
	cmd.Dir = c.Dir
	cmd.Stdin = c.Input
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, out)
	}
	return string(out), nil
}
