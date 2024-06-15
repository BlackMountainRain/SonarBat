/**
 * Package util
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/15
 */

package util

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Generates a valid JWT token with given payload and expiration
func TestGenerateTokenWithPayloadAndExpiration(t *testing.T) {
	// Arrange
	secret := []byte("my-secret")
	exp := time.Hour
	payload := map[string]interface{}{"user": "testuser"}

	// Act
	token, err := GenerateToken(secret, exp, payload)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

// Handles nil payload by initializing an empty map
func TestGenerateTokenWithNilPayload(t *testing.T) {
	// Arrange
	secret := []byte("my-secret")
	exp := time.Hour
	var payload map[string]interface{}

	// Act
	token, err := GenerateToken(secret, exp, payload)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestGenerateTokenWithEmptySecret(t *testing.T) {
	// Arrange
	var secret []byte
	exp := time.Hour
	payload := map[string]interface{}{"user": "testuser"}

	// Act
	token, err := GenerateToken(secret, exp, payload)

	// Assert
	assert.ErrorIs(t, err, jwt.ErrInvalidKey)
	assert.Empty(t, token)
}
