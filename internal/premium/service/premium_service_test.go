package service

import (
	"context"
	"errors"
	"testing"

	"github.com/dating-app-service/internal/auth/model"
	"github.com/dating-app-service/mock"
	"github.com/golang/mock/gomock"
)

func TestPremiumService_UpdateUserPremium(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var (
		mockAuthRepo    = mock.NewMockIAuthRepo(ctrl)
		mockPremiumRepo = mock.NewMockIPremiumRepo(ctrl)
	)

	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name        string
		args        args
		mockCallsFn []*gomock.Call
		wantErr     bool
	}{
		{
			name: "Successfully Update Premium User",
			args: args{
				ctx:   context.Background(),
				email: "email@email.com",
			},
			mockCallsFn: []*gomock.Call{
				mockAuthRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(model.User{ID: "user_id"}, nil),
				mockPremiumRepo.EXPECT().UpdateUserPremium(gomock.Any(), gomock.Any()).Return(nil),
			},
		},
		{
			name: "error Update Premium User",
			args: args{
				ctx:   context.Background(),
				email: "email@email.com",
			},
			mockCallsFn: []*gomock.Call{
				mockAuthRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(model.User{}, errors.New("internal server error")),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			premiumService := &PremiumService{
				repository: mockPremiumRepo,
				authRepo:   mockAuthRepo,
			}

			err := premiumService.UpdateUserPremium(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("premiumService.UpdateUserPremium() error=%v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
