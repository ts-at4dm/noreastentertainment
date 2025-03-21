package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../client/public/")))

	log.Println("Server Established on Port :5500")
	err := http.ListenAndServe(":5500", nil)
	if err != nil {
		log.Fatal(err)
	}
}
