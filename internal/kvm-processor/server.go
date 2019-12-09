package main

import (
	"./controllers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

const GET = "GET"
const POST = "POST"
const DELETE = "DELETE"
const PUT = "PUT"

var Logger = log.New()

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {

	BaseUrl := getServerUrl()
	Router := addRouterAndApi()

	Logger.Info("Server starting.......")
	Err := http.ListenAndServe(BaseUrl, Router)

	if Err != nil {
		Logger.Error(Err)
	}
}

func addRouterAndApi() *mux.Router {
	Router := mux.NewRouter()
	setHandleApis(Router)

	return Router
}

func getServerUrl() string {

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
		Logger.Warning("Port is not defined in .env setting port 8080...")
	}
	Logger.Info("Port to connect: " + Port)

	return ":" + Port
}

func setHandleApis(router *mux.Router) {
	const vmControllerPath = "/vm"

	router.HandleFunc(buildPath(vmControllerPath, "/test"), controllers.Test).Methods(GET)
	router.HandleFunc(buildPath(vmControllerPath, "/{searchCriteria}"), controllers.GetDomain).Methods(GET)
	router.HandleFunc(buildPath(vmControllerPath, ""), controllers.CreateDomain).Methods(POST)
	router.HandleFunc(buildPath(vmControllerPath, "/{searchCriteria}"), controllers.DeleteDomain).Methods(DELETE)
	router.HandleFunc(buildPath(vmControllerPath, "/reboot/{searchCriteria}"), controllers.RebootDomain).Methods(POST)

}

func buildPath(controllerPath string, methodPath string) string {
	const basePath string = "/api/v1"
	return basePath + controllerPath + methodPath
}
