package routes

import (
	controllers "github.com/chrispeterjeyaraj/exam-center/backend/pkg/controllers/authentication"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/signup", controllers.HandleSignup)
	router.HandleFunc("/signin", controllers.HandleSignin)
	// router.HandleFunc("/emailVerify", EmailVerification.HandleEmailVerification)
	// router.HandleFunc("/forgotPass", ForgotPass.HandleForgotPass)
	// router.HandleFunc("/codeAuth", ForgotPass.HandleCodeAuth)
	// router.HandleFunc("/newPass", ForgotPass.HandleNewPassword)
	router.HandleFunc("/googleAuth", controllers.HandleGoogleAuthenticate)
	router.HandleFunc("/googleAuthCode", controllers.HandleCodeAuth)

}
