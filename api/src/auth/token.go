package auth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken creates a token for the user
func CreateToken(userID uint64) (token string, err error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userID"] = userID

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	
	return
}
