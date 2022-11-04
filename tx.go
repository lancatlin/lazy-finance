package main

import (
	"bytes"
	"fmt"
	"io"
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

func (u *User) appendToFile(tx string) (err error) {
	f, err := u.File(DEFAULT_JOURNAL)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := strings.NewReader(strings.ReplaceAll(tx, "\r", "")) // Remove CR generated from browser
	_, err = io.Copy(f, buf)
	return err
}

func (u *User) executeScript(w io.Writer, name string) (err error) {
	script, ok := SCRIPTS[name]
	if !ok {
		return fmt.Errorf("%s script not found", name)
	}
	cmd := exec.Command("ledger", append([]string{"--file", DEFAULT_JOURNAL}, script...)...)
	cmd.Dir = u.Dir()
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}
