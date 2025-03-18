package handlers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/public")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/about.html")
}
func AccountHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/account.html")
}

func CalendarHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/calendar.html")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/contact.html")
}

func CorpHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/corp.html")
}

func ForgotHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/forgot.html")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/login.html")
}

func PrivateHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/private.html")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/register.html")
}

func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/response.html")
}

func SchoolsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/schools.html")
}

func WeddingsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/weddings.html")
}