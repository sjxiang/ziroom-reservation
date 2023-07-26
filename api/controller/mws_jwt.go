package controller

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sjxiang/ziroom-reservation/pkg/mws"
)



func (ctrl *Controller) JWTAuthentication() gin.HandlerFunc {
	
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("X-Api-Token")
		if tokenStr == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := mws.ValidateToken(tokenStr)
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
				"Code": -1,
				"Msg": "token expired",
			})
			return 
		}

		userID := claims["id"].(string)
		user, err := ctrl.Store.User.GetUserByID(context.TODO(), userID)
		if err != nil {
			// 查无此人
			if errors.Is(err, mongo.ErrNoDocuments) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"Msg": "unauthorized request",
				})
				return 
			}

			// 查询异常 / 服务器内部故障 （2 选 1）
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Code": -1,
				"Msg": "internal server error",
			})
			// ctx.JSON(http.StatusBadRequest, gin.H{
			// 	"Code": -1,
			// 	"Msg":  "data query exception",
			// })
			return 
		}

		// Set the current authenticated user to the context.
		ctx.Set("user", user)
		
		ctx.Next()
	}
}
