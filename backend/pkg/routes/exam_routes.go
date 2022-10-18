package routes

import (
	controllers "github.com/chrispeterjeyaraj/exam-center/backend/pkg/controllers/exams"

	"github.com/gorilla/mux"
)

func ExamRoute(router *mux.Router) {
	router.HandleFunc("/createexam", controllers.HandleExamCreate)
	router.HandleFunc("/updateexam", controllers.HandleExamsUpdate)
	router.HandleFunc("/deleteexam/{examId}", controllers.HandleDeleteExam)
	router.HandleFunc("/getallexams", controllers.GetAllExams)
	router.HandleFunc("/getexam/{examId}", controllers.GetAnExam)
}
