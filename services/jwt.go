package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	SecretKey string
}

func NewTokenService(key string) TokenService {
	return TokenService{
		SecretKey: key,
	}
}

type TokenPayload struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"user_name"`
	jwt.StandardClaims
}

func (t *TokenService) Encode(
	userID int64,
	email string,
	issuer string,
	expire time.Duration,
) (string, error) {
	if issuer == "" {
		issuer = "executionlab"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenPayload{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expire).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
		},
	})

	return token.SignedString([]byte(t.SecretKey))
}

func (t *TokenService) Decode(tokenString string) (*TokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.SecretKey), nil
	})
	if err != nil {
		if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
			return nil, fmt.Errorf("expired token")
		}

		return nil, fmt.Errorf("invalid token")
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(*TokenPayload); ok {
		return claims, nil
	}

	return nil, nil
}
