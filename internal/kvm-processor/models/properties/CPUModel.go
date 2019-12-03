package properties

import (
	def "./defaults"
	"encoding/xml"
)

type CPU struct {
	XMLName xml.Name `xml:"cpu"`

	Mode  string `xml:"mode,attr"`
	Match string `xml:"match,attr"`
	Model string `xml:"model"`
}

func InitCPUWithDefaults() CPU {
	cpu := CPU{
		XMLName: xml.Name{},
		Mode:    def.CpuMode,
		Match:   def.CpuMatch,
		Model:   def.CpuModel,
	}
	return cpu
}
