package main

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/ziroom-reservation/api/resthandler"
	"github.com/sjxiang/ziroom-reservation/api/router"
	"github.com/sjxiang/ziroom-reservation/internal/repository"
	"github.com/sjxiang/ziroom-reservation/pkg/util"
)


func Initialize() (*Server, error) {
	// 依赖倒置

	// init
	cfg, err := GetAppConfig() 
	if err != nil {
		return nil, err 
	}
	engine := gin.New()
	sugaredLogger := util.NewSugardLogger()

	// init storge
	mongoClient := initStorage(sugaredLogger)
	userRepo := repository.NewUserRepositoryImpl(sugaredLogger, mongoClient)
	
	// init handler
	userHandler := resthandler.NewUserRestHandlerImpl(sugaredLogger, userRepo)
	userRouter := router.NewUserRouterImpl(userHandler)
	
	// init router
	router := router.NewRESTRouter(sugaredLogger, userRouter)
	
	// init server
	server := NewServer(cfg, engine, router, sugaredLogger)
	
	// 有骨骼但无经络
	return server, nil
}


func initStorage(logger *zap.SugaredLogger) *mongo.Client {

	var (
		mongoEndpoint = os.Getenv("MONGO_DB_URL")
		mongoUsername = os.Getenv("MONGO_DB_USERNAME")
		mongoPassword = os.Getenv("MONGO_DB_PASSWORD")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: mongoUsername, 
		Password: mongoPassword,
	}).ApplyURI(mongoEndpoint))
	if err != nil {
		logger.Errorw("Error in startup, storage init failed.", "err", err)
	}

	return client
}

