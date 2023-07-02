package utils

import (
	"airbnb-golang/config"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	AccessToken string `json:"accessToken"`
	ExpiresAt   int64  `json:"expiresAt"`
}

type TokenPayload struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type myClaims struct {
	Payload TokenPayload `json:"payload"`
	jwt.RegisteredClaims
}

func GenerateJWT(data TokenPayload, cfg *config.Config) (*Token, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Hour * 12))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt, IssuedAt: jwt.NewNumericDate(time.Now()),
			ID: fmt.Sprintf("%d", time.Now().UnixNano()),
		},
	})
	accessToken, err := token.SignedString([]byte(cfg.App.Secret))
	if err != nil {
		return nil, err
	}
	return &Token{
		AccessToken: accessToken,
		ExpiresAt:   int64(expiresAt.Unix()),
	}, nil
}

func ValidateJWT(accessToken string, cfg *config.Config) (*TokenPayload, error) {
	token, err := jwt.ParseWithClaims(accessToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		//Don't forget to validate the argorithm is what you expect:
		// if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
		// 	return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		// }

		return []byte(cfg.App.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*myClaims)

	if !ok {
		return nil, errors.New("invalid token")
	}

	return &claims.Payload, nil

}
