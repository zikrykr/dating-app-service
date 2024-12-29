package repository

import (
	"context"

	"github.com/dating-app-service/config/db"
	"github.com/dating-app-service/internal/premium/model"
	"github.com/dating-app-service/internal/premium/port"
)

type repository struct {
	db *db.GormDB
}

func NewRepository(db *db.GormDB) port.IPremiumRepo {
	return repository{db: db}
}

func (r repository) UpdateUserPremium(ctx context.Context, userEmail string) error {
	if err := r.db.WithContext(ctx).Model(&model.UserPremium{}).Where("email = ?", userEmail).Update("is_premium", true).Error; err != nil {
		return err
	}

	return nil
}
