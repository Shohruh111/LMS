// package helper

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"log"
// 	"net/smtp"
// )

// func SendEmail(to, subject, body string) string {
// 	from := "shohruxramazonov564@gmail.com"
// 	password := "your-app-password" // Generate an app password: https://myaccount.google.com/apppasswords

// 	// Recipient's email address
// 	to = "recipient@example.com"

// 	// Email content
// 	subject = "Test Email"
// 	body = "This is a test email sent from a Golang application."

// 	// Compose the email message
// 	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body)

// 	// Set up authentication information
// 	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

// 	// Connect to the SMTP server with a secure connection
// 	server := "smtp.gmail.com"
// 	port := 587
// 	connection, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", server, port), nil)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the server: %v", err)
// 		return err.Error()
// 	}

// 	// Send the email
// 	client, err := smtp.NewClient(connection, server)
// 	if err != nil {
// 		log.Fatalf("Failed to create SMTP client: %v", err)
// 		return err.Error()
// 	}
// 	defer client.Close()

// 	if err = client.Auth(auth); err != nil {
// 		log.Fatalf("SMTP authentication failed: %v", err)
// 		return err.Error()
// 	}

// 	if err = client.Mail(from); err != nil {
// 		log.Fatalf("Failed to set sender: %v", err)
// 		return err.Error()
// 	}

// 	if err = client.Rcpt(to); err != nil {
// 		log.Fatalf("Failed to set recipient: %v", err)
// 		return err.Error()
// 	}

// 	writer, err := client.Data()
// 	if err != nil {
// 		log.Fatalf("Failed to open data connection: %v", err)
// 		return err.Error()
// 	}
// 	defer writer.Close()

// 	_, err = writer.Write([]byte(message))
// 	if err != nil {
// 		log.Fatalf("Failed to write email data: %v", err)
// 		return err.Error()
// 	}

// 	log.Println("Email sent successfully!")
// 	return ""
// }

// package helper

// import "gopkg.in/mail.v2"

// func SendEmail() error {

// 	var (
// 		to      string = "shohruxramazonov564@gmail.com"
// 		subject string = "Simple Email"
// 		body    string = "Confirm Message From SMTP Mailing system!"
// 	)

// 	from := "shohruhramozonov2662@gmail.com"
// 	password := "Shohrux$2203$"

// 	m := mail.NewMessage()
// 	m.SetHeader("From", from)
// 	m.SetHeader("To", to)
// 	m.SetHeader("Subject", subject)
// 	m.SetBody("text/plain", body)

// 	d := mail.NewDialer("smtp.gmail.com", 587, from, password)

// 	// Send the email
// 	if err := d.DialAndSend(m); err != nil {
// 		return err
// 	}

// 	return nil
// }

package helper

import (
	"math/rand"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string) (int, error) {
	// Replace these variables with your own values
	fromEmail := "shohruhramozonov2662@gmail.com"
	emailPassword := "fwve ustr jawm hvjq "

	min := 100000
	max := 999999

	verifyCode := rand.Intn(max-min+1) + min

	// Set up the email message
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Verification Code!")
	m.SetBody("text/html", strconv.Itoa(verifyCode))

	// Set up the email sender
	d := gomail.NewDialer("smtp.gmail.com", 587, fromEmail, emailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return 0, err
	}

	return verifyCode, nil

}
