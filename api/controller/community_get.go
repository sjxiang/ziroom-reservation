package controller

import (
	"context"
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func (ctrl *Controller) GetCommunityInfo(ctx *gin.Context) {
	id := ctx.Param("id")

	community, err := ctrl.Store.Community.GetCommunityByID(context.TODO(), id)
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
		"Data": community,
	})
}