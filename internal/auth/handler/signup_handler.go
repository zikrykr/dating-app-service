package handler

import (
	"net/http"

	"github.com/dating-app-service/internal/auth/payload"
	"github.com/dating-app-service/internal/auth/port"
	"github.com/dating-app-service/pkg"
	"github.com/gin-gonic/gin"
)

type SignUpHandler struct {
	signUpService port.ISignUpService
}

func NewSignUpHandler(service port.ISignUpService) port.ISignUpHandler {
	return SignUpHandler{
		signUpService: service,
	}
}

func (h SignUpHandler) SignUp(c *gin.Context) {
	var data payload.SignUpReq
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.ResponseError(c, err)
		return
	}

	ctx := c.Request.Context()

	if err := h.signUpService.SignUp(ctx, data); err != nil {
		pkg.ResponseError(c, err)
		return
	}

	c.JSON(http.StatusCreated, pkg.HTTPResponse{
		Success: true,
		Message: "User successfully registered",
	})
}
