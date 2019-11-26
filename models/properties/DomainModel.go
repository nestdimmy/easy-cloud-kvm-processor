package properties

import (
	def "./defaults"
	"encoding/xml"
	"github.com/google/logger"
	"math/rand"
	"strings"
)

type Domain struct {
	XMLName xml.Name `xml:"domain"`

	DomainType      string   `xml:"type,attr"`
	Name            string   `xml:"name"`
	UUID            string   `xml:"uuid"`
	Memory          int64    `xml:"memory"`
	CurrentMemory   int64    `xml:"currentMemory"`
	VirtualCPUs     int64    `xml:"vcpu"`
	Features        Features `xml:"features"`
	OperationSystem Os       `xml:"os"`
	CPU             CPU      `xml:"cpu"`
	Devices         Device   `xml:"devices"`
}

type Features struct {
	XMLName xml.Name `xml:"features"`
	ACPI    string   `xml:"acpi"`
	APIC    string   `xml:"apic"`
}

func InitDomain() Domain {
	features := Features{
		XMLName: xml.Name{},
		ACPI:    "",
		APIC:    "",
	}

	domain := Domain{
		XMLName:         xml.Name{},
		DomainType:      def.DomainType,
		Name:            "",
		UUID:            "",
		Memory:          0,
		CurrentMemory:   0,
		VirtualCPUs:     0,
		Features:        features,
		OperationSystem: InitOsWithDefaults(),
		CPU:             InitCPUWithDefaults(),
		Devices:         initDevicesWithDefaults(),
	}

	return domain
}

//TODO set userName and incremetnal number if name is null
func SetDomainName(domainName string, domain *Domain) {
	dName := strings.TrimSpace(domainName)
	if len(dName) == 0 {
		domain.Name = def.DomainDefaultName + string(rand.Int())
	} else {
		domain.Name = dName
	}
}

//TODO set inspect minimal value of memory
func SetMemory(memory int64, domain *Domain) {
	if memory < def.DomainMinMemory {
		logger.Warning("Memory can't be less than - " + string(def.DomainMinMemory))
		domain.Memory = def.DomainMinMemory
	} else {
		domain.Memory = memory
	}
}

//TODO set inspect minimal value of memory
func SetCurrentMemory(curMem int64, domain *Domain) {
	if curMem < def.DomainMinMemory {
		logger.Warning("Memory can't be less than - " + string(def.DomainMinMemory))
		domain.CurrentMemory = def.DomainMinMemory
	} else {
		domain.CurrentMemory = curMem
	}
}

func SetVirtualCpus(cpuCount int64, domain *Domain) {
	if cpuCount < def.DomainVirtualCpusDefaultCount {
		logger.Warning("Virtual cpus count can't be more than - " + string(def.DomainVirtualCpusDefaultCount))
		domain.VirtualCPUs = def.DomainVirtualCpusDefaultCount
	} else {
		domain.VirtualCPUs = cpuCount
	}
}
