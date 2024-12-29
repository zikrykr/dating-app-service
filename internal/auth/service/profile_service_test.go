package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/dating-app-service/internal/auth/model"
	"github.com/dating-app-service/internal/auth/payload"
	"github.com/dating-app-service/mock"
	"github.com/golang/mock/gomock"
)

func TestProfileService_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var (
		mockAuthRepo = mock.NewMockIAuthRepo(ctrl)
	)

	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name        string
		args        args
		mockCallsFn []*gomock.Call
		wantRes     payload.GetProfileResp
		wantErr     bool
	}{
		{
			name: "Successfully Get Profile",
			args: args{
				ctx:   context.Background(),
				email: "email@email.com",
			},
			mockCallsFn: []*gomock.Call{
				mockAuthRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(model.User{}, nil),
			},
			wantRes: payload.GetProfileResp{},
		},
		{
			name: "error Get Profile",
			args: args{
				ctx:   context.Background(),
				email: "email@email.com",
			},
			mockCallsFn: []*gomock.Call{
				mockAuthRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(model.User{}, errors.New("internal server error")),
			},
			wantRes: payload.GetProfileResp{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			profileService := &ProfileService{
				repository: mockAuthRepo,
			}

			gotRes, err := profileService.GetProfile(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("profileService.GetProfile() error=%v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("profileService.GetProfile() gotRes=%v want %v", gotRes, tt.wantRes)
			}
		})
	}
}
