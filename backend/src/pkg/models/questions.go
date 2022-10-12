package models

type Question struct {
	ID       int64  `json:"_id,omitempty" bson:"_id,omitempty"`
    Question 	string `json:"question" bson:"question,omitempty"`
	Options     struct {
		OptionId string `json:"optionid" bson:"optionid"`
		Option string `json:"option" bson:"option"`
	} `json:"options,omitempty" bson:"options,omitempty"`
    CorrectAnswer string `json:"correctnanswer" bson:"correctanswer,omitempty"`
}