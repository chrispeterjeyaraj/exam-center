package routes

import (
	controllers "github.com/chrispeterjeyaraj/exam-center/backend/pkg/controllers/questions"

	"github.com/gorilla/mux"
)

func QuestionRoute(router *mux.Router) {
	router.HandleFunc("/createquestion", controllers.HandleQuestionCreate)
	router.HandleFunc("/updatequestion", controllers.HandleQuestionsUpdate)
	router.HandleFunc("/deletequestion/{questionId}", controllers.HandleDeleteQuestion)
	router.HandleFunc("/getallquestions", controllers.GetAllQuestions)
	router.HandleFunc("/getquestion/{questionId}", controllers.GetQuestion)
}
