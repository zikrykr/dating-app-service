package service

import (
	"context"
	"errors"
	"testing"

	"github.com/dating-app-service/internal/auth/model"
	"github.com/dating-app-service/internal/auth/payload"
	"github.com/dating-app-service/mock"
	"github.com/golang/mock/gomock"
)

func TestSignUpService_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var (
		mockAuthRepo = mock.NewMockIAuthRepo(ctrl)
	)

	type args struct {
		ctx context.Context
		req payload.SignUpReq
	}
	tests := []struct {
		name        string
		args        args
		mockCallsFn []*gomock.Call
		wantErr     bool
	}{
		{
			name: "Successfully Sign Up",
			args: args{
				ctx: context.Background(),
				req: payload.SignUpReq{},
			},
			mockCallsFn: []*gomock.Call{
				mockAuthRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(model.User{}, nil),
				mockAuthRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil),
			},
		},
		{
			name: "error Sign Up",
			args: args{
				ctx: context.Background(),
				req: payload.SignUpReq{},
			},
			mockCallsFn: []*gomock.Call{
				mockAuthRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(model.User{}, errors.New("internal server error")),
			},
			wantErr: true,
		},
		{
			name: "error user has been registered",
			args: args{
				ctx: context.Background(),
				req: payload.SignUpReq{},
			},
			mockCallsFn: []*gomock.Call{
				mockAuthRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(model.User{ID: "user-id"}, nil),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signUpService := &SignUpService{
				repository: mockAuthRepo,
			}

			err := signUpService.SignUp(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("profileService.GetProfile() error=%v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
