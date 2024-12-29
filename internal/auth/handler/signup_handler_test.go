package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dating-app-service/mock"
	"github.com/dating-app-service/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUpHandler_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		mockSignUpService = mock.NewMockISignUpService(ctrl)

		payloadSignUp = `{
			"name":     "Some Name",
			"email":    "user@email.com",
			"password": "password123"
		}`

		invalidPayload = `{
			"name": "Some Name"
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
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/auth/sign-up", bytes.NewBufferString(payloadSignUp))
				c.Request.Header.Set("Content-Type", "application/json")
			},
			mockCallFn: func() {
				mockSignUpService.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "error",
			req: func(c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/auth/sign-up", bytes.NewBufferString(payloadSignUp))
				c.Request.Header.Set("Content-Type", "application/json")
			},
			mockCallFn: func() {
				mockSignUpService.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(errors.New("internal server error"))
			},
			wantErr: true,
		},
		{
			name: "error - validation required",
			req: func(c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/auth/sign-up", bytes.NewBufferString(invalidPayload))
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
			h := &SignUpHandler{
				signUpService: mockSignUpService,
			}
			h.SignUp(ctx)
			if tt.wantErr {
				assert.True(t, ctx.Writer.Status() != http.StatusCreated)
				return
			}

			assert.True(t, ctx.Writer.Status() == http.StatusCreated)
		})
	}

}
