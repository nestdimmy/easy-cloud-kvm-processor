package service

import (
	"../utils"
	//"encoding/xml"
	"github.com/libvirt/libvirt-go"
	//libvirtxml "github.com/libvirt/libvirt-go-xml"
)

const loggerName = "VirtualMachineServiceLogger"

var logger = utils.InitLogger(loggerName)

func CreateVM() {

	Connection, err := libvirt.NewConnect("qemy:///system")

	if err == nil {
		logger.Error("Connection fail")
	}

	logger.Info(Connection)
}

func DeleteVM(id string) {

}
