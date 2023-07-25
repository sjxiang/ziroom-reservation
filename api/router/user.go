package router

import (
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/ziroom-reservation/api/resthandler"
)

type UserRouter interface {
	InitUserRouter(userRouter *gin.RouterGroup)
}

type UserRouterImpl struct {
	userRestHandler resthandler.UserRestHandler
}

func NewUserRouterImpl(userRestHandler resthandler.UserRestHandler) *UserRouterImpl {
	return &UserRouterImpl{userRestHandler: userRestHandler}
}

func (impl UserRouterImpl) InitUserRouter(userRouter *gin.RouterGroup) {
	
	userRouter.POST("/signin", impl.userRestHandler.SignIn)
	userRouter.POST("/signup", impl.userRestHandler.SignUp)
}
