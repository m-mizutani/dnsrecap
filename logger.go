package main

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func setupLogger() error {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	logger = zapLogger.Sugar()
	return nil
}
