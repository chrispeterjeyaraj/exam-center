package controllers

import (
	"encoding/json"
	"net/http"

	database "github.com/chrispeterjeyaraj/exam-center/backend/database"
	model "github.com/chrispeterjeyaraj/exam-center/backend/pkg/models"
)

func HandleSignOut(response http.ResponseWriter, request *http.Request) {
	AuthToken := request.Header.Get("Authenticate")

	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	} else if AuthToken == "" {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Authorization Token is not provided\"}"))
		return
	}

	var code model.Code

	err := json.NewDecoder(request.Body).Decode(&code)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("{\"message\": \"Invalid request body\"}"))
		return
	}

	errDel := database.HandleRemoveCode("tokens", code.Code, AuthToken)

	if errDel {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("{\"message\": \"User Signed Out Successfully\"}"))

	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"Sign out failed\"}"))
		return
	}
}
