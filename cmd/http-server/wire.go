package main

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	
	"github.com/sjxiang/ziroom-reservation/api/controller"
	"github.com/sjxiang/ziroom-reservation/api/router"
	"github.com/sjxiang/ziroom-reservation/internal/db"
	"github.com/sjxiang/ziroom-reservation/pkg/logger"
)

// 依赖倒置
func Initialize() (*Server, error) {
	
	// set trial key for self-host users
	os.Setenv("ZIROOM_SECRET_KEY", "8xEMrWkBARcDDYQ")

	// init env
	cfg, err := GetAppConfig() 
	if err != nil {
		return nil, err 
	}

	// init gin
	engine := gin.New()
	
	// init log
	sugaredLogger := logger.NewSugardLogger()

	// init mongo
	mongoClient := initStorage(sugaredLogger)

	// init store
	userStore      := db.NewMongoUserStoreImpl(sugaredLogger, mongoClient)
	communityStore := db.NewMongoCommunityStoreImpl(sugaredLogger, mongoClient)
	roomStore      := db.NewMongoRoomStoreImpl(sugaredLogger, mongoClient, communityStore)
	bookingStore   := db.NewMongoBookingStoreImpl(sugaredLogger, mongoClient)

	store := db.NewStore(userStore, roomStore, communityStore, bookingStore)

	// init controller
	controller := controller.NewController(sugaredLogger, store)
	
	// init router
	router := router.NewRouter(controller)
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
