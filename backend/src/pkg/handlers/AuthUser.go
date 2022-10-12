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
func () AuthUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	queryParams := request.URL.Query()
	var tweets []models.users
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	searchStage := bson.D{
		{"$search", bson.D{
			{"index", "synsearch"},
			{"text", bson.D{
				{"query", queryParams.Get("username")},
				{"path", "username"},
				{"synonyms", "slang"},
				{"query", queryParams.Get("password")},
				{"path", "password"},
				{"synonyms", "slang"},
			}},
		}},
	}
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{searchStage})
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