package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/chrispeterjeyaraj/exam-center/backend/configs"
)

func HandleExamsCreate(DBname string, CollectionName string, email string, phone int, password string, fname string, lname string, uid string, created time.Time, updated time.Time, token string, code int, agent interface{}) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	_, errInsert := collection.InsertOne(ctx, bson.M{
		"email":      email,
		"phone":      phone,
		"password":   password,
		"first_name": fname,
		"last_name":  lname,
		"user_id":    uid,
		"created_at": created,
		"updated_at": updated,
		"UserAgent":  agent,
	})

	if errInsert != nil {
		return false
	}
	defer cancel()
	return true

}

func HandleUpdateExam(DBName string, CollectionName string, email string, password string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	_, err := collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"password": password}})

	if err != nil {
		return false
	}

	defer cancel()
	return true
}
