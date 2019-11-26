package templates

import (
	"../../models/dto"
	"../../models/properties"
	log "../../utils"
	"encoding/xml"
	_ "github.com/libvirt/libvirt-go"
)

func GenerateTemplateFromRequestData(domainBody dto.CreateDomainDto) string {

	var filledDomainBody = beautifyRequestContext(domainBody)

	output, err := xml.MarshalIndent(filledDomainBody, "  ", "    ")
	if err != nil {
		log.GetLogger().Error(err)
	}

	return string(output)
}

func beautifyRequestContext(domainBody dto.CreateDomainDto) properties.Domain {

	newDomain := properties.InitDomain()

	properties.SetDomainName(domainBody.Name, &newDomain)
	properties.SetVirtualCpus(domainBody.VirtualCPUs, &newDomain)
	properties.SetMemory(domainBody.Memory, &newDomain)
	properties.SetCurrentMemory(domainBody.CurrentMemory, &newDomain)

	return newDomain
}
