package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authenticate(c *gin.Context) {
	cookie, err := c.Cookie("session")
	if err == http.ErrNoCookie {
		c.AbortWithError(401, err)
		return
	}
	session, err := store.Verify(cookie)
	if err != nil {
		c.AbortWithError(401, err)
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
		c.AbortWithError(400, err)
		return
	}
	if err := store.Register(user.Email, user.Password); err != nil {
		c.AbortWithError(400, err)
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
		c.AbortWithError(400, err)
		return
	}
	token, err := store.Login(user.Email, user.Password)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}
	c.SetCookie("session", token, 60*60*24*7, "", "", false, false)
	c.Status(200)
}

func logout(c *gin.Context) {
	c.SetCookie("session", "", -1, "", "", false, false)
	c.Status(200)
}
