package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/chrispeterjeyaraj/exam-center/backend/pkg/models"
)
func GetQuestions(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var tweets []models.questions
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &tweets); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(tweets)
}