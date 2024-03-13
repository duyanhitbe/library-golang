package token

import (
	"time"

	"github.com/duyanhitbe/library-golang/db"
	"github.com/google/uuid"
)

type Payload struct {
	Subject   uuid.UUID   `json:"sub,omitempty"`
	UserID    uuid.UUID   `json:"user_id,omitempty"`
	UserRole  db.RoleEnum `json:"role,omitempty"`
	ExpiresAt time.Time   `json:"exp,omitempty"`
	IssuedAt  time.Time   `json:"iat,omitempty"`
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ExpiredTokenError
	}
	return nil
}

func NewPayload(userId uuid.UUID, userRole db.RoleEnum, duration time.Duration) *Payload {
	subject := uuid.New()
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	return &Payload{
		UserID:    userId,
		UserRole:  userRole,
		Subject:   subject,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}
}
