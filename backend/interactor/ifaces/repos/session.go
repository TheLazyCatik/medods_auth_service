package repository

import (
	"context"
	"time"

	"auth_service/domain/entities/session"

	"github.com/google/uuid"
)

type SessionRepo interface {
	CreateSession(ctx context.Context, args CreateSessionArgs) (*uuid.UUID, error)
	GetSessionByUserID(ctx context.Context, userID uuid.UUID) (*session.Session, error)
	UpdateSessionByID(ctx context.Context, id uuid.UUID, args UpdateSessionArgs) error
	DeleteSessionByID(ctx context.Context, id uuid.UUID) error
}

type CreateSessionArgs struct {
	UserID        uuid.UUID
	HashToken     string
	AccessTokenID uuid.UUID
	ExpiresAt     time.Time
}

type UpdateSessionArgs struct {
	AccessTokenID uuid.UUID
	ExpiresAt     time.Time
}
