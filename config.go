package main

import (
	"encoding/base64"
	"flag"
	"log"
	"os"
	"path"

	"github.com/gorilla/securecookie"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lancatlin/lazy-finance/auth"
)

const QUERIES_FILE = "queries.txt"
const HTPASSWD_FILE = ".htpasswd"
const DEFAULT_JOURNAL = "ledger.txt"
const ARCHETYPES_DIR = "archetypes"

type Config struct {
	Host     string
	Port     string
	DataPath string
	HashKey  []byte
}

func init() {
	config = newConfig()
}

func env(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func newConfig() Config {
	config := Config{}

	var hashKey string
	defaultDataDir := env("DATA_DIR", "data")
	defaultHost := env("HOST", "127.0.0.1")
	defaultPort := env("PORT", "8000")
	defaultHashKey := env("HASH_KEY", "")

	flag.StringVar(&config.DataPath, "d", defaultDataDir, "data folder")
	flag.StringVar(&config.Host, "b", defaultHost, "binding address")
	flag.StringVar(&config.Port, "p", defaultPort, "binding address")
	flag.StringVar(&hashKey, "s", defaultHashKey, "session secret")
	flag.Parse()

	var err error

	if hashKey == "" {
		config.HashKey = securecookie.GenerateRandomKey(32)
		log.Printf("Generate random session key: %s", base64.StdEncoding.EncodeToString(config.HashKey))
	} else {
		config.HashKey, err = base64.StdEncoding.DecodeString(hashKey)
		if err != nil {
			panic(err)
		}
	}
	store, err = auth.New(path.Join(config.DataPath, HTPASSWD_FILE), config.HashKey)
	if err != nil {
		panic(err)
	}

	return config
}
