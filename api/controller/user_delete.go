package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (ctrl *Controller) DeleteUserInfo(ctx *gin.Context) {
	userID := ctx.Param("id")
	if err := ctrl.Store.User.DeleteUser(context.TODO(), userID); err != nil {
		ctrl.logger.Infow("data query exception", "err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg":  "data query exception",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Msg":  "删除成功",
		"Data": userID,
	})
}
