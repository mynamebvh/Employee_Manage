package services

import (
	"crypto/tls"
	"employee_manage/config"

	"gopkg.in/gomail.v2"
)

func SendMail(to string, subject string, content string) (err error) {

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", config.ConfigApp.Mail.Email)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", content)

	// Settings for SMTP server
	d := gomail.NewDialer(
		config.ConfigApp.Mail.SMTP,
		config.ConfigApp.Mail.Port,
		config.ConfigApp.Mail.Email,
		config.ConfigApp.Mail.Password,
	)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err = d.DialAndSend(m); err != nil {
		return
	}

	return
}
