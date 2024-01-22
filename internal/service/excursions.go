package service

import (
	"context"

	"github.com/b0shka/walkom-backend/internal/domain"
	"github.com/b0shka/walkom-backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExcursionsService struct {
	repo repository.Excursions
}

func NewExcursionsService(repo repository.Excursions) *ExcursionsService {
	return &ExcursionsService{repo: repo}
}

func (s *ExcursionsService) GetAllExcursions(ctx context.Context) ([]domain.Excursion, error) {
	return s.repo.GetAllExcursions(ctx)
}

func (s *ExcursionsService) GetExcursionById(ctx context.Context, id primitive.ObjectID) (domain.ExcursionOpen, error) {
	return s.repo.GetExcursionById(ctx, id)
}
