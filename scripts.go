package main

import (
	"fmt"
	"io"
	"os/exec"
)

var SCRIPTS = map[string][]string{
	"balance":             {"b"},
	"register":            {"r"},
	"expenses this month": {"b", "expenses", "-b", "this month"},
}

func executeScript(w io.Writer, name string) (err error) {
	script, ok := SCRIPTS[name]
	if !ok {
		return fmt.Errorf("%s script not found", name)
	}
	cmd := exec.Command("ledger", append([]string{"-f", LEDGER_FILE}, script...)...)
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}
