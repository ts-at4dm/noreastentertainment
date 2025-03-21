package main

import (
	"log"
	"net/http"

	"github.com/ts-at4dm/noreastentertainment/server.go/contactapi"
)



func main() {
	http.Handle("/", http.FileServer(http.Dir("")))
	contactapi.SetupRoutes()

	log.Println("Server Established on Port :5500")
	err := http.ListenAndServe(":5500", nil)
	if err != nil {
		log.Fatal(err)
	}
}
