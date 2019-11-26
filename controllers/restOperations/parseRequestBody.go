package restOperations

import (
	"../../models/dto"
	log "../../utils"
	"encoding/json"
	"net/http"
)

const loggerName = "ParseLogger"

func ParseVMRequestBody(req *http.Request) dto.CreateDomainDto {
	decoder := json.NewDecoder(req.Body)

	var virtualMachine dto.CreateDomainDto

	err := decoder.Decode(&virtualMachine)
	if err != nil {
		log.GetLogger().Error(err)
	}

	return virtualMachine
}

func ParseCreateDomainRequestBody(req *http.Request) dto.CreateDomainDto {
	decoder := json.NewDecoder(req.Body)

	var domainDto dto.CreateDomainDto
	err := decoder.Decode(&domainDto)
	if err != nil {
		log.GetLogger().Error(err)
	}
	return domainDto
}
