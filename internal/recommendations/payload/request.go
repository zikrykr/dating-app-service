package payload

import "time"

type (
	GetRecommendationsFilter struct {
		UserIDNotIN   []string
		UserGenderNot string

		Limit    int
		MaxLimit int
		Page     int
		SortBy   string
	}

	GetUserRecommendationTrackerFilter struct {
		UserID      string
		TrackerDate time.Time
	}

	GetRecommendationsReq struct {
		Email string `json:"email"`
	}
)
