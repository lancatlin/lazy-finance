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
		HTML(c, 200, "index.html", nil)
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
			Scripts   map[string][]string
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

	authZone.GET("/exec", func(c *gin.Context) {
		user := getUser(c)
		name, _ := c.GetQuery("name")
		if err := user.executeScript(c.Writer, name); err != nil {
			c.AbortWithError(500, err)
			log.Println(err)
			return
		}
	})

	authZone.GET("/edit", func(c *gin.Context) {
		user := getUser(c)
		f, err := user.ReadFile(DEFAULT_JOURNAL)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		data, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		HTML(c, 200, "edit.html", string(data))
	})

	authZone.POST("/edit", func(c *gin.Context) {
		user := getUser(c)
		data := c.PostForm("data")
		err := user.overwriteFile(data)
		if err != nil {
			panic(err)
		}
		HTML(c, 200, "success.html", struct {
			Tx string
		}{data})
	})

	authZone.GET("/download", func(c *gin.Context) {
		user := getUser(c)
		c.FileAttachment(user.FilePath(DEFAULT_JOURNAL), DEFAULT_JOURNAL)
	})

	return r
}
