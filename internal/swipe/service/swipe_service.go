package service

import (
	"context"
	"errors"
	"strings"

	"github.com/dating-app-service/constants"
	authPort "github.com/dating-app-service/internal/auth/port"
	"github.com/dating-app-service/internal/swipe/model"
	"github.com/dating-app-service/internal/swipe/payload"
	"github.com/dating-app-service/internal/swipe/port"
)

type SwipeService struct {
	repository     port.ISwipeRepository
	authRepository authPort.IAuthRepo
}

func NewSwipeService(repo port.ISwipeRepository, authRepo authPort.IAuthRepo) port.ISwipeService {
	return SwipeService{
		repository:     repo,
		authRepository: authRepo,
	}
}

func (s SwipeService) CreateSwipe(ctx context.Context, userEmail string, req payload.CreateSwipeReq) error {
	// Get User Data
	userData, err := s.authRepository.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	if err := s.repository.CreateSwipe(ctx, model.UserSwipe{
		UserID:       userData.ID,
		SwipedUserID: req.SwipedUserID,
		SwipeType:    req.SwipeType,
	}); err != nil {
		if strings.EqualFold(err.Error(), constants.ErrDuplicateUniqueConstraint) {
			return errors.New("you have swiped this user")
		}

		return err
	}

	return nil
}
