package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


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

