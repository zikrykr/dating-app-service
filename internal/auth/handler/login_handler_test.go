package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dating-app-service/internal/auth/payload"
	"github.com/dating-app-service/mock"
	"github.com/dating-app-service/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		mockLoginService = mock.NewMockILoginService(ctrl)

		payloadLogin = `{
			"email": "user@email.com",
			"password": "password123"
		}`

		invalidPayloadLogin = `{
			"email": "user@email.com"
		}`
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
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBufferString(payloadLogin))
				c.Request.Header.Set("Content-Type", "application/json")
			},
			mockCallFn: func() {
				mockLoginService.EXPECT().Login(gomock.Any(), gomock.Any()).Return(payload.LoginResp{
					AccessToken: "access-token",
				}, nil)
			},
		},
		{
			name: "error",
			req: func(c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBufferString(payloadLogin))
				c.Request.Header.Set("Content-Type", "application/json")
			},
			mockCallFn: func() {
				mockLoginService.EXPECT().Login(gomock.Any(), gomock.Any()).Return(payload.LoginResp{}, errors.New("internal server error"))
			},
			wantErr: true,
		},
		{
			name: "error - validation required",
			req: func(c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBufferString(invalidPayloadLogin))
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
			h := &LoginHandler{
				loginService: mockLoginService,
			}
			h.Login(ctx)
			if tt.wantErr {
				assert.True(t, ctx.Writer.Status() != http.StatusOK)
				return
			}

			assert.True(t, ctx.Writer.Status() == http.StatusOK)
		})
	}

}
