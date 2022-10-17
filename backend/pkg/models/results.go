package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Result struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserId     primitive.ObjectID `bson:"_userid,omitempty"`
	ExamId     primitive.ObjectID `bson:"_examid,omitempty"`
	Score      int64              `json:"score" bson:"score,omitempty"`
	Result     string             `json:"result" bson:"result,omitempty"`
	Examstaken string             `json:"examstaken" bson:"examstaken,omitempty"`
}
