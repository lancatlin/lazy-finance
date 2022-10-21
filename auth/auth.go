package auth

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthStore interface {
	Register(user, pass string) error
	Authenticate(user, pass string) error
	Remove(user string) error
}

type Htpasswd struct {
	accounts map[string]string
	filePath string
}

func NewHtpasswd(path string) (AuthStore, error) {
	s := Htpasswd{
		filePath: path,
	}
	err := s.read()
	return s, err
}

func (s Htpasswd) Register(user, pass string) (err error) {
	if _, ok := s.accounts[user]; ok {
		return errors.New("user already exists")
	}
	s.accounts[user], err = hash(pass)
	if err != nil {
		return
	}
	return s.write()
}

func (s Htpasswd) Authenticate(user, pass string) (err error) {
	hashed, ok := s.accounts[user]
	if !ok {
		return errors.New("user not found")
	}
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
}

func (s Htpasswd) Remove(user string) (err error) {
	delete(s.accounts, user)
	return s.write()
}

func (s *Htpasswd) read() (err error) {
	file, err := os.OpenFile(s.filePath, os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	s.accounts = make(map[string]string)
	for fileScanner.Scan() {
		arr := strings.SplitN(fileScanner.Text(), ":", 2)
		if len(arr) < 2 {
			return fmt.Errorf("invalid data %s", arr)
		}
		s.accounts[arr[0]] = arr[1]
	}
	return nil
}

func (s *Htpasswd) write() (err error) {
	file, err := os.OpenFile(s.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to open htpasswd file: %w", err)
	}
	defer file.Close()

	for u, p := range s.accounts {
		_, err = fmt.Fprintf(file, "%s:%s\n", u, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func hash(pass string) (string, error) {
	output, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(output), err
}
