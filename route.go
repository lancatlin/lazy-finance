package main

import (
	"log"

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

	api := r.Group("/api")

	authApi := api.Group("", authenticate)

	authApi.GET("/queries", getQueries)
	authApi.GET("/templates", getTemplates)

	authZone := r.Group("", authenticate)

	authZone.GET("/dashboard", func(c *gin.Context) {
		user := getUser(c)
		queries, err := user.queries()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		templates, err := user.templates()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		HTML(c, 200, "dashboard.html", gin.H{
			"Queries":   queries,
			"Templates": templates,
		})
	})

	authZone.POST("/new", func(c *gin.Context) {
		var data TxData
		if err := c.ShouldBind(&data); err != nil {
			c.AbortWithError(400, err)
			return
		}
		user := getUser(c)
		tx, err := user.newTx(data)
		if err != nil {
			c.AbortWithError(400, err)
			log.Println(err, c.Request.Form)
			return
		}
		HTML(c, 200, "new.html", gin.H{
			"Tx": tx,
		})
	})
	authZone.POST("/submit", func(c *gin.Context) {
		user := getUser(c)
		tx := c.PostForm("tx")
		if err := user.appendToFile(tx); err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.Redirect(303, "/dashboard")
	})

	authZone.GET("/edit", func(c *gin.Context) {
		user := getUser(c)
		filename := c.Query("filename")
		list, err := user.List()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		exists := contain(list, filename)
		var data []byte
		if exists {
			data, err = user.readAllFile(filename)
			if err != nil {
				c.AbortWithError(500, err)
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
			c.AbortWithError(500, err)
			return
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
		user := getUser(c)
		response := struct {
			Query   string
			Result  string
			Queries [][2]string
		}{}
		var ok bool
		var err error
		response.Queries, err = user.queries()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		response.Query, ok = c.GetQuery("query")
		if ok && response.Query != "" {
			response.Result, err = user.query(response.Query)
			if err != nil {
				c.AbortWithError(500, err)
				return
			}
		}
		HTML(c, 200, "query.html", response)
	})

	return r
}
