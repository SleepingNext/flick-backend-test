package common

import (
	"fmt"
	"net/smtp"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_EMAIL = "flicktechtest@gmail.com"
const CONFIG_PASSWORD = "kevin99123"

func SendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + CONFIG_EMAIL + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_EMAIL, CONFIG_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
