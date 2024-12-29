package routes

import (
	"github.com/dating-app-service/internal/swipe/port"
	"github.com/gin-gonic/gin"
)

type swipeRoutes struct{}

var Routes swipeRoutes

func (r swipeRoutes) NewRoutes(router *gin.RouterGroup, swipeHandler port.ISwipeHandler) {
	// (POST /api/v1/swipe)
	router.POST("/", swipeHandler.CreateSwipe)
}
