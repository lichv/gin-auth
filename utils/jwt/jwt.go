package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Code string `json:"code"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(code, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)

	claims := Claims{
		code,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

