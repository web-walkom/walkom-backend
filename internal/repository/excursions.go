package repository

import (
	"context"

	"github.com/b0shka/walkom-backend/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExcursionsRepo struct {
	db *mongo.Collection
}

func NewExcursionsRepo(db *mongo.Database) *ExcursionsRepo {
	return &ExcursionsRepo{
		db: db.Collection(excursionsCollection),
	}
}

func (r *ExcursionsRepo) GetAllExcursions(ctx context.Context) ([]domain.Excursion, error) {
	var excursions []domain.Excursion

	cur, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &excursions); err != nil {
		return nil, err
	}

	return excursions, nil
}
