package contactapi

import (
	"fmt"
	"encoding/json"
	"net/http"
	"net/smtp"
	"strings"
	"time"
	"os"

	"github.com/joho/godotenv"
	"github.com/ts-at4dm/noreastentertainment/crmapi"
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

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Create a new instance of ContactForm with form data
		formData := ContactForm{
			FirstName:     r.FormValue("first_name"),
			LastName:      r.FormValue("last_name"),
			Phone:         r.FormValue("phone"),
			Email:         r.FormValue("email"),
			DateTime:      r.FormValue("date_time"),
			EventType:     r.FormValue("event_type"),
			Services: 	   r.Form["services"], 
			Message:       r.FormValue("message"),
		}

		// Connect to the database
		db := crmapi.ConnectDB()
		defer db.Close() 

		var err error

		// Prepare the SQL statement
		query := "INSERT INTO contacts (first_name, last_name, phone, email, date_time, event_type, services, message) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
		_, err = db.Exec(query,
			formData.FirstName,
			formData.LastName,
			formData.Phone,
			formData.Email,
			formData.DateTime,
			formData.EventType,
			joinEventServices(formData.Services), // Convert slice to a single string
			formData.Message,
		)

		if err != nil {
			http.Error(w, "Failed to insert data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Send email notification
		if err := sendEmail(formData); err != nil {
			http.Error(w, "Failed to send email: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to response page
		http.Redirect(w, r, "/response.html", http.StatusSeeOther)
		return

	}

	// Serve the contact form HTML file for non-POST requests
	http.ServeFile(w, r, "../client/public.contact.html")
}

func joinEventServices(services []string) string {
    jsonServices, err := json.Marshal(services)
    if err != nil {
        return "[]"
    }
    return string(jsonServices)
}

func init() {
	godotenv.Load("./.env")
}


func sendEmail(formData ContactForm) error {
	// Load env vars
	from := os.Getenv("GMAIL_EMAIL")
	password := os.Getenv("GMAIL_APP_PASSWORD")
	if from == "" || password == "" {
		return fmt.Errorf("GMAIL_EMAIL or GMAIL_APP_PASSWORD not set")
	}

	// Create a notification email to yourself
	toAdmin := from // Send to yourself
	
	// Parse and format the DateTime (handle empty date or parsing errors gracefully)
	eventDateFormatted := "Not specified"
	if formData.DateTime != "" {
		eventTime, err := time.Parse("2006-01-02T15:04", formData.DateTime)
		if err == nil {
			eventDateFormatted = eventTime.Format("January 2, 2006 at 15:04")
		}
	}

	// Prepare admin notification message
	adminSubject := "Subject: New Contact Form Submission\r\n"
	mime := "MIME-version: 1.0;\r\nContent-Type: text/plain; charset=\"UTF-8\";\r\n"
	
	adminBody := fmt.Sprintf(
		"You've received a new contact form submission:\n\n"+
			"First Name: %s\nLast Name: %s\nPhone: %s\nEmail: %s\n"+
			"Event Date: %s\nEvent Type: %s\nServices: %s\n\nMessage:\n%s",
		formData.FirstName, formData.LastName, formData.Phone, formData.Email,
		eventDateFormatted, formData.EventType, strings.Join(formData.Services, ", "), formData.Message,
	)
	
	adminMsg := []byte(adminSubject + mime + "\r\n" + adminBody)

	// SMTP Setup
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send admin notification email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toAdmin}, adminMsg)
	if err != nil {
		return fmt.Errorf("failed to send admin notification email: %v", err)
	}
	
	// Only send confirmation to submitter if they provided an email
	if formData.Email != "" && formData.Email != from {
		// Prepare confirmation message for the submitter
		confirmSubject := "Subject: Thank you for contacting Noreast Entertainment\r\n"
		confirmBody := fmt.Sprintf(
			"Dear %s,\n\nThank you for contacting Noreast Entertainment. We have received your submission and will get back to you shortly.\n\n"+
				"Here's a summary of your request:\n"+
				"Event Date: %s\nEvent Type: %s\nServices: %s\n\n"+
				"Best regards,\nNoreast Entertainment Team",
			formData.FirstName, eventDateFormatted, formData.EventType, strings.Join(formData.Services, ", "),
		)
		
		confirmMsg := []byte(confirmSubject + mime + "\r\n" + confirmBody)
		
		// Send confirmation email - don't return error if this fails
		smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{formData.Email}, confirmMsg)
	}
	
	return nil
}
// sendEmail sends an email using the Mailtrap API with the provided contact form data.
// func sendEmail(formData ContactForm) error {
// 	url := "https://sandbox.api.mailtrap.io/api/send/3496565"
// 	method := "POST"
	
// 	// Logic for the API key
// 	apiKey := os.Getenv("MAILTRAP_API_KEY")
// 	if apiKey == "" {
// 		return fmt.Errorf("MAILTRAP_API_KEY is not set")
// 	}
	
// 	// Parse and format the DateTime
// 	eventTime, err := time.Parse("2006-01-02T15:04", formData.DateTime)
// 	if err != nil {
// 		return fmt.Errorf("invalid date format: %v", err)
// 	}
// 	// Format the DateTime as "Month Day Year at Time in 24hr."
// 	eventDateFormatted := eventTime.Format("January 2nd 2006") + " at " + eventTime.Format("15:04")
	
// 	// Construct the payload for the email
// 	payload := map[string]interface{}{
// 		"from": map[string]string{
// 			"email": "hello@example.com",
// 			"name":  "Mailtrap Test",
// 		},
// 		"to": []map[string]string{
// 			{"email": formData.Email},
// 		},
// 		"subject": "Contact Form Submission",
// 		"text": fmt.Sprintf(
// 			"First Name: %s\nLast Name: %s\nPhone: %s\nEmail: %s\nEvent Date: %s\nEvent Type: %s\nServices: %s\nMessage:\n\n%s",
// 			formData.FirstName, formData.LastName, formData.Phone, formData.Email,
// 			eventDateFormatted, formData.EventType, strings.Join(formData.Services, ", "), formData.Message,
// 		),
// 		"category": "Contact Form",
// 	}
	
// 	// Convert payload to JSON
// 	jsonData, err := json.Marshal(payload)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal JSON: %v", err)
// 	}
	
// 	// Create HTTP request
// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return err
// 	}
	

// 	// Add the headers
// 	req.Header.Add("Authorization", apiKey)
// 	req.Header.Add("Content-Type", "application/json")
	
// 	// Send the request
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		body, _ := ioutil.ReadAll(res.Body)
// 		return fmt.Errorf("failed to send email: %s", body)
// 	}

// 	return nil
// }