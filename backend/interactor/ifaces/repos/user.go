package repository

import (
	"context"

	"auth_service/domain/entities/user"

	"github.com/google/uuid"
)

type UserRepo interface {
	CreateUser(ctx context.Context, args CreateUserArgs) (*uuid.UUID, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*user.User, error)
}

type CreateUserArgs struct {
	Email string
}
