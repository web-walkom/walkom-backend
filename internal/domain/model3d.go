package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Model3D struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Model     string             `json:"model" binding:"required"`
	Latitude  float64            `json:"latitude" binding:"required"`
	Longitude float64            `json:"longitude" binding:"required"`
}
