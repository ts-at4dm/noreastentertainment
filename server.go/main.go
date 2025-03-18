package main

import (
	"log"
	"net/http"
	"github.com/ts-at4dm/noreastentertainment/handlers"
)



func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/account", handlers.AccountHandler)
	http.HandleFunc("/calendar", handlers.CalendarHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)
	http.HandleFunc("/corp", handlers.CorpHandler)
	http.HandleFunc("/forgot", handlers.ForgotHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/private", handlers.PrivateHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/response", handlers.ResponseHandler)
	http.HandleFunc("/schools", handlers.SchoolsHandler)
	http.HandleFunc("/weddings", handlers.WeddingsHandler)

	log.Println("Server Established...")
	http.ListenAndServe(":8080", nil)

}
