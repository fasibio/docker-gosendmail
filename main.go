package main

import (
	"log"
	"os"

	"github.com/fasibio/goSendMail/properties"
	gomail "gopkg.in/gomail.v2"
)

func main() {
	log.Println("Lets go: Set envs: SMTPSERVER, SMTPPORT,EMAIL, EMAILPASSWORD")
	log.Println("command to use `gosendmail <toEmail> <subject> <text>`")

	sendEmail(os.Args[1], os.Args[3], os.Args[2])
	log.Println("Email was send")
}

func sendEmail(to string, text, header string) {
	emailCredentails := properties.GetEmailCredentails()
	m := gomail.NewMessage()
	m.SetHeader("From", emailCredentails.EmailAddress)
	m.SetHeader("To", to)
	m.SetHeader("Subject", header)
	m.SetBody("text/html", text)
	d := gomail.NewPlainDialer(emailCredentails.SMTPServer, emailCredentails.SMTPPort, emailCredentails.EmailAddress, emailCredentails.EmailPassword)
	if err := d.DialAndSend(m); err != nil {
		log.Println("Error:", err)
	}
}
