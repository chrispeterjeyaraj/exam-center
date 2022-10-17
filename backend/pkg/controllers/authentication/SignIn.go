package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	database "github.com/chrispeterjeyaraj/exam-center/backend/database"
	helper "github.com/chrispeterjeyaraj/exam-center/backend/helpers"
	model "github.com/chrispeterjeyaraj/exam-center/backend/pkg/models"
)

func HandleSignin(response http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	var user model.AuthenticationModel
	var result model.ResponseModel

	dec := json.NewDecoder(request.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&user)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		if errors.Is(err, io.EOF) {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("{\"message\": \"UnAthorized\"}"))
		}
		return
	}

	auth, email, fname, lname, userid := database.HandleAuthentication(user.Email, user.Password, "GO", "users")
	token, _, _ := helper.JWTTokenGenerator(email, fname, lname, userid)

	result.Token = token
	result.Expires_in = time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()

	if !auth {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte("{\"message\": \"Invalid Credentials\"}"))
		return
	} else if auth {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(&result)
	}
}
