package main

import (
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	r := router.Generate()

	log.Println("Server running on port 5000")

	log.Fatal(http.ListenAndServe(":5000", r))
}
