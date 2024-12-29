package handler

import (
	"errors"
	"net/http"

	"github.com/dating-app-service/constants"
	"github.com/dating-app-service/internal/premium/port"
	"github.com/dating-app-service/pkg"
	"github.com/gin-gonic/gin"
)

type PremiumHandler struct {
	premiumService port.IPremiumService
}

func NewPremiumHandler(service port.IPremiumService) port.IPremiumHandler {
	return PremiumHandler{
		premiumService: service,
	}
}

func (h PremiumHandler) UpdateUserPremium(c *gin.Context) {
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

	ctx := c.Request.Context()

	err := h.premiumService.UpdateUserPremium(ctx, emailStr)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "User premium updated successfully",
	})
}
