package routes

import (
	"github.com/dating-app-service/internal/premium/port"
	"github.com/gin-gonic/gin"
)

type (
	premiumRoutes struct{}
)

var (
	Routes premiumRoutes
)

func (r premiumRoutes) NewRoutes(router *gin.RouterGroup, PremiumHandler port.IPremiumHandler) {
	// update premium
	router.POST("/", PremiumHandler.UpdateUserPremium)
}
