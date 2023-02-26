package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Excursion struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" binding:"required"`
	Description string             `json:"description" binding:"required"`
	Photo       string             `json:"photo" binding:"required"`
	Price       int32              `json:"price" binding:"required"`
}
