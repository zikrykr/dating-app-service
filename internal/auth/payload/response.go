package payload

type (
	LoginResp struct {
		AccessToken string `json:"access_token"`
	}

	GetProfileResp struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Email           string `json:"email"`
		DateOfBirth     string `json:"date_of_birth"`
		Gender          string `json:"gender"` // enum: 'male', 'female'
		ProfileImageURL string `json:"profile_image_url"`
		Description     string `json:"description"`
		IsPremium       bool   `json:"is_premium"`
	}
)
