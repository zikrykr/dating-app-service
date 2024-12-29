package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		ID              string         `json:"id" gorm:"column:id"`
		Name            string         `json:"name" gorm:"column:name"`
		Email           string         `json:"email" gorm:"column:email"`
		DateOfBirth     string         `json:"date_of_birth" gorm:"column:date_of_birth"`
		Gender          string         `json:"gender" gorm:"column:gender"`
		ProfileImageURL string         `json:"profile_image_url" gorm:"column:profile_image_url"`
		Description     string         `json:"description" gorm:"column:description"`
		Password        string         `json:"password" gorm:"column:password"`
		CreatedAt       time.Time      `json:"created_at" gorm:"column:created_at"`
		UpdatedAt       time.Time      `json:"updated_at" gorm:"column:updated_at"`
		DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
	}
)
