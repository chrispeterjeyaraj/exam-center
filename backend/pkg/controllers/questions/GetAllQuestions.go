package controllers

import (
	"encoding/json"
	"net/http"

	database "github.com/chrispeterjeyaraj/exam-center/backend/database"
)

func GetAllQuestions(response http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	getErr, results := database.GetAllQuestions("questions")

	if getErr {
		response.WriteHeader(http.StatusOK)
		if results == nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("{\"message\": \"No Exam Found\"}"))
			return
		} else {
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(&results)
			return
		}

	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"Duplicate Data\"}"))
		return
	}

}
