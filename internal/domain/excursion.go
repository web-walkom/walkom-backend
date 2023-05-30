package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Excursion struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description" binding:"required"`
	Photos      []string           `json:"photos" binding:"required"`
}

type ExcursionOpen struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description" binding:"required"`
	Photos      []string           `json:"photos" binding:"required"`
	Placemarks  []Placemark        `json:"placemarks" binding:"required"`
	Waypoints   []Waypoint         `json:"waypoints" binding:"required"`
}
