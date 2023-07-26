package mws

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/sjxiang/ziroom-reservation/internal/types"
)



func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	var ErrUnAuthorized = errors.New("unauthorized request")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, ErrUnAuthorized
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse JWT token:", err)
		return nil, ErrUnAuthorized
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, ErrUnAuthorized
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrUnAuthorized
	}
	return claims, nil
}



func CreateTokenFromUser(user *types.User) string {
	now := time.Now()
	expires := now.Add(time.Hour * 4).Unix()
	claims := jwt.MapClaims{
		"id":      user.ID,
		"email":   user.Email,
		"expires": expires,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("failed to sign token with secret", err)
	}
	return tokenStr
}
