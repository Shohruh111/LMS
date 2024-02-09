package helper

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "shohruxramazonov564@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "shohruxramazonov564@gmail.com", "Shohrux$2203")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
