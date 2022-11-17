package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authenticate(c *gin.Context) {
	cookie, err := c.Cookie("session")
	if err == http.ErrNoCookie {
		c.Redirect(303, "/signin")
		return
	}
	session, err := store.Verify(cookie)
	if err != nil {
		c.Redirect(303, "/signin")
		return
	}
	c.Set("user", User{
		Email: session.User,
	})
	c.Next()
}

func getUser(c *gin.Context) User {
	return c.MustGet("user").(User)
}

func signup(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		HTML(c, 400, "signup.html", err)
		return
	}
	if err := store.Register(user.Email, user.Password); err != nil {
		HTML(c, 400, "signup.html", err)
		return
	}
	if err := user.Mkdir(); err != nil {
		return
	}
	signin(c)
}

func signin(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		HTML(c, 400, "signin.html", err)
		return
	}
	token, err := store.Login(user.Email, user.Password)
	if err != nil {
		HTML(c, 401, "signin.html", err)
		return
	}
	c.SetCookie("session", token, 60*60*24*7, "", "", false, false)
	c.Redirect(303, "/dashboard")
}
