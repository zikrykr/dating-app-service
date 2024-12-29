package port

import (
	"context"

	"github.com/dating-app-service/internal/auth/payload"
)

type ISignUpService interface {
	SignUp(ctx context.Context, req payload.SignUpReq) error
}
