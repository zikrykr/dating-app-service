package service

import (
	"context"
	"errors"

	authPort "github.com/dating-app-service/internal/auth/port"
	"github.com/dating-app-service/internal/premium/port"
)

type PremiumService struct {
	repository port.IPremiumRepo
	authRepo   authPort.IAuthRepo
}

func NewPremiumService(repo port.IPremiumRepo, authRepo authPort.IAuthRepo) port.IPremiumService {
	return PremiumService{
		repository: repo,
		authRepo:   authRepo,
	}
}

func (s PremiumService) UpdateUserPremium(ctx context.Context, userEmail string) error {
	userData, err := s.authRepo.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	if userData.IsPremium {
		return errors.New("user has been purchased premium")
	}

	if err := s.repository.UpdateUserPremium(ctx, userEmail); err != nil {
		return err
	}

	return nil
}
