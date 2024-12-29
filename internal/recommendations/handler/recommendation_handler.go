package handler

import (
	"errors"
	"net/http"

	"github.com/dating-app-service/constants"
	"github.com/dating-app-service/internal/recommendations/payload"
	"github.com/dating-app-service/internal/recommendations/port"
	"github.com/dating-app-service/pkg"
	"github.com/gin-gonic/gin"
)

type RecommendationHandler struct {
	recommendationService port.IRecommendationService
}

func NewRecommendationHandler(service port.IRecommendationService) port.IRecommendationHandler {
	return RecommendationHandler{
		recommendationService: service,
	}
}

func (h RecommendationHandler) GetRecommendations(c *gin.Context) {
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

	resp, err := h.recommendationService.GetRecommendation(c, payload.GetRecommendationsReq{
		Email: emailStr,
	})
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Recommendations retrieved successfully",
		Data:    resp,
	})
}
