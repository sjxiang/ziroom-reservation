package mws

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/sjxiang/ziroom-reservation/internal/db"
)

func JWTAuthentication(userStore db.UserStore) gin.HandlerFunc {
	
	return func(ctx *gin.Context) {
		authCode := ctx.GetHeader("Authorization")
		if authCode == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// SplitN 的意思是切割字符串，但是最多 N 段
		// 如果要是 N 为 0 或者负数，则是另外的含义，可以看它的文档
		authSegments := strings.SplitN(authCode, " ", 2)
		if len(authSegments) != 2 {
			// 格式不对
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := validateToken(authSegments[1])
		if err != nil {
			// 格式不对
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		expiresFloat := claims["expires"].(float64)
		expires := int64(expiresFloat)
		// Check token expiration
		if time.Now().Unix() > expires {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Msg": "token expired",
			})
			return 
		}
		userID := claims["id"].(string)
		user, err := userStore.GetUserByID(context.TODO(), userID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Msg": "unauthorized request",
			})
			return
		}

		// Set the current authenticated user to the context.
		ctx.Set("user", user)
		
		ctx.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
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
