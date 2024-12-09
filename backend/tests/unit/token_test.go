package main

import (
	token "auth_service/infra/token_manager"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateAccessToken(t *testing.T) {
	manager := token.NewJWTTokenManager("secretKey")

	userID := uuid.New()
	ip := net.ParseIP("192.168.1.1")

	signedToken, claims, err := manager.GenerateAccessToken(userID, ip)
	assert.NoError(t, err)
	assert.NotNil(t, signedToken)
	assert.NotNil(t, claims)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, ip, claims.IP)
	assert.WithinDuration(t, time.Now().Add(token.AccessTokenExpiresAt), claims.ExpiresAt.Time, time.Second*2)
}

func TestValidateAccessToken(t *testing.T) {
	manager := token.NewJWTTokenManager("secretKey")

	userID := uuid.New()
	ip := net.ParseIP("192.168.1.1")

	signedToken, expectedClaims, err := manager.GenerateAccessToken(userID, ip)
	assert.NoError(t, err)
	assert.NotNil(t, signedToken)

	claims, err := manager.ValidateAccessToken(*signedToken)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, expectedClaims.UserID, claims.UserID)
	assert.Equal(t, expectedClaims.TokenID, claims.TokenID)
	assert.Equal(t, expectedClaims.IP, claims.IP)
}

func TestGenerateRefreshToken(t *testing.T) {
	manager := token.NewJWTTokenManager("secretKey")

	userID := uuid.New()
	ip := net.ParseIP("192.168.1.1")

	_, _, err := manager.GenerateAccessToken(userID, ip)
	require.NoError(t, err)

	refreshToken, refreshClaims, err := manager.GenerateRefreshToken()

	assert.NoError(t, err)
	assert.NotNil(t, refreshToken)
	assert.NotNil(t, refreshClaims)
}

func TestValidateRefreshToken(t *testing.T) {
	manager := token.NewJWTTokenManager("secretKey")

	userID := uuid.New()
	ip := net.ParseIP("192.168.1.1")

	_, _, err := manager.GenerateAccessToken(userID, ip)
	require.NoError(t, err)
	refreshToken, _, err := manager.GenerateRefreshToken()

	assert.NoError(t, err)
	assert.NotNil(t, refreshToken)

	claims, err := manager.ValidateRefreshToken(*refreshToken)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
}

func TestEncryptAndCompareToken(t *testing.T) {
	manager := token.NewJWTTokenManager("secretKey")

	signedToken, _, err := manager.GenerateRefreshToken()
	require.NoError(t, err)

	hashToken, err := manager.GenerateHashFromToken(*signedToken)
	require.NoError(t, err)
	assert.NotNil(t, hashToken)

	isTokensEqual := manager.CompareHashTokenAndToken(*hashToken, *signedToken)

	assert.True(t, isTokensEqual)
}
