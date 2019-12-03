package properties

import (
	def "./defaults"
	"encoding/xml"
)

type Device struct {
	XMLName   xml.Name  `xml:"devices"`
	Emulator  string    `xml:"emulator"`
	Disk      Disk      `xml:"disk"`
	Interface Interface `xml:"interface"`
	Input     Input     `xml:"input"`
	Graphics  Graphics  `xml:"graphics"`
	Console   Console   `xml:"console"`
}

type Input struct {
	XMLName xml.Name `xml:"input"`
	Type    string   `xml:"type,attr"`
	Bus     string   `xml:"bus,attr"`
}
type Graphics struct {
	XMLName xml.Name `xml:"graphics"`
	Type    string   `xml:"type,attr"`
	Port    string   `xml:"port,attr"`
}
type Console struct {
	XMLName xml.Name `xml:"console"`
	Type    string   `xml:"type,attr"`
}

func initDevicesWithDefaults() Device {

	devInput := Input{
		XMLName: xml.Name{},
		Type:    def.DeviceInputType,
		Bus:     def.DeviceInputBus,
	}

	devGraphics := Graphics{
		XMLName: xml.Name{},
		Type:    def.DeviceGraphicsType,
		Port:    def.DeviceGraphicsPort,
	}

	devConsole := Console{
		XMLName: xml.Name{},
		Type:    def.DeviceConsoleType,
	}

	device := Device{
		XMLName:   xml.Name{},
		Emulator:  def.DeviceEmulator,
		Disk:      InitDiskWithDefaults(),
		Interface: InitInterfaceWithDefaults(),
		Input:     devInput,
		Graphics:  devGraphics,
		Console:   devConsole,
	}

	return device
}
