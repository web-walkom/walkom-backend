package repository

import (
	"context"
	"errors"
	"time"

	"github.com/b0shka/walkom-backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *AuthRepo) CreateVerifyEmail(ctx context.Context, email string, secret_code int32) error {
	_, err := r.db.InsertOne(ctx, bson.M{
		"email":       email,
		"secret_code": secret_code,
		"created_at":  time.Now().Unix(),
		"expired_at":  time.Now().Unix() + 900,
	})
	return err
}

func (r *AuthRepo) GetVerifyEmail(ctx context.Context, data domain.AuthCode) (domain.VerifyEmail, error) {
	var verifyEmail domain.VerifyEmail

	if err := r.db.FindOne(ctx, data).Decode(&verifyEmail); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.VerifyEmail{}, domain.ErrSecretCodeInvalid
		}
		return domain.VerifyEmail{}, err
	}

	return verifyEmail, nil
}
