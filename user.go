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
	return path.Join(DATA_DIR, u.Email, name)
}

func (u *User) File(name string) (*os.File, error) {
	return os.OpenFile(u.FilePath(name), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
}
