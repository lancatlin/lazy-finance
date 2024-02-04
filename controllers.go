package main

import (
	"github.com/gin-gonic/gin"
)

func getQueries(c *gin.Context) {
	user := getUser(c)
	queries, err := user.queries()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, queries)
}


func getTemplates(c *gin.Context) {
	user := getUser(c)
	templates, err := user.templates()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, templates)
}


