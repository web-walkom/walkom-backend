package email

import "github.com/b0shka/walkom-backend/internal/domain"

type EmailService struct {
	sender domain.EmailSender
}

func NewEmailService(sender domain.EmailSender) *EmailService {
	return &EmailService{
		sender: sender,
	}
}