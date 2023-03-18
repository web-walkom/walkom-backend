package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `json:"email" binding:"required"`
}

type UserToken struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AccessToken string             `json:"access_token" binding:"required"`
}