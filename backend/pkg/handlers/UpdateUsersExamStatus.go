package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUsersExamStatus(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")
	// we get params with mux.
	var params = mux.Vars(r)

	userid, _ := params["userid"]
	examtaken, _ := params["examtaken"]
	examtakendate, _ := params["examtakendate"]

	collection := client.Database("examcenter").Collection("users")
	id, _ := primitive.ObjectIDFromHex(userid)
	result, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"examtaken", examtaken}}},
			{"$set", bson.D{{"examtakendate", examtakendate}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode("Updated %v Document!\n")
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}
