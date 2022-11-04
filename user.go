package main

import (
	"os"
	"path"
)

type User struct {
	IsLogin  bool
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (u *User) Dir() string {
	dir := path.Join(DATA_DIR, u.Email)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}
	return dir
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
