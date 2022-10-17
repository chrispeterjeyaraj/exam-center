package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/chrispeterjeyaraj/exam-center/backend/src/db"
	"github.com/chrispeterjeyaraj/exam-center/backend/src/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Connection mongoDB with helper class
var client, err = db.ConnectDB()

func AuthUser(rw http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	//validate the request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		rw.WriteHeader(http.StatusBadRequest)
		response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}

	filter := bson.M{"Username": user.Username, "Password": user.Password}
	fmt.Println("FindOne() filter:", filter)
	collection := client.Database("examcenter").Collection("users")
	fmt.Println("Collection type:", reflect.TypeOf(collection), "\n")
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(rw).Encode(response)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	response := responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
	json.NewEncoder(rw).Encode(response)
}

// // set header.
// w.Header().Set("Content-Type", "application/json")

// var user models.User
// // we get params with mux.
// var params = mux.Vars(r)

// username, _ := params["username"]
// password, _ := params["password"]

// filter := bson.M{"Username": username, "Password": password}
// fmt.Println("FindOne() filter:", filter)
// // Declare Context type object for managing multiple API requests
// // ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

// collection := client.Database("examcenter").Collection("users")
// fmt.Println("Collection type:", reflect.TypeOf(collection), "\n")
// err := collection.FindOne(context.TODO(), filter).Decode(&user)

// if err != nil {
// 	fmt.Println("FindOne() ERROR:", err)
// 	os.Exit(1)
// } else {
// 	fmt.Println("FindOne() result:", user)
// 	fmt.Println("FindOne() Name:", user.Username)
// }

// // TODO : Need to update the user record after login, to update the last login and role if not set

// // id, _ := primitive.ObjectIDFromHex(user.ID)
// // result, err := collection.UpdateOne(
// // 	context.TODO(),
// // 	bson.M{"_id": id},
// // 	bson.D{
// // 		{"$set", bson.D{{"lastlogin", Date.now}}},
// // 		{"$set", bson.D{{"role", "user"}}},
// // 	},
// // )
// // if err != nil {
// // 	log.Fatal(err)
// // }

// json.NewEncoder(w).Encode(user)
// }
