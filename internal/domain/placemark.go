package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Placemark struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" binding:"required"`
	Photos    []string           `json:"photos" binding:"required"`
	Latitude  float64            `json:"latitude" binding:"required"`
	Longitude float64            `json:"longitude" binding:"required"`
}
