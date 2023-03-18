package service

import (
	"context"
	"errors"

	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/b0shka/walkom-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	return &UsersService{repo: repo}
}

func (s *UsersService) CreateUserIfNotExist(ctx context.Context, email string) error {
	_, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return s.repo.CreateUser(ctx, email)
		}
		return err
	}

	return nil
}

func (s *UsersService) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *UsersService) GetUserById(ctx context.Context, id primitive.ObjectID) (domain.User, error) {
	return s.repo.GetUserById(ctx, id)
}
