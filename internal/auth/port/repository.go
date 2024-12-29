package port

import (
	"context"

	"github.com/dating-app-service/internal/auth/model"
)

type IAuthRepo interface {
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, data model.User) error
}
