package port

import "github.com/gin-gonic/gin"

type ISignUpHandler interface {
	// (POST /v1/auth/sign-up)
	SignUp(c *gin.Context)
}
