package repository

import (
	"context"

	"github.com/dating-app-service/config/db"
	"github.com/dating-app-service/internal/recommendations/model"
	"github.com/dating-app-service/internal/recommendations/payload"
	"github.com/dating-app-service/internal/recommendations/port"
	"gorm.io/gorm"
)

type repository struct {
	db *db.GormDB
}

func NewRepository(db *db.GormDB) port.IRecommendationRepo {
	return repository{db: db}
}

func (r repository) GetRecommendation(ctx context.Context, filter payload.GetRecommendationsFilter) (model.Recommendation, error) {
	var (
		res     model.Recommendation
		fScopes []func(db *gorm.DB) *gorm.DB
	)

	if len(filter.UserIDNotIN) > 0 {
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("id NOT IN (?)", filter.UserIDNotIN)
		})
	}

	if filter.UserGenderNot != "" {
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("gender != ?", filter.UserGenderNot)
		})
	}

	if filter.SortBy == "" {
		filter.SortBy = "created_at DESC"
	}

	query := r.db.WithContext(ctx).Scopes(fScopes...)

	if err := query.Order(filter.SortBy).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r repository) GetUserRecommendationTracker(ctx context.Context, filter payload.GetUserRecommendationTrackerFilter) ([]model.UserRecommendationTracker, error) {
	var (
		res     []model.UserRecommendationTracker
		fScopes []func(db *gorm.DB) *gorm.DB
	)

	if filter.UserID != "" {
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", filter.UserID)
		})
	}

	if !filter.TrackerDate.IsZero() {
		trackerDate := filter.TrackerDate.Format("2006-01-02")
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("tracker_date = ?", trackerDate)
		})
	}

	if err := r.db.WithContext(ctx).Scopes(fScopes...).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r repository) CreateUserRecommendationTracker(ctx context.Context, data model.UserRecommendationTracker) error {
	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		return err
	}
	return nil
}
