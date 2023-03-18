package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VerifyEmail struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email      string             `json:"email" binding:"required"`
	SecretCode int32              `json:"secret_code" bson:"secret_code" binding:"required"`
	CreatedAt  int64              `json:"created_at" bson:"created_at" binding:"required"`
	ExpiredAt  int64              `json:"expired_at" bson:"expired_at" binding:"required"`
}
