package port

import (
	"context"

	"github.com/dating-app-service/internal/recommendations/model"
	"github.com/dating-app-service/internal/recommendations/payload"
)

type IRecommendationService interface {
	GetRecommendation(ctx context.Context, req payload.GetRecommendationsReq) (model.Recommendation, error)
}
