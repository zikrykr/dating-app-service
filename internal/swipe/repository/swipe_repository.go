package repository

import (
	"context"

	"github.com/dating-app-service/config/db"
	"github.com/dating-app-service/internal/swipe/model"
	"github.com/dating-app-service/internal/swipe/port"
)

type repository struct {
	db *db.GormDB
}

func NewRepository(db *db.GormDB) port.ISwipeRepository {
	return repository{db: db}
}

func (r repository) CreateSwipe(ctx context.Context, data model.UserSwipe) error {
	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		return err
	}

	return nil
}
