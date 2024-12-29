package port

import "github.com/gin-gonic/gin"

type ISwipeHandler interface {
	// (POST /api/v1/swipes)
	CreateSwipe(c *gin.Context)
}
