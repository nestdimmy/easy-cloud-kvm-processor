package service

import (
	"../models"
	"../models/dto"
	"../service/templates"
	"github.com/google/logger"
	"github.com/libvirt/libvirt-go"
	"github.com/libvirt/libvirt-go-xml"
)

const quemuUrl = "qemu:///system"

func CreateVM(vm dto.DomainDto) *models.VirtualMachine {

	conn, _ := libvirt.NewConnect(quemuUrl)
	template := templates.GenerateTemplateFromRequestData(vm)

	dom, err := conn.DomainDefineXMLFlags(template, 0)
	if err != nil {
		logger.Error(err)
		logger.Warning("Error while create domain with name - '" + vm.Name + "'")
	}

	err = dom.Create()
	if err != nil {
		logger.Error(err)
		logger.Warning("Error while create domain with name - '" + vm.Name + "'")
	}
	//dom.OpenGraphics() - for future open virtual machines
	domain, _ := conn.LookupDomainByName(vm.Name)
	xmldoc, _ := domain.GetXMLDesc(0)

	domainConfiguration := &libvirtxml.Domain{}
	_ = domainConfiguration.Unmarshal(xmldoc)

	logger.Info("Virtual Machine Created", domainConfiguration.UUID)

	return models.VM(domainConfiguration.ID,
		domainConfiguration.UUID,
		domainConfiguration.Name,
		domainConfiguration.Title,
		models.ExampleUser())

}

func DeleteVM(vmBody dto.DomainDto) {
	conn, _ := libvirt.NewConnect(quemuUrl)

	dom := new(libvirt.Domain)
	if len(vmBody.Name) != 0 {
		dom, _ = conn.LookupDomainByName(vmBody.Name)
	} else if vmBody.Id != 0 {
		dom, _ = conn.LookupDomainById(vmBody.Id)
	} else if len(vmBody.UUID) != 0 {
		dom, _ = conn.LookupDomainByUUID([]byte(vmBody.UUID))
	} else {
		logger.Error("Use name, id or UUID to delete VM")
	}

	destroyErr := dom.Destroy()
	if destroyErr != nil {
		logger.Error("Error while destroying domain " + vmBody.UUID)
	}
	undefinedErr := dom.Undefine()
	if undefinedErr != nil {
		logger.Error("Error while Undefying domain " + vmBody.UUID)
	}
}
