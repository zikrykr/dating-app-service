package model

import "time"

type UserRecommendationTracker struct {
	ID          int64     `json:"id" gorm:"column:id"`
	UserID      string    `json:"user_id" gorm:"column:user_id"`
	SeenUserID  string    `json:"seen_user_id" gorm:"column:seen_user_id"`
	TrackerDate time.Time `json:"tracker_date" gorm:"column:tracker_date"`
}

func (UserRecommendationTracker) TableName() string {
	return "user_recommendation_tracker"
}
