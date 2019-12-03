package properties

import (
	def "./defaults"
	"encoding/xml"
)

type Os struct {
	XMLName xml.Name `xml:"os"`
	Type    Type     `xml:"type"`
	Boot    Boot     `xml:"boot"`
}

type Type struct {
	XMLName xml.Name `xml:"type"`
	Type    string   `xml:",chardata"`
	Arch    string   `xml:"arch,attr"`
}

type Boot struct {
	XMLName xml.Name `xml:"boot"`
	Dev     string   `xml:"dev,attr"`
}

func InitOsWithDefaults() Os {

	osType := Type{
		XMLName: xml.Name{},
		Type:    def.OsTypeDefault,
		Arch:    def.OsTypeArch,
	}

	osBoot := Boot{
		XMLName: xml.Name{},
		Dev:     def.OsBootDev,
	}

	os := Os{
		XMLName: xml.Name{},
		Type:    osType,
		Boot:    osBoot,
	}

	return os
}
