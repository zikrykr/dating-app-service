package port

import (
	"context"

	"github.com/dating-app-service/internal/recommendations/model"
	"github.com/dating-app-service/internal/recommendations/payload"
	"github.com/dating-app-service/pkg"
)

type IRecommendationService interface {
	GetRecommendations(ctx context.Context, req payload.GetRecommendationsReq) ([]model.Recommendation, *pkg.Pagination, error)
}
