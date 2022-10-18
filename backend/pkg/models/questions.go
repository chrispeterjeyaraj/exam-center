package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Option struct {
	OptionId string `json:"optionid"`
	Option   string `json:"option"`
}

type Question struct {
	ID          primitive.ObjectID `bson:"_id"`
	Question_id string             `json:"question_id"`
	UserId      string             `json:"userid"`
	ExamId      string             `json:"examid"`
	Question    string             `json:"question"`
	Options     []Option           `json:"options"`
	Answer      string             `json:"answer"`
	Category    string             `json:"category"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
}

type QuestionView struct {
	Question_id string   `json:"question_id"`
	UserId      string   `json:"userid"`
	ExamId      string   `json:"examid"`
	Question    string   `json:"question"`
	Options     []Option `json:"options"`
	Answer      string   `json:"answer"`
	Category    string   `json:"category"`
}
