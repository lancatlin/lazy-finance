package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func session(c *gin.Context) {
	cookie, err := c.Cookie("session")
	if err == http.ErrNoCookie {
		return
	}
	session, err := store.Verify(cookie)
	if err != nil {
		return
	}
	c.Set("user", User{
		Email: session.User,
	})
	c.Next()
}

func authRequired(c *gin.Context) {
	if _, ok := c.Get("user"); !ok {
		c.AbortWithError(401, errors.New("unauthorized"))
		return
	}
	c.Next()
}

func getUser(c *gin.Context) User {
	return c.MustGet("user").(User)
}

// @Summary Sign up
// @Description Sign up
// @Accept json
// @Produce json
// @Param data body User true "User"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Router /signup [post]
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

// @Summary Sign in
// @Description Sign in
// @Accept json
// @Produce json
// @Param data body User true "User"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Router /signin [post]
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
	c.SetCookie("session", token, 60*60*24*7, "", "", false, true)
	c.Status(200)
}

func logout(c *gin.Context) {
	c.SetCookie("session", "", -1, "", "", false, false)
	c.Status(200)
}

// @Summary Get status
// @Description Get status
// @Produce json
// @Success 200 {boolean} bool
// @Router /status [get]
func status(c *gin.Context) {
	_, ok := c.Get("user")
	c.JSON(200, gin.H{
		"signed_in": ok,
	})
}
