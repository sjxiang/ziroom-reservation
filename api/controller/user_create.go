package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/ziroom-reservation/internal/types"
)



func (ctrl *Controller) CreateUserInfo(ctx *gin.Context) {
	var params types.CreateUserParams
	
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctrl.logger.Infow("invalid JSON request", "err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "invalid JSON request",
		})
		return
	}

	user, _ := types.NewUserFromParams(params)
	insertedUser, err := ctrl.Store.User.InsertUser(context.TODO(), user)
	if err != nil {
		ctrl.logger.Infow("data query exception", "err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg":  "data query exception",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Msg":  "注册成功",
		"Data": insertedUser,
	})
}
