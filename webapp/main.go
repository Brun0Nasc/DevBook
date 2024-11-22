package main

import (
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.LoadTemplates()
	r := router.Generate()
	
	log.Println("Server started on: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
