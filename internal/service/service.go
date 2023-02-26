package service

import (
	"context"
	"walkom/internal/domain"
	"walkom/internal/repository"
)

type Excursions interface {
	GetAllExcursions(ctx context.Context) ([]domain.Excursion, error)
}

type Services struct {
	Excursions Excursions
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Excursions: NewExcursionsService(repos.Excursions),
	}
}
