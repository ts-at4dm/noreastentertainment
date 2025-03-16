package main

import (
	"fmt"
	//"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/public/index.html")
}

func main() {
	http.HandleFunc("/", viewHandler)

	fs := http.FileServer(http.Dir("../client/public"))
	http.Handle("/static", http.StripPrefix("/static", fs))

	fd := http.FileServer(http.Dir("../client/src/styles"))
	http.Handle("/static/", http.StripPrefix("/static/", fd))

	fmt.Println("Server is runnning at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
