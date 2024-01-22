package service

import (
	"context"
	"errors"
	"time"

	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/b0shka/walkom-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) CreateUserIfNotExist(ctx context.Context, email string) error {
	_, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return s.repo.CreateUser(ctx, domain.NewUser{
				Email:     email,
				CreatedAt: time.Now().Unix(),
			})
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

func (s *UsersService) UpdateUser(ctx context.Context, user domain.UpdateUser) error {
	return s.repo.UpdateUser(ctx, user)
}
