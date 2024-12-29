package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dating-app-service/constants"
	"github.com/dating-app-service/internal/auth/payload"
	"github.com/dating-app-service/mock"
	"github.com/dating-app-service/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProfileHandler_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		mockProfileService = mock.NewMockIProfileService(ctrl)
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
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/auth/profile", nil)
				c.Request.Header.Set("Content-Type", "application/json")
			},
			mockCallFn: func() {
				mockProfileService.EXPECT().GetProfile(gomock.Any(), gomock.Any()).Return(payload.GetProfileResp{
					Name: "Some Name",
				}, nil)
			},
		},
		{
			name: "error",
			req: func(c *gin.Context) {
				c.Set(constants.CONTEXT_CLAIM_USER_EMAIL, "usererror@email.com")
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/auth/profile", nil)
				c.Request.Header.Set("Content-Type", "application/json")
			},
			mockCallFn: func() {
				mockProfileService.EXPECT().GetProfile(gomock.Any(), gomock.Any()).Return(payload.GetProfileResp{}, errors.New("internal server error"))
			},
			wantErr: true,
		},
		{
			name: "error - email not found on ctx",
			req: func(c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/auth/profile", nil)
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
			h := &ProfileHandler{
				profileService: mockProfileService,
			}
			h.GetProfile(ctx)
			if tt.wantErr {
				assert.True(t, ctx.Writer.Status() != http.StatusOK)
				return
			}

			assert.True(t, ctx.Writer.Status() == http.StatusOK)
		})
	}

}
