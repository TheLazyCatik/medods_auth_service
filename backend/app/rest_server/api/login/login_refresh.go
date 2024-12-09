package login

import (
	entity "auth_service/domain/entities/user"
)

type LoginRefreshRequest struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LoginRefreshResponse struct {
	Result LoginRefreshResult `json:"result" validate:"required"`
}

type LoginRefreshResult struct {
	AccessToken entity.TokenData `json:"access_token" validate:"required"`
}
