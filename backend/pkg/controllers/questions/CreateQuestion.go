package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	database "github.com/chrispeterjeyaraj/exam-center/backend/database"
	model "github.com/chrispeterjeyaraj/exam-center/backend/pkg/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleQuestionCreate(response http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	var question model.Question

	err := json.NewDecoder(request.Body).Decode(&question)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	question.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	question.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	question.ID = primitive.NewObjectID()
	question.Question_id = question.ID.Hex()
	insertErr := database.HandleQuestionsCreate("questions", question.Question_id, question.Question, question.Options, question.Answer, question.Category, question.Created_at, question.Updated_at, question.UserId, question.ExamId)

	if insertErr {
		response.WriteHeader(http.StatusOK)
		// json.NewEncoder(response).Encode(&result)
		response.Write([]byte("{\"message\": \"Question created successfully\"}"))

	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"Duplicate Data\"}"))
		return
	}

}
