package controllers

import (
	"../service"
	log "../utils"
	"./restOperations"
	"github.com/google/logger"
	"net/http"
)

var Test = func(w http.ResponseWriter, r *http.Request) {
	restOperations.Respond(w, restOperations.Message("domainName", log.GetInfo()))
	logger.Info("Test controller called")
}

var CreateVMDomain = func(w http.ResponseWriter, r *http.Request) {

	domainBody := restOperations.ParseCreateDomainRequestBody(r)

	restOperations.ReturnVMObject(service.CreateVM(domainBody), w)
	logger.Info("VM created")

}

var HealthCheck = func(w http.ResponseWriter, r *http.Request) {
	restOperations.Message("status", "I'm ok!")

}

var DeleteVM = func(w http.ResponseWriter, r *http.Request) {
	vmBody := restOperations.ParseVMRequestBody(r)
	service.DeleteVM(vmBody)
	logger.Info("VM deleted")
}
