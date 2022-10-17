package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

const LEDGER_FILE = "test.txt"

type TxData struct {
	Date   string
	Amount string
}

func init() {
	var err error
	tpl, err = template.ParseGlob("templates/*.txt")
	if err != nil {
		panic(err)
	}
	log.Println(tpl.DefinedTemplates())
}

func main() {
	fmt.Println("Hi")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/action", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			panic(err)
		}
		fmt.Println(r.Form)
		err := renderTx(w, r.Form)
		if err != nil {
			panic(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func renderTx(w io.Writer, params url.Values) (err error) {
	name := params.Get("action")
	data := TxData{
		Date:   time.Now().Format("2006/01/02"),
		Amount: params.Get("amount"),
	}
	f, err := os.OpenFile(LEDGER_FILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = tpl.ExecuteTemplate(f, name, data); err != nil {
		return err
	}
	if err = tpl.ExecuteTemplate(w, name, data); err != nil {
		return err
	}
	return nil
}
