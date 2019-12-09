package restOperations

import (
	"../../models/dto"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var Logger = log.New()

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func ParseVMRequestBody(req *http.Request) dto.DomainDto {
	decoder := json.NewDecoder(req.Body)

	var virtualMachine dto.DomainDto

	err := decoder.Decode(&virtualMachine)
	if err != nil {
		Logger.Error(err)
	}

	return virtualMachine
}

func ParseSearchCriteria(req *http.Request) string {
	searchCriteria := mux.Vars(req)["searchCriteria"]
	if len(searchCriteria) != 0 {
		return searchCriteria
	}
	Logger.Error("Bad searchCriteria - ", searchCriteria)
	return ""
}

func ParseCreateDomainRequestBody(req *http.Request) dto.DomainDto {
	decoder := json.NewDecoder(req.Body)

	var domainDto dto.DomainDto
	err := decoder.Decode(&domainDto)
	if err != nil {
		Logger.Error("Bad Domain Request Body", err)
	}
	return domainDto
}
