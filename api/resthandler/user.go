package resthandler

import (
	"github.com/gin-gonic/gin"
	"github.com/sjxiang/ziroom-reservation/internal/repository"
	"go.uber.org/zap"
)

// DTO，数据传输

type UserRestHandler interface {
	// 登录
	Login(c *gin.Context)
	// 注册
	Register(c *gin.Context)
	// 退出

	// 验证

	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	FindAllUsers(c *gin.Context)
}

type UserRestHandlerImpl struct {
	Logger              *zap.SugaredLogger
	UserRepo            repository.UserRepository
}

func NewUserRestHandlerImpl(Logger *zap.SugaredLogger, userRepo repository.UserRepository) *UserRestHandlerImpl {
	return &UserRestHandlerImpl{
		Logger:              Logger,
		UserRepo:            userRepo,
	}
}

func (impl UserRestHandlerImpl) SignIn(c *gin.Context) {
}

func (impl UserRestHandlerImpl) GetUser(c *gin.Context) {
}

func (impl UserRestHandlerImpl) UpdateUser(c *gin.Context) {
}

func (impl UserRestHandlerImpl) DeleteUser(c *gin.Context) {
}

func (impl UserRestHandlerImpl) CreateUser(c *gin.Context) {
}

func (impl UserRestHandlerImpl) FindAllUsers(c *gin.Context) {
}

// apiv1.Get("/user/:id", userHandler.HandleGetUser)
// apiv1.Put("/user/:id", userHandler.HandlePutUser)
// apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
// apiv1.Post("/user", userHandler.HandlePostUser)
// apiv1.Get("/user", userHandler.HandleGetUsers)

// FindAllResources(c *gin.Context)
// 	CreateResource(c *gin.Context)
// 	GetResource(c *gin.Context)
// 	UpdateResource(c *gin.Context)
// 	DeleteResource(c *gin.Context)
// 	TestConnection(c *gin.Context)
// 	GetMetaInfo(c *gin.Context)
// 	CreateOAuthToken(c *gin.Context)