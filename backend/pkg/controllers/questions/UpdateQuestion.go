package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	database "github.com/chrispeterjeyaraj/exam-center/backend/database"
	model "github.com/chrispeterjeyaraj/exam-center/backend/pkg/models"
)

func HandleQuestionsUpdate(response http.ResponseWriter, request *http.Request) {

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

	question.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateErr := database.HandleUpdateQuestion("questions", question.Question_id, question.Question, question.Options, question.Answer, question.Category, question.Updated_at)

	if updateErr {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("{\"message\": \"Question updated successfully\"}"))

	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"Duplicate Data\"}"))
		return
	}

}
