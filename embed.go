package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed app/dist/*
var fs embed.FS

func staticFiles(c *gin.Context) {
	if c.Request.Method != "GET" {
		return
	}

	requestedPath := "app/dist" + c.Request.URL.Path
	if _, err := fs.Open(requestedPath); err == nil {
		c.FileFromFS(requestedPath, http.FS(fs))
	} else {
		c.FileFromFS("app/dist/", http.FS(fs))
	}
}
