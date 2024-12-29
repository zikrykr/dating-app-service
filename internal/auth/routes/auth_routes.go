package routes

import (
	"github.com/dating-app-service/internal/auth/port"
	"github.com/gin-gonic/gin"
)

type authPublicRoutes struct{}

var PublicRoutes authPublicRoutes

func (r authPublicRoutes) NewPublicRoutes(router *gin.RouterGroup, signUpHandler port.ISignUpHandler) {
	// sign-up
	router.POST("/sign-up", signUpHandler.SignUp)
}
