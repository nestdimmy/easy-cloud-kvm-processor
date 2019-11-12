package controllers

import (
	"../utils"
	"net/http"
)

const loggerName = "Virtual Machine Controller Logger"

var logger = utils.InitLogger(loggerName)

var Test = func(w http.ResponseWriter, r *http.Request) {
	Respond(w, Message("domainName", utils.GetInfo()))
	logger.Info("Test controller called")
}

var CreateVM = func(w http.ResponseWriter, r *http.Request) {
	Respond(w, Message("domain", "VM created"))
	logger.Info("VM created")

}

var DeleteVM = func(w http.ResponseWriter, r *http.Request) {
	Respond(w, Message("domain", "VM created"))
	logger.Info("VM deleted")
}
