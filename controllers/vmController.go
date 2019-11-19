package controllers

import (
	"../service"
	"../utils"
	"./restOperations"
	"net/http"
)

const loggerName = "Virtual Machine Controller Logger"

var logger = utils.InitLogger(loggerName)

var Test = func(w http.ResponseWriter, r *http.Request) {
	restOperations.Respond(w, restOperations.Message("domainName", utils.GetInfo()))
	logger.Info("Test controller called")
}

var CreateVM = func(w http.ResponseWriter, r *http.Request) {
	vmBody := restOperations.ParseVMRequestBody(r)

	restOperations.ReturnVMObject(service.CreateVM(vmBody), w)
	logger.Info("VM created")

}

var DeleteVM = func(w http.ResponseWriter, r *http.Request) {
	restOperations.Respond(w, restOperations.Message("domain", "VM created"))
	logger.Info("VM deleted")
}
