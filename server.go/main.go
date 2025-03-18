package main

import (
	"log"
	"net/http"
	"./handlers"
)



func main() {
	http.HandleFunc("/", handlers.homeHandler)
	http.HandleFunc("/about", handlers.homeHandler)
	http.HandleFunc("/account", handlers.homeHandler)
	http.HandleFunc("/calendar", handlers.homeHandler)
	http.HandleFunc("/contact", handlers.homeHandler)
	http.HandleFunc("/corp", handlers.homeHandler)
	http.HandleFunc("/forgot", handlers.homeHandler)
	http.HandleFunc("/login", handlers.homeHandler)
	http.HandleFunc("/private", handlers.homeHandler)
	http.HandleFunc("/register", handlers.homeHandler)
	http.HandleFunc("/response", handlers.homeHandler)
	http.HandleFunc("/schools", handlers.homeHandler)
	http.HandleFunc("/weddings", handlers.homeHandler)

	log.Println("Server Established...")
	http.ListenAndServe(":8080", nil)

}
