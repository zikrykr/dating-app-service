package service

import (
	"context"
	"errors"

	"github.com/dating-app-service/internal/auth/payload"
	"github.com/dating-app-service/internal/auth/port"
	"github.com/dating-app-service/pkg"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	repository port.IAuthRepo
}

func NewLoginService(repo port.IAuthRepo) port.ILoginService {
	return LoginService{
		repository: repo,
	}
}

func (s LoginService) Login(ctx context.Context, req payload.LoginReq) (payload.LoginResp, error) {
	var res payload.LoginResp

	userData, err := s.repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return res, err
	}

	// compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(req.Password)); err != nil {
		return res, errors.New("invalid password")
	}

	// generate JWT
	resToken, err := pkg.GenerateJWT(&pkg.JWTClaims{
		UserID: userData.ID,
	})
	if err != nil {
		return res, err
	}

	res.AccessToken = resToken.AccessToken

	return res, nil
}
