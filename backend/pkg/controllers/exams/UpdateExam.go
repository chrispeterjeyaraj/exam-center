package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	database "github.com/chrispeterjeyaraj/exam-center/backend/database"
	model "github.com/chrispeterjeyaraj/exam-center/backend/pkg/models"
)

func HandleExamsUpdate(response http.ResponseWriter, request *http.Request) {

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

	exam.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateErr := database.HandleUpdateExam("exams", exam.Exam_id, exam.Examname, exam.Category, exam.Status, exam.Updated_at)

	if updateErr {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("{\"message\": \"Exam updated successfully\"}"))

	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"Duplicate Data\"}"))
		return
	}

}
