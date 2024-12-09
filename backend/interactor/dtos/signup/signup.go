package signup

import "github.com/google/uuid"

type SignupInputDTO struct {
	Email string
}

type SignupOutputDTO struct {
	UserID uuid.UUID
}
