package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

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


func (ctrl *Controller) GetUserInfo(ctx *gin.Context) {
	var (
		id = ctx.Param("id")
	)
	user, err := ctrl.Store.User.GetUserByID(context.TODO(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctrl.logger.Infow("resource not found", "id", id, "err", err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"Code": -1,
				"Msg": "resource not found",
			})
			return 
		}

		ctrl.logger.Infow("data query exception", "err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg":  "data query exception",
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Msg":  "查询成功",
		"Data": user,
	})
}


func (ctrl *Controller) GetUserList(ctx *gin.Context) {
	users, err := ctrl.Store.User.GetUsers(context.TODO())
	
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctrl.logger.Infow("resource not found", "err", err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"Code": -1,
				"Msg": "resource not found",
			})
			return  
		}

		ctrl.logger.Infow("data query exception", "err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg":  "data query exception",
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Msg":  "查询成功",
		"Data": users,
	})
}

