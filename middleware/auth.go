package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dating-app-service/config"
	"github.com/dating-app-service/constants"
	"github.com/dating-app-service/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ParseTokenFromHeader(c *gin.Context) (string, error) {
	var (
		headerToken = c.Request.Header.Get("Authorization")
		splitToken  []string
	)

	splitToken = strings.Split(headerToken, "Bearer ")

	// check valid bearer token
	if len(splitToken) <= 1 {
		return "", errors.New("invalid token")
	}

	return splitToken[1], nil
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetConfig()

		secretKey := cfg.App.JWTSecret

		// token claims
		headerToken, err := ParseTokenFromHeader(c)
		if err != nil {
			pkg.ResponseError(c, http.StatusUnauthorized, err)
			return
		}

		claims := &pkg.JWTClaims{}
		token, err := jwt.ParseWithClaims(headerToken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // check signing method
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})
		// check parse token error
		if err != nil || !token.Valid {
			pkg.ResponseError(c, http.StatusUnauthorized, err)
			return
		}

		c.Set(constants.CONTEXT_CLAIM_USER_EMAIL, claims.Email)
		c.Set(constants.CONTEXT_CLAIM_USER_ID, claims.UserID)
		c.Set(constants.CONTEXT_CLAIM_KEY, claims)

		c.Next()
	}
}
