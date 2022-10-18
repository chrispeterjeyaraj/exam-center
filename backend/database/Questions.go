package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/chrispeterjeyaraj/exam-center/backend/configs"
	"github.com/chrispeterjeyaraj/exam-center/backend/pkg/models"
)

func HandleQuestionsCreate(CollectionName string, questionid string, question string, options []models.Option, answer string, category string, created time.Time, updated time.Time, userid string, examid string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	_, errInsert := collection.InsertOne(ctx, bson.M{
		"question_id": questionid,
		"question":    question,
		"options":     options,
		"answer":      answer,
		"category":    category,
		"created_at":  created,
		"updated_at":  updated,
		"userid":      userid,
		"examid":      examid,
	})

	if errInsert != nil {
		return false
	}
	defer cancel()
	return true

}

func HandleUpdateQuestion(CollectionName string, questionid string, question string, options []models.Option, answer string, category string, updated time.Time) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	_, err := collection.UpdateOne(ctx, bson.M{"question_id": questionid}, bson.M{"$set": bson.M{
		"question_id": questionid,
		"question":    question,
		"options":     options,
		"answer":      answer,
		"category":    category,
		"updated_at":  updated,
	}})

	if err != nil {
		return false
	}

	defer cancel()
	return true
}

func GetAllQuestions(CollectionName string) (bool, []models.QuestionView) {
	var results []models.QuestionView

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	allquestions, errFind := collection.Find(ctx, bson.M{})

	if errFind != nil {
		return false, nil
	}

	for allquestions.Next(ctx) {
		var singlequestion models.QuestionView
		if errFind = allquestions.Decode(&singlequestion); errFind != nil {
			return false, nil
		}

		results = append(results, singlequestion)
	}

	defer cancel()
	return true, results
}

func GetQuestion(question_id string, CollectionName string) (bool, models.QuestionView) {
	var result models.QuestionView

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	errFind := collection.FindOne(ctx, bson.M{"question_id": question_id}).Decode(&result)

	if errFind != nil {
		return false, result
	}

	defer cancel()
	return true, result
}

func HandleDeleteQuestion(question_id string, CollectionName string) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	_, err := collection.DeleteOne(ctx, bson.M{"question_id": question_id})

	if err != nil {
		return false
	}

	defer cancel()
	return true
}
