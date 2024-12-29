package handler

import (
	"errors"
	"net/http"

	"github.com/dating-app-service/constants"
	"github.com/dating-app-service/internal/swipe/payload"
	"github.com/dating-app-service/internal/swipe/port"
	"github.com/dating-app-service/pkg"
	"github.com/gin-gonic/gin"
)

type SwipeHandler struct {
	swipeService port.ISwipeService
}

func NewSwipeHandler(service port.ISwipeService) port.ISwipeHandler {
	return SwipeHandler{
		swipeService: service,
	}
}

func (h SwipeHandler) CreateSwipe(c *gin.Context) {
	email, exists := c.Get(constants.CONTEXT_CLAIM_USER_EMAIL)
	if !exists {
		pkg.ResponseError(c, http.StatusUnauthorized, errors.New("email not found in context"))
		return
	}

	// Convert email to string
	emailStr, ok := email.(string)
	if !ok {
		pkg.ResponseError(c, http.StatusUnauthorized, errors.New("email type assertion failed"))
		return
	}

	var data payload.CreateSwipeReq
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	if err := h.swipeService.CreateSwipe(ctx, emailStr, data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, pkg.HTTPResponse{
		Success: true,
		Message: "Swipe successful",
	})
}
