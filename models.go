package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
	"time"

	cp "github.com/otiai10/copy"
)

type User struct {
	IsLogin  bool
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (u *User) Dir() string {
	dir := path.Join(DATA_DIR, u.Email)
	return dir
}

func (u *User) Mkdir() error {
	return cp.Copy(ARCHETYPES_DIR, u.Dir())
}

func (u *User) FilePath(name string) string {
	return path.Join(u.Dir(), name)
}

func (u *User) File(name string, mode int) (*os.File, error) {
	return os.OpenFile(u.FilePath(name), mode, 0644)
}

func (u *User) AppendFile(name string) (*os.File, error) {
	return u.File(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND)
}

func (u *User) ReadFile(name string) (*os.File, error) {
	return u.File(name, os.O_RDONLY|os.O_CREATE)
}

func (u *User) WriteFile(name string) (*os.File, error) {
	return u.File(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
}

func (u *User) List() ([]string, error) {
	files, err := os.ReadDir(u.Dir())
	if err != nil {
		return []string{}, fmt.Errorf("Failed to open directory: %w", err)
	}
	result := make([]string, len(files))
	for i, v := range files {
		result[i] = v.Name()
	}
	return result, nil
}

func (u *User) readAllFile(name string) (data []byte, err error) {
	f, err := u.ReadFile(name)
	if err != nil {
		return
	}
	defer f.Close()
	data, err = ioutil.ReadAll(f)
	return
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

func (u *User) overwriteFile(filename string, tx string) (err error) {
	f, err := u.WriteFile(filename)
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

func (u *User) queries() (queries [][2]string, err error) {
	f, err := u.ReadFile(QUERIES_FILE)
	if err != nil {
		err = fmt.Errorf("Failed to read queries file: %w", err)
		return
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	queries = make([][2]string, 0)
	for fileScanner.Scan() {
		arr := strings.SplitN(fileScanner.Text(), ":", 2)
		if len(arr) < 2 {
			continue
		}
		queries = append(queries, [2]string{arr[0], arr[1]})
	}
	return
}

func (u *User) templates() (templates []string, err error) {
	files, err := u.List()
	if err != nil {
		return
	}
	for _, v := range files {
		if strings.HasSuffix(v, ".tpl") {
			templates = append(templates, v)
		}
	}
	return
}

type TxData struct {
	Action      string `form:"action" binding:"required"`
	Name        string `form:"name"`
	Date        string
	Amount      string `form:"amount" binding:"required"`
	Destination string `form:"dest"`
	Source      string `form:"src"`
}

func (u *User) newTx(data TxData) (result string, err error) {
	data.Date = time.Now().Format("2006/01/02")
	var buf bytes.Buffer
	tpl, err := template.ParseFiles(u.FilePath(data.Action))
	if err != nil {
		return
	}
	err = tpl.ExecuteTemplate(&buf, data.Action, data)
	result = buf.String()
	return
}
