package models

type User struct {
	ID    	 primitive.ObjectID `bson:"_id,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	Lastlogin string `json:"lastlogin" bson:"lastlogin,omitempty"`
	Examstaken string `json:"examstaken" bson:"examstaken,omitempty"`
	Role string `json:"role" bson:"role,omitempty"`
}