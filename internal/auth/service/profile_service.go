package service

import (
	"context"

	"github.com/dating-app-service/internal/auth/payload"
	"github.com/dating-app-service/internal/auth/port"
)

type ProfileService struct {
	repository port.IAuthRepo
}

func NewProfileService(repo port.IAuthRepo) port.IProfileService {
	return ProfileService{
		repository: repo,
	}
}

func (s ProfileService) GetProfile(ctx context.Context, userEmail string) (payload.GetProfileResp, error) {
	userData, err := s.repository.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return payload.GetProfileResp{}, err
	}

	res := payload.GetProfileResp{
		ID:              userData.ID,
		Name:            userData.Name,
		Email:           userData.Email,
		DateOfBirth:     userData.DateOfBirth,
		Gender:          userData.Gender,
		ProfileImageURL: userData.ProfileImageURL,
		Description:     userData.Description,
		IsPremium:       userData.IsPremium,
	}

	return res, nil
}
