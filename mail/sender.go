package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(subject, content string, to []string) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name, fromEmailAddress, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name, fromEmailAddress, fromEmailPassword,
	}
}

func (s *GmailSender) SendEmail(subject, content string, to []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", s.name, s.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to

	auth := smtp.PlainAuth("", s.fromEmailAddress, s.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, auth)
}
