package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	database "github.com/chrispeterjeyaraj/exam-center/backend/database"
	model "github.com/chrispeterjeyaraj/exam-center/backend/pkg/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleExamCreate(response http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	var exam model.Exam

	err := json.NewDecoder(request.Body).Decode(&exam)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	exam.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	exam.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	exam.ID = primitive.NewObjectID()
	exam.Exam_id = exam.ID.Hex()
	insertErr := database.HandleDatabaseInsert("examcenter", "users", exam.Examname, exam.Category, exam.Status, exam.ExamonwerId, exam.Exam_id, exam.Created_at, exam.Updated_at)

	if insertErr {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(&result)

	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"Duplicate Data\"}"))
		return
	}

}
