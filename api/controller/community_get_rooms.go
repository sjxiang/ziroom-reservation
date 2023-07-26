package controller

import (
	"context"
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func (ctrl *Controller) GetRoomsFromCommunity(ctx *gin.Context) {
	id := ctx.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctrl.logger.Infow("invalid id given", "id", id, "err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "invalid id given",
		})
		return 
	}

	filter := bson.M{"community_id": oid}

	rooms, err := ctrl.Store.Room.GetRooms(context.TODO(), filter)
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
		"Data": rooms,
	})
}
