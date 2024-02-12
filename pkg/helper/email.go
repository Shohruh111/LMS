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
