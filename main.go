package main

import (
	"encoding/base64"
	"flag"
	"log"
	"path"
	"text/template"

	"github.com/gorilla/securecookie"
	"github.com/lancatlin/ledger-quicknote/auth"
)

var ledgerTpl *template.Template
var htmlTpl *template.Template

var DATA_DIR string
var HOST string

var store auth.AuthStore

const HTPASSWD_FILE = ".htpasswd"
const DEFAULT_JOURNAL = "ledger.txt"

func init() {
	ledgerTpl = template.Must(template.ParseGlob("tx/*"))
	flag.StringVar(&DATA_DIR, "d", "data", "data folder")
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
	store, err = auth.New(path.Join(DATA_DIR, HTPASSWD_FILE), hashKey)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := router()

	log.Fatal(r.Run(HOST))
}
