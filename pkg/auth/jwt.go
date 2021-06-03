package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type jwtClaimsDefinition struct {
	jwt.StandardClaims
}

// AuthHandler is an endpoint for apikey-based JWT authentication
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("token"))
}

// JwtGenerateToken builds and returns a new JWT
func JwtGenerateToken() (string, error) {
	claims := jwtClaimsDefinition{
		jwt.StandardClaims{
			Issuer:    os.Getenv("JWT_ISSUER"),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

// JwtValidateToken checks a token to make sure it's good
func JwtValidateToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaimsDefinition{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false
	}

	if token.Valid {
		return true
	}

	return false
}
