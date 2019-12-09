package service

import (
	"../models"
	"../models/dto"
	"./templates"
	"github.com/libvirt/libvirt-go"
	"github.com/libvirt/libvirt-go-xml"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

const quemuUrl = "qemu:///system"

var Logger = log.New()

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

var DomainStates = map[libvirt.DomainState]string{
	libvirt.DOMAIN_NOSTATE:     "Undefined",
	libvirt.DOMAIN_RUNNING:     "Running",
	libvirt.DOMAIN_BLOCKED:     "Blocked",
	libvirt.DOMAIN_PAUSED:      "Paused",
	libvirt.DOMAIN_SHUTDOWN:    "Shutdown",
	libvirt.DOMAIN_CRASHED:     "Crashed",
	libvirt.DOMAIN_PMSUSPENDED: "State is undefined",
	libvirt.DOMAIN_SHUTOFF:     "State is undefined",
}

func CreateVM(vm dto.DomainDto) (*models.VirtualMachine, error) {

	conn, _ := libvirt.NewConnect(quemuUrl)
	template := templates.GenerateTemplateFromRequestData(vm)

	dom, err := conn.DomainDefineXMLFlags(template, 0)
	if err != nil {
		Logger.Error(err)
		Logger.Warning("Error while create domain with name - '" + vm.Name + "'")
		return nil, err
	}
	err = dom.CreateWithFlags(libvirt.DOMAIN_START_FORCE_BOOT)
	if err != nil {
		Logger.Error(err)
		Logger.Warning("Error while create domain with name - '" + vm.Name + "'")
		return nil, err
	}
	//dom.OpenGraphics() - for future open virtual machines
	domain, _ := conn.LookupDomainByName(vm.Name)

	return expandDomainInfo(*domain), nil
}

func GetVirtualMachine(searchCriteria string) (*models.VirtualMachine, error) {
	domain, err := lookupDomain(searchCriteria)
	if err != nil {
		return nil, err
	}
	return expandDomainInfo(*domain), nil
}

func DeleteVM(searchCriteria string) error {

	dom, e := lookupDomain(searchCriteria)

	if e != nil {
		return e
	}

	destroyErr := dom.Destroy()
	if destroyErr != nil {

		//Log().Error("Error while destroying domain " + searchCriteria)
		return destroyErr
	}
	undefinedErr := dom.Undefine()
	if undefinedErr != nil {
		Logger.Error("Error while Undefying domain " + searchCriteria)
		return undefinedErr
	}

	return e
}

//helper methods

func lookupDomain(searchCriteria string) (*libvirt.Domain, error) {
	conn, _ := libvirt.NewConnect(quemuUrl)

	dom := new(libvirt.Domain)
	if len(searchCriteria) != 0 {
		dom, err := conn.LookupDomainByName(searchCriteria)
		if err == nil {
			return dom, nil
		}
		idUint64, err := strconv.ParseUint(searchCriteria, 10, 32)
		if err != nil {
			Logger.Error("Parse error")
		} else {
			dom, err = conn.LookupDomainById(uint32(idUint64))
			if err == nil {
				return dom, nil
			}
		}
		dom, err = conn.LookupDomainByUUID([]byte(searchCriteria))
		if err == nil {
			return dom, nil
		}
	} else {
		Logger.Error("Domain not found, please use name, id or UUID to search VM")
	}

	return dom, nil
}

func expandDomainInfo(domain libvirt.Domain) *models.VirtualMachine {

	xmlConfiguration, err := domain.GetXMLDesc(0)

	if err != nil {
		Logger.Error("Error while getting domain configuration")
	}

	domainConfiguration := &libvirtxml.Domain{}
	_ = domainConfiguration.Unmarshal(xmlConfiguration)

	domainState, _, _ := domain.GetState()

	return models.VM(domainConfiguration.ID,
		domainConfiguration.UUID,
		domainConfiguration.Name,
		DomainStates[domainState],
		domainConfiguration.Title,
		"012345678910")
}

func RebootDomain(searchCriteria string) error {
	domain, err := lookupDomain(searchCriteria)
	if err != nil {
		return err
	}

	err = domain.Reboot(libvirt.DOMAIN_REBOOT_DEFAULT)
	if err != nil {
		return err
	}
	return nil
}
