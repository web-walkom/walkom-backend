package service

import (
	"context"

	"github.com/b0shka/walkom-backend/internal/config"
	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/b0shka/walkom-backend/internal/repository"
	"github.com/b0shka/walkom-backend/pkg/email"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth interface {
	SendCodeEmail(ctx context.Context, inp domain.AuthEmail) error
	CheckSecretCode(ctx context.Context, inp domain.AuthCode) error
	CreateSession(id primitive.ObjectID) (domain.AuthToken, error)
	ParseToken(Token string) (string, error)
}

type Users interface {
	CreateUserIfNotExist(ctx context.Context, email string) error
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUserById(ctx context.Context, id primitive.ObjectID) (domain.User, error)
}

type Excursions interface {
	GetAllExcursions(ctx context.Context) ([]domain.Excursion, error)
}

type Services struct {
	Auth       Auth
	Users      Users
	Excursions Excursions
}

type Deps struct {
	Repos *repository.Repositories
	EmailService email.EmailService
	EmailConfig config.EmailConfig
	AuthConfig config.AuthConfig
}

func NewServices(deps Deps) *Services {
	return &Services{
		Auth: NewAuthService(
			deps.Repos.Auth,
			deps.EmailService,
			deps.EmailConfig,
			deps.AuthConfig,
		),
		Users:      NewUsersService(deps.Repos.Users),
		Excursions: NewExcursionsService(deps.Repos.Excursions),
	}
}
