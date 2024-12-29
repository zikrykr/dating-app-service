package model

import (
	"time"

	"gorm.io/gorm"
)

type UserSwipe struct {
	ID           int64          `json:"id" gorm:"column:id"`
	UserID       string         `json:"user_id" gorm:"column:user_id"`
	SwipedUserID string         `json:"swiped_user_id"  gorm:"column:swiped_user_id"`
	SwipeType    string         `json:"swipe_type" gorm:"column:swipe_type"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func (UserSwipe) TableName() string {
	return "user_swipes"
}
