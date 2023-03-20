package repository

import (
	"context"
	"errors"

	"github.com/b0shka/walkom-backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepo struct {
	db *mongo.Collection
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{
		db: db.Collection(usersCollection),
	}
}

func (r *UsersRepo) CreateUser(ctx context.Context, user domain.NewUser) error {
	_, err := r.db.InsertOne(ctx, user)
	return err
}

func (r *UsersRepo) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	if err := r.db.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return user, nil
}

func (r *UsersRepo) GetUserById(ctx context.Context, id primitive.ObjectID) (domain.User, error) {
	var user domain.User

	if err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return user, nil
}
