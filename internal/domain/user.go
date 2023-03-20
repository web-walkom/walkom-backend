package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewUser struct {
	Email string             `json:"email" binding:"required"`
	CreatedAt int64 `json:"created_at" bson:"created_at" binding:"required"`
}

type User struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `json:"email" binding:"required"`
	CreatedAt int64 `json:"created_at" bson:"created_at" binding:"required"`
}