package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var errInvalidToken = errors.New("invalid Token")

// CreateToken creates a token for the user
func CreateToken(userID uint64) (token string, err error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userID"] = userID

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(config.SecretKey))

	return
}

// ValidateToken checks if the token is valid
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationkey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errInvalidToken
}

// ExtracteUserID returns the user id present on token
func ExtracteUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationkey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errInvalidToken
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnVerificationkey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signature method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
