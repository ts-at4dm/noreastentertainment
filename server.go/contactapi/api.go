package contactapi

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil" 
	"time"
)

// ContactForm reps the structure of the contact form
type ContactForm struct {
	FirstName	string		
	LastName	string		
	Phone		string		
	Email		string		
	DateTime	string		
	EventType	string		
	Services	[]string		
	Message		string		
}

// SetupRoutes inits the routes for the contact api
func SetupRoutes() {
	http.HandleFunc("/contact", contactHandler)
}

// contactHandler processes the incoming contact form submission
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
		
		if err := sendEmail(formData); err != nil {
		http.Error(w, "Failed to send: "+err.Error(),
		http.StatusInternalServerError)
			return
		}
		
		http.Redirect(w, r, "/response.html", http.StatusSeeOther)
		return

	}

	http.ServeFile(w, r, "../client/public.contact.html")
}

// sendEmail sends an email using the Mailtrap API with the provided contact form data.
func sendEmail(formData ContactForm) error {
	url := "https://sandbox.api.mailtrap.io/api/send/3496565"
	method := "POST"

	// Parse and format the DateTime
	eventTime, err := time.Parse("2006-01-02T15:04", formData.DateTime)
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}
	// Format the DateTime as "March 29th 2025 at 23:10"
	eventDateFormatted := eventTime.Format("January 2nd 2006") + " at " + eventTime.Format("15:04")

	// Construct the payload for the email
	payload := fmt.Sprintf(`{
		"from": {
			"email": "hello@example.com",
			"name": "Mailtrap Test"
		},
		"to": [
			{
				"email": "%s"
			}
		],
		"subject": "Contact Form Submission",
		"text": "First Name: %s\nLast Name: %s\nPhone: %s\nEmail: %s\nEvent Date: %s\nEvent Type: %s\nServices: %s\nMessage:\n\n%s",
		"category": "Contact Form"
	}`, formData.Email, formData.FirstName, formData.LastName, formData.Phone, formData.Email, eventDateFormatted, formData.EventType, strings.Join(formData.Services, ", "), formData.Message)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		return err
	}

	// Add the headers
	req.Header.Add("Authorization", "Bearer 7912dca21d7e6e41ae931c7f64f52b8c")
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("failed to send email: %s", body)
	}

	return nil
}