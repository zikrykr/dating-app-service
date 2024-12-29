package port

import (
	"context"

	"github.com/dating-app-service/internal/swipe/payload"
)

type ISwipeService interface {
	CreateSwipe(ctx context.Context, userEmail string, req payload.CreateSwipeReq) error
}
