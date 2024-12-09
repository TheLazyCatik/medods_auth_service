package login

import (
	entity "auth_service/domain/entities/user"

	"github.com/google/uuid"
)

type PostLoginRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type PostLoginResponse struct {
	Result PostLoginResult `json:"result" validate:"required"`
}

type PostLoginResult struct {
	AccessToken  entity.TokenData `json:"access_token" validate:"required"`
	RefreshToken entity.TokenData `json:"refresh_token" validate:"required"`
}
