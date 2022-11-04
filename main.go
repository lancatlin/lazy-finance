package main

import (
	"encoding/base64"
	"flag"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
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
	IsLogin  bool
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func init() {
	ledgerTpl = template.Must(template.ParseGlob("tx/*"))
	flag.StringVar(&LEDGER_FILE, "f", "example.txt", "ledger journal file to write")
	flag.StringVar(&LEDGER_INIT, "i", "", "ledger initiation file")
	flag.StringVar(&WORKING_DIR, "w", "", "ledger working directory")
	flag.StringVar(&HOST, "b", "127.0.0.1:8000", "binding address")
	var hashKeyString string
	flag.StringVar(&hashKeyString, "s", "", "session secret")
	flag.Parse()

	var hashKey []byte
	var err error

	if hashKeyString == "" {
		hashKey = securecookie.GenerateRandomKey(32)
		log.Printf("Generate random session key: %s", base64.StdEncoding.EncodeToString(hashKey))
	} else {
		hashKey, err = base64.StdEncoding.DecodeString(hashKeyString)
		if err != nil {
			panic(err)
		}
	}
	store, err = auth.New(HTPASSWD_FILE, hashKey)
	if err != nil {
		panic(err)
	}
}

func main() {
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
		tx := c.PostForm("tx")
		if err := appendToFile(tx); err != nil {
			c.AbortWithError(500, err)
			return
		}

		HTML(c, 200, "success.html", struct {
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

func authenticate(c *gin.Context) {
	cookie, err := c.Cookie("session")
	if err == http.ErrNoCookie {
		c.Redirect(303, "/signin")
		return
	}
	session, err := store.Verify(cookie)
	if err != nil {
		c.Redirect(303, "/signin")
		return
	}
	c.Set("user", UserLogin{
		Email: session.User,
	})
	c.Next()
}
