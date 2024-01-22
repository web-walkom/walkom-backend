package email

import (
	"fmt"
	"net/smtp"

	"github.com/b0shka/walkom-backend/internal/domain"

	"github.com/jordan-wright/email"
)

func (s *EmailService) SendEmail(config domain.VerifyEmailConfig, toEmail string) error {
	e := email.NewEmail()

	e.From = fmt.Sprintf("%s <%s>", s.Name, s.FromEmail)
	e.Subject = config.Subject
	e.HTML = []byte(config.Content)
	e.To = []string{toEmail}

	smtpAuth := smtp.PlainAuth("", s.FromEmail, s.FromPassword, s.Host)
	return e.Send(fmt.Sprintf("%s:%d", s.Host, s.Port), smtpAuth)
}
