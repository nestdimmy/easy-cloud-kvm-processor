package templates

import (
	"../../models/dto"
	"../../models/properties"
	"encoding/xml"
	_ "github.com/libvirt/libvirt-go"
	log "github.com/sirupsen/logrus"
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

func GenerateTemplateFromRequestData(domainBody dto.DomainDto) string {

	var filledDomainBody = beautifyRequestContext(domainBody)

	output, err := xml.MarshalIndent(filledDomainBody, "  ", "    ")
	if err != nil {
		Logger.Error(err)
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
