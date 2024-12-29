package port

import (
	"context"

	"github.com/dating-app-service/internal/recommendations/model"
	"github.com/dating-app-service/internal/recommendations/payload"
)

type IRecommendationRepo interface {
	GetRecommendation(ctx context.Context, filter payload.GetRecommendationsFilter) (model.Recommendation, error)
	GetUserRecommendationTracker(ctx context.Context, req payload.GetUserRecommendationTrackerFilter) ([]model.UserRecommendationTracker, error)
	CreateUserRecommendationTracker(ctx context.Context, data model.UserRecommendationTracker) error
}
