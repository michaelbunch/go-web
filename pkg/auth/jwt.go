package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type jwtClaimsDefinition struct {
	jwt.StandardClaims
}

var (
	issuer = "go-web"
	secret = "my-secret-key"
)

// JwtGenerateToken builds and returns a new JWT
func JwtGenerateToken() (string, error) {
	claims := jwtClaimsDefinition{
		jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(secret)
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

// JwtValidateToken checks a token to make sure it's good
func JwtValidateToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaimsDefinition{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false
	}

	if token.Valid {
		return true
	}

	return false
}
