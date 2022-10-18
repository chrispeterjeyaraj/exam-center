package controllers

import (
	"encoding/json"
	"net/http"

	database "github.com/chrispeterjeyaraj/exam-center/backend/database"
	"github.com/gorilla/mux"
)

func GetAnExam(response http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	params := mux.Vars(request)
	examId := params["examId"]

	if examId == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("{\"message\": \"exam id missing from query parameter\"}"))
		return
	}

	exam, result := database.GetAnExam(examId, "exams")

	if !exam {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"No Exam Found\"}"))
		return
	} else if exam {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(&result)
		return
	}

}
