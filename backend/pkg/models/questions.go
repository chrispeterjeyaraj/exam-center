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
	ID         primitive.ObjectID `bson:"_id"`
	UserId     string             `json:"_userid"`
	ExamId     string             `json:"_examid"`
	Question   string             `json:"question"`
	Options    []Option           `json:"options,omitempty"`
	Answer     string             `json:"answer"`
	Category   string             `json:"category"`
	Tags       []string           `json:"tags"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
