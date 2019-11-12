package utils

import (
	"github.com/google/logger"
	"os"
)

const logPath = "webapp/logs/log.log"

func InitLogger(loggerName string) *logger.Logger {

	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()

	ServerLogger := logger.Init(loggerName, false, true, lf)
	defer ServerLogger.Close()

	return ServerLogger
}
