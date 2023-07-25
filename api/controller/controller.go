package controller

import (
	"github.com/sjxiang/ziroom-reservation/db"
	"go.uber.org/zap"
)

type Controller struct {
	logger *zap.SugaredLogger
	Store  *db.Store
}

func NewController(logger *zap.SugaredLogger, store *db.Store) *Controller {
	return &Controller{
		logger: logger,
		Store:  store,
	}
}
