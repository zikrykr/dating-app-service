package service

import (
	"context"
	"errors"
	"time"

	"github.com/dating-app-service/constants"
	authPort "github.com/dating-app-service/internal/auth/port"
	"github.com/dating-app-service/internal/recommendations/model"
	"github.com/dating-app-service/internal/recommendations/payload"
	"github.com/dating-app-service/internal/recommendations/port"
	swipePort "github.com/dating-app-service/internal/swipe/port"
	"gorm.io/gorm"
)

type RecommendationService struct {
	repository      port.IRecommendationRepo
	authRepository  authPort.IAuthRepo
	swipeRepository swipePort.ISwipeRepository
}

func NewRecommendationService(repo port.IRecommendationRepo, authRepo authPort.IAuthRepo, swipeRepo swipePort.ISwipeRepository) port.IRecommendationService {
	return RecommendationService{
		repository:      repo,
		authRepository:  authRepo,
		swipeRepository: swipeRepo,
	}
}

func (s RecommendationService) GetRecommendation(ctx context.Context, req payload.GetRecommendationsReq) (model.Recommendation, error) {
	// Get User Data
	userData, err := s.authRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return model.Recommendation{}, err
	}

	// Get User Recommendation Tracker
	// to check how much user recommendation retrieved and exclude seen on today
	currDate := time.Now()
	trackerData, err := s.repository.GetUserRecommendationTracker(ctx, payload.GetUserRecommendationTrackerFilter{
		UserID:      userData.ID,
		TrackerDate: currDate,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Recommendation{}, err
	}

	if !userData.IsPremium {
		if len(trackerData) >= constants.MAX_LIMIT_FREE_USERS {
			return model.Recommendation{}, errors.New("max limit recommendations has been reached for today")
		}
	}

	userIDNotIN := []string{
		userData.ID,
	}

	if len(trackerData) > 0 {
		for _, tracker := range trackerData {
			userIDNotIN = append(userIDNotIN, tracker.SeenUserID)
		}
	}

	// Get Swiped User
	// exclude swiped user
	swipedUser, err := s.swipeRepository.GetSwipesByUserID(ctx, userData.ID)
	if err != nil {
		return model.Recommendation{}, err
	}

	if len(swipedUser) > 0 {
		for _, user := range swipedUser {
			userIDNotIN = append(userIDNotIN, user.SwipedUserID)
		}
	}

	recommendations, err := s.repository.GetRecommendation(ctx, payload.GetRecommendationsFilter{
		UserIDNotIN:   userIDNotIN,
		UserGenderNot: userData.Gender,
		SortBy:        "date_of_birth DESC",
	})
	if err != nil {
		return model.Recommendation{}, err
	}

	// Create / Update Recommendation Tracker
	if err := s.repository.CreateUserRecommendationTracker(ctx, model.UserRecommendationTracker{
		UserID:      userData.ID,
		SeenUserID:  recommendations.ID,
		TrackerDate: currDate,
	}); err != nil {
		return model.Recommendation{}, err
	}

	return recommendations, nil
}
