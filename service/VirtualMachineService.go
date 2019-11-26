package service

import (
	"../models"
	"../models/dto"
	"../service/templates"
	log "../utils"
	"github.com/libvirt/libvirt-go"
	"github.com/libvirt/libvirt-go-xml"
)

const quemuUrl = "qemu:///system"

func CreateVM(vm dto.CreateDomainDto) *models.VirtualMachine {

	conn, _ := libvirt.NewConnect(quemuUrl)

	template := templates.GenerateTemplateFromRequestData(vm)

	_, err := conn.DomainDefineXML(template)
	if err != nil {
		log.GetLogger().Error(err)
	}

	if err != nil {
		log.GetLogger().Error("Error while create domain with name - '" + vm.Name + "'")
	} else {
		log.GetLogger().Info("Started virtual machine with name " + vm.Name)
	}

	dom, _ := conn.LookupDomainByName(vm.Name)
	xmldoc, _ := dom.GetXMLDesc(0)

	domcfg := &libvirtxml.Domain{}
	_ = domcfg.Unmarshal(xmldoc)

	log.GetLogger().Info("Virtual Machine Created", domcfg.UUID)

	return models.NewVirtualMachine(domcfg.ID, domcfg.UUID, domcfg.Name, domcfg.Title, new(models.User))

}

func DeleteVM(name string) {
	conn, _ := libvirt.NewConnect(quemuUrl)
	dom, _ := conn.LookupDomainByName(name)
	err := dom.Undefine()
	if err != nil {
		log.GetLogger().Error("Error while destroying domain " + name)
	}
}
