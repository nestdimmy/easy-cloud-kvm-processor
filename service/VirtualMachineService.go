package service

import (
	"../models"
	"../utils"
	"./templates"
	"github.com/libvirt/libvirt-go"
	"github.com/libvirt/libvirt-go-xml"
)

const loggerName = "VirtualMachineServiceLogger"
const quemuUrl = "qemu:///system"

var logger = utils.InitLogger(loggerName)

func CreateVM(vm models.VirtualMachine) *models.VirtualMachine {

	conn, _ := libvirt.NewConnect(quemuUrl)

	template := templates.FillVirtualMachineTemplate(vm)

	_, err := conn.DomainDefineXML(template)
	if err != nil {
		logger.Error(err)
	}

	if err != nil {
		logger.Error("Error while create domain with name - '" + vm.Name + "'")
	} else {
		logger.Info("Started virtual machine with name " + vm.Name)
	}

	dom, _ := conn.LookupDomainByName(vm.Name)
	xmldoc, _ := dom.GetXMLDesc(0)

	domcfg := &libvirtxml.Domain{}
	_ = domcfg.Unmarshal(xmldoc)

	logger.Info("Virtual Machine Created", domcfg.UUID)

	return models.NewVirtualMachine(domcfg.ID, domcfg.UUID, domcfg.Name, domcfg.Title, new(models.User))

}

func DeleteVM(id string) {

}
