package auth

import (
	"errors"
	"io/ioutil"
	"log"

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
	_, err := ioutil.ReadFile(path)
	if err != nil {
		return Htpasswd{}, err
	}
	return Htpasswd{}, nil
}

func (s Htpasswd) Register(user, pass string) (err error) {
	s.accounts[user], err = hash(pass)
	if err != nil {
		return
	}
	log.Println(s.accounts[user])
	return
}

func (s Htpasswd) Authenticate(user, pass string) (err error) {
	return errors.New("work in progress")
}

func (s Htpasswd) Remove(user string) (err error) {
	return errors.New("work in progress")
}

func hash(pass string) (string, error) {
	output, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(output), err
}
