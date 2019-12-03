package utils

import (
	"github.com/google/logger"
	"os"
)

const logPath = "/home/dmitry/GolandProjects/kvm-processor/webapp/logs/log.log"

var ServerLogger = InitLogger("App log")

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

func GetLogger() *logger.Logger {
	return ServerLogger
}
