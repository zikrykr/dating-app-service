package port

import "github.com/gin-gonic/gin"

type ISignUpHandler interface {
	// (POST /v1/auth/sign-up)
	SignUp(c *gin.Context)
}

type ILoginHandler interface {
	// (POST /api/v1/publics/auth/login)
	Login(c *gin.Context)
}
