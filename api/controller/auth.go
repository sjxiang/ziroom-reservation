package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/ziroom-reservation/internal/types"
	"github.com/sjxiang/ziroom-reservation/pkg/mws"
	"go.mongodb.org/mongo-driver/mongo"
)

// DTO 和 PO，分开放或者放在一起，trade off

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}


// A handler should only do:
//   - serialization of the incoming request (JSON)
//   - do some data fetching from db
//   - call some business logic
//   - return the data back the user
func (ctrl *Controller) Authenticate(ctx *gin.Context) {

	var params AuthParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctrl.logger.Infow("invalid JSON request", "err", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": -1,
			"Msg": "invalid JSON request",
		})
		return
	}

	user, err := ctrl.Store.User.GetUserByEmail(context.TODO(), params.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctrl.logger.Infow("invalid credentials", "err", err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Code": -1,
				"Msg": "invalid credentials",
			})
			return
		}

		ctrl.logger.Infow("internal server error", "err", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Code": -1,
			"Msg": "internal server error",
		})
		return
	}

	resp := AuthResponse{
		User:  user,
		Token: mws.CreateTokenFromUser(user),
	}

	ctx.JSON(http.StatusOK,gin.H{
		"Data": resp,
		"Msg":  "jwt token 下发完毕",
	})
}