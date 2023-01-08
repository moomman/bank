package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const JWTSecretSize = 32

type JwtMaker struct {
	secret string
}

func NewJwtMaker(secret string) (*JwtMaker, error) {
	if len(secret) < JWTSecretSize {
		return nil, fmt.Errorf("secret is too short")
	}

	return &JwtMaker{
		secret: secret,
	}, nil
}

func (j *JwtMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payLoad := NewPayLoad(username, duration)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, payLoad)

	return claims.SignedString([]byte(j.secret))
}

func (j *JwtMaker) VerifyToken(token string) (*PayLoad, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, ErrTokenInvalid
		}
		return []byte(j.secret), nil
	}
	claims, err := jwt.ParseWithClaims(token, &PayLoad{}, keyFunc)
	if err != nil {
		vErr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(vErr.Inner, ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}

	payload, ok := claims.Claims.(*PayLoad)
	if ok {
		return payload, nil
	}
	return nil, ErrTokenInvalid
}
