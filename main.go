package main

import (
	"fmt"
	"log"

	"github.com/lancatlin/ledger-quicknote/auth"
)

var config Config

var store auth.AuthStore

func init() {
}

func main() {
	r := router()

	log.Fatal(r.Run(fmt.Sprintf("%s:%s", config.Host, config.Port)))
}
