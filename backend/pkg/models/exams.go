package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exam struct {
	ID          primitive.ObjectID `bson:"_id"`
	Exam_id     string             `json:"exam_id"`
	Examname    string             `json:"examname"`
	Category    string             `json:"category"`
	Status      string             `json:"status"`
	ExamonwerId string             `json:"examownerid"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
}

type ExamView struct {
	Exam_id     string `json:"exam_id"`
	Examname    string `json:"examname"`
	Category    string `json:"category"`
	Status      string `json:"status"`
	ExamonwerId string `json:"examownerid"`
}
