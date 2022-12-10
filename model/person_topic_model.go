package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PersonTopicModel struct {
	ID   primitive.ObjectID `json:",omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name" validate:"required"`
}
