package port

import (
	"github.com/gin-gonic/gin"
)

type IPremiumHandler interface {
	UpdateUserPremium(c *gin.Context)
}
