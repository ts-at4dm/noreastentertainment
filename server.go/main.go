package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ts-at4dm/noreastentertainment/contactapi"
)

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		now := time.Now().Format(time.RFC1123)
		url := r.URL.Path


		if strings.HasSuffix(url, ".html") || url == "/" {
			log.Printf("Request recieved:\nTime: %s\nip: %s\nPage: %s\n\n\n", now, ip, url)
		}

		next.ServeHTTP(w,r)
	})
}

func main() {
	http.Handle("/", requestLogger(http.FileServer(http.Dir("../client/public"))))

	contactapi.SetupRoutes()

	log.Println("Server Established on Port :5500")

	err := http.ListenAndServe(":5500", nil)
	if err != nil {
		log.Fatal(err)
	}
}
