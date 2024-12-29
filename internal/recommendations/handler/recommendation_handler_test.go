package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dating-app-service/constants"
	"github.com/dating-app-service/internal/recommendations/model"
	"github.com/dating-app-service/mock"
	"github.com/dating-app-service/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRecommendHandler_GetRecommendation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		mockRecommendationService = mock.NewMockIRecommendationService(ctrl)
	)

	tests := []struct {
		name       string
		req        func(c *gin.Context)
		mockCallFn func()
		wantErr    bool
	}{
		{
			name: "success",
			req: func(c *gin.Context) {
				c.Set(constants.CONTEXT_CLAIM_USER_EMAIL, "usersuccess@email.com")
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/recommendation", nil)
				c.Request.Header.Set("Content-Type", "application/json")
			},
			mockCallFn: func() {
				mockRecommendationService.EXPECT().GetRecommendation(gomock.Any(), gomock.Any()).Return(model.Recommendation{}, nil)
			},
		},
		{
			name: "error",
			req: func(c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/recommendation", nil)
				c.Request.Header.Set("Content-Type", "application/json")
			},
			mockCallFn: func() {},
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockCallFn()

			httpRec := httptest.NewRecorder()
			ctx := pkg.GetTestGinContext(httpRec)
			tt.req(ctx)
			h := &RecommendationHandler{
				recommendationService: mockRecommendationService,
			}
			h.GetRecommendations(ctx)
			if tt.wantErr {
				assert.True(t, ctx.Writer.Status() != http.StatusOK)
				return
			}

			assert.True(t, ctx.Writer.Status() == http.StatusOK)
		})
	}

}
