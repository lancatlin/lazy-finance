package main

import (
	"bytes"
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
	f, err := u.AppendFile(DEFAULT_JOURNAL)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := strings.NewReader(strings.ReplaceAll(tx, "\r", "")) // Remove CR generated from browser
	_, err = io.Copy(f, buf)
	return err
}

func (u *User) overwriteFile(tx string) (err error) {
	f, err := u.WriteFile(DEFAULT_JOURNAL)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := strings.NewReader(strings.ReplaceAll(tx, "\r", "")) // Remove CR generated from browser
	_, err = io.Copy(f, buf)
	return err
}

func (u *User) query(query string) (result string, err error) {
	var buf bytes.Buffer

	cmd := exec.Command("ledger", "--file", DEFAULT_JOURNAL)
	cmd.Dir = u.Dir()
	cmd.Stdin = strings.NewReader(query)
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err = cmd.Run()
	return buf.String(), err
}
