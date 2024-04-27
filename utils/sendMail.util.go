package utils

import (
	"crypto/tls"
	"fmt"
	"log"

	gomail "gopkg.in/mail.v2"
)

var fromEmail string = "Anubhav Singh <mailer@anubhavsingh.dev>"
var authUser string = "mailer@anubhavsingh.dev"
var authPassword string = "Anubhav@Singh"
var smtpHost string = "mail.anubhavsingh.dev"

func SendMail(toMail string, subject string, htmlBody string) {
	m := gomail.NewMessage()

	m.SetHeader("From", fromEmail)
	m.SetHeader("To", toMail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer(smtpHost, 465, authUser, authPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.Println("Mail Sent")

}
