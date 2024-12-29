package port

import (
	"context"

	"github.com/dating-app-service/internal/auth/model"
	"github.com/dating-app-service/internal/auth/payload"
	"github.com/dating-app-service/pkg"
)

type IAuthRepo interface {
	GetUsers(ctx context.Context, filter payload.GetUserFilter) ([]model.User, *pkg.Pagination, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, data model.User) error
}
