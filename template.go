package main

import (
	"path"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
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
