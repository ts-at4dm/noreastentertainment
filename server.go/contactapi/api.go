package contactapi

import (
	"fmt"
	"net/http"
	"net/smtp"
)

type ContactForm struct {
	FirstName	string
	LastName	string
	Phone		string
	Email		string
	DateTime	string
	EventType	string
	Services	[]string
	message		string
}

func SetupRoutes() {
	http.HandleFunc("/contact", contactHandler)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		formData := ContactForm{
			FirstName:	r.FormValue("first_name"),
			LastName:	r.FormValue("last_name"),
			Phone:		r.FormValue("phone"),
			Email:		r.FormValue("email"),
			DateTime:	r.FormValue("date_time"),
			EventType:	r.FormValue("event_type"),
			Services:	r.Form["services"],
			Message:	r.FormValue("message"),
		}
		sendEmail(formData)
		http.Redirect(w, r, "/response.html", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "../client/public.contact.html")
}

func sendEmail(formData ContactForm) {
	from := "****************"
	password := "************"
	to := "*****************"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Contact Form Submission"

	services List := ""
	for _, service := range formData.Services {
		servicesList += service + ", "
	}
	if len(servicesList) > 0 {
		servicesList = servicesList[:len(servicesList)-2]
	}

	body := "First Name: " + formData.FirstName + "\n" +
			"Last Name: " + formData.LastName + "\n" +
			"Phone: " + formData.Phone + "\n" +
			"Email: " + formData.Email + "\n" +
			"Event Date: " + formData.EventDate + "\n" +
			"Event Type: " + formData.EventType + "\n" +
			"Services: " + servicesList + "\n" +
			"Message:\n" + formData.Message
	
	msg := []byte(body)

	auth := smtp.PlainAuth("", from, password, smtp)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		log.Fatal(err)
	}
}