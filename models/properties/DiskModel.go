package properties

import (
	def "./defaults"
	"encoding/xml"
)

type Disk struct {
	XMLName xml.Name `xml:"disk"`

	Type      string    `xml:"type,attr"`
	Device    string    `xml:"device,attr"`
	Driver    Driver    `xml:"driver"`
	ImgSource ImgSource `xml:"source"`
	TargetDev TargetDev `xml:"target"`
	Readonly  Readonly  `xml:"readonly"`
}

type Driver struct {
	XMLName xml.Name `xml:"driver"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
}

type ImgSource struct {
	XMLName xml.Name `xml:"source"`
	File    string   `xml:"file,attr"`
}

type TargetDev struct {
	XMLName xml.Name `xml:"target"`
	Dev     string   `xml:"dev,attr"`
	Bus     string   `xml:"bus,attr"`
}
type Readonly struct {
	XMLName xml.Name `xml:"readonly"`
}

func InitDiskWithDefaults() Disk {

	diskDriver := Driver{
		XMLName: xml.Name{},
		Name:    def.DiskDriverName,
		Type:    def.DiskDriverType,
	}

	diskSrc := ImgSource{
		XMLName: xml.Name{},
		File:    def.DiskSourceFile,
	}

	diskTrgDev := TargetDev{
		XMLName: xml.Name{},
		Dev:     def.DiskTargetDev,
		Bus:     def.DiskTargetBus,
	}

	diskReadOnly := Readonly{}

	disk := Disk{
		XMLName:   xml.Name{},
		Type:      def.DiskType,
		Device:    def.DiskDevice,
		Driver:    diskDriver,
		ImgSource: diskSrc,
		TargetDev: diskTrgDev,
		Readonly:  diskReadOnly,
	}

	return disk
}
