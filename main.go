package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"text/template"
	"time"
)

var tpl *template.Template

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
		if err := renderTx(w, r.Form); err != nil {
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
	return tpl.ExecuteTemplate(w, name, data)
}
