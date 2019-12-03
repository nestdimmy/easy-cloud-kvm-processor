package main

import (
	"./controllers"
	"./utils"
	"github.com/google/logger"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

const GET = "GET"
const POST = "POST"
const DELETE = "DELETE"
const PUT = "PUT"

func main() {

	utils.InitLogger("logger")
	BaseUrl := getServerUrl()
	Router := addRouterAndApi()
	logger.Info("123")
	Err := http.ListenAndServe(BaseUrl, Router)

	if Err != nil {
		logger.Error(Err)
	}
}

func addRouterAndApi() *mux.Router {
	Router := mux.NewRouter()
	setHandleApis(Router)

	return Router
}

func setHandleApis(router *mux.Router) {
	const vmControllerPath = "/vm"
	const health = "/health"

	router.HandleFunc(buildPath(vmControllerPath, "/test"), controllers.Test).Methods(GET)
	router.HandleFunc(buildPath(health, ""), controllers.HealthCheck).Methods(GET)
	router.HandleFunc(buildPath(vmControllerPath, "/new"), controllers.CreateVMDomain).Methods(POST)
	router.HandleFunc(buildPath(vmControllerPath, "/delete"), controllers.DeleteVM).Methods(DELETE)

}

func buildPath(controllerPath string, methodPath string) string {
	const basePath string = "/api/v1"
	return basePath + controllerPath + methodPath
}

func getServerUrl() string {

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
		logger.Warning("Port is not defined in .env setting port 8080...")
	}
	logger.Info("Port to connect: " + Port)

	return ":" + Port
}
