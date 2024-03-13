package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/duyanhitbe/library-golang/db"
	"github.com/google/uuid"
)

type JWTMaker struct {
	secret string
}

func NewJWTMaker(secret string) TokenMaker {
	return &JWTMaker{secret: secret}
}

func (maker *JWTMaker) Generate(userId uuid.UUID, userRole db.RoleEnum, duration time.Duration) (string, *Payload, error) {
	payload := NewPayload(userId, userRole, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secret))
	return token, payload, err
}

func (maker *JWTMaker) Verify(token string) (*Payload, error) {
	var keyFunc jwt.Keyfunc = func(j *jwt.Token) (interface{}, error) {
		_, ok := j.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, InvalidTokenError
		}
		return []byte(maker.secret), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	//Check which error was occurred
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ExpiredTokenError) {
			return nil, ExpiredTokenError
		}
		return nil, InvalidTokenError
	}

	//Parse claims to payload
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, InvalidTokenError
	}

	return payload, nil
}
