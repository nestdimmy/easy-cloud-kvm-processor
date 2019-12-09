package controllers

import (
	"../service"
	log "../utils"
	"./restOperations"
	"net/http"
)

var Test = func(w http.ResponseWriter, r *http.Request) {
	restOperations.SimpleMessage(log.GetInfo(), w)
}

var CreateDomain = func(w http.ResponseWriter, request *http.Request) {

	domainBody := restOperations.ParseCreateDomainRequestBody(request)

	createdDomain, err := service.CreateVM(domainBody)
	if err != nil {
		restOperations.SimpleMessage(err.Error(), w)
	} else {
		restOperations.ReturnVMObject(createdDomain, w)
	}
}

var DeleteDomain = func(w http.ResponseWriter, request *http.Request) {
	searchCriteria := restOperations.ParseSearchCriteria(request)
	err := service.DeleteVM(searchCriteria)
	if err != nil {
		restOperations.SimpleMessage(err.Error(), w)
	} else {
		restOperations.SimpleMessage("Domain with key -"+searchCriteria+" deleted", w)
	}
}

var RebootDomain = func(w http.ResponseWriter, request *http.Request) {
	searchCriteria := restOperations.ParseSearchCriteria(request)
	err := service.RebootDomain(searchCriteria)
	if err != nil {
		restOperations.SimpleMessage(err.Error(), w)
	}
	restOperations.SimpleMessage("Domain with key -"+searchCriteria+" rebooting...", w)
}

var GetDomain = func(w http.ResponseWriter, request *http.Request) {
	searchCriteria := restOperations.ParseSearchCriteria(request)

	domain, err := service.GetVirtualMachine(searchCriteria)

	if err != nil {
		restOperations.SimpleMessage(err.Error(), w)
	} else {
		restOperations.ReturnVMObject(domain, w)
	}
}
