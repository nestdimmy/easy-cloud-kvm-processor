package restOperations

import (
	"../../models/dto"
	log "../../utils"
	"encoding/json"
	"net/http"
)

func ParseVMRequestBody(req *http.Request) dto.DomainDto {
	decoder := json.NewDecoder(req.Body)

	var virtualMachine dto.DomainDto

	err := decoder.Decode(&virtualMachine)
	if err != nil {
		log.GetLogger().Error(err)
	}

	return virtualMachine
}

func ParseCreateDomainRequestBody(req *http.Request) dto.DomainDto {
	decoder := json.NewDecoder(req.Body)

	var domainDto dto.DomainDto
	err := decoder.Decode(&domainDto)
	if err != nil {
		log.GetLogger().Error(err)
	}
	return domainDto
}
