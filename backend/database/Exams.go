package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/chrispeterjeyaraj/exam-center/backend/configs"
)

func HandleExamsCreate(DBname string, CollectionName string, examname string, category string, status string, examownerid string, examid string, created time.Time, updated time.Time) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	_, errInsert := collection.InsertOne(ctx, bson.M{
		"exam_id":     examid,
		"examname":    examname,
		"category":    category,
		"status":      status,
		"examownerid": examownerid,
		"created_at":  created,
		"updated_at":  updated,
	})

	if errInsert != nil {
		return false
	}
	defer cancel()
	return true

}

func HandleUpdateExam(DBName string, CollectionName string, exam_id string, examname string, category string, status string, updated time.Time) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	_, err := collection.UpdateOne(ctx, bson.M{"exam_id": exam_id}, bson.M{"$set": bson.M{"examname": examname, "category": category, "status": status, "updated_at": updated}})

	if err != nil {
		return false
	}

	defer cancel()
	return true
}
