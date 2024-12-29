package port

import (
	"context"

	"github.com/dating-app-service/internal/swipe/model"
)

type ISwipeRepository interface {
	CreateSwipe(ctx context.Context, data model.UserSwipe) error
}
