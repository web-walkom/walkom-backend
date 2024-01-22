package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Waypoint struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Latitude  float64            `json:"latitude" binding:"required"`
	Longitude float64            `json:"longitude" binding:"required"`
	Audio     string             `json:"audio" binding:"required"`
}
