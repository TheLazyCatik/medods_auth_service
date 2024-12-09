package session

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	AccessTokenID uuid.UUID
	HashToken     string
	ExpiresAt     time.Time
}
