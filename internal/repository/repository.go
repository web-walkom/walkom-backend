package repository

import (
	"context"
	"walkom/internal/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type Excursions interface {
	GetAllExcursions(ctx context.Context) ([]domain.Excursion, error)
}

type Repositories struct {
	Excursions Excursions
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Excursions: NewExcursionsRepo(db),
	}
}
