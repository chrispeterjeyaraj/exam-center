package routes

import (
	controllers "github.com/chrispeterjeyaraj/exam-center/backend/pkg/controllers/exams"

	"github.com/gorilla/mux"
)

func ExamRoute(router *mux.Router) {
	router.HandleFunc("/createexam", controllers.HandleExamsCreate)
	// router.HandleFunc("/updateexam", controllers.HandleExamsUpdate)
}
