package domain

type EmailSender struct {
	Name string
	FromEmailAddress string
	FromEmailPassword string
}

type EmailData struct {
	Subject string
	Content string
}