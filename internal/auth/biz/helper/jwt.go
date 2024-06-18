/**
 * Package util
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/15
 */

package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// GenerateToken generates a JWT token.
func GenerateToken(secret []byte, exp time.Duration, payload map[string]interface{}) (string, error) {
	if len(secret) == 0 {
		return "", jwt.ErrInvalidKey
	}
	if payload == nil {
		payload = make(map[string]interface{})
	}
	if _, ok := payload["exp"]; !ok {
		payload["exp"] = time.Now().Add(exp).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))

	return token.SignedString(secret)
}
