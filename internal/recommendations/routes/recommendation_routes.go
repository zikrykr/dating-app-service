package routes

import (
	"github.com/dating-app-service/internal/recommendations/port"
	"github.com/gin-gonic/gin"
)

type recommendationRoutes struct{}

var Routes recommendationRoutes

func (r recommendationRoutes) NewRoutes(router *gin.RouterGroup, recommendationHandler port.IRecommendationHandler) {
	// (GET /api/v1/recommendations)
	router.GET("/", recommendationHandler.GetRecommendations)
}
