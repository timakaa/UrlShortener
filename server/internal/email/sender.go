package email

import (
	"fmt"
	"os"
	"strconv"

	gomail "gopkg.in/mail.v2"
)

func SendWelcomeEmail(to, username string) error {
	// Create new message
	m := gomail.NewMessage()
	
	// Set sender
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Welcome to Our Service!")
	
	// Set email body
	body := fmt.Sprintf(`
		<h1>Welcome, %s!</h1>
		<p>Thank you for registering with our service.</p>
		<p>We're excited to have you on board!</p>
	`, username)
	m.SetBody("text/html", body)

	// Create dialer
	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		587, // SMTP port
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
	)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func SendVerificationEmail(to string, code string) error {
	m := gomail.NewMessage()
	
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Verify Your Email")
	
	body := fmt.Sprintf(`
		<h1>Email Verification</h1>
		<p>Your verification code is: <strong>%s</strong></p>
		<p>This code will expire in 15 minutes.</p>
	`, code)
	m.SetBody("text/html", body)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASSWORD"),
	)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
