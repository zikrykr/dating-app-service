package repository

import (
	"context"

	"github.com/dating-app-service/config/db"
	"github.com/dating-app-service/internal/auth/model"
	"github.com/dating-app-service/internal/auth/payload"
	"github.com/dating-app-service/internal/auth/port"
	"github.com/dating-app-service/pkg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type repository struct {
	db *db.GormDB
}

func NewRepository(db *db.GormDB) port.IAuthRepo {
	return repository{db: db}
}

func (r repository) GetUsers(ctx context.Context, filter payload.GetUserFilter) ([]model.User, *pkg.Pagination, error) {
	var (
		res          []model.User
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

	if filter.Email != "" {
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("email = ?", filter.Email)
		})
	}

	if filter.Gender != "" {
		fScopes = append(fScopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("gender = ?", filter.Gender)
		})
	}

	query := r.db.WithContext(ctx).Scopes(fScopes...)

	if err := query.Offset(offset).Limit(filter.Limit).Find(&res).Error; err != nil {
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

func (r repository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var res model.User

	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r repository) CreateUser(ctx context.Context, data model.User) error {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	data.Password = string(hashedPw)

	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		return err
	}

	return nil
}
