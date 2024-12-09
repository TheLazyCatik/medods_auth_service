package token

import (
	"net"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccessTokenClaims struct {
	UserID  uuid.UUID `json:"user_id" validate:"required"`
	IP      net.IP    `json:"ip" validate:"required"`
	TokenID uuid.UUID `json:"token_id" validate:"required"`
	jwt.RegisteredClaims
}

func NewAccessTokenClaims(
	userID uuid.UUID,
	ip net.IP,
	tokenID uuid.UUID,
	expiresDuration time.Duration,
) *AccessTokenClaims {
	return &AccessTokenClaims{
		UserID:  userID,
		IP:      ip,
		TokenID: tokenID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresDuration).UTC()),
		},
	}
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
}

func NewRefreshTokenClaims(
	expiresDuration time.Duration,
) *RefreshTokenClaims {
	return &RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresDuration).UTC()),
		},
	}
}
