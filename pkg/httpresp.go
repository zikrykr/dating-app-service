package pkg

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HTTPResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

func ResponseError(c *gin.Context, code int, err error) {
	d := err.Error()

	logrus.WithContext(c.Request.Context()).Error(err)

	if code == 0 {
		code = http.StatusInternalServerError
	}

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
