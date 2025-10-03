package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type ContactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html")
	
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed to parse form data")
		return
	}
	
	form := ContactForm{
		Name:    strings.TrimSpace(r.FormValue("name")),
		Email:   strings.TrimSpace(r.FormValue("email")),
		Message: strings.TrimSpace(r.FormValue("message")),
	}
	
	// Validate
	if form.Name == "" || form.Email == "" || form.Message == "" {
		sendErrorResponse(w, "All fields are required")
		return
	}
	
	if !isValidEmail(form.Email) {
		sendErrorResponse(w, "Invalid email address")
		return
	}
	
	// TODO: Send email via SMTP or save to database
	// For now, just log it (in production, use proper email service)
	fmt.Printf("New contact form submission:\nName: %s\nEmail: %s\nMessage: %s\n", 
		form.Name, form.Email, form.Message)
	
	// Send success response
	response := `
		<div class="form-response success">
			<strong>[SUCCESS]</strong> Message sent! I'll get back to you soon.
		</div>
	`
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	response := fmt.Sprintf(`
		<div class="form-response error">
			<strong>[ERROR]</strong> %s
		</div>
	`, message)
	
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(response))
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
