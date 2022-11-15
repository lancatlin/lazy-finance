package main

import (
	"io/ioutil"
	"log"
	"text/template"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	r := gin.Default()
	r.HTMLRender = loadTemplates("templates")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(303, "/dashboard")
	})

	r.GET("/signup", func(c *gin.Context) {
		HTML(c, 200, "signup.html", nil)
	})

	r.GET("/signin", func(c *gin.Context) {
		HTML(c, 200, "signin.html", nil)
	})

	r.POST("/signup", signup)

	r.POST("/signin", signin)

	authZone := r.Group("", authenticate)

	authZone.GET("/dashboard", func(c *gin.Context) {
		HTML(c, 200, "dashboard.html", struct {
			Templates []*template.Template
			Scripts   map[string]string
		}{
			ledgerTpl.Templates(),
			SCRIPTS,
		})
	})

	authZone.POST("/new", func(c *gin.Context) {
		var data TxData
		if err := c.ShouldBind(&data); err != nil {
			c.AbortWithError(400, err)
			return
		}
		tx, err := newTx(data)
		if err != nil {
			c.AbortWithError(400, err)
			log.Println(err, c.Request.Form)
			return
		}
		HTML(c, 200, "new.html", struct {
			Tx string
		}{tx})
	})
	authZone.POST("/submit", func(c *gin.Context) {
		user := getUser(c)
		tx := c.PostForm("tx")
		if err := user.appendToFile(tx); err != nil {
			c.AbortWithError(500, err)
			return
		}

		HTML(c, 200, "success.html", struct {
			Tx string
		}{tx})
	})

	authZone.GET("/edit", func(c *gin.Context) {
		user := getUser(c)
		filename := c.Query("filename")
		list, err := user.List()
		if err != nil {
			panic(err)
		}
		exists := contain(list, filename)
		var data []byte
		if exists {
			f, err := user.ReadFile(filename)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			data, err = ioutil.ReadAll(f)
			if err != nil {
				panic(err)
			}
		}

		HTML(c, 200, "edit.html", gin.H{
			"Data":     string(data),
			"FileName": filename,
			"FileList": list,
			"Exists":   exists,
		})
	})

	authZone.POST("/edit", func(c *gin.Context) {
		user := getUser(c)
		filename := c.PostForm("filename")
		data := c.PostForm("data")
		err := user.overwriteFile(filename, data)
		if err != nil {
			panic(err)
		}
		HTML(c, 200, "success.html", gin.H{
			"FileName": filename,
			"Tx":       data,
		})
	})

	authZone.GET("/download", func(c *gin.Context) {
		user := getUser(c)
		c.FileAttachment(user.FilePath(DEFAULT_JOURNAL), DEFAULT_JOURNAL)
	})

	authZone.GET("/query", func(c *gin.Context) {
		response := struct {
			Query   string
			Result  string
			Scripts map[string]string
		}{Scripts: SCRIPTS}
		user := getUser(c)
		var ok bool
		var err error
		response.Query, ok = c.GetQuery("query")
		if ok && response.Query != "" {
			response.Result, err = user.query(response.Query)
			if err != nil {
				panic(err)
			}
		}
		HTML(c, 200, "query.html", response)
	})

	return r
}
