package main

import (
	"errors"
	"flag"
	"log"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/lancatlin/ledger-quicknote/auth"
)

var ledgerTpl *template.Template
var htmlTpl *template.Template

var LEDGER_FILE string
var LEDGER_INIT string
var WORKING_DIR string
var HOST string

var store auth.AuthStore

const HTPASSWD_FILE = ".htpasswd"

type UserLogin struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func init() {
	ledgerTpl = template.Must(template.ParseGlob("tx/*"))
	flag.StringVar(&LEDGER_FILE, "f", "example.txt", "ledger journal file to write")
	flag.StringVar(&LEDGER_INIT, "i", "", "ledger initiation file")
	flag.StringVar(&WORKING_DIR, "w", "", "ledger working directory")
	flag.StringVar(&HOST, "b", "127.0.0.1:8000", "binding address")
	flag.Parse()
	var err error
	store, err = auth.NewHtpasswd(HTPASSWD_FILE)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()
	r.HTMLRender = loadTemplates("templates")

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup.html", nil)
	})

	r.POST("/signup", func(c *gin.Context) {
		var user UserLogin
		if err := c.ShouldBind(&user); err != nil {
			c.HTML(400, "signup.html", err)
			return
		}
		if err := store.Register(user.Email, user.Password); err != nil {
			c.HTML(400, "signup.html", err)
			return
		}
		c.Request.SetBasicAuth(user.Email, user.Password)
		c.Redirect(303, "/dashboard")
	})

	authZone := r.Group("", basicAuth)

	authZone.GET("/dashboard", func(c *gin.Context) {
		c.HTML(200, "index.html", struct {
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
		c.HTML(200, "new.html", struct {
			Tx string
		}{tx})
	})

	authZone.POST("/submit", func(c *gin.Context) {
		tx := c.PostForm("tx")
		if err := appendToFile(tx); err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.HTML(200, "success.html", struct {
			Tx string
		}{tx})
	})

	authZone.GET("/exec", func(c *gin.Context) {
		name, _ := c.GetQuery("name")
		if err := executeScript(c.Writer, name); err != nil {
			c.AbortWithError(500, err)
			log.Println(err)
			return
		}
	})

	log.Fatal(r.Run(HOST))
}

func basicAuth(c *gin.Context) {
	var user UserLogin
	var ok bool
	user.Email, user.Password, ok = c.Request.BasicAuth()
	if !ok {
		c.Header("WWW-Authenticate", "basic realm=\"Login to continue\"")
		c.AbortWithError(401, errors.New("login required"))
		return
	}
	if err := store.Authenticate(user.Email, user.Password); err != nil {
		c.AbortWithError(401, err)
		return
	}
	c.Set("user", user)
	c.Next()
}
