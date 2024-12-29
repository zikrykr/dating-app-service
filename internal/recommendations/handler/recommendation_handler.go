package handler

import (
	"net/http"

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
	var req payload.GetRecommendationsReq

	// TODO: retrieve email from JWT
	req.Email = "alice.johnson@example.com"

	resp, err := h.recommendationService.GetRecommendation(c, req)
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
