package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	userModel "github.com/dating-app-service/internal/auth/model"
	"github.com/dating-app-service/internal/recommendations/model"
	"github.com/dating-app-service/internal/recommendations/payload"
	swipeModel "github.com/dating-app-service/internal/swipe/model"
	"github.com/dating-app-service/mock"
	"github.com/golang/mock/gomock"
)

func TestRecommendationService_GetRecommendation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var (
		mockAuthRepo           = mock.NewMockIAuthRepo(ctrl)
		mockRecommendationRepo = mock.NewMockIRecommendationRepo(ctrl)
		mockSwipeRepo          = mock.NewMockISwipeRepository(ctrl)
	)

	type args struct {
		ctx context.Context
		req payload.GetRecommendationsReq
	}
	tests := []struct {
		name        string
		args        args
		mockCallsFn []*gomock.Call
		wantRes     model.Recommendation
		wantErr     bool
	}{
		{
			name: "Successfully Get Recommendation",
			args: args{
				ctx: context.Background(),
				req: payload.GetRecommendationsReq{
					Email: "user@email.com",
				},
			},
			mockCallsFn: []*gomock.Call{
				mockAuthRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(userModel.User{ID: "user_id"}, nil),
				mockRecommendationRepo.EXPECT().GetUserRecommendationTracker(gomock.Any(), gomock.Any()).Return([]model.UserRecommendationTracker{}, nil),
				mockSwipeRepo.EXPECT().GetSwipesByUserID(gomock.Any(), gomock.Any()).Return([]swipeModel.UserSwipe{}, nil),
				mockRecommendationRepo.EXPECT().GetRecommendation(gomock.Any(), gomock.Any()).Return(model.Recommendation{}, nil),
				mockRecommendationRepo.EXPECT().CreateUserRecommendationTracker(gomock.Any(), gomock.Any()).Return(nil),
			},
		},
		{
			name: "error Get Recommendation",
			args: args{
				ctx: context.Background(),
				req: payload.GetRecommendationsReq{
					Email: "user@email.com",
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
			recommendService := &RecommendationService{
				repository:      mockRecommendationRepo,
				authRepository:  mockAuthRepo,
				swipeRepository: mockSwipeRepo,
			}

			resp, err := recommendService.GetRecommendation(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("premiumService.UpdateUserPremium() error=%v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(resp, tt.wantRes) {
				t.Errorf("profileService.GetProfile() gotRes=%v want %v", resp, tt.wantRes)
			}
		})
	}
}
