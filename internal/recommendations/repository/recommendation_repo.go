package repository

import (
	"context"

	"github.com/dating-app-service/config/db"
	"github.com/dating-app-service/internal/recommendations/model"
	"github.com/dating-app-service/internal/recommendations/payload"
	"github.com/dating-app-service/internal/recommendations/port"
	"github.com/dating-app-service/pkg"
	"gorm.io/gorm"
)

type repository struct {
	db *db.GormDB
}

func NewRepository(db *db.GormDB) port.IRecommendationRepo {
	return repository{db: db}
}

func (r repository) GetRecommendations(ctx context.Context, filter payload.GetRecommendationsFilter) ([]model.Recommendation, *pkg.Pagination, error) {
	var (
		res          []model.Recommendation
		totalRecords int64
		fScopes      []func(db *gorm.DB) *gorm.DB
	)

	pagination := &pkg.Pagination{
		CurrentPage:     int64(filter.Page),
		CurrentElements: 0,
		TotalPages:      0,
		TotalElements:   0,
		SortBy:          filter.SortBy,
	}

	if filter.Page == 0 {
		filter.Page = 1
	}

	offset := (filter.Page - 1) * filter.Limit

	if filter.UserIDNot != "" {
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("id != ?", filter.UserIDNot)
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

	if err := query.Offset(offset).Limit(filter.Limit).Order(filter.SortBy).Find(&res).Error; err != nil {
		return res, nil, err
	}

	query.Count(&totalRecords)

	// Update Pagination
	totalPage := totalRecords / int64(filter.Limit)
	if totalRecords%int64(filter.Limit) > 0 || totalRecords == 0 {
		totalPage++
	}
	pagination.TotalPages = totalPage
	pagination.CurrentElements = int64(len(res))
	pagination.TotalElements = totalRecords

	return res, pagination, nil
}
