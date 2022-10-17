package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Option struct {
	OptionId string `json:"optionid" bson:"optionid"`
	Option   string `json:"option" bson:"option"`
}

type Question struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserId   primitive.ObjectID `bson:"_userid,omitempty"`
	ExamId   primitive.ObjectID `bson:"_examid,omitempty"`
	Question string             `json:"question" bson:"question,omitempty"`
	Options  []Option           `json:"options,omitempty" bson:"options,omitempty"`
	Answer   string             `json:"answer" bson:"answer,omitempty"`
	Category string             `json:"category" bson:"category,omitempty"`
	Tags     []string           `json:"tags" bson:"tags,omitempty"`
}
