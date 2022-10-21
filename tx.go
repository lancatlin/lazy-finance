package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

type TxData struct {
	Action  string `form:"action" binding:"required"`
	Name    string `form:"name"`
	Date    string
	Amount  string `form:"amount" binding:"required"`
	Account string `form:"account"`
}

func newTx(data TxData) (result string, err error) {
	data.Date = time.Now().Format("2006/01/02")
	var buf bytes.Buffer
	err = ledgerTpl.ExecuteTemplate(&buf, data.Action, data)
	return buf.String(), nil
}

func appendToFile(tx string) (err error) {
	f, err := os.OpenFile(LEDGER_FILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := strings.NewReader(strings.ReplaceAll(tx, "\r", "")) // Remove CR generated from browser
	_, err = io.Copy(f, buf)
	return err
}

func executeScript(w io.Writer, name string) (err error) {
	script, ok := SCRIPTS[name]
	if !ok {
		return fmt.Errorf("%s script not found", name)
	}
	cmd := exec.Command("ledger", append([]string{"--init-file", LEDGER_INIT, "--file", LEDGER_FILE}, script...)...)
	cmd.Dir = WORKING_DIR
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}
