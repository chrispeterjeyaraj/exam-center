package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chrispeterjeyaraj/exam-center/backend/src/db"
	"github.com/chrispeterjeyaraj/exam-center/backend/src/pkg/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// Connection mongoDB with helper class
var client, err = db.ConnectDB()

func AuthUser(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var user []models.User
	// we get params with mux.
	var params = mux.Vars(r)

	username, _ := params["username"]
	password, _ := params["password"]

	filter := bson.M{"username": username, "password": password}
	collection := client.Database("examcenter").Collection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		db.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)
}
