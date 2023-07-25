package router

import (
	// 中间件
	// "github.com/illacloud/builder-backend/pkg/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RESTRouter struct {
	logger               *zap.SugaredLogger
	Router               *gin.RouterGroup
	UserRouter           UserRouter
}

func NewRESTRouter(logger *zap.SugaredLogger, userRouter UserRouter) *RESTRouter {
	return &RESTRouter{
		logger:               logger,
		UserRouter:           userRouter,
	}
}

func (r RESTRouter) InitRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")

	userRouter := v1.Group("/user")

	// 局部中间件
	userRouter.Use()

	// 托管给 gin 生命周期
	r.UserRouter.InitUserRouter(userRouter)	
}
