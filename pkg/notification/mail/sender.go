package mail

import (
	"fmt"
	"github.com/jordan-wright/email"
	"go-layout/config"
	"net/smtp"
)

type EmailSender interface {
	Send(subject, content string, to, cc, bcc, attachFiles []string) error
}

type Email struct {
	config *config.MailConfig
}

func NewEmailSender(config *config.MailConfig) EmailSender {
	return &Email{
		config: config,
	}
}

func (sender *Email) Send(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.config.FromName, sender.config.FromEmail)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, file := range attachFiles {
		if _, err := e.AttachFile(file); err != nil {
			return fmt.Errorf("failed to attach file %s: %w", file, err)
		}
	}

	smtpServerAddress := fmt.Sprintf("%s:%s", sender.config.Host, sender.config.Port)
	smtpAuth := smtp.PlainAuth("", sender.config.Username, sender.config.Password, sender.config.Host)

	if err := e.Send(smtpServerAddress, smtpAuth); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
