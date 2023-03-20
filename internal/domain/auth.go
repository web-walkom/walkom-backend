package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthEmail struct {
	Email string `json:"email" binding:"required"`
}

type AuthCode struct {
	Email      string `json:"email" binding:"required"`
	SecretCode int32  `json:"secret_code" bson:"secret_code" binding:"required"`
}

type AuthToken struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AccessToken string             `json:"access_token" binding:"required"`
}
