package payload

type SignUpReq struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	ProfileImageURL string `json:"profile_image_url"`
	Description     string `json:"description"`
	DateOfBirth     string `json:"date_of_birth"`
	Gender          string `json:"gender"`
	Password        string `json:"password" binding:"required"`
}

type GetUserFilter struct {
	Email  string `json:"email"`
	Gender string `json:"gender"`

	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	SortBy string `json:"sort_by"`
}
