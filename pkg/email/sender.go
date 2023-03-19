package email

import (
	"fmt"
	"github.com/b0shka/walkom-backend/internal/domain"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAddress       = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

func (s *EmailService) SendEmail(config domain.EmailConfig, toEmail string) error {
	e := email.NewEmail()

	e.From = fmt.Sprintf("%s <%s>", s.sender.Name, s.sender.FromEmailAddress)
	e.Subject = config.Subject
	e.HTML = []byte(config.Content)
	e.To = []string{toEmail}

	smtpAuth := smtp.PlainAuth(
		"",
		s.sender.FromEmailAddress,
		s.sender.FromEmailPassword,
		smtpAddress,
	)
	return e.Send(smtpServerAddress, smtpAuth)
}
