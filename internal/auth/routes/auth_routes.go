package routes

import (
	"github.com/dating-app-service/internal/auth/port"
	"github.com/gin-gonic/gin"
)

type authPublicRoutes struct{}

var PublicRoutes authPublicRoutes

func (r authPublicRoutes) NewPublicRoutes(router *gin.RouterGroup, signUpHandler port.ISignUpHandler, loginHandler port.ILoginHandler) {
	// sign-up
	router.POST("/sign-up", signUpHandler.SignUp)
	// login
	router.POST("/login", loginHandler.Login)
}
