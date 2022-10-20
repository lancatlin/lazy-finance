package auth

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

type User struct {
	user   string
	pass   string
	hashed string
}

func TestHtpasswdSuccess(t *testing.T) {
	path := "/tmp/.htpasswd"
	user1 := User{
		user:   "user",
		pass:   "password",
		hashed: "$2a$14$SQSscaF4fVO3e5dp2/.VPuVQDPKqxSagLQnN6OncTRtoQw0ie9ByK",
	}
	err := ioutil.WriteFile(path,
		[]byte(fmt.Sprintf("%s:%s", user1.user, user1.hashed)), 0640)
	if err != nil {
		t.Fatal(err)
	}
	store, err := NewHtpasswd(path)
	if err != nil {
		t.Fatal(err)
	}
	err = store.Authenticate(user1.user, user1.pass)
	if err != nil {
		t.Error(err)
	}

	user2 := User{
		user: "foo",
		pass: "bar",
	}
	err = store.Register(user2.user, user2.pass)
	if err != nil {
		t.Error(err)
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
	}
	for _, u := range []User{user1, user2} {
		if !strings.Contains(string(data), u.user) {
			t.Errorf("%s not found in htpasswd file: %s", u.user, string(data))
		}
	}
	err = store.Remove(user1.user)
	if err != nil {
		t.Error(err)
	}

	data, err = ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
	}
	if strings.Contains(string(data), user1.user) {
		t.Errorf("%s is found in htpasswd file but should be removed: %s", user1.user, string(data))
	}
}
