package repository

import (
	"context"

	"github.com/b0shka/walkom-backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth interface {
	AddVerifyEmail(ctx context.Context, verifyEmail domain.NewVerifyEmail) error
	GetVerifyEmail(ctx context.Context, inp domain.AuthCode) (domain.VerifyEmail, error)
	RemoveVerifyEmail(ctx context.Context, id primitive.ObjectID) error
}

type Users interface {
	CreateUser(ctx context.Context, user domain.NewUser) error
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUserById(ctx context.Context, id primitive.ObjectID) (domain.User, error)
}

type Excursions interface {
	GetAllExcursions(ctx context.Context) ([]domain.Excursion, error)
	GetExcursionById(ctx context.Context, id primitive.ObjectID) (domain.ExcursionOpen, error)
}

type Repositories struct {
	Auth       Auth
	Users      Users
	Excursions Excursions
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Auth:       NewAuthRepo(db),
		Users:      NewUsersRepo(db),
		Excursions: NewExcursionsRepo(db),
	}
}
