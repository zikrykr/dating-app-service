package service

import (
	"context"
	"errors"

	"github.com/dating-app-service/internal/auth/model"
	"github.com/dating-app-service/internal/auth/payload"
	"github.com/dating-app-service/internal/auth/port"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SignUpService struct {
	repository port.IAuthRepo
}

func NewSignUpService(repo port.IAuthRepo) port.ISignUpService {
	return SignUpService{
		repository: repo,
	}
}

func (s SignUpService) SignUp(ctx context.Context, req payload.SignUpReq) error {
	// check if user has been registered
	if err := s.checkUserRegistered(ctx, req.Email); err != nil {
		return err
	}

	id := uuid.New()

	data := model.User{
		ID:              id.String(),
		Name:            req.Name,
		Email:           req.Email,
		DateOfBirth:     req.DateOfBirth,
		Gender:          req.Gender,
		ProfileImageURL: req.ProfileImageURL,
		Description:     req.Description,
	}

	if err := s.repository.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}

func (s SignUpService) checkUserRegistered(ctx context.Context, email string) error {
	// check if email has been registered
	resEmail, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if resEmail.ID != "" {
		return errors.New("email has been registered")
	}

	return nil
}
