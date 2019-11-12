package main

import (
	"./controllers"
	"./utils"
	"github.com/google/logger"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

const loggerName = "ServerLogger"
const GET = "GET"
const POST = "POST"
const DELETE = "DELETE"
const PUT = "PUT"

func main() {

	ServerLogger := utils.InitLogger(loggerName)
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
	router.HandleFunc(buildPath(vmControllerPath, "/new"), controllers.CreateVM).Methods(POST)
	router.HandleFunc(buildPath(vmControllerPath, "/delete"), controllers.DeleteVM).Methods(POST)

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
