package handler

import (
	"bytes"
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

func TestSwipeHandler_CreateSwipe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		mockSwipeService = mock.NewMockISwipeService(ctrl)

		payload = `
		{
				"swiped_user_id": "f2e4df28-7a9e-4539-9782-6c9a67d8e0b6",
				"swipe_type": "pass"
		}
		`
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
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/swipe", bytes.NewBufferString(payload))
				c.Request.Header.Set("Content-Type", "application/json")
			},
			mockCallFn: func() {
				mockSwipeService.EXPECT().CreateSwipe(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "error",
			req: func(c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/swipe", bytes.NewBufferString(payload))
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
			h := &SwipeHandler{
				swipeService: mockSwipeService,
			}
			h.CreateSwipe(ctx)
			if tt.wantErr {
				assert.True(t, ctx.Writer.Status() != http.StatusCreated)
				return
			}

			assert.True(t, ctx.Writer.Status() == http.StatusCreated)
		})
	}

}
