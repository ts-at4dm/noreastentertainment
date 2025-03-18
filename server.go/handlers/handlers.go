package handlers

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/public")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/about.html")
}
func accountHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/account.html")
}

func calendarHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/calendar.html")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/contact.html")
}

func corpHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/corp.html")
}

func forgotHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/forgot.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/login.html")
}

func privateHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/private.html")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/register.html")
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/response.html")
}

func schoolsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/schools.html")
}

func weddingsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/src/pages/weddings.html")
}