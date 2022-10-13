package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/chrispeterjeyaraj/exam-center/backend/src/db"
	"github.com/chrispeterjeyaraj/exam-center/backend/src/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Book array
	var questions []models.Question

	// bson.M{},  we passed empty filter. So we want to get all data.
	collection := client.Database("examcenter").Collection("questions")
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		db.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var question models.Question
		// & character returns the memory address of the following variable.
		err := cur.Decode(&questions) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		questions = append(questions, question)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(questions) // encode similar to serialize process.
}
