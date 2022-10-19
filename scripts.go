package main

import (
	"io"
	"os/exec"
)

var SCRIPTS = map[string][]string{
	"balance":  {"b"},
	"register": {"r"},
}

func executeScript(w io.Writer, name string) (err error) {
	cmd := exec.Command("ledger", append([]string{"-f", LEDGER_FILE}, SCRIPTS[name]...)...)
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}
