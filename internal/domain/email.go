package domain

type EmailSender struct {
	Name string
	FromEmailAddress string
	FromEmailPassword string
}

type EmailConfig struct {
	Subject string
	Content string
}