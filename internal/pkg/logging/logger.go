package logging

import "go.uber.org/zap"

func NewLogger(isDevelopment bool) (*zap.Logger, error) {
	var config zap.Config
	if isDevelopment {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}
	return config.Build()
}
