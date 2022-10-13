package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/chrispeterjeyaraj/exam-center/backend/src/pkg/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/authuser", handlers.AuthUser).Methods("GET")
	router.HandleFunc("/questions", handlers.GetQuestions).Methods("GET")
	http.ListenAndServe(":4000", router)
}
