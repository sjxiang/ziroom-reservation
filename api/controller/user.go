package controller

import (
	"context"
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sjxiang/ziroom-reservation/db"
	"github.com/sjxiang/ziroom-reservation/types"
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
			ctx.JSON(http.StatusNotFound, gin.H{
				"Code": -1,
				"Msg": "not found",
			})
			return 
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Code": -1,
			"Msg": "internal server error",
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
			ctx.JSON(http.StatusNotFound, gin.H{
				"Code": -1,
				"Msg": "not found",
			})
			return 
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Code": -1,
			"Msg": "internal server error",
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Msg":  "查询成功",
		"Data": users,
	})
}

