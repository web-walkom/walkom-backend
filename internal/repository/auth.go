package repository

import (
	"context"
	"errors"

	"github.com/b0shka/walkom-backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepo struct {
	db *mongo.Collection
}

func NewAuthRepo(db *mongo.Database) *AuthRepo {
	return &AuthRepo{
		db: db.Collection(verifyEmailsCollection),
	}
}

func (r *AuthRepo) AddVerifyEmail(ctx context.Context, verifyEmail domain.NewVerifyEmail) error {
	_, err := r.db.InsertOne(ctx, verifyEmail)
	return err
}

func (r *AuthRepo) GetVerifyEmail(ctx context.Context, inp domain.AuthCode) (domain.VerifyEmail, error) {
	var verifyEmail domain.VerifyEmail

	if err := r.db.FindOne(ctx, inp).Decode(&verifyEmail); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.VerifyEmail{}, domain.ErrSecretCodeInvalid
		}
		return domain.VerifyEmail{}, err
	}

	return verifyEmail, nil
}

func (r *AuthRepo) RemoveVerifyEmail(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
