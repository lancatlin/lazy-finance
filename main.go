package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

var ledgerTpl *template.Template
var htmlTpl *template.Template

var LEDGER_FILE string
var LEDGER_INIT string
var WORKING_DIR string
var HOST string

type TxData struct {
	Action  string `form:"action" binding:"required"`
	Name    string `form:"name"`
	Date    string
	Amount  string `form:"amount" binding:"required"`
	Account string `form:"account"`
}

func init() {
	ledgerTpl = template.Must(template.ParseGlob("tx/*"))
	flag.StringVar(&LEDGER_FILE, "f", "example.txt", "ledger journal file to write")
	flag.StringVar(&LEDGER_INIT, "i", "", "ledger initiation file")
	flag.StringVar(&WORKING_DIR, "w", "", "ledger working directory")
	flag.StringVar(&HOST, "b", "127.0.0.1:8000", "binding address")
	flag.Parse()
}

func main() {
	r := gin.Default()
	r.HTMLRender = loadTemplates("templates")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", struct {
			Templates []*template.Template
			Scripts   map[string][]string
		}{
			ledgerTpl.Templates(),
			SCRIPTS,
		})
	})

	r.POST("/new", func(c *gin.Context) {
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

	r.POST("/submit", func(c *gin.Context) {
		tx := c.PostForm("tx")
		if err := appendToFile(tx); err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.HTML(200, "success.html", struct {
			Tx string
		}{tx})
	})

	r.GET("/exec", func(c *gin.Context) {
		name, _ := c.GetQuery("name")
		if err := executeScript(c.Writer, name); err != nil {
			c.AbortWithError(500, err)
			log.Println(err)
			return
		}
	})

	log.Fatal(r.Run(HOST))
}

func newTx(data TxData) (result string, err error) {
	data.Date = time.Now().Format("2006/01/02")
	var buf bytes.Buffer
	err = ledgerTpl.ExecuteTemplate(&buf, data.Action, data)
	return buf.String(), nil
}

func appendToFile(tx string) (err error) {
	f, err := os.OpenFile(LEDGER_FILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := strings.NewReader(strings.ReplaceAll(tx, "\r", "")) // Remove CR generated from browser
	_, err = io.Copy(f, buf)
	return err
}

func executeScript(w io.Writer, name string) (err error) {
	script, ok := SCRIPTS[name]
	if !ok {
		return fmt.Errorf("%s script not found", name)
	}
	cmd := exec.Command("ledger", append([]string{"--init-file", LEDGER_INIT, "--file", LEDGER_FILE}, script...)...)
	cmd.Dir = WORKING_DIR
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}
