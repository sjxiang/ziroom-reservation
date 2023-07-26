package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"github.com/sjxiang/ziroom-reservation/internal/types"
)


func (ctrl *Controller) AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		user := ctx.MustGet("user").(*types.User)
	
		if !user.IsAdmin {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		
		ctx.Next()	
	}
}
