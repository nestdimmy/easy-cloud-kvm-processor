package templates

import (
	"../../models/dto"
	"../../models/properties"
	"encoding/xml"
	"github.com/google/logger"
	_ "github.com/libvirt/libvirt-go"
)

func GenerateTemplateFromRequestData(domainBody dto.DomainDto) string {

	var filledDomainBody = beautifyRequestContext(domainBody)

	output, err := xml.MarshalIndent(filledDomainBody, "  ", "    ")
	if err != nil {
		logger.Error(err)
	}

	return string(output)
}

func beautifyRequestContext(domainBody dto.DomainDto) properties.Domain {

	newDomain := properties.InitDomain()

	properties.SetDomainName(domainBody.Name, &newDomain)
	properties.SetVirtualCpus(domainBody.VirtualCPUs, &newDomain)
	properties.SetMemory(domainBody.Memory, &newDomain)
	properties.SetCurrentMemory(domainBody.CurrentMemory, &newDomain)

	return newDomain
}
