package routes

import (
	"github.com/chrispeterjeyaraj/exam-center/backend/pkg/controllers"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	// router.HandleFunc("/user", controllers.CreateUser()).Methods("POST")
	// router.HandleFunc("/user/{userId}", controllers.GetAUser()).Methods("GET")
	// router.HandleFunc("/user/{userId}", controllers.EditAUser()).Methods("PUT")
	// router.HandleFunc("/user/{userId}", controllers.DeleteAUser()).Methods("DELETE")
	// router.HandleFunc("/users", controllers.GetAllUser()).Methods("GET")

	// router.HandleFunc("/signup", Signup.HandleSignup)
	router.HandleFunc("/signin", controllers.HandleSignin)
	// router.HandleFunc("/emailVerify", EmailVerification.HandleEmailVerification)
	// router.HandleFunc("/forgotPass", ForgotPass.HandleForgotPass)
	// router.HandleFunc("/codeAuth", ForgotPass.HandleCodeAuth)
	// router.HandleFunc("/newPass", ForgotPass.HandleNewPassword)
	router.HandleFunc("/googleAuth", controllers.HandleGoogleAuthenticate)
	router.HandleFunc("/googleAuthCode", controllers.HandleCodeAuth)

}
