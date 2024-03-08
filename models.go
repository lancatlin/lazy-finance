package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/lancatlin/lazy-finance/model"
	cp "github.com/otiai10/copy"
)

type User struct {
	IsLogin  bool
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (u *User) Dir() string {
	dir := path.Join(config.DataPath, u.Email)
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
		return []string{}, fmt.Errorf("failed to open directory: %w", err)
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
	data, err = io.ReadAll(f)
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

	cmd := exec.Command("hledger")
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
		err = fmt.Errorf("failed to read queries file: %w", err)
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

func (u *User) templates() (templates []model.Template, err error) {
	f, err := u.ReadFile("templates.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read templates.json: %w", err)
	}
	defer f.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read from buffer: %w", err)
	}

	templates, err = model.LoadTemplates(buf.String())
	if err != nil {
		return nil, fmt.Errorf("failed to load templates: %w", err)
	}
	return templates, nil
}

func (u *User) newTx(tx model.Transaction) error {
	if err := tx.Validate(); err != nil {
		return err
	}
	txString, err := tx.Generate()
	if err != nil {
		return err
	}
	err = u.appendToFile(txString)
	if err != nil {
		return err
	}
	return nil
}
