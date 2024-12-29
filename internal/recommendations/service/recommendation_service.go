package service

import (
	"context"

	authPort "github.com/dating-app-service/internal/auth/port"
	"github.com/dating-app-service/internal/recommendations/model"
	"github.com/dating-app-service/internal/recommendations/payload"
	"github.com/dating-app-service/internal/recommendations/port"
	"github.com/dating-app-service/pkg"
)

type RecommendationService struct {
	repository     port.IRecommendationRepo
	authRepository authPort.IAuthRepo
}

func NewRecommendationService(repo port.IRecommendationRepo, authRepo authPort.IAuthRepo) port.IRecommendationService {
	return RecommendationService{
		repository:     repo,
		authRepository: authRepo,
	}
}

func (s RecommendationService) GetRecommendations(ctx context.Context, req payload.GetRecommendationsReq) ([]model.Recommendation, *pkg.Pagination, error) {
	userData, err := s.authRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return []model.Recommendation{}, nil, err
	}

	recommendations, pagination, err := s.repository.GetRecommendations(ctx, payload.GetRecommendationsFilter{
		UserIDNot:     userData.ID,
		UserGenderNot: userData.Gender,

		Limit:  pkg.ValidateLimit(req.Limit),
		Page:   pkg.ValidatePage(req.Page),
		SortBy: "date_of_birth DESC",
	})
	if err != nil {
		return []model.Recommendation{}, nil, err
	}

	return recommendations, pagination, nil
}
