package gate

import (
	"context"
	"fmt"
	"net/smtp"
	"os"
	"time"
)

type EmailGate struct {
	auth smtp.Auth
}

func NewEmailGate() *EmailGate {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_HOST"),
	)

	return &EmailGate{
		auth: auth,
	}
}

func (e *EmailGate) SendNewIPLoginNotificationEmail(ctx context.Context, ipAddress string, email string, timestamp time.Time) error {

	to := email
	from := os.Getenv("SMTP_USERNAME")
	mesage := []byte("Log in from new IP address: " + ipAddress + " at " + timestamp.String())

	addr := fmt.Sprintf("%s:%s", os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"))
	err := smtp.SendMail(addr, e.auth, from, []string{to}, mesage)

	return err
}
