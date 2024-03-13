package token

import "errors"

var (
	ExpiredTokenError = errors.New("Token was expired")
	InvalidTokenError = errors.New("Invalid token")
)
