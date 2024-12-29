package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dating-app-service/constants"
	"github.com/dating-app-service/mock"
	"github.com/dating-app-service/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPremiumHandler_UpdateUserPremium(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		mockPremiumService = mock.NewMockIPremiumService(ctrl)
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
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/premium", nil)
				c.Request.Header.Set("Content-Type", "application/json")
			},
			mockCallFn: func() {
				mockPremiumService.EXPECT().UpdateUserPremium(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "error",
			req: func(c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/premium", nil)
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
			h := &PremiumHandler{
				premiumService: mockPremiumService,
			}
			h.UpdateUserPremium(ctx)
			if tt.wantErr {
				assert.True(t, ctx.Writer.Status() != http.StatusOK)
				return
			}

			assert.True(t, ctx.Writer.Status() == http.StatusOK)
		})
	}

}
