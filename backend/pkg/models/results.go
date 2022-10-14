package models

type User struct {
	ID    	 primitive.ObjectID `bson:"_id,omitempty"`
	UserId   primitive.ObjectID `bson:"_userid,omitempty"`
	ExamId   primitive.ObjectID `bson:"_examid,omitempty"`
	score int64 `json:"score" bson:"score,omitempty"`
	result string `json:"result" bson:"result,omitempty"`
	Examstaken string `json:"examstaken" bson:"examstaken,omitempty"`
}