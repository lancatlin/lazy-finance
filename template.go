package main

import (
	"path"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(path.Join(templatesDir, "layouts", "*.html"))
	if err != nil {
		panic(err)
	}

	includes, err := filepath.Glob(path.Join(templatesDir, "*.html"))
	if err != nil {
		panic(err)
	}

	for _, include := range includes {
		r.AddFromFiles(filepath.Base(include), append(layouts, include)...)
	}
	return r
}

type Data interface{}

type Page struct {
	Data
	User UserLogin
}

func HTML(c *gin.Context, status int, name string, data interface{}) {
	output := Page{
		Data: data,
	}
	_, ok := c.Get("user")
	if ok {
		output.User = c.MustGet("user").(UserLogin)
	}
	c.HTML(status, name, output)
}
