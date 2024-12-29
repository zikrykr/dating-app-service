package pkg

import (
	"time"

	"github.com/dating-app-service/config"
	"github.com/golang-jwt/jwt/v4"
)

const (
	JWT_SUBJECT = "access_token"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type JWTResp struct {
	AccessToken string `json:"access_token"`
}

func GenerateJWT(claimsData *JWTClaims) (JWTResp, error) {
	config := config.GetConfig()

	issuedAt := time.Now()
	expAt := issuedAt.Add(60 * time.Minute)

	claimsData.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    config.App.Name,
		Subject:   JWT_SUBJECT,
		ExpiresAt: jwt.NewNumericDate(expAt),
		IssuedAt:  jwt.NewNumericDate(issuedAt),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claimsData)

	signedToken, err := token.SignedString([]byte(config.App.JWTSecret))
	if err != nil {
		return JWTResp{}, err
	}

	return JWTResp{
		AccessToken: signedToken,
	}, nil
}
