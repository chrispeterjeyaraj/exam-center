package models

type User struct {
	ID       int64  `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
}