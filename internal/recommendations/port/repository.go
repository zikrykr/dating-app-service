package port

import (
	"context"

	"github.com/dating-app-service/internal/recommendations/model"
	"github.com/dating-app-service/internal/recommendations/payload"
	"github.com/dating-app-service/pkg"
)

type IRecommendationRepo interface {
	GetRecommendations(ctx context.Context, filter payload.GetRecommendationsFilter) ([]model.Recommendation, *pkg.Pagination, error)
}
