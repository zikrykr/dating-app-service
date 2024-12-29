package model

type Recommendation struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	ProfileImageURL string `json:"profile_image_url"`
	DateOfBirth     string `json:"date_of_birth"`
	Gender          string `json:"gender"`
	Description     string `json:"description"`
}

func (Recommendation) TableName() string {
	return "users"
}
