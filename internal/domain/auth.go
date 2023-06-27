package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthEmail struct {
	Email string `json:"email" binding:"required"`
}

type ResultSendCode struct {
	Status bool   `json:"status" binding:"required"`
	Error  string `json:"error"`
}

type AuthCode struct {
	Email      string `json:"email" binding:"required"`
	SecretCode int32  `json:"secret_code" bson:"secret_code" binding:"required"`
}

type ResultCheckCode struct {
	Status      bool               `json:"status" binding:"required"`
	Error       string             `json:"error"`
	ID          primitive.ObjectID `json:"id"`
	Email       string             `json:"email"`
	AccessToken string             `json:"access_token"`
}

type AuthToken struct {
	ID          primitive.ObjectID `json:"id" binding:"required"`
	AccessToken string             `json:"access_token" binding:"required"`
}
