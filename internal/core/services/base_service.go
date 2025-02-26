package services

import (
	"evconn/internal/pkg/logging"

	"go.uber.org/zap"
)

type BaseService struct {
	logger *zap.Logger
}

func NewBaseService() *BaseService {
	logger, _ := logging.NewLogger(true)
	return &BaseService{
		logger: logger,
	}
}
