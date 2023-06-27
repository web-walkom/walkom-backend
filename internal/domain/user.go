package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewUser struct {
	Email     string `json:"email" binding:"required"`
	Photo     string `json:"photo" binding:"required"`
	FirstName string `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string `json:"last_name" bson:"last_name" binding:"required"`
	CreatedAt int64  `json:"created_at" bson:"created_at" binding:"required"`
}

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" binding:"required"`
	Photo     string             `json:"photo" binding:"required"`
	FirstName string             `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string             `json:"last_name" bson:"last_name" binding:"required"`
	CreatedAt int64              `json:"created_at" bson:"created_at" binding:"required"`
}

type UpdateUser struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Photo     string             `json:"photo" binding:"required"`
	FirstName string             `json:"first_name" bson:"first_name" binding:"required"`
	LastName  string             `json:"last_name" bson:"last_name" binding:"required"`
}

type ResultUpdateUser struct {
	Status bool `json:"status" binding:"required"`
}
