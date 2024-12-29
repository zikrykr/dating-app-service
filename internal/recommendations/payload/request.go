package payload

type (
	GetRecommendationsFilter struct {
		UserIDNot     string
		UserGenderNot string

		Limit  int
		Page   int
		SortBy string
	}

	GetRecommendationsReq struct {
		Email string `json:"email"`
		Limit int    `form:"limit"`
		Page  int    `form:"page"`
	}
)
