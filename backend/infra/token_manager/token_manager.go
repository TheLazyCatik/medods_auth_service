package token

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTTokenManager struct {
	secretKey string
}

const (
	AccessTokenExpiresAt  time.Duration = 15 * time.Minute
	RefreshTokenExpiresAt time.Duration = 24 * time.Hour
)

func NewJWTTokenManager(jwtSecretKey string) *JWTTokenManager {
	return &JWTTokenManager{
		secretKey: jwtSecretKey,
	}
}

func (tm *JWTTokenManager) GenerateAccessToken(userID uuid.UUID, ip net.IP) (*string, *AccessTokenClaims, error) {
	claims := NewAccessTokenClaims(
		userID,
		ip,
		uuid.New(),
		AccessTokenExpiresAt,
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := token.SignedString([]byte(tm.secretKey))
	if err != nil {
		return nil, nil, err
	}

	return &signedToken, claims, nil
}

func (tm *JWTTokenManager) ValidateAccessToken(tokenStr string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tm.secretKey), nil
	})

	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("error claims type")
	}
}

func (tm *JWTTokenManager) GenerateRefreshToken() (*string, *RefreshTokenClaims, error) {
	claims := NewRefreshTokenClaims(
		RefreshTokenExpiresAt,
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := token.SignedString([]byte(tm.secretKey))
	if err != nil {
		return nil, nil, err
	}

	return &signedToken, claims, nil
}

func (tm *JWTTokenManager) ValidateRefreshToken(tokenStr string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tm.secretKey), nil
	})

	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*RefreshTokenClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("error claims type")
	}
}

func (tm *JWTTokenManager) GenerateHashFromToken(tokenString string) (*string, error) {
	hash := sha512.Sum512([]byte(tokenString))
	hashStr := hex.EncodeToString(hash[:])
	return &hashStr, nil
}

func (tm *JWTTokenManager) CompareHashTokenAndToken(hashTokenStr string, tokenString string) bool {
	hash := sha512.Sum512([]byte(tokenString))
	return hashTokenStr == hex.EncodeToString(hash[:])
}
