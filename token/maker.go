package token

import (
	"time"

	"github.com/duyanhitbe/library-golang/db"
	"github.com/google/uuid"
)

type TokenMaker interface {
	Generate(userId uuid.UUID, userRole db.RoleEnum, duration time.Duration) (string, *Payload, error)
	Verify(token string) (*Payload, error)
}
