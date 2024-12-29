package pkg

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HTTPResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Pagination struct {
	CurrentPage     int64  `json:"current_page"`
	CurrentElements int64  `json:"current_elements"`
	TotalPages      int64  `json:"total_pages"`
	TotalElements   int64  `json:"total_elements"`
	SortBy          string `json:"sort_by"`
}

func ResponseError(c *gin.Context, err error) {
	d := err.Error()

	code := http.StatusInternalServerError

	// if request cancelled
	if c.Request.Context().Err() == context.Canceled {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) || strings.Contains(err.Error(), "not found") {
		code = http.StatusNotFound
	}

	c.AbortWithStatusJSON(code, HTTPResponse{
		Success: false,
		Message: d,
	})
}
