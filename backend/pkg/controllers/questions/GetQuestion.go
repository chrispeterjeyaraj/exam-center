package controllers

import (
	"encoding/json"
	"net/http"

	database "github.com/chrispeterjeyaraj/exam-center/backend/database"
	"github.com/gorilla/mux"
)

func GetQuestion(response http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	params := mux.Vars(request)
	questionId := params["questionId"]

	if questionId == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("{\"message\": \"question id missing from query parameter\"}"))
		return
	}

	question, result := database.GetQuestion(questionId, "questions")

	if !question {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"No Question Found\"}"))
		return
	} else if question {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(&result)
		return
	}

}
