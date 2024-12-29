package service

import (
	"context"
	"errors"
	"testing"

	userModel "github.com/dating-app-service/internal/auth/model"
	"github.com/dating-app-service/internal/swipe/payload"
	"github.com/dating-app-service/mock"
	"github.com/golang/mock/gomock"
)

func TestSwipeService_CreateSwipe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var (
		mockAuthRepo  = mock.NewMockIAuthRepo(ctrl)
		mockSwipeRepo = mock.NewMockISwipeRepository(ctrl)
	)

	type args struct {
		ctx       context.Context
		userEmail string
		req       payload.CreateSwipeReq
	}
	tests := []struct {
		name        string
		args        args
		mockCallsFn []*gomock.Call
		wantErr     bool
	}{
		{
			name: "Successfully Create Swipe",
			args: args{
				ctx:       context.Background(),
				userEmail: "email@email.com",
				req: payload.CreateSwipeReq{
					SwipedUserID: "id",
					SwipeType:    "like",
				},
			},
			mockCallsFn: []*gomock.Call{
				mockAuthRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(userModel.User{ID: "user_id"}, nil),
				mockSwipeRepo.EXPECT().CreateSwipe(gomock.Any(), gomock.Any()).Return(nil),
			},
		},
		{
			name: "error Create Swipe",
			args: args{
				ctx:       context.Background(),
				userEmail: "email@email.com",
				req: payload.CreateSwipeReq{
					SwipedUserID: "id",
					SwipeType:    "like",
				},
			},
			mockCallsFn: []*gomock.Call{
				mockAuthRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(userModel.User{}, errors.New("internal server error")),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swipeService := &SwipeService{
				authRepository: mockAuthRepo,
				repository:     mockSwipeRepo,
			}

			err := swipeService.CreateSwipe(tt.args.ctx, tt.args.userEmail, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("premiumService.UpdateUserPremium() error=%v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
