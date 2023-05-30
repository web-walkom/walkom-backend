package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Waypoint struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Latitude    float64            `json:"latitude" binding:"required"`
	Longitude   float64            `json:"longitude" binding:"required"`
	PlacemarkId int                `json:"placemarkId"`
	Audio       string             `json:"audio" binding:"required"`
}
