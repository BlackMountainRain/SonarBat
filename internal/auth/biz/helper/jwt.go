/**
 * Package util
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/15
 */

package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims

	UID       string `json:"uid"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Status    bool   `json:"status"`
	AvatarURL string `json:"avatar_url"`
}

// GenerateToken generates a JWT token.
func GenerateToken(secret []byte, claims Claims,
	issuer string, exp time.Duration) (string, error) {
	if len(secret) == 0 {
		return "", jwt.ErrInvalidKey
	}

	now := time.Now()

	claims.Issuer = issuer
	claims.Subject = claims.UID
	claims.ExpiresAt = jwt.NewNumericDate(now.Add(exp))
	claims.NotBefore = jwt.NewNumericDate(now)
	claims.IssuedAt = jwt.NewNumericDate(now)
	claims.ID = uuid.New().String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}
