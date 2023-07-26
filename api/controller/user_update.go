package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/ziroom-reservation/internal/db"
	"github.com/sjxiang/ziroom-reservation/internal/types"
)


func (ctrl *Controller) UpdateUserInfo(ctx *gin.Context) {
	var (
		params types.UpdateUserParams
		userID = ctx.Param("id")
	)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctrl.logger.Infow("invalid JSON request", "err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "invalid JSON request",
		})
		return
	}

	filter := db.Map{"_id": userID}
	
	if err := ctrl.Store.User.UpdateUser(context.TODO(), filter, params); err != nil {
		ctrl.logger.Infow("data query exception", "err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg":  "data query exception",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Msg":  "更新成功",
		"Data": userID,
	})
}
