package routes

import (
	"github.com/dating-app-service/internal/auth/port"
	"github.com/gin-gonic/gin"
)

type (
	authPublicRoutes struct{}
	authRoutes       struct{}
)

var (
	PublicRoutes authPublicRoutes
	Routes       authRoutes
)

func (r authPublicRoutes) NewPublicRoutes(router *gin.RouterGroup, signUpHandler port.ISignUpHandler, loginHandler port.ILoginHandler) {
	// sign-up
	router.POST("/sign-up", signUpHandler.SignUp)
	// login
	router.POST("/login", loginHandler.Login)
}

func (r authRoutes) NewRoutes(router *gin.RouterGroup, profileHandler port.IProfileHandler) {
	// get profile
	router.GET("/profile", profileHandler.GetProfile)
}
