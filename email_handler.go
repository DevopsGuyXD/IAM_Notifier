package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"os/exec"

	util "github.com/DevopsGuyXD/IAM-Access-Key-Rotation/Utils"
	gomail "gopkg.in/mail.v2"
)

//======================= EMAIL HANDLER =======================
func email_handler(user_name string) {

	util.InitEnvFile()

	account_id, err := exec.Command("aws", "sts", "get-caller-identity", "--query", "Account", "--output", "text").Output(); util.CheckForNil(err)

	email := gomail.NewMessage()

	email_subject := "AWS securty - Rotate access key"
	email_body := fmt.Sprintf("<h2>Greetings user,</h2><p>You are receiving this email as a gentle reminder to decommission your existing <b>Access Key</b> as it is now older than 90 days.</p><p>This is currently above the proposed compliance limit. We trust you will do the needful at the earliest.</p><p><b>Account ID: </b>%v</p>", string(account_id))
	email_connection := os.Getenv("EMAIL_CONNECTION")
	port := 587

	email.SetHeader("From", os.Getenv("EMAIL_SENDER_ID"))
	email.SetHeader("To", user_name)
	email.SetHeader("Subject", email_subject)
	email.SetBody("text/html", email_body)

	d := gomail.NewDialer(email_connection, port, os.Getenv("EMAIL_SENDER_ID"), os.Getenv("EMAIL_SENDER_PASSWORD"))

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = d.DialAndSend(email)
	util.CheckForNil(err)
}