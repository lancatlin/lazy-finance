package main

import "github.com/gin-gonic/gin"

func signup(c *gin.Context) {
	var user UserLogin
	if err := c.ShouldBind(&user); err != nil {
		HTML(c, 400, "signup.html", err)
		return
	}
	if err := store.Register(user.Email, user.Password); err != nil {
		HTML(c, 400, "signup.html", err)
		return
	}
	signin(c)
}

func signin(c *gin.Context) {
	var user UserLogin
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
