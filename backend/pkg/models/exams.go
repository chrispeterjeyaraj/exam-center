package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exam struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Examname  string             `json:"examname" bson:"examname,omitempty"`
	Category  string             `json:"category" bson:"category,omitempty"`
	Status    string             `json:"status" bson:"status,omitempty"`
	Examonwer primitive.ObjectID `bson:"_examonwer,omitempty"`
}
