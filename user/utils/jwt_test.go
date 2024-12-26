package utils

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	uid := uint64(12345)
	username := "testuser"

	tokenStr, err := GenerateTokn(uid, username)
	assert.NoError(t, err, "Error generating token")
	assert.NotEmpty(t, tokenStr, "Generated token should not be empty")
}

func TestParseToken(t *testing.T) {
	uid := uint64(12345)
	username := "testuser"

	// Generate a token
	tokenStr, err := GenerateTokn(uid, username)
	assert.NoError(t, err, "Error generating token")
	t.Logf("Generated token: %s", tokenStr)

	// Parse the token
	claims, err := ParseToken(tokenStr)
	if err != nil {
		t.Errorf("Error parsing token: %v", err)
	}
	assert.NoError(t, err, "Error parsing token")
	assert.NotNil(t, claims, "Claims should not be nil")
	if claims != nil {
		assert.Equal(t, uid, claims.UID, "UID should match")
		assert.Equal(t, username, claims.Username, "Username should match")
	}
}

func TestParseInvalidToken(t *testing.T) {
	invalidToken := "invalid.token.string"

	claims, err := ParseToken(invalidToken)
	assert.Error(t, err, "Expected error for invalid token")
	assert.Nil(t, claims, "Claims should be nil for invalid token")
}

func TestTokenExpiration(t *testing.T) {
	uid := uint64(12345)
	username := "testuser"

	// Generate a token with a short expiration time
	claims := MyClaims{
		UID:      uid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "ShortVideo_Server",
			Subject:   "yamon",
			Audience:  "sv_user",
			ExpiresAt: time.Now().Add(1 * time.Second).Unix(), // 1 second expiration
			NotBefore: time.Now().Add(time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(SignKey))
	assert.NoError(t, err, "Error generating token")

	// Wait for the token to expire
	time.Sleep(2 * time.Second)

	// Parse the expired token
	parsedClaims, err := ParseToken(tokenStr)
	assert.Error(t, err, "Expected error for expired token")
	assert.Nil(t, parsedClaims, "Claims should be nil for expired token")
}
