package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL represents the url to communication with the api
	APIURL = ""
	// Port where the web app is running
	Port = 0
	// HashKey is used to authenticate the cookie
	HashKey []byte
	// BlockKey is used to cryptogaph cookie's data
	BlockKey []byte
)

// Load initializes the enviroment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
