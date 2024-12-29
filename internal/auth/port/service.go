package port

import (
	"context"

	"github.com/dating-app-service/internal/auth/payload"
)

type ISignUpService interface {
	SignUp(ctx context.Context, req payload.SignUpReq) error
}

type ILoginService interface {
	Login(ctx context.Context, req payload.LoginReq) (payload.LoginResp, error)
}

type IProfileService interface {
	GetProfile(ctx context.Context, userEmail string) (payload.GetProfileResp, error)
}
