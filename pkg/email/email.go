package email

type EmailService struct {
	Name string
	FromEmail string
	FromPassword string
	Host string
	Port int
}

func NewEmailService(name, fromEmail, fromPassword, host string, port int) *EmailService {
	return &EmailService{
		Name: name,
		FromEmail: fromEmail,
		FromPassword: fromPassword,
		Host: host,
		Port: port,
	}
}