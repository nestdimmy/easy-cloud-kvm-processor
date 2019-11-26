package properties

import (
	def "./defaults"
	"encoding/xml"
)

type Interface struct {
	XMLName       xml.Name      `xml:"interface"`
	Type          string        `xml:"type,attr"`
	SourceNetwork SourceNetwork `xml:"source"`
	//MacAddress Mac `xml:"mac"`
	Model Model `xml:"model"`
}

type SourceNetwork struct {
	XMLName xml.Name `xml:"source"`
	Network string   `xml:"network,attr"`
}
type Mac struct {
	XMLName xml.Name `xml:"mac"`
	Network string   `xml:"address,attr"`
}
type Model struct {
	XMLName xml.Name `xml:"model"`
	Network string   `xml:"type,attr"`
}

func InitInterfaceWithDefaults() Interface {

	intSource := SourceNetwork{
		XMLName: xml.Name{},
		Network: def.InterfaceSourceNetwork,
	}

	intModel := Model{
		XMLName: xml.Name{},
		Network: def.InterfaceModelType,
	}

	interf := Interface{
		XMLName:       xml.Name{},
		Type:          def.InterfaceType,
		SourceNetwork: intSource,
		Model:         intModel,
	}

	return interf
}
