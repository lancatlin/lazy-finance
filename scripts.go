package main

import (
	"fmt"
	"io"
	"os/exec"
)

func executeScript(w io.Writer, name string) (err error) {
	script, ok := SCRIPTS[name]
	if !ok {
		return fmt.Errorf("%s script not found", name)
	}
	cmd := exec.Command("ledger", append([]string{"--init-file", LEDGER_INIT}, script...)...)
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}
