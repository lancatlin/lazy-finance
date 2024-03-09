package ledger

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

type Command struct {
	Query
	Dir   string
	Input io.Reader
}

func NewCommand(query Query, dir string, input io.Reader) Command {
	return Command{
		Query: query,
		Dir:   dir,
		Input: input,
	}
}

func (c Command) Execute() (string, error) {
	cmd := exec.Command("hledger", c.GetArgs()...)
	log.Println(cmd.Args)
	cmd.Dir = c.Dir
	cmd.Stdin = c.Input
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, out)
	}
	return string(out), nil
}
