package service

import (
	"context"

	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/b0shka/walkom-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth interface {
	CreateVerifyEmail(ctx context.Context, email string, secret_code int32) error
	CheckSecretCode(ctx context.Context, data domain.AuthCode) error
	CreateSession(ctx context.Context, id primitive.ObjectID) (domain.UserToken, error)
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

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Auth:       NewAuthService(repos.Auth),
		Users:      NewUsersService(repos.Users),
		Excursions: NewExcursionsService(repos.Excursions),
	}
}
