package controllers

import (
	"net/http"

	database "github.com/chrispeterjeyaraj/exam-center/backend/database"
	"github.com/gorilla/mux"
)

func HandleDeleteExam(response http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	params := mux.Vars(request)
	examId := params["examId"]

	deleteErr := database.HandleDeleteExam(examId, "exams")

	if deleteErr {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("{\"message\": \"Exam deleted successfully\"}"))

	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"could not delete data\"}"))
		return
	}

}
