package model

type UserPremium struct {
	Email     string `json:"email"`
	IsPremium string `json:"is_premium"`
}

func (UserPremium) TableName() string {
	return "users"
}
