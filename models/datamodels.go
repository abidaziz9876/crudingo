package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataModels struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name"`
	Age        int                `json:"age"`
	Salary     int                `json:"salary"`
	Department string             `json:"department"`
}
