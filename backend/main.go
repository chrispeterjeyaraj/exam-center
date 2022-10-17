package main

import (
	"log"
	"net/http"

	"github.com/chrispeterjeyaraj/exam-center/backend/configs"
	"github.com/chrispeterjeyaraj/exam-center/backend/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router) //add this

	log.Fatal(http.ListenAndServe(":4000", router))
}
