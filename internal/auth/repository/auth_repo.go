package repository

import (
	"context"

	"github.com/dating-app-service/config/db"
	"github.com/dating-app-service/internal/auth/model"
	"github.com/dating-app-service/internal/auth/port"
	"golang.org/x/crypto/bcrypt"
)

type repository struct {
	db *db.GormDB
}

func NewRepository(db *db.GormDB) port.IAuthRepo {
	return repository{db: db}
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
