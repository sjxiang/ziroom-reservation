package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/sjxiang/ziroom-reservation/api/router"
	"github.com/sjxiang/ziroom-reservation/api/mws"
)

type Config struct {
	ZIROOM_SERVER_HOST string 
	ZIROOM_SERVER_PORT string 
	ZIROOM_SERVER_MODE string 
}

func GetAppConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		ZIROOM_SERVER_HOST: os.Getenv("ZIROOM_SERVER_HOST"),
		ZIROOM_SERVER_PORT: os.Getenv("ZIROOM_SERVER_PORT"),
		ZIROOM_SERVER_MODE: os.Getenv("ZIROOM_SERVER_MODE"),
	}, nil
}

type Server struct {
	engine     *gin.Engine
	restRouter *router.RESTRouter
	logger     *zap.SugaredLogger
	cfg        *Config
}

func NewServer(cfg *Config, engine *gin.Engine, restRouter *router.RESTRouter, logger *zap.SugaredLogger) *Server {
	return &Server{
		cfg:        cfg,
		engine:     engine,
		restRouter: restRouter,
		logger:     logger,
	}
}

func (server *Server) Start() {
	server.logger.Infow("Starting server")
	
	gin.SetMode(server.cfg.ZIROOM_SERVER_MODE)

	// 全局中间件，cors
	server.engine.Use(gin.CustomRecovery(mws.CorsHandleRecovery))
	server.engine.Use(mws.Cors())

	// 经络（打通任督二脉）
	server.restRouter.InitRouter(server.engine.Group("/api"))

	err := server.engine.Run(server.cfg.ZIROOM_SERVER_HOST + ":" + server.cfg.ZIROOM_SERVER_PORT)
	if err != nil {
		server.logger.Errorw("Error in startup", "err", err)
		os.Exit(2)
	}
}
