package controller

import (
	"go.uber.org/zap"

	"github.com/sjxiang/ziroom-reservation/internal/db"
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
