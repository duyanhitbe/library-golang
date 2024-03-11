package validations

import (
	"github.com/google/uuid"
)

func ParseUUID(id string) (*uuid.UUID, error) {
	result, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
