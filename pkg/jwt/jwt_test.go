package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAccessToken(t *testing.T) {
	userID := "12345"
	tokenString, err := GenerateAccessToken(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	claims, err := ValidateToken(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, "ewallet", claims.Issuer)
	assert.WithinDuration(t, time.Now().Add(time.Hour), claims.ExpiresAt.Time, time.Minute)
}

func TestGenerateRefreshToken(t *testing.T) {
	userID := "12345"
	tokenString, err := GenerateRefreshToken(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	claims, err := ValidateToken(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, "ewallet", claims.Issuer)
	assert.WithinDuration(t, time.Now().Add(time.Hour*24*7), claims.ExpiresAt.Time, time.Minute)
}

func TestValidateToken(t *testing.T) {
	userID := "12345"
	tokenString, err := GenerateAccessToken(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	claims, err := ValidateToken(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, "ewallet", claims.Issuer)

	// Test with an invalid token
	invalidTokenString := tokenString + "invalid"
	_, err = ValidateToken(invalidTokenString)
	assert.Error(t, err)
}
