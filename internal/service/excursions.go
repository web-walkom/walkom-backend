package service

import (
	"context"

	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/b0shka/walkom-backend/internal/repository"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type ExcursionsService struct {
	repo repository.Excursions
}

func NewExcursionsService(repo repository.Excursions) *ExcursionsService {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	return &ExcursionsService{repo: repo}
}

func (s *ExcursionsService) GetAllExcursions(ctx context.Context) ([]domain.Excursion, error) {
	return s.repo.GetAllExcursions(ctx)
}
