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
	SendEmail(
		subject string,
		content string,
		to string,
		cc string,
		bcc string,
		attachFiles []string,
	) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name, fromEmailAddress, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(
	subject string,
	content string,
	to string,
	cc string,
	bcc string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.Subject = subject
	e.Text = []byte(content)
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.To = []string{to}
	e.Cc = []string{cc}
	e.Bcc = []string{bcc}
	for _, file := range attachFiles {
		_, err := e.AttachFile(file)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", file, err)
		}
	}

	auth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, auth)
}
