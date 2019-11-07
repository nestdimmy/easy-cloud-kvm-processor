package main

import (
	"easy-cloud-kvm-processor/controllers"
	"github.com/google/logger"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

const logPath = "webapp/logs/server.log"
const GET = "GET"
const POST = "POST"
const DELETE = "DELETE"
const PUT = "PUT"

func main() {

	ServerLogger := initLogger()
	BaseUrl := getServerUrl(ServerLogger)
	Router := addRouterAndApi()

	Err := http.ListenAndServe(BaseUrl, Router)

	if Err != nil {
		ServerLogger.Error(Err)
	}
}

func addRouterAndApi() *mux.Router {
	Router := mux.NewRouter()
	setHandleApis(Router)

	return Router
}
func setHandleApis(router *mux.Router) {
	const vmControllerPath = "/vm"
	router.HandleFunc(buildPath(vmControllerPath, "/test"), controllers.Test).Methods(GET)
}

func buildPath(controllerPath string, methodPath string) string {
	const basePath string = "/api/v1"
	return basePath + controllerPath + methodPath
}

func getServerUrl(ServerLogger *logger.Logger) string {

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
		ServerLogger.Warning("Port is not defined in .env setting port 8080...")
	}
	ServerLogger.Info("Port to connect: " + Port)

	return ":" + Port
}

func initLogger() *logger.Logger {

	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()

	ServerLogger := logger.Init("LoggerExample", false, true, lf)
	defer ServerLogger.Close()

	return ServerLogger

}
