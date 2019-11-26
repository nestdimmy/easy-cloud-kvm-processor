package controllers

import (
	"../service"
	log "../utils"
	"./restOperations"
	"net/http"
)

var Test = func(w http.ResponseWriter, r *http.Request) {
	restOperations.Respond(w, restOperations.Message("domainName", log.GetInfo()))
	log.GetLogger().Info("Test controller called")
}

var CreateVMDomain = func(w http.ResponseWriter, r *http.Request) {

	domainBody := restOperations.ParseCreateDomainRequestBody(r)

	restOperations.ReturnVMObject(service.CreateVM(domainBody), w)
	log.GetLogger().Info("VM created")

}

var CreateVM = func(w http.ResponseWriter, r *http.Request) {
	vmBody := restOperations.ParseVMRequestBody(r)

	restOperations.ReturnVMObject(service.CreateVM(vmBody), w)
	log.GetLogger().Info("VM created")

}

var DeleteVM = func(w http.ResponseWriter, r *http.Request) {
	vmBody := restOperations.ParseVMRequestBody(r)
	service.DeleteVM(vmBody.Name)
	log.GetLogger().Info("VM deleted")
}
