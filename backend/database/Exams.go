package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/chrispeterjeyaraj/exam-center/backend/configs"
	models "github.com/chrispeterjeyaraj/exam-center/backend/pkg/models"
)

func HandleExamsCreate(CollectionName string, examname string, category string, status string, examownerid string, examid string, created time.Time, updated time.Time) bool {
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

func HandleUpdateExam(CollectionName string, exam_id string, examname string, category string, status string, updated time.Time) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	_, err := collection.UpdateOne(ctx, bson.M{"exam_id": exam_id}, bson.M{"$set": bson.M{"examname": examname, "category": category, "status": status, "updated_at": updated}})

	if err != nil {
		return false
	}

	defer cancel()
	return true
}

func GetAllExams(CollectionName string) (bool, []models.ExamView) {
	var results []models.ExamView

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	allExams, errFind := collection.Find(ctx, bson.M{})

	if errFind != nil {
		return false, nil
	}

	for allExams.Next(ctx) {
		var singleExam models.ExamView
		if errFind = allExams.Decode(&singleExam); errFind != nil {
			return false, nil
		}

		results = append(results, singleExam)
	}

	defer cancel()
	return true, results
}

func GetAnExam(exam_id string, CollectionName string) (bool, models.ExamView) {
	var result models.ExamView

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	errFind := collection.FindOne(ctx, bson.M{"exam_id": exam_id}).Decode(&result)

	if errFind != nil {
		return false, result
	}

	defer cancel()
	return true, result
}

func HandleDeleteExam(exam_id string, CollectionName string) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	_, err := collection.DeleteOne(ctx, bson.M{"exam_id": exam_id})

	if err != nil {
		return false
	}

	defer cancel()
	return true
}
