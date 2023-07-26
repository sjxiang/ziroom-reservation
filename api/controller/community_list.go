package controller

import (
	"context"
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sjxiang/ziroom-reservation/internal/db"
)


type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

type CommunityQueryParams struct {
	db.Pagination
	Rating int
}


func (ctrl *Controller) GetCommunityInfoList(ctx *gin.Context) {
	var params CommunityQueryParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctrl.logger.Infow("invalid JSON request", "err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "invalid JSON request",
		})
		return
	}

	filter := db.Map{
		"rating": params.Rating,
	}
		
	communities, err := ctrl.Store.Community.GetCommunitys(context.TODO(), filter, &params.Pagination)
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

	resp := ResourceResp{
		Data:    communities,
		Results: len(communities),
		Page:    int(params.Page),
	}

	ctx.JSON(http.StatusOK, resp)
}
