package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/chrispeterjeyaraj/exam-center/backend/db"
)

func main() {
	DB := db.Init()
	// h := handlers.New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/authuser", AuthUser).Methods("GET")
	router.HandleFunc("/questions", GetQuestions).Methods("GET")
	http.ListenAndServe(":4000", router)
}
